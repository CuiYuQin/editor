package sql

import (
	"backend/common"
	"fmt"
)

type User struct {
	UserId     string `json:"userId"`
	UserName   string `json:"userName"`
	Tree       string `json:"tree"`
	DelectTree string `json:"delectTree"`
	CreateTime string `json:"createTime"`
}

func GetTree(requestData map[string]interface{}, responseData *common.ResponseData) {
	rows, err := db.Query("Select user_id , user_name, tree, delete_tree, create_time from user "+
		" where user_name = ?", requestData["userName"])
	common.CheckError(err)
	defer rows.Close()

	var user User
	for rows.Next() {
		err = rows.Scan(&user.UserId, &user.UserName, &user.Tree, &user.DelectTree, &user.CreateTime)
		common.CheckError(err)
	}
	fmt.Println(user)
	responseData.Objects = user
	return
}
