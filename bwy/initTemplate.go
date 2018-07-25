package bwy


import (
	"fmt"
	"html/template"
	"time"
	"strconv"
	"io"
)

var MyTemplate *template.Template

func InitTemplate(fileName ...string) (err error){

	MyTemplate = MyTemplate.Funcs(template.FuncMap{"unescaped": unescaped,"strtotime": strtotime})

	MyTemplate ,err = MyTemplate.ParseFiles(fileName...)
	//MyTemplate = template.Must()
	if err != nil{
		fmt.Println("parse file err:",err)
		return
	}
	return
}

func View(wr io.Writer, name string, data interface{}) {

	MyTemplate.ExecuteTemplate(wr, name, data)
}

//添加函数方法
type FuncsFuncMap map[string]interface{}

// html 解析
func unescaped (x string) template.HTML { return template.HTML(x) }

// 时间戳转时间
func strtotime(timestamp string) string {
	times ,_ := strconv.ParseInt(timestamp,10,64)
	return time.Unix(times, 0).Format("2006-01-02 15:04:05") //设置时间戳 使用模板格式化为日期字符串
}
