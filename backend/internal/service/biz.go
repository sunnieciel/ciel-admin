package service

import (
	"context"
	"freekey-backend/api"
	v1 "freekey-backend/api/v1"
	"freekey-backend/internal/consts"
	"freekey-backend/internal/logic"
	"freekey-backend/internal/model"
	"freekey-backend/internal/model/do"
	"freekey-backend/internal/model/entity"
	"freekey-backend/utility/utils/xjwt"
	"freekey-backend/utility/utils/xpwd"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/guid"
	"net/http"
)

var Biz = sBiz{}

type sBiz struct{}

// -----------Banner----------------------

func (s sBiz) AddBanner(ctx context.Context, in *entity.Banner) error {
	return logic.Biz.AddBanner(ctx, in)
}
func (s sBiz) GetBannerById(ctx context.Context, id uint64) (*entity.Banner, error) {
	return logic.Biz.GetBannerById(ctx, id)
}
func (s sBiz) ListBanner(ctx context.Context, req *v1.ListBannerReq) ([]*entity.Banner, *api.PageRes, error) {
	menu, total, err := logic.Biz.ListBanner(ctx, req)
	if err != nil {
		return nil, nil, err
	}
	return menu, logic.Sys.MakePageInfo(req.Page, req.Size, int64(total)), nil
}
func (s sBiz) DelBanner(ctx context.Context, id uint64) error {
	return logic.Biz.DelBanner(ctx, id)
}
func (s sBiz) UpdateBanner(ctx context.Context, data *v1.UpdateBannerReq) error {
	return logic.Biz.UpdateBanner(ctx, data)
}

// -----------User----------------------

func (s sBiz) GetUserById(ctx context.Context, id uint64) (*entity.User, error) {
	return logic.Biz.GetUserById(ctx, id)
}
func (s sBiz) GetUserIdFromCtx(ctx context.Context) uint64 {
	return logic.Biz.GetUserIdFromCtx(ctx)
}
func (s sBiz) GetUserInfo(ctx context.Context, uid uint64) (*v1.GetUserInfoRes, error) {
	var out v1.GetUserInfoRes
	user, err := logic.Biz.GetUserById(ctx, uid)
	if err != nil {
		return nil, nil
	}
	out.Id = user.Id
	out.Uname = user.Uname
	out.Nickname = user.Nickname
	out.Icon = consts.ImgPrefix + user.Icon
	out.Summary = user.Summary
	out.Email = user.Email
	out.Phone = user.Phone

	wallet, err := logic.Biz.GetWalletByUid(ctx, user.Id)
	if err != nil {
		return nil, err
	}
	out.Balance = wallet.Balance
	return &out, nil
}
func (s sBiz) ListUser(ctx context.Context, req *v1.ListUserReq) ([]*entity.User, *api.PageRes, error) {
	menu, total, err := logic.Biz.ListUser(ctx, req)
	if err != nil {
		return nil, nil, err
	}
	return menu, logic.Sys.MakePageInfo(req.Page, req.Size, int64(total)), nil
}
func (s sBiz) DelUser(ctx context.Context, id uint64) error {
	_, err := logic.Biz.GetUserById(ctx, id)
	if err != nil {
		return err
	}
	return logic.Biz.DelUser(ctx, id)
}
func (s sBiz) UpdateUser(ctx context.Context, data *v1.UpdateUserReq) error {
	return logic.Biz.UpdateUser(ctx, data)
}
func (s sBiz) UpdateUserUname(ctx context.Context, uname string, id uint64) error {
	count, err := logic.Biz.CountUserByUname(ctx, uname)
	if err != nil {
		return err
	}
	if count != 0 {
		return consts.ErrUnameExist
	}
	return logic.Biz.UpdateUserUname(ctx, uname, id)
}
func (s sBiz) UpdateUserPassByAdmin(ctx context.Context, pass string, id uint64) error {
	return logic.Biz.UpdateUserPass(ctx, pass, id)
}
func (s sBiz) UpdateUserPass(ctx context.Context, oldPass string, pass string, uid uint64) error {
	d, err := logic.Biz.GetUserById(ctx, uid)
	if err != nil {
		return err
	}
	if err = logic.Biz.CheckUserPass(d, oldPass); err != nil {
		return consts.ErrOldPass
	}
	if err = logic.Biz.UpdateUserPass(ctx, pass, uid); err != nil {
		return err
	}
	return nil
}
func (s sBiz) UpdateUserNickName(ctx context.Context, nickname string, uid uint64) error {
	d, err := logic.Biz.GetUserById(ctx, uid)
	if err != nil {
		return err
	}
	d.Nickname = nickname
	if err = logic.Biz.SaveUser(ctx, d); err != nil {
		return err
	}
	return nil
}
func (s sBiz) UpdateUserIcon(ctx context.Context, icon string, uid uint64) error {
	return logic.Biz.UpdateUserIcon(ctx, icon, uid)
}
func (s sBiz) MiddlewareUserAuth(r *ghttp.Request) {
	userInfo, err := xjwt.UserInfo(r)
	if err != nil {
		r.Response.WriteStatus(http.StatusForbidden, consts.ErrAuth.Error())
		r.Exit()
	}
	r.SetParam(consts.TokenUserIdKey, userInfo.Uid)
	r.Middleware.Next()
}
func (s sBiz) MiddlewareAuthMaybe(r *ghttp.Request) {
	userInfo, err := xjwt.UserInfo(r)
	if err == nil {
		r.SetParam(consts.TokenUserIdKey, userInfo.Uid)
	}
	r.Middleware.Next()
}
func (s sBiz) Register(ctx context.Context, uname string, pass string) (string, error) {
	count, err := logic.Biz.CountUserByUname(ctx, uname)
	if err != nil {
		return "", err
	}
	if count != 0 {
		return "", consts.ErrUnameExist
	}

	var uid int64
	if err = g.DB("sys").Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		d := entity.User{}
		d.Uname = uname
		d.Nickname = uname
		d.Pass = xpwd.GenPwd(pass)
		d.Status = 1
		d.JoinIp = ghttp.RequestFromCtx(ctx).GetClientIp()
		icon, err := logic.Sys.GetRandomIconFromFile(ctx)
		if err != nil {
			return err
		}
		d.Icon = icon
		id, err := logic.Biz.AddUserGetId(ctx, tx, &d)
		if err != nil {
			return err
		}
		uid = id
		if err = logic.Biz.AddWallet(ctx, tx, id); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return "", err
	}
	token, err := xjwt.GenToken(uname, uint64(uid), 0)
	if err != nil {
		return "", err
	}
	return token, nil
}
func (s sBiz) Login(ctx context.Context, uname string, pass string) (string, error) {
	d, err := logic.Biz.GetUserByUname(ctx, uname)
	if err != nil {
		return "", err
	}
	if err = logic.Biz.CheckUserPass(d, pass); err != nil {
		return "", err
	}
	if err = logic.Biz.AddUserLoginLog(ctx, d.Id); err != nil {
		return "", err
	}
	token, err := xjwt.GenToken(d.Uname, d.Id, 0)
	if err != nil {
		return "", err
	}
	return token, nil
}

// -----------UserLoginLog----------------------

func (s sBiz) ListUserLoginLog(ctx context.Context, req *v1.ListUserLoginLogReq) ([]*model.UserLoginLog, *api.PageRes, error) {
	menu, total, err := logic.Biz.ListUserLoginLog(ctx, req)
	if err != nil {
		return nil, nil, err
	}
	return menu, logic.Sys.MakePageInfo(req.Page, req.Size, int64(total)), nil
}
func (s sBiz) DelUserLoginLog(ctx context.Context, id uint64) error {
	return logic.Biz.DelUserLoginLog(ctx, id)
}
func (s sBiz) ClearUserLoginLog(ctx context.Context) error {
	return logic.Biz.ClearUserLoginLog(ctx)
}

// ----------------Wallet-----------------------

func (s sBiz) GetWalletById(ctx context.Context, id uint64) (*entity.Wallet, error) {
	return logic.Biz.GetWalletById(ctx, id)
}
func (s sBiz) GetWalletReport(ctx context.Context, uname string, begin string, end string) (*v1.GetReportRes, error) {
	return logic.Biz.GetWalletReport(ctx, uname, begin, end)
}
func (s sBiz) ListWallet(ctx context.Context, req *v1.ListWalletReq) ([]*model.Wallet, *api.PageRes, error) {
	list, total, err := logic.Biz.ListWallet(ctx, req)
	if err != nil {
		return nil, nil, err
	}
	return list, logic.Sys.MakePageInfo(req.Page, req.Size, int64(total)), nil
}
func (s sBiz) UpdateWallet(ctx context.Context, data *v1.UpdateWalletReq) error {
	return logic.Biz.UpdateWallet(ctx, data)
}
func (s sBiz) UpdateWalletPassByAdmin(ctx context.Context, id uint64, pass string) error {
	return logic.Biz.UpdateWalletPass(ctx, pass, id)
}
func (s sBiz) UpdateWalletMoneyByAdmin(ctx context.Context, req *v1.UpdateWalletByAdminReq) error {
	return g.DB("sys").Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		transId := guid.S()
		if err := logic.Biz.WalletKit(ctx, tx, req.Type, req.Uid, req.Money, transId, req.Desc); err != nil {
			return err
		}
		return nil
	})
}

// --- WalletChangeType -----------------------------------------------------------------

func (s sBiz) AddWalletChangeType(ctx context.Context, in *entity.WalletChangeType) error {
	return logic.Biz.AddWalletChangeType(ctx, in)
}
func (s sBiz) GetWalletChangeTypeById(ctx context.Context, id uint64) (*entity.WalletChangeType, error) {
	return logic.Biz.GetWalletChangeTypeById(ctx, id)
}
func (s sBiz) ListWalletChangeType(ctx context.Context, req *v1.ListWalletChangeTypeReq) ([]*entity.WalletChangeType, *api.PageRes, error) {
	menu, total, err := logic.Biz.ListWalletChangeType(ctx, req)
	if err != nil {
		return nil, nil, err
	}
	return menu, logic.Sys.MakePageInfo(req.Page, req.Size, int64(total)), nil
}
func (s sBiz) ListWalletChangeTypeOptions(ctx context.Context) ([]*v1.ListWalletChangeTypeOptionsRes, error) {
	return logic.Biz.ListWalletChangeTypeOptions(ctx)
}
func (s sBiz) DelWalletChangeType(ctx context.Context, id uint64) error {
	return logic.Biz.DelWalletChangeType(ctx, id)
}
func (s sBiz) UpdateWalletChangeType(ctx context.Context, data *v1.UpdateWalletChangeTypeReq) error {
	return logic.Biz.UpdateWalletChangeType(ctx, data)
}

// --- WalletChangeLog -----------------------------------------------------------------

func (s sBiz) AddWalletChangeLog(ctx context.Context, in *do.WalletChangeLog) error {
	return logic.Biz.AddWalletChangeLog(ctx, in)
}
func (s sBiz) GetWalletChangeLogById(ctx context.Context, id uint64) (*entity.WalletChangeLog, error) {
	return logic.Biz.GetWalletChangeLogById(ctx, id)
}
func (s sBiz) ListWalletChangeLog(ctx context.Context, req *v1.ListWalletChangeLogReq) ([]*model.WalletChangeLog, *api.PageRes, error) {
	menu, total, err := logic.Biz.ListWalletChangeLog(ctx, req)
	if err != nil {
		return nil, nil, err
	}
	return menu, logic.Sys.MakePageInfo(req.Page, req.Size, int64(total)), nil
}
func (s sBiz) DelWalletChangeLog(ctx context.Context, id uint64) error {
	return logic.Biz.DelWalletChangeLog(ctx, id)
}
func (s sBiz) UpdateWalletChangeLog(ctx context.Context, data *v1.UpdateWalletChangeLogReq) error {
	return logic.Biz.UpdateWalletChangeLog(ctx, data)
}

// --- WalletStatisticsLog -----------------------------------------------------------------

func (s sBiz) AddWalletStatisticsLog(ctx context.Context, in *do.WalletStatisticsLog) error {
	return logic.Biz.AddWalletStatisticsLog(ctx, in)
}
func (s sBiz) GetWalletStatisticsLogById(ctx context.Context, id uint64) (*entity.WalletStatisticsLog, error) {
	return logic.Biz.GetWalletStatisticsLogById(ctx, id)
}
func (s sBiz) GetReport(ctx context.Context, uname string, begin string, end string) (*v1.GetReportRes, error) {
	return logic.Biz.GetReport(ctx, uname, begin, end)
}
func (s sBiz) ListWalletStatisticsLog(ctx context.Context, req *v1.ListWalletStatisticsLogReq) ([]*model.WalletStatisticsLog, *api.PageRes, error) {
	menu, total, err := logic.Biz.ListWalletStatisticsLog(ctx, req)
	if err != nil {
		return nil, nil, err
	}
	return menu, logic.Sys.MakePageInfo(req.Page, req.Size, int64(total)), nil
}
func (s sBiz) DelWalletStatisticsLog(ctx context.Context, id uint64) error {
	return logic.Biz.DelWalletStatisticsLog(ctx, id)
}
func (s sBiz) UpdateWalletStatisticsLog(ctx context.Context, data *v1.UpdateWalletStatisticsLogReq) error {
	return logic.Biz.UpdateWalletStatisticsLog(ctx, data)
}

// --- TopUp -----------------------------------------------------------------

func (s sBiz) GetTopUpById(ctx context.Context, id uint64) (*entity.TopUp, error) {
	return logic.Biz.GetTopUpById(ctx, id)
}
func (s sBiz) ListTopUp(ctx context.Context, req *v1.ListTopUpReq) ([]*model.TopUp, *api.PageRes, error) {
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

func (s sBiz) CreateTopUp(ctx context.Context, t int, money float64, desc string, uid uint64) error {
	count, err := logic.Biz.CheckWaitTopUpCount(ctx, uid)
	if err != nil {
		return err
	}
	if count != 0 {
		return consts.ErrHasOrderNotFinish
	}
	if err = logic.Biz.AddTopUp(ctx, t, money, desc, uid); err != nil {
		return err
	}
	return nil
}

func (s sBiz) UpdateTopUpByAdmin(ctx context.Context, t int, id uint64, aid uint64) error {
	topUp, err := logic.Biz.GetTopUpById(ctx, id)
	if err != nil {
		return err
	}
	if topUp.Status != 1 {
		return consts.ErrTopUpStatusIsNotWait
	}

	topUp.Aid = aid
	switch t {
	case 2: // 拒绝
		topUp.Status = 3
		if err = logic.Biz.SaveTopUp(ctx, topUp); err != nil {
			return err
		}
	case 1: // 通过
		if err = g.DB("sys").Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
			topUp.Status = 2
			if err = logic.Biz.SaveTopUpTx(ctx, tx, topUp); err != nil {
				return err
			}
			if err = logic.Biz.WalletKit(ctx, tx, int(topUp.ChangeType), topUp.Uid, topUp.Money, topUp.TransId, ""); err != nil {
				return err
			}
			return nil
		}); err != nil {
			return err
		}
	}
	return nil
}

func (s sBiz) ListTopUpForWeb(ctx context.Context, page int64, size int64, status int, uid uint64) (*api.PageRes, []*model.TopUpForWeb, error) {
	total, list, err := logic.Biz.ListTopUpForWeb(ctx, page, size, status, uid)
	if err != nil {
		return nil, nil, err
	}
	return logic.Sys.MakePageInfo(page, size, total), list, nil
}

func (s sBiz) ListWalletChangeLogForWeb(ctx context.Context, page int64, size int64, t int, uid uint64) (*api.PageRes, []*model.WalletChangeLogForWeb, error) {
	total, list, err := logic.Biz.ListWalletChangeLogForWeb(ctx, page, size, t, uid)
	if err != nil {
		return nil, nil, err
	}
	return logic.Sys.MakePageInfo(page, size, total), list, nil
}
