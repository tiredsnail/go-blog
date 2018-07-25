package models

import (
	"www/bwy/db"
	"fmt"
	"www/bwy"
	"database/sql"
	"strconv"
	"math"
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
		fmt.Println("评论添加失败:Prepare")
	}
	rows, err := stmt.Exec(&data.Article_id,&data.Username,&data.Email,&data.Url,&data.Content,&data.Pid,&data.Ip,&data.Created_at)
	if err != nil {
		fmt.Println("评论添加失败:Exec")
	}
	comment_id,err := rows.LastInsertId()
	stmt.Close()
	return comment_id,err
}

func CommentList(page int,num int, article_id string) (*ArticleListData) {
	bwy.MyLog("--------------------------------------")
	//time.Sleep(3000000000)
	ALD := ArticleListData{}
	DB := db.Db{}
	list,err := DB.Table("`blog_comment` a").Order("a.comment_id desc").Select("a.comment_id,a.nick,a.url,a.content,a.pid,a.created_at,b.nick as pnick,b.content as pcontent").Join("left join `blog_comment` b on a.pid = b.comment_id").Where("a.article_id=1").Limit(strconv.Itoa(((page-1)*num))+","+strconv.Itoa(num)).Get()

	//DB.MysqlConn().Conn.Query()

	if err != nil {
		return &ALD
	}
	ALD.List = list

	count, _ := DB.Table("blog_comment").Where("article_id=1").Count()
	ALD.CurrentPage = page
	ALD.Page = int(math.Ceil(float64(count) / float64(num) ))


	//sqls := "select a.comment_id,a.nick,a.url,a.content,a.pid,a.created_at,b.nick as pnick,b.content as bcontent from `blog_comment` a left join `blog_comment` b on a.pid = b.comment_id"
	//if MysqlConn == nil {
	//	MysqlConn, _ = sql.Open("mysql", config.CONFIG["database|mysqlUser"]+":"+config.CONFIG["database|mysqlPwd"]+"@tcp("+config.CONFIG["database|mysqlHost"]+":"+config.CONFIG["database|mysqlPort"]+")/"+config.CONFIG["database|mysqlDatabase"])
	//}
	//select_rows ,err := MysqlConn.Query(sqls);
	//if err != nil {
	//	fmt.Println("错误",err.Error())
	//	return &ALD
	//}
	//var data []map[string]string
	//for select_rows.Next() {
	//	columns, _ := select_rows.Columns()
	//
	//	scanArgs := make([]interface{}, len(columns))
	//	values := make([]interface{}, len(columns))
	//
	//	for i := range values {
	//		scanArgs[i] = &values[i]
	//	}
	//
	//	//将数据保存到 record 字典
	//	err = select_rows.Scan(scanArgs...)
	//	record := make(map[string]string)
	//	for i, col := range values {
	//		if col != nil {
	//			record[columns[i]] = string(col.([]byte))
	//		}
	//	}
	//	data = append(data, record)
	//}
	//select_rows.Close()
	//
	//ALD.List = data

	return &ALD
}
var MysqlConn *sql.DB
