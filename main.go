package main

import (
	"fmt"
	"gvb-server/core"
	"gvb-server/global"
)

func main() {
	core.InitConf()
	fmt.Println(global.Config)
}
