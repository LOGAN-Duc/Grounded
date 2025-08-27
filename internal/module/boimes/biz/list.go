package boimebiz

import (
	"context"

	"example.com/m/internal/common"
	boimesmodel "example.com/m/internal/module/boimes/model"
)

type ListBoimeStore interface {
	List(ctx context.Context, paging common.Paging,
		filter *boimesmodel.FilterBiomeRequest, moreKeys []string) ([]boimesmodel.Biome, int64, error)
}
type listBoimeBiz struct {
	store ListBoimeStore
}

func NewListBoimeBiz(store ListBoimeStore) *listBoimeBiz {
	return &listBoimeBiz{
		store: store,
	}
}
func (biz *listBoimeBiz) List(ctx context.Context, paging common.Paging,
	filter *boimesmodel.FilterBiomeRequest, moreKeys []string) ([]boimesmodel.Biome, int64, error) {
	items, total, err := biz.store.List(ctx, paging, filter, moreKeys)
	if err != nil {
		return nil, 0, err
	}
	return items, total, nil
}
