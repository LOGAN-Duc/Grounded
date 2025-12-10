package landmarkstore

import (
	"context"

	"example.com/m/internal/common"
	landmarkmodel "example.com/m/internal/module/landmarks/model"
)

func (s *landmarkStore) ListLandmarks(ctx context.Context, paging common.Paging,
	filter *landmarkmodel.ListLandmarksRequest, moreKeys []string) ([]*landmarkmodel.Landmark, int64, error) {
	var landmarks []*landmarkmodel.Landmark
	query := s.db.WithContext(ctx).Table(landmarkmodel.Landmark{}.TableName()).Where("status = ?", 1)
	var total int64
	if filter.Search != "" {
		query = query.Where("name LIKE ?", "%"+filter.Search+"%")
	}
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	for _, key := range moreKeys {
		query = query.Preload(key)
	}
	if paging.FakeCursor != "" {
		query = query.Where("id < ?", paging.FakeCursor)
	} else {
		offset := (paging.Page - 1) * paging.Limit
		query = query.Offset(int(offset))
	}
	err := query.Limit(int(paging.Limit)).Find(&landmarks).Error
	if err != nil {
		return nil, 0, err
	}
	return landmarks, total, nil
}
