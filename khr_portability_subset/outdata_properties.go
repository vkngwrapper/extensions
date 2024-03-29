package khr_portability_subset

/*
#define VK_ENABLE_BETA_EXTENSIONS 1
#include <stdlib.h>
#include "../vulkan/vulkan.h"
#include "../vulkan/vulkan_beta.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v2/common"
	"unsafe"
)

// PhysicalDevicePortabilitySubsetProperties describes additional properties supported by a portable implementation
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevicePortabilitySubsetPropertiesKHR.html
type PhysicalDevicePortabilitySubsetProperties struct {
	// MinVertexInputBindingStrideAlignment indicates the minimum alignment for vertex input strides
	MinVertexInputBindingStrideAlignment int

	common.NextOutData
}

func (o *PhysicalDevicePortabilitySubsetProperties) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkPhysicalDevicePortabilitySubsetPropertiesKHR)
	}

	outData := (*C.VkPhysicalDevicePortabilitySubsetPropertiesKHR)(preallocatedPointer)
	outData.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PORTABILITY_SUBSET_PROPERTIES_KHR
	outData.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDevicePortabilitySubsetProperties) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	outData := (*C.VkPhysicalDevicePortabilitySubsetPropertiesKHR)(cDataPointer)
	o.MinVertexInputBindingStrideAlignment = int(outData.minVertexInputBindingStrideAlignment)

	return outData.pNext, nil
}
