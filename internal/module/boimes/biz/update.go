package boimebiz

import (
	"context"

	boimesmodel "example.com/m/internal/module/boimes/model"
)

type BoimesStore interface {
	Update(ctx context.Context, id uint, data interface{}, resourceIds []int) error
	FindById(ctx context.Context, id int) (*boimesmodel.Biome, error)
}

type BiomeBiz struct {
	store BoimesStore
}

func NewBiomeBiz(store BoimesStore) *BiomeBiz {
	return &BiomeBiz{store: store}
}

func (b *BiomeBiz) UpdateBiome(ctx context.Context, id int, req boimesmodel.UpdateBiomeRequest, resourceIds []int) error {
	updates := map[string]interface{}{}
	biome, err := b.store.FindById(ctx, id)
	if err != nil {
		return err
	}
	if req.Name != "" {

	}
}
