package controllers

import (
	"net/http"
	"www/bwy"
	"www/app/models"
)

func About(w http.ResponseWriter,r *http.Request) {
	rd := RetData{
		Title: "关于 - 白乌鸦",
		Nav: LayoutType(),
		Archive: models.Archive(),
	}
	MyTemplate := bwy.InitTemplate()
	//模板
	MyTemplate.ParseFiles("./views/about.html", "./views/common/_header.html")
	MyTemplate.ExecuteTemplate(w, "about", rd)
}
