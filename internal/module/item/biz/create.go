package itembiz

import (
	"context"
	"errors"

	itemmodel "example.com/m/internal/module/item/model"
)

type CreateItemStore interface {
	Create(ctx context.Context, resource *itemmodel.Item) error
	FindByName(ctx context.Context, nameRequest string) (*itemmodel.Item, error)
}

type createItemBiz struct {
	store CreateItemStore
}

func NewCreateItemBiz(store CreateItemStore) *createItemBiz {
	return &createItemBiz{
		store: store,
	}
}
func (biz *createItemBiz) Create(ctx context.Context, req *itemmodel.CreateItemRequest) error {
	_, err := biz.store.FindByName(ctx, req.Name)
	if err == nil {
		return errors.New("resource with this name already exists")
	}
	resource := &itemmodel.Item{
		Name:       req.Name,
		Code:       req.Code,
		ItemTypeId: req.ItemTypeId,
	}

	return biz.store.Create(ctx, resource)
}
