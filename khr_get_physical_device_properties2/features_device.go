package khr_get_physical_device_properties2

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/core1_0"
	"unsafe"
)

// PhysicalDeviceFeatures2 describes the fine-grained features that can be supported
// by an implementation
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html
type PhysicalDeviceFeatures2 struct {
	// Features describes the fine-grained features of the Vulkan 1.0 API
	Features core1_0.PhysicalDeviceFeatures

	common.NextOptions
	common.NextOutData
}

func (o *PhysicalDeviceFeatures2) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceFeatures2KHR{})))
	}

	data := (*C.VkPhysicalDeviceFeatures2KHR)(preallocatedPointer)
	data.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FEATURES_2_KHR
	data.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceFeatures2) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	data := (*C.VkPhysicalDeviceFeatures2KHR)(cDataPointer)

	(&o.Features).PopulateFromCPointer(unsafe.Pointer(&data.features))

	return data.pNext, nil
}

func (o PhysicalDeviceFeatures2) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceFeatures2KHR{})))
	}

	data := (*C.VkPhysicalDeviceFeatures2KHR)(preallocatedPointer)
	data.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FEATURES_2_KHR
	data.pNext = next
	_, err := o.Features.PopulateCPointer(allocator, unsafe.Pointer(&data.features))

	return preallocatedPointer, err
}
