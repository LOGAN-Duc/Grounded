package itembiz

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"

	itemmodel "example.com/m/internal/module/item/model"
	itemtypemodel "example.com/m/internal/module/item_type/model"
)

type CreateItemStore interface {
	Create(ctx context.Context, resource *itemmodel.Item) error
	FindByName(ctx context.Context, nameRequest string) (*itemmodel.Item, error)
}
type ItemTypeFindByCode interface {
	FindById(ctx context.Context, id int) (*itemtypemodel.ItemType, error)
}
type FileUploader interface {
	Upload(file *multipart.FileHeader) (string, error)
}
type createItemBiz struct {
	store    CreateItemStore
	itemType ItemTypeFindByCode
	uploader FileUploader
}

func NewCreateItemBiz(store CreateItemStore, itemType ItemTypeFindByCode, uploader FileUploader) *createItemBiz {
	return &createItemBiz{
		store:    store,
		itemType: itemType,
		uploader: uploader,
	}
}
func (biz *createItemBiz) Create(ctx context.Context, req *itemmodel.CreateItemRequest, file *multipart.FileHeader) error {

	// Upload file tại biz
	if file != nil {
		url, err := biz.uploader.Upload(file)
		if err != nil {
			return err
		}
		req.UrlImage = url
	}
	_, err := biz.store.FindByName(ctx, req.Name)
	if err == nil {
		return errors.New("item with this name already exists")
	}
	fmt.Println("req.Itemtype", req.ItemTypeId)
	itemType, err := biz.itemType.FindById(ctx, req.ItemTypeId)
	if err != nil {
		return errors.New("itemType not found")
	}
	resource := &itemmodel.Item{
		Name:       req.Name,
		Code:       req.Code,
		UrlImage:   req.UrlImage,
		ItemTypeId: int(itemType.Id),
	}

	return biz.store.Create(ctx, resource)
}
