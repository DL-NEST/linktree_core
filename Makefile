
#VERSION		?= $(shell git describe --tags --always --dirty --match=v* 2> /dev/null || cat $(PWD)/.version 2> /dev/null || echo v0)

build:
	gox -os "linux windows" -arch "amd64 arm arm64" -verbose -output "./deploy/build/{{.OS}}_{{.Arch}}/linktree_ctl" linktree_core/cmd/main
	gox -os "darwin" -arch "arm64" -verbose -output "./deploy/build/{{.OS}}_{{.Arch}}/linktree_ctl" linktree_core/cmd/main

build_docker:
	@cd deploy/docker \
	&& go env -w GO111MODULE=on && go env -w GOPROXY=https://goproxy.cn,direct \
	&& go env -w CGO_ENABLED=0 && go env  && go mod tidy \
	&& go build -o ./linktree-ctl linktree_core \
	&& docker build -t dlnest/linktree_core .

generate-swagger:
	swag init