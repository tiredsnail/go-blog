package controllers

import (
	"net/http"
	"fmt"
	"go-blog/app/models"
	"strings"
	"time"
	"strconv"
	"go-blog/bwy"
	"go-blog/bwy/db"
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
	ret := RetJson{}
	if data.Username == "" || data.Email == "" || data.Content == "" {
		ret.Status = "err"
		ret.Msg = "请求参数错误"
		ret.Success(w)
		return
	}
	//判断文章id是否存在


	//判断昵称是否使用关键词
	if err := checklogin(w ,r); err != nil {
		if data.Username == "管理员" {
			ret.Status = "err"
			ret.Msg = "昵称不能使用关键词 : (管理员)"
			ret.Success(w)
			return
		}
	}


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
	//判断是否是回复
	if data.Pid != "" {
		//查询接收人邮件
		pdata := models.CommentFirst(data.Pid,"email,content")

		//加入回复任务队列 - 数据库
		commentEmailQueue := make(map[string]string)
		commentEmailQueue["comment_id"] = strconv.Itoa(int(resdb))
		commentEmailQueue["email"] = pdata["email"]
		commentEmailQueue["created_at"] = data.Created_at
		commentEmailQueue["content"] = `<html><body><h3>你在【白乌鸦】网站中的评论有人回复了 ,</h3>
		<p>点击查看完整内容:<a href="https://www.baiwuya.cn">https://www.baiwuya.cn/post/`+data.Article_id+`</a></p>
		<b>你的评论内容 :</b>
		<p><xmp>`+pdata["content"]+`</xmp></p>
		<b>回复内容 :</b>
		<xmp>`+data.Content+`</xmp>
		</body>
		</html>
		`
		_ ,err = models.CommentEmailQueuePut(commentEmailQueue)
		if err != nil {
			bwy.MyLog("加入评论回复任务队列失败["+err.Error()+"],comment_id:"+commentEmailQueue["comment_id"])
		}
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


func Tests(w http.ResponseWriter,r *http.Request) {
	DB := db.Db{}
	DB.MysqlConnect()
	list,err := DB.Table("blog_comment_email").Select("*").Limit("1,50").Where("state=0").Get()
	//list,err := DB.Table("blog_article").Select("*").Where("state=0").Order("state asc").Limit("1,10").Get()

	fmt.Println(list)
	fmt.Println(err)
}