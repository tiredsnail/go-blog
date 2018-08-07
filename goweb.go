package main

import (
	"net/http"
	"go-blog/engine"

	"flag"
)


func main() {
	//go func() {
	//	http.ListenAndServe("0.0.0.0:6060", nil) // 启动默认的 http 服务，可以使用自带的路由
	//}()

	ConfPath := flag.String("cpath", "/Users/wangzhigang/go/src/go-blog/config.conf", "config file")

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
	// 设置静态目录
	fsh := http.FileServer(http.Dir("./views/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fsh))
}

