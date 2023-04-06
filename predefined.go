package notion

import "github.com/google/uuid"

type URLReference struct {
	URL string `json:"url"`
}

type PageReference struct {
	Id uuid.UUID `json:"id"`
}
