package repository_test

import (
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	helper.ClearAllTables()

	// var (
	// 	firebaseID     = "uid"
	// 	user           = entity.NewUser(firebaseID)
	// 	userRepository = repository.NewUserRepositoryImpl(dbm)
	// )

	// err := userRepository.Create(user)
	// if err != nil {
	// 	t.Fatalf("user should be created but got error: %v", err)
	// }

	// if user.ID == 0 {
	// 	t.Fatalf("user.ID should be not 0, but got 0")
	// }

	// user, err = userRepository.FindByID(user.ID)
	// if err != nil {
	// 	t.Fatal(err)
	// }
}
