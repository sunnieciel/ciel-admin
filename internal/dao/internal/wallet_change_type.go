// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// WalletChangeTypeDao is the data access object for table u_wallet_change_type.
type WalletChangeTypeDao struct {
	table   string                  // table is the underlying table name of the DAO.
	group   string                  // group is the database configuration group name of current DAO.
	columns WalletChangeTypeColumns // columns contains all the column names of Table for convenient usage.
}

// WalletChangeTypeColumns defines and stores column names for table u_wallet_change_type.
type WalletChangeTypeColumns struct {
	Id          string //
	Title       string //
	SubTitle    string //
	Type        string // 1 add; 2 reduce
	Class       string //
	Desc        string //
	Status      string //
	CountStatus string // Whether this field needs statistics, 1 true 2 false
}

// walletChangeTypeColumns holds the columns for table u_wallet_change_type.
var walletChangeTypeColumns = WalletChangeTypeColumns{
	Id:          "id",
	Title:       "title",
	SubTitle:    "sub_title",
	Type:        "type",
	Class:       "class",
	Desc:        "desc",
	Status:      "status",
	CountStatus: "count_status",
}

// NewWalletChangeTypeDao creates and returns a new DAO object for table data access.
func NewWalletChangeTypeDao() *WalletChangeTypeDao {
	return &WalletChangeTypeDao{
		group:   "default",
		table:   "u_wallet_change_type",
		columns: walletChangeTypeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *WalletChangeTypeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *WalletChangeTypeDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *WalletChangeTypeDao) Columns() WalletChangeTypeColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *WalletChangeTypeDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *WalletChangeTypeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *WalletChangeTypeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
