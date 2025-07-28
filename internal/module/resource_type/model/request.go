package resourcetypemodel

type CreateResourceTypeRequest struct {
	Name string `json:"name"`
}
type ListResourceTypeRequest struct {
	Search string `json:"search"`
}
