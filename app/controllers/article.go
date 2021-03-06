package controllers

import (
	"net/http"
	"go-blog/app/models"
	"strconv"
	"strings"
	"go-blog/snail-web/config"
	"html/template"
	"go-blog/snail-web"
	"time"
	"fmt"
	"go-blog/snail-web/db"
	"go-blog/snail-web/_template"
	"go-blog/snail-web/const"
)

type RetData struct {
	Description string
	Title 		string
	Nav			[]map[string]string
	ArticleList	*models.PageData
	Archive		[]map[string]string
	ArticleData map[string]string
}
func Index(w http.ResponseWriter,r *http.Request) {

	//接收参数
	page := 1
	req := strings.Split(r.URL.Path, "/")
	if len(req) > 2 {
		page ,_ = strconv.Atoi(req[2])
		if page ==0 {page++}
	}

	//分页使用
	_const.REQUEST_URI = ""
	rd := RetData{
		Title: config.CONFIG["init#appIndexName"],
		Nav: LayoutType(),
		ArticleList: models.ArticlePostList(page, 3,"","article_id,type_url,type_name,headline,summary,updated_at,comm,pv"),
		//Archive: models.Archive(),
	}

	snail_web.FuncMap = template.FuncMap{"Tpages": _template.Tpages}
	snail_web.Views("", w, "index", rd,
		"resources/views/index.html",
		"resources/views/common/_header.html",
		"resources/views/common/_list.html",
		"resources/views/common/_rside.html")

	//uEnc := base64.URLEncoding.EncodeToString([]byte(r.URL.String()))
	//file,_ := os.OpenFile("storage/framework/views/"+uEnc, os.O_CREATE|os.O_WRONLY, 0755)
	//MyTemplate.MyTemplate.ExecuteTemplate(file, "index", rd)

}


/**
	分类文章列表
*/
func TypeArticleList(w http.ResponseWriter,r *http.Request) {
	page := 1
	//接收参数
	req := strings.Split(r.URL.Path, "/")
	if len(req) > 4 {
		page ,_ = strconv.Atoi(req[4])
		if page ==0 {page++}
	}
	types := req[2]

	//分页使用
	_const.REQUEST_URI = "/type/"+types

	rd := RetData{
		Title: types+ " - "+config.CONFIG["init#appName"],
		Nav: LayoutType(),
		ArticleList: models.ArticlePostList(page, 3,"type_url='"+types+"'","article_id,type_url,type_name,headline,summary,updated_at,comm,pv"),
	}

	snail_web.FuncMap = template.FuncMap{"Tpages": _template.Tpages}
	snail_web.Views("", w, "list", rd,
		"./resources/views/list.html", "./resources/views/common/_list.html", "./resources/views/common/_nav.html")
}

//归档页面
func Archive(w http.ResponseWriter,r *http.Request) {
	rd := RetData{
		Title: "归档 - "+config.CONFIG["init#appName"],
		Archive: models.Archive(),
	}

	snail_web.Views("", w, "archive", rd,
		"./resources/views/archive.html", "./resources/views/common/_header.html")
}

//归档文章列表
func ArchiveArticleList(w http.ResponseWriter,r *http.Request) {
	page := 1
	//接收参数
	req := strings.Split(r.URL.Path, "/")
	if len(req) > 4 {
		page ,_ = strconv.Atoi(req[4])
		if page ==0 {page++}
	}
	created_at := req[2]

	//分页使用
	_const.REQUEST_URI = "/archive/"+created_at

	theTime ,err := time.ParseInLocation("2006-01",created_at,time.Local)
	if err != nil {
		fmt.Println("时间格式化失败")
	}
	year, month, _ := theTime.Date()
	thisMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	start := thisMonth.AddDate(0, 0, 0).Format("2006-01-02")
	end := thisMonth.AddDate(0, 1, -1).Format("2006-01-02")
	theTime, _ = time.ParseInLocation("2006-01-02", start, time.Local) //使用模板在对应时区转化为time.time类型
	startTime := theTime.Unix()
	theTime, _ = time.ParseInLocation("2006-01-02", end, time.Local) //使用模板在对应时区转化为time.time类型
	endTime := theTime.Unix()


	rd := RetData{
		Title: "归档|"+created_at+" - "+config.CONFIG["init#appIndexName"],
		Description: "归档日期:"+created_at+" - 白乌鸦",
		//Nav: LayoutType(),
		ArticleList: models.ArticlePostList(page, 3,"created_at>'"+strconv.Itoa(int(startTime))+"' and created_at<'"+strconv.Itoa(int(endTime))+"'","article_id,type_url,type_name,headline,summary,updated_at,comm,pv"),
		Archive: models.Archive(),
	}


	snail_web.FuncMap = template.FuncMap{"Tpages": _template.Tpages}
	snail_web.Views("", w, "list", rd,
		"./resources/views/list.html", "./resources/views/common/_list.html", "./resources/views/common/_nav.html")
}

//文章详情
func Post(w http.ResponseWriter,r *http.Request) {
	//接收参数
	req := strings.Split(r.URL.Path, "/")

	rd := &RetData {
		Title: config.CONFIG["init#appName"],
		Nav: LayoutType(),
		ArticleData: models.ArticlePosts(req[2]),
	}
	DB := db.Db{}
	DB.MysqlConnect()
	db.MysqlConn.Exec("UPDATE blog_article set pv=pv+1 WHERE article_id="+req[2])

	snail_web.Views("", w, "post", rd,
		"./resources/views/post.html", "./resources/views/common/_header.html", "./resources/views/common/_rside.html")
}

