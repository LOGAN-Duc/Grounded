package itembiz

import (
	"context"

	"example.com/m/internal/common"
	itemmodel "example.com/m/internal/module/item/model"
)

type ListItemStore interface {
	List(ctx context.Context, paging common.Paging,
		filter *itemmodel.ListItemRequest, moreKeys []string) ([]itemmodel.Item, int64, error)
}
type listItemBiz struct {
	store ListItemStore
}

func NewListItemBiz(store ListItemStore) *listItemBiz {
	return &listItemBiz{
		store: store,
	}
}
func (biz *listItemBiz) List(ctx context.Context, paging common.Paging,
	filter *itemmodel.ListItemRequest, moreKeys []string) ([]itemmodel.Item, int64, error) {
	items, total, err := biz.store.List(ctx, paging, filter, moreKeys)
	if err != nil {
		return nil, 0, err
	}
	return items, total, nil
}
