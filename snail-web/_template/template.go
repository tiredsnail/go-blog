package _template

import (
	"html/template"
	"strconv"
	"go-blog/snail-web/const"
)
//分页
func Tpages(page int,currentPage int) template.HTML {

	var pages string
	if 1 != currentPage{
		pages = "<li><a href='"+_const.REQUEST_URI+"/'> 首页 </a></li>"
	}
	if currentPage > 2 {
		pages += "<li><a href='"+_const.REQUEST_URI+"/page/"+strconv.Itoa(currentPage-1)+"'>上一页</a></li>"
	}
	pages += "<li><a class='active' href='javascript:;'>"+strconv.Itoa(currentPage)+"</a></li>"
	if currentPage < page-1 {
		pages += "<li><a href='"+_const.REQUEST_URI+"/page/"+strconv.Itoa(currentPage+1)+"'>下一页</a></li>"
	}
	if currentPage < page {
		pages += "<li><a href='"+_const.REQUEST_URI+"/page/"+strconv.Itoa(page)+"'> 末页 </a></li>"
	}
	return template.HTML(pages)
}