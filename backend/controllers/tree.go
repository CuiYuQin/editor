package controllers

import (
	"backend/common"
	"backend/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetTree(w http.ResponseWriter, r *http.Request) {

	fmt.Println("GetTree请求:", r.Method)
	common.SetHeader(w)

	//请求的数据
	requestData := common.GetData(r)
	//返回的数据体
	responseData := common.ResponseData{}
	sql.GetTree(requestData, &responseData)
	res, _ := json.Marshal(responseData)
	w.Write(res)
}

func InsertTree(w http.ResponseWriter, r *http.Request) {
	fmt.Println("InsertTree请求:", r.Method)
	common.SetHeader(w)

	//请求的数据
	requestData := common.GetData(r)
	requestDataTree, err := json.Marshal(requestData["tree"])
	common.CheckError(err)
	requestData["tree"] = requestDataTree
	requestDataDeleteTree, err := json.Marshal(requestData["deleteTree"])
	common.CheckError(err)
	requestData["deleteTree"] = requestDataDeleteTree

	//返回的数据体
	responseData := common.ResponseData{}
	//数据库操作
	sql.InsertTree(requestData, &responseData)
	res, _ := json.Marshal(responseData)
	w.Write(res)
}

func UpdateTree(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UpdateTree请求:", r.Method)
	common.SetHeader(w)

	//请求的数据
	requestData := common.GetData(r)
	requestDataTree, err := json.Marshal(requestData["tree"])
	common.CheckError(err)
	requestData["tree"] = requestDataTree
	requestDataDeleteTree, err := json.Marshal(requestData["deleteTree"])
	common.CheckError(err)
	requestData["deleteTree"] = requestDataDeleteTree
	//返回的数据体
	responseData := common.ResponseData{}
	sql.UpdateTree(requestData, &responseData)
	res, _ := json.Marshal(responseData)
	w.Write(res)
}
