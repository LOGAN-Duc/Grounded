package landmarkstore

import (
	"context"

	landmarkmodel "example.com/m/internal/module/landmarks/model"
)

func (s *landmarkStore) Find(ctx context.Context, id int64, moreKeys []string) (*landmarkmodel.Landmark, error) {
	var landmark landmarkmodel.Landmark
	query := s.db.WithContext(ctx).Table(landmarkmodel.Landmark{}.TableName()).Where("id = ? AND status = ?", id, 1)
	for _, key := range moreKeys {
		query = query.Preload(key)
	}

	err := query.First(&landmark).Error
	if err != nil {
		return nil, err
	}
	return &landmark, nil
}
func (s *landmarkStore) FindByName(ctx context.Context, name string, moreKeys []string) (*landmarkmodel.Landmark, error) {
	var landmark landmarkmodel.Landmark
	query := s.db.WithContext(ctx).Table(landmarkmodel.Landmark{}.TableName()).Where("name = ? AND status = ?", name, 1)

	for _, key := range moreKeys {
		query = query.Preload(key)
	}
	err := query.First(&landmark).Error

	if err != nil {
		return nil, err
	}
	return &landmark, nil
}
