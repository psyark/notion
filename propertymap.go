package notion

type PropertyValueMap map[string]PropertyValue

func (m PropertyValueMap) Get(id string) *PropertyValue {
	for _, v := range m {
		if v.Id == id {
			return &v
		}
	}
	return nil
}
