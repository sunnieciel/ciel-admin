package logic

import (
	"context"
	"fmt"
	"freekey-backend/api"
	"freekey-backend/api/v1"
	"freekey-backend/internal/consts"
	"freekey-backend/internal/dao"
	"freekey-backend/internal/model"
	"freekey-backend/internal/model/entity"
	"freekey-backend/utility/utils/res"
	"freekey-backend/utility/utils/xhtml"
	"freekey-backend/utility/utils/xjwt"
	"freekey-backend/utility/utils/xpwd"
	"freekey-backend/utility/utils/xstr"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	"github.com/yudeguang/ratelimit"
	"math"
	"path"
	"strings"
	"time"
)

var Sys = lSys{}

type lSys struct{}

func (l lSys) MakePageInfo(page, size, total interface{}) *api.PageRes {
	if size == 0 {
		size = 10
	}
	totalPage := math.Ceil(gconv.Float64(total) / gconv.Float64(size)) //这里计算总页数时，要向上取整
	if totalPage <= 0 {
		totalPage = 1
	}
	if gconv.Int64(page) == 0 {
		page = 1
	}
	return &api.PageRes{
		TotalPage: int64(totalPage),
		Total:     gconv.Int64(total),
		Page:      gconv.Int64(page),
		Size:      gconv.Int64(size),
	}
}

func (l lSys) SwitchEnvironment(current, dev, server string, t int) error {
	switch t {
	case 0: // env
		if err := gfile.Rename(current, server); err != nil {
			return err
		}
		if err := gfile.Rename(dev, current); err != nil {
			return err
		}
	case 1: // server
		if err := gfile.Rename(current, dev); err != nil {
			return err
		}
		if err := gfile.Rename(server, current); err != nil {
			return err
		}
	}
	return nil
}

// -----------------Ws----------------------------

func (l lSys) WsSendAdminMsg(ctx context.Context, d *model.AdminMsg) error {
	toAdmin, err := l.GetAdminByUname(ctx, d.ToUname)
	if err != nil {
		return err
	}
	d.ToUid = uint64(toAdmin.Id)
	d.FromUname = l.GetUnameFromCtx(ctx)
	return Ws.NoticeAdmin(ctx, d, uint64(toAdmin.Id))
}
func (l lSys) WsNoticeAdmins(ctx context.Context, d *model.AdminMsg) error {
	return Ws.NoticeAdmins(ctx, d)
}

func (l lSys) MiddlewareCORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}
func (l lSys) MiddlewareWhiteIp(r *ghttp.Request) {
	if consts.WhiteIps != "" {
		ip := r.GetClientIp()
		ips := consts.WhiteIps
		if !gstr.Contains(ips, ip) {
			if err := l.UpdateDictWhiteIps(r.Context()); err != nil {
				res.Err(err, r)
			}
			if !gstr.Contains(consts.WhiteIps, ip) {
				res.Err(fmt.Errorf("%s ip error", r.GetClientIp()), r)
			}
		}
	}
	r.Middleware.Next()
}
func (l lSys) CreateRateLimit(fn func(rule *ratelimit.Rule)) *ratelimit.Rule {
	r := ratelimit.NewRule()
	fn(r)
	return r
}

// ---------------Menu---------------------------

func (l lSys) AddMenu(ctx context.Context, menu *entity.Menu) error {
	if _, err := dao.Menu.Ctx(ctx).Insert(menu); err != nil {
		return err
	}
	return nil
}
func (l lSys) GetMenuByName(ctx context.Context, name string) (*entity.Menu, error) {
	var data entity.Menu
	one, err := dao.Menu.Ctx(ctx).One("name", name)
	if err != nil {
		return nil, nil
	}
	if one.IsEmpty() {
		return nil, consts.ErrDataNotFound
	}
	if err = one.Struct(&data); err != nil {
		return nil, err
	}
	return &data, nil
}
func (l lSys) GetMenuByPath(ctx context.Context, path string) (*entity.Menu, error) {
	var data entity.Menu
	one, err := dao.Menu.Ctx(ctx).Where("path", path).One()
	if err != nil {
		return nil, err
	}
	if one.IsEmpty() {
		return nil, consts.ErrDataNotFound
	}
	if err = one.Struct(&data); err != nil {
		return nil, err
	}
	return &data, nil
}
func (l lSys) GetMenuById(ctx context.Context, id uint64) (*entity.Menu, error) {
	var data entity.Menu
	one, err := dao.Menu.Ctx(ctx).WherePri(id).One()
	if err != nil {
		return nil, err
	}
	if one.IsEmpty() {
		return nil, consts.ErrDataNotFound
	}
	if err = one.Struct(&data); err != nil {
		return nil, err
	}
	return &data, nil
}
func (l lSys) ListMenu(ctx context.Context, req *v1.ListMenuReq) ([]*entity.Menu, int, error) {
	var data = make([]*entity.Menu, 0)
	db := dao.Menu.Ctx(ctx)
	if req.Pid != 0 {
		db = db.Where("pid", req.Pid)
	}
	if req.Name != "" {
		db = db.Where("name", req.Name)
	}
	count, err := db.Count()
	if err != nil {
		return nil, 0, err
	}
	if err = db.Page(int(req.Page), int(req.Size)).Order("sort").Scan(&data); err != nil {
		return nil, 0, err
	}
	return data, count, nil
}
func (l lSys) ListMenuByPid(ctx context.Context, id int) ([]*entity.Menu, error) {
	var data = make([]*entity.Menu, 0)
	err := dao.Menu.Ctx(ctx).Scan(&data, "pid", id)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func (l lSys) DelMenu(ctx context.Context, id uint64) error {
	_, err := dao.Menu.Ctx(ctx).Delete("id", id)
	return err
}
func (l lSys) SaveMenu(ctx context.Context, i *entity.Menu) error {
	if _, err := dao.Menu.Ctx(ctx).Save(i); err != nil {
		return err
	}
	return nil
}

// -----------------API----------------------------

func (l lSys) AddApi(ctx context.Context, menu *entity.Api) error {
	if _, err := dao.Api.Ctx(ctx).Insert(menu); err != nil {
		return err
	}
	return nil
}
func (l lSys) GetApiById(ctx context.Context, id uint64) (*entity.Api, error) {
	var data entity.Api
	one, err := dao.Api.Ctx(ctx).WherePri(id).One()
	if err != nil {
		return nil, err
	}
	if one.IsEmpty() {
		return nil, consts.ErrDataNotFound
	}
	if err = one.Struct(&data); err != nil {
		return nil, err
	}
	return &data, nil
}
func (l lSys) GetApiByMethodAndUrl(ctx context.Context, method string, url string) (*entity.Api, error) {
	var data entity.Api
	one, err := dao.Api.Ctx(ctx).One("method = ? and url = ?", method, url)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	if one.IsEmpty() {
		return nil, consts.ErrDataNotFound
	}
	if err = one.Struct(&data); err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	return &data, nil
}
func (l lSys) ListApi(ctx context.Context, req *v1.ListApiReq) ([]*entity.Api, int, error) {
	var data = make([]*entity.Api, 0)
	db := dao.Api.Ctx(ctx)
	if req.Desc != "" {
		db = db.WhereLike("desc", xstr.Like(req.Desc))
	}
	if req.Url != "" {
		db = db.WhereLike("url", xstr.Like(req.Url))
	}
	if req.Method != "" {
		db = db.Where("method", req.Method)
	}
	if req.Type != "" {
		db = db.Where("type", req.Type)
	}
	if req.Group != "" {
		db = db.Where("group", req.Group)
	}
	count, err := db.Count()
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, 0, err
	}
	if err = db.Page(int(req.Page), int(req.Size)).Order("group,id desc,method").Scan(&data); err != nil {
		g.Log().Error(ctx, err)
		return nil, 0, err
	}
	return data, count, nil
}
func (l lSys) DelApi(ctx context.Context, id uint64) error {
	if _, err := dao.Api.Ctx(ctx).Delete("id", id); err != nil {
		return err
	}
	return nil
}
func (l lSys) UpdateApi(ctx context.Context, data *v1.UpdateApiReq) error {
	if _, err := dao.Api.Ctx(ctx).OmitEmpty().Update(data, "id", data.Id); err != nil {
		return err
	}
	return nil
}

// -----------------Role----------------------------

func (l lSys) AddRole(ctx context.Context, role *entity.Role) error {
	if _, err := dao.Role.Ctx(ctx).Insert(role); err != nil {
		return err
	}
	return nil
}
func (l lSys) GetRoleOptions(ctx context.Context) (string, error) {
	all, err := dao.Role.Ctx(ctx).All()
	if err != nil {
		g.Log().Error(ctx, err)
		return "", err
	}
	var arr []string
	for i, record := range all {
		arr = append(arr, fmt.Sprintf("%d:%s:%s", record["id"].Uint64(), record["name"], xhtml.SwitchTagClass(i)))
	}
	return strings.Join(arr, ","), nil
}
func (l lSys) GetRoleById(ctx context.Context, id uint64) (*entity.Role, error) {
	var data entity.Role
	one, err := dao.Role.Ctx(ctx).WherePri(id).One()
	if err != nil {
		return nil, err
	}
	if one.IsEmpty() {
		return nil, consts.ErrDataNotFound
	}
	if err = one.Struct(&data); err != nil {
		return nil, err
	}
	return &data, nil
}
func (l lSys) ListRole(ctx context.Context, req *v1.ListRoleReq) ([]*entity.Role, int, error) {
	var data = make([]*entity.Role, 0)
	db := dao.Role.Ctx(ctx)
	count, err := db.Count()
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, 0, err
	}
	if err = db.Page(int(req.Page), int(req.Size)).Order("id desc").Scan(&data); err != nil {
		g.Log().Error(ctx, err)
		return nil, 0, err
	}
	return data, count, nil
}
func (l lSys) DelRole(ctx context.Context, id uint64) error {
	if _, err := dao.Role.Ctx(ctx).Delete("id", id); err != nil {
		return err
	}
	return nil
}
func (l lSys) DelRoleMenusByRid(ctx context.Context, rid interface{}) error {
	if _, err := dao.RoleMenu.Ctx(ctx).Delete("rid", rid); err != nil {
		return err
	}
	return nil
}
func (l lSys) UpdateRole(ctx context.Context, data *v1.UpdateRoleReq) error {
	if _, err := dao.Role.Ctx(ctx).OmitEmpty().Update(data, "id", data.Id); err != nil {
		return err
	}
	return nil
}
func (l lSys) CheckRoleApi(ctx context.Context, rid int, uri, method string) bool {
	if uri == "/" {
		return true
	}
	count, _ := g.DB("sys").Ctx(ctx).Model("s_role t1").
		LeftJoin("s_role_api t2 on t1.id = t2.rid").
		LeftJoin("s_api t3 on t2.aid = t3.id").
		Where("t3.url = ? and  t1.id = ?  and t3. method = ?", uri, rid, method).
		Count()
	if count == 1 {
		return false
	}
	return true
}
func (l lSys) ListAllRole(ctx context.Context) ([]*entity.Role, error) {
	var d = make([]*entity.Role, 0)
	err := dao.Role.Ctx(ctx).Scan(&d)
	return d, err
}
func (l lSys) MixRoleOptions(all []*entity.Role) (string, error) {
	var arr []string
	for i, record := range all {
		arr = append(arr, fmt.Sprintf("%d:%s:%s", record.Id, record.Name, xhtml.SwitchTagClass(i)))
	}
	return strings.Join(arr, ","), nil
}

// -----------------RoleMenu----------------------------

func (l lSys) AddMenus(ctx context.Context, rid int, ids []int) error {
	return g.DB("sys").Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		for _, item := range ids {
			if _, err := tx.Ctx(ctx).Model(dao.RoleMenu.Table()).Replace(g.Map{
				"rid": rid,
				"mid": item,
			}); err != nil {
				return err
			}
		}
		return nil
	})
}
func (l lSys) ListRoleMenu(ctx context.Context, req *v1.ListRoleMenuReq) ([]*model.RoleMenu, int, error) {
	var data = make([]*model.RoleMenu, 0)
	db := g.DB("sys").Model(dao.RoleMenu.Table() + " t1").
		LeftJoin("s_menu t2 on t1.mid = t2.id").
		LeftJoin("s_role t3 on t1.rid =t3.id").Ctx(ctx)
	if req.Rid != 0 {
		db = db.Where("t1.rid", req.Rid)
	}
	count, err := db.Count()
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, 0, err
	}
	if err = db.Page(int(req.Page), int(req.Size)).Fields("t1.id,t2.name menu_name,t2.type,t3.name role_name").Order("sort").Scan(&data); err != nil {
		g.Log().Error(ctx, err)
		return nil, 0, err
	}
	return data, count, nil
}
func (l lSys) ListRoleNoMenus(ctx context.Context, rid interface{}) ([]*v1.ListRoleNoMenusRes, error) {
	var data []*v1.ListRoleNoMenusRes
	array, err := dao.RoleMenu.Ctx(ctx).Array("mid", "rid", rid)
	if err != nil {
		return nil, err
	}
	db := dao.Menu.Ctx(ctx)
	if len(array) != 0 {
		db = db.WhereNotIn("id", array)
	}
	err = db.Order("sort").Fields("").Scan(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
func (l lSys) DelRoleMenu(ctx context.Context, id uint64) error {
	if _, err := dao.RoleMenu.Ctx(ctx).Delete("id", id); err != nil {
		return err
	}
	return nil
}
func (l lSys) ClearRoleMenu(ctx context.Context, rid uint64) error {
	if _, err := dao.RoleMenu.Ctx(ctx).Delete("rid", rid); err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	return nil
}

// -----------------RoleRole----------------------------

func (l lSys) AddRoleApis(ctx context.Context, rid int, ids []int) error {
	for _, item := range ids {
		_, err := dao.RoleApi.Ctx(ctx).Replace(g.Map{"rid": rid, "aid": item})
		if err != nil {
			return err
		}
	}
	return nil
}
func (l lSys) ListRoleApi(ctx context.Context, req *v1.ListRoleApiReq) ([]*model.RoleApi, int, error) {
	var data = make([]*model.RoleApi, 0)
	db := g.DB("sys").Model(dao.RoleApi.Table() + " t1").
		LeftJoin("s_role t2 on t1.rid = t2.id").
		LeftJoin("s_api t3 on t1.aid = t3.id").Ctx(ctx)
	if req.Rid != 0 {
		db = db.Where("t1.rid", req.Rid)
	}
	count, err := db.Count()
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, 0, err
	}
	if err = db.Page(int(req.Page), int(req.Size)).Order("t3.group desc").
		Fields("t1.id,t2.name roleName,t3.url path,t3.group ,t3.method ,t3.desc ").Scan(&data); err != nil {
		g.Log().Error(ctx, err)
		return nil, 0, err
	}
	return data, count, nil
}
func (l lSys) ListRoleNoApis(ctx context.Context, rid interface{}) ([]*v1.ListRoleNoApisRes, error) {
	var data = make([]*v1.ListRoleNoApisRes, 0)
	array, err := dao.RoleApi.Ctx(ctx).Array("aid", "rid", rid)
	if err != nil {
		return nil, err
	}
	db := dao.Api.Ctx(ctx)
	if len(array) != 0 {
		db = db.WhereNotIn("id", array)
	}
	if err = db.Order("group").Scan(&data); err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	return data, nil
}
func (l lSys) DelRoleApi(ctx context.Context, id uint64) error {
	if _, err := dao.RoleApi.Ctx(ctx).Delete("id", id); err != nil {
		return err
	}
	return nil
}
func (l lSys) ClearRoleApi(ctx context.Context, rid uint64) error {
	if _, err := dao.RoleApi.Ctx(ctx).Delete("rid", rid); err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	return nil
}

// -----------------Admin----------------------------

func (l lSys) AddAdmin(ctx context.Context, in *entity.Admin) error {
	if _, err := dao.Admin.Ctx(ctx).Insert(in); err != nil {
		return err
	}
	return nil
}
func (l lSys) GetAdminByUname(ctx context.Context, uname string) (*entity.Admin, error) {
	var data entity.Admin
	one, err := dao.Admin.Ctx(ctx).One("uname", uname)
	if err != nil {
		return nil, err
	}
	if one.IsEmpty() {
		return nil, consts.ErrDataNotFound
	}
	err = one.Struct(&data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
func (l lSys) GetUnameFromCtx(ctx context.Context) string {
	return ghttp.RequestFromCtx(ctx).Get(consts.TokenAdminUname).String()
}
func (l lSys) GetAdminIdFromCtx(ctx context.Context) uint64 {
	return ghttp.RequestFromCtx(ctx).Get(consts.TokenAdminIdKey).Uint64()
}
func (l lSys) GetAdminJwtInfo(r *ghttp.Request) (*xjwt.MyClaims, error) {
	userInfo, err := xjwt.UserInfo(r)
	if err != nil {
		return nil, err
	}
	return userInfo, nil
}

func (l lSys) GetAdminById(ctx context.Context, id uint64) (*entity.Admin, error) {
	var data entity.Admin
	one, err := dao.Admin.Ctx(ctx).WherePri(id).One()
	if err != nil {
		return nil, err
	}
	if one.IsEmpty() {
		return nil, consts.ErrDataNotFound
	}
	if err = one.Struct(&data); err != nil {
		return nil, err
	}
	return &data, nil
}
func (l lSys) ListAdmin(ctx context.Context, req *v1.ListAdminReq) ([]*entity.Admin, int, error) {
	var data = make([]*entity.Admin, 0)
	db := g.DB("sys").Model(dao.Admin.Table() + " t1").Ctx(ctx)
	if req.Id != 0 {
		db = db.Where("t1.id", req.Id)
	}
	if req.Uname != "" {
		db = db.WhereLike("t1.uname", xstr.Like(req.Uname))
	}
	if req.Rid != 0 {
		db = db.Where("t1.rid", req.Rid)
	}
	if req.Status != 0 {
		db = db.Where("t1.status", req.Status)
	}
	count, err := db.Count()
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, 0, err
	}
	if err = db.Page(int(req.Page), int(req.Size)).Order("id desc").Fields("t1.*").Scan(&data); err != nil {
		g.Log().Error(ctx, err)
		return nil, 0, err
	}
	return data, count, nil
}
func (l lSys) ListMenus(ctx context.Context, rid int, pid int) ([]*model.Menu, error) {
	var d = make([]*model.Menu, 0)
	menus, err := l.doMenus(ctx, rid, pid)
	if err != nil {
		return nil, err
	}
	d = append(d, menus...)
	return d, err
}
func (l lSys) doMenus(ctx context.Context, rid, pid int) ([]*model.Menu, error) {
	var data = make([]*model.Menu, 0)
	err := g.DB("sys").Ctx(ctx).Model(dao.RoleMenu.Table()+" t1").
		LeftJoin(dao.Menu.Table()+" t2 on t1.mid = t2.id").
		Fields("t2.*").
		Where("t1.rid = ? and t2.pid = ?", rid, pid).
		Order("t2.sort").
		Scan(&data)
	if err != nil {
		return nil, err
	}
	for _, item := range data {
		if item.Type == 2 {
			children, err := l.doMenus(ctx, rid, item.Id)
			if err != nil {
				return nil, err
			}
			item.Children = children
		} else {
			item.Children = make([]*model.Menu, 0)
		}
	}
	return data, nil
}
func (l lSys) DelAdmin(ctx context.Context, id uint64) error {
	if _, err := dao.Admin.Ctx(ctx).Delete("id", id); err != nil {
		return err
	}
	return nil
}

func (l lSys) UpdateAdmin(ctx context.Context, data *v1.UpdateAdminReq) error {
	if _, err := dao.Admin.Ctx(ctx).OmitEmpty().Update(data, "id", data.Id); err != nil {
		return err
	}
	return nil
}
func (l lSys) UpdateAdminUname(ctx context.Context, id uint64, uname string) error {
	if _, err := dao.Admin.Ctx(ctx).Update(g.Map{"uname": uname}, "id", id); err != nil {
		return err
	}
	return nil
}
func (l lSys) UpdateAdminPass(ctx context.Context, id interface{}, pwd string) error {
	_, err := dao.Admin.Ctx(ctx).Update(g.Map{"pwd": xpwd.GenPwd(pwd)}, "id", id)
	if err != nil {
		return err
	}
	return nil
}

func (l lSys) CheckAdminPass(admin *entity.Admin, pass string) error {
	if !xpwd.ComparePassword(admin.Pwd, pass) {
		return consts.ErrLogin
	}
	return nil
}
func (l lSys) CheckAdminStatus(admin *entity.Admin) error {
	if admin.Status == 2 {
		return consts.ErrAuthNotEnough
	}
	return nil
}
func (l lSys) AddAdminLoginLog(ctx context.Context, adminId int) error {
	d := entity.AdminLoginLog{}
	d.Uid = adminId
	d.Ip = ghttp.RequestFromCtx(ctx).GetClientIp()
	d.Status = 1
	if _, err := dao.AdminLoginLog.Ctx(ctx).Insert(&d); err != nil {
		return err
	}
	return nil
}
func (l lSys) CountAdminByUname(ctx context.Context, uname string) (int, error) {
	count, err := dao.Admin.Ctx(ctx).Count("uname", uname)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// -----------------Dict----------------------------

func (l lSys) AddDict(ctx context.Context, in *entity.Dict) error {
	d, _ := l.GetDictByKey(ctx, in.K)
	if d != nil {
		return consts.ErrDataAlreadyExist
	}
	if _, err := dao.Dict.Ctx(ctx).Insert(in); err != nil {
		return err
	}
	return nil
}
func (l lSys) GetDictById(ctx context.Context, id uint64) (*entity.Dict, error) {
	var data entity.Dict
	one, err := dao.Dict.Ctx(ctx).WherePri(id).One()
	if err != nil {
		return nil, err
	}
	if one.IsEmpty() {
		return nil, consts.ErrDataNotFound
	}
	if err = one.Struct(&data); err != nil {
		return nil, err
	}
	return &data, nil
}
func (l lSys) GetDictByKey(ctx context.Context, s string) (*entity.Dict, error) {
	var data entity.Dict
	one, err := dao.Dict.Ctx(ctx).One("k", s)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	if one.IsEmpty() {
		return nil, consts.ErrDataNotFound
	}
	if err = one.Struct(&data); err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	return &data, nil
}
func (l lSys) ListDict(ctx context.Context, req *v1.ListDictReq) ([]*entity.Dict, int, error) {
	var data = make([]*entity.Dict, 0)
	db := dao.Dict.Ctx(ctx)
	count, err := db.Count()
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, 0, err
	}
	if err = db.Page(int(req.Page), int(req.Size)).Order("id desc").Scan(&data); err != nil {
		g.Log().Error(ctx, err)
		return nil, 0, err
	}
	return data, count, nil
}
func (l lSys) UpdateDictWhiteIps(ctx context.Context) error {
	d, err := l.GetDictByKey(ctx, "white_ips")
	if err != nil {
		return err
	}
	consts.WhiteIps = d.V
	return nil
}
func (l lSys) DelDict(ctx context.Context, id uint64) error {
	if _, err := dao.Dict.Ctx(ctx).Delete("id", id); err != nil {
		return err
	}
	return nil
}
func (l lSys) UpdateDict(ctx context.Context, data *v1.UpdateDictReq) error {
	if _, err := dao.Dict.Ctx(ctx).Update(data, "id", data.Id); err != nil {
		return err
	}
	return nil
}

// -----------------OperationLog----------------------------

func (l lSys) AddOperationLog(ctx context.Context, log *entity.OperationLog) error {
	_, err := dao.OperationLog.Ctx(ctx).Insert(log)
	return err
}
func (l lSys) ListOperationLog(ctx context.Context, req *v1.ListOperationLogReq) ([]*model.OperationLog, int, error) {
	var data = make([]*model.OperationLog, 0)
	db := g.DB("sys").Model(dao.OperationLog.Table() + " t1").
		LeftJoin("s_admin t2 on t1.uid = t2.id").
		LeftJoin(dao.Api.Table() + " t3 on t1.method = t3.method and t1.uri = t3.url").Ctx(ctx)
	if req.Desc != "" {
		db = db.WhereLike("t3.desc", xstr.Like(req.Desc))
	}
	if req.Uname != "" {
		db = db.WhereLike("t2.uname", xstr.Like(req.Uname))
	}
	if req.Ip != "" {
		db = db.WhereLike("t1.ip", xstr.Like(req.Ip))
	}
	if req.Content != "" {
		db = db.WhereLike("t1.content", xstr.Like(req.Content))
	}
	if req.Uri != "" {
		db = db.WhereLike("t1.uri", xstr.Like(req.Uri))
	}
	if req.Response != "" {
		db = db.WhereLike("t1.response", xstr.Like(req.Response))
	}
	if req.Method != "" {
		db = db.Where("t1.method", req.Method)
	}
	count, err := db.Count()
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, 0, err
	}
	if err = db.Page(int(req.Page), int(req.Size)).Order("t1.id desc").Fields("t1.*,t2.uname,t3.desc").Scan(&data); err != nil {
		g.Log().Error(ctx, err)
		return nil, 0, err
	}
	return data, count, nil
}
func (l lSys) DelOperationLog(ctx context.Context, id uint64) error {
	if _, err := dao.OperationLog.Ctx(ctx).Delete("id", id); err != nil {
		return err
	}
	return nil
}
func (l lSys) ClearOperationLog(ctx context.Context) error {
	if _, err := dao.OperationLog.Ctx(ctx).Delete("id !=0"); err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	return nil
}
func (l lSys) MakeOperationLogWithMiddleware(r *ghttp.Request) *entity.OperationLog {
	begin := time.Now().UnixMilli()
	r.Middleware.Next()
	var content string
	switch r.Method {
	case "DELETE":
		content = r.RequestURI
	case "POST", "PUT":
		content = fmt.Sprint(r.GetBodyString())
		if content == "" {
			content = r.Request.PostForm.Encode()
		}
		if content == "" {
			content = r.Request.Form.Encode()
		}
		if len(content) > 233 {
			content = fmt.Sprint(gstr.SubStrRune(content, 0, 233), "...")
		}
	default:
		r.Middleware.Next()
	}
	end := time.Now().UnixMilli()
	log := entity.OperationLog{
		Uid:      int(l.GetAdminIdFromCtx(r.Context())),
		Content:  content,
		Method:   r.Method,
		Uri:      r.URL.Path,
		Ip:       r.GetClientIp(),
		UseTime:  int(end - begin),
		Response: r.Response.BufferString(),
	}
	return &log
}

// -----------------LoginLog----------------------------

func (l lSys) ListLoginLog(ctx context.Context, req *v1.ListAdminLoginLogReq) ([]*model.AdminLoginLog, int, error) {
	var data = make([]*model.AdminLoginLog, 0)
	db := g.DB("sys").Model(dao.AdminLoginLog.Table() + " t1").
		LeftJoin("s_admin t2 on t1.uid = t2.id").Ctx(ctx)
	if req.Uname != "" {
		db = db.WhereLike("t2.uname", xstr.Like(req.Uname))
	}
	if req.Ip != "" {
		db = db.WhereLike("t1.ip", xstr.Like(req.Ip))
	}
	count, err := db.Count()
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, 0, err
	}
	if err = db.Page(int(req.Page), int(req.Size)).Order("t1.id desc").
		Fields("t1.*,t2.uname").Scan(&data); err != nil {
		g.Log().Error(ctx, err)
		return nil, 0, err
	}
	return data, count, nil
}
func (l lSys) DelLoginLog(ctx context.Context, id uint64) error {
	if _, err := dao.AdminLoginLog.Ctx(ctx).Delete("id", id); err != nil {
		return err
	}
	return nil
}
func (l lSys) ClearLoginLog(ctx context.Context) error {
	if _, err := dao.AdminLoginLog.Ctx(ctx).Delete("id !=0"); err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	return nil
}

// -----------------File----------------------------

func (l lSys) UploadFile(ctx context.Context, group int, file *ghttp.UploadFile) (string, error) {
	if file == nil {
		return "", consts.ErrImgCannotBeEmpty
	}
	fileName := fmt.Sprint(grand.S(6), path.Ext(file.Filename))
	file.Filename = fileName
	datePre := time.Now().Format("2006/01")
	rootFilePath, err := g.Cfg().Get(ctx, "server.rootFilePath")
	if err != nil {
		return "", err
	}
	rootPath := gfile.Pwd() + rootFilePath.String()
	mixPath := fmt.Sprintf("%s/%d/%s/", rootPath, group, datePre)
	_, err = file.Save(mixPath)
	if err != nil {
		return "", err
	}
	dbName := fmt.Sprintf("%d/%s/%s", group, datePre, file.Filename)
	f := entity.File{
		Url:    dbName,
		Group:  group,
		Status: 1,
	}
	if err = l.AddFile(ctx, &f); err != nil {
		return "", err
	}
	return dbName, err
}
func (l lSys) GetFileById(ctx context.Context, id uint64) (*entity.File, error) {
	var data entity.File
	one, err := dao.File.Ctx(ctx).WherePri(id).One()
	if err != nil {
		return nil, err
	}
	if one.IsEmpty() {
		return nil, consts.ErrDataNotFound
	}
	if err = one.Struct(&data); err != nil {
		return nil, err
	}
	return &data, nil
}
func (l lSys) ListFile(ctx context.Context, req *v1.ListFileReq) ([]*entity.File, int, error) {
	var data = make([]*entity.File, 0)
	db := dao.File.Ctx(ctx)
	if req.Id != 0 {
		db = db.Where("id", req.Id)
	}
	if req.Url != "" {
		db = db.WhereLike("url", xstr.Like(req.Url))
	}
	if req.Group != 0 {
		db = db.Where("group", req.Group)
	}
	count, err := db.Count()
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, 0, err
	}
	if err = db.Page(int(req.Page), int(req.Size)).Order("id desc").Scan(&data); err != nil {
		g.Log().Error(ctx, err)
		return nil, 0, err
	}
	return data, count, nil
}
func (l lSys) GetRandomIconFromFile(ctx context.Context) (string, error) {
	value, err := dao.File.Ctx(ctx).OrderRandom().Value("url", "`group` = 1")
	if err != nil {
		return "", err
	}
	return value.String(), nil
}
func (l lSys) ListIconsFromFile(ctx context.Context) ([]string, error) {
	array, err := dao.File.Ctx(ctx).Array("url", "`group`=1")
	if err != nil {
		return nil, err
	}
	var r []string
	for _, i := range array {
		r = append(r, i.String())
	}
	return r, nil
}
func (l lSys) DelFile(ctx context.Context, id uint64) error {
	if _, err := dao.File.Ctx(ctx).Delete("id", id); err != nil {
		return err
	}
	return nil
}
func (l lSys) UpdateFile(ctx context.Context, data *v1.UpdateFileReq) error {
	if _, err := dao.File.Ctx(ctx).OmitEmpty().Update(data, "id", data.Id); err != nil {
		return err
	}
	return nil
}
func (l lSys) AddFile(ctx context.Context, e *entity.File) error {
	_, err := dao.File.Ctx(ctx).Insert(e)
	return err
}
func (l lSys) RemoveFile(ctx context.Context, path string) error {
	if !gfile.Exists(path) {
		g.Log().Warningf(ctx, "path:%v is not exists", path)
		return nil
	}
	if !gfile.IsFile(path) {
		g.Log().Warningf(ctx, "path:%v is not file", path)
		return nil
	}
	if err := gfile.Remove(path); err != nil {
		g.Log().Errorf(ctx, "remove File error path is %v,err:%v", path, err.Error())
		return fmt.Errorf("remove file error path is %v", path)
	}
	g.Log().Debugf(ctx, "Remove File success path is %v", path)
	return nil
}
