package khr_bind_memory2

import (
	"github.com/vkngwrapper/core/v3/common"
)

//go:generate mockgen -source extiface.go -destination ./mocks/extension.go -package mock_bind_memory2

// ExtensionDriver contains all the commands for the khr_bind_memory2 extension
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_bind_memory2.html
type ExtensionDriver interface {
	// BindBufferMemory2 binds DeviceMemory to Buffer objects
	//
	// device - The core1_0.Device which owns the core1_0.DeviceMemory and core1_0.Buffer
	//
	// options - A slice of BindBufferMemoryInfo structures describing Buffer objects and memory to bind
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkBindBufferMemory2.html
	BindBufferMemory2(options []BindBufferMemoryInfo) (common.VkResult, error)
	// BindImageMemory2 binds DeviceMemory to Image objects
	//
	// device - The core1_0.Device which owns the core1_0.DeviceMemory and core1_0.Buffer
	//
	// options - A slice of BindImageMemoryInfo structures describing Image objects and memory to bind
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkBindImageMemory2.html
	BindImageMemory2(options []BindImageMemoryInfo) (common.VkResult, error)
}
