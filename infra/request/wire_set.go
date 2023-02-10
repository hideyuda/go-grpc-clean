package request

import "github.com/google/wire"

var WireSet = wire.NewSet(
	NewUserRequestImpl,
)
