package resourcetypestore

import (
	"context"

	resourcetypemodel "example.com/m/internal/module/resource_type/model"
)

func (s *resourcesTypeStore) Create(ctx context.Context, req *resourcetypemodel.ResourceType) error {
	return s.mysql.WithContext(ctx).Table(resourcetypemodel.ResourceType{}.TableName()).Create(req).Error
}
