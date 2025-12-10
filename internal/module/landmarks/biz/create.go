package landmarkbiz

import (
	"context"
	"errors"
	"mime/multipart"

	landmarkmodel "example.com/m/internal/module/landmarks/model"
)

type Landmark interface {
	Find(ctx context.Context, id int64, moreKeys []string) (*landmarkmodel.Landmark, error)
	FindByName(ctx context.Context, name string, moreKeys []string) (*landmarkmodel.Landmark, error)
	CreateLandmark(ctx context.Context, landmark *landmarkmodel.Landmark) error
}
type FileUploader interface {
	Upload(file *multipart.FileHeader) (string, error)
}
type landmarkBiz struct {
	landmarkStore Landmark
	fileUploader  FileUploader
}

func NewLandmarkBiz(landmarkStore Landmark, fileUploader FileUploader) *landmarkBiz {
	return &landmarkBiz{landmarkStore: landmarkStore, fileUploader: fileUploader}
}
func (biz *landmarkBiz) CreateLandmark(ctx context.Context, landmark *landmarkmodel.CreateLandmarkRequest, file *multipart.FileHeader) error {
	_, err := biz.landmarkStore.FindByName(ctx, landmark.Name, nil)
	if err == nil {
		return errors.New("landmark already exists")
	}
	if file != nil {
		url, err := biz.fileUploader.Upload(file)
		if err != nil {
			return err
		}
		landmark.UrlImage = url
	}
	landmarkModel := &landmarkmodel.Landmark{
		Name:        landmark.Name,
		Description: landmark.Description,
		UrlImage:    landmark.UrlImage,
	}
	return biz.landmarkStore.CreateLandmark(ctx, landmarkModel)

}
