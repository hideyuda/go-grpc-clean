package utility

import (
	"fmt"

	"github.com/google/uuid"
)

// UUIDの生成
// 例：c5db1800-ce4c-11de-9ce6-97579ef38e45
func CreateUUID() (uuidObj uuid.UUID) {
	uuidObj, err := uuid.NewUUID()
	if err != nil {
		fmt.Println(err)
	}

	return uuidObj
}
