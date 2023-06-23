package main

import (
	_ "clickhouse-demo/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"
	"clickhouse-demo/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
