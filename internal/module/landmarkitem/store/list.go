package landmarkitemstore

import (
	"context"

	landmarkitemmodel "example.com/m/internal/module/landmarkitem/model"
	"gorm.io/gorm"
)

func (s *landmarkItemStore) ListLandmarkItemsByLandmarkID(ctx context.Context, filter landmarkitemmodel.LandmarkListRequest) ([]*landmarkitemmodel.LandMarkItem, error) {
	var items []*landmarkitemmodel.LandMarkItem
	query := s.db.WithContext(ctx).Model(&landmarkitemmodel.LandMarkItem{})
	if filter.LandMarkID != 0 {
		query = query.Where("landmark_id = ?", filter.LandMarkID)
	}
	if filter.SearchText != "" {
		likePattern := "%" + filter.SearchText + "%"
		query = query.Where("name LIKE ?", likePattern)
	}
	query = query.Preload("Item", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "name", "url_image") // chọn các cột cần thiết
	}).Preload("Resource", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "name", "url_image")
	})

	if err := query.Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}
