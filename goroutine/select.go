package main

import "fmt"

func callerA(c chan string) {
	c <- "Hello World!"
	close(c)
}

func callerB(c chan string) {
	c <- "Hola Mundo!"
	close(c)
}

func main() {
	a, b := make(chan string), make(chan string)
	go callerA(a)
	go callerB(b)
	var msg string
	ok1, ok2 := true, true
	for i := 0; i < 5; i++ {
		select {
		case msg, ok1 = <-a:	//通道关闭后，ok1将被置为false
			if ok1{
				fmt.Printf("%s from A\n", msg)
			}

		case msg, ok2 = <-b:
			if ok2{
				fmt.Printf("%s from B\n", msg)
			}

		default:
			fmt.Println("default")
		}
	}
}
