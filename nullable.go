package notion

import "encoding/json"

// Nullable は JSONでomitemptyを指定しつつ、非empty値としてnullの存在を許す型です
type Nullable[T any] []struct {
	isNull bool
	value  T
}

func (n Nullable[T]) IsEmpty() bool {
	return len(n) == 0
}

func (n *Nullable[T]) SetEmpty() {
	*n = (*n)[0:0]
}

func (n Nullable[T]) IsNull() bool {
	return !n.IsEmpty() && (n)[0].isNull
}

func (n *Nullable[T]) SetNull() {
	if n.IsEmpty() {
		*n = append(*n, struct {
			isNull bool
			value  T
		}{})
	}
	(*n)[0].isNull = true
}

func (n Nullable[T]) Value() T {
	if !n.IsEmpty() {
		return (n)[0].value
	}
	var value T
	return value
}

func (n *Nullable[T]) SetValue(value T) {
	if n.IsEmpty() {
		*n = append(*n, struct {
			isNull bool
			value  T
		}{})
	}
	(*n)[0].isNull = false
	(*n)[0].value = value
}

func (n *Nullable[T]) UnmarshalJSON(data []byte) error {
	if n.IsEmpty() {
		*n = append(*n, struct {
			isNull bool
			value  T
		}{})
	}

	if string(data) == "null" {
		(*n)[0].isNull = true
	} else {
		(*n)[0].isNull = false
		return json.Unmarshal(data, &(*n)[0].value)
	}

	return nil
}

func (n Nullable[T]) MarshalJSON() ([]byte, error) {
	if n[0].isNull {
		return []byte("null"), nil
	}
	return json.Marshal(n[0].value)
}
