package resourcestore

import (
	"context"

	"example.com/m/internal/common"
	resourcemodel "example.com/m/internal/module/resource/model"
)

func (s *resourceStore) List(ctx context.Context, paging common.Paging,
	filter *resourcemodel.ListResourcesRequest, moreKeys []string) ([]resourcemodel.Resources, int64, error) {
	var resources []resourcemodel.Resources
	query := s.mysql.WithContext(ctx).Table(resourcemodel.Resources{}.TableName()).Where("status = ?", 1)

	if filter.Search != "" {
		query = query.Where("name LIKE ?", "%"+filter.Search+"%")
	}
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if paging.FakeCursor != "" {
		query = query.Where("id < ?", paging.FakeCursor)
	} else {
		offset := (paging.Page - 1) * paging.Limit
		query = query.Offset(int(offset))
	}
	for _, key := range moreKeys {
		query = query.Preload(key)
	}
	err := query.Limit(int(paging.Limit)).Find(&resources).Error
	if err != nil {
		return nil, 0, err
	}
	return resources, total, nil
}
