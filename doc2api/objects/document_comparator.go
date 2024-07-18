package objects

import (
	"fmt"
	"reflect"
)

// DocumentComparator は、クライアント実装を最新に保つために、オンラインのドキュメントの更新を検知するための機能を提供します。
type DocumentComparator struct {
	index    int
	elements []DocumentElement
	builder  *CodeBuilder
}

func (c *DocumentComparator) ExpectBlock(block *Block) *Match[*Block] {
	if reflect.DeepEqual(c.elements[c.index], block) {
		c.index++
		return &Match[*Block]{element: block, builder: c.builder}
	}

	panic(fmt.Errorf("mismatch:\nwant: %#v\ngot:  %#v", block, c.elements[c.index]))
}

func (c *DocumentComparator) ExpectParameter(parameter *Parameter) *Match[*Parameter] {
	if reflect.DeepEqual(c.elements[c.index], parameter) {
		c.index++
		return &Match[*Parameter]{element: parameter, builder: c.builder}
	}

	panic(fmt.Errorf("mismatch:\nwant: %#v\ngot:  %#v", parameter, c.elements[c.index]))
}

// RequestBuilderForUndocumented は、ドキュメントに書かれていないことをコードに反映するために CodeBuilderを提供します
func (c *DocumentComparator) RequestBuilderForUndocumented(fn func(b *CodeBuilder)) {
	fn(c.builder)
}

// finish は比較を終了します
func (c *DocumentComparator) finish() {
	if c.index < len(c.elements) {
		next := c.elements[c.index]
		panic(fmt.Sprintf("比較されていないエレメントが存在します: %+v", next))
	}
}

type Match[T DocumentElement] struct {
	element T
	builder *CodeBuilder
	// globalBuilder *CodeBuilder
}

func (e *Match[T]) Output(fn func(e T, b *CodeBuilder)) *Match[T] {
	fn(e.element, e.builder)
	return e
}

// func (e *Match[T]) OutputGlobal(fn func(e T, b *CodeBuilder)) *Match[T] {
// 	fn(e.element, e.globalBuilder)
// 	return e
// }
