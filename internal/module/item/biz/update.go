package itembiz

import (
	"context"
	"errors"
	"fmt"
	"strings"

	itemmodel "example.com/m/internal/module/item/model"
	itemresourcemodel "example.com/m/internal/module/item_resource/model"
	resourcemodel "example.com/m/internal/module/resource/model"
)

type UpdateItemStore interface {
	UpadteWithInterFace(ctx context.Context, id int, datas map[string]interface{}) error
	FindByID(ctx context.Context, id int) (*itemmodel.Item, error)
}
type FindByItemIdAndResourceIdStore interface {
	FindByItemId(ctx context.Context, id int) ([]*itemresourcemodel.ItemResource, error)
}
type FindResourceStore interface {
	FindByID(ctx context.Context, id int) (*resourcemodel.Resources, error)
}
type updateItemStore struct {
	store             UpdateItemStore
	itemResourceStore FindByItemIdAndResourceIdStore
	resource          FindResourceStore
}

func NewUpdateResourceBiz(store UpdateItemStore, itemResourceStore FindByItemIdAndResourceIdStore, resource FindResourceStore) *updateItemStore {
	return &updateItemStore{
		store:             store,
		itemResourceStore: itemResourceStore,
		resource:          resource,
	}
}
func (biz *updateItemStore) UpdateWithInterface(ctx context.Context, id int, req itemmodel.UpdateItemRequest) error {
	updates := map[string]interface{}{}
	item, err := biz.store.FindByID(ctx, id)
	if err != nil {
		return errors.New("item not found")
	}

	listItemResources, err := biz.itemResourceStore.FindByItemId(ctx, id)
	if err != nil {
		return errors.New("item-resources not found")
	}

	if req.Code == "" {
		var builder strings.Builder
		for _, itemResource := range listItemResources {
			resource, err := biz.resource.FindByID(ctx, itemResource.ResourceId)
			if err != nil {
				return fmt.Errorf("resource not found for id %d: %w", itemResource.ResourceId, err)
			}
			if resource.Status == "1" && itemResource.Quantity > 0 {
				for i := 0; i < itemResource.Quantity; i++ {
					builder.WriteString(resource.Code)
				}
			}
		}
		newCode := builder.String()

		if newCode != item.Code {
			updates["code"] = newCode
		}
	} else {
		if req.Code != item.Code {
			updates["code"] = req.Code
		}
	}
	if req.Name != "" && req.Name != item.Name {
		updates["name"] = req.Name
	}
	if req.ItemTypeId != nil && *req.ItemTypeId != item.ItemTypeId {
		updates["item_type_id"] = *req.ItemTypeId
	}

	if len(updates) == 0 {
		return nil
	}

	err = biz.store.UpadteWithInterFace(ctx, id, updates)
	if err != nil {
		return fmt.Errorf("failed to update item: %w", err)
	}

	return nil
}
