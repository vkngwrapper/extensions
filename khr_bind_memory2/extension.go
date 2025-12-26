package khr_bind_memory2

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
	"github.com/vkngwrapper/extensions/v3/khr_bind_memory2/loader"
)

// VulkanExtension is an implementation of the Extension interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanExtension struct {
	driver khr_bind_memory2_loader.Loader
}

// CreateExtensionFromDevice produces an Extension object from a Device with
// khr_bind_memory2 loaded
func CreateExtensionFromDevice(device core.Device) *VulkanExtension {
	if !device.IsDeviceExtensionActive(ExtensionName) {
		return nil
	}
	return CreateExtensionFromDriver(khr_bind_memory2_loader.CreateLoaderFromCore(device.Driver()))
}

// CreateExtensionFromDriver generates an Extension from a loader.Loader object- this is usually
// used in tests to build an Extension from mock drivers
func CreateExtensionFromDriver(driver khr_bind_memory2_loader.Loader) *VulkanExtension {
	return &VulkanExtension{
		driver: driver,
	}
}

func (e *VulkanExtension) BindBufferMemory2(device core.Device, options []BindBufferMemoryInfo) (common.VkResult, error) {
	if device.Handle() == 0 {
		panic("device cannot be uninitialized")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	optionPtr, err := common.AllocOptionSlice[C.VkBindBufferMemoryInfoKHR, BindBufferMemoryInfo](arena, options)
	if err != nil {
		return core1_0.VKErrorUnknown, err
	}

	return e.driver.VkBindBufferMemory2KHR(device.Handle(), loader.Uint32(len(options)), (*khr_bind_memory2_loader.VkBindBufferMemoryInfoKHR)(unsafe.Pointer(optionPtr)))
}

func (e *VulkanExtension) BindImageMemory2(device core.Device, options []BindImageMemoryInfo) (common.VkResult, error) {
	if device.Handle() == 0 {
		panic("device cannot be uninitialized")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	optionPtr, err := common.AllocOptionSlice[C.VkBindImageMemoryInfoKHR, BindImageMemoryInfo](arena, options)
	if err != nil {
		return core1_0.VKErrorUnknown, err
	}

	return e.driver.VkBindImageMemory2KHR(device.Handle(), loader.Uint32(len(options)), (*khr_bind_memory2_loader.VkBindImageMemoryInfoKHR)(unsafe.Pointer(optionPtr)))
}

var _ Extension = &VulkanExtension{}
