package logic

import (
	"ciel-admin/internal/dao"
	"ciel-admin/internal/model/do"
	"ciel-admin/internal/model/entity"
	"ciel-admin/utility/utils/xpwd"
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/guid"
)

var Gold = lGold{}

type lGold struct {
}

func (l lGold) UpdatePassByAdmin(ctx context.Context, pass string, uid uint64) error {
	if _, err := dao.Gold.Ctx(ctx).Update(do.Gold{Pass: xpwd.GenPwd(pass)}, "uid", uid); err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	return nil
}

func (l lGold) TopUpByAdmin(ctx context.Context, t int, uid uint64, amount float64, desc string) error {
	if err := g.DB().Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		// 修改用户金币
		goldInfo, err := l.ChangeGold(ctx, tx, t, uid, amount)
		if err != nil {
			return err
		}
		// 创建账变记录
		transId := guid.S()
		if desc == "" {
			desc = "人工充值"
		}
		if err = GoldChangeLog.Add(ctx, tx, transId, t, uid, amount, goldInfo.Balance, desc); err != nil {
			return err
		}
		// 创建账变统计
		if err = GoldStatisticsLog.Add(ctx, tx, t, uid, amount); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

func (l lGold) DeductByAdmin(ctx context.Context, t int, uid uint64, amount float64) error {
	if err := g.DB().Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		goldInfo, err := l.ChangeGold(ctx, tx, t, uid, amount)
		if err != nil {
			return err
		}
		// 创建账变记录
		transId := guid.S()
		if err = GoldChangeLog.Add(ctx, tx, transId, t, uid, amount, goldInfo.Balance, "人工扣除"); err != nil {
			return err
		}
		// 创建账变统计
		if err = GoldStatisticsLog.Add(ctx, tx, t, uid, amount); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

func (l lGold) ChangeGold(ctx context.Context, tx *gdb.TX, t int, uid uint64, amount float64) (*entity.Gold, error) {
	gold, err := dao.Gold.GetByUidTx(ctx, tx, uid)
	if err != nil {
		return nil, err
	}
	gold.Balance += amount
	if gold.Balance < 0 {
		gold.Balance = 0
	}
	var data = do.Gold{Balance: gold.Balance}
	if _, err = tx.Model(dao.Gold.Table()).WherePri(gold.Id).Data(data).Update(); err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	return gold, nil
}
