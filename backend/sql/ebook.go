package sql

import (
	"backend/common"
	"fmt"
)

type Ebook struct {
	EbookId     string `json:"ebookId"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	CreateTime  string `json:"createTime"`
	Content     string `json:"content"`
	ContentHtml string `json:"contentHtml"`
}

func GetEbook(requestData map[string]interface{}, responseData *common.ResponseData) {
	rows, err := db.Query("Select ebookId , title, author, createTime, content, contentHtml from ebook "+
		" where ebookId = ?", requestData["ebookId"])
	common.CheckError(err)
	defer rows.Close()

	var ebook Ebook
	for rows.Next() {
		err = rows.Scan(&ebook.EbookId, &ebook.Title, &ebook.Author, &ebook.CreateTime, &ebook.Content, &ebook.ContentHtml)
		common.CheckError(err)
	}
	// fmt.Println(ebook)

	responseData.Object = ebook
	responseData.Message = "成功获取"
	responseData.Status = 200
}

func InsertEbook(requestData map[string]interface{}, responseData *common.ResponseData) {

	res, err := db.Exec("insert into ebook (title, author, createTime, content, contentHtml)"+
		"values(?,?,now(),?,?)", requestData["title"], requestData["author"], requestData["content"], requestData["contentHtml"])
	common.CheckError(err)

	//插入数据的主键id
	lastInsertID, err := res.LastInsertId()
	common.CheckError(err)
	fmt.Println("LastInsertID:", lastInsertID)

	//影响行数
	rowsaffected, err := res.RowsAffected()
	common.CheckError(err)
	fmt.Println("RowsAffected:", rowsaffected)

	requestData["ebookId"] = lastInsertID
	GetEbook(requestData, responseData)
	responseData.Message = "成功生成"
	responseData.Status = 200
}

func UpdateEbook(requestData map[string]interface{}, responseData *common.ResponseData) {

	res, err := db.Exec("update ebook set title = ? , content = ? , contentHtml = ? where ebookId = ?",
		requestData["title"], requestData["content"], requestData["contentHtml"], requestData["ebookId"])
	common.CheckError(err)

	//影响行数
	rowsaffected, err := res.RowsAffected()
	common.CheckError(err)
	fmt.Println("RowsAffected:", rowsaffected)

	responseData.Object = requestData
	responseData.Message = "成功更改"
	responseData.Status = 200
}

func UpdateEbookTitle(requestData map[string]interface{}, responseData *common.ResponseData) {

	res, err := db.Exec("update ebook set title = ? where ebookId = ?",
		requestData["title"], requestData["ebookId"])
	common.CheckError(err)

	//影响行数
	rowsaffected, err := res.RowsAffected()
	common.CheckError(err)
	fmt.Println("RowsAffected:", rowsaffected)

	responseData.Object = requestData
	responseData.Message = "成功更改"
	responseData.Status = 200
}


