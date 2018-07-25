package main

import (
	"fmt"
	"net/http"
	"www/engine"
	"flag"
)


func main() {
	//配置文件
	ConfPath := flag.String("cpath", "/Users/wangzhigang/go/src/www/config.conf", "config file")
	engine.Inits(*ConfPath)

	static()	//静态文件处理
	http.HandleFunc("/", engine.Engine)
	err := http.ListenAndServe("0.0.0.0:8880", nil)
	if err != nil {
		fmt.Println("http listen failed")
	}
}


func static() {
	//h := http.FileServer(http.Dir("/Users/wangzhigang/go/src/www/static/css/"))
	//http.Handle("/static/css/", http.StripPrefix("/static/css/", h)) // 启动静态文件服务
	//Header().Set("Expires", time.Now().Format("MON, 02 Jan 2006 15:04:05 GMT"))
}

