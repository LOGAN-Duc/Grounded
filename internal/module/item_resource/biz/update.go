package itemresourcebiz

import (
	"context"

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

// func (biz *updateItemResourceBiz) Update(ctx context.Context, id int, idR int, req []itemresourcemodel.UpdateItemRrsourceRequest) error {
// 	datas, err := biz.store.FindByItemIdAndResourceId(ctx, id, idR)
// 	if err != nil {
// 		return errors.New("item dosen't exist")
// 	}

// }
