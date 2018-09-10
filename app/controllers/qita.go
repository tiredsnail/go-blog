package controllers

import (
	"net/http"
	"go-blog/bwy"
	"go-blog/app/models"
	"go-blog/bwy/config"
)

func About(w http.ResponseWriter,r *http.Request) {
	rd := RetData{
		Title: "关于 - "+config.CONFIG["init#appIndexName"],
		Nav: LayoutType(),
		Archive: models.Archive(),
	}
	MyTemplate := bwy.InitTemplate()
	//模板
	MyTemplate.ParseFiles("./resources/views/about.html", "./resources/views/common/_header.html")
	MyTemplate.ExecuteTemplate(w, "about", rd)
}
