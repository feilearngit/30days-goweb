package main

import "testing"

func BenchmarkPrint2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		print2()
	}
}

func BenchmarkGoPrint2(b  *testing.B)  {
	for i := 0; i < b.N; i ++{
		goPrint2()
	}
}
