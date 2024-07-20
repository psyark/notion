package notion

// TODO
type Comment struct{}

// TODO
type PageOrDatabase struct{}

// TODO 生成
type SearchFilter struct {
	Value    string `json:"value"`    // The value of the property to filter the results by. Possible values for object type include page or database. Limitation: Currently the only filter allowed is object which will filter by type of object (either page or database)
	Property string `json:"property"` // The name of the property to filter by. Currently the only property you can filter by is the object type. Possible values include object. Limitation: Currently the only filter allowed is object which will filter by type of object (either page or database)
}
