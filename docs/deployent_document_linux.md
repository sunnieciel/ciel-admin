## 服务器

- 系统： Ubuntu 20.04.5 LTS
- cup： 16核
- 内存：8G
- SSD： 512G

## 环境安装

1. nginx
2. mysql 8.0
3. PM2管理器
4. Linux工具箱

### 5. golang 安装

ssh 登录远程服务器

下载golang

```
wget -c https://dl.google.com/go/go1.19.1.linux-amd64.tar.gz -O - | sudo tar -xz -C /usr/local   
```

配置环境变量

```
vim /etc/profile

export PATH=$PATH:/usr/local/go/bin

source /etc/profile

```

测试一下golang是否安装好

 ```
go version
go version go1.19.1 linux/amd64
```

### 6. gf 安装

下载gf

```
wget https://github.com/gogf/gf/archive/refs/tags/v2.3.3.zip

```

解压gf

```
unzip v2.3.3.zip
```

制作gf工具包

```
cd gf-2.3.3/cmd/gf/
go build
```

制作好gf后，将gf 存放到有配置环境变量的文件夹，方便后面使用

```
cp gf /usr/local/bin/
```

测试gf 是否安装成功

```
cd 

gf version

GoFrame CLI Tool v2.3.3, https://goframe.org
GoFrame Version: cannot find go.mod
CLI Installed At: /usr/local/bin/gf
Current is a custom installed version, no installation information.
```

## 后台项目部署

1. ssh 登录服务器创建项目文件夹
2. 使用 goland tools Deploment 上传后台项目
3. 创建数据库
4. 导入sql
5. 修改项目配置文件，改成服务器数据库信息

### 6.进入项目文件夹后运行项目

```
gf run main.go
```

如果发现没有什么问题，那就说明项目配置成功了。

退出 使用 nohup 后台运行项目

```
nohup gf run main.go &
```

看看项目是否已经运行

```
netstat -tunlp

tcp6       0      0 :::2033                 :::*                    LISTEN      110423/./main  
```

看看已经后台运行成功了,下面开始部署前台吧

### 7. 宝塔开启后台接口的端口

- 可以把 mysql 3306
- 项目端口 2033
- 前台端口 3000

一起打开吧

## 前台项目部署

1. ssh 登录服务器创建项目文件夹
2. 使用 goland tools Deploment 上传前台项目

### 3. 运行后台项目

进入前台项目文件夹

安装项目

```
npm install
```

构建项目

```
npm run build
```

使用 pm2 来管理程序

```
pm2 start npm --name love -- start
```

进入服务器的ip端口看看吧

```
http://XXX.XX.XX.XXX:3000/backend/sys/login
```

到此 前后台的项目部署完成

## 其他服务器事件

1. 设置时区
2. 添加定时任务，每日同步时间
3. 添加定时任务，每日备份数据库
