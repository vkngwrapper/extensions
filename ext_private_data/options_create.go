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

type PrivateDataSlotCreateInfo struct {
	Flags PrivateDataSlotCreateFlags

	common.NextOptions
}

func (o PrivateDataSlotCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof([1]C.VkPrivateDataSlotCreateInfoEXT{})))
	}

	features := (*C.VkPrivateDataSlotCreateInfoEXT)(preallocatedPointer)
	features.sType = C.VK_STRUCTURE_TYPE_PRIVATE_DATA_SLOT_CREATE_INFO
	features.pNext = next
	features.flags = C.VkPrivateDataSlotCreateFlags(o.Flags)

	return preallocatedPointer, nil
}
