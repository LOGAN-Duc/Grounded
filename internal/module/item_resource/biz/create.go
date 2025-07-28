package itemresourcebiz

import (
	"context"
	"errors"
	"fmt"

	itemmodel "example.com/m/internal/module/item/model"
	itemresourcemodel "example.com/m/internal/module/item_resource/model"
	resourcemodel "example.com/m/internal/module/resource/model"
	"gorm.io/gorm"
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
func (biz *createItemResourceBiz) Create(ctx context.Context, id int, req []itemresourcemodel.CreateItemRrsourceRequest) error {
	for _, itemReq := range req {
		// Kiểm tra xem item-resource đã tồn tại chưa
		// Tìm resource theo ID
		resource, err := biz.resourceStore.FindByID(ctx, itemReq.ResourceId)
		if err != nil {
			return errors.New("resource does not exist")
		}
		// Tìm item theo ID
		item, err := biz.itemStore.FindByID(ctx, id)
		if err != nil {
			return errors.New("item does not exist")
		}

		_, err = biz.store.FindByItemIdAndResourceId(ctx, id, itemReq.ResourceId)
		if err == nil {
			return fmt.Errorf("resource %s already exists in %s", resource.Name, item.Name)
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		// Cập nhật code cho item
		for i := 0; i < itemReq.Quantity; i++ {
			item.Code += resource.Code
		}

		// Cập nhật item
		updates := map[string]interface{}{
			"code": item.Code,
		}
		err = biz.itemStore.UpadteWithInterFace(ctx, id, updates)
		if err != nil {
			return errors.New("failed to update item")
		}

		// Tạo mới bản ghi item-resource
		data := &itemresourcemodel.ItemResource{
			ItemId:     id,
			ResourceId: itemReq.ResourceId,
			Quantity:   itemReq.Quantity,
		}

		// Lưu item-resource vào cơ sở dữ liệu
		if err := biz.store.Create(ctx, data); err != nil {
			return err
		}
	}
	return nil
}
