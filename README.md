# linkTree-core
一个智能家居私有化管理系统，mqtt的消息中间件使用emqx

## Prerequisites
* It supports go 1.19+

## 服务层的边界
api遵循mvc的架构
- router层是对api的路由注册
- controller层是对请求数据的处理和对数据的返回进行处理
- service是对接口的操作返回到control层
- dao层是定义获取数据的接口


## emqx的连接
### exhook连接设置grpc进程池为1000（添加对速度没有影响）
redis存储速度为  
平均10000条每秒的速率，不使用exhook时6-7w条每秒的速率


## Install
```shell script
go mod tidy
```

## Run
```shell script
go run main.go start -c ./ -l ./logs
```

## build
使用gox对项目进行构建分包
```shell script
gox -os "linux windows darwin" -arch "amd64 arm arm64" -verbose -output "./deploy/build/{{.OS}}_{{.Arch}}/linktree_ctl"
```
