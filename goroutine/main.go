package main

import (
	"fmt"
	"time"
)

//并发是同时处理多项任务，并行是同时执行多项任务
//通道 channel
//ch := make(chan int)  //无缓冲区通道
//ch := make(chan int, 10) //有缓冲区通道

func printNumbers1() {
	for i := 0; i < 10; i++ {
		//fmt.Printf("%d ", i)
	}
}

func printLetters1() {
	for i := 'A'; i < 'A'+10; i++ {
		//fmt.Printf("%c ", i)
	}
}

func printNumbers2(){
	for i := 0; i < 10; i ++{
		time.Sleep(1*time.Microsecond)
		//fmt.Printf("%d ", i)
	}
}

func printLetters2(){
	for i := 'A'; i < 'A' + 10; i++{
		time.Sleep(1*time.Microsecond)
		//fmt.Printf("%c ", i)
	}
}

func printNumbers3(w chan bool){ //bool类型通道
	for i := 0; i < 10; i ++{
		time.Sleep(1*time.Microsecond)
		fmt.Printf("%d ", i)
	}
	w <- true
}

func printLetters3(w chan bool){
	for i := 'A'; i < 'A' + 10; i++{
		time.Sleep(1*time.Microsecond)
		fmt.Printf("%c ", i)
	}
	w <- true
}

func print1() {
	printNumbers1()
	printLetters1()
}

func goPrint1() {
	go printNumbers1()
	go printLetters1()
}

func print2(){
	printNumbers2()
	printLetters2()
}

func goPrint2(){
	go printNumbers2()
	go printLetters2()
}

func main() {
	w1, w2 := make(chan bool), make(chan bool)
	go printNumbers3(w1)
	go printLetters3(w2)
	<- w1
	<- w2
}
