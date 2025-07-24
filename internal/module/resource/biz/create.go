package resourcebiz

import (
	"context"
	"errors"

	resourcemodel "example.com/m/internal/module/resource/model"
	resourcetypemodel "example.com/m/internal/module/resource_type/model"
)

type CreateResourcesStore interface {
	Create(ctx context.Context, resource *resourcemodel.Resources) error
	FindByName(ctx context.Context, nameRequest string) (*resourcemodel.Resources, error)
}

type FindByIdResourceTypeStore interface {
	FindById(ctx context.Context, id int) (*resourcetypemodel.ResourceType, error)
}
type createResourcesBiz struct {
	store      CreateResourcesStore
	typerstore FindByIdResourceTypeStore
}

func NewCreateResourcesBiz(store CreateResourcesStore, typerstore FindByIdResourceTypeStore) *createResourcesBiz {
	return &createResourcesBiz{
		store:      store,
		typerstore: typerstore,
	}
}
func (biz *createResourcesBiz) Create(ctx context.Context, req *resourcemodel.CreateResourcesRequest) error {
	_, err := biz.store.FindByName(ctx, req.Name)
	if err == nil {
		return errors.New("resource with this name already exists")
	}
	_, err = biz.typerstore.FindById(ctx, req.ResourceTypeId)
	if err != nil {
		return errors.New(" resource type does not exist")
	}
	resource := &resourcemodel.Resources{
		Name:           req.Name,
		Code:           req.Code,
		ResourceTypeId: req.ResourceTypeId,
	}

	return biz.store.Create(ctx, resource)
}
