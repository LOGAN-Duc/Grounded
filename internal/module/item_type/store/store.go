package itemtypestore

import (
	"gorm.io/gorm"
)

type itemTypeStore struct {
	mysql *gorm.DB
}

func NewItemTypeStore(mysql *gorm.DB) *itemTypeStore {
	return &itemTypeStore{
		mysql: mysql,
	}
}
