package injector

import (
	"github.com/google/wire"
	"github.com/labstack/echo/v5"
)

// Set app
var Set = wire.NewSet(wire.Struct(new(Injector), "*"))

// Injector app
type Injector struct {
	App *echo.Echo
}
