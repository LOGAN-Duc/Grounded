package resourcebiz

import (
	"context"
	"errors"

	resourcemodel "example.com/m/internal/module/resource/model"
)

type UpdateResouceStore interface {
	UpadteWithInterFace(ctx context.Context, id int, datas map[string]interface{}) error
	FindByID(ctx context.Context, id int) (*resourcemodel.Resources, error)
}

type updateResourceStore struct {
	store UpdateResouceStore
}

func NewUpdateResourceBiz(store UpdateResouceStore) *updateResourceStore {
	return &updateResourceStore{
		store: store,
	}
}

func (biz *updateResourceStore) UpadteWithInterFace(ctx context.Context, id int, req resourcemodel.UpdateResourcesRequest) error {
	updates := make(map[string]interface{})
	resource, err := biz.store.FindByID(ctx, id)
	if err != nil {
		return errors.New("resource not found")
	}

	if req.Name != "" && req.Name != resource.Name {
		updates["name"] = req.Name
	}
	if req.Code != "" && req.Code != resource.Code {
		updates["code"] = req.Code
	}
	if req.ResourceTypeId != nil && *req.ResourceTypeId != resource.ResourceTypeId {
		updates["resource_type_id"] = *req.ResourceTypeId
	}
	if len(updates) == 0 {
		return nil
	}
	err = biz.store.UpadteWithInterFace(ctx, int(resource.Id), updates)
	if err != nil {
		return err
	}
	return nil
}
