package khr_timeline_semaphore

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/pkg/errors"
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/core1_0"
	"unsafe"
)

// SemaphoreSignalInfo contains information about a Semaphore signal operation
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreSignalInfo.html
type SemaphoreSignalInfo struct {
	// Semaphore is the Semaphore object to signal
	Semaphore core1_0.Semaphore
	// Value is the value to signal
	Value uint64

	common.NextOptions
}

func (o SemaphoreSignalInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkSemaphoreSignalInfoKHR{})))
	}

	if o.Semaphore == nil {
		return nil, errors.New("the 'Semaphore' field of SemaphoreSignalInfo must be non-nil")
	}

	info := (*C.VkSemaphoreSignalInfoKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_SEMAPHORE_SIGNAL_INFO_KHR
	info.pNext = next
	info.semaphore = C.VkSemaphore(unsafe.Pointer(o.Semaphore.Handle()))
	info.value = C.uint64_t(o.Value)

	return preallocatedPointer, nil
}
