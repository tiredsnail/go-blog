package engine
import (
	"net/http"
	"go-blog/snail-web"
	bc "go-blog/snail-web/config"
	"go-blog/config"
)
//type App struct {
//	Request 	*http.Request
//	ResponseWriter http.ResponseWriter
//}

func Engine(w http.ResponseWriter,r *http.Request) {
	//定义常量
	//_const.Request(r)

	snail_web.Match(&w, r) //匹配路由
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

