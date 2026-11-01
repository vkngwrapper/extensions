package ext_image_robustness

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

type PhysicalDeviceImageRobustnessFeatures struct {
	RobustImageAccess bool

	common.NextOptions
	common.NextOutData
}

func (o PhysicalDeviceImageRobustnessFeatures) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof([1]C.VkPhysicalDeviceImageRobustnessFeaturesEXT{})))
	}

	features := (*C.VkPhysicalDeviceImageRobustnessFeaturesEXT)(preallocatedPointer)
	features.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_ROBUSTNESS_FEATURES_EXT
	features.pNext = next
	features.robustImageAccess = C.VkBool32(0)
	if o.RobustImageAccess {
		features.robustImageAccess = C.VkBool32(1)
	}

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceImageRobustnessFeatures) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof([1]C.VkPhysicalDeviceImageRobustnessFeaturesEXT{})))
	}

	features := (*C.VkPhysicalDeviceImageRobustnessFeaturesEXT)(preallocatedPointer)
	features.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_ROBUSTNESS_FEATURES_EXT
	features.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceImageRobustnessFeatures) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	features := (*C.VkPhysicalDeviceImageRobustnessFeaturesEXT)(cDataPointer)
	o.RobustImageAccess = features.robustImageAccess != C.VkBool32(0)

	return features.pNext, nil
}
