package logic

import (
	"ciel-admin/apiv1"
	"ciel-admin/internal/consts"
	"ciel-admin/internal/dao"
	"ciel-admin/internal/model/do"
	"ciel-admin/internal/model/entity"
	"ciel-admin/utility/utils/xjwt"
	"ciel-admin/utility/utils/xpwd"
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"strings"
)

var (
	User = user{}
)

type user struct {
}

func (l user) UpdateUname(ctx context.Context, uname string, id uint64) error {
	count, err := dao.User.Ctx(ctx).Count("uname", uname)
	if err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	if count != 0 {
		return consts.ErrUnameExist
	}
	if err = g.Validator().Rules("password").Data(uname).Run(ctx); err != nil {
		return consts.ErrUnameFormat
	}
	if _, err = dao.User.Ctx(ctx).WherePri(id).Update(g.Map{"uname": uname}); err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	return nil
}

func (l user) UpdatePass(ctx context.Context, pass string, id uint64) error {
	if err := g.Validator().Rules("password").Data(pass).Run(ctx); err != nil {
		return consts.ErrPassFormat
	}
	if _, err := dao.User.Ctx(ctx).WherePri(id).Update(g.Map{"pass": xpwd.GenPwd(pass)}); err != nil {
		g.Log().Error(ctx)
		return err
	}
	return nil
}

func (l user) Register(ctx context.Context, uname string, pass string, ip string) (*apiv1.LoginVo, error) {
	var (
		resVo apiv1.LoginVo
	)
	count, err := dao.User.Ctx(ctx).Count("uname", uname)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	if count != 0 {
		return nil, consts.ErrUnameExist
	}
	var (
		userData = entity.User{
			Uname:    uname,
			Nickname: uname,
			JoinIp:   ip,
			Pass:     xpwd.GenPwd(pass),
			Status:   1,
		}
	)
	icon, err := File.RandomUserIcon(ctx)
	if err != nil {
		return nil, err
	}
	userData.Icon = icon
	if err = g.DB().Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		uid, err := tx.Model(dao.User.Table()).InsertAndGetId(userData)
		if err != nil {
			g.Log().Error(ctx, err)
			return err
		}
		userData.Id = uint64(uid)
		var gold = entity.Gold{}
		gold.Uid = userData.Id
		if _, err = tx.Model(dao.Gold.Table()).Insert(gold); err != nil {
			g.Log().Error(ctx, err)
			return err
		}
		vo, err := l.loginVo(ctx, tx, userData.Id)
		resVo = *vo
		return nil
	}); err != nil {
		return nil, err
	}
	return &resVo, nil
}

func (l user) loginVo(ctx context.Context, tx *gdb.TX, id uint64) (*apiv1.LoginVo, error) {
	var data apiv1.LoginVo
	userData, err := dao.User.GetByIdTx(ctx, tx, id)
	if err != nil {
		return nil, err
	}
	data.Uname = userData.Uname
	data.Nickname = userData.Nickname
	data.Email = userData.Email
	data.Phone = userData.Phone
	data.Summary = userData.Summary
	if strings.HasPrefix(userData.Icon, "http") {
		data.Icon = userData.Icon
	} else {
		data.Icon = consts.ImgPrefix + userData.Icon
	}
	gold, err := dao.Gold.GetByUidTx(ctx, tx, id)
	if err != nil {
		return nil, err
	}
	data.GoldStatus = gold.Status
	token, err := xjwt.GenToken(userData.Uname, userData.Id, 0)
	if err != nil {
		return nil, err
	}
	data.Token = token
	return &data, nil
}

func (l user) Login(ctx context.Context, uname string, pass string, ip string) (*apiv1.LoginVo, error) {
	var data apiv1.LoginVo
	userData, err := dao.User.GetByUname(ctx, uname)
	if err != nil {
		return nil, err
	}
	if userData.PassErrorCount > 6 {
		return nil, consts.ErrPassErrorTooMany
	}
	if !xpwd.ComparePassword(userData.Pass, pass) {
		userData.PassErrorCount++
		if userData.PassErrorCount >= 6 {
			userData.Status = consts.UserStatusLock
		}
		if _, err = dao.User.Ctx(ctx).Save(userData); err != nil {
			return nil, err
		}
		return nil, consts.ErrLogin
	}
	if err = g.DB().Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		var loginLog = entity.UserLoginLog{
			Uid: userData.Id,
			Ip:  ip,
		}
		if userData.PassErrorCount != 0 {
			if _, err = tx.Model(dao.User.Table()).WherePri(userData.Id).Update(do.User{PassErrorCount: 0}); err != nil {
				g.Log().Error(ctx, err)
				return err
			}
		}
		if _, err = tx.Model(dao.UserLoginLog.Table()).Insert(loginLog); err != nil {
			g.Log().Error(ctx, err)
			return err
		}
		vo, err := l.loginVo(ctx, tx, userData.Id)
		if err != nil {
			return err
		}
		data = *vo
		return nil
	}); err != nil {
		return nil, err
	}
	return &data, nil
}

func (l user) UpdatePassByUser(ctx context.Context, oldPass string, newPass string, id uint64) error {
	userData, err := dao.User.GetById(ctx, id)
	if err != nil {
		return err
	}
	if !xpwd.ComparePassword(userData.Pass, oldPass) {
		return consts.ErrOldPassNotMatch
	}
	data := do.User{Pass: xpwd.GenPwd(newPass)}
	if _, err := dao.User.Ctx(ctx).WherePri(id).Update(data); err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	return nil
}

func (l user) UpdateNickname(ctx context.Context, nickname string, uid uint64) error {
	if len(nickname) > 16 {
		return consts.ErrMaxLengthSixTy
	}
	if _, err := dao.User.Ctx(ctx).WherePri(uid).Update(do.User{Nickname: nickname}); err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	return nil
}

func (l user) UpdateIcon(ctx context.Context, icon string, uid uint64) error {
	if _, err := dao.User.Ctx(ctx).WherePri(uid).Data(do.User{Icon: icon}).Update(); err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	return nil
}
