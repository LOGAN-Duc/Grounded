package itemstore

import (
	"context"

	itemmodel "example.com/m/internal/module/item/model"
	"gorm.io/gorm"
)

func (s *itemStore) UpadteWithInterFace(ctx context.Context, id int, datas map[string]interface{}) error {
	return s.mysql.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Table(itemmodel.Item{}.TableName()).Where("id = ?", id).Updates(&datas).Error; err != nil {
			return err
		}
		return nil
	})
}
