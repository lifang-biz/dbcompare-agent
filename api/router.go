package api

import (
	"lifang.biz/dbcompare-client/conf"
	"lifang.biz/dbcompare-client/db"
	"net/http"
)

func routers() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		RenderOKJson(w, "welcome")
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

	//数据库的表结构和索引 http://127.0.0.1:9100/table-info?database=xlhd&table=admin
	http.HandleFunc("/table-info", func(w http.ResponseWriter, r *http.Request) {
		database := GetUrlArg( r,"database")
		tableName := GetUrlArg( r,"table")
		if len(tableName) < 1 || len(database) < 1{
			RenderErrorJson(w,"请提供参数")
		}else{
			type res struct {
				Columns []map[string]string
				Indexes []map[string]string
			}
			RenderOKJson(w, res{
				Columns: db.ShowTableColumns(database,tableName),
				Indexes: db.ShowTableIndexs(database,tableName),
			})
		}
	})

	// 刷新配置
	http.HandleFunc("/config/reload", func(w http.ResponseWriter, r *http.Request) {
		conf.Reload()
		RenderOKJson(w, conf.Config())
	})
}
