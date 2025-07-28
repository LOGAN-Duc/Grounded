package itemresourcestore

import (
	"context"

	itemresourcemodel "example.com/m/internal/module/item_resource/model"
	"gorm.io/gorm"
)

func (s *itemResourceStore) Update(ctx context.Context, id int, datas map[string]interface{}) error {
	return s.mysql.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Table(itemresourcemodel.ItemResource{}.TableName()).Where("id = ?", id).Updates(&datas).Error; err != nil {
			return err
		}
		return nil
	})
}
