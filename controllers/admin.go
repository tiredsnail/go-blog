package controllers

import (
	"net/http"
	"www/bwy"
	"fmt"
	"errors"
	"io"
	"crypto/rand"
	"encoding/base64"
	"crypto/md5"
	"encoding/hex"
	"time"
)
type AdminData struct {
	Title		string
	Prompt		string
}
type checkloginstruct struct {
	Maxage 	int64			//过期时间
	Token	string			//token
}
var checklogindata checkloginstruct
func checklogin(w http.ResponseWriter,r *http.Request) (errs error) {
	cookie, err := r.Cookie("checklogindata")
	//fmt.Println("cookie:",cookie.Value)
	//fmt.Println("checklogindata:",checklogindata)

	if err != nil {
		//fmt.Fprintln(w, "Domain:", cookie.Domain)
		//fmt.Fprintln(w, "Expires:", cookie.Expires)
		//fmt.Fprintln(w, "Name:", cookie.Name)
		//fmt.Fprintln(w, "Value:", cookie.Value)
		//返回登录页面

		//panic(fmt.Sprintf("invalid suit %v"))
		return errors.New("cookie读取失败")

	}
	//if checklogindata == nil {
	//	return errors.New("token不存在")
	//}
	//查找
	if checklogindata.Token != cookie.Value || cookie.Value == ""{
		return errors.New("token不匹配")
	}
	//判断 token 是否过期
	if time.Now().Unix() > checklogindata.Maxage {
		return errors.New("token已失效")
	}
	//刷新 token 过期时间
	checklogindata.Maxage = time.Now().Unix() + 3600
	return nil
}


//后台登录
func Admin_Login(w http.ResponseWriter,r *http.Request) {
	rd := AdminData {
		Title: "登录 - 后台管理",
	}
	if r.Method == "POST" {
		number := r.PostFormValue("bwy_number")
		password := r.PostFormValue("bwy_password")
		if number == "admin" && password == "123456" {
			//生成唯一id
			unique := make([]byte, 48)
			if _, err := io.ReadFull(rand.Reader, unique); err == nil {
				uniqueid := base64.URLEncoding.EncodeToString(unique)
				h := md5.New()
				h.Write([]byte(uniqueid))
				uniqueid = hex.EncodeToString(h.Sum(nil))
				//设置 cookie
				cookie := http.Cookie{Name: "checklogindata", Value: uniqueid, Path: "/", MaxAge: 86400}
				http.SetCookie(w, &cookie)
				//设置 session
				checklogindata.Token = uniqueid
				checklogindata.Maxage = time.Now().Unix() + 3600
				//跳转后台
				w.Header().Set("Location", "/admin")
				w.WriteHeader(302)
				return
			}
			rd.Prompt = "设置cookie错误"
		} else {
			rd.Prompt = "账号信息错误"
		}

	}

	MyTemplate := bwy.InitTemplate()
	//模板
	MyTemplate.ParseFiles("./views/admin/login.html", "./views/common/_admin_login.html")
	MyTemplate.ExecuteTemplate(w, "login", rd)
}

//首页
func Admin_Index(w http.ResponseWriter,r *http.Request) {
	if err := checklogin(w ,r); err != nil {
		fmt.Println(err)
		w.Header().Set("Location", "/admin/login")
		w.WriteHeader(302)
		return
	}
	//return
	rd := RetData{
		Title: "后台管理",
		//Archive: models.Archive(),
	}
	MyTemplate := bwy.InitTemplate()
	//模板
	MyTemplate.ParseFiles("./views/admin/index.html", "./views/common/_admin_lside.html")
	MyTemplate.ExecuteTemplate(w, "index", rd)
}

//文章列表
func Admin_ArticleList(w http.ResponseWriter,r *http.Request) {
	rd := RetData{
		Title: "文章列表 - 后台管理",
		//Archive: models.Archive(),
	}
	MyTemplate := bwy.InitTemplate()
	//模板
	MyTemplate.ParseFiles("./views/admin/article_list.html", "./views/common/_admin_lside.html")
	MyTemplate.ExecuteTemplate(w, "article_list", rd)
}
//文章 添加|修改
func Admin_ArticleCreate(w http.ResponseWriter,r *http.Request) {
	//判断 - 修改|添加

	//判断 get | post

	rd := RetData{
		Title: "文章(添加|修改) - 后台管理",
		//Archive: models.Archive(),
	}
	MyTemplate := bwy.InitTemplate()
	//模板
	MyTemplate.ParseFiles("./views/admin/article_list.html", "./views/common/_admin_lside.html")
	MyTemplate.ExecuteTemplate(w, "article_list", rd)
}
func Admin_ArticleDelete(w http.ResponseWriter,r *http.Request) {

}

//退出登录
func Admin_outLogin(w http.ResponseWriter,r *http.Request) {
	checklogindata.Token = ""
	checklogindata.Maxage = 0

	w.Header().Set("Location", "/admin/login")
	w.WriteHeader(301)
}