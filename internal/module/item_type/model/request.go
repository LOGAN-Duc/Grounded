package itemtypemodel

type CreateItemTypeRequest struct {
	Name string `json:"name"`
}
type ListItemTypeRequest struct {
	Search string `json:"search"`
}
