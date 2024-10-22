package uuid

import "github.com/google/uuid"

func NewId() string {
	return uuid.New().String()
}

func IsValid(id string) bool {
	_, err := uuid.Parse(id)
	return err == nil
}
