package itemstore

import (
	"context"

	itemmodel "example.com/m/internal/module/item/model"
)

func (s *itemStore) Create(ctx context.Context, resource *itemmodel.Item) error {
	return s.mysql.WithContext(ctx).Table(itemmodel.Item{}.TableName()).Create(resource).Error
}
