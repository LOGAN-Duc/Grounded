package resourcestore

import (
	"context"

	resourcemodel "example.com/m/internal/module/resource/model"
)

func (s *resourceStore) Create(ctx context.Context, resource *resourcemodel.Resources) error {
	return s.mysql.WithContext(ctx).Table(resourcemodel.Resources{}.TableName()).Create(resource).Error
}
