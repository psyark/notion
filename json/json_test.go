package json

import (
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestJSON(t *testing.T) {
	type TestA struct {
		Nil []int `json:"nil"`
	}
	type TestB struct {
		Nil  []int `json:"nil,omitempty"`
		Len0 []int `json:"len0,omitempty"`
	}

	assert.Equal(t, `{"nil":[]}`, marshal(TestA{nil}))                      // - スライス型の nil は [] にエンコードする
	assert.Equal(t, `{"len0":[]}`, marshal(TestB{Nil: nil, Len0: []int{}})) // - スライス型の nil は empty として扱う一方、長さ0のスライスは empty として扱わない
	marshal(TestA{[]int{1, 2, 3}})                                          // 長さのある配列も正しくエンコードできる
}

func marshal(data any) string {
	return string(lo.Must(Marshal(data)))
}
