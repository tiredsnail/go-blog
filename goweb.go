package main

import (
	"net/http"
	"www/engine"
	"flag"
	 //_ "net/http/pprof" // pprof 的 http 路由注册在自带路由上
)


func main() {
	//go func() {
	//	http.ListenAndServe("0.0.0.0:6060", nil) // 启动默认的 http 服务，可以使用自带的路由
	//}()

	ConfPath := flag.String("cpath", "/Users/wangzhigang/go/src/www/config.conf", "config file")

	//配置文件
	//dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//ConfPath := flag.String("cpath", dir+"/config.conf", "config file")
	engine.Inits(*ConfPath)
	static()	//静态文件处理
	http.HandleFunc("/", engine.Engine)
	http.ListenAndServe("0.0.0.0:8880", nil)



}


func static() {
	//h := http.FileServer(http.Dir("/Users/wangzhigang/go/src/www/static/css/"))
	//http.Handle("/static/css/", http.StripPrefix("/static/css/", h)) // 启动静态文件服务
	//Header().Set("Expires", time.Now().Format("MON, 02 Jan 2006 15:04:05 GMT"))
}

