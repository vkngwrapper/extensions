package ext_host_query_reset

import (
	"github.com/vkngwrapper/core/v3/core1_0"
)

//go:generate mockgen -source extiface.go -destination ./mocks/extension.go -package mock_host_query_reset

// ExtensionDriver contains all the commands for the ext_host_query_reset extension
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_host_query_reset.html
type ExtensionDriver interface {
	// ResetQueryPool resets queries in the provided core1_0.QueryPool
	//
	// queryPool - the core1_0.QueryPool to reset
	//
	// firstQuery - The initial query index to reset
	//
	// queryCount - The number of queries to reset
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkResetQueryPool.html
	ResetQueryPool(queryPool core1_0.QueryPool, firstQuery, queryCount int)
}
