package main

import (
	"context"
	"log"
	"os"
)

func main() {
	// 是否跟着参数
	var args = os.Args
	if len(args) < 2 {
		log.Fatal("下载了一个锤子 。。。 ")
	}

	// 链接
	fileURL := args[1]

	err := DoWget(context.Background(), fileURL)
	if err != nil {
		panic(err)
	}
}
