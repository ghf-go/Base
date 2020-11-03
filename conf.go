package Base

import (
	"fmt"
	"time"
)

type GError struct {
	Code int
	Msg string
}
// ConfMysql Mysql 连接配置
type ConfMysql struct {
	Host string
	User string
	Pass string
	DbName string
	TablePrefix string
	MaxConLiftTime time.Duration
	MaxOpenCons int
	MaxIdleCons int
}

// AppConf 应用程序配置
type AppConf struct {
	Dbs map[string]ConfMysql
}

// GetDbConf 获取应用程序的数据库配置
func (a AppConf) GetDbConf(name string)  ConfMysql{
	ret , ok := a.Dbs[name]
	if ok{
		return ret
	}
	panic(GError{1001,fmt.Sprintf("应用没有配置数据库连接 : %v",name)})
}