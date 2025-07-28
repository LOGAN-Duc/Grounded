package resourcetypestore

import (
	"context"

	resourcetypemodel "example.com/m/internal/module/resource_type/model"
)

func (s *resourcesTypeStore) FindById(ctx context.Context, id int) (*resourcetypemodel.ResourceType, error) {
	var data resourcetypemodel.ResourceType
	err := s.mysql.WithContext(ctx).Table(resourcetypemodel.ResourceType{}.TableName()).Where("id = ?", id).First(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (s *resourcesTypeStore) FindByName(ctx context.Context, id string) (*resourcetypemodel.ResourceType, error) {
	var data resourcetypemodel.ResourceType
	err := s.mysql.WithContext(ctx).Table(resourcetypemodel.ResourceType{}.TableName()).Where("name = ?", id).First(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}
