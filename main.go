package main

import (
	"ciel-admin/internal/cmd"
	_ "ciel-admin/internal/packed"
	"github.com/gogf/gf/v2/os/gctx"
	_ "net/http/pprof"
)

func main() {
	cmd.Main.Run(gctx.New())
}

/*
todo
	- c_diary 日记表
		- 属性
			- aid 管理员ID
			- title 标题
			- level 级别  普通 稀有 传承 唯一 史诗 传说
			- content 内容
			- sign1
			- sign2
			- sign3
			- sign4
			- sign5
			- sign6
			- sign7
*/
