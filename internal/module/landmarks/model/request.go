package landmarkmodel

type ListLandmarksRequest struct {
	Search string `form:"search" json:"search"`
}
type CreateLandmarkRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	UrlImage    string `json:"urlImage"`
}
type UpdateLandmarkRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	UrlImage    string `json:"urlImage"`
}
