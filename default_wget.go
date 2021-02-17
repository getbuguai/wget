package main

import (
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"
	"time"
)

func init() {
	if FilterUrlMap == nil {
		FilterUrlMap = make(map[string]UrlFilterInterface)
	}
	FilterUrlMap[DefaultWgetHostName] = DefaultWget{}
}

const DefaultWgetHostName = "default"

type DefaultWget struct {
}

func (wGet DefaultWget) GetSavePath(url string) string {
	return GetSavePathByFileName(GetUrlFileName(url))
}

func (wGet DefaultWget) FilterHostName() string {
	return DefaultWgetHostName
}

func (wGet DefaultWget) ClientDo(ctx context.Context, url string) (body []byte, err error) {
	return DefaultGetByte(ctx, url)
}

func (wGet DefaultWget) FilterUrl(oldURL string) (newURL string, err error) {
	_, err = url.Parse(oldURL)
	return oldURL, err
}

// IsUrl 链接是否是 http 的请求
func IsUrl(urllink string) bool {
	_, err := url.Parse(urllink)
	if err != nil {
		return false
	}
	return true
}

// GetUrlFileName 获取下载文件的 文件名称，没有名称的以时间为结果
func GetUrlFileName(urlFile string) string {

	u, err := url.Parse(urlFile)
	// fmt.Println(u.Path)
	// 解释 len(u.Path) < 2 域名后跟着 ‘/’ 会显示 为 Path，所以要去除 /
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

// GetSavePathByFileName 通过文件的名称获取保存的位置
func GetSavePathByFileName(fName ...string) string {
	l := 1 + len(fName)
	joins := make([]string, l)
	joins[0] = WgetPwd

	for i := range fName {
		joins[i+1] = fName[i]
	}
	return filepath.Join(joins...)
}

// DefaultGetByte 默认的 get 资源
func DefaultGetByte(ctx context.Context, fileURL string) ([]byte, error) {
	res, err := http.Get(fileURL)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, errors.New("获取不到信息 。。。 ")
	}

	return ioutil.ReadAll(res.Body)
}

// DownloadFile 下载 链接文件 并保存到指定的路径
func DownloadFile(fileURL, savePath string) error {

	body, err := DefaultGetByte(context.Background(), fileURL)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(savePath, body, 0644)
}
