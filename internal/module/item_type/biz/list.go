package itemtypeypebiz

import (
	"context"

	"example.com/m/internal/common"
	itemtypemodel "example.com/m/internal/module/item_type/model"
)

type ListItemTypeStore interface {
	List(ctx context.Context, paging common.Paging, req *itemtypemodel.ListItemTypeRequest, moreKeys []string) ([]itemtypemodel.ItemType, int64, error)

}


type listItemTypeStore struct {
	store ListItemTypeStore
}

func NewListTypeBiz(store ListItemTypeStore) *listItemTypeStore {
	return &listItemTypeStore{
		store: store,
	}
}
func (biz *listItemTypeStore) List(ctx context.Context, paging common.Paging,
	filter *itemtypemodel.ListItemTypeRequest, moreKeys []string) ([]itemtypemodel.ItemType, int64, error) {
	items, total, err := biz.store.List(ctx, paging, filter, moreKeys)
	if err != nil {
		return nil, 0, err
	}
	return items, total, nil
}
