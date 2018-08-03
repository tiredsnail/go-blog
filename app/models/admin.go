package models

import (
	"www/bwy/db"
	"log"
)

//文章 添加|修改 - 提交
type ArticleTable struct {
	Article_id		string
	Headline		string
	Type_url		string
	Type_name		string
	State			string
	Summary			string
	Content			string
	Created_at		string
	Updated_at		string
}

func ArticleInsert(data *ArticleTable) (int64 ,error) {
	DB := db.Db{}
	DB.MysqlConnect()	//取sql连接
	stmt, err := db.MysqlConn.Prepare(`INSERT blog_article (headline, type_url, type_name, state, summary, content, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		return 0,err
	}
	rows, err := stmt.Exec(&data.Headline,&data.Type_url,&data.Type_name,&data.State,&data.Summary,&data.Content,&data.Created_at,&data.Updated_at)
	if err != nil {
		return 0,err
	}
	article_id,err := rows.LastInsertId()
	stmt.Close()
	return article_id,err
}
func ArticleUpdate(data *ArticleTable) (error) {
	DB := db.Db{}
	DB.MysqlConnect()	//取sql连接
	stmt, err := db.MysqlConn.Prepare("UPDATE blog_article set headline=?, type_url=?, type_name=?, state=?, summary=?, content=?, updated_at=? WHERE article_id=?")
	if err != nil {
		return err
	}
	res, err := stmt.Exec(data.Headline, data.Type_url, data.Type_name, data.State, data.Summary, data.Content, data.Updated_at, data.Article_id)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	stmt.Close()
	return nil
}

func ArticleDelete(article_id string) (error) {
	DB := db.Db{}
	DB.MysqlConnect()	//取sql连接
	stmt, err := db.MysqlConn.Prepare("DELETE FROM blog_article WHERE article_id=?")
	if err != nil {
		return err
	}

	res, err := stmt.Exec(article_id)
	stmt.Close()
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func ArticleGetFind(article_id string) *ArticleTable {
	DB := db.Db{}
	DB.MysqlConnect()
	select_sql := "SELECT `article_id`,`type_url`,`headline`,`content`,`state`,`summary` FROM `blog_article` WHERE article_id="+article_id

	data := ArticleTable{}
	select_err :=db.MysqlConn.QueryRow(select_sql).Scan(
		&data.Article_id,
		&data.Type_url,
		&data.Headline,
		&data.Content,
		&data.State,
		&data.Summary,
	)

	if select_err != nil { //如果没有查询到任何数据就进入if中err：no rows in result set
		log.Println(select_err)
		return &data
	}
	//log.Println(data)
	return &data
}