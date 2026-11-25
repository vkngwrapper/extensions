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

type PhysicalDevicePrivateDataFeatures struct {
	PrivateData bool

	common.NextOptions
	common.NextOutData
}

func (o PhysicalDevicePrivateDataFeatures) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof([1]C.VkPhysicalDevicePrivateDataFeaturesEXT{})))
	}

	features := (*C.VkPhysicalDevicePrivateDataFeaturesEXT)(preallocatedPointer)
	features.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PRIVATE_DATA_FEATURES
	features.pNext = next
	features.privateData = C.VkBool32(0)

	if o.PrivateData {
		features.privateData = C.VkBool32(1)
	}

	return preallocatedPointer, nil
}

func (o *PhysicalDevicePrivateDataFeatures) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof([1]C.VkPhysicalDevicePrivateDataFeaturesEXT{})))
	}

	features := (*C.VkPhysicalDevicePrivateDataFeaturesEXT)(preallocatedPointer)
	features.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PRIVATE_DATA_FEATURES
	features.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDevicePrivateDataFeatures) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	features := (*C.VkPhysicalDevicePrivateDataFeaturesEXT)(cDataPointer)
	o.PrivateData = features.privateData != C.VkBool32(0)

	return features.pNext, nil
}
