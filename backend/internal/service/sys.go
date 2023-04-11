package service

import (
	"context"
	"errors"
	"fmt"
	"freekey-backend/api"
	"freekey-backend/api/v1"
	"freekey-backend/internal/consts"
	"freekey-backend/internal/logic"
	"freekey-backend/internal/model"
	"freekey-backend/internal/model/entity"
	"freekey-backend/utility/utils/res"
	"freekey-backend/utility/utils/xcaptcha"
	"freekey-backend/utility/utils/xjwt"
	"freekey-backend/utility/utils/xpwd"
	"freekey-backend/utility/utils/xtrans"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/yudeguang/ratelimit"
	"net/http"
	"strings"
	"time"
)

var (
	Sys = sSys{}
	Ws  = sWs{}
)

type sWs struct{}

func (s sWs) GetUserWs(r *ghttp.Request) {
	logic.Ws.GetUserWs(r)
}
func (s sWs) GetAdminWs(r *ghttp.Request) {
	logic.Ws.GetAdminWs(r)
}

type sSys struct{}

func (s sSys) SwitchEnvironment(projectPath string, t int) error {
	backendCurrent := fmt.Sprint(projectPath, "/backend/manifest/config/config.yaml")
	backendDev := fmt.Sprint(projectPath, "/backend/manifest/config/config-dev.yaml")
	backendServer := fmt.Sprint(projectPath, "/backend/manifest/config/config-server.yaml")
	err := logic.Sys.SwitchEnvironment(backendCurrent, backendDev, backendServer, t)
	if err != nil {
		return err
	}
	frontendCurrent := fmt.Sprint(projectPath, "/frontend/.env")
	frontendDev := fmt.Sprint(projectPath, "/frontend/.env-dev")
	frontendServer := fmt.Sprint(projectPath, "/frontend/.env-server")
	err = logic.Sys.SwitchEnvironment(frontendCurrent, frontendDev, frontendServer, t)
	return nil
}

func (sSys) MiddlewareHandlerResponse(r *ghttp.Request) {
	r.Middleware.Next()

	if r.Response.BufferLength() > 0 {
		return
	}

	var (
		msg  string
		err  = r.GetError()
		resp = r.GetHandlerResponse()
		code = gerror.Code(err)
	)
	if err != nil {
		if code == gcode.CodeNil {
			code = gcode.CodeInternalError
		}
		msg = err.Error()
	} else {
		if r.Response.Status > 0 && r.Response.Status != http.StatusOK {
			msg = http.StatusText(r.Response.Status)
			switch r.Response.Status {
			case http.StatusNotFound:
				code = gcode.CodeNotFound
			case http.StatusForbidden:
				code = gcode.CodeNotAuthorized
			default:
				code = gcode.CodeUnknown
			}
			// It creates error as it can be retrieved by other middlewares.
			err = gerror.NewCode(code, msg)
			r.SetError(err)
		} else {
			code = gcode.CodeOK
		}
	}
	lang := r.GetHeader("lang")
	if lang == "" {
		lang = "zh"
	}
	m := xtrans.T(lang, fmt.Sprint("t", code.Code()))
	if gstr.Contains(m, "t") {
		m = msg
	}
	r.Response.WriteJson(ghttp.DefaultHandlerResponse{
		Code:    code.Code(),
		Message: m,
		Data:    resp,
	})
}

// -----------Ws----------------------

func (s sSys) WsSendAdminMsg(ctx context.Context, d *model.AdminMsg) error {
	return logic.Sys.WsSendAdminMsg(ctx, d)
}
func (s sSys) WsNoticeAdmins(ctx context.Context, d *model.AdminMsg) error {
	return logic.Sys.WsNoticeAdmins(ctx, d)
}

func (s sSys) Init(ctx context.Context) {
	// set log
	g.Log().SetFlags(glog.F_FILE_LONG | glog.F_TIME_DATE | glog.F_TIME_MILLI)
	// bind funcMap
	// set imgPrefix
	get, err := g.Cfg().Get(ctx, "server.imgPrefix")
	if err != nil {
		panic(err)
	}
	consts.ImgPrefix = get.String()
	// setWhiteIps
	if err = logic.Sys.UpdateDictWhiteIps(ctx); err != nil {
		panic(err)
	}
	xjwt.Init()
}
func (s sSys) MiddlewareCORS(r *ghttp.Request) {
	logic.Sys.MiddlewareCORS(r)
}
func (s sSys) MiddlewareWhiteIp(r *ghttp.Request) {
	logic.Sys.MiddlewareWhiteIp(r)
}
func (s sSys) MiddleIpRateLimit(r *ghttp.Request) {
	ip := r.GetClientIp()
	ok := IpRateLimitRole.AllowVisit(ip)
	if !ok {
		res.Err(errors.New("YOUR ACCESS IS ABNORMAL"), r)
	}
	r.Middleware.Next()
}
func (s sSys) MiddlewareActionLock(r *ghttp.Request) {
	var (
		ctx = r.Context()
	)
	sessionId := r.Get(r.Session.Id()).Uint64()
	key := fmt.Sprintf("userActionLock%d", sessionId)
	v, err := gcache.Get(ctx, key)
	if err != nil {
		res.Err(err, r)
	}
	if !v.IsEmpty() {
		res.Err(consts.ErrActionFast, r)
	} else {
		if err = gcache.Set(ctx, key, 1, time.Second*10); err != nil {
			res.Err(err, r)
		}
	}
	r.Middleware.Next()
	_, _ = gcache.Remove(ctx, key)
}

var IpRateLimitRole = logic.Sys.CreateRateLimit(func(r *ratelimit.Rule) {
	r.AddRule(time.Hour, 200)
	r.AddRule(time.Minute, 20)
	r.AddRule(time.Second, 5)
})

//  ---------menu------------------

func (s sSys) AddMenu(ctx context.Context, menu *entity.Menu) error {
	localMenu, _ := logic.Sys.GetMenuByName(ctx, menu.Name)
	if localMenu != nil {
		return consts.ErrDataAlreadyExist
	}
	return logic.Sys.AddMenu(ctx, menu)
}
func (s sSys) GetMenuById(ctx context.Context, id uint64) (*entity.Menu, error) {
	return logic.Sys.GetMenuById(ctx, id)
}
func (s sSys) GetMenuBypath(ctx context.Context, path string) (*entity.Menu, error) {
	return logic.Sys.GetMenuByPath(ctx, path)
}
func (s sSys) ListMenu(ctx context.Context, req *v1.ListMenuReq) ([]*entity.Menu, *api.PageRes, error) {
	menu, total, err := logic.Sys.ListMenu(ctx, req)
	if err != nil {
		return nil, nil, err
	}
	return menu, logic.Sys.MakePageInfo(req.Page, req.Size, int64(total)), nil
}
func (s sSys) DelMenu(ctx context.Context, id uint64) error {
	return logic.Sys.DelMenu(ctx, id)
}
func (s sSys) UpdateMenu(ctx context.Context, data *v1.UpdateMenuReq) error {
	return logic.Sys.SaveMenu(ctx, data.Menu)
}
func (s sSys) UpdateMenuSort(ctx context.Context, sort int, id uint64) error {
	change := func(in float64) float64 {
		arr := strings.Split(fmt.Sprintf("%.2f", in), ".")
		resStr := fmt.Sprintf("%d.%s", sort, arr[1])
		return gconv.Float64(resStr)
	}
	pMenu, err := logic.Sys.GetMenuById(ctx, id)
	if err != nil {
		return err
	}
	pMenu.Sort = change(pMenu.Sort)
	if err = logic.Sys.SaveMenu(ctx, pMenu); err != nil {
		return err
	}
	arr, err := logic.Sys.ListMenuByPid(ctx, pMenu.Id)
	if err != nil {
		return err
	}
	for _, i := range arr {
		i.Sort = change(i.Sort)
		if err = logic.Sys.SaveMenu(ctx, i); err != nil {
			return err
		}
	}
	return nil
}

// -----------API----------------------

func (s sSys) AddApi(ctx context.Context, menu *entity.Api) error {
	return logic.Sys.AddApi(ctx, menu)
}
func (s sSys) GetApiById(ctx context.Context, id uint64) (*entity.Api, error) {
	return logic.Sys.GetApiById(ctx, id)
}
func (s sSys) ListApi(ctx context.Context, req *v1.ListApiReq) ([]*entity.Api, *api.PageRes, error) {
	menu, total, err := logic.Sys.ListApi(ctx, req)
	if err != nil {
		return nil, nil, err
	}
	return menu, logic.Sys.MakePageInfo(req.Page, req.Size, int64(total)), nil
}
func (s sSys) DelApi(ctx context.Context, id uint64) error {
	return logic.Sys.DelApi(ctx, id)
}
func (s sSys) UpdateApi(ctx context.Context, data *v1.UpdateApiReq) error {
	return logic.Sys.UpdateApi(ctx, data)
}
func (s sSys) AddGroupApi(ctx context.Context, group string, url string) (int, error) {
	url = "/backend/" + url
	var data = []*entity.Api{
		{Group: group, Url: url, Method: "GET"},
		{Group: group, Url: fmt.Sprint(url, "/list"), Method: "GET"},
		{Group: group, Url: url, Method: "POST"},
		{Group: group, Url: url, Method: "PUT"},
		{Group: group, Url: url, Method: "DELETE"},
	}
	var count int
	for _, i := range data {
		_, err := logic.Sys.GetApiByMethodAndUrl(ctx, i.Method, i.Url)
		if err == consts.ErrDataNotFound {
			count++
			if err = logic.Sys.AddApi(ctx, i); err != nil {
				return 0, err
			}
		}
	}
	return count, nil
}

// -----------Role----------------------

func (s sSys) AddRole(ctx context.Context, role *entity.Role) error {
	return logic.Sys.AddRole(ctx, role)
}
func (s sSys) GetRoleOptions(ctx context.Context) (string, error) {
	all, err := logic.Sys.ListAllRole(ctx)
	if err != nil {
		g.Log().Error(ctx, err)
		return "", err
	}
	return logic.Sys.MixRoleOptions(all)
}
func (s sSys) GetRoleById(ctx context.Context, id uint64) (*entity.Role, error) {
	return logic.Sys.GetRoleById(ctx, id)
}
func (s sSys) ListRole(ctx context.Context, req *v1.ListRoleReq) (*api.PageRes, []*entity.Role, error) {
	data, total, err := logic.Sys.ListRole(ctx, req)
	if err != nil {
		return nil, nil, err
	}
	return logic.Sys.MakePageInfo(req.Page, req.Size, total), data, nil
}
func (s sSys) DelRole(ctx context.Context, id uint64) error {
	return logic.Sys.DelRole(ctx, id)
}
func (s sSys) DelRoleMenus(ctx context.Context, rid interface{}) error {
	_, err := logic.Sys.GetRoleById(ctx, gconv.Uint64(rid))
	if err != nil {
		return err
	}
	return logic.Sys.DelRoleMenusByRid(ctx, rid)
}
func (s sSys) UpdateRole(ctx context.Context, data *v1.UpdateRoleReq) error {
	return logic.Sys.UpdateRole(ctx, data)
}

// -----------RoleMenu----------------------

func (s sSys) AddMenus(ctx context.Context, rid int, menuIds []int) error {
	return logic.Sys.AddMenus(ctx, rid, menuIds)
}
func (s sSys) ListRoleMenu(ctx context.Context, req *v1.ListRoleMenuReq) ([]*model.RoleMenu, *api.PageRes, error) {
	menu, total, err := logic.Sys.ListRoleMenu(ctx, req)
	if err != nil {
		return nil, nil, err
	}
	return menu, logic.Sys.MakePageInfo(req.Page, req.Size, int64(total)), nil
}
func (s sSys) ListRoleNoMenus(ctx context.Context, rid int) ([]*v1.ListRoleNoMenusRes, error) {
	return logic.Sys.ListRoleNoMenus(ctx, rid)
}
func (s sSys) DelRoleMenu(ctx context.Context, id uint64) error {
	return logic.Sys.DelRoleMenu(ctx, id)
}
func (s sSys) ClearRoleMenu(ctx context.Context, rid uint64) error {
	return logic.Sys.ClearRoleMenu(ctx, rid)
}

// -----------Sys----------------------

func (s sSys) AddRoleApis(ctx context.Context, rid int, apiIds []int) error {
	return logic.Sys.AddRoleApis(ctx, rid, apiIds)
}
func (s sSys) ListRoleApi(ctx context.Context, req *v1.ListRoleApiReq) ([]*model.RoleApi, *api.PageRes, error) {
	menu, total, err := logic.Sys.ListRoleApi(ctx, req)
	if err != nil {
		return nil, nil, err
	}
	return menu, logic.Sys.MakePageInfo(req.Page, req.Size, int64(total)), nil
}
func (s sSys) ListRoleNoApis(ctx context.Context, rid interface{}) ([]*v1.ListRoleNoApisRes, error) {
	return logic.Sys.ListRoleNoApis(ctx, rid)
}
func (s sSys) DelRoleApi(ctx context.Context, id uint64) error {
	return logic.Sys.DelRoleApi(ctx, id)
}
func (s sSys) ClearRoleApi(ctx context.Context, rid uint64) error {
	return logic.Sys.ClearRoleApi(ctx, rid)
}

// -----------Admin----------------------

func (s sSys) AddAdmin(ctx context.Context, in *entity.Admin) error {
	count, err := logic.Sys.CountAdminByUname(ctx, in.Uname)
	if err != nil {
		return err
	}
	if count != 0 {
		return consts.ErrUnameExist
	}
	in.Pwd = xpwd.GenPwd(in.Pwd)
	return logic.Sys.AddAdmin(ctx, in)
}
func (s sSys) GetAdminById(ctx context.Context, id uint64) (*entity.Admin, error) {
	return logic.Sys.GetAdminById(ctx, id)
}
func (s sSys) GetAdminUidFromCtx(ctx context.Context) uint64 {
	return logic.Sys.GetAdminIdFromCtx(ctx)
}
func (s sSys) GetAdminInfo(ctx context.Context) (*model.Admin, error) {
	loginInfo, err := logic.Sys.GetAdminJwtInfo(ghttp.RequestFromCtx(ctx))
	if err != nil {
		return nil, err
	}
	admin, err := s.GetAdminById(ctx, gconv.Uint64(loginInfo.Uid))
	if err != nil {
		return nil, err
	}
	menus, err := logic.Sys.ListMenus(ctx, int(loginInfo.Rid), -1)
	if err != nil {
		return nil, err
	}
	out := model.Admin{
		Id:       admin.Id,
		Uname:    admin.Uname,
		Nickname: admin.Nickname,
		Email:    admin.Email,
		Phone:    admin.Phone,
		Menus:    menus,
	}
	return &out, nil

}
func (s sSys) ListAdmin(ctx context.Context, req *v1.ListAdminReq) ([]*entity.Admin, *api.PageRes, error) {
	menu, total, err := logic.Sys.ListAdmin(ctx, req)
	if err != nil {
		return nil, nil, err
	}
	return menu, logic.Sys.MakePageInfo(req.Page, req.Size, int64(total)), nil
}
func (s sSys) DelAdmin(ctx context.Context, id uint64) error {
	return logic.Sys.DelAdmin(ctx, id)
}
func (s sSys) AdminLogin(ctx context.Context, req *v1.AdminLoginReq) (string, error) {
	if !xcaptcha.Store.Verify(req.Id, req.Captcha, true) {
		return "", consts.ErrCaptcha
	}
	admin, err := logic.Sys.GetAdminByUname(ctx, req.Uname)
	if err != nil {
		return "", err
	}
	if err = logic.Sys.CheckAdminPass(admin, req.Pass); err != nil {
		return "", err
	}
	if err = logic.Sys.CheckAdminStatus(admin); err != nil {
		return "", err
	}
	if err = logic.Sys.AddAdminLoginLog(ctx, admin.Id); err != nil {
		return "", err
	}
	token, err := xjwt.GenToken(admin.Uname, uint64(admin.Id), uint64(admin.Rid))
	if err != nil {
		return "", err
	}
	return token, nil

}
func (s sSys) UpdateAdmin(ctx context.Context, data *v1.UpdateAdminReq) error {
	return logic.Sys.UpdateAdmin(ctx, data)
}
func (s sSys) UpdateAdminUname(ctx context.Context, id uint64, uname string) error {
	count, err := logic.Sys.CountAdminByUname(ctx, uname)
	if err != nil {
		return err
	}
	if count != 0 {
		return consts.ErrUnameExist
	}
	return logic.Sys.UpdateAdminUname(ctx, id, uname)
}
func (s sSys) UpdateAdminPass(ctx context.Context, id uint64, pwd string) error {
	return logic.Sys.UpdateAdminPass(ctx, id, pwd)
}
func (s sSys) UpdateAdminPassSelf(ctx context.Context, pass string) error {
	return logic.Sys.UpdateAdminPass(ctx, s.GetAdminUidFromCtx(ctx), pass)
}
func (s sSys) MiddlewareAdminAuth(r *ghttp.Request) {
	adminInfo, err := logic.Sys.GetAdminJwtInfo(r)
	if err != nil {
		r.Response.WriteStatus(http.StatusForbidden, consts.ErrAuth.Error())
		r.Exit()
	}

	if !logic.Sys.CheckRoleApi(r.Context(), int(adminInfo.Rid), r.URL.Path, r.Method) {
		res.Err(consts.ErrAuthNotEnough, r)
	}
	r.SetParam(consts.TokenAdminIdKey, adminInfo.Uid)
	r.Middleware.Next()
}
func (s sSys) MiddlewareAdminActionLog(r *ghttp.Request) {
	switch {
	case r.RequestURI == "/backend/login",
		r.RequestURI == "/backend/operationLog/delClear",
		r.Method == "GET":
		r.Middleware.Next()
		return
	}
	log := logic.Sys.MakeOperationLogWithMiddleware(r)
	ctx := r.Context()
	if err := logic.Sys.AddOperationLog(ctx, log); err != nil {
		g.Log().Error(ctx, err)
	}
}

// -----------Dict----------------------

func (s sSys) AddDict(ctx context.Context, in *entity.Dict) error {
	return logic.Sys.AddDict(ctx, in)
}
func (s sSys) GetDictById(ctx context.Context, id uint64) (*entity.Dict, error) {
	return logic.Sys.GetDictById(ctx, id)
}
func (s sSys) GetDictByKey(ctx context.Context, key string) (string, error) {
	d, err := logic.Sys.GetDictByKey(ctx, key)
	if err != nil {
		return "", err
	}
	return d.V, err
}
func (s sSys) ListDict(ctx context.Context, req *v1.ListDictReq) ([]*entity.Dict, *api.PageRes, error) {
	menu, total, err := logic.Sys.ListDict(ctx, req)
	if err != nil {
		return nil, nil, err
	}
	return menu, logic.Sys.MakePageInfo(req.Page, req.Size, int64(total)), nil
}
func (s sSys) DelDict(ctx context.Context, id uint64) error {
	return logic.Sys.DelDict(ctx, id)
}
func (s sSys) UpdateDict(ctx context.Context, data *v1.UpdateDictReq) error {
	if err := logic.Sys.UpdateDict(ctx, data); err != nil {
		return err
	}
	switch data.K {
	case "white_ips":
		if err := logic.Sys.UpdateDictWhiteIps(ctx); err != nil {
			return err
		}
	}
	return nil
}

// -----------LoginLog----------------------

func (s sSys) ListLoginLog(ctx context.Context, req *v1.ListAdminLoginLogReq) ([]*model.AdminLoginLog, *api.PageRes, error) {
	menu, total, err := logic.Sys.ListLoginLog(ctx, req)
	if err != nil {
		return nil, nil, err
	}
	return menu, logic.Sys.MakePageInfo(req.Page, req.Size, int64(total)), nil
}
func (s sSys) DelLoginLog(ctx context.Context, id uint64) error {
	return logic.Sys.DelLoginLog(ctx, id)
}
func (s sSys) ClearLoginLog(ctx context.Context) error {
	return logic.Sys.ClearLoginLog(ctx)
}

// -----------OperationLog----------------------

func (s sSys) AddOperationLog(ctx context.Context, in *entity.OperationLog) error {
	return logic.Sys.AddOperationLog(ctx, in)
}
func (s sSys) ListOperationLog(ctx context.Context, req *v1.ListOperationLogReq) ([]*model.OperationLog, *api.PageRes, error) {
	menu, total, err := logic.Sys.ListOperationLog(ctx, req)
	if err != nil {
		return nil, nil, err
	}
	return menu, logic.Sys.MakePageInfo(req.Page, req.Size, int64(total)), nil
}
func (s sSys) DelOperationLog(ctx context.Context, id uint64) error {
	return logic.Sys.DelOperationLog(ctx, id)
}
func (s sSys) ClearOperationLog(ctx context.Context) error {
	return logic.Sys.ClearOperationLog(ctx)
}

// -----------File----------------------

func (s sSys) UploadsFiles(ctx context.Context, files ghttp.UploadFiles, group int) ([]string, string, error) {
	DbNames := make([]string, 0)
	for _, i := range files {
		dbName, err := logic.Sys.UploadFile(ctx, group, i)
		if err != nil {
			return nil, "", err
		}
		DbNames = append(DbNames, dbName)
	}
	return DbNames, consts.ImgPrefix, nil
}
func (s sSys) GetFileById(ctx context.Context, id uint64) (*entity.File, error) {
	return logic.Sys.GetFileById(ctx, id)
}
func (s sSys) ListFile(ctx context.Context, req *v1.ListFileReq) ([]*entity.File, *api.PageRes, error) {
	menu, total, err := logic.Sys.ListFile(ctx, req)
	if err != nil {
		return nil, nil, err
	}
	return menu, logic.Sys.MakePageInfo(req.Page, req.Size, int64(total)), nil
}
func (s sSys) DelFile(ctx context.Context, id uint64) error {
	file, err := logic.Sys.GetFileById(ctx, id)
	if err != nil {
		return err
	}
	rootFilePath, err := g.Cfg().Get(ctx, "server.rootFilePath")
	if err != nil {
		g.Log().Error(ctx, err)
		return err
	}

	if err = logic.Sys.RemoveFile(ctx, gfile.Pwd()+rootFilePath.String()+file.Url); err != nil {
		return err
	}
	if err = logic.Sys.DelFile(ctx, id); err != nil {
		return err
	}
	return nil
}
func (s sSys) UpdateFile(ctx context.Context, data *v1.UpdateFileReq) error {
	return logic.Sys.UpdateFile(ctx, data)
}
