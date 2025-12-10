package resourcemodel

type CreateResourcesRequest struct {
	Name           string `json:"name" form:"name"`
	Code           string `json:"code" form:"code"`
	ResourceTypeId int    `json:"resourceTypeId" form:"resourceTypeId"`
	UrlImage       string `json:"urlImage" form:"urlImage"`
}
type UpdateResourcesRequest struct {
	Name           string `json:"name" form:"name"`
	Code           string `json:"code" form:"code"`
	ResourceTypeId *int   `json:"resourceTypeId" form:"resourceTypeId"`
}

type ListResourcesRequest struct {
	Search string `json:"search" form:"search"`
}
