package base

import "github.com/google/uuid"

func GetID() string {
	return uuid.New().String()
}
