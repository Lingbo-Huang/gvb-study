package global

import (
	"gorm.io/gorm"
	"gvb-server/config"
)

var (
	Config *config.Config
	DB     *gorm.DB
)
