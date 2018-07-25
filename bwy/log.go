package bwy

import (
	"os"
	"fmt"
	"time"
	"www/bwy/config"
)

func MyLog(logs string) {
	file, err := os.OpenFile(config.CONFIG["default|logPath"]+"error.log", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil { //抛出错误信息
		panic(fmt.Sprintf("os.Open file error: %s", err.Error()))
	}

	defer file.Close()

	_, err = file.WriteString("["+time.Now().Format("2006-01-02 15:04:05")+"]:"+logs+"\n")
	if err != nil {
		panic(fmt.Sprintf("file.WriteString file error: %s" ,err.Error()))
	}
}
