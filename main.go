package main

import (
	_ "processon/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"processon/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
