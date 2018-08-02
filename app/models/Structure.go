package models

/**
	文章列表
*/
type PageData struct {
	Page		int						//总页码
	CurrentPage	int						//当前页码
	List		[]map[string]string		//list
}
