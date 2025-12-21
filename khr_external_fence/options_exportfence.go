package khr_external_fence

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/extensions/v3/khr_external_fence_capabilities"
)

// ExportFenceCreateInfo specifies handle types that can be exported from a Fence
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportFenceCreateInfo.html
type ExportFenceCreateInfo struct {
	// HandleTypes specifies one or more Fence handle types the application can export from
	// the resulting Fence
	HandleTypes khr_external_fence_capabilities.ExternalFenceHandleTypeFlags

	common.NextOptions
}

func (o ExportFenceCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkExportFenceCreateInfoKHR{})))
	}

	info := (*C.VkExportFenceCreateInfoKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_EXPORT_FENCE_CREATE_INFO_KHR
	info.pNext = next
	info.handleTypes = C.VkExternalFenceHandleTypeFlags(o.HandleTypes)

	return preallocatedPointer, nil
}
