package objectdoc

import "github.com/dave/jennifer/jen"

var undocumentedRequestID = &field{
	name:      "request_id",
	typeCode:  jen.String(),
	comment:   "UNDOCUMENTED",
	omitEmpty: true,
}
