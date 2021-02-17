package main

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
)

// FilterUrlMap 过滤网站 的 列表
var FilterUrlMap map[string]UrlFilterInterface
var WgetPwd string

func init() {
	if FilterUrlMap == nil {
		FilterUrlMap = map[string]UrlFilterInterface{}
	}

	var err error
	// 当前路径
	WgetPwd, err = os.Getwd()
	if err != nil {
		fmt.Println("执行命令的地方我找不到 。。。 ")
		panic(err)
	}
}

// UrlFilterInterface 特定网站的下载方式接口
type UrlFilterInterface interface {
	// GetSavePath 文件保存的路径包含文件名
	GetSavePath(url string) string
	// FilterHostName 过滤的域名的 HOSTS
	FilterHostName() string
	// ClientDo 获取下载的文件
	ClientDo(ctx context.Context, url string) (body []byte, err error)
	// FilterUrl url 的处理，如 a.com --> b.com
	FilterUrl(oldURL string) (newURL string, err error)
	// 完成以上 接口之后 使用 init 函数 在 FilterUrlMap 中进行注册
}

// DoWget 开始执行
func DoWget(ctx context.Context, urlLink string) error {

	if FilterUrlMap == nil {
		panic("初始化失败 。。。 ")
	}

	if len(urlLink) == 0 {
		return errors.New("没有输入 地址信息 ")
	}

	uu, err := url.Parse(urlLink)
	if err != nil {
		return err
	}

	fWget, ok := FilterUrlMap[uu.Host].(UrlFilterInterface)
	if !ok {
		fWget = FilterUrlMap[DefaultWgetHostName]
	}

	savePath := fWget.GetSavePath(urlLink)

	urlLink, err = fWget.FilterUrl(urlLink)
	if err != nil {
		return err
	}

	body, err := fWget.ClientDo(ctx, urlLink)
	if err != nil {
		return err
	}

	err = os.MkdirAll(filepath.Dir(savePath), 0644)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(savePath, body, 0644)

}
