package controller

import (
	"context"
	"errors"
	"freekey-backend/api"
	"freekey-backend/api/v1"
	"freekey-backend/internal/consts"
	"freekey-backend/internal/model"
	"freekey-backend/internal/service"
	"freekey-backend/utility/utils/res"
	"freekey-backend/utility/utils/xcaptcha"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	captcha "github.com/mojocn/base64Captcha"
)

var Sys = cSys{}

type cSys struct{}

// ----------------Ws-----------------------

func (c cSys) WsGetConnectForAdmin(r *ghttp.Request) {
	service.Ws.GetAdminWs(r)
}
func (c cSys) WsSendMsg(r *ghttp.Request) {
	var (
		ctx = r.Context()
		d   model.AdminMsg
	)
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	if err := service.Sys.WsSendAdminMsg(ctx, &d); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
func (c cSys) WsNoticeAdmins(r *ghttp.Request) {
	var (
		ctx = r.Context()
		d   model.AdminMsg
	)
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	if err := service.Sys.WsNoticeAdmins(ctx, &d); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}

// ----------------Menu-----------------------

func (c cSys) GetMenuById(ctx context.Context, req *v1.GetMenuReq) (res *v1.GetMenuRes, err error) {
	data, err := service.Sys.GetMenuById(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &v1.GetMenuRes{Data: data}, nil
}
func (c cSys) GetMenuByPath(ctx context.Context, req *v1.GetMenuReqByPathReq) (res *v1.GetMenuRes, err error) {
	menu, err := service.Sys.GetMenuBypath(ctx, req.Path)
	if err != nil {
		return nil, err
	}
	return &v1.GetMenuRes{Data: menu}, nil
}
func (c cSys) ListMenu(ctx context.Context, req *v1.ListMenuReq) (res *v1.ListMenuRes, err error) {
	Menu, pageRes, err := service.Sys.ListMenu(ctx, req)
	return &v1.ListMenuRes{List: Menu, PageRes: pageRes}, nil
}
func (c cSys) AddMenu(ctx context.Context, req *v1.AddMenuReq) (res *api.DefaultRes, err error) {
	if err = service.Sys.AddMenu(ctx, req.Menu); err != nil {
		return nil, err
	}
	return
}
func (c cSys) DelMenu(ctx context.Context, req *v1.DelMenuReq) (res *api.DefaultRes, err error) {
	if err = service.Sys.DelMenu(ctx, req.Id); err != nil {
		return nil, err
	}
	return
}
func (c cSys) UpdateMenu(ctx context.Context, req *v1.UpdateMenuReq) (res *api.DefaultRes, err error) {
	if err = service.Sys.UpdateMenu(ctx, req); err != nil {
		return nil, err
	}
	return
}
func (c cSys) SortMenu(ctx context.Context, req *v1.MenuSortReq) (res *api.DefaultRes, err error) {
	if err = service.Sys.UpdateMenuSort(ctx, req.Sort, req.Id); err != nil {
		return nil, err
	}
	return
}

// ----------------API-----------------------

func (c cSys) GetAPIById(ctx context.Context, req *v1.GetApiReq) (res *v1.GetApiRes, err error) {
	data, err := service.Sys.GetApiById(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &v1.GetApiRes{Data: data}, nil
}
func (c cSys) ListAPI(ctx context.Context, req *v1.ListApiReq) (res *v1.ListApiRes, err error) {
	Api, pageRes, err := service.Sys.ListApi(ctx, req)
	return &v1.ListApiRes{List: Api, PageRes: pageRes}, nil
}
func (c cSys) AddAPI(ctx context.Context, req *v1.AddApiReq) (res *api.DefaultRes, err error) {
	if err = service.Sys.AddApi(ctx, req.Api); err != nil {
		return nil, err
	}
	return
}
func (c cSys) AddAPIGroup(ctx context.Context, req *v1.AddApiGroupReq) (res *v1.AddApiGroupRes, err error) {
	count, err := service.Sys.AddGroupApi(ctx, req.Group, req.Url)
	if err != nil {
		return nil, err
	}
	return &v1.AddApiGroupRes{Count: count}, nil
}
func (c cSys) DelAPI(ctx context.Context, req *v1.DelApiReq) (res *api.DefaultRes, err error) {
	if err = service.Sys.DelApi(ctx, req.Id); err != nil {
		return nil, err
	}
	return
}
func (c cSys) UpdateAPI(ctx context.Context, req *v1.UpdateApiReq) (res *api.DefaultRes, err error) {
	if err = service.Sys.UpdateApi(ctx, req); err != nil {
		return nil, err
	}
	return
}

// ----------------Role-----------------------

func (c cSys) AddRole(ctx context.Context, menu *v1.AddRoleReq) (_ *api.DefaultRes, _ error) {
	if err := service.Sys.AddRole(ctx, menu.Data); err != nil {
		return nil, err
	}
	return
}
func (c cSys) GetRoleById(ctx context.Context, req *v1.GetRoleReq) (*v1.GetRoleRes, error) {
	role, err := service.Sys.GetRoleById(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &v1.GetRoleRes{Data: role}, nil
}
func (c cSys) GetOptions(ctx context.Context, _ *api.DefaultReq) (*v1.GetRoleOptionsRes, error) {
	options, err := service.Sys.GetRoleOptions(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.GetRoleOptionsRes{Options: options}, err
}
func (c cSys) ListRole(ctx context.Context, req *v1.ListRoleReq) (res *v1.ListRoleRes, err error) {
	pageRes, data, err := service.Sys.ListRole(ctx, req)
	if err != nil {
		return nil, err
	}
	return &v1.ListRoleRes{PageRes: pageRes, List: data}, nil
}
func (c cSys) DelRole(ctx context.Context, req *v1.DelRoleReq) (_ *api.DefaultRes, _ error) {
	err := service.Sys.DelRole(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return
}
func (c cSys) DelRoleMenus(ctx context.Context, in *v1.DelRoleMenuReq) (_ *api.DefaultRes, _ error) {
	if err := service.Sys.DelRoleMenus(ctx, in.Id); err != nil {
		return nil, err
	}
	return
}
func (c cSys) UpdateRole(ctx context.Context, req *v1.UpdateRoleReq) (_ *api.DefaultRes, _ error) {
	if err := service.Sys.UpdateRole(ctx, req); err != nil {
		return nil, err
	}
	return
}

// ----------------RoleMenu-----------------------

func (c cSys) ListRoleMenu(ctx context.Context, req *v1.ListRoleMenuReq) (res *v1.ListRoleMenuRes, err error) {
	RoleMenu, pageRes, err := service.Sys.ListRoleMenu(ctx, req)
	return &v1.ListRoleMenuRes{List: RoleMenu, PageRes: pageRes}, nil
}
func (c cSys) ListRoleNoMenus(ctx context.Context, req *v1.ListRoleNoMenusReq) (res []*v1.ListRoleNoMenusRes, err error) {
	return service.Sys.ListRoleNoMenus(ctx, req.Rid)
}
func (c cSys) DelRoleMenu(ctx context.Context, req *v1.DelRoleMenuReq) (res *api.DefaultRes, err error) {
	if err = service.Sys.DelRoleMenu(ctx, req.Id); err != nil {
		return nil, err
	}
	return
}
func (c cSys) ClearRoleMenu(ctx context.Context, req *v1.DelClearRoleMenuReq) (res *api.DefaultRes, err error) {
	if err = service.Sys.ClearRoleMenu(ctx, req.Rid); err != nil {
		return nil, err
	}
	return
}
func (c cSys) AddRoleMenus(ctx context.Context, req *v1.AddRoleMenusReq) (res *api.DefaultRes, err error) {
	if err = service.Sys.AddMenus(ctx, req.Rid, req.Mids); err != nil {
		return nil, err
	}
	return
}

// ----------------RoleApi-----------------------

func (c cSys) ListRoleApi(ctx context.Context, req *v1.ListRoleApiReq) (res *v1.ListRoleApiRes, err error) {
	RoleApi, pageRes, err := service.Sys.ListRoleApi(ctx, req)
	return &v1.ListRoleApiRes{List: RoleApi, PageRes: pageRes}, nil
}
func (c cSys) ListRoleNoApis(ctx context.Context, req *v1.ListRoleNoApisReq) (res []*v1.ListRoleNoApisRes, err error) {
	return service.Sys.ListRoleNoApis(ctx, req.Rid)
}
func (c cSys) AddRoleApis(ctx context.Context, req *v1.AddRoleApisReq) (res *api.DefaultRes, err error) {
	if err = service.Sys.AddRoleApis(ctx, req.Rid, req.Apis); err != nil {
		return nil, err
	}
	return
}
func (c cSys) DelRoleApi(ctx context.Context, req *v1.DelRoleApiReq) (res *api.DefaultRes, err error) {
	if err = service.Sys.DelRoleApi(ctx, req.Id); err != nil {
		return nil, err
	}
	return
}
func (c cSys) ClearRoleApi(ctx context.Context, req *v1.DelRoleApiClearReq) (res *api.DefaultRes, err error) {
	if err = service.Sys.ClearRoleApi(ctx, req.Rid); err != nil {
		return nil, err
	}
	return
}

// ----------------Admin-----------------------

func (c cSys) AddAdmin(ctx context.Context, req *v1.AddAdminReq) (res *api.DefaultRes, err error) {
	if req.Pwd == "" {
		return nil, consts.ErrPassEmpty
	}
	if req.Nickname == "" {
		req.Nickname = req.Uname
	}
	if req.Email != "" {
		if err := g.Validator().Rules("email").Data(req.Email).Run(ctx); err != nil {
			return nil, consts.ErrFormatEmail
		}
	}
	if err = service.Sys.AddAdmin(ctx, req.Admin); err != nil {
		return nil, err
	}
	return
}
func (c cSys) AdminLogin(ctx context.Context, req *v1.AdminLoginReq) (res *v1.AdminLoginRes, err error) {
	token, err := service.Sys.AdminLogin(ctx, req)
	if err != nil {
		return nil, err
	}
	return &v1.AdminLoginRes{Token: token}, nil
}
func (c cSys) GetCaptcha(_ context.Context, req *v1.CaptchaReq) (res *v1.CaptchaRes, err error) {
	var driver = xcaptcha.NewDriver().ConvertFonts()
	cc := captcha.NewCaptcha(driver, xcaptcha.Store)
	_, content, answer := cc.Driver.GenerateIdQuestionAnswer()
	item, _ := cc.Driver.DrawCaptcha(content)
	_ = cc.Store.Set(req.Id, answer)
	return &v1.CaptchaRes{Img: item.EncodeB64string()}, nil
}
func (c cSys) GetAdminInfo(ctx context.Context, _ *v1.AdminInfoReq) (res *v1.AdminInfoRes, err error) {
	info, err := service.Sys.GetAdminInfo(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.AdminInfoRes{Info: info, Menus: info.Menus}, nil
}
func (c cSys) GetAdminById(ctx context.Context, req *v1.GetAdminReq) (res *v1.GetAdminRes, err error) {
	data, err := service.Sys.GetAdminById(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &v1.GetAdminRes{Data: data}, nil
}
func (c cSys) ListAdmin(ctx context.Context, req *v1.ListAdminReq) (res *v1.ListAdminRes, err error) {
	Admin, pageRes, err := service.Sys.ListAdmin(ctx, req)
	return &v1.ListAdminRes{List: Admin, PageRes: pageRes}, nil
}
func (c cSys) DelAdmin(ctx context.Context, req *v1.DelAdminReq) (res *api.DefaultRes, err error) {
	if err = service.Sys.DelAdmin(ctx, req.Id); err != nil {
		return nil, err
	}
	return
}
func (c cSys) UpdateAdmin(ctx context.Context, req *v1.UpdateAdminReq) (res *api.DefaultRes, err error) {
	if err = service.Sys.UpdateAdmin(ctx, req); err != nil {
		return nil, err
	}
	return
}
func (c cSys) UpdateAdminUname(ctx context.Context, req *v1.UpdateAdminUnameReq) (res *api.DefaultRes, err error) {
	if err = service.Sys.UpdateAdminUname(ctx, req.Id, req.Uname); err != nil {
		return nil, err
	}
	return
}
func (c cSys) UpdateAdminPass(ctx context.Context, req *v1.UpdateAdminPassReq) (res *api.DefaultRes, err error) {
	if err = service.Sys.UpdateAdminPass(ctx, req.Id, req.Pass); err != nil {
		return nil, err
	}
	return
}
func (c cSys) UpdateAdminPassBySelf(ctx context.Context, req *v1.UpdateAdminPassSelfReq) (res *api.DefaultRes, err error) {
	if err = service.Sys.UpdateAdminPassSelf(ctx, req.Pass); err != nil {
		return nil, err
	}
	return
}

// ----------------Dict-----------------------

func (c cSys) GetDictById(ctx context.Context, req *v1.GetDictReq) (res *v1.GetDictRes, err error) {
	data, err := service.Sys.GetDictById(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &v1.GetDictRes{Data: data}, nil
}
func (c cSys) GetDictByKey(ctx context.Context, req *v1.DictReq) (res *v1.DictRes, err error) {
	data, err := service.Sys.GetDictByKey(ctx, req.Key)
	if err != nil {
		return nil, err
	}
	return &v1.DictRes{Value: data}, nil
}
func (c cSys) ListDict(ctx context.Context, req *v1.ListDictReq) (res *v1.ListDictRes, err error) {
	Dict, pageRes, err := service.Sys.ListDict(ctx, req)
	return &v1.ListDictRes{List: Dict, PageRes: pageRes}, nil
}
func (c cSys) AddDict(ctx context.Context, req *v1.AddDictReq) (res *api.DefaultRes, err error) {
	if err = service.Sys.AddDict(ctx, req.Dict); err != nil {
		return nil, err
	}
	return
}
func (c cSys) DelDict(ctx context.Context, req *v1.DelDictReq) (res *api.DefaultRes, err error) {
	if err = service.Sys.DelDict(ctx, req.Id); err != nil {
		return nil, err
	}
	return
}
func (c cSys) UpdateDict(ctx context.Context, req *v1.UpdateDictReq) (res *api.DefaultRes, err error) {
	if err = service.Sys.UpdateDict(ctx, req); err != nil {
		return nil, err
	}
	return
}

// ----------------LoginLog-----------------------

func (c cSys) ListLoginLog(ctx context.Context, req *v1.ListAdminLoginLogReq) (res *v1.ListAdminLoginLogRes, err error) {
	AdminLoginLog, pageRes, err := service.Sys.ListLoginLog(ctx, req)
	return &v1.ListAdminLoginLogRes{List: AdminLoginLog, PageRes: pageRes}, nil
}
func (c cSys) DelLoginLog(ctx context.Context, req *v1.DelAdminLoginLogReq) (res *api.DefaultRes, err error) {
	if err = service.Sys.DelLoginLog(ctx, req.Id); err != nil {
		return nil, err
	}
	return
}
func (c cSys) ClearLoginLog(ctx context.Context, _ *v1.DelClearAdminLoginLogReq) (res *api.DefaultRes, err error) {
	if err = service.Sys.ClearLoginLog(ctx); err != nil {
		return nil, err
	}
	return
}

// ----------------OperationLog-----------------------

func (c cSys) ListOperationLog(ctx context.Context, req *v1.ListOperationLogReq) (res *v1.ListOperationLogRes, err error) {
	OperationLog, pageRes, err := service.Sys.ListOperationLog(ctx, req)
	return &v1.ListOperationLogRes{List: OperationLog, PageRes: pageRes}, nil
}
func (c cSys) DelOperationLog(ctx context.Context, req *v1.DelOperationLogReq) (res *api.DefaultRes, err error) {
	if err = service.Sys.DelOperationLog(ctx, req.Id); err != nil {
		return nil, err
	}
	return
}
func (c cSys) ClearOperationLog(ctx context.Context, _ *v1.DelClearOperationLogReq) (res *api.DefaultRes, err error) {
	if err = service.Sys.ClearOperationLog(ctx); err != nil {
		return nil, err
	}
	return
}

// ----------------File-----------------------

func (c cSys) GetFileById(ctx context.Context, req *v1.GetFileReq) (res *v1.GetFileRes, err error) {
	data, err := service.Sys.GetFileById(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &v1.GetFileRes{Data: data}, nil
}
func (c cSys) ListFile(ctx context.Context, req *v1.ListFileReq) (res *v1.ListFileRes, err error) {
	File, pageRes, err := service.Sys.ListFile(ctx, req)
	return &v1.ListFileRes{List: File, PageRes: pageRes}, nil
}
func (c cSys) DelFile(ctx context.Context, req *v1.DelFileReq) (res *api.DefaultRes, err error) {
	if err = service.Sys.DelFile(ctx, req.Id); err != nil {
		return nil, err
	}
	return
}
func (c cSys) UpdateFile(ctx context.Context, req *v1.UpdateFileReq) (res *api.DefaultRes, err error) {
	if err = service.Sys.UpdateFile(ctx, req); err != nil {
		return nil, err
	}
	return
}
func (c cSys) UploadFiles(ctx context.Context, in *v1.UploadFileReq) (res *v1.UploadFilesRes, err error) {
	files := ghttp.RequestFromCtx(ctx).GetUploadFiles("file")
	if len(files) == 0 {
		return nil, errors.New("file can't be empty")
	}
	group := ghttp.RequestFromCtx(ctx).GetQuery("group")
	if group.IsEmpty() {
		in.Group = 1
	} else {
		in.Group = group.Int()
	}
	uploadsFiles, imgPrefix, err := service.Sys.UploadsFiles(ctx, files, in.Group)
	if err != nil {
		return nil, err
	}
	return &v1.UploadFilesRes{DbNames: uploadsFiles, ImgPrefix: imgPrefix}, nil
}
