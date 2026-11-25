package ext_private_data

import (
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3/common"
)

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"

type DevicePrivateDataCreateInfo struct {
	PrivateDataSlotRequestCount int

	common.NextOptions
}

func (o DevicePrivateDataCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof([1]C.VkDevicePrivateDataCreateInfoEXT{})))
	}

	features := (*C.VkDevicePrivateDataCreateInfoEXT)(preallocatedPointer)
	features.sType = C.VK_STRUCTURE_TYPE_DEVICE_PRIVATE_DATA_CREATE_INFO
	features.pNext = next
	features.privateDataSlotRequestCount = C.uint32_t(o.PrivateDataSlotRequestCount)

	return preallocatedPointer, nil
}
