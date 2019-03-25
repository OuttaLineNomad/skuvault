package sales

// ReleaseHeldQuantities is a automatically generated struct from json provided by sku vault's api docs.
type ReleaseHeldQuantities struct {
	SkusToRelease struct {
		String int64 `json:"String"`
	} `json:"SkusToRelease"`
}

// ReleaseHeldQuantitiesResponse is a automatically generated struct from json provided by sku vault's api docs.
type ReleaseHeldQuantitiesResponse struct {
	ReleasedQuantities struct {
		SKUAsKey int64 `json:"SKU_as_key"`
	} `json:"ReleasedQuantities"`
}
