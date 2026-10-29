package ext_extended_dynamic_state2

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

type PhysicalDeviceExtendedDynamicState2Features struct {
	ExtendedDynamicState2                   bool
	ExtendedDynamicState2LogicOp            bool
	ExtendedDynamicState2PatchControlPoints bool

	common.NextOptions
	common.NextOutData
}

func (o PhysicalDeviceExtendedDynamicState2Features) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof([1]C.VkPhysicalDeviceExtendedDynamicState2FeaturesEXT{})))
	}

	features := (*C.VkPhysicalDeviceExtendedDynamicState2FeaturesEXT)(preallocatedPointer)
	features.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTENDED_DYNAMIC_STATE_2_FEATURES_EXT
	features.pNext = next
	features.extendedDynamicState2 = C.VkBool32(0)
	features.extendedDynamicState2LogicOp = C.VkBool32(0)
	features.extendedDynamicState2PatchControlPoints = C.VkBool32(0)

	if o.ExtendedDynamicState2 {
		features.extendedDynamicState2 = C.VkBool32(1)
	}
	if o.ExtendedDynamicState2LogicOp {
		features.extendedDynamicState2LogicOp = C.VkBool32(1)
	}
	if o.ExtendedDynamicState2PatchControlPoints {
		features.extendedDynamicState2PatchControlPoints = C.VkBool32(1)
	}

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceExtendedDynamicState2Features) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof([1]C.VkPhysicalDeviceExtendedDynamicState2FeaturesEXT{})))
	}

	features := (*C.VkPhysicalDeviceExtendedDynamicState2FeaturesEXT)(preallocatedPointer)
	features.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTENDED_DYNAMIC_STATE_2_FEATURES_EXT
	features.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceExtendedDynamicState2Features) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	features := (*C.VkPhysicalDeviceExtendedDynamicState2FeaturesEXT)(cDataPointer)
	o.ExtendedDynamicState2 = features.extendedDynamicState2 != C.VkBool32(0)
	o.ExtendedDynamicState2LogicOp = features.extendedDynamicState2LogicOp != C.VkBool32(0)
	o.ExtendedDynamicState2PatchControlPoints = features.extendedDynamicState2PatchControlPoints != C.VkBool32(0)

	return features.pNext, nil
}
