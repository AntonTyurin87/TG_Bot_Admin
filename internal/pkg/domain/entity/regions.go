package entity

// GetAllRegionsResponse ...
type GetAllRegionsResponse struct {
	Regions []*Region
}

// GetRegions ...
func (r *GetAllRegionsResponse) GetRegions() []*Region {
	if r == nil {
		return nil
	}
	return r.Regions
}

// Region ...
type Region struct {
	ID          int32  `json:"id"`
	NameRu      string `json:"name_ru"`
	Description string `json:"description"`
}

// GetID ...
func (r *Region) GetID() int32 {
	if r == nil {
		return 0
	}
	return r.ID
}

// GetNameRu ...
func (r *Region) GetNameRu() string {
	if r == nil {
		return ""
	}
	return r.NameRu
}

// GetDescription ...
func (r *Region) GetDescription() string {
	if r == nil {
		return ""
	}
	return r.Description
}
