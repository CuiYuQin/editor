package sql

import (
	"backend/common"
	"fmt"
)

type User struct {
	Id     string `json:"id"`
	UserName   string `json:"userName"`
	PassWord   string `json:"PassWord"`
	Tree       string `json:"tree"`
	DeleteTree string `json:"deleteTree"`
	CreateTime string `json:"createTime"`
}

func Login(requestData map[string]interface{}, responseData *common.ResponseData) {
	rows, err := db.Query("Select id , username, password, tree, delete_tree, create_time from user "+
		" where username = ?", requestData["userName"])
	common.CheckError(err)
	defer rows.Close()

	var user User
	for rows.Next() {
		err = rows.Scan(&user.Id, &user.UserName, &user.PassWord, &user.Tree, &user.DeleteTree, &user.CreateTime)
		common.CheckError(err)
	}
	// fmt.Println(user)

	responseData.Object = user
	if user.Id != "" {
		if user.PassWord == requestData["PassWord"] {
			responseData.Message = "成功登录"
			responseData.Status = 200
			return
		}
		responseData.Message = "密码错误，请重新登录！"
		responseData.Status = 300
		return
	}
	responseData.Message = "不存在这个账号，请先注册！"
	responseData.Status = 400

}

func GetTree(requestData map[string]interface{}, responseData *common.ResponseData) {
	rows, err := db.Query("Select id , username, tree, delete_tree, create_time from user "+
		" where username = ?", requestData["userName"])
	common.CheckError(err)
	defer rows.Close()

	var user User
	for rows.Next() {
		err = rows.Scan(&user.Id, &user.UserName, &user.Tree, &user.DeleteTree, &user.CreateTime)
		common.CheckError(err)
	}
	// fmt.Println(user)

	responseData.Object = user
	responseData.Message = "成功获取"
	responseData.Status = 200
}

func InsertTree(requestData map[string]interface{}, responseData *common.ResponseData) {

	GetTree(requestData,responseData)
	if responseData.Object.(User).Id != "" {
		responseData.Message = "已有账号！"
		responseData.Status = 300
		return
	}

	res, err := db.Exec("insert into user (username, password, types, tree, delete_tree, create_time)"+
		"values(?,?,?,?,?,now())", requestData["userName"], requestData["PassWord"],requestData["types"],requestData["tree"], requestData["deleteTree"])
	common.CheckError(err)

	//插入数据的主键id
	lastInsertID, err := res.LastInsertId()
	common.CheckError(err)
	fmt.Println("LastInsertID:", lastInsertID)

	//影响行数
	rowsaffected, err := res.RowsAffected()
	common.CheckError(err)
	fmt.Println("RowsAffected:", rowsaffected)

	GetTree(requestData, responseData)
	responseData.Message = "成功生成"
	responseData.Status = 200
}

func UpdateTree(requestData map[string]interface{}, responseData *common.ResponseData) {

	res, err := db.Exec("update user set username = ? , tree = ? , delete_tree = ? where id = ?",
		requestData["userName"], requestData["tree"], requestData["deleteTree"], requestData["id"])
	common.CheckError(err)

	//影响行数
	rowsaffected, err := res.RowsAffected()
	common.CheckError(err)
	fmt.Println("RowsAffected:", rowsaffected)

	responseData.Object = requestData
	responseData.Message = "成功更改"
	responseData.Status = 200
}
