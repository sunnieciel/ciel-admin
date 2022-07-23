package sys

import (
	"ciel-admin/internal/consts"
	"ciel-admin/internal/dao"
	"ciel-admin/internal/model/bo"
	"ciel-admin/internal/model/do"
	"ciel-admin/internal/service/sys/view"
	"ciel-admin/utility/utils/xpwd"
	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcache"
	"strings"
	"time"
)

func Login(ctx context.Context, id, code, uname string, pwd string, ip string) error {
	if !Store.Verify(id, code, true) {
		return errors.New("验证码错误")
	}
	admin, err := dao.Admin.GetByUname(ctx, uname)
	if err != nil {
		return err
	}
	if !xpwd.ComparePassword(admin.Pwd, pwd) {
		return consts.ErrLogin
	}

	if admin.Status == 2 {
		return consts.ErrAuthNotEnough
	}
	menus, err := Menus(ctx, admin.Rid, -1)
	if err != nil {
		return err
	}
	if err = setAdmin(ctx, &bo.Admin{Admin: admin, Menus: menus}); err != nil {
		return err
	}
	if _, err = dao.AdminLoginLog.Ctx(ctx).Insert(do.AdminLoginLog{Uid: admin.Id, Ip: ip}); err != nil {
		return err
	}
	dao.Admin.Ctx(ctx).Update(do.Admin{UnreadMsgCount: 1}, "id", admin.Id)
	return nil
}
func Logout(ctx context.Context) error {
	return RemoveAdmin(ctx)
}
func UpdateAdminPwd(ctx context.Context, pwd string, pwd2 string) error {
	admin, err := GetAdmin(ghttp.RequestFromCtx(ctx))
	if err != nil {
		return err
	}
	u, err := dao.Admin.GetByUname(ctx, admin.Admin.Uname)
	if err != nil {
		return err
	}
	if !xpwd.ComparePassword(u.Pwd, pwd) {
		return errors.New("old password not match")
	}
	u.Pwd = xpwd.GenPwd(pwd2)
	err = RemoveAdmin(ctx)
	if err != nil {
		return err
	}
	return dao.Admin.Update(ctx, u)
}
func UpdateAdminUname(ctx context.Context, id, uname interface{}) error {
	count, err := dao.Admin.Ctx(ctx).Count("uname", uname)
	if err != nil {
		return err
	}
	if count != 0 {
		return consts.ErrUnameExist
	}
	if _, err = dao.Admin.Ctx(ctx).Update(g.Map{"uname": uname}, "id", id); err != nil {
		return err
	}
	return nil
}
func UpdateAdminPwdWithoutOldPwd(ctx context.Context, id, pwd interface{}) error {
	_, err := dao.Admin.Ctx(ctx).Update(g.Map{"pwd": xpwd.GenPwd(pwd.(string))}, "id", id)
	if err != nil {
		return err
	}
	return nil
}
func ClearAdminLog(ctx context.Context) error {
	if _, err := dao.AdminLoginLog.Ctx(ctx).Delete("id is not null"); err != nil {
		return err
	}

	return nil
}

func GetAllAdminOptions(ctx context.Context) (*gvar.Var, error) {
	return gcache.GetOrSetFunc(ctx, "", func(ctx context.Context) (value interface{}, err error) {
		all, err := dao.Admin.Ctx(ctx).All()
		options := make([]string, 0)
		for index, i := range all {
			options = append(options, fmt.Sprintf("%v:%v:%v", i["id"], i["uname"], view.SwitchTagClass(index)))
		}
		return gvar.New(strings.Join(options, ",")), nil
	}, time.Second*10)
}

func AddAdminUnReadMsg(ctx context.Context, uid int) error {
	if _, err := dao.Admin.Ctx(ctx).Where("id", uid).Increment("unread_msg_count", 1); err != nil {
		return err
	}
	return nil
}

// GetAdminUnreadMsgCount get admin unread msg count
func GetAdminUnreadMsgCount(ctx context.Context) (*gvar.Var, error) {
	admin, err := GetAdmin(ghttp.RequestFromCtx(ctx))
	if err != nil {
		return nil, err
	}
	v, err := dao.Admin.Ctx(ctx).Value("unread_msg_count", "id", admin.Admin.Id)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// ClearUnreadMsg clear unread msg
func ClearUnreadMsg(ctx context.Context) error {
	admin, err := GetAdmin(ghttp.RequestFromCtx(ctx))
	if err != nil {
		return err
	}
	_, err = dao.Admin.Ctx(ctx).Update(do.Admin{UnreadMsgCount: 0}, "id", admin.Admin.Id)
	if err != nil {
		return err
	}
	return nil
}

//ClearAdminMessage clear admin message by group
func ClearAdminMessage(ctx context.Context, group string) error {
	if _, err := dao.AdminMessage.Ctx(ctx).Where("group", group).Delete(); err != nil {
		return err
	}
	return nil
}
