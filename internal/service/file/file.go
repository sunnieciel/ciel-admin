package file

import (
	"ciel-admin/internal/dao"
	"ciel-admin/internal/logic"
	"ciel-admin/internal/model/entity"
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
)

func Upload(ctx context.Context, r *ghttp.Request) error {
	return logic.File.Upload(ctx, r)
}
func GetById(ctx context.Context, id interface{}) (*entity.File, error) {
	return dao.File.GetById(ctx, id)
}
