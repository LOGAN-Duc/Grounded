package resourcetypebiz

import (
	"context"

	"example.com/m/internal/common"
	resourcetypemodel "example.com/m/internal/module/resource_type/model"
)

type ListResourceTypeStore interface {
	List(ctx context.Context, paging common.Paging, req *resourcetypemodel.ListResourceTypeRequest, moreKeys []string) ([]resourcetypemodel.ResourceType, int64, error)
}

type listResourceTypeStore struct {
	store ListResourceTypeStore
}

func NewListTypeBiz(store ListResourceTypeStore) *listResourceTypeStore {
	return &listResourceTypeStore{
		store: store,
	}
}
func (biz *listResourceTypeStore) List(ctx context.Context, paging common.Paging,
	filter *resourcetypemodel.ListResourceTypeRequest, moreKeys []string) ([]resourcetypemodel.ResourceType, int64, error) {
	items, total, err := biz.store.List(ctx, paging, filter, moreKeys)
	if err != nil {
		return nil, 0, err
	}
	return items, total, nil
}
