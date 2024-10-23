package handler

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// 实现上传文件的接口
// 这里注意，首字母大写才可以在包外调用
func UpLoadHandler(w http.ResponseWriter, r *http.Request) {
	// 如果请求是Get, 那么返回页面
	if r.Method == http.MethodGet {
		// 返回上传的html页面
		file, err := os.Open("./static/view/index.html")
		if err != nil { // 错误处理
			http.Error(w, err.Error(), http.StatusInternalServerError) // http包里面的错误处理
			return
		}
		defer file.Close() // 记得关闭文件
		// 执行到这里说明上传已经成功
		// 写入w
		_, err = io.Copy(w, file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	} else if r.Method == http.MethodPost {
		// 上传文件，存储到本地目录
		// 实现上传文件的请求
		file, header, err := r.FormFile("file") // 这个函数有三个返回值， 文件数据流，文件的元数据、err
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file.Close() // 关闭文件
		// 创建一个新的文件，存储在服务端本地
		newFile, err := os.Create("./temp/" + header.Filename)
		if err != nil {
			fmt.Println("Failed to upload the file:%s", header.Filename)
			return
		}
		defer newFile.Close()
		// 把上传的数据拷贝进来
		_, err = io.Copy(newFile, file)
		if err != nil {
			fmt.Println("Failed to upload the file:%s", header.Filename)
			return
		}
		// 上传成功后实现重定向页面
		http.Redirect(w, r, "/file/upload/suc", http.StatusFound) // 3XX代表重定向

	}
}

func UploadSucHandler(w http.ResponseWriter, r *http.Request) {
	_, err := io.WriteString(w, "Upload finished")
	if err != nil {
		fmt.Println("Err happen:%s", err.Error())
		return
	}
}
