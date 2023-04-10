package logic

import (
	"context"
	"fmt"
	v1 "freekey-backend/api/v1"
	"freekey-backend/internal/consts"
	"freekey-backend/internal/dao"
	"freekey-backend/internal/model"
	"freekey-backend/internal/model/do"
	"freekey-backend/internal/model/entity"
	"freekey-backend/utility/utils/xpwd"
	"freekey-backend/utility/utils/xstr"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/guid"
	"math"
	"time"
)

var (
	Biz = lBiz{}
)

type lBiz struct {
}

// -----------------Banner----------------------------

func (l lBiz) AddBanner(ctx context.Context, in *entity.Banner) error {
	if _, err := dao.Banner.Ctx(ctx).Insert(in); err != nil {
		return err
	}
	return nil
}
func (l lBiz) GetBannerById(ctx context.Context, id uint64) (*entity.Banner, error) {
	var data entity.Banner
	one, err := dao.Banner.Ctx(ctx).WherePri(id).One()
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
func (l lBiz) ListBanner(ctx context.Context, req *v1.ListBannerReq) ([]*entity.Banner, int, error) {
	var data = make([]*entity.Banner, 0)
	db := dao.Banner.Ctx(ctx)
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
func (l lBiz) DelBanner(ctx context.Context, id uint64) error {
	if _, err := dao.Banner.Ctx(ctx).Delete("id", id); err != nil {
		return err
	}
	return nil
}
func (l lBiz) UpdateBanner(ctx context.Context, data *v1.UpdateBannerReq) error {
	if _, err := dao.Banner.Ctx(ctx).OmitEmpty().Update(data, "id", data.Id); err != nil {
		return err
	}
	return nil
}

// -----------------User----------------------------

func (l lBiz) AddUserLoginLog(ctx context.Context, uid uint64) error {
	if _, err := dao.UserLoginLog.Ctx(ctx).Insert(entity.UserLoginLog{
		Uid: uid,
		Ip:  ghttp.RequestFromCtx(ctx).GetClientIp(),
	}); err != nil {
		return err
	}
	return nil
}
func (l lBiz) AddUserGetId(ctx context.Context, tx gdb.TX, e *entity.User) (int64, error) {
	return tx.Ctx(ctx).Model(dao.User.Table()).InsertAndGetId(e)
}
func (l lBiz) GetUserById(ctx context.Context, id uint64) (*entity.User, error) {
	var data entity.User
	one, err := dao.User.Ctx(ctx).WherePri(id).One()
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
func (l lBiz) GetUserByUname(ctx context.Context, uname string) (*entity.User, error) {
	var data entity.User
	one, err := dao.User.Ctx(ctx).One("uname", uname)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	if one.IsEmpty() {
		return nil, consts.ErrLogin
	}
	if err = one.Struct(&data); err != nil {
		return nil, err
	}
	return &data, nil
}
func (l lBiz) GetUserByPhone(ctx context.Context, phone string) (*entity.User, error) {
	var data entity.User
	one, err := dao.User.Ctx(ctx).One("phone", phone)
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
func (l lBiz) GetUserIdFromCtx(ctx context.Context) uint64 {
	return ghttp.RequestFromCtx(ctx).Get(consts.TokenUserIdKey).Uint64()
}
func (l lBiz) ListUser(ctx context.Context, req *v1.ListUserReq) ([]*entity.User, int, error) {
	var data = make([]*entity.User, 0)
	db := dao.User.Ctx(ctx)
	if req.Id != 0 {
		db = db.WherePri(req.Id)
	}
	if req.Phone != "" {
		db = db.WhereLike("phone", xstr.Like(req.Phone))
	}
	if req.Country != "" {
		db = db.WhereLike("country", xstr.Like(req.Country))
	}
	if req.MemberCode != "" {
		db = db.WhereLike("member_code", xstr.Like(req.MemberCode))
	}
	if req.Vip != "" {
		db = db.Where("vip", req.Vip)
	}
	if req.Boss1 != "" {
		db = db.Where("boss1", req.Boss1)
	}
	if req.Boss2 != "" {
		db = db.Where("boss2", req.Boss2)
	}
	if req.Boss3 != "" {
		db = db.Where("boss3", req.Boss3)
	}
	if req.Uname != "" {
		db = db.WhereLike("uname", xstr.Like(req.Uname))
	}
	if req.JoinIp != "" {
		db = db.WhereLike("join_ip", xstr.Like(req.JoinIp))
	}
	if req.Status != 0 {
		db = db.Where("status", req.Status)
	}
	if req.Desc != "" {
		db = db.WhereLike("desc", req.Desc)
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
func (l lBiz) DelUser(ctx context.Context, id uint64) error {
	if _, err := g.DB("sys").Model("u_user").Delete("id", id); err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	if _, err := g.DB("sys").Model("u_user_login_log").Delete("uid", id); err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	if _, err := g.DB("sys").Model("u_wallet").Delete("uid", id); err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	if _, err := g.DB("sys").Model("u_wallet_change_log").Delete("uid", id); err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	if _, err := g.DB("sys").Model("u_wallet_statistics_log").Delete("uid", id); err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	return nil
}
func (l lBiz) UpdateUser(ctx context.Context, data *v1.UpdateUserReq) error {
	if _, err := dao.User.Ctx(ctx).OmitEmpty().Update(data, "id", data.Id); err != nil {
		return err
	}
	return nil
}
func (l lBiz) UpdateUserUname(ctx context.Context, uname string, id uint64) error {
	if _, err := dao.User.Ctx(ctx).WherePri(id).Update(g.Map{"uname": uname}); err != nil {
		return err
	}
	return nil
}
func (l lBiz) UpdateUserPass(ctx context.Context, pass string, id uint64) error {
	if _, err := dao.User.Ctx(ctx).WherePri(id).Update(g.Map{"pass": xpwd.GenPwd(pass)}); err != nil {
		return err
	}
	return nil
}
func (l lBiz) UpdateUserIcon(ctx context.Context, icon string, uid uint64) error {
	if _, err := dao.User.Ctx(ctx).WherePri(uid).Update(do.User{Icon: icon}); err != nil {
		return err
	}
	return nil
}
func (l lBiz) SaveUser(ctx context.Context, d *entity.User) error {
	if _, err := dao.User.Ctx(ctx).Save(&d); err != nil {
		return err
	}
	return nil
}
func (l lBiz) GetUidByUname(ctx context.Context, uname string) (uint64, error) {
	v, err := gcache.GetOrSetFunc(ctx, fmt.Sprintf("GetUidByUname%s", uname), func(ctx context.Context) (value interface{}, err error) {
		return dao.User.Ctx(ctx).Value("id", "uname", uname)
	}, time.Minute*30)
	if err != nil {
		return 0, err
	}
	return v.Uint64(), err
}
func (l lBiz) GetUnameById(ctx context.Context, uid uint64) (string, error) {
	v, err := gcache.GetOrSetFunc(ctx, fmt.Sprintf("GetUnameById%d", uid), func(ctx context.Context) (value interface{}, err error) {
		return dao.User.Ctx(ctx).Value("uname", "id", uid)
	}, time.Minute*30)
	if err != nil {
		return "", err
	}
	return v.String(), err
}
func (l lBiz) CountUserByUname(ctx context.Context, uname string) (int, error) {
	return dao.User.Ctx(ctx).Count("uname", uname)
}
func (l lBiz) CheckUserPass(user *entity.User, pass string) error {
	if !xpwd.ComparePassword(user.Pass, pass) {
		return consts.ErrLogin
	}
	return nil
}

// -----------------UserLoginLog----------------------------

func (l lBiz) ListUserLoginLog(ctx context.Context, req *v1.ListUserLoginLogReq) ([]*model.UserLoginLog, int, error) {
	var data = make([]*model.UserLoginLog, 0)
	db := g.DB("sys").Model(dao.UserLoginLog.Table() + " t1").LeftJoin("u_user t2 on t1.uid = t2.id").Ctx(ctx)
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
	if err = db.Page(int(req.Page), int(req.Size)).Order("t1.id desc").Fields("t1.*,t2.uname").Scan(&data); err != nil {
		g.Log().Error(ctx, err)
		return nil, 0, err
	}
	return data, count, nil
}
func (l lBiz) DelUserLoginLog(ctx context.Context, id uint64) error {
	if _, err := dao.UserLoginLog.Ctx(ctx).Delete("id", id); err != nil {
		return err
	}
	return nil
}
func (l lBiz) ClearUserLoginLog(ctx context.Context) error {
	if _, err := dao.UserLoginLog.Ctx(ctx).Delete("id !=0"); err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	return nil
}

// ----------------Wallet-----------------------

func (l lBiz) AddWallet(ctx context.Context, tx gdb.TX, uid int64) error {
	_, err := tx.Ctx(ctx).Model(dao.Wallet.Table()).Insert(entity.Wallet{Uid: uint64(uid), Status: 1})
	if err != nil {
		return err
	}
	return nil
}
func (l lBiz) AddChangeLog(ctx context.Context, tx gdb.TX, transId string, t int, uid uint64, amount float64, balance float64, desc string) error {
	data := do.WalletChangeLog{
		TransId: transId,
		Uid:     uid,
		Type:    t,
		Amount:  amount,
		Balance: balance,
		Desc:    desc,
	}
	if _, err := tx.Model(dao.WalletChangeLog.Table()).Insert(&data); err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	return nil
}
func (l lBiz) AddStatisticsLog(ctx context.Context, tx gdb.TX, t int, uid uint64, amount float64) error {
	todayLog, err := l.GetStatisticsTodayLog(ctx, tx, uid)
	if err != nil {
		if err != consts.ErrDataNotFound {
			return err
		}
		now := gtime.Now()
		data := g.Map{
			"uid":                 uid,
			"created_date":        now,
			fmt.Sprintf("t%d", t): math.Abs(amount),
		}
		if _, err = tx.Model(dao.WalletStatisticsLog.Table()).Insert(data); err != nil {
			return err
		}
		return nil
	}
	if _, err = tx.Model(dao.WalletStatisticsLog.Table()).
		WherePri(todayLog.Id).
		Increment(fmt.Sprintf("t%d", t), math.Abs(amount)); err != nil {
		return err
	}
	return nil
}
func (l lBiz) ChangeWallet(ctx context.Context, tx gdb.TX, t int, uid uint64, amount float64) (*entity.Wallet, error) {
	wallet, err := l.GetByUidTx(ctx, tx, uid)
	if err != nil {
		return nil, err
	}
	wallet.Balance += amount
	if wallet.Balance < 0 {
		return nil, consts.ErrBalance
	}
	d := g.Map{}
	d["balance"] = wallet.Balance
	switch t {
	// todo
	}
	if _, err = tx.Model(dao.Wallet.Table()).WherePri(wallet.Id).Data(d).Update(); err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	return wallet, nil
}
func (l lBiz) GetByUidTx(ctx context.Context, tx gdb.TX, id uint64) (*entity.Wallet, error) {
	var data entity.Wallet
	one, err := tx.Ctx(ctx).Model(dao.Wallet.Table()).One("uid", id)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	if one.IsEmpty() {
		g.Log().Errorf(ctx, "%d 钱包信息不存在", id)
		return nil, consts.ErrDataNotFound
	}
	if err = one.Struct(&data); err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	return &data, nil
}
func (l lBiz) GetWalletReport(ctx context.Context, uname string, begin string, end string) (*v1.GetReportRes, error) {
	if begin == "" {
		begin = gtime.Now().AddDate(0, -6, 0).StartOfDay().String()
	}
	db := g.DB("sys").Model(dao.WalletStatisticsLog.Table()+" t1").Ctx(ctx).
		FieldSum("t1.t1", "t1").
		FieldSum("t1.t2", "t2").
		FieldSum("t1.t3", "t3").
		FieldSum("t1.t4", "t4").
		FieldSum("t1.t5", "t5").
		FieldSum("t1.t6", "t6").
		FieldSum("t1.t7", "t7").
		FieldSum("t1.t8", "t8").
		FieldSum("t1.t9", "t9").
		FieldSum("t1.t10", "t10").
		FieldSum("t1.t11", "t11").
		FieldSum("t1.t12", "t12").
		FieldSum("t1.t13", "t13").
		WhereGTE("t1.created_date", begin)
	if end != "" {
		db = db.WhereLTE("t1.created_date", end)
	}
	var out v1.GetReportRes
	if uname != "" {
		db = db.LeftJoin("u_user t2 on t1.uid = t2.id").Where("t2.uname", uname)
	}
	err := db.Scan(&out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
func (l lBiz) GetStatisticsTodayLog(ctx context.Context, tx gdb.TX, uid uint64) (*entity.WalletStatisticsLog, error) {
	var data entity.WalletStatisticsLog
	one, err := tx.Model(dao.WalletStatisticsLog.Table()).Where("uid = ? and created_date>=?", uid, gtime.Date()).One()
	if err != nil {
		g.Log().Error(ctx, err)
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
func (l lBiz) GetStatisticsTodayLogNoTx(ctx context.Context, uid uint64) (*entity.WalletStatisticsLog, error) {
	var data entity.WalletStatisticsLog
	one, err := dao.WalletStatisticsLog.Ctx(ctx).Where("uid = ? and created_date>=?", uid, gtime.Date()).One()
	if err != nil {
		g.Log().Error(ctx, err)
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
func (l lBiz) GetWalletById(ctx context.Context, id uint64) (*entity.Wallet, error) {
	var data entity.Wallet
	one, err := dao.Wallet.Ctx(ctx).WherePri(id).One()
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
func (l lBiz) ListWallet(ctx context.Context, req *v1.ListWalletReq) ([]*model.Wallet, int, error) {
	var data = make([]*model.Wallet, 0)
	db := g.DB("sys").Model(dao.Wallet.Table() + " t1").LeftJoin("u_user t2 on t1.uid = t2.id").Ctx(ctx)
	if req.Trc20Address != "" {
		db = db.WhereLike("t1.trc20_address", req.Trc20Address)
	}
	if req.Uname != "" {
		db = db.WhereLike("t2.uname", xstr.Like(req.Uname))
	}
	if req.Status != "" {
		db = db.Where("t1.status", req.Status)
	}
	if req.Desc != "" {
		db = db.WhereLike("t1.desc", xstr.Like(req.Desc))
	}
	if req.Balance != 0 {
		db = db.WhereGTE("t1.balance", req.Balance)
	}
	count, err := db.Count()
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, 0, err
	}
	if err = db.Page(int(req.Page), int(req.Size)).Order("t1.id desc").Fields("t1.*,t2.uname").Scan(&data); err != nil {
		g.Log().Error(ctx, err)
		return nil, 0, err
	}
	return data, count, nil
}
func (l lBiz) GetWalletByUid(ctx context.Context, uid uint64) (*entity.Wallet, error) {
	var data entity.Wallet
	one, err := dao.Wallet.Ctx(ctx).Where("uid", uid).One()
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
func (l lBiz) UpdateWallet(ctx context.Context, in *v1.UpdateWalletReq) error {
	d := do.Wallet{}
	d.Desc = in.Desc
	if _, err := dao.Wallet.Ctx(ctx).Update(in, "id", in.Id); err != nil {
		return err
	}
	return nil
}
func (l lBiz) UpdateWalletPass(ctx context.Context, pass string, id uint64) error {
	if _, err := dao.Wallet.Ctx(ctx).WherePri(id).Update(do.Wallet{Pass: xpwd.GenPwd(pass)}); err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	return nil
}
func (l lBiz) WalletKit(ctx context.Context, tx gdb.TX, changeType int, uid uint64, money float64, transId string, desc string) error {
	// add user money
	wallet, err := l.ChangeWallet(ctx, tx, changeType, uid, money)
	if err != nil {
		return err
	}
	// add change log
	if err = l.AddChangeLog(ctx, tx, transId, changeType, uid, money, wallet.Balance, desc); err != nil {
		return err
	}
	// add statistics log
	if err = l.AddStatisticsLog(ctx, tx, changeType, uid, money); err != nil {
		return err
	}
	return nil
}

// --- WalletChangeType -----------------------------------------------------------------

func (l lBiz) AddWalletChangeType(ctx context.Context, in *entity.WalletChangeType) error {
	if _, err := dao.WalletChangeType.Ctx(ctx).Insert(in); err != nil {
		return err
	}
	return nil
}
func (l lBiz) GetWalletChangeTypeById(ctx context.Context, id uint64) (*entity.WalletChangeType, error) {
	var data entity.WalletChangeType
	one, err := dao.WalletChangeType.Ctx(ctx).Where("id", id).One()
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
func (l lBiz) ListWalletChangeType(ctx context.Context, req *v1.ListWalletChangeTypeReq) ([]*entity.WalletChangeType, int, error) {
	var data = make([]*entity.WalletChangeType, 0)
	db := dao.WalletChangeType.Ctx(ctx)
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
func (l lBiz) ListWalletChangeTypeOptions(ctx context.Context) ([]*v1.ListWalletChangeTypeOptionsRes, error) {
	var data []*v1.ListWalletChangeTypeOptionsRes
	if err := dao.WalletChangeType.Ctx(ctx).Order("id").Scan(&data); err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	return data, nil
}
func (l lBiz) DelWalletChangeType(ctx context.Context, id uint64) error {
	if _, err := dao.WalletChangeType.Ctx(ctx).Delete("id", id); err != nil {
		return err
	}
	return nil
}
func (l lBiz) UpdateWalletChangeType(ctx context.Context, data *v1.UpdateWalletChangeTypeReq) error {
	if _, err := dao.WalletChangeType.Ctx(ctx).OmitEmpty().Update(data, "id", data.Id); err != nil {
		return err
	}
	return nil
}

// --- WalletChangeLog -----------------------------------------------------------------

func (l lBiz) AddWalletChangeLog(ctx context.Context, in *do.WalletChangeLog) error {
	if _, err := dao.WalletChangeLog.Ctx(ctx).Insert(in); err != nil {
		return err
	}
	return nil
}
func (l lBiz) GetWalletChangeLogById(ctx context.Context, id uint64) (*entity.WalletChangeLog, error) {
	var data entity.WalletChangeLog
	one, err := dao.WalletChangeLog.Ctx(ctx).WherePri(id).One()
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
func (l lBiz) ListWalletChangeLog(ctx context.Context, req *v1.ListWalletChangeLogReq) ([]*model.WalletChangeLog, int, error) {
	var data = make([]*model.WalletChangeLog, 0)
	db := g.DB("sys").Model(dao.WalletChangeLog.Table() + " t1").LeftJoin("u_user t2 on t1.uid = t2.id").Ctx(ctx)
	if req.TransId != "" {
		db = db.WhereLike("t1.trans_id", req.TransId)
	}
	if req.Uname != "" {
		db = db.WhereLike("t2.uname", req.Uname)
	}
	if req.Type != "" {
		db = db.Where("t1.type", req.Type)
	}
	if req.Desc != "" {
		db = db.WhereLike("t1.desc", req.Desc)
	}
	count, err := db.Count()
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, 0, err
	}
	if err = db.Page(int(req.Page), int(req.Size)).Order("t1.id desc").Fields("t1.*,t2.uname").Scan(&data); err != nil {
		g.Log().Error(ctx, err)
		return nil, 0, err
	}
	return data, count, nil
}
func (l lBiz) DelWalletChangeLog(ctx context.Context, id uint64) error {
	if _, err := dao.WalletChangeLog.Ctx(ctx).Delete("id", id); err != nil {
		return err
	}
	return nil
}
func (l lBiz) UpdateWalletChangeLog(ctx context.Context, data *v1.UpdateWalletChangeLogReq) error {
	if _, err := dao.WalletChangeLog.Ctx(ctx).OmitEmpty().Update(data, "id", data.Id); err != nil {
		return err
	}
	return nil
}

// --- WalletStatisticsLog -----------------------------------------------------------------

func (l lBiz) AddWalletStatisticsLog(ctx context.Context, in *do.WalletStatisticsLog) error {
	if _, err := dao.WalletStatisticsLog.Ctx(ctx).Insert(in); err != nil {
		return err
	}
	return nil
}
func (l lBiz) GetReport(ctx context.Context, uname string, begin string, end string) (*v1.GetReportRes, error) {
	if begin == "" {
		begin = gtime.Now().AddDate(0, -6, 0).StartOfDay().String()
	}
	db := g.DB("sys").Model(dao.WalletStatisticsLog.Table()+" t1").Ctx(ctx).
		FieldSum("t1.t1", "t1").
		FieldSum("t1.t2", "t2").
		FieldSum("t1.t3", "t3").
		FieldSum("t1.t4", "t4").
		FieldSum("t1.t5", "t5").
		FieldSum("t1.t6", "t6").
		FieldSum("t1.t7", "t7").
		FieldSum("t1.t8", "t8").
		FieldSum("t1.t9", "t9").
		FieldSum("t1.t10", "t10").
		FieldSum("t1.t11", "t11").
		FieldSum("t1.t12", "t12").
		FieldSum("t1.t13", "t13").
		WhereGTE("t1.created_date", begin)
	if end != "" {
		db = db.WhereLTE("t1.created_date", end)
	}
	var out v1.GetReportRes
	if uname != "" {
		db = db.LeftJoin("u_user t2 on t1.uid = t2.id").Where("t2.uname", uname)
	}
	err := db.Scan(&out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
func (l lBiz) GetWalletStatisticsLogById(ctx context.Context, id uint64) (*entity.WalletStatisticsLog, error) {
	var data entity.WalletStatisticsLog
	one, err := dao.WalletStatisticsLog.Ctx(ctx).WherePri(id).One()
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
func (l lBiz) ListWalletStatisticsLog(ctx context.Context, req *v1.ListWalletStatisticsLogReq) ([]*model.WalletStatisticsLog, int, error) {
	var data = make([]*model.WalletStatisticsLog, 0)
	db := g.DB("sys").Model(dao.WalletStatisticsLog.Table() + " t1").LeftJoin("u_user t2 on t1.uid = t2.id").Ctx(ctx)
	if req.Uname != "" {
		db = db.WhereLike("t2.uname", xstr.Like(req.Uname))
	}
	if req.Begin != "" {
		db = db.WhereGTE("t1.created_date", req.Begin)
	}
	if req.End != "" {
		db = db.WhereLTE("t1.created_date", req.End)
	}
	count, err := db.Count()
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, 0, err
	}
	if err = db.Page(int(req.Page), int(req.Size)).Order("t1.id desc").Fields("t1.*,t2.uname").Scan(&data); err != nil {
		g.Log().Error(ctx, err)
		return nil, 0, err
	}
	return data, count, nil
}
func (l lBiz) DelWalletStatisticsLog(ctx context.Context, id uint64) error {
	if _, err := dao.WalletStatisticsLog.Ctx(ctx).Delete("id", id); err != nil {
		return err
	}
	return nil
}
func (l lBiz) UpdateWalletStatisticsLog(ctx context.Context, data *v1.UpdateWalletStatisticsLogReq) error {
	if _, err := dao.WalletStatisticsLog.Ctx(ctx).OmitEmpty().Update(data, "id", data.Id); err != nil {
		return err
	}
	return nil
}

// --- TopUp -----------------------------------------------------------------

func (l lBiz) AddTopUp(ctx context.Context, t int, money float64, desc string, uid uint64) error {
	d := entity.TopUp{}
	d.Uid = uid
	d.Money = money
	d.Desc = desc
	d.ChangeType = uint(t)
	d.TransId = guid.S()
	d.Status = 1
	d.Ip = ghttp.RequestFromCtx(ctx).GetClientIp()
	if _, err := dao.TopUp.Ctx(ctx).Insert(&d); err != nil {
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
func (l lBiz) ListTopUp(ctx context.Context, in *v1.ListTopUpReq) ([]*model.TopUp, int, error) {
	var d = make([]*model.TopUp, 0)
	db := g.DB("sys").Model(dao.TopUp.Table() + " t1").
		LeftJoin(dao.User.Table() + " t2 on t1.uid = t2.id")

	if in.Uname != "" {
		db = db.WhereLike("t2.uname", xstr.Like(in.Uname))
	}
	if in.TransId != "" {
		db = db.WhereLike("t1.trans_id", xstr.Like(in.TransId))
	}
	if in.ChangeType != "" {
		db = db.Where("t1.change_type", in.ChangeType)
	}
	if in.Desc != "" {
		db = db.WhereLike("t1.desc", xstr.Like(in.Desc))
	}
	if in.Aid != "" {
		db = db.Where("t1.aid", in.Aid)
	}
	if in.Status != "" {
		db = db.Where("t1.status", in.Status)
	}
	count, err := db.Count()
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, 0, err
	}
	if err = db.Page(int(in.Page), int(in.Size)).
		Fields("t1.*,t2.uname").
		Order("t1.id desc").Scan(&d); err != nil {
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

func (l lBiz) CheckWaitTopUpCount(ctx context.Context, uid uint64) (int, error) {
	count, err := dao.TopUp.Ctx(ctx).Count("uid = ? and `status`=1", uid)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (l lBiz) SaveTopUp(ctx context.Context, in *entity.TopUp) error {
	if _, err := dao.TopUp.Ctx(ctx).Save(in); err != nil {
		return err
	}
	return nil
}
func (l lBiz) SaveTopUpTx(ctx context.Context, tx gdb.TX, in *entity.TopUp) error {
	if _, err := tx.Ctx(ctx).Model(dao.TopUp.Table()).Save(in); err != nil {
		return err
	}
	return nil
}

func (l lBiz) ListTopUpForWeb(ctx context.Context, page int64, size int64, status int, uid uint64) (int, []*model.TopUpForWeb, error) {
	var list = make([]*model.TopUpForWeb, 0)
	db := dao.TopUp.Ctx(ctx).Where("uid", uid)
	if status != 0 {
		db = db.Where("status", status)
	}
	count, err := db.Count()
	if err != nil {
		return 0, nil, err
	}
	if err = db.Page(int(page), int(size)).OrderDesc("id").Scan(&list); err != nil {
		return 0, nil, err
	}
	return count, list, nil
}

func (l lBiz) ListWalletChangeLogForWeb(ctx context.Context, page int64, size int64, t int, uid uint64) (int, []*model.WalletChangeLogForWeb, error) {
	var list = make([]*model.WalletChangeLogForWeb, 0)
	db := dao.WalletChangeLog.Ctx(ctx).Where("uid", uid)
	if t != 0 {
		db = db.Where("type", t)
	}
	count, err := db.Count()
	if err != nil {
		return 0, nil, err
	}
	if err = db.Page(int(page), int(size)).OrderDesc("id").Scan(&list); err != nil {
		return 0, nil, err
	}
	return count, list, nil
}
