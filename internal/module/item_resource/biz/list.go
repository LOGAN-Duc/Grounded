package itemresourcebiz

import (
	"context"

	"example.com/m/internal/common"
	itemresourcemodel "example.com/m/internal/module/item_resource/model"
)

type ListItemResourceStore interface {
	List(ctx context.Context, paging common.Paging, filter *itemresourcemodel.ListItemRrsourceRequest, moreKeys []string) ([]itemresourcemodel.ItemResource, int64, error)
}
type listItemResourceBiz struct {
	store ListItemResourceStore
}

func NewListItemResourceBiz(store ListItemResourceStore) *listItemResourceBiz {
	return &listItemResourceBiz{
		store: store,
	}
}
func (biz *listItemResourceBiz) List(ctx context.Context, paging common.Paging,
	filter *itemresourcemodel.ListItemRrsourceRequest, moreKeys []string) ([]itemresourcemodel.ItemResource, int64, error) {
	items, total, err := biz.store.List(ctx, paging, filter, moreKeys)
	if err != nil {
		return nil, 0, err
	}
	return items, total, nil
}
