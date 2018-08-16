package main

import (
	"net/http"
	"go-blog/engine"
	"flag"
	"fmt"
	"path/filepath"
	"os"
)


func main() {
	//go func() {
	//	http.ListenAndServe("0.0.0.0:6060", nil) // 启动默认的 http 服务，可以使用自带的路由
	//}()
	//配置文件
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err.Error())
	}
	dir = "/Users/wangzhigang/go/src/go-blog"		//开发环境
	ConfPath := flag.String("cpath", dir+"/config.conf", "config file")
	flag.Parse()
	engine.Inits(*ConfPath)

	static()	//静态文件处理
	http.HandleFunc("/", engine.Engine)
	go func() {
		err = http.ListenAndServeTLS(":443", "./storage/chain/full_chain.pem","./storage/chain/private.key", nil)
		if err != nil {
			fmt.Println("开启HTTPS协议失败",err.Error())
		}
	}()

	err = http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Println("没有权限开启80端口,已经开启8080端口")
		http.ListenAndServe(":8080", nil)
	}


}


func static() {
	// 设置静态目录
	fsh := http.FileServer(http.Dir("./views/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fsh))
}

