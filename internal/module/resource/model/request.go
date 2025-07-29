package resourcemodel

type CreateResourcesRequest struct {
	Name           string `json:"name"`
	Code           string `json:"code"`
	ResourceTypeId int    `json:"resourceTypeId"`
}
type UpdateResourcesRequest struct {
	Name           string `json:"name"`
	Code           string `json:"code"`
	ResourceTypeId *int   `json:"resourceTypeId"`
}

type ListResourcesRequest struct {
	Search string `json:"search" form:"search"`
}
