package controllers

import (
	"net/http"
	"go-blog/bwy"
	"errors"
	"io"
	"crypto/rand"
	"encoding/base64"
	"crypto/md5"
	"encoding/hex"
	"time"
	"go-blog/app/models"
	"strings"
	"strconv"
	"html/template"
	"go-blog/bwy/config"
)
type AdminData struct {
	Title				string
	Prompt				string
	ArticlePageData		*models.PageData
	LayoutType			[]map[string]string
	Article				*models.ArticleTable
	CommentPageData		*models.PageData
}
type checkloginstruct struct {
	Maxage 	int64			//过期时间
	Token	string			//token
}
var checklogindata checkloginstruct
func checklogin(w http.ResponseWriter,r *http.Request) (errs error) {
	cookie, err := r.Cookie("checklogindata")

	if err != nil {
		//返回登录页面
		return errors.New("cookie读取失败")
	}

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
		number := r.PostFormValue("bwyNumber")
		password := r.PostFormValue("bwyPassword")
		if number == config.CONFIG["init#appNumber"] && password == config.CONFIG["init#appPassword"] {
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
	MyTemplate.ParseFiles("./resources/views/admin/login.html", "./resources/views/common/_admin_login.html")
	MyTemplate.ExecuteTemplate(w, "login", rd)
}

//首页
func Admin_Index(w http.ResponseWriter,r *http.Request) {
	if err := checklogin(w ,r); err != nil {
		w.Header().Set("Location", "/admin/login")
		w.WriteHeader(302)
		return
	}
	//return
	rd := AdminData{
		Title: "后台管理",
	}
	MyTemplate := bwy.InitTemplate()
	//模板
	MyTemplate.ParseFiles("./resources/views/admin/index.html", "./resources/views/common/_admin_lside.html")
	MyTemplate.ExecuteTemplate(w, "index", rd)
}

//文章列表
func Admin_ArticleList(w http.ResponseWriter,r *http.Request) {
	if err := checklogin(w ,r); err != nil {
		w.Header().Set("Location", "/admin/login")
		w.WriteHeader(302)
		return
	}
	//接收参数
	page := 1
	req := strings.Split(r.URL.Path, "/")
	if len(req) > 4 {
		page ,_ = strconv.Atoi(req[4])
		if page ==0 {page++}
	}
	//分页使用
	URL_PATH = "/admin/article"
	rd := AdminData{
		Title: "文章列表 - 后台管理",
		ArticlePageData: models.ArticlePostList(page, 10,"","*"),
	}
	MyTemplate := bwy.InitTemplate()
	MyTemplate.Funcs(template.FuncMap{"mypages": mypages})
	//模板
	MyTemplate.ParseFiles("./resources/views/admin/article_list.html", "./resources/views/common/_admin_lside.html")
	MyTemplate.ExecuteTemplate(w, "article_list", rd)
}
//文章 添加|修改
func Admin_ArticleCreate(w http.ResponseWriter,r *http.Request) {
	if err := checklogin(w ,r); err != nil {
		w.Header().Set("Location", "/admin/login")
		w.WriteHeader(302)
		return
	}
	//判断 - 修改|添加
	rd := AdminData{
		Title: "文章(添加|修改) - 后台管理",
	}
	if len(r.URL.Query()["article_id"]) > 0 { //修改
		rd.LayoutType = LayoutType()
		rd.Article = models.ArticleGetFind(r.URL.Query()["article_id"][0])
	} else {
		rd.LayoutType = LayoutType()
	}


	MyTemplate := bwy.InitTemplate()
	//模板
	MyTemplate.ParseFiles("./resources/views/admin/article_create.html", "./resources/views/common/_admin_lside.html")
	MyTemplate.ExecuteTemplate(w, "article_create", rd)
}
func Admin_ArticleCreateButton(w http.ResponseWriter,r *http.Request) {
	if err := checklogin(w ,r); err != nil {
		w.Header().Set("Location", "/admin/login")
		w.WriteHeader(302)
		return
	}
	//post 数据
	data := models.ArticleTable{
		Headline: r.PostFormValue("headline"),
		Type_url: r.PostFormValue("type_url"),
		State: r.PostFormValue("state"),
		Summary: r.PostFormValue("summary"),
		Content: r.PostFormValue("content"),
		Created_at: strconv.Itoa(int(time.Now().Unix())),
		Updated_at: strconv.Itoa(int(time.Now().Unix())),
	}
	types := LayoutType()
	for i:=0; i < len(types); i++  {
		if data.Type_url ==  types[i]["url"]{
			data.Type_name = types[i]["name"];
		}
	}
	rd := AdminData{
		Title: "文章(添加|修改) - 后台管理",
	}
	//判断 - 修改|添加
	var err error
	if r.PostFormValue("article_id") != "" { //修改
		rd.LayoutType = LayoutType()
		data.Article_id = r.PostFormValue("article_id")
		err = models.ArticleUpdate(&data)
	} else {
		_,err = models.ArticleInsert(&data)
	}

	if err != nil {	//判断是否成功
		rd.Prompt = "操作失败"
	} else {
		rd.Prompt = "操作成功"
	}
	MyTemplate := bwy.InitTemplate()
	//模板
	MyTemplate.ParseFiles("./resources/views/admin/prompt.html", "./resources/views/common/_admin_lside.html")
	MyTemplate.ExecuteTemplate(w, "prompt", rd)
}
//删除
func Admin_ArticleDelete(w http.ResponseWriter,r *http.Request) {
	if err := checklogin(w ,r); err != nil {
		w.Header().Set("Location", "/admin/login")
		w.WriteHeader(302)
		return
	}
	rd := AdminData{
		Title: "文章删除 - 后台管理",
	}
	//删除
	if len(r.URL.Query()["article_id"]) > 0 { //修改
		err := models.ArticleDelete(r.URL.Query()["article_id"][0])
		if err != nil {	//判断是否成功
			rd.Prompt = "删除失败"
		} else {
			rd.Prompt = "删除成功"
		}
	} else {
		rd.Prompt = "参数错误"
	}
	MyTemplate := bwy.InitTemplate()
	//模板
	MyTemplate.ParseFiles("./resources/views/admin/prompt.html", "./resources/views/common/_admin_lside.html")
	MyTemplate.ExecuteTemplate(w, "prompt", rd)
}


//获取七牛token

//退出登录
func Admin_outLogin(w http.ResponseWriter,r *http.Request) {
	if err := checklogin(w ,r); err != nil {
		w.Header().Set("Location", "/admin/login")
		w.WriteHeader(302)
		return
	}
	checklogindata.Token = ""
	checklogindata.Maxage = 0

	w.Header().Set("Location", "/admin/login")
	w.WriteHeader(301)
}


/**
 * 评论
*/
func Admin_CommentList(w http.ResponseWriter,r *http.Request) {
	if err := checklogin(w ,r); err != nil {
		w.Header().Set("Location", "/admin/login")
		w.WriteHeader(302)
		return
	}
	//接收参数
	page := 1
	req := strings.Split(r.URL.Path, "/")
	if len(req) > 4 {
		page ,_ = strconv.Atoi(req[4])
		if page ==0 {page++}
	}
	//分页使用
	URL_PATH = "/admin/comment"
	rd := AdminData{
		Title: "评论列表 - 后台管理",
		CommentPageData: models.CommentList(page, 10,"1"),
	}
	MyTemplate := bwy.InitTemplate()
	MyTemplate.Funcs(template.FuncMap{"mypages": mypages})
	//模板
	MyTemplate.ParseFiles("./resources/views/admin/comment_list.html", "./resources/views/common/_admin_lside.html")
	MyTemplate.ExecuteTemplate(w, "comment_list", rd)
}
func Admin_CommentState(w http.ResponseWriter,r *http.Request) {

}
func Admin_CommentDelete(w http.ResponseWriter,r *http.Request) {

}