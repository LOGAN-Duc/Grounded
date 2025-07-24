package resourcetypestore

import (
	"gorm.io/gorm"
)

type resourcesTypeStore struct {
	mysql *gorm.DB
}

func NewResourceTypeStore(mysql *gorm.DB) *resourcesTypeStore {
	return &resourcesTypeStore{
		mysql: mysql,
	}
}
