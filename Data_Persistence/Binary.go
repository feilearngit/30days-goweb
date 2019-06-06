package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"reflect"
	"time"
)

type Box struct {
	Id     int
	Name   string
	Author string
}

func storeBox(data interface{}, filename string) {
	buffer := new(bytes.Buffer)
	fmt.Println("buffer type is %v", reflect.TypeOf(buffer))
	encoder := gob.NewEncoder(buffer) //转为二进制数据
	err := encoder.Encode(data)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(filename, buffer.Bytes(), 0600)
	if err != nil {
		panic(err)
	}
}

func loadBox(data interface{}, filename string) {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	buffer := bytes.NewBuffer(raw)
	doc := gob.NewDecoder(buffer)	//解码器
	err = doc.Decode(data)
	if err != nil {
		panic(err)
	}
}

func main() {
	box := Box{Id:1, Name:"xxx", Author:"xxx-xxx"}
	storeBox(box, "box1")
	var boxRead Box
	time.Sleep(time.Duration(2) * time.Second)		//休眠2s
	loadBox(&boxRead, "box1") //此处传入指针，用于赋值
	fmt.Println(boxRead)

}
