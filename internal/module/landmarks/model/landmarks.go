package landmarkmodel

import (
	"example.com/m/internal/common"
	landmarkitemmodel "example.com/m/internal/module/landmarkitem/model"
)

type Landmark struct {
	common.MySqlModel
	Name          string                            `json:"name" gorm:"column:name;not null"`
	Description   string                            `json:"description" gorm:"column:description"`
	UrlImage      string                            `json:"urlImage" gorm:"column:url_image"`
	LandmarkItems []*landmarkitemmodel.LandMarkItem `json:"landmark_items" gorm:"-"`
}

func (Landmark) TableName() string {
	return "landmarks"
}
