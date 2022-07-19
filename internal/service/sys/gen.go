package sys

import (
	"ciel-admin/internal/logic"
	"ciel-admin/internal/model/bo"
	"ciel-admin/utility/utils/xcmd"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/text/gstr"
)

var GenCommon = &gcmd.Command{
	Name:        "g",
	Brief:       "生成代码",
	Description: "生成代码",
	Func:        GenCode,
	Arguments: []gcmd.Argument{
		{Name: "queryField", Short: "q", Brief: "查询字段,eg: t1.*,t2.uname uname"},
		{Name: "orderBy", Short: "o", Brief: "排序，eg: t1.id desc"},
		{Name: "t1", Short: "t1", Brief: "主查询表,说明添加了关联字段,记得后面添加查询字段"},
		{Name: "t2", Short: "t2", Brief: "关联表,eg: u_user t2 on t2.id = t1.uid"},
		{Name: "t3", Short: "t3", Brief: "关联表,eg: u_login_log t3  on t3.uid = t2.id"},
		{Name: "t4", Short: "t4", Brief: "关联表,eg: u_wallet t4  on t4.uid = t2.id"},
		{Name: "t5", Short: "t5", Brief: "关联表,eg: u_user_info t5  on t5.uid = t2.id"},
		{Name: "t6", Short: "t6", Brief: "关联表,eg: u_operation t6  on t6.uid = t2.id"},
		{Name: "hideAddBtn", Short: "hadd", Brief: "隐藏添加按钮 eg 1"},
		{Name: "hideEditBtn", Short: "hedit", Brief: "隐藏修改按钮 eg 1"},
		{Name: "hideDelBtn", Short: "hdel", Brief: "隐藏删除按钮 eg 1"},
		{Name: "hideStatus", Short: "hs", Brief: "隐藏状态 eg 1"},
		{Name: "pageDesc", Short: "desc", Brief: "页面描述"},
		{Name: "pageLogo", Short: "logo", Brief: "页面图标,为空将随机生成"},
		{Name: "pageName", Short: "name", Brief: "页面菜单名称"},
		{Name: "genType", Short: "g", Brief: "生成类型1 crud 2静态页面"},
		{Name: "htmlGroup", Short: "group", Brief: "html页面分组分组文件夹 '项目/resource/template/你输入的分组文件' eg sys "},
	},
}

// Fields 根据表名查询字段信息
func Fields(ctx context.Context, tableName string) ([]*gdb.TableField, error) {
	return logic.Fields(ctx, tableName)
}

// Tables 查询当前数据库连接的所有表k
func Tables(ctx context.Context) ([]string, error) {
	return g.DB().Tables(ctx)
}

// GenCode 生成code
func GenCode(ctx context.Context, p *gcmd.Parser) error {
	d := bo.GenConf{}
	logic.GenCodeGreet(ctx)
	genType := logic.GenCodeSetConf(ctx, &d, p)
start:
	switch genType {
	case 1: // CRUD
		err := logic.CRUDBefore(ctx, &d)
		logic.CRUDParseFields(ctx, &d)
		if err = GenFile(ctx, d); err != nil {
			panic(err)
		}
	case 2: // 生成静态页面
		d.StructName = xcmd.MustScan("页面文件名称(不用加html):").String()
		if err := GenStaticHtmlFile(ctx, d); err != nil {
			panic(err)
		}
	default:
		glog.Warning(ctx, "该类型暂不支持")
		genType = xcmd.MustScan("生成类型(1 curd 2 静态页面):").Int()
		goto start
	}
	g.Log().Notice(ctx, "\n\n\n ██████╗ ██╗  ██╗    ███████╗██╗███╗   ██╗██╗███████╗██╗  ██╗███████╗██████╗ \n██╔═══██╗██║ ██╔╝    ██╔════╝██║████╗  ██║██║██╔════╝██║  ██║██╔════╝██╔══██╗\n██║   ██║█████╔╝     █████╗  ██║██╔██╗ ██║██║███████╗███████║█████╗  ██║  ██║\n██║   ██║██╔═██╗     ██╔══╝  ██║██║╚██╗██║██║╚════██║██╔══██║██╔══╝  ██║  ██║\n╚██████╔╝██║  ██╗    ██║     ██║██║ ╚████║██║███████║██║  ██║███████╗██████╔╝\n ╚═════╝ ╚═╝  ╚═╝    ╚═╝     ╚═╝╚═╝  ╚═══╝╚═╝╚══════╝╚═╝  ╚═╝╚══════╝╚═════╝ \n                                                                             \n")
	return nil
}
func GenFile(ctx context.Context, d bo.GenConf) (err error) {
	if d.StructName == "" {
		return fmt.Errorf("结构体名称不能为空")
	}
	if err = logic.GenMenu(ctx, d, func(name string) string {
		return fmt.Sprintf("/%s/path", gstr.CaseCamelLower(name))
	}, ""); err != nil {
		return err
	}
	if err = logic.GenController(d); err != nil {
		return err
	}
	if err = logic.GenRouter(d); err != nil {
		return err
	}
	if err = logic.GenHtml(d); err != nil {
		return err
	}
	if err = logic.GenApi(ctx, d.HtmlGroup, d.StructName, d.PageName); err != nil {
		return err
	}
	return
}
func GenStaticHtmlFile(ctx context.Context, d bo.GenConf) error {
	return logic.GenStaticHtmlFile(ctx, d)
}
