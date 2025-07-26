package itemresourcestore

import (
	"context"

	itemresourcemodel "example.com/m/internal/module/item_resource/model"
)

func (s *itemResourceStore) FindByItemIdAndResourceId(ctx context.Context, itemId int, resourceId int) (*itemresourcemodel.ItemResource, error) {
	db := s.mysql.WithContext(ctx).Table(itemresourcemodel.ItemResource{}.TableName())
	var data *itemresourcemodel.ItemResource
	if err := db.Where("item_id = ? AND resource_id = ? AND status = ?", itemId, resourceId, "1").First(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}
func (s *itemResourceStore) FindByResourceId(ctx context.Context, id int) (*itemresourcemodel.ItemResource, error) {
	db := s.mysql.WithContext(ctx).Table(itemresourcemodel.ItemResource{}.TableName())
	var data *itemresourcemodel.ItemResource
	if err := db.Where("itemresource_id = ? AND status = ?", id, "1").First(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}
func (s *itemResourceStore) FindByItemId(ctx context.Context, id int) ([]*itemresourcemodel.ItemResource, error) {
	db := s.mysql.WithContext(ctx).Table(itemresourcemodel.ItemResource{}.TableName())
	var data []*itemresourcemodel.ItemResource
	if err := db.Where("item_id = ? AND status = ?", id, "1").Find(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}
func (s *itemResourceStore) FindById(ctx context.Context, id int) (*itemresourcemodel.ItemResource, error) {
	db := s.mysql.WithContext(ctx).Table(itemresourcemodel.ItemResource{}.TableName())
	var data *itemresourcemodel.ItemResource
	if err := db.Where("id = ?", id).First(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}
