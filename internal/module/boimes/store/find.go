package boimesstore

import (
	"context"

	boimesmodel "example.com/m/internal/module/boimes/model"
)

func (s *boimesStore) FindByName(ctx context.Context, nameRequest string) (*boimesmodel.Biome, error) {
	var data boimesmodel.Biome
	err := s.mysql.WithContext(ctx).Table(boimesmodel.Biome{}.TableName()).Where("name = ?", nameRequest).First(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}
func (s *boimesStore) FindById(ctx context.Context, id int) (*boimesmodel.Biome, error) {
	var data boimesmodel.Biome
	err := s.mysql.WithContext(ctx).Table(boimesmodel.Biome{}.TableName()).Where("id = ?", id).First(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}
