package itemtypeypebiz

import (
	"context"
	"errors"

	itemtypemodel "example.com/m/internal/module/item_type/model"
	temtypemodel "example.com/m/internal/module/item_type/model"
)

type CreateItemTypeStore interface {
	Create(ctx context.Context, req *temtypemodel.ItemType) error
	FindByName(ctx context.Context, id string) (*itemtypemodel.ItemType, error)
}

type createItemTypeBiz struct {
	store CreateItemTypeStore
}

func NewCreateItemTypeBiz(store CreateItemTypeStore) *createItemTypeBiz {
	return &createItemTypeBiz{
		store: store,
	}
}

func (biz *createItemTypeBiz) Create(ctx context.Context, req *temtypemodel.CreateItemTypeRequest) error {
	_, err := biz.store.FindByName(ctx, req.Name)
	if err == nil {
		return errors.New("name already exists")
	}
	itemType := &temtypemodel.ItemType{
		Name: req.Name,
	}
	return biz.store.Create(ctx, itemType)
}
