package db

import (
	"fmt"
	"go.uber.org/zap"
	"lifang.biz/dbcompare-client/conf"
	"lifang.biz/dbcompare-client/pkg"
	"xorm.io/xorm"
	"xorm.io/xorm/log"

	_ "github.com/go-sql-driver/mysql"
)

var engine *xorm.Engine

//建立连接
func Setup() {
	ConnStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4",
		conf.Config().Database.User,
		conf.Config().Database.Password,
		conf.Config().Database.Host,
		conf.Config().Database.Name)

	var err error

	engine, err = xorm.NewEngine("mysql", ConnStr)
	if err != nil {
		pkg.Logger.Error("NewEngine error: ", zap.String("msg",err.Error()))
	}

	level := log.LOG_ERR
	if conf.Config().App.RunMode != "pro" {
		level = log.LOG_DEBUG
	}

	xormLogger := &pkg.XormLogger{}
	xormLogger.SetLevel(level)

	err = engine.Ping()
	if err != nil {
		pkg.Logger.Error("Connect to mysql error: ", zap.String("msg",err.Error()))
	} else {
		pkg.Logger.Info("Connect to mysql OK ")
	}

	if conf.Config().App.RunMode != "pro" {
		engine.ShowSQL(true)
	}
}

// 关于 information_schema 的介绍  https://blog.csdn.net/kikajack/article/details/80065753
// ShowDBs 显示所有数据库
func ShowDBs() []map[string]string {
	sql := "select * from schemata"
	return queryString(sql)
}

func ShowTables(database string) []map[string]string {
	sql := fmt.Sprintf("select * from tables where table_schema='%s'",database)
	return queryString(sql)
}

// ShowTableInfo show full columns from admin
func ShowTableColumns(database string,table string) []map[string]string{
	sql := fmt.Sprintf("select * from columns where table_schema='%s' and table_name='%s'",database,table)
	return queryString(sql)
}
func ShowTableIndexs(database string,table string) []map[string]string{
	sql := fmt.Sprintf("select * from statistics where table_schema='%s' and table_name='%s'",database,table)
	return queryString(sql)
}

// queryString 执行 sql
// 所有的 mysql show 语句：https://dev.mysql.com/doc/refman/5.6/en/show.html
func queryString(sql string) []map[string]string {
	m,err := engine.QueryString(sql)
	if err != nil {
		pkg.Logger.Error("SQL error: ", zap.String("sql",sql),zap.String("msg",err.Error()))
	}
	return m
}