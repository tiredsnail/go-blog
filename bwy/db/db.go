package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"go-blog/bwy/config"
	"go-blog/bwy"
	"strings"
	"log"
)
var MysqlConn	 	*sql.DB
type Db struct {
	tables		string
	selects 	string
	wheres		string
	limits		string
	joins		string
	orders		string
}

//func init() {
//	fmt.Println("初始化")
//	//if MysqlConn == nil {
//	//	MysqlConn,_ = sql.Open("mysql", config.CONFIG["database|mysqlUser"]+":"+config.CONFIG["database|mysqlPwd"]+"@tcp("+config.CONFIG["database|mysqlHost"]+":"+config.CONFIG["database|mysqlPort"]+")/"+config.CONFIG["database|mysqlDatabase"]+"?parseTime=true")
//	//}
//}

func (DB *Db) MysqlConnect() *Db {
	if MysqlConn == nil {
		MysqlConn, _ = sql.Open("mysql", config.CONFIG["database|mysqlUser"]+":"+config.CONFIG["database|mysqlPwd"]+"@tcp("+config.CONFIG["database|mysqlHost"]+":"+config.CONFIG["database|mysqlPort"]+")/"+config.CONFIG["database|mysqlDatabase"])
		MysqlConn.SetMaxOpenConns(100)		//最大连接数
		MysqlConn.SetMaxIdleConns(50)		//空闲连接数
	}
	//MysqlConn.Ping()
	return DB
}

//table
func (DB *Db) Table(table string) *Db {
	DB.MysqlConnect()
	DB.tables = table
	return DB
}

//条件
func (DB *Db) Where(where string) *Db {
	if where != "" {
		DB.wheres = " WHERE "+where
	}
	return DB
}

//查询字段
func (DB *Db) Select(selects string) *Db {
	DB.selects = selects
	return DB
}

//limit
func (DB *Db) Limit(limit string) *Db {
	if limit != "" {
		DB.limits = " LIMIT " + limit
	}
	return DB
}

func (DB *Db) Join(join string) *Db {
	DB.joins = join
	return DB
}

func (DB *Db) Order(order string) *Db {
	if order != "" {
		DB.orders = " ORDER BY "+order;
	}
	return DB
}

//查询
func (DB *Db) Get() ([]map[string]string,error) {
	select_sql := "SELECT "+DB.selects+" FROM "+DB.tables
	if DB.joins != "" {
		select_sql += " "+DB.joins
	}
	if DB.wheres != "" {
		select_sql += DB.wheres
	}
	if DB.orders != "" {
		select_sql += DB.orders
	}
	if DB.limits != "" {
		select_sql += DB.limits
	}
	var data []map[string]string
	//查询多条
	select_rows,err := MysqlConn.Query(select_sql)
	if err != nil {
		bwy.MyLog("MySql错误:...bwy/db/db.go line 95 [error:"+err.Error()+"]")
		return data,err
	}
	for select_rows.Next() {
		columns, _ := select_rows.Columns()

		scanArgs := make([]interface{}, len(columns))
		values := make([]interface{}, len(columns))

		for i := range values {
			scanArgs[i] = &values[i]
		}

		//将数据保存到 record 字典
		err = select_rows.Scan(scanArgs...)
		record := make(map[string]string)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}
		data = append(data, record)
	}
	select_rows.Close()
	return data,err
}

//查询单条
func (DB *Db) First(selects string) (data map[string]string) {
	select_sql := "SELECT "+selects+" FROM "+DB.tables
	if DB.wheres != "" {
		select_sql += DB.wheres
	}

	columns := strings.Split(selects,",")
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))

	for i := range values {
		scanArgs[i] = &values[i]
	}
	select_err := MysqlConn.QueryRow(select_sql).Scan(scanArgs...)
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

//删除
func delete() {}

//更新
func update() {}

//添加单条
func (DB *Db) Insert(data map[string]string) {

}

//添加多条
func (DB *Db) InsertAll() {

}

//count
func (DB *Db) Count() (int,error) {
	sql := "SELECT count(*) FROM `"+DB.tables+"`"
	if DB.wheres != "" {
		sql += DB.wheres
	}
	var count int
	err := MysqlConn.QueryRow(sql).Scan(&count)
	if err != nil {
		bwy.MyLog("MySql错误:...bwy/db/db.go line 149 [error:"+err.Error()+"]")
		return 0,err
	}
	//MysqlConn.Close()	 | 不需要关闭
	return count,err
}

//原始sql
func (DB *Db) _DoExec()  {

}


/**
 * 启动事务
 * @return void
*/
func startTrans() {}

/**
* 用于非自动提交状态下面的查询提交
* @return boolen
*/
func commit() {}

/**
 * 事务回滚
 * @return boolen
*/
func rollback() {}