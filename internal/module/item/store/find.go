package itemstore

import (
	"context"

	itemmodel "example.com/m/internal/module/item/model"
)

func (s *itemStore) FindByName(ctx context.Context, nameRequest string) (*itemmodel.Item, error) {
	var item itemmodel.Item
	err := s.mysql.WithContext(ctx).Table(itemmodel.Item{}.TableName()).Where("name = ?", nameRequest).First(&item).Error
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *itemStore) FindByID(ctx context.Context, id int) (*itemmodel.Item, error) {
	var item itemmodel.Item
	err := s.mysql.WithContext(ctx).Table(itemmodel.Item{}.TableName()).Where("id = ?", id).First(&item).Error
	if err != nil {
		return nil, err
	}
	return &item, nil
}
