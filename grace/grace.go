package grace

import (
	"github.com/jenchik/helium/module"
)

// Module graceful context.
// nolint:gochecknoglobals
var Module = module.Module{
	{Constructor: NewGracefulContext},
}
