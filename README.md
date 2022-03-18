# go-http-routers-testing

Go 路由的性能测试工具

<https://caixw.github.io/go-http-routers-testing/>

## 安装

```shell
go get github.com/caixw/go-http-routers-testing
```

### 运行测试

main.go 可以有以下参数：

- test.benchtime 指定每个测试的运行时间；

```shell
go run main.go -test.benchtime=5s
```

之后通过以下代码编译成 html：

```shell
go run html.go
```

如果要本地查看数据，请运行：

```shell
go run ./serve.go
```

## 协作

与 api 定义有的内容定义在 apis 目录下，如果不同的风格需求可以在此目录下添加即可;

路由定义在 routers 目录下，新路由只需要在此目录下定义 Load 方法即可；

不要提交 /docs/data 下的数据，该数据由本人统一产生。

## 版权

本项目采用 [MIT](https://opensource.org/licenses/MIT) 开源授权许可证，完整的授权说明可在 [LICENSE](LICENSE) 文件中找到。
