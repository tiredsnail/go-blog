package controllers

import (
	"net/http"
	"html/template"
	"go-blog/bwy"
	"fmt"
	"go-blog/app/models"
	"strconv"
	"strings"
	"time"
)

type RetData struct {
	Description string
	Title 		string
	Nav			[]map[string]string
	ArticleList	*models.PageData
	Archive		[]map[string]string
	ArticleData map[string]string
}
var URL_PATH string
func Index(w http.ResponseWriter,r *http.Request) {
	//接收参数
	page := 1
	req := strings.Split(r.URL.Path, "/")
	if len(req) > 2 {
		page ,_ = strconv.Atoi(req[2])
		if page ==0 {page++}
	}

	//分页使用
	URL_PATH = ""

	rd := RetData{
		Title: "白乌鸦 - 一个phper的博客",
		Nav: LayoutType(),
		ArticleList: models.ArticlePostList(page, 3,"","article_id,type_url,type_name,headline,summary,updated_at,comm,pv"),
		//Archive: models.Archive(),
	}
	MyTemplate := bwy.InitTemplate()
	MyTemplate.Funcs(template.FuncMap{"mypages": mypages})
	//模板
	MyTemplate.ParseFiles("./views/index.html", "./views/common/_list.html", "./views/common/_header.html", "./views/common/_rside.html")
	MyTemplate.ExecuteTemplate(w, "index", rd)
}

//分类文章列表
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
	URL_PATH = "/type/"+types

	rd := RetData{
		Title: types+ " - 白乌鸦 - 一个phper的博客",
		Nav: LayoutType(),
		ArticleList: models.ArticlePostList(page, 3,"type_url='"+types+"'","article_id,type_url,type_name,headline,summary,updated_at,comm,pv"),
	}

	MyTemplate := bwy.InitTemplate().Funcs(template.FuncMap{"mypages": mypages})
	//模板
	MyTemplate.ParseFiles("./views/list.html", "./views/common/_list.html", "./views/common/_nav.html")
	MyTemplate.ExecuteTemplate(w, "list", rd)
}

//归档页面
func Archive(w http.ResponseWriter,r *http.Request) {
	rd := RetData{
		Title: "归档 - 白乌鸦",
		Archive: models.Archive(),
	}
	MyTemplate := bwy.InitTemplate()
	//模板
	MyTemplate.ParseFiles("./views/archive.html", "./views/common/_header.html")
	MyTemplate.ExecuteTemplate(w, "archive", rd)
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
	URL_PATH = "/archive/"+created_at

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
		Title: "归档|"+created_at+" - 白乌鸦",
		Description: "归档日期:"+created_at+" - 白乌鸦",
		//Nav: LayoutType(),
		ArticleList: models.ArticlePostList(page, 3,"created_at>'"+strconv.Itoa(int(startTime))+"' and created_at<'"+strconv.Itoa(int(endTime))+"'","article_id,type_url,type_name,headline,summary,updated_at,comm,pv"),
		Archive: models.Archive(),
	}

	MyTemplate := bwy.InitTemplate().Funcs(template.FuncMap{"mypages": mypages})

	//模板
	MyTemplate.ParseFiles("./views/list.html", "./views/common/_list.html", "./views/common/_nav.html")
	MyTemplate.ExecuteTemplate(w, "list", rd)
}

//文章详情
func Post(w http.ResponseWriter,r *http.Request) {
	//接收参数
	req := strings.Split(r.URL.Path, "/")

	rd := &RetData {
		Title: "白乌鸦",
		Nav: LayoutType(),
		ArticleData: models.ArticlePosts(req[2]),
	}

	MyTemplate := bwy.InitTemplate()
	MyTemplate.ParseFiles("./views/post.html", "./views/common/_header.html", "./views/common/_rside.html")
	MyTemplate.ExecuteTemplate(w, "post", &rd)

	//MyTemplate := bwy.InitTemplate()
	//MyTemplate.Funcs(template.FuncMap{"mypages": mypages})
	//MyTemplate.ParseFiles("./views/test.html","./views/common/_test.html")
	//MyTemplate.ExecuteTemplate(w, "test", "")

}

//关于页面



/**
	最新评论文章
*/
func ArticleNewComment() {

}


//分页
func mypages(page int,currentPage int) template.HTML {
	var pages string
	if 1 != currentPage{
		pages = "<li><a href='"+URL_PATH+"/'> 首页 </a></li>"
	}
	if currentPage > 2 {
		pages += "<li><a href='"+URL_PATH+"/page/"+strconv.Itoa(currentPage-1)+"'>上一页</a></li>"
	}
	pages += "<li><a class='active' href='javascript:;'>"+strconv.Itoa(currentPage)+"</a></li>"
	if currentPage < page-1 {
		pages += "<li><a href='"+URL_PATH+"/page/"+strconv.Itoa(currentPage+1)+"'>下一页</a></li>"
	}
	if currentPage < page {
		pages += "<li><a href='"+URL_PATH+"/page/"+strconv.Itoa(page)+"'> 末页 </a></li>"
	}
	return template.HTML(pages)
}
