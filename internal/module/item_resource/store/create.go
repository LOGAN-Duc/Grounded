package itemresourcestore

import (
	"context"

	itemresourcemodel "example.com/m/internal/module/item_resource/model"
)

func (s *itemResourceStore) Create(ctx context.Context, data *itemresourcemodel.ItemResource) error {
	if err := s.mysql.WithContext(ctx).Create(data).Error; err != nil {
		return err
	}
	return nil
}
 