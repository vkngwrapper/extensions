package khr_device_group_creation

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/pkg/errors"
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/core1_0"
	"unsafe"
)

// DeviceGroupDeviceCreateInfo creates a logical Device from multiple PhysicalDevice objects
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupDeviceCreateInfo.html
type DeviceGroupDeviceCreateInfo struct {
	// PhysicalDevices is a slice of PhysicalDevice objects belonging to the same Device group
	PhysicalDevices []core1_0.PhysicalDevice

	common.NextOptions
}

func (o DeviceGroupDeviceCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkDeviceGroupDeviceCreateInfoKHR{})))
	}

	if len(o.PhysicalDevices) < 1 {
		return nil, errors.New("must include at least one physical device in DeviceGroupDeviceCreateInfo")
	}

	createInfo := (*C.VkDeviceGroupDeviceCreateInfoKHR)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_DEVICE_GROUP_DEVICE_CREATE_INFO_KHR
	createInfo.pNext = next

	count := len(o.PhysicalDevices)
	createInfo.physicalDeviceCount = C.uint32_t(count)
	physicalDevicesPtr := (*C.VkPhysicalDevice)(allocator.Malloc(count * int(unsafe.Sizeof([1]C.VkPhysicalDevice{}))))
	physicalDevicesSlice := ([]C.VkPhysicalDevice)(unsafe.Slice(physicalDevicesPtr, count))

	for i := 0; i < count; i++ {
		if o.PhysicalDevices[i] == nil {
			return nil, errors.Errorf("khr_device_group_creation.DeviceGroupDeviceCreateInfo.PhysicalDevices cannot "+
				"have nil elements, but element %d is nil", i)
		}
		physicalDevicesSlice[i] = C.VkPhysicalDevice(unsafe.Pointer(o.PhysicalDevices[i].Handle()))
	}
	createInfo.pPhysicalDevices = physicalDevicesPtr
	return preallocatedPointer, nil
}
