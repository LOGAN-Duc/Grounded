package itemresourcebiz

import (
	"context"
	"errors"

	itemmodel "example.com/m/internal/module/item/model"
	itemresourcemodel "example.com/m/internal/module/item_resource/model"
	resourcemodel "example.com/m/internal/module/resource/model"
)

type CreateItemRrsourceStore interface {
	Create(ctx context.Context, data *itemresourcemodel.ItemResource) error
	FindByItemIdAndResourceId(ctx context.Context, itemId int, resourceId int) (*itemresourcemodel.ItemResource, error)
	FindByResourceId(ctx context.Context, id int) (*itemresourcemodel.ItemResource, error)
}
type FindResourceStore interface {
	FindByID(ctx context.Context, id int) (*resourcemodel.Resources, error)
}
type FindItemStore interface {
	FindByID(ctx context.Context, id int) (*itemmodel.Item, error)
	UpadteWithInterFace(ctx context.Context, id int, datas map[string]interface{}) error
}
type createItemResourceBiz struct {
	store         CreateItemRrsourceStore
	resourceStore FindResourceStore
	itemStore     FindItemStore
}

func NewCreateItemResourceBiz(store CreateItemRrsourceStore, resourceStore FindResourceStore, itemStore FindItemStore) *createItemResourceBiz {
	return &createItemResourceBiz{
		store:         store,
		resourceStore: resourceStore,
		itemStore:     itemStore,
	}
}

func (biz *createItemResourceBiz) Create(ctx context.Context, req itemresourcemodel.CreateItemRrsourceRequest) error {
	// Kiểm tra xem item-resource đã tồn tại chưa
	_, err := biz.store.FindByItemIdAndResourceId(ctx, req.ItemId, req.ResourceId)
	if err == nil {
		// Kiểm tra lỗi nếu không tìm thấy item-resource
		return errors.New("item-resource does not exist")
	}

	// Tìm resource theo ID
	resource, err := biz.resourceStore.FindByID(ctx, req.ResourceId)
	if err != nil {
		return errors.New("resource does not exist")
	}

	// Tìm item theo ID
	item, err := biz.itemStore.FindByID(ctx, req.ItemId)
	if err != nil {
		return errors.New("item does not exist")
	}

	// Cập nhật code cho item
	for i := 0; i < req.Quantity; i++ {
		item.Code += resource.Code
	}
	
	// Tạo bản đồ cập nhật
	updates := map[string]interface{}{
		"code": item.Code,
	}

	// Cập nhật item trong cơ sở dữ liệu
	err = biz.itemStore.UpadteWithInterFace(ctx, req.ItemId, updates)
	if err != nil {
		return errors.New("failed to update item")
	}

	// Tạo mới một bản ghi item-resource
	data := &itemresourcemodel.ItemResource{
		ItemId:     req.ItemId,
		ResourceId: req.ResourceId,
		Quantity:   req.Quantity,
	}

	// Lưu item-resource vào cơ sở dữ liệu
	return biz.store.Create(ctx, data)
}
