package engine
import (
	"net/http"
	"go-blog/bwy"
	bc "go-blog/bwy/config"
	"go-blog/config"
)
//type App struct {
//	Request 	*http.Request
//	ResponseWriter http.ResponseWriter
//}

func Engine(w http.ResponseWriter,r *http.Request) {
	//常量定义
	//bwy.CONSTS_URL_PATH = r.URL.Path

	//通过 url path 判断查询路由 调用匹配方法
	// ....
	//_ = App {
	//	Request:r,
	//	ResponseWriter:w,
	//}
	bwy.Match(&w, r)	//匹配路由
}

func Inits(ConfPath string) {
	//初始化配置文件
	myConfig := new(bc.Config)
	myConfig.InitConfig(ConfPath)


}

func init() {
	//初始化路由
	config.Route()	//路由规则
}

