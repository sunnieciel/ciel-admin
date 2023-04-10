# 充值订单开发示例

> Goland LiveTemplates 模版文本在 [这里](demo-live-templates.md)

## 1 后端代码创建

### 1.1 创建数据表

```text
create table w_top_up
(
    `id`          bigint unsigned auto_increment primary key,
    `uid`         bigint unsigned not null comment '用户id',
    `trans_id`    varchar(64) comment '交易id',
    `money`       decimal(12, 2) unsigned comment '充值金额',
    `change_type` int unsigned comment '账变类型 最好配置与此表 u_wallet_change_type 中的相对应。 1 支付宝充值, 2 微信充值, 3 Paypal充值',
    `ip`          varchar(64) comment '用户操作ip',
    `desc`        varchar(64) comment '备注',
    `status`      int unsigned default 1 comment '订单状态 1 等待, 2 成功, 3 失败',
    `aid`         bigint unsigned comment '管理员id',
    created_at    datetime not null default current_timestamp comment '创建时间',
    updated_at    datetime not null default current_timestamp comment '修改时间'
) engine MyISAM;
```

### 1.2 检查配置并生成 `dao` 文件

检查配置 `/freekey-admin/backend/hack/config.yaml`

在 `removePrefix` 后添加  `w_`

```yaml
gfcli:
  gen:
    dao:
      - link: "mysql:root:12345678@tcp(localhost:3306)/freekey_admin"
        tables: ""
        removePrefix: "s_,u_,c_,w_"
        descriptionTag: true
        noModelComment: true
        group: "sys"
        clear: true
        overwriteDao: true
```

> `tables`为空时`gf` 默认生成所有的数据表

进入 `/freekey-admin/backend` 执行 `gf gen dao`

```text
ciel@cieldeMacBook-Pro backend % gf gen dao
...
generated: internal/model/entity/wallet.go
generated: internal/model/entity/wallet_change_log.go
generated: internal/model/entity/wallet_change_type.go
generated: internal/model/entity/wallet_statistics_log.go
done!
ciel@cieldeMacBook-Pro backend % 
```

### 1.3 生成 api文件

进入 `/freekey-admin/backend/api/v1/biz.go`

使用 `goland` 的 `Live Templates` 在 `go` 文件中 输入 `ciel-api`

```text
//--- TopUp ---------------------------------------------------------

type AddTopUpReq struct {
	g.Meta `tags:"后台" dc:"添加"`
	*do.TopUp
}
type GetTopUpReq struct {
	g.Meta `tags:"后台" dc:"查询一条数据"`
	Id     uint64 `v:"required"`
}
type GetTopUpRes struct {
	Data *entity.TopUp `json:"data"`
}

type ListTopUpReq struct {
	g.Meta `tags:"后台" dc:"查询列表数据"`
	api.PageReq
}
type ListTopUpRes struct {
	List []*entity.TopUp `json:"list"`
	*api.PageRes
}

type DelTopUpReq struct {
	g.Meta `tags:"后台" dc:"删除"`
	Id     uint64 `v:"required"`
}

type UpdateTopUpReq struct {
	g.Meta `tags:"后台" dc:"修改菜单"`
	*do.TopUp
}
```

### 1.4 生成 logic 文件

进入 `/freekey-admin/backend/internal/logic/biz.go`

使用 `goland` 的 `Live Templates` 在 `go` 文件中 输入 `ciel-logic`

```text
// --- TopUp -----------------------------------------------------------------

func (l lBiz) AddTopUp(ctx context.Context, in *do.TopUp) error {
	if _, err := dao.TopUp.Ctx(ctx).Insert(in); err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	return nil
}
func (l lBiz) GetTopUpById(ctx context.Context, id uint64) (*entity.TopUp, error) {
	var d entity.TopUp
	one, err := dao.TopUp.Ctx(ctx).WherePri(id).One()
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	if one.IsEmpty() {
		return nil, consts.ErrDataNotFound
	}
	if err = one.Struct(&d); err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	return &d, nil
}
func (l lBiz) ListTopUp(ctx context.Context, in *v1.ListTopUpReq) ([]*entity.TopUp, int, error) {
	var d = make([]*entity.TopUp, 0)
	db := dao.TopUp.Ctx(ctx)
	count, err := db.Count()
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, 0, err
	}
	if err = db.Page(int(in.Page), int(in.Size)).Order("id desc").Scan(&d); err != nil {
		g.Log().Error(ctx, err)
		return nil, 0, err
	}
	return d, count, nil
}
func (l lBiz) DelTopUp(ctx context.Context, id uint64) error {
	_, err := l.GetTopUpById(ctx, id)
	if err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	if _, err = dao.TopUp.Ctx(ctx).WherePri(id).Delete(); err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	return nil
}
func (l lBiz) UpdateTopUp(ctx context.Context, in *v1.UpdateTopUpReq) error {
	if _, err := dao.TopUp.Ctx(ctx).OmitEmpty().WherePri(in.Id).Update(in); err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	return nil
}

```

### 1.5 生成 service 文件

进入 `/freekey-admin/backend/internal/service/biz.go`

使用 `goland` 的 `Live Templates` 在 `go` 文件中 输入 `ciel-service`

```text
// --- TopUp -----------------------------------------------------------------

func (s sBiz) AddTopUp(ctx context.Context, in *do.TopUp) error {
	return logic.Biz.AddTopUp(ctx, in)
}
func (s sBiz) GetTopUpById(ctx context.Context, id uint64) (*entity.TopUp, error) {
	return logic.Biz.GetTopUpById(ctx, id)
}
func (s sBiz) ListTopUp(ctx context.Context, req *v1.ListTopUpReq) ([]*entity.TopUp, *api.PageRes, error) {
	d, total, err := logic.Biz.ListTopUp(ctx, req)
	if err != nil {
		return nil, nil, err
	}
	return d, logic.Sys.MakePageInfo(req.Page, req.Size, int64(total)), nil
}
func (s sBiz) DelTopUp(ctx context.Context, id uint64) error {
	return logic.Biz.DelTopUp(ctx, id)
}
func (s sBiz) UpdateTopUp(ctx context.Context, data *v1.UpdateTopUpReq) error {
	return logic.Biz.UpdateTopUp(ctx, data)
}

```

### 1.6 生成 controller 文件

进入 `/freekey-admin/backend/internal/controller/biz.go`

使用 `goland` 的 `Live Templates` 在 `go` 文件中 输入 `ciel-controller`

```text
// --- TopUp-----------------------------------------------------------------

func (c cBiz) GetTopUpById(ctx context.Context, req *v1.GetTopUpReq) (res *v1.GetTopUpRes, err error) {
	data, err := service.Biz.GetTopUpById(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &v1.GetTopUpRes{Data: data}, nil
}
func (c cBiz) ListTopUp(ctx context.Context, req *v1.ListTopUpReq) (res *v1.ListTopUpRes, err error) {
	d, pageRes, err := service.Biz.ListTopUp(ctx, req)
	return &v1.ListTopUpRes{List: d, PageRes: pageRes}, nil
}
func (c cBiz) AddTopUp(ctx context.Context, req *v1.AddTopUpReq) (res *api.DefaultRes, err error) {
	if err = service.Biz.AddTopUp(ctx, req.TopUp); err != nil {
		return nil, err
	}
	return
}
func (c cBiz) DelTopUp(ctx context.Context, req *v1.DelTopUpReq) (res *api.DefaultRes, err error) {
	if err = service.Biz.DelTopUp(ctx, req.Id); err != nil {
		return nil, err
	}
	return
}
func (c cBiz) UpdateTopUp(ctx context.Context, req *v1.UpdateTopUpReq) (res *api.DefaultRes, err error) {
	if err = service.Biz.UpdateTopUp(ctx, req); err != nil {
		return nil, err
	}
	return
}

```

### 1.7 生成 cmd 路由

进入 `/freekey-admin/backend/internal/cmd/cmd.go`

使用 `goland` 的 `Live Templates` 在 `sysRouters` 方法下 输入 `ciel-cmd`

```text
	// --- TopUp -----------------------------------------------------------------
	g.Group("/topUp", func(g *ghttp.RouterGroup) {
	    g.GET("/", controller.Biz.GetTopUpById)
	    g.GET("/list", controller.Biz.ListTopUp)
	    g.Middleware(service.Sys.MiddlewareActionLock)
	    g.POST("/", controller.Biz.AddTopUp)
	    g.DELETE("/", controller.Biz.DelTopUp)
	    g.PUT("/", controller.Biz.UpdateTopUp)
	})

```

好,到这里后端的代码就生成完啦,下面我们来生成前台代码. 重启一下后台程序.

## 2. 前台代码生成

### 2.1 生成页面js

进入 `/freekey-admin/frontend/pages/backend/wallet`

创建 `topUp.js`

使用 `goland` 的 `Live Templates` 在 `sysRouters` 方法下 输入 `ciel-js`

```text
import useSWR from "swr";
import {useState} from "react";
import {listData} from "../../../libs/api-admin";
import {handleDel, objToParams,} from "../../../libs/utils";
import {FullScreenLoading} from "../../../compoents/common";
import {AddPage, Footer, Headers, Nav, PageButtons, PageInfo, SearchInput, Td, UpdatePage,} from "../../../compoents/sys-page";

const pageConf = {
    name: '充值订单', path: '/topUp',
    // 页面配置，这里可以根据数据库的字段进行方便地添加或修改
    fields: [
        {field: 'id', name: 'Id', renderFn: (d) => d.id, search: 1}, // 
        {field: 'uname', name: 'Uname', search: 1, required: 1},
        {field: 'transId', name: 'TransId'},
        {field: 'changeType', name: 'ChangeType'},
        {field: 'ip', name: 'Ip'},
    ]
}

export default function TopUp() {
    const [query, setQuery] = useState() // 查询参数
    const [showType, setShowType] = useState(1) // 1 主页 2添加 3修改
    const [id, setId] = useState() // 修改数据时使用
    return (<>
        <Headers/>
        <div className={'wrapper'}>
            <div className="w">
                <div className={'wrapper-left'}>
                    <Nav/>
                    {showType === 1 && <TopUpPage query={query} setQuery={setQuery} setShowType={setShowType} setId={setId}/>}
                    {showType === 2 && <AddPage pageConf={pageConf} setShowType={setShowType}/>}
                    {showType === 3 && <UpdatePage pageConf={pageConf} setShowType={setShowType} id={id}/>}
                </div>
            </div>
        </div>
        <Footer/>
    </>)
}

const TopUpPage = ({query, setQuery, setShowType, setId}) => {
    const [tempQuery, setTempQuery] = useState({})
    const s = objToParams(query)
    const {data, isLoading, mutate, error} = useSWR(`/${pageConf.path}/list?${s !== undefined ? s : ''}`, listData)
    if (error ) return
    return <>
        <PageInfo>
            <div className={'cell p-3 flex-center'}>
                <span className={'btn-info ml-12 mr-auto'} onClick={() => setShowType(2)}>添加</span>
                <SearchInput pageConf={pageConf} tempQuery={tempQuery} setTempQuery={setTempQuery} setQuery={setQuery}></SearchInput>
            </div>
        </PageInfo>
        <div className={'box-02 no-bottom-border'}>
            {isLoading
                ? <FullScreenLoading/>
                : <>
                    {data && data.list.length === 0
                        ? <div className={'cell color-desc-02 fs-13'}>暂无数据</div>
                        : <>
                            <div className={'cell flex-center p-3'}><PageButtons query={query} setTempQuery={setTempQuery} setQuery={setQuery} totalPage={data && data.total_page} total={data && data.total}/></div>
                            <table className={'table-1'}>
                                <tbody>
                                <tr>{pageConf.fields.filter(i => !i.hide).map((i, index) => <th key={index}>{i.name}</th>)}
                                    <th>操作</th>
                                </tr>
                                {data && data.list.map((i, index) => <tr key={index}><Td pageConf={pageConf} data={i}/>
                                    <td>
                                        <button className={'btn-warning mr-6'} onClick={() => setId(i.id) & setShowType(3)}>修改</button>
                                        <button className={'btn-danger'} onClick={() => handleDel(pageConf.path, i.id, mutate)}>删除</button>
                                    </td>
                                </tr>)}
                                </tbody>
                            </table>
                            <div className={'cell-tools p-3 flex-center'}><PageButtons totalPage={data && data.total_page} total={data && data.total} query={query} setTempQuery={setTempQuery} setQuery={setQuery}/></div>
                        </>
                    }
                </>
            }
        </div>
    </>
}
```

#### pageConf 说明

pageConf 的常量，它包含了一个页面的配置信息，具体如下：

- name：页面的名称。
- path：页面的路径。
- fields：页面的字段信息，它是一个数组，包含了多个对象，每个对象代表一个字段。每个字段对象包含以下属性：
    - field：字段名，比如“id”、“uname”等等。
    - name：字段的显示名称，比如“ID”、“用户名”等等。
    - renderFn：字段的渲染函数，这里使用箭头函数定义，接收一个参数 d 返回当前数据,用户可以自己获取信息自定义渲染。
    - search：字段是否支持搜索，值为 1 代表支持搜索。
    - required：字段是否为必填项，如果存在该属性并且值为 1，代表该字段是必填项。
    - disabled：字段是否为禁用状态，如果存在该属性并且值为 1，代表该字段为禁用状态。
    - hide：字段是否隐藏，如果存在该属性并且值为 1，代表该字段为隐藏状态。
    - type：字段的类型，如果存在该属性，代表该字段为特殊类型。这里包含了两种类型：select ,img 和 textarea,
    - options：当 type 为 select 时，该属性代表下拉框的选项。eg: "1:正常:tag-success,2:禁用:tag-danger"
    - editHide：当 type 为 select 时，该属性代表编辑时是否隐藏该字段。

### 2.2 创建充值订单菜单

登录后台进入 http://localhost:3000/backend/sys/menu

在钱包分组下创建 菜单名称为 `充值订单`， 排序为 `6.3`, 页面地址为 `/backend/wallet/topUp`

> nextjs 自动识别 `pages` 文件夹下面的所有导出为 js的页面。 所以这里直接填写相应的路径信息即可。

### 2.3 角色添加页面权限

进入 http://localhost:3000/backend/sys/role

对超级管理员的菜单权限进行修改，添加我们刚添加的 `充值订单页面`

刷新一下浏览器，就可以看到，钱包下面已经有充值订单页面了。

好的，前后台文件生成就到这里。下面我们进行创建充值订单前台api进行开发。

[充值api开发](demo-topUp-interface.md)