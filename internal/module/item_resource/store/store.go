package itemresourcestore

import "gorm.io/gorm"

type itemResourceStore struct {
	mysql *gorm.DB
}

func NewItemResourceStore(mysql *gorm.DB) *itemResourceStore {
	return &itemResourceStore{
		mysql: mysql,
	}
}
