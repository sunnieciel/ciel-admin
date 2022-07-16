# free key admin 自由钥匙 - 个人的后台管理系统

## 框架追求

用越少的技术、越简单的方式来实现一个面向于web独立开发者的后台管理系统。另外不妨将此作为自己的秘密基地，记录自己的事情，记下自己喜欢的格言或语录等等。

## 框架说明

free key admin(自由钥匙) 一个V2EX风格go语言写的个人后台管理系统，提供了权限管理、CURD文件生成、本地文件上传等基本功能。采用`服务端渲染`，整个项目一键启动。

使用到的技术:

- 后端: goframe v2.0.6
- 前端: html,css3,javascript, jquery-3.6.0,vue3,element-plus(vue相关的框架只在少量的页面中有使用.)

## 实现的基本功能

- 角色权限管理 ok
- 字典管理 ok
- 本地文件上传 ok
- CURD文件生成 ok
- 记事本功能 ok
- 常用网站导航 ok
- WebSocket后台聊天功能 todo
- 常用工具集合 todo

这是菜单页面 （图标选择的是哆啦A梦 ）![](./documents/home1.png) ![](./documents/home2.png)

> 更多浏览在页末

## 环境安装

1. golang环境 安装
2. mysql 安装
3. redis 安装
4. gf 安装  https://goframe.org/pages/viewpage.action?pageId=1115782

## 快速开始

1. 下载项目 https://github.com/1211ciel/ciel-admin/releases/tag/v.1.0.9
2. 创建数据库并导入项目`resource`目录下的`ciel.sql`
3. 修改`config.yaml`文件配置:`mysql`,`redis` 将数据库信息替换成您自己的
4. 启动项目：到项目根目录下执行 `gf run main.go`
5. 浏览器打开 http://localhost:1211/login  默认用户名 `ciel` 密码 `1`

## 项目结构

```text
├── api
├── internal
│   ├── cmd   入口指令
│   ├── const
│   ├── controller
│   ├── dao
│   ├── model
│   └── service
├── manifest
├── resource
├── utility  //通用工具
├── go.mod
└── main.go 
```

详情见：https://goframe.org/pages/viewpage.action?pageId=30740166

## css使用

### `.tag`

`.tag-info` `.tag-success` `.tag-priamry` `.tag-wanrnig` `.tag-brown` `.tag-purple` `.tag-danger`

![](./documents/tag.png) <br> ![](./documents/tag2.png)

### 更多class

https://github.com/1211ciel/ciel-admin/blob/master/documents/class.md
![](./documents/class1.png) <br>
![](./documents/class2.png)

## 反馈

使用过程中如果遇见什么问题，可以发issue或发邮件1211sciel@gmail.com 我看到后会及时回复的。

## v.1.1.0 更新计划(进行中)

目前一直在修改、更新、整理，相关改变会比较大。

**前端**

- 将vue去掉，采用全服务端渲染方式。（做自己小项目使用的技术或框架,自然是越简单,越少越好。）
- css整理为3个文件
    - base.css 对结构进行描述
    - white.css white模式色彩描述
    - dark.css dark模式色彩描述

**后端**

- 随着goframe更新到v2.0.6 框架同时进行了更新。
- `service`层使用包方式进行划分。

> `service`层使用包方式进行划分，需要定义好各个模块的依赖关系。每个模块做自己的逻辑事情，不可循环依赖。
> 查询事务都交给`dao`层处理，`service`层只做自己分内的逻辑业务。
> `service` 层之间进行调用或业务增加，请确认好依赖关系，让各模块做好自己分内的事情，防止关系混乱，导致循环依赖。

## 更多预览

导航页面  ![](./documents/home5.png) ![](./documents/home6.png)
备忘录工具页面 ![](./documents/home3.png) ![](./documents/home4.png)
