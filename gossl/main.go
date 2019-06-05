package main

import (
	"encoding/base64"
	json2 "encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"time"
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

func setCookie(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name:     "first_cookie",
		Value:    "Go Web Programming",
		HttpOnly: true,
	}
	c2 := http.Cookie{
		Name:     "second_cookie",
		Value:    "Manning Publications Co.",
		HttpOnly: true,
	}

	//两种cookie设置方式
	w.Header().Set("Set-Cookie", c1.String())
	http.SetCookie(w, &c2) //此处传入cookie结构的指针

}

func getCookie(w http.ResponseWriter, r *http.Request) {
	c1, err := r.Cookie("first_cookie")
	if err != nil{
		fmt.Fprintln(w, "Cannot get the first cookie")
	}
	cs := r.Cookies()
	fmt.Fprintln(w, c1)
	fmt.Fprintln(w, cs)
	//h := r.Header["Cookie"]
	//fmt.Fprintln(w, h)
}

func jsonExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	post := &Post{
		User:    "feifei jiang",
		Threads: []string{"first", "second", "third"},
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

func setMessage(w http.ResponseWriter, r *http.Request){
	msg := []byte("Hello World!")
	c := http.Cookie{
		Name:"flash",
		Value:base64.URLEncoding.EncodeToString(msg),
	}

	http.SetCookie(w, &c)
}

func showMessage(w http.ResponseWriter, r *http.Request){
	c, err := r.Cookie("flash")
	if err != nil{
		if err == http.ErrNoCookie{
			fmt.Fprintln(w, "No message found")
		}
	}else{
		rc := http.Cookie{
			Name:"flash",
			MaxAge:-1,
			Expires:time.Unix(1,0),
		}
		http.SetCookie(w, &rc)
		val, _ := base64.URLEncoding.DecodeString(c.Value)
		fmt.Fprintln(w, string(val))
	}
}

func _template(w http.ResponseWriter, r *http.Request){
	t, _ := template.ParseFiles("temp.html")
	//msg := []byte("Go Web Programming")

	//msg = append(msg, []byte("Hello"))

	t.Execute(w, "Hello World!")
}

func main() {
	//handler := MyHandler{}
	server := http.Server{
		Addr: "127.0.0.1:8080",
		//Handler: &handler,
	}
	http.HandleFunc("/template", _template)
	http.HandleFunc("/set_message", setMessage)
	http.HandleFunc("/show_message", showMessage)
	http.HandleFunc("/set_cookie", setCookie)
	http.HandleFunc("/get_cookie", getCookie)
	http.HandleFunc("/body", body)
	http.HandleFunc("/process", process)
	http.HandleFunc("/header", headers)
	http.HandleFunc("/write", writeExample)
	http.HandleFunc("/writeheader", writeHeaderExample)
	http.HandleFunc("/redirect", headerExample)
	http.HandleFunc("/json", jsonExample)
	//http2.ConfigureServer(&server, &http2.Server{})
	server.ListenAndServe()
	//server.ListenAndServe("cert.pem", "key.pem")
}
