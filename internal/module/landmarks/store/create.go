package landmarkstore

import (
	"context"

	landmarkmodel "example.com/m/internal/module/landmarks/model"
)

func (s *landmarkStore) CreateLandmark(ctx context.Context, landmark *landmarkmodel.Landmark) error {
	return s.db.WithContext(ctx).Table(landmarkmodel.Landmark{}.TableName()).Create(landmark).Error
}
