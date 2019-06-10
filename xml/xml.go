package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

const flag int = 1

type Post struct {
	XMLName  xml.Name  `xml:"post"`
	Id       string    `xml:"id,attr"`
	Content  string    `xml:"content"`
	Author   Author    `xml:"author"`
	Xml      string    `xml:",innerxml"`
	Comments []Comment `xml:"comments>comment"` //跳跃式访问
}

type Comment struct {
	Id      string `xml:"id,attr"`
	Content string `xml:"content"`
	Author  Author `xml:"author"`
}

type Author struct {
	Id   string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

func createXML() {
	post := Post{
		Id:      "1",
		Content: "Hello Golang",
		Author: Author{
			Id:   "2",
			Name: "Alice",
		},
	}

	xmlFile, err := os.Create("create.xml")
	if err != nil {
		fmt.Println("Error creating xml file")
		return
	}
	encoder := xml.NewEncoder(xmlFile)
	encoder.Indent("", "\t")
	err = encoder.Encode(&post)
	if err != nil {
		fmt.Println("Error encoding xml file")
		return
	}
}

func main() {

	createXML()

	xmlFile, err := os.Open("post.xml")
	if err != nil {
		fmt.Println("Error opening XML file:", err)
		return
	}
	defer xmlFile.Close()

	if flag == 0 {
		xmlData, err := ioutil.ReadAll(xmlFile)
		if err != nil {
			fmt.Println("Error reading XML data:", err)
			return
		}
		var post Post
		xml.Unmarshal(xmlData, &post) //将XML结构解封到结构里面，适合小型XML文件
		fmt.Println(post.XMLName)
		fmt.Println(post.Id)
		fmt.Println(post.Content)
		fmt.Println(post.Author.Name)
		fmt.Println(post.Xml)
	} else if flag == 1 {
		//fmt.Println("flag is 1")
		decoder := xml.NewDecoder(xmlFile)
		for {
			t, err := decoder.Token()
			if err == io.EOF {
				//fmt.Println("err = io.EOF")
				break
			}
			if err != nil {
				fmt.Println("Error decoding XML into tokens:", err)
				return
			}

			//fmt.Println("err != io.EOF")
			switch se := t.(type) {
			case xml.StartElement:
				if se.Name.Local == "comment" {
					var comment Comment
					//fmt.Println("strat decode")
					decoder.DecodeElement(&comment, &se)
					fmt.Println(comment.Id)
					fmt.Println(comment.Author)
					fmt.Println(comment.Content)
				}
			}
		}

	}

}
