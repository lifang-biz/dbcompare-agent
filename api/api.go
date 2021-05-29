package api

import (
	"encoding/json"
	"go.uber.org/zap"
	"lifang.biz/dbcompare-client/conf"
	"lifang.biz/dbcompare-client/pkg"
	"log"
	"net/http"
	"strconv"
)

func Start()  {
	addr := "0.0.0.0:" + strconv.Itoa(conf.Config().App.Port)
	s := &http.Server{
		Addr:           addr,
		MaxHeaderBytes: 1 << 30,
	}

	routers()

	pkg.Logger.Info("http.startHttpServer ok, listening", zap.String("address",addr))
	log.Fatalln(s.ListenAndServe())
}


func RenderJson(w http.ResponseWriter, v interface{}) {
	bs, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//跨域支持
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(bs)
}

type JsonResult struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Msg string      `json:"msg"`
}

//获取URL的GET参数
func GetUrlArg(r *http.Request,name string)string{
	var arg string
	values := r.URL.Query()
	arg=values.Get(name)
	return arg
}


func RenderOKJson(w http.ResponseWriter, data interface{}) {
	RenderJson(w, &JsonResult{
		Code:    200,
		Data:    data,
		Msg: "",
	})
}

func RenderErrorJson(w http.ResponseWriter, msg string) {
	RenderJson(w, &JsonResult{
		Code:    404,
		Data:    nil,
		Msg: msg,
	})
}

func RenderErrorJsonWithCode(w http.ResponseWriter, code int ,msg string) {
	RenderJson(w, &JsonResult{
		Code:    code,
		Data:    nil,
		Msg: msg,
	})
}
