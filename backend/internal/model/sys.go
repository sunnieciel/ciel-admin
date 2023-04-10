package model

import (
	"github.com/gogf/gf/v2/os/gtime"
)

type Admin struct {
	Id       int `json:"id"           description:""`
	Rid      int
	Uname    string  `json:"uname"        description:""`
	Nickname string  `json:"nickname"     description:""`
	Email    string  `json:"email"        description:""`
	Phone    string  `json:"phone"        description:""`
	Menus    []*Menu `json:"menus"`
}

type AdminMsg struct {
	FromUname string `json:"from_uname"`
	Position  string `json:"position"`
	ToUname   string `json:"to_uname"`
	ToUid     uint64 `json:"to_uid"`
	Msg       string `v:"required" json:"msg"`
	Type      string `v:"required" d:"info" json:"type"` // 用于通知类型
}

type Menu struct {
	Id       int     `json:"id"        description:""`
	Pid      int     `json:"pid"       description:""`
	Icon     string  `json:"icon"      description:""`
	BgImg    string  `json:"bgImg"     description:""`
	Name     string  `json:"name"      description:""`
	Path     string  `json:"path"      description:""`
	Sort     float64 `json:"sort"      description:""`
	Type     int     `json:"type"      description:"1normal 2group"`
	Desc     string  `json:"desc"      description:""`
	FilePath string  `json:"filePath"  description:""`
	Status   int     `json:"status"    description:""`
	Children []*Menu `json:"children"`
}

type RoleMenu struct {
	Id       int    `json:"id"`
	RoleName string `json:"role_name"`
	MenuName string `json:"menu_name"`
	Type     int    `json:"type"`
}
type RoleApi struct {
	Id       uint64 `json:"id"`
	RoleName string `json:"role_name"`
	Path     string `json:"path"`
	Group    string `json:"group"`
	Method   string `json:"method"`
	Desc     string `json:"desc"`
}

type OperationLog struct {
	Id        int         `json:"id"        description:""`
	Uid       int         `json:"uid"       description:""`
	Uname     string      `json:"uname"`
	Desc      string      `json:"desc"`
	Content   string      `json:"content"   description:""`
	Response  string      `json:"response"  description:""`
	Method    string      `json:"method"    description:""`
	Uri       string      `json:"uri"       description:""`
	Ip        string      `json:"ip"        description:""`
	UseTime   int         `json:"useTime"   description:""`
	CreatedAt *gtime.Time `json:"createdAt" description:""`
}
type AdminLoginLog struct {
	Id        int         `json:"id"        description:""`
	Uid       int         `json:"uid"       description:""`
	Uname     string      `json:"uname"`
	Ip        string      `json:"ip"        description:""`
	Area      string      `json:"area"      description:""`
	Status    int         `json:"status"    description:""`
	CreatedAt *gtime.Time `json:"createdAt" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" description:""`
}
