package itemresourcebiz

import (
	"context"
	"errors"
	"fmt"
	"strings"

	itemresourcemodel "example.com/m/internal/module/item_resource/model"
)

type DeleteItemResourceStore interface {
	Delete(ctx context.Context, id int) error
	FindById(ctx context.Context, id int) (*itemresourcemodel.ItemResource, error)
}

type deleteItemResourceBiz struct {
	store         DeleteItemResourceStore
	resourceStore FindResourceStore
	itemStore     FindItemStore
}

func NewDeleteItemResourceBiz(store DeleteItemResourceStore, resourceStore FindResourceStore,
	itemStore FindItemStore) *deleteItemResourceBiz {
	return &deleteItemResourceBiz{
		store:         store,
		resourceStore: resourceStore,
		itemStore:     itemStore,
	}
}
func (biz *deleteItemResourceBiz) Delete(ctx context.Context, id int) error {
	itemResource, err := biz.store.FindById(ctx, id)
	if err != nil {
		return errors.New("item-resource not found")
	}
	resource, err := biz.resourceStore.FindByID(ctx, itemResource.ResourceId)
	if err != nil {
		return errors.New("resource does not exist")
	}
	fmt.Print("itemid...", itemResource.ItemId)
	item, err := biz.itemStore.FindByID(ctx, itemResource.ItemId)
	if err != nil {
		return errors.New("item does not exist")
	}

	// Cập nhật code cho item
	for i := 0; i < itemResource.Quantity; i++ {
		item.Code = strings.Replace(item.Code, resource.Code, "", 1)
	}

	// Tạo bản đồ cập nhật
	updates := map[string]interface{}{
		"code": item.Code,
	}
	err = biz.itemStore.UpadteWithInterFace(ctx, itemResource.ItemId, updates)
	if err != nil {
		return errors.New("failed to update item")
	}
	return biz.store.Delete(ctx, id)
}
