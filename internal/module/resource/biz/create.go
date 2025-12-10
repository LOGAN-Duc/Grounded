package resourcebiz

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"

	resourcemodel "example.com/m/internal/module/resource/model"
	resourcetypemodel "example.com/m/internal/module/resource_type/model"
)

type CreateResourcesStore interface {
	Create(ctx context.Context, resource *resourcemodel.Resources) error
	FindByName(ctx context.Context, nameRequest string) (*resourcemodel.Resources, error)
}
type FileUploader interface {
	Upload(file *multipart.FileHeader) (string, error)
}
type FindByIdResourceTypeStore interface {
	FindById(ctx context.Context, id int) (*resourcetypemodel.ResourceType, error)
}
type createResourcesBiz struct {
	store        CreateResourcesStore
	typerstore   FindByIdResourceTypeStore
	fileUploader FileUploader
}

func NewCreateResourcesBiz(store CreateResourcesStore, typerstore FindByIdResourceTypeStore, fileUploader FileUploader) *createResourcesBiz {
	return &createResourcesBiz{
		store:        store,
		typerstore:   typerstore,
		fileUploader: fileUploader,
	}
}
func (biz *createResourcesBiz) Create(ctx context.Context, req *resourcemodel.CreateResourcesRequest, file *multipart.FileHeader) error {
	if file != nil {
		url, err := biz.fileUploader.Upload(file)
		if err != nil {
			return err
		}
		req.UrlImage = url
	}
	_, err := biz.store.FindByName(ctx, req.Name)
	if err == nil {
		return errors.New("resource with this name already exists")
	}
	fmt.Println("ResourceTypeId", req.ResourceTypeId)
	resourceType, err := biz.typerstore.FindById(ctx, req.ResourceTypeId)
	if err != nil || resourceType == nil {
		return errors.New("resource type does not exist")
	}

	resource := &resourcemodel.Resources{
		Name:           req.Name,
		UrlImage:       req.UrlImage,
		Code:           req.Code,
		ResourceTypeId: int(resourceType.Id),
	}

	return biz.store.Create(ctx, resource)
}
