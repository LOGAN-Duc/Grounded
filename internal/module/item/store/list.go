package itemstore

import (
	"context"

	"example.com/m/internal/common"
	itemmodel "example.com/m/internal/module/item/model"
)

func (s *itemStore) List(ctx context.Context, paging common.Paging,
	filter *itemmodel.ListItemRequest, moreKeys []string) ([]itemmodel.Item, int64, error) {
	var items []itemmodel.Item
	query := s.mysql.WithContext(ctx).Table(itemmodel.Item{}.TableName()).Where("status = ?", 1)

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
	err := query.Limit(int(paging.Limit)).Find(&items).Error
	if err != nil {
		return nil, 0, err
	}
	return items, total, nil
}
