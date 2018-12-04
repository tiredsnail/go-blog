package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/smtp"
	"strconv"
	"strings"
	"time"
	"www/bwy/db"
	"www/engine"
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

	msg := []byte("To: " + to + "\r\nFrom: <" + nick + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
}

func main() {
	ConfPath := flag.String("cpath", "/Users/wangzhigang/go/src/www/config.conf", "config file")
	engine.Inits(*ConfPath)
	//list,err := DB.Table("blog_comment_email").Select("*").Limit("1,50").Where("state=0").Get()
	//
	//fmt.Println(list)
	//fmt.Println(err)

	user := "erppcc@163.com"
	password := "erppcc163mm"
	host := "smtp.163.com:25"
	//
	//user := "admin@baiwuya.cn"
	//password := "erppccyx11."
	//host := "smtp.mxhichina.com:25"
	nick := "白乌鸦"
	subject := "BaiWy 评论回复通知 ... "
	//如果没有任务 10 秒读取一次数据库 , 判断任务数量 , 并发处理. 最多3并发 . 300数据
	// email
	ticker := time.NewTicker(time.Second * 10) //10秒执行一次
	go func() {                                //协程
		for {
			//查询是否有任务
			DB := db.Db{}
			list, err := DB.Table("blog_comment_email").Select("comment_email_id,email,content,error_num").Limit("1,50").Where("state=0").Get()
			if err != nil {
				<-ticker.C //10秒执行一次
				continue
			}
			fmt.Println(time.Now().Unix())
			if len(list) > 0 {
				for _, v := range list {
					emailerr := SendToMail(user, password, host, v["email"], subject, v["content"], "html", nick)
					if emailerr != nil {
						//DB := db.Db{}
						//DB.MysqlConnect()	//取sql连接
						stmt, err := db.MysqlConn.Prepare("UPDATE blog_comment_email set state=?, error_num=?, error_msg=?, start_time=? WHERE comment_email_id=?")
						if err != nil {
							continue
						}
						error_num, _ := strconv.Atoi(v["error_num"])
						_, err = stmt.Exec(2, error_num+1, emailerr.Error(), time.Now().Unix(), v["comment_email_id"])
						if err != nil {
							continue
						}
						//num, err := res.RowsAffected()
						//if err != nil {
						//	continue
						//}
						stmt.Close()
					} else {
						stmt, err := db.MysqlConn.Prepare("UPDATE blog_comment_email set state=?, start_time=? WHERE comment_email_id=?")
						if err != nil {
							continue
						}
						res, err := stmt.Exec(1, time.Now().Unix(), v["comment_email_id"])
						if err != nil {
							continue
						}
						num, err := res.RowsAffected()
						fmt.Println("Send success", v["email"], num)

					}
				}
			} else { //没有 等待10秒
				<-ticker.C //10秒执行一次
			}

		}
	}()

	http.HandleFunc("/monitor", func(writer http.ResponseWriter, request *http.Request) {

	})

	http.ListenAndServe(":14315", nil)
}
