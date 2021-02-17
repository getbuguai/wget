package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func Test_Main(t *testing.T) {
	fileURL := ""
	dir := "./temp"
	if !IsUrl(fileURL) {
		fmt.Println("链接看不懂 。。。 ")
		return
	}

	// 获取文件保存路径
	savePath := filepath.Join(dir, GetUrlFileName(fileURL))

	// 下载路径
	err := DownloadFile(fileURL, savePath)
	if err != nil {
		fmt.Println("能力有限 。。。 ")
		panic(err)
	}
	fmt.Println("下载完成 。。。 ")
}

func Test_makedirAll(t *testing.T) {
	dir := "./temp/aaa/aaa"
	err := os.MkdirAll(dir, 0644)
	if err != nil {
		panic(err)
	}

	t.Log("ok ...")
}

func Test_Wget(t *testing.T) {
	fileURL := ""
	WgetPwd = "./temp"
	err := DoWget(context.Background(), fileURL)
	if err != nil {
		panic(err)
	}

}
