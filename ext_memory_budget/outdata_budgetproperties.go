package ext_memory_budget

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

// PhysicalDeviceMemoryBudgetProperties specifies PhysicalDevice memory budget and usage
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMemoryBudgetPropertiesEXT.html
type PhysicalDeviceMemoryBudgetProperties struct {
	HeapBudget [common.MaxMemoryHeaps]int
	HeapUsage  [common.MaxMemoryHeaps]int

	common.NextOutData
}

func (o *PhysicalDeviceMemoryBudgetProperties) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(C.sizeof_struct_VkPhysicalDeviceMemoryBudgetPropertiesEXT))
	}
	info := (*C.VkPhysicalDeviceMemoryBudgetPropertiesEXT)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MEMORY_BUDGET_PROPERTIES_EXT
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceMemoryBudgetProperties) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	outData := (*C.VkPhysicalDeviceMemoryBudgetPropertiesEXT)(cDataPointer)

	for i := 0; i < common.MaxMemoryHeaps; i++ {
		o.HeapBudget[i] = int(outData.heapBudget[i])
		o.HeapUsage[i] = int(outData.heapUsage[i])
	}

	return outData.pNext, nil
}
