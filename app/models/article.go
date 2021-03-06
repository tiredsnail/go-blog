package models

import (
	"go-blog/snail-web"
	"go-blog/snail-web/db"
	"log"
	"math"
	"strconv"
)

type ArticleStructure struct {
}

type Article interface {
	ArticleAdd(data *ArticleStructure) //添加
	ArticlePostDel()                   //删除
	ArticlePostEdit()                  //修改
	ArticlePost()                      //查询
	ArticlePostList()                  //列表
}

func ArticleAdd(data *ArticleStructure) {

}

func ArticlePostList(page int, num int, where string, selects string) *PageData {
	//time.Sleep(3000000000)
	DB := db.Db{}
	ALD := PageData{}
	list, err := DB.Table("blog_article").Select(selects).Where(where).Order("article_id desc").Limit(strconv.Itoa(((page - 1) * num)) + "," + strconv.Itoa(num)).Get()
	if err != nil {
		return &ALD
	}
	ALD.List = list
	count, _ := DB.Table("blog_article").Where(where).Count()
	ALD.CurrentPage = page
	ALD.Page = int(math.Ceil(float64(count) / float64(num)))

	return &ALD
}

//最新评论文章
func ArticleCommentData() []map[string]string {
	DB := db.Db{}
	list, err := DB.Table("blog_article").Select("*").Limit("0,10").Get()
	if err != nil {

	}
	return list
}

//归档 list
func Archive() (list []map[string]string) {
	DB := db.Db{}
	DB.MysqlConnect()
	select_sql := "select FROM_UNIXTIME(created_at, '%Y-%m') as created_at, count(*) as cnt from `blog_article` group by FROM_UNIXTIME(created_at, '%Y-%m') desc"

	select_rows, err := db.MysqlConn.Query(select_sql)
	if err != nil {
		snail_web.MyLog("MySql错误:...models/article.go line 58 [error:" + err.Error() + "]")
		return list
	}
	for select_rows.Next() {
		var created_at string
		var cnt string
		record := make(map[string]string)
		if err := select_rows.Scan(&created_at, &cnt); err != nil {
			log.Println(err)
		}
		record["created_at"] = created_at
		record["cnt"] = cnt
		list = append(list, record)
	}
	select_rows.Close()
	return list
}

//文章详情
type ArticleData struct {
	Article_id string
	Type_name  string
	Headline   string
	Content    string
	Updated_at string
	Comm       string
	Pv         string
}

func ArticlePost(id string) (data *ArticleData) {
	DB := db.Db{}
	DB.MysqlConnect()
	select_sql := "SELECT `article_id`,`type_name`,`headline`,`content`,`updated_at`,`comm`,`pv` FROM `blog_article` WHERE article_id=" + id

	select_err := db.MysqlConn.QueryRow(select_sql).Scan(
		data.Article_id,
		data.Type_name,
		data.Headline,
		data.Content,
		data.Updated_at,
		data.Comm,
		data.Pv,
	)

	if select_err != nil { //如果没有查询到任何数据就进入if中err：no rows in result set
		log.Println(select_err)
		return data
	}
	//log.Println(data)
	return data
}

func ArticlePosts(id string) map[string]string {

	DB := db.Db{}
	data := DB.Table("blog_article").Where("article_id=" + id).First("article_id,type_name,headline,content,updated_at,comm,pv")
	return data
}
