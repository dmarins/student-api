package dtos

type (
	StudentCreateInput struct {
		Name string `json:"name" validate:"required,max=200"`
	}

	StudentUpdateInput struct {
		ID   string `json:"id" validate:"required,max=36"`
		Name string `json:"name" validate:"required,max=200"`
	}

	StudentOutput struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}
)
