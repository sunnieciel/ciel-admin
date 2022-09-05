package user

import (
	"ciel-admin/apiv1"
	"ciel-admin/internal/dao"
	"ciel-admin/internal/logic"
	"ciel-admin/internal/model/entity"
	"ciel-admin/utility/utils/xicon"
	"context"
)

func UpdateUname(ctx context.Context, uname string, id uint64) error {
	return logic.User.UpdateUname(ctx, uname, id)
}
func UpdatePassByAdmin(ctx context.Context, pass string, id uint64) error {
	return logic.User.UpdatePass(ctx, pass, id)
}

func UpdatePassByUser(ctx context.Context, oldPass, newPass string, id uint64) error {
	return logic.User.UpdatePassByUser(ctx, oldPass, newPass, id)
}
func UpdateNickname(ctx context.Context, nickname string, uid uint64) error {
	return logic.User.UpdateNickname(ctx, nickname, uid)
}
func UpdateIcon(ctx context.Context, icon string, uid uint64) error {
	return logic.User.UpdateIcon(ctx, icon, uid)

}
func Icons(ctx context.Context) []string {
	var arr = make([]string, 0)
	for i := 0; i < 20; i++ {
		arr = append(arr, xicon.GenIcon())
	}
	return arr
}

func Register(ctx context.Context, uname, pass, ip string) (*apiv1.LoginVo, error) {
	return logic.User.Register(ctx, uname, pass, ip)
}
func Login(ctx context.Context, uname, pass, ip string) (*apiv1.LoginVo, error) {
	return logic.User.Login(ctx, uname, pass, ip)
}
func GetById(ctx context.Context, id uint64) (*entity.User, error) {
	return dao.User.GetById(ctx, id)
}
