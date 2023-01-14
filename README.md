# linkTree-server
一个智能家居私有化管理系统，mqtt的消息中间件使用emqx

## Prerequisites
* It supports go 1.13+

## Install
```shell script
go get github.com/eclipse/paho.mqtt.golang
go get github.com/gorilla/websocket
go get golang.org/x/net/proxy
```

## Run
```shell script
go run main.go start -c ./ -l ./logs
```

## build
```shell script
goreleaser --rm-dist --snapshot -o ./param/dist/
```
```shell script
gox -os "linux windows" -arch amd64 -verbose -output "./param/dist/{{.Dir}}_{{.OS}}_{{.Arch}}"
```
