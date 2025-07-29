package itemresourcestore

import (
	"context"

	"example.com/m/internal/common"
	itemresourcemodel "example.com/m/internal/module/item_resource/model"
)

func (s *itemResourceStore) List(ctx context.Context, paging common.Paging, filter *itemresourcemodel.ListItemRrsourceRequest, moreKeys []string) ([]itemresourcemodel.ItemResource, int64, error) {
	var items []itemresourcemodel.ItemResource
	query := s.mysql.WithContext(ctx).Table(itemresourcemodel.ItemResource{}.TableName()).Where("status = ?", 1)

	if filter.Search != "" {
		query = query.Where("item_id = ?", filter.Search)
	}
	if filter.ItemId != nil {
		query = query.
			Where("item_id = ?", filter.ItemId)
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
