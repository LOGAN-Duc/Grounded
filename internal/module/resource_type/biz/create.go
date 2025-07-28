package resourcetypebiz

import (
	"context"
	"errors"

	resourcetypemodel "example.com/m/internal/module/resource_type/model"
)

type CreateResourceTypeStore interface {
	FindByName(ctx context.Context, id string) (*resourcetypemodel.ResourceType, error)
	Create(ctx context.Context, req *resourcetypemodel.ResourceType) error
}

type createResourceTypeBiz struct {
	store CreateResourceTypeStore
}

func NewCreateResourceTypeBiz(store CreateResourceTypeStore) *createResourceTypeBiz {
	return &createResourceTypeBiz{
		store: store,
	}
}

func (biz *createResourceTypeBiz) Create(ctx context.Context, req *resourcetypemodel.CreateResourceTypeRequest) error {
	if req.Name == "" {
		return errors.New("please! input name type of resource")
	}
	_, err := biz.store.FindByName(ctx, req.Name)
	if err == nil {
		return errors.New("name already exists")
	}
	reasourceType := &resourcetypemodel.ResourceType{
		Name: req.Name,
	}
	return biz.store.Create(ctx, reasourceType)
}
