package itemtypestore

import (
	"context"

	itemtypemodel "example.com/m/internal/module/item_type/model"
)

func (s *itemTypeStore) FindById(ctx context.Context, id int) (*itemtypemodel.ItemType, error) {
	var data itemtypemodel.ItemType
	err := s.mysql.WithContext(ctx).Table(itemtypemodel.ItemType{}.TableName()).Where("id = ?", id).First(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}
