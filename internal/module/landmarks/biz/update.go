package landmarkbiz

import (
	"context"
	"errors"

	landmarkmodel "example.com/m/internal/module/landmarks/model"
)

type UpdateLandmarkRequest interface {
	Find(ctx context.Context, id int64, moreKeys []string) (*landmarkmodel.Landmark, error)
	UpdateLandmark(ctx context.Context, landmarkId uint, datas map[string]interface{}) error
}
type landmarkUpdateBiz struct {
	landmarkStore UpdateLandmarkRequest
}

func NewLandmarkUpdateBiz(landmarkStore UpdateLandmarkRequest) *landmarkUpdateBiz {
	return &landmarkUpdateBiz{landmarkStore: landmarkStore}
}
func (biz *landmarkUpdateBiz) UpdateLandmark(ctx context.Context, landmarkId uint, req landmarkmodel.UpdateLandmarkRequest) error {
	landmarkIdt, err := biz.landmarkStore.Find(ctx, int64(landmarkId), nil)
	if err != nil {
		return err
	}
	datas := make(map[string]interface{})
	if req.Name != "" && req.Name != landmarkIdt.Name {
		datas["name"] = req.Name
	}
	if req.Description != "" && req.Description != landmarkIdt.Description {
		datas["description"] = req.Description
	}
	if req.UrlImage != "" && req.UrlImage != landmarkIdt.UrlImage {
		datas["url_image"] = req.UrlImage
	}
	if len(datas) == 0 {
		return errors.New("no data to update")
	}
	return biz.landmarkStore.UpdateLandmark(ctx, landmarkId, datas)
}
