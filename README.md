## genshowdoc 一个生成 showdoc 文档的工具

genshowdoc 基于 go-zero 开发，根据定义的 .api 文件，一行命令就可以生成 showdoc 文档

#### 安装

go 1.16 以下使用
```sh
go get -u github.com/MasterJoyHunan/genshowdoc
```
go 1.16 及以上使用
```sh
go install github.com/MasterJoyHunan/genshowdoc@v1.0.0
```

#### 在项目下定义 you-app.api 文件

[api语法指南](https://go-zero.dev/cn/docs/design/grammar)

you-app.api 文件内容示例

```api
syntax = "v1"

info(
	title: "some app"
)

type bookRequest {
    Name string `json:"name"` // 姓名
    Age int `json:"age"`      // 年龄
}

type bookResponse {
    Code int `json:"code"` // 业务码
    Msg string `json:"msg"` // 业务消息
}

@server(
    jwt: Auth
    group: book
    middleware: SomeMiddleware,CorsMiddleware
    prefix: /v1
)

service someapp {
    @doc "获取所有书本信息"
    @handler getBookList
    get /book (bookRequest) returns (bookResponse)

    @doc "获取书本信息"
    @handler getBook
    get /book/:id (bookRequest) returns (bookResponse)

    @doc "添加书本信息"
    @handler addBook
    post /book (bookRequest) returns (bookResponse)

    @doc "获取书本信息"
    @handler editBook
    put /book/:id (bookRequest) returns (bookResponse)
}
```

#### 生成文档

```sh
genshowdoc --api="http://host:port/server/index.php?s=/api/item/updateByApi" --key="you_key" --token="you_token" xxx.api
```

参数描述

* `--api` api 请求地址
* `--key` 项目的 key
* `--token` 项目的 token

### 其他

如果觉得该项目对你有所帮助，请不要吝啬你的小手，帮忙点个 stars

如果对本项目有更好的建议或意见，欢迎提交 pr / issues，或者联系本人 tanwuyang88@gmail.com

再次感谢 [go-zero](https://github.com/zeromicro/go-zero)

### 协议

[MIT](https://github.com/MasterJoyHunan/genshowdoc/blob/master/LICENSE)