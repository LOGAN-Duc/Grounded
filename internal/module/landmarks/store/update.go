package landmarkstore

import (
	"context"

	landmarkmodel "example.com/m/internal/module/landmarks/model"
)

func (s *landmarkStore) UpdateLandmark(ctx context.Context, landmarkId uint, datas map[string]interface{}) error {
	return s.db.WithContext(ctx).Table(landmarkmodel.Landmark{}.TableName()).
		Where("id = ?", landmarkId).Updates(&datas).Error
}
