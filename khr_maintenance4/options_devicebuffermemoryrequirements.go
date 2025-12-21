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

// DeviceBufferMemoryRequirements is used to provide BufferCreateInfo data to
// the DeviceBufferMemoryRequirements extension method
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceBufferMemoryRequirementsKHR.html
type DeviceBufferMemoryRequirements struct {
	CreateInfo core1_0.BufferCreateInfo

	common.NextOptions
}

func (o DeviceBufferMemoryRequirements) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(C.sizeof_struct_VkDeviceBufferMemoryRequirements))
	}

	info := (*C.VkDeviceBufferMemoryRequirementsKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_DEVICE_BUFFER_MEMORY_REQUIREMENTS
	info.pNext = next
	createInfo, err := common.AllocOptions(allocator, o.CreateInfo)
	if err != nil {
		return nil, err
	}

	info.pCreateInfo = (*C.VkBufferCreateInfo)(createInfo)

	return preallocatedPointer, nil
}
