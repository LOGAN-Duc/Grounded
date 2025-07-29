package itemresourcestore

import (
	"context"

	itemresourcemodel "example.com/m/internal/module/item_resource/model"
)

func (s *itemResourceStore) Delete(ctx context.Context) error {
	return s.mysql.WithContext(ctx).Table(itemresourcemodel.ItemResource{}.TableName()).Where("quantity = ?", 0).Update("status", "0").Error
}
