package main

import (
	"testing"
	"time"
)

//go test -v -short -parallel 3   ---3表示最多并行3个测试用例
func TestParallel_1(t *testing.T) {
	t.Parallel()
	time.Sleep(1 * time.Second) //耗时1s
}

func TestParallel_2(t *testing.T) {
	t.Parallel()
	time.Sleep(2 * time.Second) //耗时2s

}

func TestParallel_3(t *testing.T) {
	t.Parallel()
	time.Sleep(3 * time.Second) //耗时3秒
}
