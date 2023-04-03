package doc2api

type docElement interface {
	equals(docElement) bool
}

type headingElement string

func (e headingElement) equals(e2 docElement) bool {
	if e2, ok := e2.(headingElement); ok {
		return string(e) == string(e2)
	}
	return false
}

type codeElement string

func (e codeElement) equals(e2 docElement) bool {
	if e2, ok := e2.(codeElement); ok {
		return string(e) == string(e2)
	}
	return false
}
