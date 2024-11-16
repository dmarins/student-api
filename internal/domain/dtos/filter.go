package dtos

const (
	FILTER_NAME = "name"
)

type Filter struct {
	Name *string `query:"name"`
}
