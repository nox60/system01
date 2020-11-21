package dao

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

const (
	DATABASE = "system01"
	CHARSET  = "utf8"
)

var MysqlDb *sql.DB
var MysqlDbErr error

// 初始化链接
func Init(DB_USER string, DB_HOST string, DB_PASSWORD string, DB_PORT string) {

	dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DATABASE, CHARSET)

	// 打开连接失败
	MysqlDb, MysqlDbErr = sql.Open("mysql", dbDSN)
	//defer MysqlDb.Close();
	if MysqlDbErr != nil {
		log.Println("dbDSN: " + dbDSN)
		panic("数据源配置不正确: " + MysqlDbErr.Error())
	}

	// 最大连接数
	MysqlDb.SetMaxOpenConns(100)
	// 闲置连接数
	MysqlDb.SetMaxIdleConns(20)
	// 最大连接周期
	MysqlDb.SetConnMaxLifetime(100 * time.Second)

	if MysqlDbErr = MysqlDb.Ping(); nil != MysqlDbErr {
		panic("数据库链接失败: " + MysqlDbErr.Error())
	}

}
