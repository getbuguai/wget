package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	// 是否跟着参数
	var args = os.Args
	if len(args) < 2 {
		fmt.Println("下载了一个锤子 。。。 ")
		return
	}

	// 当前路径
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("执行命令的地方我找不到 。。。 ")
		panic(err)
	}

	// 检查链接
	fileURL := os.Args[1]
	if !IsUrl(fileURL) {
		fmt.Println("链接看不懂 。。。 ")
		return
	}

	// 获取文件保存路径
	savePath := filepath.Join(dir, GetUrlFileName(fileURL))

	// 下载路径
	err = DownloadFile(fileURL, savePath)
	if err != nil {
		fmt.Println("能力有限 。。。 ")
		panic(err)
	}
	fmt.Println("下载完成 。。。 ")
}

// IsUrl 链接是否是 http 的请求
func IsUrl(urllink string) bool {
	_, err := url.Parse(urllink)
	if err != nil {
		return false
	}
	return true
}

// DownloadFile 下载 链接文件 并保存到指定的路径
func DownloadFile(fileURL, savePath string) error {

	res, err := http.Get(fileURL)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return errors.New("获取不到信息 。。。 ")
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(savePath, body, 0644)
}

// GetUrlFileName 获取下载文件的 文件名称，没有名称的以时间为结果
func GetUrlFileName(urlFile string) string {

	u, err := url.Parse(urlFile)
	// fmt.Println(u.Path)
	// len(u.Path) < 2 域名后跟着 ‘/’ 会显示 为 Path，所以要去除 /
	if err != nil || len(u.Path) < 2 {

		return time.Now().Format("2006-01-02T15_04_05")
		//TODO return time.Now().Format("2006-01-02T15_04_05.html")
	}

	i := strings.LastIndexByte(u.Path, '/')
	if i == -1 {
		return u.Path
	}
	return u.Path[i+1:]
	// strings.Contains(u.Path, "/")

}
