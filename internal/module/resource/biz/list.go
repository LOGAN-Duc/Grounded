package resourcebiz

import (
	"context"

	"example.com/m/internal/common"
	resourcemodel "example.com/m/internal/module/resource/model"
)

type ListResourcesStore interface {
	List(ctx context.Context, paging common.Paging,
		filter *resourcemodel.ListResourcesRequest, moreKeys []string) ([]resourcemodel.Resources, int64, error)
}
type listResourcesBiz struct {
	store ListResourcesStore
}

func NewListResourcesBiz(store ListResourcesStore) *listResourcesBiz {
	return &listResourcesBiz{
		store: store,
	}
}
func (biz *listResourcesBiz) List(ctx context.Context, paging common.Paging,
	filter *resourcemodel.ListResourcesRequest, moreKeys []string) ([]resourcemodel.Resources, int64, error) {
	resources, total, err := biz.store.List(ctx, paging, filter, moreKeys)
	if err != nil {
		return nil, 0, err
	}
	return resources, total, nil
}
