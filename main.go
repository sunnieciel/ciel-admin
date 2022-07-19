package main

import (
	"ciel-admin/internal/cmd"
	_ "ciel-admin/internal/packed"
	"ciel-admin/internal/service/sys"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gctx"
	_ "net/http/pprof"
)

func main() {
	cmd.Main.AddCommand(&gcmd.Command{
		Name:        "hello",
		Brief:       "这是一个欢迎对话命令",
		Description: "我在学习gf的gcmd包",
		Arguments: []gcmd.Argument{
			{Name: "name", Short: "n"},
		},
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			name := ""
			n := parser.GetOpt("n")
			if n.IsEmpty() {
				name = gcmd.Scan("What's your name?\n")
			} else {
				name = n.String()
			}
			age := gcmd.Scanf("Hello %s,how old are you?\n", name)
			fmt.Printf("> %s's age is :%s! bye!", name, age)
			return nil
		},
	})
	cmd.Main.AddCommand(sys.GenCommon)
	cmd.Main.Run(gctx.New())
}
