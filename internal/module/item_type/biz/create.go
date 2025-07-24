package itemtypeypebiz

import (
	"context"
	"errors"

	temtypemodel "example.com/m/internal/module/item_type/model"
)

type CreateItemTypeStore interface {
	Create(ctx context.Context, req *temtypemodel.ItemType) error
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
	if req.Name == "" {
		return errors.New("please! input name type of Item")
	}
	itemType := &temtypemodel.ItemType{
		Name: req.Name,
	}
	return biz.store.Create(ctx, itemType)
}
