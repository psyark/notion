package notion

import "fmt"

type PropertyValueMap map[string]PropertyValue

func (m PropertyValueMap) Get(id string) (*PropertyValue, error) {
	for _, v := range m {
		if v.Id == id {
			return &v, nil
		}
	}
	return nil, fmt.Errorf("not exists")
}
