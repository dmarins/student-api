package dtos

type StudentInput struct {
	Name string `json:"name" validate:"required,max=200"`
}

type StudentOutput struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}