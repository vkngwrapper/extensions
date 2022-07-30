package khr_dedicated_allocation

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/driver"
	"unsafe"
)

// MemoryDedicatedRequirements describes dedicated allocation requirements of Buffer and Image
// resources
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryDedicatedRequirements.html
type MemoryDedicatedRequirements struct {
	// PrefersDedicatedAllocation specifies that the implementation would prefer a dedicated
	// allocation for this resource. The application is still free to suballocate the resource
	// but it may get better performance if a dedicated allocation is used
	PrefersDedicatedAllocation bool
	// RequiresDedicatedAllocation specifies that a dedicated allocation is required for this resource
	RequiresDedicatedAllocation bool

	common.NextOutData
}

func (o *MemoryDedicatedRequirements) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkMemoryDedicatedRequirementsKHR{})))
	}

	outData := (*C.VkMemoryDedicatedRequirementsKHR)(preallocatedPointer)
	outData.sType = C.VK_STRUCTURE_TYPE_MEMORY_DEDICATED_REQUIREMENTS_KHR
	outData.pNext = next

	return preallocatedPointer, nil
}

func (o *MemoryDedicatedRequirements) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	outData := (*C.VkMemoryDedicatedRequirementsKHR)(cDataPointer)
	o.RequiresDedicatedAllocation = driver.VkBool32(outData.requiresDedicatedAllocation) != driver.VkBool32(0)
	o.PrefersDedicatedAllocation = driver.VkBool32(outData.prefersDedicatedAllocation) != driver.VkBool32(0)

	return outData.pNext, nil
}
