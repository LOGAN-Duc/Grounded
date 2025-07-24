package cpnmysqldb

import (
	"fmt"

	"example.com/m/internal/common"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/opentelemetry/tracing"
)

type SQLGORM struct {
	DB *gorm.DB
}

func (s *SQLGORM) Connect(cfg *common.Config) (*gorm.DB, error) {
	mysqlConfig := cfg.Mysqldb
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=True&loc=Local", mysqlConfig.User, mysqlConfig.Password, mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		common.Log(err, "connect mysql error")
	}

	sqlDB, err := db.DB()
	if err != nil {
		common.Log(err, "get sql.DB error")
		return nil, err
	}

	if err := sqlDB.Ping(); err != nil {
		common.Log(err, "ping mysql error")
		return nil, err
	}

	// Set connection pool settings
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	if err := db.Use(tracing.NewPlugin()); err != nil {
		common.Log(err, "error when use tracing plugin")
	}

	return db, err
}
