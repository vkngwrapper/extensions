package amd_device_coherent_memory

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

// PhysicalDeviceCoherentMemoryFeatures indicates support for device coherent memory
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceCoherentMemoryFeaturesAMD.html
type PhysicalDeviceCoherentMemoryFeatures struct {
	DeviceCoherentMemory bool

	common.NextOptions
	common.NextOutData
}

func (o PhysicalDeviceCoherentMemoryFeatures) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(C.sizeof_struct_VkPhysicalDeviceCoherentMemoryFeaturesAMD))
	}

	info := (*C.VkPhysicalDeviceCoherentMemoryFeaturesAMD)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_COHERENT_MEMORY_FEATURES_AMD
	info.pNext = next
	info.deviceCoherentMemory = C.VkBool32(0)

	if o.DeviceCoherentMemory {
		info.deviceCoherentMemory = C.VkBool32(1)
	}

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceCoherentMemoryFeatures) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(C.sizeof_struct_VkPhysicalDeviceCoherentMemoryFeaturesAMD))
	}
	info := (*C.VkPhysicalDeviceCoherentMemoryFeaturesAMD)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_COHERENT_MEMORY_FEATURES_AMD
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceCoherentMemoryFeatures) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	outData := (*C.VkPhysicalDeviceCoherentMemoryFeaturesAMD)(cDataPointer)

	o.DeviceCoherentMemory = outData.deviceCoherentMemory != C.VkBool32(0)

	return outData.pNext, nil
}
