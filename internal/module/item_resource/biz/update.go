package itemresourcebiz

import (
	"context"
	"errors"

	itemresourcemodel "example.com/m/internal/module/item_resource/model"
)

type UpdateItemResourceRequest interface {
	Update(ctx context.Context, id int, datas map[string]interface{}) error
	FindByItemIdAndResourceId(ctx context.Context, itemId int, resourceId int) (*itemresourcemodel.ItemResource, error)
}

type updateItemResourceBiz struct {
	store UpdateItemResourceRequest
}

func NewUpdateItemResourceBiz(store UpdateItemResourceRequest) *updateItemResourceBiz {
	return &updateItemResourceBiz{
		store: store,
	}
}

func (biz *updateItemResourceBiz) Update(ctx context.Context, idR int, req []itemresourcemodel.UpdateItemRrsourceRequest) error {
	for _, resourceReq := range req {
		itemResource, err := biz.store.FindByItemIdAndResourceId(ctx, idR, resourceReq.ResourceId)
		if err != nil {
			return errors.New("reource not found")
		}

		updates := map[string]interface{}{}
		if *resourceReq.Quantity != itemResource.Quantity {
			updates["quantity"] = resourceReq.Quantity
		}

		// Nếu không có gì để cập nhật, bỏ qua
		if len(updates) > 0 {
			// Lưu thay đổi vào store
			if err := biz.store.Update(ctx, int(itemResource.Id), updates); err != nil {
				return err
			}
		}
	}
	return nil
}
