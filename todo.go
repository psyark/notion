package notion

import (
	"encoding/json"

	"github.com/samber/lo"
)

// TODO
type Comment struct{}

// TODO 生成
type SearchFilter struct {
	Value    string `json:"value"`    // The value of the property to filter the results by. Possible values for object type include page or database. Limitation: Currently the only filter allowed is object which will filter by type of object (either page or database)
	Property string `json:"property"` // The name of the property to filter by. Currently the only property you can filter by is the object type. Possible values include object. Limitation: Currently the only filter allowed is object which will filter by type of object (either page or database)
}

func (p *Pagination[T]) UnmarshalJSON(data []byte) error {
	type Alias Pagination[T]

	var zero T
	if string(lo.Must(json.Marshal(zero))) != "null" {
		return json.Unmarshal(data, (*Alias)(p))
	}

	t := struct {
		*Alias
		Results []pageOrDatabaseUnmarshaler `json:"results"`
	}{
		(*Alias)(p),
		[]pageOrDatabaseUnmarshaler{},
	}
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	p.Results = lo.Map(t.Results, func(u pageOrDatabaseUnmarshaler, index int) T { return u.value.(T) })
	return nil
}
