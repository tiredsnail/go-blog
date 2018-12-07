package controllers

import (
	"net/http"
	"go-blog/snail-web"
	"go-blog/app/models"
	"go-blog/snail-web/config"
)

func LayoutType() []map[string]string {
	data := []map[string]string{
		{"name":"GoLang","url":"golang"},
		{"name":"Mysql","url":"mysql"},
		{"name":"PHP","url":"php"},
		{"name":"other","url":"other"},
	}
	return data
}

func About(w http.ResponseWriter, r *http.Request) {
	rd := RetData{
		Title: "关于 - "+config.CONFIG["init#appIndexName"],
		Nav: LayoutType(),
		Archive: models.Archive(),
	}
	snail_web.Views("", w, "index", rd,
		"./resources/views/about.html", "./resources/views/common/_header.html")
}
