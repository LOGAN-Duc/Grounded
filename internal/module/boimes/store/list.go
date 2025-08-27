package boimesstore

import (
	"context"

	"example.com/m/internal/common"
	boimesmodel "example.com/m/internal/module/boimes/model"
)

func (s *boimesStore) List(ctx context.Context, paging common.Paging,
	filter *boimesmodel.FilterBiomeRequest, moreKeys []string) ([]boimesmodel.Biome, int64, error) {
	var datas []boimesmodel.Biome
	query := s.mysql.WithContext(ctx).Table(boimesmodel.Biome{}.TableName()).Where("status = ?", 1)

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
	query = query.Preload("Items").Preload("Items.ItemType")

	for _, key := range moreKeys {
		query = query.Preload(key)
	}
	err := query.Limit(int(paging.Limit)).Find(&datas).Error
	if err != nil {
		return nil, 0, err
	}
	return datas, total, nil
}
