package models

import (
	"go-blog/bwy/db"
	"strconv"
	"math"
	"go-blog/bwy"
	"log"
	"strings"
)


type CommentData struct {
	Username 	string
	Email		string
	Url			string
	Content		string
	Article_id	string
	Pid			string
	Ip			string
	Created_at	string
}

type Comment interface {
	Add()	//添加
	Del()	//删除
	Edit()	//修改
	Post()	//查询
	List()	//列表
}

func CommentAdd(data *CommentData) (int64,error) {
	DB := db.Db{}
	DB.MysqlConnect()	//取sql连接
	stmt, err := db.MysqlConn.Prepare(`INSERT blog_comment (article_id, nick, email, url, content, Pid, ip, created_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		bwy.MyLog("评论添加失败 failed Prepare")
		return 0,err
	}
	rows, err := stmt.Exec(&data.Article_id,&data.Username,&data.Email,&data.Url,&data.Content,&data.Pid,&data.Ip,&data.Created_at)
	if err != nil {
		bwy.MyLog("评论添加失败 failed Exec")
		return 0,err
	}
	comment_id,err := rows.LastInsertId()
	stmt.Close()
	return comment_id,err
}

func CommentList(page int,num int, article_id string) (*PageData) {
	//time.Sleep(3000000000)
	ALD := PageData{}
	DB := db.Db{}
	list,err := DB.Table("`blog_comment` a").Order("a.comment_id desc").Select("a.comment_id,a.article_id,a.nick,a.url,a.content,a.pid,a.created_at,b.nick as pnick,b.content as pcontent").Join("left join `blog_comment` b on a.pid = b.comment_id").Where("a.article_id="+article_id).Limit(strconv.Itoa(((page-1)*num))+","+strconv.Itoa(num)).Get()

	//DB.MysqlConn().Conn.Query()

	if err != nil {
		return &ALD
	}
	ALD.List = list

	count, _ := DB.Table("blog_comment").Where("article_id="+article_id).Count()
	ALD.CurrentPage = page
	ALD.Page = int(math.Ceil(float64(count) / float64(num) ))

	return &ALD
}

func CommentFirst(comment_id string,selects string) (map[string]string) {
	DB := db.Db{}
	DB.MysqlConnect()
	select_sql := "SELECT "+selects+" FROM `blog_comment` WHERE comment_id="+comment_id

	columns := strings.Split(selects,",")
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))

	for i := range values {
		scanArgs[i] = &values[i]
	}

	select_err :=db.MysqlConn.QueryRow(select_sql).Scan(scanArgs...)
	//将数据保存到 record 字典
	record := make(map[string]string)
	for i, col := range values {
		if col != nil {
			record[columns[i]] = string(col.([]byte))
		}
	}

	if select_err != nil { //如果没有查询到任何数据就进入if中err：no rows in result set
		log.Println(select_err)
		return record
	}

	//log.Println(data)
	return record
}


func CommentEmailQueuePut(data map[string]string) (int64,error) {
	DB := db.Db{}
	DB.MysqlConnect()	//取sql连接
	stmt, err := db.MysqlConn.Prepare(`INSERT blog_comment_email (comment_id, email, created_at, content) VALUES (?, ?, ?, ?)`)
	if err != nil {
		return 0,err
	}

	rows, err := stmt.Exec(data["comment_id"],data["email"],data["created_at"],data["content"])
	if err != nil {
		return 0,err
	}
	comment_id,err := rows.LastInsertId()
	stmt.Close()
	return comment_id,err
}
