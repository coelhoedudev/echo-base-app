package util

import "github.com/google/uuid"

type UUID struct{}

func (UUID) Create() string {
	result, _ := uuid.NewV7()
	return result.String()
}
