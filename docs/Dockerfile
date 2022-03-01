FROM golang:alpine

## docker镜像构建
## [docker安装使用请参考官方](https://www.docker.com/)
## 1. 根据情况修改Dockerfile文件
## 2. 在项目根目录下使用命令docker build -t <你要出的进行名>:<版本号> .

# 设置环境变量
ENV GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GOPROXY="https://goproxy.cn,direct"
# 移动到工作目录：/home/cutego 这个目录 是项目代码 放在linux上
WORKDIR /home/cutego
# 将代码复制到容器中
COPY .. .
# 将我们的代码编译成二进制可执行文件  可执行文件名为 app
RUN go build -o app .
# 移动到用于存放生成的二进制文件的 /dist 目录
WORKDIR /dist
# 创建日志文件目录
RUN mkdir /home/logs
# 创建文件存储目录
RUN mkdir /home/upload
# 将二进制文件从 /home/cutego 目录复制到这里
RUN cp /home/cutego/app .
# 将配置文件放入与app同级目录
RUN cp /home/cutego/application*.yml .
# 声明服务端口
EXPOSE 21366
# 启动容器时运行的命令
CMD ["/dist/app"]