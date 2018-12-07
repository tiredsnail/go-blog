package snail_web


import (
	"html/template"
	"time"
	"strconv"
	"io"
)

type Template interface {
	InitTemplate(Tname string)
}

var FuncMap template.FuncMap

func InitTemplate(Tname string) *template.Template {
	t := template.New(Tname)
	//自定义公共 模板方法
	t = t.Funcs(template.FuncMap{"unescaped": unescaped,"strtotime": strtotime})

	return t

	//MyTemplate ,err = MyTemplate.ParseFiles(fileName...)
	////MyTemplate = template.Must()
	//if err != nil{
	//	fmt.Println("parse file err:",err)
	//	return MyTemplate,err
	//}
}

func Views(Tname string, wr io.Writer, name string, data interface{},filenames ...string) {
	t := template.New(Tname)
	//自定义公共 模板方法
	t = t.Funcs(template.FuncMap{"unescaped": unescaped,"strtotime": strtotime})
	//FuncMap := template.FuncMap{"unescaped": unescaped,"strtotime": strtotime}
	t = t.Funcs(FuncMap)

	t = template.Must(t.ParseFiles(filenames ...))

	t.ExecuteTemplate(wr, name, data)
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
