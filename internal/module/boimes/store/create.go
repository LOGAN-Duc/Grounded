package boimesstore

import (
	"context"

	boimesmodel "example.com/m/internal/module/boimes/model"
	"gorm.io/gorm"
)

func (s *boimesStore) Create(ctx context.Context, data *boimesmodel.Biome, resourceIds []int) error {
	return s.mysql.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(data).Error; err != nil {
			return err
		}
		if len(resourceIds) > 0 {
			var biomeResources []boimesmodel.BiomeResource
			for _, rid := range resourceIds {
				biomeResources = append(biomeResources, boimesmodel.BiomeResource{
					BiomeId:    &data.Id,
					ResourceId: (*int)(&rid),
				})
			}
			if err := tx.Table(boimesmodel.BiomeResource{}.TableName()).Create(&biomeResources).Error; err != nil {
				return err
			}
		}

		return nil
	})
}
