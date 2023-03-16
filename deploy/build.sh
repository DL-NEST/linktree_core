cd ../
gox -os "linux windows" -arch "amd64 arm arm64" -verbose -output "./deploy/build/{{.OS}}_{{.Arch}}/linktree_ctl" linktree_core/cmd/main
gox -os "darwin" -arch "arm64" -verbose -output "./deploy/build/{{.OS}}_{{.Arch}}/linktree_ctl" linktree_core/cmd/main