package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type Post struct {
	Id       int       `json:"id"`
	Content  string    `json:"content"`
	Author   Author    `json:"author"`
	Comments []Comment `json:"comments"`
}

type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Comment struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

const flag = 1

func createJson()  {
	post := Post{
		Id:1,
		Content:"Hello Json",
		Author:Author{
			Id:2,
			Name:"David",
		},
		Comments:[]Comment{
			{
				Id:3,
				Content:"Hava a good day!",
				Author:"Elly",
			},
			{
				Id:4,
				Content:"See you tomorrow!",
				Author:"Fermat",
			},
		},
	}

	output, err := json.MarshalIndent(&post, "", "\t")
	if err != nil{
		fmt.Println("Error marshalling to JSON:", err)
		return
	}
	err = ioutil.WriteFile("create.json", output, 0644)
	if err != nil{
		fmt.Println("Error writing JSON to file:", err)
		return
	}

	jsonFile, err := os.Create("create2.json")
	if err != nil{
		fmt.Println("Error creating JSON file:", err)
		return
	}
	encoder := json.NewEncoder(jsonFile)
	err = encoder.Encode(&post)
	if err != nil{
		fmt.Println("Error encode JSON file:", err)
		return
	}

}


func main() {
	createJson()

	jsonFile, err := os.Open("post.json")
	if err != nil {
		fmt.Println("Error open json file:", err)
		return
	}
	defer jsonFile.Close()
	if flag == 0{

		//如果JSON数据来源于字符串或者内存，用Unmarshal更好
		jsonData, err := ioutil.ReadAll(jsonFile)
		if err != nil {
			fmt.Println("Error reading JSON data:", err)
			return
		}

		var post Post
		json.Unmarshal(jsonData, &post)
		fmt.Println(post)
	}else if flag == 1{

		//如果JSON数据来自io.Reader流，用Decoder更好
		decoder := json.NewDecoder(jsonFile)
		for{
			var post Post
			err := decoder.Decode(&post)
			if err == io.EOF{
				break
			}
			if err != nil{
				fmt.Println("Error decode JSON data:", err)
				return
			}
			fmt.Println(post)
		}
	}


}
