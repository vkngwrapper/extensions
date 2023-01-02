package khr_bind_memory2_shim

import (
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/core1_0"
	"github.com/vkngwrapper/core/v2/core1_1"
	"github.com/vkngwrapper/extensions/v2/khr_bind_memory2"
)

//go:generate mockgen -source shim.go -destination ../mocks/shim.go -package mock_bind_memory2

type Shim interface {
	// BindBufferMemory2 binds DeviceMemory to Buffer objects
	//
	// options - A slice of BindBufferMemoryInfo structures describing Buffer objects and memory to bind
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkBindBufferMemory2.html
	BindBufferMemory2(options []core1_1.BindBufferMemoryInfo) (common.VkResult, error)
	// BindImageMemory2 binds DeviceMemory to Image objects
	//
	// options - A slice of BindImageMemoryInfo structures describing Image objects and memory to bind
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkBindImageMemory2.html
	BindImageMemory2(options []core1_1.BindImageMemoryInfo) (common.VkResult, error)
}

type VulkanShim struct {
	device    core1_0.Device
	extension khr_bind_memory2.Extension
}

// Compiler check that VulkanShim satisfies Shim
var _ Shim = &VulkanShim{}

func NewShim(device core1_0.Device, extension khr_bind_memory2.Extension) *VulkanShim {
	if device == nil {
		panic("device cannot be nil")
	}
	if extension == nil {
		panic("extension cannot be nil")
	}
	return &VulkanShim{
		extension: extension,
		device:    device,
	}
}

func (s *VulkanShim) BindBufferMemory2(options []core1_1.BindBufferMemoryInfo) (common.VkResult, error) {
	inOptions := make([]khr_bind_memory2.BindBufferMemoryInfo, 0, len(options))
	for _, option := range options {
		inOptions = append(inOptions, (khr_bind_memory2.BindBufferMemoryInfo)(option))
	}

	return s.extension.BindBufferMemory2(s.device, inOptions)
}

func (s *VulkanShim) BindImageMemory2(options []core1_1.BindImageMemoryInfo) (common.VkResult, error) {
	inOptions := make([]khr_bind_memory2.BindImageMemoryInfo, 0, len(options))
	for _, option := range options {
		inOptions = append(inOptions, (khr_bind_memory2.BindImageMemoryInfo)(option))
	}

	return s.extension.BindImageMemory2(s.device, inOptions)
}
