package bwy


import (
	"html/template"
	"time"
	"strconv"
	"io"
)

type Template interface {
	InitTemplate(Tname string)
}
var Templates struct {
	MyTemplate *template.Template
}

func (t *Templates) InitTemplate(Tname string) {
	t.MyTemplate = t.MyTemplate.New(Tname)
	//自定义公共 模板方法
	t.MyTemplate = t.MyTemplate.Funcs(template.FuncMap{"unescaped": unescaped,"strtotime": strtotime})
	//return MyTemplate
	//MyTemplate ,err = MyTemplate.ParseFiles(fileName...)
	////MyTemplate = template.Must()
	//if err != nil{
	//	fmt.Println("parse file err:",err)
	//	return MyTemplate,err
	//}
}

func (t *Templates) Views(wr io.Writer, name string, data interface{},filenames ...string) {

	//MyTemplate.ExecuteTemplate(wr, name, data)
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
