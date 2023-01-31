package repository

import "github.com/google/wire"

var WireSet = wire.NewSet(
	NewUserRepositoryImpl,
	NewChatGroupRepositoryImpl,
	NewChatUserRepositoryImpl,
	NewChatRepositoryImpl,
)
