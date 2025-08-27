package boimesstore

import (
	"context"
	"unsafe"

	boimesmodel "example.com/m/internal/module/boimes/model"
	"gorm.io/gorm"
)

func (s *boimesStore) Update(ctx context.Context, id uint, data interface{}, resourceIds []int) error {
	return s.mysql.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&boimesmodel.Biome{}).Where("id = ?", id).Updates(data).Error; err != nil {
			return err
		}
		if err := tx.Where("biome_id = ?", id).Delete(&boimesmodel.BiomeResource{}).Error; err != nil {
			return err
		}
		if len(resourceIds) > 0 {
			var biomeResources []boimesmodel.BiomeResource
			for _, rid := range resourceIds {
				biomeResources = append(biomeResources, boimesmodel.BiomeResource{
					BiomeId:    (*int)(unsafe.Pointer(&id)),
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
