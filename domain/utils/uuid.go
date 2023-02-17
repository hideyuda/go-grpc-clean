package utils

import (
	"fmt"

	"github.com/google/uuid"
)

// Uuidの生成
// 例：c5db1800-ce4c-11de-9ce6-97579ef38e45
func CreateUuid() (uuidObj uuid.UUID) {
	uuidObj, err := uuid.NewUUID()
	if err != nil {
		fmt.Println(err)
	}

	return uuidObj
}

func CreateUuidString() (uuidStr string) {
	uuidStr = CreateUuid().String()

	return uuidStr
}

func ParseUuid(uuidStr string) (uuidObj uuid.UUID, err error) {
	uuidObj, err = uuid.Parse(uuidStr)
	if err != nil {
		fmt.Println(err)
	}

	return uuidObj, err
}
