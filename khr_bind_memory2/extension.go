package khr_bind_memory2

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/core1_0"
	"github.com/vkngwrapper/core/driver"
	"github.com/vkngwrapper/extensions/khr_bind_memory2/driver"
	"unsafe"
)

// VulkanExtension is an implementation of the Extension interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanExtension struct {
	driver khr_bind_memory2_driver.Driver
}

// CreateExtensionFromDevice produces an Extension object from a Device with
// khr_bind_memory2 loaded
func CreateExtensionFromDevice(device core1_0.Device) *VulkanExtension {
	if !device.IsDeviceExtensionActive(ExtensionName) {
		return nil
	}
	return CreateExtensionFromDriver(khr_bind_memory2_driver.CreateDriverFromCore(device.Driver()))
}

// CreateExtensionFromDriver generates an Extension from a driver.Driver object- this is usually
// used in tests to build an Extension from mock drivers
func CreateExtensionFromDriver(driver khr_bind_memory2_driver.Driver) *VulkanExtension {
	return &VulkanExtension{
		driver: driver,
	}
}

func (e *VulkanExtension) BindBufferMemory2(device core1_0.Device, options []BindBufferMemoryInfo) (common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	optionPtr, err := common.AllocOptionSlice[C.VkBindBufferMemoryInfoKHR, BindBufferMemoryInfo](arena, options)
	if err != nil {
		return core1_0.VKErrorUnknown, err
	}

	return e.driver.VkBindBufferMemory2KHR(device.Handle(), driver.Uint32(len(options)), (*khr_bind_memory2_driver.VkBindBufferMemoryInfoKHR)(unsafe.Pointer(optionPtr)))
}

func (e *VulkanExtension) BindImageMemory2(device core1_0.Device, options []BindImageMemoryInfo) (common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	optionPtr, err := common.AllocOptionSlice[C.VkBindImageMemoryInfoKHR, BindImageMemoryInfo](arena, options)
	if err != nil {
		return core1_0.VKErrorUnknown, err
	}

	return e.driver.VkBindImageMemory2KHR(device.Handle(), driver.Uint32(len(options)), (*khr_bind_memory2_driver.VkBindImageMemoryInfoKHR)(unsafe.Pointer(optionPtr)))
}

var _ Extension = &VulkanExtension{}
