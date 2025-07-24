package resourcestore

import (
	"context"

	resourcemodel "example.com/m/internal/module/resource/model"
)

func (s *resourceStore) UpadteWithInterFace(ctx context.Context, id int, datas map[string]interface{}) error {
	return s.mysql.WithContext(ctx).Table(resourcemodel.Resources{}.TableName()).Where("id = ?", id).Updates(&datas).Error
}
