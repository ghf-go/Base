package Base

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)
// GModel 数据Model
type GModel struct {
	ConName string
	DbCon *sql.DB
	TableName string
	Pk string
	Data map[string]interface{}
}
// GQuery 查询构造器
//type GQuery struct {
//	Where map[string]interface{}
//	Order map[string]string
//	Limit int
//	Skip int
//	Group []string
//	Having map[string]interface{}
//}
// GetMysql 获取Mysql 链接
func GetMysql(conf ConfMysql) *sql.DB{
	db,err := sql.Open("mysql",fmt.Sprintf("%v:%v@%v/%v",conf.User,conf.Pass,conf.Host,conf.DbName))
	if err != nil {
		panic(GError{1002,err.Error()})
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(conf.MaxConLiftTime)

	db.SetMaxOpenConns(conf.MaxOpenCons)
	db.SetMaxIdleConns(conf.MaxIdleCons)
	return db
}

func (m GModel ) Save()  {

}
func (m GModel ) Insert()  {

}
func (m GModel ) Find(id int64)  {

}
func (m GModel ) Delete(id int64)  {

}



