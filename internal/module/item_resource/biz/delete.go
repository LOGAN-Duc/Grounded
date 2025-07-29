package itemresourcebiz

import (
	"context"
)

type DeleteItemResourceStore interface {
	Delete(ctx context.Context) error
}

type deleteItemResourceBiz struct {
	store DeleteItemResourceStore
}

func NewDeleteItemResourceBiz(store DeleteItemResourceStore) *deleteItemResourceBiz {
	return &deleteItemResourceBiz{
		store: store,
	}
}
func (biz *deleteItemResourceBiz) Delete(ctx context.Context) error {

	return biz.store.Delete(ctx)
}
