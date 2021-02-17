# wget 
一个简易的 wget 下载小工具，支持所有 http 格式的文件下载，如果不支持，你可以补充一下。

## 背景

windows 上命令行不能下载文件，或者我不知道怎么下载。

于是就写了一个简单的下载工具。

有好的建议可以提出来，[进行改进。](https://github.com/getbuguai/wget/issues)

## 构建

```
git clone https://github.com/getbuguai/wget.git
cd wget && go build 

# 以工具包的形式, 会自动安装到 GOBIN
go get -u github.com/getbuguai/wget

```
## 使用方式

wget https://cdn.jsdelivr.net/gh/getbuguai/flutter-app1/assets/img/git.jpg

## 贡献

代码支持进行拓展，可以为指定网站指定下载方式。

具体的实现就是实现一下接口:

```go
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
```

> 默认情况使用 default_wget 中的下载方式。

## 赞助

<details>
  <summary>支付宝</summary>
   <img src="https://cdn.jsdelivr.net/gh/getbuguai/getbuguai/zhifubao.jpg"
     alt="支付宝收款">
     加载失败访问: https://cdn.jsdelivr.net/gh/getbuguai/getbuguai/zhifubao.jpg
</details>

<details>
    <summary>微信</summary>
     <img class="fit-picture"
     src="https://cdn.jsdelivr.net/gh/getbuguai/getbuguai/weixin.png"
     alt="微信收款">
     加载失败访问: https://cdn.jsdelivr.net/gh/getbuguai/getbuguai/weixin.png
</details>

> 如有问题可以进行备注  
> 默认为请作者喝杯奶茶  
> 感谢投喂 !!! 

## 找到我：这个程序不太乖 

- BiliBili 主页: [https://space.bilibili.com/278413353 ](https://space.bilibili.com/278413353)
- GitHub 主页: [https://github.com/getbuguai](https://github.com/getbuguai) 
- 码云 主页: [https://gitee.com/getbuguai](https://gitee.com/getbuguai) 
- QQ群: [ 不乖的程序交流群: [487090042](https://qm.qq.com/cgi-bin/qm/qr?k=4E_QbhCpe0O2QVPU_UFi-AFMLOmxpXrw&jump_from=webapi) ]