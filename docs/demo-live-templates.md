## Live Templates

### ciel-api

缩写 `ciel-api`

模版文本

```text
//--- $Menu$$END$ ---------------------------------------------------------

type Add$Menu$Req struct {
	g.Meta `tags:"$group$" dc:"添加"`
	*do.$Menu$
}
type Get$Menu$Req struct {
	g.Meta `tags:"$group$" dc:"查询一条数据"`
	Id     uint64 `v:"required"`
}
type Get$Menu$Res struct {
	Data *entity.$Menu$ `json:"data"`
}

type List$Menu$Req struct {
	g.Meta `tags:"$group$" dc:"查询列表数据"`
	api.PageReq
}
type List$Menu$Res struct {
	List []*entity.$Menu$ `json:"list"`
	*api.PageRes
}

type Del$Menu$Req struct {
	g.Meta `tags:"$group$" dc:"删除"`
	Id     uint64 `v:"required"`
}

type Update$Menu$Req struct {
	g.Meta `tags:"$group$" dc:"修改菜单"`
	*do.$Menu$
}

```

### ciel-logic

缩写 `ciel-logic`

模版文本

```text
// --- $Menu$$END$ -----------------------------------------------------------------

func (l l$Biz$) Add$Menu$(ctx context.Context, in *do.$Menu$) error {
	if _, err := dao.$Menu$.Ctx(ctx).Insert(in); err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	return nil
}
func (l l$Biz$) Get$Menu$ById(ctx context.Context, id uint64) (*entity.$Menu$, error) {
	var d entity.$Menu$
	one, err := dao.$Menu$.Ctx(ctx).WherePri(id).One()
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
func (l l$Biz$) List$Menu$(ctx context.Context, in *$v2$.List$Menu$Req) ([]*entity.$Menu$, int, error) {
	var d = make([]*entity.$Menu$, 0)
	db := dao.$Menu$.Ctx(ctx)
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
func (l l$Biz$) Del$Menu$(ctx context.Context, id uint64) error {
	_, err := l.Get$Menu$ById(ctx, id)
	if err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	if _, err = dao.$Menu$.Ctx(ctx).WherePri(id).Delete(); err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	return nil
}
func (l l$Biz$) Update$Menu$(ctx context.Context, in *$v2$.Update$Menu$Req) error {
	if _, err := dao.$Menu$.Ctx(ctx).OmitEmpty().WherePri(in.Id).Update(in); err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	return nil
}

```

### ciel-service

缩写 `ciel-service`

模版文本

```text
// --- $Menu$$END$ -----------------------------------------------------------------

func (s s$Biz$) Add$Menu$(ctx context.Context, in *do.$Menu$) error {
	return logic.$Biz$.Add$Menu$(ctx, in)
}
func (s s$Biz$) Get$Menu$ById(ctx context.Context, id uint64) (*entity.$Menu$, error) {
	return logic.$Biz$.Get$Menu$ById(ctx, id)
}
func (s s$Biz$) List$Menu$(ctx context.Context, req *$v1$.List$Menu$Req) ([]*entity.$Menu$, *api.PageRes, error) {
	d, total, err := logic.$Biz$.List$Menu$(ctx, req)
	if err != nil {
		return nil, nil, err
	}
	return d, logic.Sys.MakePageInfo(req.Page, req.Size, int64(total)), nil
}
func (s s$Biz$) Del$Menu$(ctx context.Context, id uint64) error {
	return logic.$Biz$.Del$Menu$(ctx, id)
}
func (s s$Biz$) Update$Menu$(ctx context.Context, data *$v1$.Update$Menu$Req) error {
	return logic.$Biz$.Update$Menu$(ctx, data)
}

```

### ciel-controller

缩写 `ciel-controller`

模版文本

```text
// --- $menu$$END$-----------------------------------------------------------------

func (c c$Biz$) Get$menu$ById(ctx context.Context, req *$v1$.Get$menu$Req) (res *$v1$.Get$menu$Res, err error) {
	data, err := service.$Biz$.Get$menu$ById(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &$v1$.Get$menu$Res{Data: data}, nil
}
func (c c$Biz$) List$menu$(ctx context.Context, req *$v1$.List$menu$Req) (res *$v1$.List$menu$Res, err error) {
	d, pageRes, err := service.$Biz$.List$menu$(ctx, req)
	return &$v1$.List$menu$Res{List: d, PageRes: pageRes}, nil
}
func (c c$Biz$) Add$menu$(ctx context.Context, req *$v1$.Add$menu$Req) (res *api.DefaultRes, err error) {
	if err = service.$Biz$.Add$menu$(ctx, req.$menu$); err != nil {
		return nil, err
	}
	return
}
func (c c$Biz$) Del$menu$(ctx context.Context, req *$v1$.Del$menu$Req) (res *api.DefaultRes, err error) {
	if err = service.$Biz$.Del$menu$(ctx, req.Id); err != nil {
		return nil, err
	}
	return
}
func (c c$Biz$) Update$menu$(ctx context.Context, req *$v1$.Update$menu$Req) (res *api.DefaultRes, err error) {
	if err = service.$Biz$.Update$menu$(ctx, req); err != nil {
		return nil, err
	}
	return
}

```

### ciel-cmd

缩写 `ciel-cmd`

模版文本

```text
// --- $Menu$$END$ -----------------------------------------------------------------
g.Group("/$menu$", func(g *ghttp.RouterGroup) {
    g.GET("/", controller.$Biz$.Get$Menu$ById)
    g.GET("/list", controller.$Biz$.List$Menu$)
    g.Middleware(service.Sys.MiddlewareActionLock)
    g.POST("/", controller.$Biz$.Add$Menu$)
    g.DELETE("/", controller.$Biz$.Del$Menu$)
    g.PUT("/", controller.$Biz$.Update$Menu$)
})

```

这里涉及到了 `$Menu$` 和 `$menu$`  如果只想输入一个`$Menu$`就完成`$menu$`。那么在 Live Templates 的 `Edit Variables` 里面可以设置

| Name | Expression      | Skip if defined |
|------|-----------------|-----------------|
| Biz  |                 | false           |
| Menu |                 | false           |
| menu | camelCase(Menu) | true            |

### ciel-js

缩写 `ciel-js`

模板文本

```text
import useSWR from "swr";
import {useState} from "react";
import {toast} from "react-toastify";
import {listData} from "../../../libs/api-admin";
import {handleDel, objToParams,} from "../../../libs/utils";
import {FullScreenLoading} from "../../../compoents/common";
import {AddPage, Footer, Headers, Nav, PageButtons, PageInfo, SearchInput, Td, UpdatePage,} from "../../../compoents/sys-page";
const pageConf = {
    name: '$zhName$', path: '/$name$',
    fields: [
        {field: '$f0$', name: '$F0$', renderFn: (d) => d.$f0$, search: 1},
        {field: '$f1$', name: '$F1$', search: 1, required: 1},
        {field: '$f2$', name: '$F2$'},
        {field: '$f3$', name: '$F3$'},
        {field: '$f4$', name: '$F4$'},
    ]$END$
}

export default function $Name$() {
    const [query, setQuery] = useState() // 查询参数
    const [showType, setShowType] = useState(1) // 1 主页 2添加 3修改
    const [id, setId] = useState() // 修改数据时使用
    return (<>
        <Headers/>
        <div className={'wrapper'}>
            <div className="w">
                <div className={'wrapper-left'}>
                    <Nav/>
                    {showType === 1 && <$Name$Page query={query} setQuery={setQuery} setShowType={setShowType} setId={setId}/>}
                    {showType === 2 && <AddPage pageConf={pageConf} setShowType={setShowType}/>}
                    {showType === 3 && <UpdatePage pageConf={pageConf} setShowType={setShowType} id={id}/>}
                </div>
            </div>
        </div>
        <Footer/>
    </>)
}

const $Name$Page = ({query, setQuery, setShowType, setId}) => {
    const [tempQuery, setTempQuery] = useState({})
    const s = objToParams(query)
    const {data, isLoading, mutate, error} = useSWR(`/${pageConf.path}/list?${s !== undefined ? s : ''}`, listData)
    if (error||!data) return
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
                            <div className={'cell flex-center p-3'}><PageButtons query={query} setTempQuery={setTempQuery} setQuery={setQuery} totalPage={data&&data.total_page} total={data&&data.total}/></div>
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
                            <div className={'cell-tools p-3 flex-center'}><PageButtons totalPage={data&&data.total_page} total={data&&data.total} query={query} setTempQuery={setTempQuery} setQuery={setQuery}/></div>
                        </>
                    }
                </>
            }
        </div>
    </>
}
```