package sys

import (
	"ciel-admin/internal/dao"
	"ciel-admin/internal/model/entity"
	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	"path"
	"time"
)

func UploadFile(ctx context.Context, r *ghttp.Request) error {
	files := r.GetUploadFiles("file")
	if len(files) == 0 {
		return errors.New("file can't be empty")
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
		_, err = dao.File.Ctx(ctx).Insert(entity.File{
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
func GetFileById(ctx context.Context, id interface{}) (*entity.File, error) {
	return dao.File.GetById(ctx, id)
}
func RemoveFile(ctx context.Context, path string) error {
	if !gfile.Exists(path) {
		g.Log().Warningf(ctx, "path:%v is not exists", path)
		return nil
	}
	if !gfile.IsFile(path) {
		g.Log().Warningf(ctx, "path:%v is not file", path)
		return nil
	}
	if err := gfile.Remove(path); err != nil {
		g.Log().Errorf(ctx, "remove File error path is %v,err:%v", path, err.Error())
		return fmt.Errorf("remove file error path is %v", path)
	}
	g.Log().Debugf(ctx, "Remove File success path is %v", path)
	return nil
}
