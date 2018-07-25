package engine
import (
	"net/http"
	"html/template"
	"www/bwy"
	bc "www/bwy/config"
	"www/config"
)
func Engine(w http.ResponseWriter,r *http.Request) {
	bwy.MyTemplate = template.New("")

	//常量定义
	bwy.CONSTS_URL_PATH = r.URL.Path

	//通过 url path 判断查询路由 调用匹配方法
	// ....
	config.Route()	//路由规则
	bwy.Match(w, r)	//匹配路由
}

func Inits(ConfPath string) {
	//初始化配置文件
	myConfig := new(bc.Config)
	myConfig.InitConfig(ConfPath)
}