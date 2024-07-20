package notion

import "encoding/json"

// Nullable は empty | null | 値あり の三種類の状態を取りうるデータ型で、
// omitempty JSONタグと併せて使います。
// empty -> 出力はスキップされます
// null -> null が出力されます
// 値あり -> その値が出力されます
type Nullable[T any] []struct {
	isNull bool
	value  T
}

func (n Nullable[T]) IsEmpty() bool {
	return len(n) == 0
}

func (n *Nullable[T]) SetEmpty() {
	*n = Nullable[T]{}
}

func (n Nullable[T]) IsNull() bool {
	return !n.IsEmpty() && (n)[0].isNull
}

func (n *Nullable[T]) SetNull() {
	*n = Nullable[T]{{isNull: true}}
}

func (n Nullable[T]) Value() (value T) {
	if !n.IsEmpty() {
		return (n)[0].value
	}
	return
}

func (n *Nullable[T]) SetValue(value T) {
	if n.IsEmpty() {
		*n = Nullable[T]{{value: value}}
	}
}

func (n *Nullable[T]) UnmarshalJSON(data []byte) error {
	*n = Nullable[T]{{}}
	if string(data) == "null" {
		(*n)[0].isNull = true
	} else {
		(*n)[0].isNull = false
		return json.Unmarshal(data, &(*n)[0].value)
	}
	return nil
}

// TODO omitemptyを指定しなかった場合のエラーを抑制
func (n Nullable[T]) MarshalJSON() ([]byte, error) {
	if n[0].isNull {
		return []byte("null"), nil
	}
	return json.Marshal(n[0].value)
}
