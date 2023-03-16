echo "build executableFile"
go env -w GO111MODULE=on && go env -w GOPROXY=https://goproxy.cn,direct && go env -w CGO_ENABLED=0 && go env  && go mod tidy
go build -o ./linktree-ctl linktree_core/cmd/main
echo "build image"
docker build -t dlnest/linktree_core .
echo "delete executableFile"
rm linktree-ctl