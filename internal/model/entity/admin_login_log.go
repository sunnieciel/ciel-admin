// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminLoginLog is the golang structure for table admin_login_log.
type AdminLoginLog struct {
	Id        int         `json:"id"        description:""`
	Uid       int         `json:"uid"       description:""`
	Ip        string      `json:"ip"        description:""`
	Area      string      `json:"area"      description:""`
	CreatedAt *gtime.Time `json:"createdAt" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" description:""`
}
