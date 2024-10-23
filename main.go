package main

import (
	"FileStore/handler"
	"fmt"
	"net/http"
)

func main() {
	// 注册第一个路由
	http.HandleFunc("/file/upload", handler.UpLoadHandler)
	// 注册第二个路由
	http.HandleFunc("/file/upload/suc", handler.UploadSucHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Println("Failed to start server %s", err.Error())
		return
	}
}
