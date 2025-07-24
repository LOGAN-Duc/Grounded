package itemtypestore

import (
	"context"

	itemtypemodel "example.com/m/internal/module/item_type/model"
)

func (s *itemTypeStore) Create(ctx context.Context, req *itemtypemodel.ItemType) error {
	return s.mysql.WithContext(ctx).Table(itemtypemodel.ItemType{}.TableName()).Create(req).Error
}
