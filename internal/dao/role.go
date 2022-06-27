// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"ciel-admin/internal/consts"
	"ciel-admin/internal/dao/internal"
	"ciel-admin/internal/model/entity"
	"context"
)

// internalRoleDao is internal type for wrapping internal DAO implements.
type internalRoleDao = *internal.RoleDao

// roleDao is the data access object for table s_role.
// You can define custom methods on it to extend its functionality as you wish.
type roleDao struct {
	internalRoleDao
}

func (d roleDao) GetById(ctx context.Context, rid interface{}) (*entity.Role, error) {
	var data entity.Role
	one, err := d.Ctx(ctx).One("id", rid)
	if err != nil {
		return nil, err
	}
	if one.IsEmpty() {
		return nil, consts.ErrDataNotFound
	}
	err = one.Struct(&data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

var (
	// Role is globally public accessible object for table s_role operations.
	Role = roleDao{
		internal.NewRoleDao(),
	}
)

// Fill with you ideas below.
