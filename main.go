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
