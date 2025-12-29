package khr_bind_memory2

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
	"github.com/vkngwrapper/extensions/v3/khr_bind_memory2/loader"
)

// VulkanExtensionDriver is an implementation of the ExtensionDriver interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanExtensionDriver struct {
	driver khr_bind_memory2_loader.Loader
	device core1_0.Device
}

// CreateExtensionDriverFromCoreDriver produces an ExtensionDriver object from a Device with
// khr_bind_memory2 loaded
func CreateExtensionDriverFromCoreDriver(coreDriver core1_0.DeviceDriver) ExtensionDriver {
	device := coreDriver.Device()
	if !device.IsDeviceExtensionActive(ExtensionName) {
		return nil
	}
	return CreateExtensionDriverFromLoader(khr_bind_memory2_loader.CreateLoaderFromCore(coreDriver.Loader()), device)
}

// CreateExtensionDriverFromLoader generates an ExtensionDriver from a loader.Loader object- this is usually
// used in tests to build an ExtensionDriver from mock drivers
func CreateExtensionDriverFromLoader(driver khr_bind_memory2_loader.Loader, device core1_0.Device) *VulkanExtensionDriver {
	return &VulkanExtensionDriver{
		driver: driver,
		device: device,
	}
}

func (e *VulkanExtensionDriver) BindBufferMemory2(options ...BindBufferMemoryInfo) (common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	optionPtr, err := common.AllocOptionSlice[C.VkBindBufferMemoryInfoKHR, BindBufferMemoryInfo](arena, options)
	if err != nil {
		return core1_0.VKErrorUnknown, err
	}

	return e.driver.VkBindBufferMemory2KHR(e.device.Handle(), loader.Uint32(len(options)), (*khr_bind_memory2_loader.VkBindBufferMemoryInfoKHR)(unsafe.Pointer(optionPtr)))
}

func (e *VulkanExtensionDriver) BindImageMemory2(options ...BindImageMemoryInfo) (common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	optionPtr, err := common.AllocOptionSlice[C.VkBindImageMemoryInfoKHR, BindImageMemoryInfo](arena, options)
	if err != nil {
		return core1_0.VKErrorUnknown, err
	}

	return e.driver.VkBindImageMemory2KHR(e.device.Handle(), loader.Uint32(len(options)), (*khr_bind_memory2_loader.VkBindImageMemoryInfoKHR)(unsafe.Pointer(optionPtr)))
}

var _ ExtensionDriver = &VulkanExtensionDriver{}
