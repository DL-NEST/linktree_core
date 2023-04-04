# linkTree-core
一个智能家居私有化管理系统，mqtt的消息中间件使用emqx  
项目希望部署简单依赖少所以不考虑使用kafka和InfluxDB，可以考虑插件化添加

## Prerequisites
* It supports go 1.19+

## emqx的连接
### exhook连接设置grpc进程池为1000（添加对速度没有影响）
平均13000条/秒的速率，不使用exhook时6-7w条每秒的速率

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
