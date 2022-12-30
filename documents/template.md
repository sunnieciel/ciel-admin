
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
