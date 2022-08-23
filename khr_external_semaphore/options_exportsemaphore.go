package khr_external_semaphore

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/extensions/v2/khr_external_semaphore_capabilities"
	"unsafe"
)

// ExportSemaphoreCreateInfo specifies handle types that can be exported from a Semaphore
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportSemaphoreCreateInfo.html
type ExportSemaphoreCreateInfo struct {
	// HandleTypes specifies one or more Semaphore handle types the application can export
	// from the resulting Semaphore
	HandleTypes khr_external_semaphore_capabilities.ExternalSemaphoreHandleTypeFlags

	common.NextOptions
}

func (o ExportSemaphoreCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkExportSemaphoreCreateInfoKHR{})))
	}

	info := (*C.VkExportSemaphoreCreateInfoKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_EXPORT_SEMAPHORE_CREATE_INFO_KHR
	info.pNext = next
	info.handleTypes = C.VkExternalSemaphoreHandleTypeFlags(o.HandleTypes)

	return preallocatedPointer, nil
}
