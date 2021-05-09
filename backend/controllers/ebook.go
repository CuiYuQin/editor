package controllers

import (
	"backend/common"
	"backend/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetEbook(w http.ResponseWriter, r *http.Request) {

	fmt.Println("GetEbook请求:", r.Method)
	common.SetHeader(w)

	//请求的数据
	requestData := common.GetData(r)
	//返回的数据体
	responseData := common.ResponseData{}
	sql.GetEbook(requestData, &responseData)
	res, _ := json.Marshal(responseData)
	w.Write(res)
}

func InsertEbook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("InsertEbook请求:", r.Method)
	common.SetHeader(w)

	//请求的数据
	requestData := common.GetData(r)
	//返回的数据体
	responseData := common.ResponseData{}
	sql.InsertEbook(requestData, &responseData)
	res, _ := json.Marshal(responseData)
	w.Write(res)
}

func UpdateEbook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UpdateEbook请求:", r.Method)
	common.SetHeader(w)

	//请求的数据
	requestData := common.GetData(r)
	//返回的数据体
	responseData := common.ResponseData{}
	sql.UpdateEbook(requestData, &responseData)
	res, _ := json.Marshal(responseData)
	w.Write(res)
}
func UpdateEbookTitle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UpdateEbookTitle请求:", r.Method)
	common.SetHeader(w)

	//请求的数据
	requestData := common.GetData(r)
	//返回的数据体
	responseData := common.ResponseData{}
	sql.UpdateEbookTitle(requestData, &responseData)
	res, _ := json.Marshal(responseData)
	w.Write(res)
}
