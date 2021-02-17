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

type UrlFilterInterface interface {
	GetSavePath(url string) string
	FilterHostName() string
	ClientDo(ctx context.Context, url string) (body []byte, err error)
	FilterUrl(oldURL string) (newURL string, err error)
}

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
