package doc2api

type builder struct{}

func (b *builder) addClass() *class {
	return &class{}
}

type class struct {
	name    string
	comment string
}
