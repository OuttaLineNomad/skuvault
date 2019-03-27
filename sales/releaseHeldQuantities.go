package sales

// ReleaseHeldQuantities is a automatically generated struct from json provided by sku vault's api docs.
type ReleaseHeldQuantities struct {
	SkusToRelease map[string]int `json:"SkusToRelease"`
}

// ReleaseHeldQuantitiesResponse is a automatically generated struct from json provided by sku vault's api docs.
type ReleaseHeldQuantitiesResponse struct {
	ReleasedQuantities map[string]int `json:"ReleasedQuantities"`
}
