package main

import (
	_ "ciel-begin/internal/packed"

	"ciel-begin/internal/cmd"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	cmd.Main.Run(gctx.New())
}
