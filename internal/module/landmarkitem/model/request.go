package landmarkitemmodel

type LandmarkAddRequest struct {
	ItemIds     []int `json:"item_ids" form:"item_ids"`
	ResourceIds []int `json:"resource_ids" form:"resource_ids"`
}
