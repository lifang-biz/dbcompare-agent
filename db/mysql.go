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
		conf.Config().DatabaseSetting.User,
		conf.Config().DatabaseSetting.Password,
		conf.Config().DatabaseSetting.Host,
		conf.Config().DatabaseSetting.Name)

	var err error

	engine, err = xorm.NewEngine("mysql", ConnStr)
	if err != nil {
		pkg.Logger.Error("NewEngine error: ", zap.String("msg",err.Error()))
	}

	level := log.LOG_ERR
	if conf.Config().AppSetting.RunMode != "pro" {
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

	if conf.Config().AppSetting.RunMode != "pro" {
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
	sql := fmt.Sprintf("select * from tables where table_schema='%s' order by table_name asc",database)
	return queryString(sql)
}

// ShowTableInfo show full columns from admin
func ShowTableColumns(database string,table string) []map[string]string{
	sql := fmt.Sprintf("select * from columns where table_schema='%s' and table_name='%s'",database,table)
	return queryString(sql)
}
func ShowTableIndexs(database string,table string) []map[string]string{
	// 一个索引有多个字段的话，会有多行，需要把字段合一下
	sql := fmt.Sprintf("select *,GROUP_CONCAT(COLUMN_NAME) as COLUMN_NAME_ALL from statistics where table_schema='%s' and table_name='%s' group by INDEX_NAME",database,table)
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