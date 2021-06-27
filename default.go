package helium

import (
	"github.com/jenchik/helium/module"
	"github.com/jenchik/helium/service"
)

// DefaultApp defines default helium application and provides service.Module.
// nolint:gochecknoglobals
var DefaultApp = module.New(newDefaultApp).Append(service.Module)

func newDefaultApp(svc service.Group) App { return svc }
