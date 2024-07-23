package json

import (
	"reflect"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
	"github.com/modern-go/reflect2"
)

type customExtension struct {
	jsoniter.DummyExtension
}

func (*customExtension) DecorateEncoder(typ reflect2.Type, encoder jsoniter.ValEncoder) jsoniter.ValEncoder {
	if typ.Kind() == reflect.Slice {
		return &sliceEncoder{encoder}
	}
	return encoder
}

type sliceEncoder struct {
	originEncoder jsoniter.ValEncoder
}

func (encoder *sliceEncoder) IsEmpty(ptr unsafe.Pointer) bool {
	return ptr == nil || (*reflect.SliceHeader)(ptr).Data == 0
}

func (encoder *sliceEncoder) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	sliceHeader := (*reflect.SliceHeader)(ptr)
	if sliceHeader.Len == 0 {
		stream.WriteEmptyArray()
	} else {
		encoder.originEncoder.Encode(ptr, stream)
	}
}
