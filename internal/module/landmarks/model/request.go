package landmarkmodel

type ListLandmarksRequest struct {
	Search string `form:"search" json:"search"`
}
type CreateLandmarkRequest struct {
	Name        string `json:"name" binding:"required" form:"name"`
	Description string `json:"description" form:"description"`
	UrlImage    string `json:"urlImage" form:"urlImage"`
}
type UpdateLandmarkRequest struct {
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
	UrlImage    string `json:"urlImage" form:"urlImage"`
}
