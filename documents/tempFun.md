# 自定义模版函数文档

## index页面查询

### input

使用场景：页面查询输入

- 参数1 input框的 name属性
- 参数2 描述
- 参数3 从当前Query map中获取对应值

```text
{{input "pid" "pid" .Query}}
输出
<label class="input">pid<input type="text" name="pid" value="" onkeydown="if(event.keyCode===13)this.form.submit()"> </label>
```

### options

使用场景：页面查询选择框

- 参数1 select框 name属性
- 参数2 描述
- 参数3 选项 eg 1:正常:tag-info,2:禁用:tag-danger 格式为 "值1:描述1:类名1,值2:描述2:类名2,值n:描述n:类名n"
- 参数4 从当前Query map中获取对应值

```text
{{options "status" "状态" .Config.options.status .Query}}
输出
<label class="input">状态 <select name="status" onchange="this.form.submit()"><option value="">请选择</option><option value="1" class="tag-info">正常</option><option value="2" class="tag-danger">禁用</option></select></label>
```

## 页面 table

### th

使用场景： 页面 table 设置 th标题

用小写逗号分割

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

### td

使用场景：页面 table 设置 td

```html

{{td "ID" .id}}
输出
<td data-label="ID">136</td>
```

### tdImg

使用场景：页面 table 设置 td

- 参数2 如果不以 `http` 开头 则会自动拼接本地图片前缀

> 本地图片前缀在 config.yml 中进行配置

```html

{{tdImg "图标" "1/2022/03/FdI4Yw.gif"}}
输出
<td data-label="图标"><a href="http://localhost:2033/upload/1/2022/03/FdI4Yw.gif" target="_blank"><img class="s-icon" src="http://localhost:2033/upload/1/2022/03/FdI4Yw.gif" alt="not fond"></a></td>
```

### tdChoose

使用场景：页面 table 设置 td 当该字段为选项时

```html
{{tdChoose "状态" "1:正常:tag-info,2:禁用:tag-danger" 1}}
输出
<td data-label="状态"><span class="tag-info">正常</span></td>
```

#### searchPageSize

使用场景： 每个页面查询都需要设置默认自动生成

```html
{{searchPageSize .Query}}
输出
<input id="page" name="page" value="1" hidden="">
<input name="size" value="10" hidden="">
```

等等, 更多使用请自行查看文件 `ciel-admin/internal/service/view/view.go`

可根据自己的需求进行封装使用。