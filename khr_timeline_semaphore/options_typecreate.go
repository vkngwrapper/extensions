package khr_timeline_semaphore

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/common"
	"unsafe"
)

// SemaphoreTypeCreateInfo specifies the type of a newly-created Semaphore
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreTypeCreateInfo.html
type SemaphoreTypeCreateInfo struct {
	// SemaphoreType specifies the type of the Semaphore
	SemaphoreType SemaphoreType
	// InitialValue is the initial payload value if SemaphoreType is SemaphoreTypeTimeline
	InitialValue uint64

	common.NextOptions
}

func (o SemaphoreTypeCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkSemaphoreTypeCreateInfoKHR{})))
	}

	info := (*C.VkSemaphoreTypeCreateInfoKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_SEMAPHORE_TYPE_CREATE_INFO_KHR
	info.pNext = next
	info.semaphoreType = C.VkSemaphoreType(o.SemaphoreType)
	info.initialValue = C.uint64_t(o.InitialValue)

	return preallocatedPointer, nil
}
