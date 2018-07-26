package config

import (
	"www/bwy"
	"www/controllers"
)

func Route() {
	//首页
	bwy.RouteAny("^/(page)?/?[0-9]*/?$", controllers.Index)

	//分类文章列表页
	bwy.RouteAny(`^/type/[a-z]+(/page/)?[0-9]?/?$`, controllers.TypeArticleList)

	//归档
	bwy.RouteAny("^/archive/?$",controllers.Archive)
	//归档文章列表页
	bwy.RouteAny(`^/archive/[1-9]\d{3}-(0[1-9]|1[0-2])(/page/)?[0-9]?/?$`, controllers.ArchiveArticleList)

	//文章详情页
	bwy.RouteAny("^/(post/)?[0-9]+/?$", controllers.Post)

	//评论列表
	bwy.RouteAny("^/comment/list/?$",controllers.CommentList)

	//评论添加
	bwy.RouteAny("^/comment/add/?$",controllers.CommentAdd)
}
