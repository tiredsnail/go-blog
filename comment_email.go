package main

import (
	"fmt"
	"net/smtp"
	"strings"
	"time"
)

func SendToMail(user, password, host, to, subject, body, mailtype string, nick string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: <" + nick + ">\r\nSubject: " +subject+ "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
}

func main() {
	//如果没有任务 10 秒读取一次数据库 , 判断任务数量 , 并发处理. 最多3并发 . 300数据
	// email
	ticker := time.NewTicker(time.Second * 10)		//10秒执行一次
	go func() { //协程
		for {
			//查询是否有任务
			//code ...
			// ....
			if true {
				// code ...
				// ...
			} else { //没有 等待10秒
				<-ticker.C		//10秒执行一次
			}

		}
	}()
	//user := "erppcc@163.com"
	//password := "erppcc163mm"
	//host := "smtp.163.com:25"

	user := "admin@baiwuya.cn"
	password := "erppccyx11."
	host := "smtp.mxhichina.com:25"
	nick := "白乌鸦"

	to := "erppcc@163.com"

	subject := "GoLang 系统通知 ... "

	body := `
		<html>
		<body>
		<h3>
		你在【白乌鸦】网站中的评论有人回复了 ,
		</h3>
		<p>点击查看完整内容:<a href="https://www.baiwuya.cn">https://www.baiwuya.cn/post/1</a></p>
		<b>评论内容 :</b>
		<p><xmp><p>一到晚上就睡不着 . 敲代码敲到2点多.我的情况还有的救吗? ...</p></xmp></p>
		<b>回复内容 :</b>
		<xmp><p>你的情况比较严重 . 多半是没得救了 ...</p></xmp>
		</body>
		</html>
		`
	fmt.Println("send email")
	err := SendToMail(user, password, host, to, subject, body, "html",nick)
	if err != nil {
		fmt.Println("Send mail error!")
		fmt.Println(err)
	} else {
		fmt.Println("Send mail success!")
	}

}