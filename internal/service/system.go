package service

import (
	"ciel-admin/api/v1"
	"ciel-admin/internal/consts"
	"ciel-admin/internal/dao"
	"ciel-admin/internal/logic"
	"ciel-admin/internal/model"
	"ciel-admin/internal/model/entity"
	"ciel-admin/utility/utils/res"
	"ciel-admin/utility/utils/xjwt"
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/glog"
	"net/http"
)

var (
	System = sSystem{}
	Role   = sRole{}
	Admin  = sAdmin{}
	Dict   = sDict{}
	File   = sFile{}
	Gen    = sGen{}
	Ws     = sWs{}
	User   = sUser{}
	Wallet = sWallet{}
)

type sSystem struct{}

func (s sSystem) Init(ctx context.Context) {
	// set log
	g.Log().SetFlags(glog.F_FILE_LONG | glog.F_TIME_DATE | glog.F_TIME_MILLI)
	// bind funcMap
	g.View().BindFuncMap(View.BindFuncMap())
	// set imgPrefix
	get, err := g.Cfg().Get(ctx, "server.imgPrefix")
	if err != nil {
		panic(err)
	}
	consts.ImgPrefix = get.String()
	// setWhiteIps
	if err = logic.Dict.UpdateWhiteIps(ctx); err != nil {
		panic(err)
	}
	xjwt.Init()
}
func (s sSystem) Add(ctx context.Context, table, data interface{}, dbGroup ...string) error {
	return logic.System.Add(ctx, table, data, dbGroup...)
}

func (s sSystem) GetById(ctx context.Context, table, id interface{}, dbGroup ...string) (gdb.Record, error) {
	return logic.System.GetById(ctx, table, id, dbGroup...)
}
func (s sSystem) GetNodeInfo(ctx context.Context, path string) (*entity.Menu, error) {
	return logic.System.GetNodeInfo(ctx, path)
}
func (s sSystem) GetMsgFromSession(r *ghttp.Request) string {
	return logic.System.GetMsgFromSession(r)
}
func (s sSystem) GetAdminId(ctx context.Context) (int, error) {
	session, err := Admin.GetInfoFromSession(ghttp.RequestFromCtx(ctx).Session)
	if err != nil {
		return 0, err
	}
	return session.Admin.Id, nil
}
func (s sSystem) List(ctx context.Context, c model.Search, dbGroup ...string) (int, gdb.List, error) {
	return logic.System.List(ctx, c, dbGroup...)
}
func (s sSystem) ListAllDict(ctx context.Context) (g.Map, error) {
	return logic.System.ListAllDict(ctx)
}
func (s sSystem) ListBanners(ctx context.Context) ([]*v1.BannerRes, error) {
	return logic.System.ListBanners(ctx)
}

func (s sSystem) Del(ctx context.Context, table, id interface{}, dbGroup ...string) error {
	return logic.System.Del(ctx, table, id, dbGroup...)
}
func (s sSystem) Update(ctx context.Context, table, id, data interface{}, dbGroup ...string) error {
	return logic.System.Update(ctx, table, id, data, dbGroup...)
}
func (s sSystem) UpdateMenuSort(ctx context.Context, sort int, id uint64) error {
	return logic.Menu.UpdateGroupSort(ctx, sort, id)
}

func (s sSystem) MiddlewareCORS(r *ghttp.Request) {
	logic.System.MiddlewareCORS(r)
}
func (s sSystem) MiddlewareWhiteIp(r *ghttp.Request) {
	logic.System.MiddlewareWhiteIp(r)
}

type sRole struct{}

func (s sRole) AddApi(ctx context.Context, rid int, apiIds []int) error {
	return logic.Role.AddApi(ctx, rid, apiIds)
}
func (s sRole) AddMenu(ctx context.Context, rid int, menuIds []int) error {
	return logic.Role.AddMenu(ctx, rid, menuIds)
}

func (s sRole) GetById(ctx context.Context, id interface{}) (*entity.Role, error) {
	return logic.Role.GetById(ctx, id)
}
func (s sRole) GetRoleOptions(ctx context.Context) (string, error) {
	return logic.Role.GetRoleOptions(ctx)
}
func (s sRole) ListRoleNoMenus(ctx context.Context, rid interface{}) (gdb.List, error) {
	return logic.Role.ListRoleNoMenus(ctx, rid)
}
func (s sRole) ListRoleNoApis(ctx context.Context, rid interface{}) (gdb.List, error) {
	return logic.Role.ListRoleNoApis(ctx, rid)
}

func (s sRole) DelApis(ctx context.Context, rid interface{}, t int) error {
	return logic.Role.DelApis(ctx, rid, t)
}
func (s sRole) DelMenus(ctx context.Context, rid interface{}) error {
	return logic.Role.DelMenus(ctx, rid)
}

type sAdmin struct{}

func (s sAdmin) Add(ctx context.Context, in entity.Admin) error {
	return logic.Admin.Add(ctx, in)
}
func (s sAdmin) Login(ctx context.Context, id, code, uname, pwd, ip string) error {
	return logic.Admin.Login(ctx, id, code, uname, pwd, ip)
}
func (s sAdmin) AddMessage(ctx context.Context, uname string, title string, url string, t int) error {
	return logic.Admin.AddMessage(ctx, uname, title, url, t)
}

func (s sAdmin) GetInfoFromSession(r *ghttp.Session) (*model.Admin, error) {
	return logic.Session.GetAdmin(r)
}
func (s sAdmin) ListNotifications(ctx context.Context, page int, size int) (int64, []*entity.AdminMessage, error) {
	session, err := s.GetInfoFromSession(ghttp.RequestFromCtx(ctx).Session)
	if err != nil {
		return 0, nil, err
	}

	return logic.Admin.ListNotifications(ctx, page, size, session.Admin.Id)
}

func (s sAdmin) DelNotification(ctx context.Context, id uint64) error {
	session, err := s.GetInfoFromSession(ghttp.RequestFromCtx(ctx).Session)
	if err != nil {
		return err
	}
	return logic.System.DelFun(ctx, "sys", func(ctx context.Context, db gdb.DB) error {
		if _, err = db.Model(dao.AdminMessage.Table()).Delete("id = ? and aid = ?", id, session.Admin.Id); err != nil {
			return err
		}
		return nil
	})
}
func (s sAdmin) DelOperationLogs(ctx context.Context) error {
	return logic.Admin.DelOperationLogs(ctx)
}
func (s sAdmin) DelNotifications(ctx context.Context) error {
	session, err := s.GetInfoFromSession(ghttp.RequestFromCtx(ctx).Session)
	if err != nil {
		return err
	}
	return logic.Admin.DelNotifications(ctx, session.Admin.Id)
}
func (s sAdmin) DelLoginLogs(ctx context.Context) error {
	return logic.Admin.DelLoginLogs(ctx)
}
func (s sAdmin) Logout(ctx context.Context) error {
	return logic.Session.DelAdmin(ctx)
}

func (s sAdmin) UpdatePwd(ctx context.Context, pwd, pwd2 string) error {
	return logic.Admin.UpdatePwd(ctx, pwd, pwd2)
}
func (s sAdmin) UpdateUname(ctx context.Context, id, uname interface{}) error {
	return logic.Admin.UpdateUname(ctx, id, uname)
}
func (s sAdmin) UpdatePwdWithoutOld(ctx context.Context, id, pwd interface{}) error {
	return logic.Admin.UpdatePwdWithoutOldPwd(ctx, id, pwd)
}

func (s sAdmin) MiddlewareAuth(r *ghttp.Request) {
	logic.Admin.MiddlewareAuth(r)
}
func (s sAdmin) MiddlewareLock(r *ghttp.Request) {
	logic.Admin.MiddlewareLock(r)
}
func (s sAdmin) MiddlewareActionLog(r *ghttp.Request) {
	logic.Admin.MiddlewareActionLog(r)
}
func (s sAdmin) MiddlewareUnread(r *ghttp.Request) {
	session, err := s.GetInfoFromSession(r.Session)
	if err != nil {
		r.Middleware.Next()
	} else {
		if err = logic.Admin.MiddlewareUnread(r, session.Admin.Id); err != nil {
			res.Err(err, r)
		}
		r.Middleware.Next()
	}
}

type sDict struct{}

func (s sDict) GetByKey(ctx context.Context, key string) (string, error) {
	return logic.Dict.GetByKeyString(ctx, key)
}
func (s sDict) GetApiGroupOptions(ctx context.Context) (string, error) {
	return logic.Dict.TakeApiGroupOptions(ctx)
}
func (s sDict) UpdateWhiteIps(ctx context.Context, v ...string) error {
	return logic.Dict.UpdateWhiteIps(ctx, v...)
}

type sFile struct{}

func (s sFile) Uploads(ctx context.Context, r *ghttp.Request) error {
	return logic.File.Uploads(ctx, r)
}
func (s sFile) Upload(ctx context.Context, group int) (*v1.UploadFileRes, error) {
	return logic.File.Upload(ctx, group)
}
func (s sFile) GetById(ctx context.Context, id interface{}) (*entity.File, error) {
	return logic.File.GetById(ctx, id)
}

type sGen struct{}

func (s sGen) GenMenuLevel1(ctx context.Context) (string, error) {
	return logic.Gen.MenuLeve1(ctx)
}
func (s sGen) GetTables(ctx context.Context, db string) (string, error) {
	return logic.Gen.TakeTables(ctx, db)
}
func (s sGen) Gen(ctx context.Context, table, group, menu, prefix, apiGroup, htmlGroup, dbGroup string) error {
	return logic.Gen.Gen(ctx, table, group, menu, prefix, apiGroup, htmlGroup, dbGroup)
}

type sWs struct{}

func (s sWs) GetUserWs(r *ghttp.Request) {
	logic.Ws.GetUserWs(r)
}
func (s sWs) GetAdminWs(r *ghttp.Request) {
	logic.Ws.GetAdminWs(r)
}
func (s sWs) NoticeUser(ctx context.Context, uid int, msg interface{}) error {
	return logic.Ws.NoticeUser(ctx, uid, msg)
}
func (s sWs) NoticeUsers(ctx context.Context, msg interface{}) error {
	return logic.Ws.NoticeUsers(ctx, msg)
}
func (s sWs) NoticeAdmin(ctx context.Context, msg interface{}, adminId int) error {
	return logic.Ws.NoticeAdmin(ctx, msg, adminId)
}
func (s sWs) NoticeAdmins(ctx context.Context, msg interface{}) error {
	return logic.Ws.NoticeAdmins(ctx, msg)
}

type sUser struct{}

func (u sUser) Register(ctx context.Context, input *v1.RegisterReq) (*v1.LoginRes, error) {
	return logic.User.Add(ctx, input)
}
func (u sUser) Login(ctx context.Context, input *v1.LoginReq) (*v1.LoginRes, error) {
	loginVo, err := logic.User.Login(ctx, input)
	if err != nil {
		return nil, err
	}
	return loginVo, err
}

func (u sUser) GetUserInfo(ctx context.Context) (*v1.LoginRes, error) {
	uid := logic.User.GetUidFromCtx(ctx)
	return logic.User.GetInfo(ctx, uid)
}
func (u sUser) GetById(ctx context.Context, id uint64) (*entity.User, error) {
	return logic.User.GetById(ctx, id)
}
func (u sUser) GetUidFromCookie(ctx context.Context) uint64 {
	return logic.User.GetUidFromCookie(ctx)
}
func (u sUser) GetUidFromCtx(ctx context.Context) uint64 {
	return ghttp.RequestFromCtx(ctx).Get(consts.UidKey).Uint64()
}
func (u sUser) ListIcons(ctx context.Context) ([]string, error) {
	return logic.File.ListIcons(ctx)
}

func (u sUser) Del(ctx context.Context, id uint64) error {
	return logic.User.Del(ctx, id)
}
func (u sUser) Logout(ctx context.Context) {
	logic.User.Logout(ctx)
}
func (u sUser) DelLoginLogs(ctx context.Context) error {
	return logic.User.DelLoinLogs(ctx)
}

func (u sUser) UpdateUname(ctx context.Context, uname string, id uint64) error {
	return logic.User.UpdateUname(ctx, uname, id)
}
func (u sUser) UpdatePassByAdmin(ctx context.Context, pass string, id uint64) error {
	return logic.User.UpdatePass(ctx, pass, id)
}
func (u sUser) UpdatePassByUser(ctx context.Context, in *v1.UpdatePassReq) error {
	id := logic.User.GetUidFromCtx(ctx)
	return logic.User.UpdatePassByUser(ctx, in, id)
}
func (u sUser) UpdateNickname(ctx context.Context, in *v1.UpdateNicknameReq) error {
	return logic.User.UpdateNickname(ctx, in.Nickname, u.GetUidFromCtx(ctx))
}
func (u sUser) UpdateIcon(ctx context.Context, icon string) error {
	return logic.User.UpdateIcon(ctx, icon, u.GetUidFromCtx(ctx))
}

func (u sUser) MiddlewareAuth(c *ghttp.Request) {
	userInfo, err := xjwt.UserInfo(c)
	if err != nil {
		c.Response.WriteStatus(http.StatusForbidden, consts.ErrAuth.Error())
		c.Exit()
	}
	c.SetParam(consts.UidKey, userInfo.Uid)
	c.Middleware.Next()
}

type sWallet struct{}

func (s sWallet) AddTopUp(ctx context.Context, req *v1.CreateTopUpReq, uid uint64) error {
	return logic.Wallet.AddTopUp(ctx, req.Money, req.ChangeTypeId, uid)
}

func (s sWallet) GetById(ctx context.Context, uid uint64) (*entity.Wallet, error) {
	return logic.Wallet.GetByUid(ctx, uid)
}
func (s sWallet) GetInfo(ctx context.Context) (*v1.WalletInfoRes, error) {
	return logic.Wallet.GetInfo(ctx, User.GetUidFromCtx(ctx))
}
func (s sWallet) GetStatisticsLogReport(ctx context.Context, begin, end, uname string) (gdb.Record, error) {
	return logic.Wallet.GetStatisticsLogReport(ctx, begin, end, uname)
}
func (s sWallet) GetChangeTypeOptions(ctx context.Context) (string, error) {
	return logic.Wallet.GetChangeTypeOptions(ctx)
}
func (s sWallet) GetChangeTypeTopUpOptions(ctx context.Context) (string, error) {
	return logic.Wallet.GetChangeTypeTopUpOptions(ctx)
}
func (s sWallet) GetChangeTypeDeductOptions(ctx context.Context) (string, error) {
	return logic.Wallet.GetChangeTypeDeductOptions(ctx)
}
func (s sWallet) GetStatisticsLogFieldsNeedToBeCountedOptions(ctx context.Context) (string, error) {
	return logic.Wallet.TakeStatisticsLogFieldsNeedToBeCountedOptionsIntoStr(ctx)
}
func (s sWallet) GetStatisticsLogFieldsNeedToBeCountedOptionsIntoArray(ctx context.Context) ([]string, error) {
	return logic.Wallet.TakeStatisticsLogFieldsNeedToBeCountedOptionsIntoArray(ctx)
}

func (s sWallet) ListTopUpCategory(ctx context.Context) ([]*v1.TopUpCategoryRes, error) {
	return logic.Wallet.ListTopUpCategory(ctx)
}
func (s sWallet) ListTopUp(ctx context.Context, req *v1.ListTopUpReq) ([]*model.TopUpItem, *v1.PageRes, error) {
	total, items, err := logic.Wallet.ListTopUpByUid(ctx, req.Page, req.Size, req.Status, User.GetUidFromCtx(ctx))
	if err != nil {
		return nil, nil, err
	}
	return items, Common.GetPageInfo(req.Page, req.Size, total), nil
}
func (s sWallet) ListChangeTypes(ctx context.Context) ([]*v1.ListChangeTypesRes, error) {
	return logic.Wallet.ListChangeTypes(ctx)
}
func (s sWallet) ListChangeLogs(ctx context.Context, req *v1.ListChangeLogReq) ([]*model.ChangeLogItem, *v1.PageRes, error) {
	total, items, err := logic.Wallet.ListChangeLogs(ctx, req.Page, req.Size, req.Type, User.GetUidFromCtx(ctx))
	if err != nil {
		return nil, nil, err
	}

	return items, Common.GetPageInfo(req.Page, req.Size, total), nil
}

func (s sWallet) DelChangeLogs(ctx context.Context) error {
	return logic.Wallet.DelChangeLogs(ctx)
}
func (s sWallet) DelStatisticsLogs(ctx context.Context) error {
	return logic.Wallet.DelStatisticsLogs(ctx)
}

func (s sWallet) UpdatePassByAdmin(ctx context.Context, pass string, uid uint64) error {
	return logic.Wallet.UpdatePassByAdmin(ctx, pass, uid)
}
func (s sWallet) UpdatePass(ctx context.Context, req *v1.WalletUpdatePassReq, uid uint64) error {
	return logic.Wallet.UpdatePass(ctx, req.OldPass, req.NewPass, uid)
}
func (s sWallet) UpdateSetPass(ctx context.Context, req *v1.WalletSetPassReq, uid uint64) error {
	return logic.Wallet.UpdateSetPass(ctx, req.Pass, uid)
}
func (s sWallet) UpdateTopUpApplication(ctx context.Context, orderId uint64, operationType int64, aid int) error {
	return logic.Wallet.UpdateTopUpApplication(ctx, orderId, operationType, aid)
}
func (s sWallet) UpdateTopUpByAdmin(ctx context.Context, t int, uid uint64, amount float64, desc string) error {
	return logic.Wallet.UpdateTopUpByAdmin(ctx, t, uid, amount, desc)
}
func (s sWallet) UpdateDeductByAdmin(ctx context.Context, t int, uid uint64, amount float64) error {
	return logic.Wallet.UpdateDeductByAdmin(ctx, t, uid, amount)
}
