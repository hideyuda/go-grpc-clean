package driver

import "github.com/google/wire"

var WireSet = wire.NewSet(
	NewFirebaseImpl,
)
