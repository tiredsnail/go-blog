package controllers

import (
	"net/http"
	"fmt"
	"www/models"
	"strings"
	"time"
	"strconv"
)

func CommentAdd(w http.ResponseWriter,r *http.Request) {
	//post 数据
	data := models.CommentData{
		Username:r.PostFormValue("bwy_username"),
		Email:r.PostFormValue("bwy_email"),
		Url:r.PostFormValue("bwy_url"),
		Content:r.PostFormValue("bwy_comment"),
		Article_id:r.PostFormValue("article_id"),
		Pid:r.PostFormValue("pid"),
	}
	data.Created_at = strconv.Itoa(int(time.Now().Unix()))
	fmt.Println(data.Created_at)
	ret := RetJson{}
	if data.Username == "" || data.Email == "" || data.Content == "" {
		ret.Status = "err"
		ret.Msg = "请求参数错误"
		ret.Success(w)
		return
	}
	//判断文章id是否存在

	//判断是否是回复
	if data.Pid != "" {
		//加入回复任务队列
	}

	//判断昵称是否使用关键词
	if data.Username == "管理员" {
		ret.Status = "err"
		ret.Msg = "昵称不能使用关键词 : (管理员)"
		ret.Success(w)
		return
	}

	fmt.Println(data)

	//获取 ip
	ip := r.RemoteAddr
	//去除端口
	data.Ip = strings.Split(ip, ":")[0]

	//判断 10分钟内 ip 评论次数 < 100
	//DB := db.Db{}
	//created_at := time.Now().Unix() - 60*10
	//if count, err := DB.Table("blog_comment").Where("ip='"+data.Ip+"' and created_at>"+strconv.Itoa(int(created_at)) ).Count();err != nil || count > 100 {
	//	ret.Status = "err"
	//	ret.Msg = "评论过于频繁..10分钟100次限制"
	//	ret.Success(w)
	//	return
	//}


	resdb,err := models.CommentAdd(&data)
	if err != nil {
		ret.Status = "err"
		ret.Msg = "服务器繁忙..请稍候再试"
		ret.Success(w)
		return
	}
	ret.Status = "ok"
	ret.Msg = "评论成功"
	ret.Data = resdb
	ret.Success(w)
	return
}

func CommentList(w http.ResponseWriter,r *http.Request) {
	article_id := r.PostFormValue("article_id")
	page := 1
	if r.PostFormValue("page") != "" {
		page ,_ = strconv.Atoi(r.PostFormValue("page"))
		if page == 0 {page++}
	}

	data := models.CommentList(page,5,article_id)

	ret := RetJson{}
	ret.Status = "ok"
	ret.Msg = "查询成功"
	ret.Data = data
	ret.Success(w)
	return
}
