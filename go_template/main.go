package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("template.html")
	rand.Seed(time.Now().Unix())
	t.Execute(w, rand.Intn(10) > 5) //返回随机数是否大于5
}

func rangeX(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("range.html")
	msg := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	t.Execute(w, msg)
}

func withX(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("with.html")
	t.Execute(w, "Hello")
}

func includeX(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("include.html", "t2.html")
	t.Execute(w, "Hello World")
}

func formatDate(t time.Time) string {
	layout := "2006-01-02"
	return t.Format(layout)
}

func pipeX(w http.ResponseWriter, r *http.Request) {
	funcMap := template.FuncMap{"fdate": formatDate}
	t := template.New("pipe.html").Funcs(funcMap)
	t, _ = t.ParseFiles("pipe.html")
	t.Execute(w, time.Now())
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/pipe", pipeX)
	http.HandleFunc("/include", includeX)
	http.HandleFunc("/with", withX)
	http.HandleFunc("/range", rangeX)
	http.HandleFunc("/template", process)
	server.ListenAndServe()
}
