package sql

import (
	"backend/common"
	"fmt"
)

type User struct {
	UserId     string `json:"userId"`
	UserName   string `json:"userName"`
	Tree       string `json:"tree"`
	DeleteTree string `json:"deleteTree"`
	CreateTime string `json:"createTime"`
}

func GetTree(requestData map[string]interface{}, responseData *common.ResponseData) {
	rows, err := db.Query("Select user_id , user_name, tree, delete_tree, create_time from user "+
		" where user_name = ?", requestData["userName"])
	common.CheckError(err)
	defer rows.Close()

	var user User
	for rows.Next() {
		err = rows.Scan(&user.UserId, &user.UserName, &user.Tree, &user.DeleteTree, &user.CreateTime)
		common.CheckError(err)
	}
	// fmt.Println(user)

	responseData.Object = user
	responseData.Message = "成功获取"
	responseData.Status = 200
}

func InsertTree(requestData map[string]interface{}, responseData *common.ResponseData) {

	res, err := db.Exec("insert into user (user_name, tree, delete_tree, create_time)"+
		"values(?,?,?,now())", requestData["userName"], requestData["tree"], requestData["deleteTree"])
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

	res, err := db.Exec("update user set user_name = ? , tree = ? , delete_tree = ? where user_id = ?",
		requestData["userName"], requestData["tree"], requestData["deleteTree"], requestData["userId"])
	common.CheckError(err)

	//影响行数
	rowsaffected, err := res.RowsAffected()
	common.CheckError(err)
	fmt.Println("RowsAffected:", rowsaffected)

	responseData.Object = requestData
	responseData.Message = "成功更改"
	responseData.Status = 200
}
