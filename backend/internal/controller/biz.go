package controller

import (
	"context"
	"freekey-backend/api"
	v1 "freekey-backend/api/v1"
	"freekey-backend/internal/consts"
	"freekey-backend/internal/service"
	"github.com/gogf/gf/v2/frame/g"
)

var Biz = cBiz{}

type cBiz struct{}

// ----------------User-----------------------

func (c cBiz) GetUserById(ctx context.Context, req *v1.GetUserReq) (res *v1.GetUserRes, err error) {
	data, err := service.Biz.GetUserById(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &v1.GetUserRes{Data: data}, nil
}
func (c cBiz) GetUserInfo(ctx context.Context, _ *v1.GetUserInfoReq) (res *v1.GetUserInfoRes, err error) {
	info, err := service.Biz.GetUserInfo(ctx, service.Biz.GetUserIdFromCtx(ctx))
	if err != nil {
		return nil, err
	}
	return info, nil
}
func (c cBiz) ListUser(ctx context.Context, req *v1.ListUserReq) (res *v1.ListUserRes, err error) {
	User, pageRes, err := service.Biz.ListUser(ctx, req)
	return &v1.ListUserRes{List: User, PageRes: pageRes}, nil
}
func (c cBiz) DelUser(ctx context.Context, req *v1.DelUserReq) (res *api.DefaultRes, err error) {
	if err = service.Biz.DelUser(ctx, req.Id); err != nil {
		return nil, err
	}
	return
}
func (c cBiz) UpdateUser(ctx context.Context, req *v1.UpdateUserReq) (res *api.DefaultRes, err error) {
	if err = service.Biz.UpdateUser(ctx, req); err != nil {
		return nil, err
	}
	return
}
func (c cBiz) UpdateUserUname(ctx context.Context, req *v1.UpdateUnameReq) (res *api.DefaultRes, err error) {

	if err = g.Validator().Rules("password").Data(req.Uname).Run(ctx); err != nil {
		return nil, consts.ErrUnameFormat
	}
	if err = service.Biz.UpdateUserUname(ctx, req.Uname, req.Id); err != nil {
		return nil, err
	}
	return
}
func (c cBiz) UpdateUserPassByAdmin(ctx context.Context, req *v1.UpdatePassForBackendReq) (res *api.DefaultRes, err error) {
	if err = g.Validator().Rules("password").Data(req.Pass).Run(ctx); err != nil {
		return nil, consts.ErrPassFormat
	}
	if err = service.Biz.UpdateUserPassByAdmin(ctx, req.Pass, req.Id); err != nil {
		return nil, err
	}
	return
}
func (c cBiz) UpdateUserPass(ctx context.Context, req *v1.UpdateUserPassReq) (res *api.DefaultRes, err error) {
	if err = service.Biz.UpdateUserPass(ctx, req.OldPass, req.Pass, service.Biz.GetUserIdFromCtx(ctx)); err != nil {
		return nil, err
	}
	return
}
func (c cBiz) UpdateUserNickname(ctx context.Context, req *v1.UpdateNicknameReq) (res *api.DefaultRes, err error) {
	if err = service.Biz.UpdateUserNickName(ctx, req.Nickname, service.Biz.GetUserIdFromCtx(ctx)); err != nil {
		return nil, err
	}
	return
}

func (c cBiz) Register(ctx context.Context, req *v1.RegisterReq) (res *v1.LoginRes, err error) {
	if err = g.Validator().Rules("required|passport").Data(req.Uname).Run(ctx); err != nil {
		return nil, consts.ErrUnameFormat
	}
	if err = g.Validator().Rules("required|password").Data(req.Pass).Run(ctx); err != nil {
		return nil, consts.ErrPassFormat
	}
	token, err := service.Biz.Register(ctx, req.Uname, req.Pass)
	if err != nil {
		return nil, err
	}
	return &v1.LoginRes{Token: token}, nil
}
func (c cBiz) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	if err = g.Validator().Rules("required|passport").Data(req.Uname).Run(ctx); err != nil {
		return nil, consts.ErrUnameFormat
	}
	if err = g.Validator().Rules("required|password").Data(req.Pass).Run(ctx); err != nil {
		return nil, consts.ErrPassFormat
	}
	token, err := service.Biz.Login(ctx, req.Uname, req.Pass)
	if err != nil {
		return nil, err
	}
	return &v1.LoginRes{Token: token}, nil
}

// ----------------UserLoginLog-----------------------

func (c cBiz) ListUserLoginLog(ctx context.Context, req *v1.ListUserLoginLogReq) (res *v1.ListUserLoginLogRes, err error) {
	UserLoginLog, pageRes, err := service.Biz.ListUserLoginLog(ctx, req)
	return &v1.ListUserLoginLogRes{List: UserLoginLog, PageRes: pageRes}, nil
}
func (c cBiz) DelUserLoginLog(ctx context.Context, req *v1.DelUserLoginLogReq) (res *api.DefaultRes, err error) {
	if err = service.Biz.DelUserLoginLog(ctx, req.Id); err != nil {
		return nil, err
	}
	return
}
func (c cBiz) ClearUserLoginLog(ctx context.Context, _ *v1.DelClearUserLoginLogsReq) (res *api.DefaultRes, err error) {
	if err = service.Biz.ClearUserLoginLog(ctx); err != nil {
		return nil, err
	}
	return
}

// ----------------Wallet-----------------------

func (c cBiz) GetWalletById(ctx context.Context, req *v1.GetWalletReq) (res *v1.GetWalletRes, err error) {
	data, err := service.Biz.GetWalletById(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &v1.GetWalletRes{Data: data}, nil
}
func (c cBiz) GetWalletReport(ctx context.Context, req *v1.GetReportReq) (res *v1.GetReportRes, err error) {
	return service.Biz.GetWalletReport(ctx, req.Uname, req.Begin, req.End)
}
func (c cBiz) ListWallet(ctx context.Context, req *v1.ListWalletReq) (res *v1.ListWalletRes, err error) {
	Wallet, pageRes, err := service.Biz.ListWallet(ctx, req)
	return &v1.ListWalletRes{List: Wallet, PageRes: pageRes}, nil
}
func (c cBiz) UpdateWallet(ctx context.Context, req *v1.UpdateWalletReq) (res *api.DefaultRes, err error) {
	if err = service.Biz.UpdateWallet(ctx, req); err != nil {
		return nil, err
	}
	return
}
func (c cBiz) UpdateWalletPassByAdmin(ctx context.Context, req *v1.UpdatePassForBackendReq) (res *api.DefaultRes, err error) {
	if err = service.Biz.UpdateWalletPassByAdmin(ctx, req.Id, req.Pass); err != nil {
		return nil, err
	}
	return
}
func (c cBiz) UpdateWalletMoneyByAdmin(ctx context.Context, req *v1.UpdateWalletByAdminReq) (res *api.DefaultRes, err error) {
	if err = service.Biz.UpdateWalletMoneyByAdmin(ctx, req); err != nil {
		return nil, err
	}
	return
}

// ---WalletChangeType-----------------------------------------------------------------

func (c cBiz) GetWalletChangeTypeById(ctx context.Context, req *v1.GetWalletChangeTypeReq) (res *v1.GetWalletChangeTypeRes, err error) {
	data, err := service.Biz.GetWalletChangeTypeById(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &v1.GetWalletChangeTypeRes{Data: data}, nil
}
func (c cBiz) ListWalletChangeType(ctx context.Context, req *v1.ListWalletChangeTypeReq) (res *v1.ListWalletChangeTypeRes, err error) {
	WalletChangeType, pageRes, err := service.Biz.ListWalletChangeType(ctx, req)
	return &v1.ListWalletChangeTypeRes{List: WalletChangeType, PageRes: pageRes}, nil
}
func (c cBiz) ListWalletChangeTypeOptions(ctx context.Context, _ *v1.ListWalletChangeTypeOptionsReq) (res []*v1.ListWalletChangeTypeOptionsRes, err error) {
	return service.Biz.ListWalletChangeTypeOptions(ctx)
}
func (c cBiz) AddWalletChangeType(ctx context.Context, req *v1.AddWalletChangeTypeReq) (res *api.DefaultRes, err error) {
	if err = service.Biz.AddWalletChangeType(ctx, req.WalletChangeType); err != nil {
		return nil, err
	}
	return
}
func (c cBiz) DelWalletChangeType(ctx context.Context, req *v1.DelWalletChangeTypeReq) (res *api.DefaultRes, err error) {
	if err = service.Biz.DelWalletChangeType(ctx, req.Id); err != nil {
		return nil, err
	}
	return
}
func (c cBiz) UpdateWalletChangeType(ctx context.Context, req *v1.UpdateWalletChangeTypeReq) (res *api.DefaultRes, err error) {
	if err = service.Biz.UpdateWalletChangeType(ctx, req); err != nil {
		return nil, err
	}
	return
}

// --- WalletChangeLog -----------------------------------------------------------------

func (c cBiz) GetWalletChangeLogById(ctx context.Context, req *v1.GetWalletChangeLogReq) (res *v1.GetWalletChangeLogRes, err error) {
	data, err := service.Biz.GetWalletChangeLogById(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &v1.GetWalletChangeLogRes{Data: data}, nil
}
func (c cBiz) ListWalletChangeLog(ctx context.Context, req *v1.ListWalletChangeLogReq) (res *v1.ListWalletChangeLogRes, err error) {
	WalletChangeLog, pageRes, err := service.Biz.ListWalletChangeLog(ctx, req)
	return &v1.ListWalletChangeLogRes{List: WalletChangeLog, PageRes: pageRes}, nil
}
func (c cBiz) ListWalletChangeLogForWeb(ctx context.Context, req *v1.ListWalletChangeLogForWebReq) (res *v1.ListWalletChangeLogForWebRes, err error) {
	pageRes, list, err := service.Biz.ListWalletChangeLogForWeb(ctx, req.Page, req.Size, req.Type, service.Biz.GetUserIdFromCtx(ctx))
	if err != nil {
		return nil, err
	}
	return &v1.ListWalletChangeLogForWebRes{PageRes: pageRes, List: list}, nil
}
func (c cBiz) AddWalletChangeLog(ctx context.Context, req *v1.AddWalletChangeLogReq) (res *api.DefaultRes, err error) {
	if err = service.Biz.AddWalletChangeLog(ctx, req.WalletChangeLog); err != nil {
		return nil, err
	}
	return
}
func (c cBiz) DelWalletChangeLog(ctx context.Context, req *v1.DelWalletChangeLogReq) (res *api.DefaultRes, err error) {
	if err = service.Biz.DelWalletChangeLog(ctx, req.Id); err != nil {
		return nil, err
	}
	return
}
func (c cBiz) UpdateWalletChangeLog(ctx context.Context, req *v1.UpdateWalletChangeLogReq) (res *api.DefaultRes, err error) {
	if err = service.Biz.UpdateWalletChangeLog(ctx, req); err != nil {
		return nil, err
	}
	return
}

// --- WalletStatisticsLog -----------------------------------------------------------------

func (c cBiz) GetWalletStatisticsLogById(ctx context.Context, req *v1.GetWalletStatisticsLogReq) (res *v1.GetWalletStatisticsLogRes, err error) {
	data, err := service.Biz.GetWalletStatisticsLogById(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &v1.GetWalletStatisticsLogRes{Data: data}, nil
}
func (c cBiz) ListWalletStatisticsLog(ctx context.Context, req *v1.ListWalletStatisticsLogReq) (res *v1.ListWalletStatisticsLogRes, err error) {
	WalletStatisticsLog, pageRes, err := service.Biz.ListWalletStatisticsLog(ctx, req)
	return &v1.ListWalletStatisticsLogRes{List: WalletStatisticsLog, PageRes: pageRes}, nil
}
func (c cBiz) AddWalletStatisticsLog(ctx context.Context, req *v1.AddWalletStatisticsLogReq) (res *api.DefaultRes, err error) {
	if err = service.Biz.AddWalletStatisticsLog(ctx, req.WalletStatisticsLog); err != nil {
		return nil, err
	}
	return
}
func (c cBiz) DelWalletStatisticsLog(ctx context.Context, req *v1.DelWalletStatisticsLogReq) (res *api.DefaultRes, err error) {
	if err = service.Biz.DelWalletStatisticsLog(ctx, req.Id); err != nil {
		return nil, err
	}
	return
}
func (c cBiz) UpdateWalletStatisticsLog(ctx context.Context, req *v1.UpdateWalletStatisticsLogReq) (res *api.DefaultRes, err error) {
	if err = service.Biz.UpdateWalletStatisticsLog(ctx, req); err != nil {
		return nil, err
	}
	return
}

func (c cBiz) UpdateUserIcon(ctx context.Context, req *v1.UpdateIconReq) (res *api.DefaultRes, err error) {
	service.Biz.UpdateUserIcon(ctx, req.Icon, service.Biz.GetUserIdFromCtx(ctx))
	return
}

// ----------------Banner-----------------------

func (c cSys) AddBanner(ctx context.Context, req *v1.AddBannerReq) (res *api.DefaultRes, err error) {
	if err = service.Biz.AddBanner(ctx, req.Banner); err != nil {
		return nil, err
	}
	return
}
func (c cSys) GetBannerById(ctx context.Context, req *v1.GetBannerReq) (res *v1.GetBannerRes, err error) {
	data, err := service.Biz.GetBannerById(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &v1.GetBannerRes{Data: data}, nil
}
func (c cSys) ListBanner(ctx context.Context, req *v1.ListBannerReq) (res *v1.ListBannerRes, err error) {
	Banner, pageRes, err := service.Biz.ListBanner(ctx, req)
	return &v1.ListBannerRes{List: Banner, PageRes: pageRes}, nil
}
func (c cSys) DelBanner(ctx context.Context, req *v1.DelBannerReq) (res *api.DefaultRes, err error) {
	if err = service.Biz.DelBanner(ctx, req.Id); err != nil {
		return nil, err
	}
	return
}
func (c cSys) UpdateBanner(ctx context.Context, req *v1.UpdateBannerReq) (res *api.DefaultRes, err error) {
	if err = service.Biz.UpdateBanner(ctx, req); err != nil {
		return nil, err
	}
	return
}

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
func (c cBiz) CreateTopUp(ctx context.Context, req *v1.CreateTopUpReq) (res *api.DefaultRes, err error) {
	if err = service.Biz.CreateTopUp(ctx, req.Type, req.Money, req.Desc, service.Biz.GetUserIdFromCtx(ctx)); err != nil {
		return nil, err
	}
	return
}
func (c cBiz) UpdateTopUpByAdmin(ctx context.Context, req *v1.UpdateTopUpByAdminReq) (res *api.DefaultRes, err error) {
	if err = service.Biz.UpdateTopUpByAdmin(ctx, req.Type, req.Id, service.Sys.GetAdminUidFromCtx(ctx)); err != nil {
		return nil, err
	}
	return
}
func (c cBiz) ListTopUpForWeb(ctx context.Context, req *v1.ListTopUpForWebReq) (res *v1.ListTopUpForWebRes, err error) {
	pageRes, list, err := service.Biz.ListTopUpForWeb(ctx, req.Page, req.Size, req.Status, service.Biz.GetUserIdFromCtx(ctx))
	if err != nil {
		return nil, err
	}
	return &v1.ListTopUpForWebRes{PageRes: pageRes, List: list}, nil
}
