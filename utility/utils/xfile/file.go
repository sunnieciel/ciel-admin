package xfile

import (
	"bufio"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"os"
)

// ReadLine read line from file
func ReadLine(path string, lineNum int) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	fileScanner := bufio.NewScanner(file)
	lineCount := 1
	for fileScanner.Scan() {
		if lineCount == lineNum {
			return fileScanner.Text(), nil
		}
		lineCount++
	}
	if err = file.Close(); err != nil {
		return "", err
	}
	return "", nil
}

func Remove(ctx context.Context, path string) error {
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
