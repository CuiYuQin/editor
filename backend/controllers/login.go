package controllers

import (
	"backend/common"
	"backend/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {

	fmt.Println("GetEbook请求:", r.Method)
	common.SetHeader(w)

	//请求的数据
	requestData := common.GetData(r)
	//返回的数据体
	responseData := common.ResponseData{}
	sql.Login(requestData, &responseData)
	res, _ := json.Marshal(responseData)
	w.Write(res)
}
