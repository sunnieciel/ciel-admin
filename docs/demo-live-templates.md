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

