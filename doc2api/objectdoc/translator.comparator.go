package objectdoc

import "fmt"

type comparator struct {
	index    int
	elements []docElement
}

func (c *comparator) nextMustBlock(element blockElement, callback func(e blockElement)) {
	if c.index >= len(c.elements) {
		panic("no more elements")
	}

	elem := c.elements[c.index]
	if block, ok := elem.(*blockElement); ok {
		if *block != element {
			panic(fmt.Sprintf("mismatch:\nremote: %#v\nlocal:  %#v", *block, element))
		}
	} else {
		panic(fmt.Sprintf("the element is not block but %T", elem))
	}

	callback(element)
	c.index++
}

func (c *comparator) nextMustParameter(element parameterElement, callback func(e parameterElement)) {
	if c.index >= len(c.elements) {
		panic("no more elements")
	}

	elem := c.elements[c.index]
	if param, ok := elem.(*parameterElement); ok {
		if *param != element {
			panic(fmt.Sprintf("mismatch:\nremote: %#v\nlocal:  %#v", *param, element))
		}
	} else {
		panic(fmt.Sprintf("the element is not param but %T", elem))
	}

	callback(element)
	c.index++
}
