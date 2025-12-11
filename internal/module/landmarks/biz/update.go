package landmarkbiz

import (
	"context"
	"errors"

	landmarkitemmodel "example.com/m/internal/module/landmarkitem/model"
	landmarkmodel "example.com/m/internal/module/landmarks/model"
	resourcemodel "example.com/m/internal/module/resource/model"
)

type UpdateLandmarkRequest interface {
	Find(ctx context.Context, id int64, moreKeys []string) (*landmarkmodel.Landmark, error)
	UpdateLandmark(ctx context.Context, landmarkId uint, datas map[string]interface{}) error
}
type RecountResource interface {
	FindByID(ctx context.Context, id int) (*resourcemodel.Resources, error)
	UpadteWithInterFace(ctx context.Context, id int, datas map[string]interface{}) error
}
type LandmarkUpdateBizInterface interface {
	CreateLandmarkItem(ctx context.Context, data landmarkitemmodel.LandMarkItem) error
}

type landmarkUpdateBiz struct {
	landmarkStore   UpdateLandmarkRequest
	recountResource RecountResource
	landmarkItem    LandmarkUpdateBizInterface
}

func NewLandmarkUpdateBiz(landmarkStore UpdateLandmarkRequest, recountResource RecountResource, landmarkItem LandmarkUpdateBizInterface) *landmarkUpdateBiz {
	return &landmarkUpdateBiz{landmarkStore: landmarkStore, recountResource: recountResource, landmarkItem: landmarkItem}
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

func (biz *landmarkUpdateBiz) AddItemsAndResources(ctx context.Context, landmarkId int, req landmarkitemmodel.LandmarkAddRequest) error {
	// Lấy landmark hiện tại kèm items & resources
	landmarkIdt, err := biz.landmarkStore.Find(ctx, int64(landmarkId), []string{"LandmarkItems"})
	if err != nil {
		return err
	}

	// --- Thêm Items ---
	existingItemMap := make(map[int]bool)
	for _, item := range landmarkIdt.LandmarkItems {
		existingItemMap[item.ItemId] = true
	}

	for _, itemId := range req.ItemIds {
		if existingItemMap[itemId] {
			continue // bỏ qua nếu đã tồn tại
		}

		newItem := landmarkitemmodel.LandMarkItem{
			LandmarkId: int(landmarkId),
			ItemId:     itemId,
		}

		if err := biz.landmarkItem.CreateLandmarkItem(ctx, newItem); err != nil {
			return err
		}
	}

	// --- Thêm Resources ---
	existingResourceMap := make(map[int]bool)
	for _, item := range landmarkIdt.LandmarkItems {
		existingResourceMap[item.ResourceId] = true
	}

	for _, resId := range req.ResourceIds {
		if existingResourceMap[resId] {
			continue // bỏ qua nếu đã tồn tại
		}

		newResource := landmarkitemmodel.LandMarkItem{
			LandmarkId: int(landmarkId),
			ResourceId: resId,
		}

		if err := biz.landmarkItem.CreateLandmarkItem(ctx, newResource); err != nil {
			return err
		}
	}

	return nil
}
