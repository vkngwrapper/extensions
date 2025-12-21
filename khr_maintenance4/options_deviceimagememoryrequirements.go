package khr_maintenance4

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
)

// DeviceImageMemoryRequirements is used to provide ImageCreateInfo data to the
// DeviceImageMemoryRequirements extension method
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceImageMemoryRequirementsKHR.html
type DeviceImageMemoryRequirements struct {
	CreateInfo core1_0.ImageCreateInfo

	common.NextOptions
}

func (o DeviceImageMemoryRequirements) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(C.sizeof_struct_VkDeviceImageMemoryRequirements))
	}

	info := (*C.VkDeviceImageMemoryRequirementsKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_DEVICE_IMAGE_MEMORY_REQUIREMENTS
	info.pNext = next
	createInfo, err := common.AllocOptions(allocator, o.CreateInfo)
	if err != nil {
		return nil, err
	}

	info.pCreateInfo = (*C.VkImageCreateInfo)(createInfo)

	return preallocatedPointer, nil
}
