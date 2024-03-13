package id

import "github.com/google/uuid"

type ID = uuid.UUID

func NewID() ID {
	return uuid.New()
}

func StringToID(s string) (ID, error) {
	return uuid.Parse(s)
}

func IsNil(id ID) bool {
	return id == uuid.Nil
}
