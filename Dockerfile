FROM centos

COPY ./param/dist/linktree_server_linux_amd64 /usr/linktree/linktree
COPY ./param/config.yaml /usr/linktree/config.yaml

WORKDIR /usr/linktree

# 端口
EXPOSE 5523

# 运行参数
CMD /usr/linktree/linktree start -c /usr/linktree