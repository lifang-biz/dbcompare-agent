package api

import (
	"lifang.biz/dbcompare-client/conf"
	"lifang.biz/dbcompare-client/db"
	"net/http"
)

func routers() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		RenderOKJson(w, "welcome to db compare")
	})

	//数据库列表 http://127.0.0.1:9100/databases
	http.HandleFunc("/databases", func(w http.ResponseWriter, r *http.Request) {
		RenderOKJson(w, db.ShowDBs())
	})

	//数据库的表列表 http://127.0.0.1:9100/tables?database=xlhd
	http.HandleFunc("/tables", func(w http.ResponseWriter, r *http.Request) {
		database := GetUrlArg( r,"database")
		if len(database) < 1 {
			RenderErrorJson(w,"请提供数据库名称参数")
		}else{
			RenderOKJson(w, db.ShowTables(database))
		}
	})

	//数据库的表结构 http://127.0.0.1:9100/table-columns?database=xlhd&table=admin
	http.HandleFunc("/table-columns", func(w http.ResponseWriter, r *http.Request) {
		database := GetUrlArg( r,"database")
		tableName := GetUrlArg( r,"table")
		if len(tableName) < 1 || len(database) < 1{
			RenderErrorJson(w,"请提供参数")
		}else{
			RenderOKJson(w, db.ShowTableColumns(database,tableName))
		}
	})

	//数据库的表索引  http://127.0.0.1:9100/table-index?database=xlhd&table=admin
	http.HandleFunc("/table-index", func(w http.ResponseWriter, r *http.Request) {
		database := GetUrlArg( r,"database")
		tableName := GetUrlArg( r,"table")
		if len(tableName) < 1 || len(database) < 1{
			RenderErrorJson(w,"请提供参数")
		}else{
			RenderOKJson(w, db.ShowTableIndexs(database,tableName))
		}
	})

	// 刷新配置
	http.HandleFunc("/config/reload", func(w http.ResponseWriter, r *http.Request) {
		conf.Reload()
		RenderOKJson(w, conf.Config())
	})
}
