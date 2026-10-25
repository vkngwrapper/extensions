package ext_extended_dynamic_state

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3/common"
)

type PhysicalDeviceExtendedDynamicStateFeatures struct {
	ExtendedDynamicState bool

	common.NextOutData
	common.NextOptions
}

func (o PhysicalDeviceExtendedDynamicStateFeatures) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof([1]C.VkPhysicalDeviceExtendedDynamicStateFeaturesEXT{})))
	}

	features := (*C.VkPhysicalDeviceExtendedDynamicStateFeaturesEXT)(preallocatedPointer)
	features.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTENDED_DYNAMIC_STATE_FEATURES_EXT
	features.pNext = next
	features.extendedDynamicState = C.VkBool32(0)
	if o.ExtendedDynamicState {
		features.extendedDynamicState = C.VkBool32(1)
	}

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceExtendedDynamicStateFeatures) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof([1]C.VkPhysicalDeviceExtendedDynamicStateFeaturesEXT{})))
	}

	features := (*C.VkPhysicalDeviceExtendedDynamicStateFeaturesEXT)(preallocatedPointer)
	features.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTENDED_DYNAMIC_STATE_FEATURES_EXT
	features.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceExtendedDynamicStateFeatures) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	features := (*C.VkPhysicalDeviceExtendedDynamicStateFeaturesEXT)(cDataPointer)
	o.ExtendedDynamicState = features.extendedDynamicState != C.VkBool32(0)

	return features.pNext, nil
}
