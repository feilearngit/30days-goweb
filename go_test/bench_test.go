package main

import "testing"

//go test -run x -bench
//run 接不存在的功能测试名字，所以的功能测试都将被忽略
func BenchmarkDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		decode("post.json")
	}
}

func BenchmarkUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		unmarshal("post.json")
	}
}
