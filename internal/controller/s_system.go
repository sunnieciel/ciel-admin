package controller

import (
	"ciel-admin/internal/consts"
	"ciel-admin/internal/model/bo"
	"ciel-admin/internal/service/sys"
	"ciel-admin/utility/utils/res"
	"ciel-admin/utility/utils/xurl"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"sort"
)

type (
	home struct{}
	cSys struct{}
	gen  struct{}
	ws   struct{}
)

var (
	Home = &home{}
	Sys  = &cSys{}
	Gen  = &gen{}
	Ws   = &ws{}
)

// ---home-------------------------------------------------------------------

func (c *home) IndexPage(r *ghttp.Request) {
	res.Page(r, "/index.html", g.Map{"icon": "/resource/image/v2ex.png"})
}

// ---system-----------------------------------------------------------------

func (s cSys) Path(r *ghttp.Request) {
	path := r.GetQuery("path")
	res.Page(r, path.String())
}

func (s cSys) Level1(r *ghttp.Request) {
	level1, err := sys.MenusLevel1(r.Context())
	if err != nil {
		res.Err(err, r)
	}
	res.OkData(level1, r)
}

func (s cSys) GetDictByKey(r *ghttp.Request) {
	data, err := sys.DictGetByKey(r.Context(), r.Get("key").String())
	if err != nil {
		res.Err(err, r)
	}
	res.OkData(data, r)
}

// ---roleMenu-------------------------------------------------------------------

func (c *cRoleMenu) RoleNoMenus(r *ghttp.Request) {
	rid := r.GetQuery("rid")
	data, err := sys.RoleNoMenu(r.Context(), rid)
	if err != nil {
		res.Err(err, r)
	}
	res.OkData(data, r)
}
func (c *cRoleMenu) RoleNoApis(r *ghttp.Request) {
	rid := r.GetQuery("rid")
	data, err := sys.RoleNoApi(r.Context(), rid)
	if err != nil {
		res.Err(err, r)
	}
	res.OkData(data, r)
}

//  ---admin-------------------------------------------------------------------

func (c *cAdmin) LoginPage(r *ghttp.Request) {
	res.Page(r, "login.html")
}
func (c *cAdmin) Login(r *ghttp.Request) {
	var d struct {
		Uname string `form:"uname"`
		Pwd   string `form:"pwd"`
	}
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	if err := sys.Login(r.Context(), d.Uname, d.Pwd); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
func (c *cAdmin) Logout(r *ghttp.Request) {
	err := sys.Logout(r.Context())
	if err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
func (c *cAdmin) UpdatePwd(r *ghttp.Request) {
	var d struct {
		OldPwd string `v:"required"`
		NewPwd string `v:"required"`
	}
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	if err := sys.UpdateAdminPwd(r.Context(), d.OldPwd, d.NewPwd); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
func (c *cAdmin) UpdateUname(r *ghttp.Request) {
	var d struct {
		Uname string `v:"required"`
		Id    int64  `v:"required"`
	}
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	if err := sys.UpdateAdminUname(r.Context(), d.Id, d.Uname); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
func (c *cAdmin) UpdatePwdWithoutOldPwd(r *ghttp.Request) {
	var d struct {
		Pwd string `v:"required"`
		Id  string `v:"required"`
	}
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	if err := sys.UpdateAdminPwdWithoutOldPwd(r.Context(), d.Id, d.Pwd); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}

// ---Gen Code-------------------------------------------------------------------

func (c gen) Path(r *ghttp.Request) {
	icon, err := sys.Icon(r.Context(), r.URL.Path)
	if err != nil {
		res.Err(err, r)
	}
	res.Page(r, "/sys/gen.html", g.Map{"icon": icon})
}
func (c gen) Tables(r *ghttp.Request) {
	data, err := sys.Tables(r.Context())
	if err != nil {
		res.Err(err, r)
	}
	res.OkData(data, r)
}
func (c gen) Fields(r *ghttp.Request) {
	var d struct {
		Table string `v:"required#名表不能为空"`
	}
	err := r.Parse(&d)
	if err != nil {
		res.Err(err, r)
	}
	data, err := sys.Fields(r.Context(), d.Table)
	res.OkData(data, r)
}

func (c gen) GenFile(r *ghttp.Request) {
	var d bo.GenConf
	// set genConf
	genConf := r.Get("genConf")
	if err := genConf.Struct(&d); err != nil {
		res.Err(err, r)
	}
	if err := d.SetUrlPrefix(); err != nil {
		res.Err(err, r)
	}
	// set fields
	d.Fields = make([]*bo.GenFiled, 0)
	for _, v := range r.Get("fields").MapStrVarDeep() {
		f := &bo.GenFiled{}
		if err := v.Struct(f); err != nil {
			res.Err(err, r)
		}
		if f.FieldType == "select" {
			f.Options = make([]*bo.FieldOption, 0)
			for _, v := range v.Map()["Options"].(map[string]interface{}) {
				f.Options = append(f.Options, &bo.FieldOption{
					Value: v.(map[string]interface{})["Value"].(string),
					Type:  v.(map[string]interface{})["Type"].(string),
					Label: v.(map[string]interface{})["Name"].(string),
				})
				if gstr.IsNumeric(fmt.Sprint(v.(map[string]interface{})["Value"])) {
					sort.Slice(f.Options, func(i, j int) bool { return gconv.Int(f.Options[i].Value) < gconv.Int(f.Options[j].Value) })
				}
			}
		}
		d.Fields = append(d.Fields, f)
	}
	sort.Slice(d.Fields, func(i, j int) bool {
		return d.Fields[i].Index < d.Fields[j].Index
	})
	if len(d.Fields) == 0 {
		res.Err(errors.New("字段不能为空"), r)
	}
	g.Dump(d)
	if err := sys.GenFile(r.Context(), &d); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}

// --- Ws ------------------------------------------------------------------------

func (w ws) GetUserWs(r *ghttp.Request) {
	sys.GetUserWs(r)
}
func (w ws) GetAdminWs(r *ghttp.Request) {
	sys.GetAdminWs(r)
}
func (w ws) NoticeUser(r *ghttp.Request) {
	var d struct {
		Uid     int `v:"required"`
		OrderId int `v:"required"`
	}
	err := r.Parse(&d)
	if err != nil {
		res.Err(err, r)
	}
	err = sys.NoticeUser(gctx.New(), d.Uid, d)
	if err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
func (w ws) NoticeAdmin(r *ghttp.Request) {
	var d struct {
		Msg string `v:"required" json:"msg"`
	}
	err := r.Parse(&d)
	if err != nil {
		res.Err(err, r)
	}
	err = sys.NoticeAllAdmin(r.Context(), d)
	if err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}

func (c *cFile) Upload(r *ghttp.Request) {
	msg := fmt.Sprintf(consts.MsgPrimary, "上传成功")
	if err := sys.UploadFile(r.Context(), r); err != nil {
		msg = fmt.Sprintf(consts.MsgPrimary, err.Error())
	}
	r.Session.Set("msg", msg)
	r.Response.RedirectTo("/file/path/add?" + xurl.ToUrlParams(r.GetQueryMap()))
}
