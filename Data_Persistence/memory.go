package main

import "fmt"

var PostById map[int]*Post
var PostsByAuthor map[string][]*Post

func store(post Post) {
	PostById[post.Id] = &post
	PostsByAuthor[post.Author] = append(PostsByAuthor[post.Author], &post)
}

func main() {
	PostById = make(map[int]*Post)
	PostsByAuthor = make(map[string][]*Post)

	post1 := Post{Id: 1, Content: "Hello World!", Author: "jiangjiang"}
	post2 := Post{Id: 2, Content: "Bonjour Monde!", Author: "xiongxiong"}
	post3 := Post{Id: 3, Content: "Hola Mundo!", Author: "feifei"}
	post4 := Post{Id: 4, Content: "Greetings Earthlings!", Author: "feifei"}

	store(post1)
	store(post2)
	store(post3)
	store(post4)

	fmt.Println(PostById[1])
	fmt.Println(PostById[2])

	for _, post := range PostsByAuthor["feifei"] {
		fmt.Println(post)
	}
	for _, post := range PostsByAuthor["jiangjiang"] {
		fmt.Println(post)
	}
}
