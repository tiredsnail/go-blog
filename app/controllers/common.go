package controllers

import (
	"net/http"
	"encoding/json"
	"fmt"
	"html/template"
	"strconv"
)

type RetJson struct {
	Status		string			`json:"status"`
	Msg			string			`json:"msg"`
	Data		interface{}		`json:"data"`
}

func LayoutType() []map[string]string {
	data := []map[string]string{
		{"name":"GoLang","url":"golang"},
		{"name":"Mysql","url":"mysql"},
		{"name":"PHP","url":"php"},
		{"name":"other","url":"other"},
	}
	return data
}


func (ret *RetJson) Success(w http.ResponseWriter) {
	retjson, _ := json.Marshal(ret)
	fmt.Fprintf(w, string(retjson) )
	return
}

func Error404(w http.ResponseWriter,r *http.Request) {

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