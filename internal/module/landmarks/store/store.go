package landmarkstore

import "gorm.io/gorm"

type landmarkStore struct {
	db *gorm.DB
}

func NewLandmarkStore(db *gorm.DB) *landmarkStore {
	return &landmarkStore{db: db}
}
