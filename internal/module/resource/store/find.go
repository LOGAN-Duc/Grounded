package resourcestore

import (
	"context"

	resourcemodel "example.com/m/internal/module/resource/model"
)

func (s *resourceStore) FindByName(ctx context.Context, nameRequest string) (*resourcemodel.Resources, error) {
	var resource resourcemodel.Resources
	err := s.mysql.WithContext(ctx).Table(resourcemodel.Resources{}.TableName()).Where("name = ?", nameRequest).First(&resource).Error
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

func (s *resourceStore) FindByID(ctx context.Context, id int) (*resourcemodel.Resources, error) {
	var resource resourcemodel.Resources
	err := s.mysql.WithContext(ctx).Table(resourcemodel.Resources{}.TableName()).Where("id = ?", id).First(&resource).Error
	if err != nil {
		return nil, err
	}
	return &resource, nil
}
