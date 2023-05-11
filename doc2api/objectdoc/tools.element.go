package objectdoc

type fieldOption func(f *field)

var omitEmpty fieldOption = func(f *field) {
	f.omitEmpty = true
}

var discriminatorNotEmpty fieldOption = func(f *field) {
	f.discriminatorNotEmpty = true
}
