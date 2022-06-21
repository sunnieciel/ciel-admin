package sys

import (
	"ciel-admin/internal/service/internal/dao"
	"context"
	"github.com/gogf/gf/v2/database/gdb"
)

func ThingOptions(ctx context.Context) (gdb.Result, error) {
	return dao.Thing.Ctx(ctx).
		//WhereNotIn("type", []int{1}).
		Fields("id,name,type").All()
}
