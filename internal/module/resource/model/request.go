package resourcemodel

type CreateResourcesRequest struct {
	Name           string `json:"name" binding:"required"`
	Code           string `json:"code" binding:"required"`
	ResourceTypeId int    `json:"resourceTypeId" binding:"required"`
}
type UpdateResourcesRequest struct {
	Name           string `json:"name"`
	Code           string `json:"code"`
	ResourceTypeId *int   `json:"resourceTypeId"`
}

type ListResourcesRequest struct {
	Search string `json:"search" form:"search"`
}
