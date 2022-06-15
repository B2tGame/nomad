// +build !noop

package catalog

import (
	"github.com/hashicorp/nomad/devices/gpu/noop"
)

// This file is where all builtin plugins should be registered in the catalog.
// Plugins with build restrictions should be placed in the appropriate
// register_XXX.go file.
func init() {
	Register(noop.PluginID, noop.PluginConfig)
}
