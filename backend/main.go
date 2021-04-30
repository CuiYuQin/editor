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
	mux.HandleFunc("/gettree", controllers.GetTree)

	//端口
	fmt.Println("Web服务器启动...端口:7000")
	err = http.ListenAndServe(":7000", mux)
	if err != nil {
		panic(err)
	}
}
