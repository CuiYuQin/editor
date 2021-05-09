package main

import (
	"backend/controllers"
	"backend/sql"
	"fmt"
	"net/http"
)

func main() {

	err := sql.Initdb()
	if err != nil {
		fmt.Printf("init db failed,err:%v\n", err)
		return
	}
	fmt.Println("DB connected!")

	mux := http.NewServeMux()
	//路由
	//login
	mux.HandleFunc("/login", controllers.Login)
	//tree // user
	mux.HandleFunc("/gettree", controllers.GetTree)
	mux.HandleFunc("/inserttree", controllers.InsertTree)
	mux.HandleFunc("/updatetree", controllers.UpdateTree)
	//ebook
	mux.HandleFunc("/getebook", controllers.GetEbook)
	mux.HandleFunc("/insertebook", controllers.InsertEbook)
	mux.HandleFunc("/updateebook", controllers.UpdateEbook)
	mux.HandleFunc("/updateebooktitle", controllers.UpdateEbookTitle)

	//ggb
	mux.HandleFunc("/ggb", controllers.Ggb)
	mux.Handle("/", http.FileServer(http.Dir("./")))

	//端口
	fmt.Println("Web服务器启动...端口:7000")
	err = http.ListenAndServe(":7000", mux)
	if err != nil {
		panic(err)
	}
}
