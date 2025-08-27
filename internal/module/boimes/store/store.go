package boimesstore

import "gorm.io/gorm"

type boimesStore struct {
	mysql *gorm.DB
}

func NewBoimesStore(mysql *gorm.DB) *boimesStore {
	return &boimesStore{
		mysql: mysql,
	}
}
