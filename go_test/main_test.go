package main

import (
	"testing"
	"time"
)

/*
1、使用go test -v -cover来测试
2、可是有Skip跳过单个测试用例
3、使用-short跳过耗时过长的函数
4、使用-parallel进行并行测试
*/

func TestLongRunningTest(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping long-running test in short mode")
	}
	time.Sleep(10 * time.Second)
}

func TestDecode(t *testing.T) {
	post, err := decode("post.json")
	if err != nil {
		t.Error(err)
	}
	if post.Id != 1 {
		t.Error("Wrong id, was expecting 1 but got", post.Id)
	}
	if post.Content != "Hello Json!" {
		t.Error("Wrong content, was expecting 'Hello Json' but got", post.Content)
	}
}

func TestEncode(t *testing.T) {
	t.Skip("Skipping encoding for now")
}
