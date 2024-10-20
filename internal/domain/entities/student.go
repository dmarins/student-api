package entities

import "github.com/dmarins/student-api/internal/infrastructure/uuid"

type Student struct {
	ID   string
	Name string
}

func NewStudent(name string) *Student {
	return &Student{
		ID:   uuid.NewId(),
		Name: name,
	}
}
