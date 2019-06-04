package main

import (
	json2 "encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type MyHandler struct {
}

type Post struct {
	User    string
	Threads []string
}

func body(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	fmt.Fprintln(w, string(body))
}

func writeExample(w http.ResponseWriter, r *http.Request) {
	str := `<html>
<head><title>Go Web Programming</title></head>
<body><h1>Hello World</h1></body>
</html>`
	w.Write([]byte(str))
}

func writeHeaderExample(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	fmt.Fprintln(w, "No such service, try next door")
}

func headerExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "https://www.baidu.com")
	w.WriteHeader(302)
}

func process(w http.ResponseWriter, r *http.Request) {
	//r.ParseForm()
	//fmt.Fprintln(w, r.PostForm)
	//使用formfile处理上传文件、
	file, fileHeader, err := r.FormFile("uploaded")
	if err == nil {
		data, err := ioutil.ReadAll(file)
		if err == nil {
			fmt.Fprintln(w, fileHeader)
			fmt.Fprintln(w, string(data))
		}
	}

	//r.ParseMultipartForm(1024) //从multipart编码中取1024字节数据
	//fileHeader := r.MultipartForm.File["uploaded"][0]
	//file, err := fileHeader.Open()
	//if err == nil{
	//	data, err := ioutil.ReadAll(file)
	//	if err == nil{
	//		fmt.Fprintln(w, string(data))
	//	}
	//}
	//fmt.Fprintln(w, r.MultipartForm)
}

func jsonExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	post := &Post{
		User:"feifei jiang",
		Threads:[]string{"first", "second", "third"},
	}
	json, _ := json2.Marshal(post)
	w.Write(json)
}

func headers(w http.ResponseWriter, r *http.Request) {
	h := r.Header.Get("Accept-Encoding")
	fmt.Fprintln(w, h)
}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func main() {
	//handler := MyHandler{}
	server := http.Server{
		Addr: "127.0.0.1:8080",
		//Handler: &handler,
	}
	http.HandleFunc("/body", body)
	http.HandleFunc("/process", process)
	http.HandleFunc("/header", headers)
	http.HandleFunc("/write", writeExample)
	http.HandleFunc("writeheader", writeHeaderExample)
	http.HandleFunc("redirect", headerExample)
	http.HandleFunc("/json", jsonExample)
	//http2.ConfigureServer(&server, &http2.Server{})
	server.ListenAndServe()
	//server.ListenAndServe("cert.pem", "key.pem")
}
