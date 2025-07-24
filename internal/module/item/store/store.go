package itemstore

import (
	"gorm.io/gorm"
)

type itemStore struct {
	mysql *gorm.DB
}

func NewItemStore(mysql *gorm.DB) *itemStore {
	return &itemStore{
		mysql: mysql,
	}
}
