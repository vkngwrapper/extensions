package khr_device_group

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v2/common"
	"unsafe"
)

// DeviceGroupCommandBufferBeginInfo sets the initial device mask for a CommandBuffer
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupCommandBufferBeginInfo.html
type DeviceGroupCommandBufferBeginInfo struct {
	// DeviceMask is the initial value of the CommandBuffer object's device mask
	DeviceMask uint32

	common.NextOptions
}

func (o DeviceGroupCommandBufferBeginInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkDeviceGroupCommandBufferBeginInfoKHR{})))
	}

	createInfo := (*C.VkDeviceGroupCommandBufferBeginInfoKHR)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_DEVICE_GROUP_COMMAND_BUFFER_BEGIN_INFO_KHR
	createInfo.pNext = next
	createInfo.deviceMask = C.uint32_t(o.DeviceMask)

	return preallocatedPointer, nil
}
