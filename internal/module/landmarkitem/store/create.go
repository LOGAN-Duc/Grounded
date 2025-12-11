package landmarkitemstore

import (
	"context"

	landmarkitemmodel "example.com/m/internal/module/landmarkitem/model"
)

func (store *landmarkItemStore) CreateLandmarkItem(ctx context.Context, data landmarkitemmodel.LandMarkItem) error {
	return store.db.WithContext(ctx).Table(landmarkitemmodel.LandMarkItem{}.TableName()).Create(&data).Error
}
