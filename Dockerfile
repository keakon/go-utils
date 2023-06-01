FROM hub.yizhisec.com/mirror/ubuntu:18.04

# 安装依赖
RUN apt update && apt install -y ca-certificates && update-ca-certificates \
    && sed -i "s@http://\(archive\|security\).ubuntu.com@https://mirror.yizhisec.com@g" /etc/apt/sources.list \
    && apt update \
    && apt upgrade -y \
    && DEBIAN_FRONTEND=noninteractive apt install -y --no-install-recommends wget tar build-essential pkg-config tzdata \
    && ln -fs /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && dpkg-reconfigure -f noninteractive tzdata

# 安装 golang 环境
RUN wget https://artifactory.yizhisec.com/artifactory/filestore/go/go1.19.4.linux-amd64.tar.gz \
    && tar -C /usr/local -xzf go*.linux-amd64.tar.gz \
    && rm go*.linux-amd64.tar.gz


ENV GOPROXY="https://goproxy.cn"  \
    PATH="${PATH}:/usr/local/go/bin"

# 测试
WORKDIR /utils
COPY . /utils
RUN go test -v -covermode=atomic -v ./... -gcflags=all=-l
