package landmarkitemstore

import "gorm.io/gorm"

type landmarkItemStore struct {
	db *gorm.DB
}

func NewLandmarkItemStore(db *gorm.DB) *landmarkItemStore {
	return &landmarkItemStore{db: db}
}
