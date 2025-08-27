package itembiz

import (
	"context"
	"errors"

	boimesmodel "example.com/m/internal/module/boimes/model"
	itemmodel "example.com/m/internal/module/item/model"
)

type CreateItemStore interface {
	Create(ctx context.Context, resource *itemmodel.Item) error
	FindByName(ctx context.Context, nameRequest string) (*itemmodel.Item, error)
}
type FindBoimeStore interface {
	FindById(ctx context.Context, nameRequest int) (*boimesmodel.Biome, error)
}
type createItemBiz struct {
	store      CreateItemStore
	boimeStore FindBoimeStore
}

func NewCreateItemBiz(store CreateItemStore, boimeStore FindBoimeStore) *createItemBiz {
	return &createItemBiz{
		store:      store,
		boimeStore: boimeStore,
	}
}
func (biz *createItemBiz) Create(ctx context.Context, req *itemmodel.CreateItemRequest) error {
	_, err := biz.store.FindByName(ctx, req.Name)
	if err == nil {
		return errors.New("resource with this name already exists")
	}
	_, err = biz.boimeStore.FindById(ctx, *req.BoimeId)
	if err != nil {
		return errors.New("boime dosen't exists")
	}
	resource := &itemmodel.Item{
		Name:       req.Name,
		Code:       req.Code,
		ItemTypeId: req.ItemTypeId,
		BiomeId:    req.BoimeId,
	}

	return biz.store.Create(ctx, resource)
}
