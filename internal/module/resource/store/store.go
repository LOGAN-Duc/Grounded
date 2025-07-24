package resourcestore

import (
	"gorm.io/gorm"
)

type resourceStore struct {
	mysql *gorm.DB
}

func NewResourcesStore(mysql *gorm.DB) *resourceStore {
	return &resourceStore{
		mysql: mysql,
	}
}
