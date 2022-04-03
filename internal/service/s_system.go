package service

import (
	"ciel-admin/internal/consts"
	"ciel-admin/internal/model/bo"
	"ciel-admin/internal/model/entity"
	"ciel-admin/internal/service/internal/dao"
	"ciel-admin/manifest/config"
	"ciel-admin/utility/utils/res"
	"ciel-admin/utility/utils/xjwt"
	"ciel-admin/utility/utils/xpwd"
	"ciel-admin/utility/utils/xredis"
	"ciel-admin/utility/utils/xstr"
	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gxml"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gview"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	"net/http"
	"net/url"
	"path"
	"sort"
	"strings"
	"time"
)

const (
	AdminSessionKey = "adminInfo"
	Uid             = "userInfoKey"
)

type (
	system     struct{}
	role       struct{}
	session    struct{}
	view       struct{}
	rss        struct{}
	file       struct{}
	admin      struct{}
	middleware struct{}
	gen        struct{ ModName string }
)

var (
	sSystem     = &system{}
	insRole     = &role{}
	sRss        = &rss{}
	sView       = &view{}
	sFile       = &file{}
	sAdmin      = &admin{}
	sMiddleware = &middleware{}
	sGen        = newGen()
)

// ---System --------------------------------------------------------------------------------------------------

func System() *system { return sSystem }
func (s *system) List(ctx context.Context, c *config.SearchConf) (count int, data gdb.List, err error) {
	db := g.DB().Ctx(ctx).Model(c.T1 + " t1")
	if c.T2 != "" {
		db = db.LeftJoin(c.T2)
	}
	if c.T3 != "" {
		db = db.LeftJoin(c.T3)
	}
	if c.T4 != "" {
		db = db.LeftJoin(c.T4)
	}
	if c.T5 != "" {
		db = db.LeftJoin(c.T5)
	}

	conditions := c.FilterConditions(ctx)
	if len(conditions) > 0 {
		for _, item := range conditions {
			field := item.Field
			if g.IsEmpty(item.Value) {
				continue
			}
			if !strings.Contains(field, ".") {
				field = "t1." + field
			}
			if item.Like {
				db = db.WhereLike(field, xstr.Like(gconv.String(item.Value)))
			} else {
				db = db.Where(field, item.Value)
			}
		}
	}

	if count, err = db.Count(); err != nil {
		return
	}
	var o = "t1.id desc"
	if c.OrderBy != "" {
		o = c.OrderBy
	}
	if c.SearchFields == "" {
		c.SearchFields = "t1.*"
	}
	all, err := db.Page(c.Page, c.Size).Fields(c.SearchFields).Order(o).All()
	if all.IsEmpty() {
		return
	}
	data = all.List()
	return
}
func (s *system) Add(ctx context.Context, table, data interface{}) error {
	if _, err := g.DB().Ctx(ctx).Model(table).Insert(data); err != nil {
		glog.Error(ctx, err)
		return err
	}
	return nil
}
func (s *system) Del(ctx context.Context, table, id interface{}) (err error) {
	if _, err = g.DB().Ctx(ctx).Model(table).Delete("id", id); err != nil {
		glog.Error(ctx, err)
		return
	}
	return
}
func (s *system) Update(ctx context.Context, table string, id, data interface{}) error {
	// 空值过滤
	_, err := g.DB().Model(table).Where("id", id).Data(data).Update()
	if err != nil {
		glog.Error(ctx, err)
		return err
	}
	return nil
}
func (s *system) GetById(ctx context.Context, table, id interface{}) (gdb.Record, error) {
	one, err := g.DB().Ctx(ctx).Model(table).One("id", id)
	if err != nil {
		glog.Error(ctx, err)
		return nil, err
	}
	return one, nil
}
func (s *system) Icon(ctx context.Context, path string) (string, error) {
	menu, err := dao.Menu.GetByPath(ctx, path)
	if err != nil {
		return "", nil
	}
	if menu.Icon == "" {
		return "", nil
	}
	return consts.ImgPrefix + menu.Icon, err
}
func (s *system) Init() {
	get, err := g.Cfg().Get(gctx.New(), "server.imgPrefix")
	if err != nil {
		panic(err)
	}
	consts.ImgPrefix = get.String()
}

//  ---Role --------------------------------------------------------------------------------------------------

func Role() *role { return insRole }
func (s *role) RoleNoMenu(ctx context.Context, rid interface{}) (interface{}, error) {
	return dao.RoleMenu.RoleNoMenu(ctx, rid)
}
func (s *role) AddRoleMenu(ctx context.Context, rid int, mid []int) error {
	return dao.RoleMenu.AddRoleMenu(ctx, rid, mid)
}
func (s *role) RoleNoApi(ctx context.Context, rid interface{}) (gdb.List, error) {
	return dao.RoleApi.RoleNoApi(ctx, rid)
}
func (s *role) AddRoleApi(ctx context.Context, rid int, aid []int) error {
	return dao.RoleApi.AddRoleApi(ctx, rid, aid)
}
func (s *role) CheckRoleApi(ctx context.Context, rid int, uri string, method string) bool {
	if strings.Contains(uri, "?") {
		uri = strings.Split(uri, "?")[0]
	}
	if uri == "/" {
		return true
	}
	count, _ := g.DB().Ctx(ctx).Model("s_role t1").
		LeftJoin("s_role_api t2 on t1.id = t2.rid").
		LeftJoin("s_api t3 on t2.aid = t3.id").
		Where("t3.url = ? and t3.method = ? and t1.id = ?  ", uri, method, rid).
		Count()
	if count == 1 {
		return false
	}
	return true
}
func (s *role) Menus(ctx context.Context, rid int, pid int) ([]*bo.Menu, error) {
	var d = make([]*bo.Menu, 0)
	get, err := g.Cfg().Get(ctx, "rss")
	if err != nil {
		return nil, err
	}
	array := get.Array()
	if len(array) > 0 {
		children := make([]*bo.Menu, 0)
		d = append(d, &bo.Menu{
			Name: "首页",
			Children: []*bo.Menu{
				{Name: "V2EX", Path: "/"},
				{Name: "Github", Path: "/sys/path/github"},
				{Name: "豆瓣阅读", Path: "/sys/path/douban"},
				{Name: "开源中国", Path: "/sys/path/oschina"},
			},
		})
		for _, item := range array {
			split := gstr.Split(fmt.Sprint(item), ":")
			children = append(children, &bo.Menu{
				Name: split[0],
				Path: split[1],
			})
		}
	}
	menus, err := dao.RoleMenu.Menus(ctx, rid, pid)
	if err != nil {
		return nil, err
	}
	d = append(d, menus...)
	return d, err
}

func (s *role) Roles(ctx context.Context) (gdb.Result, error) {
	all, err := dao.Role.Ctx(ctx).All()
	if err != nil {
		return nil, err
	}
	return all, nil
}

// ---session ------------------------------------------------------------

func Session() *session { return &session{} }
func (s session) SetAdmin(ctx context.Context, data *bo.Admin) error {
	return g.RequestFromCtx(ctx).Session.Set(AdminSessionKey, data)
}
func (s session) GetAdmin(r *ghttp.Request) (*bo.Admin, error) {
	get, err := r.Session.Get(AdminSessionKey)
	var data *bo.Admin
	err = get.Scan(&data)
	if err != nil {
		return nil, err
	}
	return data, err
}
func (s session) RemoveAdmin(ctx context.Context) error {
	return g.RequestFromCtx(ctx).Session.Remove(AdminSessionKey)
}
func (s session) AdminIsLogin(r *ghttp.Request) error {
	user, err := s.GetAdmin(r)
	if err != nil {
		return err
	}
	if user == nil {
		return consts.ErrNotAuth
	}
	return nil
}

// ---view-------------------------------------------------------------------------------

func View() *view { return sView }
func (s *view) QueryEscape(str string) string {
	unescape, err := url.QueryUnescape(str)
	if err != nil {
		panic(err)
	}
	return unescape
}
func (s *view) BindFuncMap() gview.FuncMap {
	return gview.FuncMap{}
}

// ---Rss-------------------------------------------------------------------------------

func Rss() *rss { return sRss }
func (s *rss) fetchXml(ctx context.Context, url string) (map[string]interface{}, error) {
	num := 0
	max := 5
retry:
	data, err := gclient.New().Timeout(time.Second*3).Get(ctx, url)
	if err != nil {
		num++
		if num > max {
			return nil, errors.New(fmt.Sprintf("获取RSS数据失败,已重试%d次,请稍后重试", max))
		}
		glog.Infof(ctx, "获取RSS失败,重试中...%d", num)
		goto retry
	}
	return gxml.DecodeWithoutRoot([]byte(data.ReadAllString()))
}
func (s *rss) Feftch(ctx context.Context, url string) (map[string]interface{}, error) {
	return s.fetchXml(ctx, url)
}

// ---File-------------------------------------------------------------------------------

func File() *file { return sFile }
func (f file) Upload(ctx context.Context, r *ghttp.Request) error {
	files := r.GetUploadFiles("file")
	if len(files) == 0 {
		res.Err(errors.New("file can't be empty"), r)
	}
	for _, file := range files {
		fileName := fmt.Sprint(grand.S(6), path.Ext(file.Filename))
		file.Filename = fileName
	}
	datePre := time.Now().Format("2006/01")
	group := r.Get("group").String()
	if group == "" || group == "undefined" {
		group = "1"
	}
	rootFilePath, err := g.Cfg().Get(ctx, "server.rootFilePath")
	if err != nil {
		return err
	}
	rootPath := gfile.Pwd() + rootFilePath.String()
	mixPath := fmt.Sprintf("%s/%s/%s/", rootPath, group, datePre)
	_, err = files.Save(mixPath)
	if err != nil {
		return err
	}
	for _, file := range files {
		dbName := fmt.Sprintf("%s/%s/%s", group, datePre, file.Filename)
		_, err := dao.File.Ctx(ctx).Insert(entity.File{
			Url:    dbName,
			Group:  gconv.Int(group),
			Status: 1,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
func (f file) GetById(ctx context.Context, id uint64) (*entity.File, error) {
	return dao.File.GetById(ctx, id)
}

// ---Admin-------------------------------------------------------------------------------

func Admin() *admin { return sAdmin }
func (s *admin) Login(ctx context.Context, uname string, pwd string) error {
	admin, err := dao.Admin.GetByUname(ctx, uname)
	if err != nil {
		return err
	}
	if !xpwd.ComparePassword(admin.Pwd, pwd) {
		return consts.ErrLogin
	}

	if admin.Status == 2 {
		return consts.ErrAuthNotEnough
	}
	menus, err := Role().Menus(ctx, admin.Rid, -1)
	if err != nil {
		return err
	}
	if err = Session().SetAdmin(ctx, &bo.Admin{Admin: admin, Menus: menus}); err != nil {
		return err
	}
	return nil
}
func (s *admin) Logout(ctx context.Context) error {
	return Session().RemoveAdmin(ctx)
}
func (s *admin) UpdateAdminPwd(ctx context.Context, pwd string, pwd2 string) error {
	admin, err := Session().GetAdmin(ghttp.RequestFromCtx(ctx))
	if err != nil {
		return err
	}
	u, err := dao.Admin.GetByUname(ctx, admin.Admin.Uname)
	if err != nil {
		return err
	}
	if !xpwd.ComparePassword(u.Pwd, pwd) {
		return errors.New("old password not match")
	}
	u.Pwd = xpwd.GenPwd(pwd2)
	err = Session().RemoveAdmin(ctx)
	if err != nil {
		return err
	}
	return dao.Admin.Update(ctx, u)
}

// ---middleware-----------------------------------------------------------------------

func Middleware() *middleware { return sMiddleware }
func (s *middleware) CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}
func (s *middleware) AuthAdmin(r *ghttp.Request) {
	user, err := Session().GetAdmin(r)
	if err != nil || user == nil {
		r.Response.RedirectTo("/login")
		return
	}
	b := Role().CheckRoleApi(r.Context(), user.Admin.Rid, r.RequestURI, r.Method)
	if !b {
		res.Err(errors.New("没有权限"), r)
	}
	r.Middleware.Next()
}
func (s *middleware) UserAuth(c *ghttp.Request) {
	userInfo, err := xjwt.UserInfo(c)
	if err != nil {
		c.Response.WriteStatus(http.StatusForbidden, consts.ErrAuth.Error())
		c.Exit()
	}
	c.SetParam(Uid, userInfo.Uid)
	c.Middleware.Next()
}
func (s *middleware) LockAction(r *ghttp.Request) {
	uid := r.Get(Uid).Uint64()
	if uid == 0 {
		getAdmin, err := Session().GetAdmin(r)
		if err != nil {
			res.Err(err, r)
		}
		uid = uint64(getAdmin.Admin.Id)
		if uid == 0 {
			err := errors.New("uid is empty")
			glog.Error(nil, err)
			res.Err(err, r)
		}
	}
	lock, err := xredis.UserLock(uid)
	if err != nil {
		res.Err(err, r)
	}
	r.Middleware.Next()
	lock.Unlock()
}

// ---Gen Code-------------------------------------------------------------------------------

func newGen() *gen {
	ctx := context.Background()
	get, err := g.Cfg().Get(ctx, "server.modName")
	if err != nil {
		panic(err)
	}
	return &gen{ModName: get.String()}
}
func Gen() *gen {
	return sGen
}
func (s *gen) Fields(ctx context.Context, tableName string) (map[string]*gdb.TableField, error) {
	fields, err := g.DB().Ctx(ctx).Model(tableName).TableFields(tableName)
	if err != nil {
		return nil, err
	}
	return fields, nil
}
func (s gen) Tables(ctx context.Context) ([]string, error) {
	return g.DB().Tables(ctx)
}
func (s *gen) GenCode(ctx context.Context, b *bo.GenCodeInfo) error {
	// gen controller
	template, err := s.fileFactory(b, 0)
	if err != nil {
		glog.Error(ctx, err)
		return err
	}
	err = s.SaveFile(b.Table, template, "", 0)

	// gen cmd
	template, err = s.fileFactory(b, 1)
	if err != nil {
		return err
	}
	err = s.SaveFile(b.Table, template, "", 1)

	//// gen html
	//template, err = s.fileFactory(b, 2)
	//if err != nil {
	//	return err
	//}
	//err = s.SaveFile(b.Table, template, b.Category, 2)

	// gen api
	s.genApi(ctx, b.Category, b.StructName)
	return nil
}

// fileFactory  t 0:controller 1:cmd 2:html
func (s gen) fileFactory(b *bo.GenCodeInfo, t int) (string, error) {
	switch t {
	case 0: // controller
		return s.makeControllerStr(b)
	case 1: // cmd
		return s.makeCmdStr(b)
	case 2: // html
		return s.makeHtmlStr(b)
	}
	return "", nil

}
func (s *gen) makeControllerStr(b *bo.GenCodeInfo) (string, error) {
	filePath := fmt.Sprintf("%s/manifest/gen_code_template/controller.text", gfile.MainPkgPath())
	template := gfile.GetContents(filePath)
	// set modName
	template = strings.ReplaceAll(template, "ciel-admin", s.ModName)
	// set category
	template = strings.ReplaceAll(template, "category", gstr.CaseCamelLower(b.Category))
	// set html file name
	template = strings.ReplaceAll(template, "fileName", b.Table)
	// set structName
	template = strings.ReplaceAll(template, "menu", gstr.CaseCamelLower(b.StructName))
	template = strings.ReplaceAll(template, "Menu", gstr.CaseCamel(b.StructName))
	// set search config
	temps := ""
	for _, i := range b.Fields {
		switch i.SearchType {
		case "like":
			temps += fmt.Sprintf(`{Field: "%s", Like: true,QueryField: "%s"},`, i.Name, i.QueryField)
		case "=":
			temps += fmt.Sprintf(`{Field: "%s",QueryField: "%s"},`, i.Name, i.QueryField)
		case ">":
			temps += fmt.Sprintf(`{Field: "%s",GT: true,QueryField: "%s"},`, i.Name, i.QueryField)
		case "<":
			temps += fmt.Sprintf(`{Field: "%s",LT: true,QueryField: "%s"},`, i.Name, i.QueryField)
		case ">=":
			temps += fmt.Sprintf(`{Field: "%s",GTE: true,QueryField: "%s"},`, i.Name, i.QueryField)
		case "<=":
			temps += fmt.Sprintf(`{Field: "%s",LTE: true,QueryField: "%s"},`, i.Name, i.QueryField)
		case "in":
			temps += fmt.Sprintf(`{Field: "%s",In: true,QueryField: "%s"},`, i.Name, i.QueryField)
		}
	}
	if temps != "" {
		template = strings.ReplaceAll(template, `{Field: "id"},`, temps)
	}
	// set t1 t2 t3 t4 t5 t6
	manyTable := fmt.Sprintf(` T1:"%s",`, b.Table)
	if b.T2 != "" {
		manyTable += fmt.Sprintf(` T2:"%s",`, b.T2)
	}
	if b.T3 != "" {
		manyTable += fmt.Sprintf(` T3:"%s",`, b.T3)
	}
	if b.T4 != "" {
		manyTable += fmt.Sprintf(` T4:"%s",`, b.T4)
	}
	if b.T5 != "" {
		manyTable += fmt.Sprintf(` T5:"%s",`, b.T5)
	}
	if b.T6 != "" {
		manyTable += fmt.Sprintf(` T6:"%s",`, b.T6)
	}
	if b.OrderBy != "" {
		manyTable += fmt.Sprintf(`OrderBy: "%s",`, b.OrderBy)
	}
	if b.QueryField != "" {
		manyTable += fmt.Sprintf(`SearchFields: "%s",`, b.QueryField)
	}
	template = strings.ReplaceAll(template, "[Tables]", manyTable)
	return template, nil
}
func (s *gen) makeCmdStr(b *bo.GenCodeInfo) (string, error) {
	filePath := fmt.Sprintf("%s/manifest/gen_code_template/cmd.text", gfile.MainPkgPath())
	template := gfile.GetContents(filePath)
	// set structName
	template = strings.ReplaceAll(template, "menu", gstr.CaseCamelLower(b.StructName))
	template = strings.ReplaceAll(template, "Menu", gstr.CaseCamel(b.StructName))
	return template, nil
}
func (s *gen) makeHtmlStr(b *bo.GenCodeInfo) (string, error) {
	filePath := fmt.Sprintf("%s/manifest/gen_code_template/menu.html", gfile.MainPkgPath())
	template := gfile.GetContents(filePath)
	// set [desc]
	template = strings.ReplaceAll(template, "[desc]", b.Desc)
	// set [menu]
	template = strings.ReplaceAll(template, "[menu]", gstr.CaseCamelLower(b.StructName))
	template = strings.ReplaceAll(template, "[Menu]", gstr.CaseCamel(b.StructName))

	// set tr and td
	var (
		tr, td string
	)

	sort.Slice(b.Fields, func(i, j int) bool {
		return b.Fields[i].Sort < b.Fields[j].Sort
	})
	for _, i := range b.Fields {
		width := 80
		switch i.Name {
		case "id":
			width = 30
		case "created_at", "updated_at":
			width = 120
		default:

		}

		itemField := i.QueryField
		if itemField == "" {
			itemField = i.Name
		}
		title := i.Title
		if title == "" {
			title = itemField
		}

		tr += fmt.Sprintf("                  <th width='%d'>%s</th>\n", width, title)
		if i.Name == "status" {
			td += fmt.Sprintf("                     <td> <el-tag v-if='item.status==1'>启用</el-tag> <el-tag type='danger' v-if='item.status==2'>禁用</el-tag> </td> \n ")
		} else {

			td += fmt.Sprintf("                    <td v-text='item.%s'>\n", itemField)
		}
	}
	template = strings.ReplaceAll(template, `[<th width="30">ID</th>]`, tr)
	template = strings.ReplaceAll(template, `[<td v-text="item.id">]`, td)

	// set search
	search := ""
	for _, i := range b.Fields {
		if i.SearchType == "" {
			continue
		}
		itemField := i.QueryField
		if itemField == "" {
			itemField = i.Name
		}
		title := i.Title
		if title == "" {
			title = itemField
		}
		if i.Name == "status" {
			search += fmt.Sprintf(`
                        <label class="input">状态
                            <select name="status" v-model="search.status" @change="getList()">
                                <option value="">请选择</option>
                                <option value="1">启用</option>
                                <option value="2">禁用</option>
                            </select>
                        </label>
`)
		} else {
			search += fmt.Sprintf(" <label class='input'>%s<input type='text'  v-model='search.%s' autocomplete='off' @keyup.enter='getList()'></label> \n", title, itemField)
		}
	}
	template = strings.ReplaceAll(template, `[<label class="input">PID<input type="text" name="pid" v-model="search.pid" autocomplete="off" @keyup.enter="getList()"></label>]`, search)

	// set details
	details := ""
	for _, i := range b.Fields {
		disabled := ""
		switch i.DetailsType {
		case "no-show":
			continue
		case "disabled":
			disabled = "disabled"
		}
		itemField := i.QueryField
		if itemField == "" {
			itemField = i.Name
		}
		title := i.Title
		if title == "" {
			title = itemField
		}
		if i.Name == "status" {
			details += fmt.Sprintf(`                <li>
                    <label class="input">状态<select v-model="details.status">
                        <option label="正常" :value="1"></option>
                        <option label="关闭" :value="2"></option>
                    </select></label>
                </li>`)
		} else {
			details += fmt.Sprintf("                <li><label class='input'>%s<input v-model='details.%s' %s></label></li>\n", title, itemField, disabled)
		}
	}
	template = strings.ReplaceAll(template, `<li><label class="input">名称<input v-model="details.name"></label></li>`, details)
	return template, nil
}
func (s *gen) SaveFile(fileName string, fileStr string, category string, t int) error {
	switch t {
	case 0: // controller
		p := fmt.Sprintf("%s/internal/controller/%s.go", gfile.MainPkgPath(), fileName)
		return gfile.PutContents(p, fileStr)
	case 1: // cmd
		p := fmt.Sprintf("%s/internal/cmd/sys_router.go", gfile.MainPkgPath())
		stat, _ := gfile.Stat(p)
		if err := gfile.Truncate(p, (int)(stat.Size()-2)); err != nil {
			return err
		}
		return gfile.PutContentsAppend(p, fileStr)
	case 2: // html
		p := fmt.Sprintf("%s/resource/template/%s/%s.html", gfile.MainPkgPath(), category, fileName)
		return gfile.PutContents(p, fileStr)
	}
	return nil
}

func (s *gen) genApi(ctx context.Context, category string, name string) error {
	name = gstr.CaseCamelLower(name)
	array := []*entity.Api{
		{Url: fmt.Sprintf("/%s/del", name), Method: "DELETE", Group: category, Desc: fmt.Sprintf("删除%s", name), Status: 1},
		{Url: fmt.Sprintf("/%s/post", name), Method: "POST", Group: category, Desc: fmt.Sprintf("添加%s", name), Status: 1},
		{Url: fmt.Sprintf("/%s/put", name), Method: "PUT", Group: category, Desc: fmt.Sprintf("修改%s", name), Status: 1},
	}
	for _, i := range array {
		count, err := dao.Api.Ctx(ctx).Count("url", i.Url)
		if err != nil {
			return err
		}
		if count != 0 {
			continue
		}
		if _, err = dao.Api.Ctx(ctx).Insert(i); err != nil {
			return err
		}
	}
	return nil
}
