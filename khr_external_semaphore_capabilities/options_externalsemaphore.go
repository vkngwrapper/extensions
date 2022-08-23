package khr_external_semaphore_capabilities

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

// PhysicalDeviceExternalSemaphoreInfo specifies Semaphore creation parameters
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceExternalSemaphoreInfo.html
type PhysicalDeviceExternalSemaphoreInfo struct {
	// HandleType specifies the external Semaphore handle type for which capabilities will
	// be returned
	HandleType ExternalSemaphoreHandleTypeFlags

	common.NextOptions
}

func (o PhysicalDeviceExternalSemaphoreInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceExternalSemaphoreInfoKHR{})))
	}

	info := (*C.VkPhysicalDeviceExternalSemaphoreInfoKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_SEMAPHORE_INFO_KHR
	info.pNext = next
	info.handleType = C.VkExternalSemaphoreHandleTypeFlagBits(o.HandleType)

	return preallocatedPointer, nil
}
