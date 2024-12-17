package main

import (
	"gvb-server/core"
	"gvb-server/global"
)

func main() {
	// 初始化配置文件
	core.InitConf()
	// 初始化数据库连接
	global.DB = core.InitGorm()
}
