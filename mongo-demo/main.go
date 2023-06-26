package main

import (
	_ "mongo-demo/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"
	"mongo-demo/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
