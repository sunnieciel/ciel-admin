# 🍉 freekey admin 自由钥匙 - 个人的后台管理系统

```
███████╗██████╗ ███████╗███████╗    ██╗  ██╗███████╗██╗   ██╗     █████╗ ██████╗ ███╗   ███╗██╗███╗   ██╗
██╔════╝██╔══██╗██╔════╝██╔════╝    ██║ ██╔╝██╔════╝╚██╗ ██╔╝    ██╔══██╗██╔══██╗████╗ ████║██║████╗  ██║
█████╗  ██████╔╝█████╗  █████╗      █████╔╝ █████╗   ╚████╔╝     ███████║██║  ██║██╔████╔██║██║██╔██╗ ██║
██╔══╝  ██╔══██╗██╔══╝  ██╔══╝      ██╔═██╗ ██╔══╝    ╚██╔╝      ██╔══██║██║  ██║██║╚██╔╝██║██║██║╚██╗██║
██║     ██║  ██║███████╗███████╗    ██║  ██╗███████╗   ██║       ██║  ██║██████╔╝██║ ╚═╝ ██║██║██║ ╚████║
╚═╝     ╚═╝  ╚═╝╚══════╝╚══════╝    ╚═╝  ╚═╝╚══════╝   ╚═╝       ╚═╝  ╚═╝╚═════╝ ╚═╝     ╚═╝╚═╝╚═╝  ╚═══╝
```

## v1.1 morri 版本正在更新中

如果你想先尝试，可以Fork 一份到自己的github(直接clone 文件可能有点大,因为包含以前老的版本).

## 🍓 追求

用越少的技术、越简单的方式来实现一个面向于web独立开发者的后台管理系统。

## 🍌 说明

freekey admin(自由钥匙) 一个V2EX风格go语言写的个人后台管理系统，提供了权限管理、本地文件上传等基本功能。采用`服务端渲染`，整个项目一键启动。

使用到的技术:

- 后端: goframe v2.0.6, mysql,redis
- 前端: html,css,javascript, jquery-3.6.0

## 🍒 内置功能

已完成

- 角色权限管理
- 字典管理
- 本地文件上传
- 用户管理
    - 后台修改用户名
    - 后台修改密码
    - 前台用户登录
    - 前台用户注册
- 用户登录日志
- 用户金币
    - 后台人工充值
    - 后台人工扣除
- 充值类型
- 账变记录
- 账变统计
- 账变报表

计划中 (todo)

- 商品买卖
- 主题购买
- 商品二手交易所

这是菜单页面 （图标选择的是哆啦A梦 ）![](./documents/home1.png) ![](./documents/home2.png)

## 🌴 项目结构

```text
├── hack
├── internal   
│   ├── cmd        入口指令                         ---👌 "如果没什么事的话，就让我们开始吧。"
│   ├── consts     常量                            --- 🙂 "请相信，我是永远不会改变的~"
│   ├── controller 控制层     提供对外接口           --- 😀 "您需要的所有功能我都可以为您实现,请告诉我吧~"
│   ├── dao        数据查询层  数据库查询操作在此层执行 --- 🤪 "数据库相关的信息我最了解。"
│   ├── logic      逻辑层     业务逻辑在此层进行处理,  --- 😎 "不管你的逻辑有多复杂，交给在下吧～"
│   ├── model      实体层     包含 bo do entity     --- 🤗 "我这里包含了你可能会用到的所有模型实体。"
│   └── service    服务层     为控制层提供服务支持     --- 🙂 "我只衷心为您服务~"
│       ├── admin                                 --- 🫡 "您想注册一个管理员？"
│       ├── dict                                  --- 😛 "查字典、查字典"
│       ├── file                                  --- 🧐 "让我们上传一个文件吧"
│       ├── role                                  --- 🥺 "这里有好几种角色"
│       ├── sys                                   --- 🙂 "这里是系统服务"
│       ├── view                                  --- 🍓 "我是为html自定义的模板函数,你想享受尽情地编写代码吗？"
│       └── ws                                    --- 🍉 "我提供了Websocket服务"
├── main.go                                       --- 🌞 "go!"
├── manifest   // 配置文件                          --- 🌼 "有什么事情可以先告诉我."
├── resource   // 静态或其他资源                     --- 🍀 "我是html 请到这里找来我吧"
└── utility    // 通用工具                          --- 🐶 "如果其他项目也想使用的话，请交给我."

      "我们一直在这里，等你!"
🌞 👌 🙂 😀 🤪 😎 🤗 🙂 🌞 🌼 🍀 🐶  
```

详情见：https://goframe.org/pages/viewpage.action?pageId=30740166

## 🌸 环境安装

1. golang环境 安装
2. mysql 安装
3. redis 安装
4. gf 安装  https://goframe.org/pages/viewpage.action?pageId=1115782

## ✨️ 快速开始

1. 下载项目 https://github.com/1211ciel/ciel-admin/releases/tag/v.1.0.9
2. 创建数据库并导入项目`resource`目录下的`ciel.sql`
3. 修改`config.yaml`文件配置:`mysql`,`redis` 将数据库信息替换成您自己的
4. 启动项目：到项目根目录下执行 `gf run main.go`
5. 浏览器打开 http://localhost:2033/login  默认用户名 `ciel` 密码 `1`

## 🍀 CRUD代码生成

[查看示例](https://github.com/1211ciel/ciel-admin/blob/master/documents/gen.md)

---

## 👌 自定义模版函数使用

### 使用实例

您可以像下面这样写简便的书写`html`代码.

#### 输入

```html

<table class="table-1">
    <tr>{{th "id,pid,名称,图标,背景图,路径,排序,类型,状态,创建时间,操作"}}</tr>
    {{range .list}}
    <tr>
        {{td "ID" .id}}
        {{td "PID" .pid}}
        {{td "名称" .name}}
        {{tdImg "图标" .icon}}
        {{tdImg "背景图" .bg_img}}
        {{td "路径" .path}}
        {{td "排序" .sort}}
        {{tdChoose "类型" $.Config.options.menuType .type}}
        {{tdChoose "状态" $.Config.options.status .status}}
        {{td "创建时间" .created_at}}
        <td align="center">
            {{a "tag-warning" (concat "/admin/menu/edit/" .id) "修改" $.Query}}
            {{aDel (concat "/admin/menu/del/" .id) $.Query}}
        </td>
    </tr>
    {{end}}
</table
```

#### 输出

```html

<table class="table-1">
    <tbody>
    <tr>
        <th>id</th>
        <th>pid</th>
        <th>名称</th>
        <th>图标</th>
        <th>背景图</th>
        <th>路径</th>
        <th>排序</th>
        <th>类型</th>
        <th>状态</th>
        <th>创建时间</th>
        <th>操作</th>
    </tr>
    <tr>
        <td data-label="ID">2</td>
        <td data-label="PID">1</td>
        <td data-label="名称">菜单</td>
        <td data-label="图标"><a href="http://localhost:2033/upload/1/2022/03/FdI4Yw.gif" target="_blank"><img class="s-icon" src="http://localhost:2033/upload/1/2022/03/FdI4Yw.gif" alt="not fond"></a></td>
        <td data-label="背景图"><span class="Tag-normal">暂无图片</span></td>
        <td data-label="路径">/admin/menu</td>
        <td data-label="排序">1.1</td>
        <td data-label="类型"><span class="tag-info">菜单</span></td>
        <td data-label="状态"><span class="tag-info">正常</span></td>
        <td data-label="创建时间">2022-02-16 19:14:13</td>
        <td align="center">
            <a class="tag-primary" href="/admin/menu/edit/2">修改</a>
            <a class="tag-purple" href="#" onclick="if(confirm('确认删除?')){location.href='/admin/menu/del/2?'}">删除</a>
        </td>
    </tr>
    ...
    </tbody>
</table>
```

[goframe 文档链接]( https://goframe.org/pages/viewpage.action?pageId=1114228)

### 封装示例

#### a

 ```html
 {{a "tag-primary mr-auto" "/admin/menu/add" "添加"}}
输出
<a class="tag-purple mr-auto" href="/admin/menu/add">添加</a> 
 ```

#### input

```html
{{input "pid" "pid" .Query}}
输出
<label class="input">pid<input type="text" name="pid" value="" onkeydown="if(event.keyCode===13)this.form.submit()"> </label>
```

#### th

```html
{{th "id,pid,名称,图标,背景图,路径,排序,类型,状态,创建时间,操作"}}
输出
<tr>
    <th>id</th>
    <th>pid</th>
    <th>名称</th>
    <th>图标</th>
    <th>背景图</th>
    <th>路径</th>
    <th>排序</th>
    <th>类型</th>
    <th>状态</th>
    <th>创建时间</th>
    <th>操作</th>
</tr>
```

[查看更多](https://github.com/1211ciel/ciel-admin/blob/master/documents/tempFun.md)
> [源码](https://github.com/1211ciel/ciel-admin/tree/master/internal/service/view)


---

## 🌈 css使用

### `.tag`

`.tag-info` `.tag-success` `.tag-priamry` `.tag-wanrnig` `.tag-brown` `.tag-purple` `.tag-danger`

![](./documents/tag.png) <br> ![](./documents/tag2.png)

### 💦 更多class

https://github.com/1211ciel/ciel-admin/blob/master/documents/class.md

![](./documents/class1.png) <br>
![](./documents/class2.png)

## 🥵 反馈

使用过程中如果遇见什么问题，可以发issue或发邮件1211sciel@gmail.com 我看到后会及时回复的。
