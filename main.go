package main

import (
	_ "check/internal/packed"

	_ "check/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"check/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
