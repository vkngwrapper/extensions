package khr_external_memory

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/extensions/v3/khr_external_memory_capabilities"
)

// ExportMemoryAllocateInfo specifies exportable handle types for a DeviceMemory object
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMemoryAllocateInfo.html
type ExportMemoryAllocateInfo struct {
	// HandleTypes specifies one or more memory handle types the application can export from
	// the resulting allocation
	HandleTypes khr_external_memory_capabilities.ExternalMemoryHandleTypeFlags

	common.NextOptions
}

func (o ExportMemoryAllocateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkExportMemoryAllocateInfoKHR{})))
	}

	info := (*C.VkExportMemoryAllocateInfoKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_EXPORT_MEMORY_ALLOCATE_INFO_KHR
	info.pNext = next
	info.handleTypes = C.VkExternalMemoryHandleTypeFlagsKHR(o.HandleTypes)

	return preallocatedPointer, nil
}

func (o ExportMemoryAllocateInfo) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkExportMemoryAllocateInfoKHR)(cDataPointer)
	return info.pNext, nil
}
