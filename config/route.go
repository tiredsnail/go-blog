package config

import (
	"go-blog/snail-web"
	"go-blog/app/controllers"
)

func Route() {
	//首页
	snail_web.RouteAny("^/(page)?/?[0-9]*/?$", controllers.Index)

	//分类文章列表页
	snail_web.RouteAny(`^/type/[a-z]+(/page/)?[0-9]*/?$`, controllers.TypeArticleList)

	//归档
	snail_web.RouteAny("^/archive/?$",controllers.Archive)
	//归档文章列表页
	snail_web.RouteAny(`^/archive/[1-9]\d{3}-(0[1-9]|1[0-2])(/page/)?[0-9]?/?$`, controllers.ArchiveArticleList)

	//文章详情页
	snail_web.RouteAny("^/(post/)?[0-9]+/?$", controllers.Post)

	//评论列表
	snail_web.RouteAny("^/comment/list/?$",controllers.CommentList)

	//评论添加
	snail_web.RouteAny("^/comment/add/?$",controllers.CommentAdd)

	//关于
	snail_web.RouteAny("^/about/?$",controllers.About)

	//测试
	snail_web.RouteAny("^/tests/?$",controllers.Tests)


	//后台
	snail_web.RouteAny("^/admin/login/?$",controllers.Admin_Login) //登录
	snail_web.RouteAny("^/admin/?$",controllers.Admin_Index)       //首页

	snail_web.RouteAny("^/admin/article(/page/)?[0-9]*/?$",controllers.Admin_ArticleList)        //文章列表
	snail_web.RouteAny("^/admin/article/create/?$",controllers.Admin_ArticleCreate)              //文章添加|修改
	snail_web.RouteAny("^/admin/article/create_button/?$",controllers.Admin_ArticleCreateButton) //文章添加|修改
	snail_web.RouteAny("^/admin/article/delete/?$",controllers.Admin_ArticleDelete)              //文章删除

	snail_web.RouteAny("^/admin/comment(/page/)?[0-9]*/?$",controllers.Admin_CommentList) //评论列表
	snail_web.RouteAny("^/admin/comment/state/?$",controllers.Admin_CommentState)         //评论状态修改
	snail_web.RouteAny("^/admin/comment/delete/?$",controllers.Admin_CommentDelete)       //评论删除


	snail_web.RouteAny("^/admin/outlogin/?$",controllers.Admin_outLogin) //退出登录

}
