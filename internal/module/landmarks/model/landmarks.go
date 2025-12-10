package landmarkmodel

import "example.com/m/internal/common"

type Landmark struct {
	common.MySqlModel
	Name        string `json:"name" gorm:"column:name;not null"`
	Description string `json:"description" gorm:"column:description"`
	UrlImage    string `json:"urlImage" gorm:"column:url_image"`
}

func (Landmark) TableName() string {
	return "landmarks"
}

// func (r *Landmark) Mark() {
// 	r.MySqlModel.GenUID(common.DBTypeLandmark)
// }
