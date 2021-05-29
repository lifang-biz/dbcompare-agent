package main

import (
	"lifang.biz/dbcompare-client/api"
	"lifang.biz/dbcompare-client/conf"
	"lifang.biz/dbcompare-client/db"
	"lifang.biz/dbcompare-client/pkg"
)

func main()  {
	conf.Setup()
	pkg.SetupLogger()
	db.Setup()

	api.Start() //启动 API 接口 这个在主线程里跑，卡住线程，防止程序运行直接退出
}
