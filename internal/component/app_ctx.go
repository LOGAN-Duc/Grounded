package component

import (
	"example.com/m/internal/common"
	"gorm.io/gorm"
)

type appCtx struct {
	config  *common.Config
	mysqldb *gorm.DB
}

type AppContext interface {
	GetConfig() *common.Config
	GetMySqlDB() *gorm.DB
}

func NewAppContext(
	config *common.Config,
	mysqldb *gorm.DB,
) AppContext {
	return &appCtx{
		config:  config,
		mysqldb: mysqldb,
	}
}

func (ctx *appCtx) GetConfig() *common.Config {
	return ctx.config
}

func (ctx *appCtx) GetMySqlDB() *gorm.DB {
	return ctx.mysqldb
}
