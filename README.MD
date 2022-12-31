# 🍉 freekey admin 自由钥匙 - 个人的后台管理系统

```
███████╗██████╗ ███████╗███████╗    ██╗  ██╗███████╗██╗   ██╗     █████╗ ██████╗ ███╗   ███╗██╗███╗   ██╗
██╔════╝██╔══██╗██╔════╝██╔════╝    ██║ ██╔╝██╔════╝╚██╗ ██╔╝    ██╔══██╗██╔══██╗████╗ ████║██║████╗  ██║
█████╗  ██████╔╝█████╗  █████╗      █████╔╝ █████╗   ╚████╔╝     ███████║██║  ██║██╔████╔██║██║██╔██╗ ██║
██╔══╝  ██╔══██╗██╔══╝  ██╔══╝      ██╔═██╗ ██╔══╝    ╚██╔╝      ██╔══██║██║  ██║██║╚██╔╝██║██║██║╚██╗██║
██║     ██║  ██║███████╗███████╗    ██║  ██╗███████╗   ██║       ██║  ██║██████╔╝██║ ╚═╝ ██║██║██║ ╚████║
╚═╝     ╚═╝  ╚═╝╚══════╝╚══════╝    ╚═╝  ╚═╝╚══════╝   ╚═╝       ╚═╝  ╚═╝╚═════╝ ╚═╝     ╚═╝╚═╝╚═╝  ╚═══╝

v 1.2
```

## 🍓 追求

用越少的技术、越简单的方式来实现一个面向于web独立开发者的后台管理系统。

## 🍌 说明

freekey admin(自由钥匙) 一个V2EX风格go语言写的个人后台管理系统，提供了权限管理、本地文件上传等基本功能。采用`服务端渲染`，整个项目一键启动。

使用到的技术:

- 后端: goframe v2.2.5, mysql,redis
- 前端: html,css,javascript, jquery-3.6.0

## 🍒 内置功能

### 后台

角色权限管理、字典管理、本地文件上传、用户管理 、用户登录日志、用户钱包、人工充值、人工扣除 、充值类型、账变记录、账变统计、账变报表

这是菜单页面 （图标选择的是哆啦A梦 ）![](./documents/home1.png) ![](./documents/home2.png)

### API

完成一般系统都会涉及的主要接口。 http://localhost:2033/swagger (启动项目可直接访问)

- 系统: 查询所有字典、查询Banner图列表、查询单个字典信息、上传图片 (到本地服务器)
- 用户: 可选头像、用户登录信息、登录、注册、修改头像 、修改昵称、修改密码
- 钱包: 创建充值订单、设置钱包密码、查询钱包信息、查询账变记录、查询账变类型、查询充值记录、查询充值类型、修改钱包密码

### 计划中模块

- 活动模块 todo
- 商品买卖 todo
- 商品主题购买 todo
- 商品二手交易所 todo

---

## 🌸 环境安装

1. golang环境 安装
2. mysql 安装
3. redis 安
4. gf 安装  https://goframe.org/pages/viewpage.action?pageId=1115782

--- 

## ✨️ 快速开始

1. 下载项目 https://github.com/1211ciel/ciel-admin/archive/refs/tags/v1.1.zip
2. 创建数据库并导入项目`resource`目录下的`freekey_system.sql`
3. 修改`config.yaml`文件配置:`mysql`,`redis` 将数据库信息替换成您自己的
4. 启动项目：到项目根目录下执行 `gf run main.go`
5. 浏览器打开 http://localhost:2033/login  默认用户名 `ciel` 密码 `1`

--- 

## 🍀 [CRUD代码生成](https://github.com/1211ciel/ciel-admin/blob/master/documents/gen.md)

--- 

## 🌴 [项目结构](https://goframe.org/pages/viewpage.action?pageId=30740166)

```text
├── api
├── hack
├── internal   
│   ├── cmd        入口指令                         ---👌 "如果没什么事的话，就让我们开始吧。"
│   ├── consts     常量                            --- 🙂 "请相信，我是永远不会改变的~"
│   ├── controller 控制层     提供对外接口           --- 😀 "您需要的所有功能我都可以为您实现,请告诉我吧~"
│   ├── dao        数据查询层  数据库查询操作在此层执行 --- 🤪 "数据库相关的信息我最了解。"
│   ├── logic      逻辑层     业务逻辑在此层进行处理,  --- 😎 "不管你的逻辑有多复杂，交给在下吧～"
│   ├── model      实体层     包含 bo do entity     --- 🤗 "我这里包含了你可能会用到的所有模型实体。"
│   └── service    服务层     为控制层提供服务支持     --- 🙂 "我只衷心为您服务~"
├── main.go                                       --- 🌞 "go!"
├── manifest   // 配置文件                          --- 🌼 "有什么事情可以先告诉我."
├── resource   // 静态或其他资源                     --- 🍀 "我是html 请到这里找来我吧"
└── utility    // 通用工具                          --- 🐶 "如果其他项目也想使用的话，请交给我."

      "我们一直在这里，等你!"
🌞 👌 🙂 😀 🤪 😎 🤗 🙂 🌞 🌼 🍀 🐶  
```

--- 

## 🌼 包与分层职责及方法名的建议

`dao`包、`model.entity`包为 `gf` 工具自动生成，最好不要在工具自动生成的包下面写代码。

控制层、服务层、逻辑层的代码各自的职责都划分好，具体执行的逻辑代码最好都写到 `logic` 包下。这样如果后面有对数据库新的增加或删除，可以直接删除工具自动生成的包，重新生成即可。

**方法名建议**：

- IndexXXX 服务端渲染的页面
- AddXXX 添加数据
- GetXXX 查询单个数据
- ListXXX 查询集合数据
- DelXXX 删除单个数据
- DelXXXs 删除集合数据¬
- UpdateXXX 更新单个数据
- UpdateXXXs 更新多个数据
- CheckXXX 相关规则的检查
- CountXXX 统计某些数据
- TakeXXX 查询多个组合的数据

**方法顺序排列**：(从易到难) 渲染页面，添加数据，查询数据，删除数据，修改数据，其他。

例如：

```text
func (c cRole) Index(r *ghttp.Request) {}             // 角色页面
func (c cRole) IndexAdd(r *ghttp.Request) {}          // 角色添加页面
func (c cRole) IndexEdit(r *ghttp.Request) {}         // 角色修改页面
func (c cRole) Add(r *ghttp.Request) {}               // 添加角色操作
func (c cRole) ListRoleNoMenus(r *ghttp.Request) {}   // 查询角色不包含的菜单集合
func (c cRole) ListRoleNoApis(r *ghttp.Request) {}    // 查询角色不包含的api集合
func (c cRole) Del(r *ghttp.Request) {}               // 删除角色操作
func (c cRole) DelRoles(r *ghttp.Request) {}          // 删除多个角色操作
func (c cRole) Update(r *ghttp.Request) {}            // 修改角色操作
```

这样分的好处在于，扩展新的功能时可以直接定位到各自的位置，让自己的思路清晰，代码也更可观。

---

## 👌 [自定义模版函数使用](./documents/template.md)

---

## 🌈 [css使用]()

### `.tag`

`.tag-info` `.tag-success` `.tag-priamry` `.tag-wanrnig` `.tag-brown` `.tag-purple` `.tag-danger`

![](./documents/tag.png) <br> ![](./documents/tag2.png)

### 💦 [更多class](https://github.com/1211ciel/ciel-admin/blob/master/documents/class.md)

![](./documents/class1.png) <br>
![](./documents/class2.png)

--- 

## 🥵 反馈

使用过程中如果遇见什么问题，可以发issue或发邮件1211sciel@gmail.com 我看到后会及时回复的。

