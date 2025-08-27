package boimebiz

import (
	"context"
	"errors"

	boimesmodel "example.com/m/internal/module/boimes/model"
	itemmodel "example.com/m/internal/module/item/model"
)

type CreateBoimeStore interface {
	Create(ctx context.Context, data *boimesmodel.Biome, resourceIds []int) error
	FindByName(ctx context.Context, nameRequest string) (*boimesmodel.Biome, error)
}
type UpdateItemStore interface {
	UpadteWithInterFace(ctx context.Context, id int, datas map[string]interface{}) error
	FindByID(ctx context.Context, id int) (*itemmodel.Item, error)
}
type createBoimeBiz struct {
	store     CreateBoimeStore
	itemStore UpdateItemStore
}

func NewCreateBoimeBiz(store CreateBoimeStore, itemStore UpdateItemStore) *createBoimeBiz {
	return &createBoimeBiz{
		store:     store,
		itemStore: itemStore,
	}
}
func (biz *createBoimeBiz) Create(ctx context.Context, req boimesmodel.CreateBiomeRequest) error {
	_, err := biz.store.FindByName(ctx, req.Name)
	if err == nil {
		return errors.New("biome with this name already exists")
	}
	data := &boimesmodel.Biome{
		Name:        req.Name,
		Url:         req.Url,
		Description: req.Description,
	}
	if err := biz.store.Create(ctx, data, req.ResourceIds); err != nil {
		return err
	}
	for _, itemID := range req.ItemIds {
		item, err := biz.itemStore.FindByID(ctx, itemID)
		if err != nil {
			return errors.New("item doesn't exist")
		}
		if item.BiomeId == nil {
			updateData := map[string]interface{}{
				"biome_id": data.Id,
			}
			if err := biz.itemStore.UpadteWithInterFace(ctx, itemID, updateData); err != nil {
				return err
			}
		} else {

		}
	}

	return nil
}
