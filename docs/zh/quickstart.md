# 快速入门 Quick Start

## 安装
```
go get github.com/headwindfly/clevergo
```

## Hello World
```
package main

import (
	"log"
	"github.com/headwindfly/clevergo"
)


func helloCleverGo(ctx *clevergo.Context) {
	ctx.SetBodyString("Hello CleverGo.")
}

func main() {
	// 创建路由器实例
	router := clevergo.NewRouter()

	// 注册路由处理器
	router.GET("/", clevergo.HandlerFunc(helloCleverGo))

	// 启动 Server
	log.Fatal(clevergo.ListenAndServe(":8080", router.Handler))
}
```
然后访问 http://127.0.0.1:8080 即可看到“Hello CleverGo.”字样。

## 使用案例
[Examples](/examples)