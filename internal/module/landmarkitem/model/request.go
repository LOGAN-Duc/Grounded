package landmarkitemmodel

type LandmarkAddRequest struct {
	ItemIds     []int `json:"item_ids" form:"item_ids"`
	ResourceIds []int `json:"resource_ids" form:"resource_ids"`
}
type LandmarkUpdateRequest struct {
	ItemIds     []int `json:"item_ids" form:"item_ids"`
	ResourceIds []int `json:"resource_ids" form:"resource_ids"`
}
type LandmarkDeleteRequest struct {
	ID int `json:"id" form:"id" binding:"required"`
}
type LandmarkListRequest struct {
	SearchText string `json:"search_text" form:"search_text"`
	LandMarkID int    `json:"landmark_id" form:"landmark_id"`
}
