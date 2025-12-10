package landmarkbiz

import (
	"context"

	"example.com/m/internal/common"
	landmarkmodel "example.com/m/internal/module/landmarks/model"
)

type ListLandmarksRequest interface {
	ListLandmarks(ctx context.Context, paging common.Paging,
		filter *landmarkmodel.ListLandmarksRequest, moreKeys []string) ([]*landmarkmodel.Landmark, int64, error)
}
type landmarkListBiz struct {
	landmarkStore ListLandmarksRequest
}

func NewLandmarkListBiz(landmarkStore ListLandmarksRequest) *landmarkListBiz {
	return &landmarkListBiz{landmarkStore: landmarkStore}
}
func (biz *landmarkListBiz) ListLandmarks(ctx context.Context, paging common.Paging,
	filter *landmarkmodel.ListLandmarksRequest, moreKeys []string) ([]*landmarkmodel.Landmark, int64, error) {
	return biz.landmarkStore.ListLandmarks(ctx, paging, filter, moreKeys)
}
