package itemtypestore

import (
	"context"

	"example.com/m/internal/common"
	itemtypemodel "example.com/m/internal/module/item_type/model"
)

func (s *itemTypeStore) List(ctx context.Context, paging common.Paging, req *itemtypemodel.ListItemTypeRequest, moreKeys []string) ([]itemtypemodel.ItemType, int64, error) {
	var datas []itemtypemodel.ItemType
	query := s.mysql.WithContext(ctx).Table(itemtypemodel.ItemType{}.TableName()).Where("status = ?", "1")

	if req.Search != "" {
		query.Where("name Like ?", "%"+req.Search+"%")
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
	err := query.Limit(int(paging.Limit)).Find(&datas).Error
	if err != nil {
		return nil, 0, err
	}
	return datas, total, nil
}
