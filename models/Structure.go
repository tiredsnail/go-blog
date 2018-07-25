package models

/**
	分类列表
*/
type ArticleListData struct {
	Page		int						//总页码
	CurrentPage	int						//当前页码
	List		[]map[string]string		//文章list
}
