package landmarkitembiz

import (
	"context"

	landmarkitemmodel "example.com/m/internal/module/landmarkitem/model"
)

type LandmarkItemListStore interface {
	ListLandmarkItemsByLandmarkID(ctx context.Context, filter landmarkitemmodel.LandmarkListRequest) ([]*landmarkitemmodel.LandMarkItem, error)
}
type landmarkItemBiz struct {
	store LandmarkItemListStore
}

func NewLandmarkItemBiz(store LandmarkItemListStore) *landmarkItemBiz {
	return &landmarkItemBiz{store: store}
}

func (biz *landmarkItemBiz) ListLandmarkItemsByLandmarkID(ctx context.Context, filter landmarkitemmodel.LandmarkListRequest) ([]*landmarkitemmodel.LandMarkItem, error) {
	return biz.store.ListLandmarkItemsByLandmarkID(ctx, filter)
}
