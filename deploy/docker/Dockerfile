FROM centos:7

MAINTAINER DL-NEST<dlnestmit@gmail.com>

ENV LANG=en_US.utf8

RUN set -ex \
    && echo "LANG=en_US.utf8" > /etc/locale.conf

COPY ./linktree-ctl /home/linktree/linktree-ctl
# 工作区
WORKDIR /home/linktree

VOLUME /home/linktree
# 端口
EXPOSE 5523
# 启动运行
CMD ["./linktree-ctl","start"]