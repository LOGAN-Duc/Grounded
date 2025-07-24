package itembiz

import (
	"context"
	"errors"

	itemmodel "example.com/m/internal/module/item/model"
)

type UpdateItemStore interface {
	UpadteWithInterFace(ctx context.Context, id int, datas map[string]interface{}) error
	FindByID(ctx context.Context, id int) (*itemmodel.Item, error)
}

type updateItemStore struct {
	store UpdateItemStore
}

func NewUpdateResourceBiz(store UpdateItemStore) *updateItemStore {
	return &updateItemStore{
		store: store,
	}
}

func (biz *updateItemStore) UpadteWithInterFace(ctx context.Context, id int, req itemmodel.UpdateItemRequest) error {
	updates := make(map[string]interface{})
	item, err := biz.store.FindByID(ctx, id)
	if err != nil {
		return errors.New("item not found")
	}

	if req.Name != "" && req.Name != item.Name {
		updates["name"] = req.Name
	}
	if req.Code != "" && req.Code != item.Code {
		updates["code"] = req.Code
	}
	if *req.ItemTypeId != item.ItemTypeId {
		updates["item_type_id"] = *req.ItemTypeId
	}
	if len(updates) == 0 {
		return errors.New("no data to update")
	}
	err = biz.store.UpadteWithInterFace(ctx, int(item.Id), updates)
	if err != nil {
		return err
	}
	return nil
}
