package ext_host_query_reset_shim

import (
	"github.com/vkngwrapper/core/v2/core1_0"
	"github.com/vkngwrapper/extensions/v2/ext_host_query_reset"
)

//go:generate mockgen -source shim.go -destination ../mocks/shim.go -package mock_host_query_reset

// Shim provides a bridge between ext_host_query_reset and the core 1.2 version, allowing code to handle
// both in a single interface
type Shim interface {
	// Reset resets queries in the provided core1_0.QueryPool
	//
	// firstQuery - The initial query index to reset
	//
	// queryCount - The number of queries to reset
	Reset(firstQuery, queryCount int)
}

// Compiler check that VulkanShim satisfies Shim
var _ Shim = &VulkanShim{}

type VulkanShim struct {
	extension ext_host_query_reset.Extension
	queryPool core1_0.QueryPool
}

func NewShim(extension ext_host_query_reset.Extension, queryPool core1_0.QueryPool) *VulkanShim {
	if extension == nil {
		panic("extension cannot be nil")
	}
	if queryPool == nil {
		panic("queryPool cannot be nil")
	}
	return &VulkanShim{
		extension: extension,
		queryPool: queryPool,
	}
}

func (s *VulkanShim) Reset(firstQuery, queryCount int) {
	s.extension.ResetQueryPool(s.queryPool, firstQuery, queryCount)
}
