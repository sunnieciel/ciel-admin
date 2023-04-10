package v1

import (
	"freekey-backend/api"
	"freekey-backend/internal/model"
	"freekey-backend/internal/model/do"
	"freekey-backend/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

type CaptchaReq struct {
	g.Meta `tags:"后台"`
	Id     string `v:"required" in:"query" dc:"随机验证码ID" `
}
type CaptchaRes struct {
	g.Meta `tags:"后台"`
	Img    string `json:"img" dc:"图片 (base64)"`
}
type GetMenuReqByPathReq struct {
	g.Meta `tags:"后台"`
	Path   string `v:"required" dc:"查询路径"`
	api.Authorization
}
type MenuSortReq struct {
	g.Meta `tags:"后台"`
	Sort   int    `v:"required" dc:"排序的数字"`
	Id     uint64 `v:"required" dc:"一级菜单ID"`
}

// AddMenuReq 添加
type AddMenuReq struct {
	g.Meta `tags:"后台"`
	*entity.Menu
}

// GetMenuReq 获取
type GetMenuReq struct {
	g.Meta `tags:"后台"`
	Id     uint64 `v:"required"`
}
type GetMenuRes struct {
	g.Meta `tags:"后台"`
	Data   *entity.Menu `json:"data"`
}

// ListMenuReq 集合
type ListMenuReq struct {
	g.Meta `tags:"后台"`
	api.PageReq
	Pid  int64
	Name string
}
type ListMenuRes struct {
	g.Meta `tags:"后台"`
	List   []*entity.Menu `json:"list"`
	*api.PageRes
}

// DelMenuReq 删除
type DelMenuReq struct {
	g.Meta `tags:"后台"`
	Id     uint64 `v:"required"`
}

// UpdateMenuReq 修改
type UpdateMenuReq struct {
	g.Meta `tags:"后台"`
	*entity.Menu
}

// --- Role  -------------------------------------------

// AddApiReq 添加
type AddApiReq struct {
	g.Meta `tags:"后台"`
	*entity.Api
}
type AddApiGroupReq struct {
	g.Meta `tags:"后台"`
	Group  string `v:"required"`
	Url    string `v:"required"`
}
type AddApiGroupRes struct {
	g.Meta `tags:"后台"`
	Count  int `json:"count"`
}

// GetApiReq 获取
type GetApiReq struct {
	g.Meta `tags:"后台"`
	Id     uint64 `v:"required"`
}
type GetApiRes struct {
	g.Meta `tags:"后台"`
	Data   *entity.Api `json:"data"`
}

// ListApiReq 集合
type ListApiReq struct {
	g.Meta `tags:"后台"`
	api.PageReq
	Url    string
	Method string
	Group  string
	Type   string
	Desc   string
}
type ListApiRes struct {
	g.Meta `tags:"后台"`
	List   []*entity.Api `json:"list"`
	*api.PageRes
}

// DelApiReq 删除
type DelApiReq struct {
	g.Meta `tags:"后台"`
	Id     uint64 `v:"required"`
}

// UpdateApiReq 修改
type UpdateApiReq struct {
	g.Meta `tags:"后台"`
	*entity.Api
}

// AddRoleReq 添加
type AddRoleReq struct {
	g.Meta `tags:"后台"`
	Data   *entity.Role
}

// GetRoleReq 获取
type GetRoleReq struct {
	g.Meta `tags:"后台"`
	Id     uint64 `v:"required"`
}
type GetRoleRes struct {
	Data *entity.Role `json:"data"`
}
type GetRoleOptionsReq struct {
	g.Meta `tags:"后台"`
}
type GetRoleOptionsRes struct {
	Options string `json:"options"`
}

// ListRoleReq 集合
type ListRoleReq struct {
	g.Meta `tags:"后台"`
	api.PageReq
}
type ListRoleRes struct {
	*api.PageRes
	List []*entity.Role `json:"list"`
}

// DelRoleReq 删除
type DelRoleReq struct {
	g.Meta `tags:"后台"`
	Id     uint64 `v:"required"`
}

// UpdateRoleReq 修改
type UpdateRoleReq struct {
	g.Meta `tags:"后台"`
	*entity.Role
}

//--- RoleApi ---------------------------------------------------------

// AddRoleApiReq 添加
type AddRoleApiReq struct {
	g.Meta `tags:"后台"`
	*entity.RoleApi
}
type AddRoleApisReq struct {
	g.Meta `tags:"后台"`
	Rid    int   `v:"required"`
	Apis   []int `v:"required"`
}

// GetRoleApiReq 获取
type GetRoleApiReq struct {
	g.Meta `tags:"后台"`
	Id     uint64 `v:"required"`
}
type GetRoleApiRes struct {
	g.Meta `tags:"后台"`
	Data   *entity.RoleApi `json:"data"`
}

// ListRoleApiReq 集合
type ListRoleApiReq struct {
	g.Meta `tags:"后台"`
	api.PageReq
	Rid int
}
type ListRoleApiRes struct {
	g.Meta `tags:"后台"`
	List   []*model.RoleApi `json:"list"`
	*api.PageRes
}

type ListRoleNoApisReq struct {
	g.Meta `tags:"后台"`
	Rid    uint64 `v:"required"`
}
type ListRoleNoApisRes struct {
	g.Meta `tags:"后台"`
	Id     uint64 `json:"id"`
	Url    string `json:"url"`
	Method string `json:"method"`
	Group  string `json:"group"`
	Type   int    `json:"type"`
	Desc   string `json:"desc"`
}

type ListRoleNoMenusReq struct {
	g.Meta `tags:"后台"`
	Rid    int
}
type ListRoleNoMenusRes struct {
	g.Meta `tags:"后台"`
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Type   int    `json:"type"`
}

// DelRoleApiReq 删除
type DelRoleApiReq struct {
	g.Meta `tags:"后台"`
	Id     uint64 `v:"required"`
}
type DelRoleApiClearReq struct {
	g.Meta `tags:"后台"`
	Rid    uint64 `v:"required"`
}

// UpdateRoleApiReq 修改
type UpdateRoleApiReq struct {
	g.Meta `tags:"后台"`
	*entity.RoleApi
}

//--- RoleMenu ---------------------------------------------------------

// AddRoleMenuReq 添加
type AddRoleMenuReq struct {
	g.Meta `tags:"后台"`
	*entity.RoleMenu
}

type AddRoleMenusReq struct {
	g.Meta `tags:"后台"`
	Rid    int   `v:"required"`
	Mids   []int `v:"required"`
}

// GetRoleMenuReq 获取
type GetRoleMenuReq struct {
	g.Meta `tags:"后台"`
	Id     uint64 `v:"required"`
}
type GetRoleMenuRes struct {
	g.Meta `tags:"后台"`
	Data   *entity.RoleMenu `json:"data"`
}

// ListRoleMenuReq 集合
type ListRoleMenuReq struct {
	g.Meta `tags:"后台"`
	api.PageReq
	Rid int
}
type ListRoleMenuRes struct {
	g.Meta `tags:"后台"`
	List   []*model.RoleMenu `json:"list"`
	*api.PageRes
}

// DelRoleMenuReq 删除
type DelRoleMenuReq struct {
	g.Meta `tags:"后台"`
	Id     uint64 `v:"required"`
}
type DelClearRoleMenuReq struct {
	g.Meta `tags:"后台"`
	Rid    uint64 `v:"required"`
}

// UpdateRoleMenuReq 修改
type UpdateRoleMenuReq struct {
	g.Meta `tags:"后台"`
	*entity.RoleMenu
}

//--- Aid ---------------------------------------------------------

// AddAdminReq 添加
type AddAdminReq struct {
	g.Meta `tags:"后台"`
	*entity.Admin
}
type AdminLoginReq struct {
	g.Meta  `tags:"后台"`
	Uname   string `v:"required" dc:"用户名"`
	Pass    string `v:"required" dc:"密码"`
	Id      string `v:"required" dc:"验证码ID"`
	Captcha string `v:"required" dc:"验证码"`
}
type AdminLoginRes struct {
	g.Meta `tags:"后台"`
	Token  string `json:"token"`
}

// GetAdminReq 获取
type GetAdminReq struct {
	g.Meta `tags:"后台"`
	Id     uint64 `v:"required"`
}
type GetAdminRes struct {
	g.Meta `tags:"后台"`
	Data   *entity.Admin `json:"data"`
}
type AdminInfoReq struct {
	g.Meta `tags:"后台"`
	api.Authorization
}
type AdminInfoRes struct {
	g.Meta `tags:"后台"`
	Info   *model.Admin  `json:"info"`
	Menus  []*model.Menu `json:"menus"`
}

// ListAdminReq 集合
type ListAdminReq struct {
	g.Meta `tags:"后台"`
	api.PageReq
	Id     uint64
	Uname  string
	Status uint64
	Rid    uint64
}
type ListAdminRes struct {
	g.Meta `tags:"后台"`
	List   []*entity.Admin `json:"list"`
	*api.PageRes
}

// DelAdminReq 删除
type DelAdminReq struct {
	g.Meta `tags:"后台"`
	Id     uint64 `v:"required"`
}

// UpdateAdminReq 修改
type UpdateAdminReq struct {
	g.Meta `tags:"后台"`
	*entity.Admin
}
type UpdateAdminUnameReq struct {
	g.Meta `tags:"后台"`
	Uname  string `v:"required"`
	Id     uint64 `v:"required"`
}
type UpdateAdminPassReq struct {
	g.Meta `tags:"后台"`
	Pass   string `v:"required"`
	Id     uint64 `v:"required"`
}
type UpdateAdminPassSelfReq struct {
	g.Meta `tags:"后台"`
	Pass   string `v:"required"`
}

//--- Dict ---------------------------------------------------------

// AddDictReq 添加
type AddDictReq struct {
	g.Meta `tags:"后台"`
	*entity.Dict
}

// GetDictReq 获取
type GetDictReq struct {
	g.Meta `tags:"后台"`
	Id     uint64 `v:"required"`
}
type GetDictRes struct {
	g.Meta `tags:"后台"`
	Data   *entity.Dict `json:"data"`
}

// ListDictReq 集合
type ListDictReq struct {
	g.Meta `tags:"后台"`
	api.PageReq
}
type ListDictRes struct {
	g.Meta `tags:"后台"`
	List   []*entity.Dict `json:"list"`
	*api.PageRes
}

// DelDictReq 删除
type DelDictReq struct {
	g.Meta `tags:"后台"`
	Id     uint64 `v:"required"`
}

// UpdateDictReq 修改
type UpdateDictReq struct {
	g.Meta `tags:"后台"`
	*entity.Dict
}

//--- File ---------------------------------------------------------

// AddFileReq 添加
type AddFileReq struct {
	g.Meta `tags:"后台"`
	*entity.File
}

// GetFileReq 获取
type GetFileReq struct {
	g.Meta `tags:"后台"`
	Id     uint64 `v:"required"`
}
type GetFileRes struct {
	g.Meta `tags:"后台"`
	Data   *entity.File `json:"data"`
}

// ListFileReq 集合
type ListFileReq struct {
	g.Meta `tags:"后台"`
	Id     int
	api.PageReq
	Url   string
	Group int
}
type ListFileRes struct {
	g.Meta `tags:"后台"`
	List   []*entity.File `json:"list"`
	*api.PageRes
}

// DelFileReq 删除
type DelFileReq struct {
	g.Meta `tags:"后台"`
	Id     uint64 `v:"required"`
}

// UpdateFileReq 修改
type UpdateFileReq struct {
	g.Meta `tags:"后台"`
	*entity.File
}

//--- OperationLog ---------------------------------------------------------

// AddOperationLogReq 添加
type AddOperationLogReq struct {
	g.Meta `tags:"后台"`
	*entity.OperationLog
}

// GetOperationLogReq 获取
type GetOperationLogReq struct {
	g.Meta `tags:"后台"`
	Id     uint64 `v:"required"`
}
type GetOperationLogRes struct {
	g.Meta `tags:"后台"`
	Data   *entity.OperationLog `json:"data"`
}

// ListOperationLogReq 集合
type ListOperationLogReq struct {
	g.Meta `tags:"后台"`
	api.PageReq
	Uname    string
	Ip       string
	Method   string
	Content  string
	Uri      string
	Response string
	Desc     string
}
type ListOperationLogRes struct {
	g.Meta `tags:"后台"`
	List   []*model.OperationLog `json:"list"`
	*api.PageRes
}

// DelOperationLogReq 删除
type DelOperationLogReq struct {
	g.Meta `tags:"后台"`
	Id     uint64 `v:"required"`
}
type DelClearOperationLogReq struct {
	g.Meta `tags:"后台"`
}

// UpdateOperationLogReq 修改
type UpdateOperationLogReq struct {
	g.Meta `tags:"后台"`
	*entity.OperationLog
}

//--- AdminLoginLog ---------------------------------------------------------

// AddAdminLoginLogReq 添加
type AddAdminLoginLogReq struct {
	g.Meta `tags:"后台"`
	*entity.AdminLoginLog
}

// GetAdminLoginLogReq 获取
type GetAdminLoginLogReq struct {
	g.Meta `tags:"后台"`
	Id     uint64 `v:"required"`
}
type GetAdminLoginLogRes struct {
	g.Meta `tags:"后台"`
	Data   *entity.AdminLoginLog `json:"data"`
}

// ListAdminLoginLogReq 集合
type ListAdminLoginLogReq struct {
	g.Meta `tags:"后台"`
	api.PageReq
	Uname string
	Ip    string
}
type ListAdminLoginLogRes struct {
	g.Meta `tags:"后台"`
	List   []*model.AdminLoginLog `json:"list"`
	*api.PageRes
}

// DelAdminLoginLogReq 删除
type DelAdminLoginLogReq struct {
	g.Meta `tags:"后台"`
	Id     uint64 `v:"required"`
}
type DelClearAdminLoginLogReq struct {
	g.Meta `tags:"后台"`
}

// UpdateAdminLoginLogReq 修改
type UpdateAdminLoginLogReq struct {
	g.Meta `tags:"后台"`
	*entity.AdminLoginLog
}

//--- Banner ---------------------------------------------------------

// AddBannerReq 添加
type AddBannerReq struct {
	g.Meta `tags:"后台"`
	*entity.Banner
}

// GetBannerReq 获取
type GetBannerReq struct {
	g.Meta `tags:"后台"`
	Id     uint64 `v:"required"`
}
type GetBannerRes struct {
	g.Meta `tags:"后台"`
	Data   *entity.Banner `json:"data"`
}

// ListBannerReq 集合
type ListBannerReq struct {
	g.Meta `tags:"后台"`
	api.PageReq
}
type ListBannerRes struct {
	g.Meta `tags:"后台"`
	List   []*entity.Banner `json:"list"`
	*api.PageRes
}

// DelBannerReq 删除
type DelBannerReq struct {
	g.Meta `tags:"后台"`
	Id     uint64 `v:"required"`
}

// UpdateBannerReq 修改
type UpdateBannerReq struct {
	g.Meta `tags:"后台"`
	*entity.Banner
}

//--- User ---------------------------------------------------------

// AddUserReq 添加
type AddUserReq struct {
	g.Meta `tags:"后台"`
	*entity.User
}

// GetUserReq 获取
type GetUserReq struct {
	g.Meta `tags:"后台"`
	Id     uint64 `v:"required"`
}
type GetUserRes struct {
	g.Meta `tags:"后台"`
	Data   *entity.User `json:"data"`
}

// ListUserReq 集合
type ListUserReq struct {
	g.Meta `tags:"后台"`
	api.PageReq
	Uname      string
	JoinIp     string
	Status     int
	Id         uint64
	Desc       string
	Phone      string
	Country    string
	MemberCode string
	Vip        string
	Boss1      string
	Boss2      string
	Boss3      string
}
type ListUserRes struct {
	g.Meta `tags:"后台"`
	List   []*entity.User `json:"list"`
	*api.PageRes
}

// DelUserReq 删除
type DelUserReq struct {
	g.Meta `tags:"后台"`
	Id     uint64 `v:"required"`
}

// UpdateUserReq 修改
type UpdateUserReq struct {
	g.Meta `tags:"后台"`
	*entity.User
}
type UpdateUnameReq struct {
	g.Meta `tags:"后台"`
	Uname  string `v:"required"`
	Id     uint64 `v:"required"`
}
type UpdatePassForBackendReq struct {
	g.Meta `tags:"后台"`
	Id     uint64 `v:"required"`
	Pass   string `v:"required"`
}

//--- UserLoginLog ---------------------------------------------------------

// AddUserLoginLogReq 添加
type AddUserLoginLogReq struct {
	g.Meta `tags:"后台"`
	*entity.UserLoginLog
}

// GetUserLoginLogReq 获取
type GetUserLoginLogReq struct {
	g.Meta `tags:"后台"`
	Id     uint64 `v:"required"`
}
type GetUserLoginLogRes struct {
	g.Meta `tags:"后台"`
	Data   *entity.UserLoginLog `json:"data"`
}

// ListUserLoginLogReq 集合
type ListUserLoginLogReq struct {
	g.Meta `tags:"后台"`
	api.PageReq
	Uname string
	Ip    string
}
type ListUserLoginLogRes struct {
	g.Meta `tags:"后台"`
	List   []*model.UserLoginLog `json:"list"`
	*api.PageRes
}

// DelUserLoginLogReq 删除
type DelUserLoginLogReq struct {
	g.Meta `tags:"后台"`
	Id     uint64 `v:"required"`
}
type DelClearUserLoginLogsReq struct {
	g.Meta `tags:"后台"`
}

// UpdateUserLoginLogReq 修改
type UpdateUserLoginLogReq struct {
	g.Meta `tags:"后台"`
	*entity.UserLoginLog
}

//--- WalletChangeType ---------------------------------------------------------

// AddWalletChangeTypeReq 添加
type AddWalletChangeTypeReq struct {
	g.Meta `tags:"后台"`
	*entity.WalletChangeType
}

// GetWalletChangeTypeReq 获取
type GetWalletChangeTypeReq struct {
	g.Meta `tags:"后台"`
	Id     uint64 `v:"required"`
}
type GetWalletChangeTypeRes struct {
	g.Meta `tags:"后台"`
	Data   *entity.WalletChangeType `json:"data"`
}

// ListWalletChangeTypeReq 集合
type ListWalletChangeTypeReq struct {
	g.Meta `tags:"后台"`
	api.PageReq
}
type ListWalletChangeTypeRes struct {
	g.Meta `tags:"后台"`
	List   []*entity.WalletChangeType `json:"list"`
	*api.PageRes
}

type ListWalletChangeTypeOptionsReq struct {
	g.Meta `tags:"后台"`
}
type ListWalletChangeTypeOptionsRes struct {
	Id       uint64 `json:"id"          description:""`
	Title    string `json:"title"       description:""`
	SubTitle string `json:"subTitle"    description:""`
	Type     uint   `json:"type"        description:"1 add; 2 reduce"`
	g.Meta   `tags:"后台"`
	Class    string `json:"class"       description:""`
	Desc     string `json:"desc"        description:""`
	Status   uint   `json:"status"      description:""`
}

// DelWalletChangeTypeReq 删除
type DelWalletChangeTypeReq struct {
	g.Meta `tags:"后台"`
	Id     uint64 `v:"required"`
}

// UpdateWalletChangeTypeReq 修改
type UpdateWalletChangeTypeReq struct {
	g.Meta `tags:"后台"`
	*entity.WalletChangeType
}

//--- Wallet ---------------------------------------------------------

// AddWalletReq 添加
type AddWalletReq struct {
	g.Meta `tags:"后台"`
	*entity.Wallet
}

// GetWalletReq 获取
type GetWalletReq struct {
	g.Meta `tags:"后台"`
	Id     uint64 `v:"required"`
}
type GetWalletRes struct {
	g.Meta `tags:"后台"`
	Data   *entity.Wallet `json:"data"`
}

// ListWalletReq 集合
type ListWalletReq struct {
	g.Meta `tags:"后台"`
	api.PageReq
	Uname        string
	Balance      float64
	Desc         string
	Status       string
	Trc20Address string
}
type ListWalletRes struct {
	g.Meta `tags:"后台"`
	List   []*model.Wallet `json:"list"`
	*api.PageRes
}

// DelWalletReq 删除
type DelWalletReq struct {
	g.Meta `tags:"后台"`
	Id     uint64 `v:"required"`
}

// UpdateWalletReq 修改
type UpdateWalletReq struct {
	g.Meta `tags:"后台"`
	*do.Wallet
}
type UpdateWalletPassReq struct {
	g.Meta `tags:"后台"`
	Id     int64  `v:"required"`
	Pass   string `v:"required"`
}
type UpdateWalletByAdminReq struct {
	g.Meta `tags:"后台"`
	Uid    uint64
	Money  float64
	Type   int
	Desc   string
}

//--- WalletChangeLog ---------------------------------------------------------

// AddWalletChangeLogReq 添加
type AddWalletChangeLogReq struct {
	g.Meta `tags:"后台"`
	*do.WalletChangeLog
}

// GetWalletChangeLogReq 获取
type GetWalletChangeLogReq struct {
	g.Meta `tags:"后台"`
	Id     uint64 `v:"required"`
}
type GetWalletChangeLogRes struct {
	g.Meta `tags:"后台"`
	Data   *entity.WalletChangeLog `json:"data"`
}

// ListWalletChangeLogReq 集合
type ListWalletChangeLogReq struct {
	g.Meta `tags:"后台"`
	api.PageReq
	TransId string
	Uname   string
	Type    string
	Desc    string
}
type ListWalletChangeLogRes struct {
	g.Meta `tags:"后台"`
	List   []*model.WalletChangeLog `json:"list"`
	*api.PageRes
}

// DelWalletChangeLogReq 删除
type DelWalletChangeLogReq struct {
	g.Meta `tags:"后台"`
	Id     uint64 `v:"required"`
}

// UpdateWalletChangeLogReq 修改
type UpdateWalletChangeLogReq struct {
	g.Meta `tags:"后台"`
	*do.WalletChangeLog
}

//--- WalletStatisticsLog ---------------------------------------------------------

// AddWalletStatisticsLogReq 添加
type AddWalletStatisticsLogReq struct {
	g.Meta `tags:"后台"`
	*do.WalletStatisticsLog
}

// GetWalletStatisticsLogReq 获取
type GetWalletStatisticsLogReq struct {
	g.Meta `tags:"后台"`
	Id     uint64 `v:"required"`
}
type GetWalletStatisticsLogRes struct {
	g.Meta `tags:"后台"`
	Data   *entity.WalletStatisticsLog `json:"data"`
}

type GetReportReq struct {
	g.Meta            `tags:"后台"`
	Uname, Begin, End string
}

type GetReportRes struct {
	g.Meta `tags:"后台"`
	T1     float64 `json:"t1"          description:""`
	T2     float64 `json:"t2"          description:""`
	T3     float64 `json:"t3"          description:""`
	T4     float64 `json:"t4"          description:""`
	T5     float64 `json:"t5"          description:""`
	T6     float64 `json:"t6"          description:""`
	T7     float64 `json:"t7"          description:""`
	T8     float64 `json:"t8"          description:""`
	T9     float64 `json:"t9"          description:""`
	T10    float64 `json:"t10"         description:""`
	T11    float64 `json:"t11"         description:""`
	T12    float64 `json:"t12"         description:""`
	T13    float64 `json:"t13"         description:""`
}

// ListWalletStatisticsLogReq 集合
type ListWalletStatisticsLogReq struct {
	g.Meta `tags:"后台"`
	api.PageReq
	Uname string
	Begin string
	End   string
}
type ListWalletStatisticsLogRes struct {
	g.Meta `tags:"后台"`
	List   []*model.WalletStatisticsLog `json:"list"`
	*api.PageRes
}

// DelWalletStatisticsLogReq 删除
type DelWalletStatisticsLogReq struct {
	g.Meta `tags:"后台"`
	Id     uint64 `v:"required"`
}

// UpdateWalletStatisticsLogReq 修改
type UpdateWalletStatisticsLogReq struct {
	g.Meta `tags:"后台"`
	*do.WalletStatisticsLog
}

type UploadFileReq struct {
	g.Meta `tags:"系统" dc:"上传图片"`
	Group  int `json:"group" dc:"分组:1头像,2图片,3动图,4音频,5文件" d:"2" in:"query"`
}
type UploadFilesRes struct {
	DbNames   []string `json:"db_names" dc:"图片"`
	ImgPrefix string   `json:"img_prefix" dc:"图片前缀"`
}

type DictReq struct {
	g.Meta `tags:"系统" summary:"字典接口" dc:"recharge_range:充值金额范围,trc_address:trc充值订单,customer_link:客户链接,domain_name:域名" `
	Key    string `json:"key" in:"query"`
}
type DictRes struct {
	Value string `json:"value" dc:"值"`
}
