package common

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//结构化返回信息
type ResponseData struct {
	Objects interface{} `json:"objects"` //主要数据内容
	ErrMsg  string      `json:"errMsg"`  //提示信息内容
	Status  int        `json:"status"`  //状态码
}

//获取请求数据
func GetData(r *http.Request) (requestData  map[string]interface{}) {
	r.ParseForm()
	defer r.Body.Close()
	boby, err := ioutil.ReadAll(r.Body)
	CheckError(err)
	err = json.Unmarshal(boby, &requestData)
	CheckError(err)
	return 
}

//设置请求头
func SetHeader(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")                           //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type,Authorization") //header的类型
	w.Header().Set("content-type", "application/json")                           //返回数据格式是json
}

//错误处理
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
