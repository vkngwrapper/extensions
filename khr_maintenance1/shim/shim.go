package khr_maintenance1_shim

import (
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/core1_1"
	"github.com/vkngwrapper/extensions/v3/khr_maintenance1"
)

//go:generate mockgen -source shim.go -destination ../mocks/shim.go -package mock_maintenance1

// Shim contains all commands for the khr_maintenance1 extension
type Shim interface {
	// TrimCommandPool trims a CommandPool
	//
	// flags - Reserved for future use
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkTrimCommandPool.html
	TrimCommandPool(flags core1_1.CommandPoolTrimFlags)
}

type VulkanShim struct {
	extension   khr_maintenance1.Extension
	commandPool core1_0.CommandPool
}

// Compiler check that VulkanShim satisfies Shim
var _ Shim = &VulkanShim{}

func NewShim(extension khr_maintenance1.Extension, pool core1_0.CommandPool) *VulkanShim {
	if extension == nil {
		panic("extension cannot be nil")
	}
	if pool == nil {
		panic("pool cannot be nil")
	}
	return &VulkanShim{
		extension:   extension,
		commandPool: pool,
	}
}

func (s *VulkanShim) TrimCommandPool(flags core1_1.CommandPoolTrimFlags) {
	s.extension.TrimCommandPool(s.commandPool, khr_maintenance1.CommandPoolTrimFlags(flags))
}
