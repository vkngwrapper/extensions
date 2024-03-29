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

// SemaphoreWaitInfo contains information about the Semaphore wait condition
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreWaitInfo.html
type SemaphoreWaitInfo struct {
	// Flags specifies additional parameters for the Semaphore wait operation
	Flags SemaphoreWaitFlags
	// Semaphores is a slice of Semaphore objects to wait on
	Semaphores []core1_0.Semaphore
	// Values is a slice of timeline Semaphore values
	Values []uint64

	common.NextOptions
}

func (o SemaphoreWaitInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkSemaphoreWaitInfoKHR{})))
	}

	if len(o.Semaphores) != len(o.Values) {
		return nil, errors.Errorf("the SemaphoreWaitInfo 'Semaphores' list has %d elements, but the 'Values' list has %d elements- these lists must be the same size", len(o.Semaphores), len(o.Values))
	}

	info := (*C.VkSemaphoreWaitInfoKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_SEMAPHORE_WAIT_INFO_KHR
	info.pNext = next
	info.flags = C.VkSemaphoreWaitFlags(o.Flags)

	count := len(o.Semaphores)
	info.semaphoreCount = C.uint32_t(count)
	info.pSemaphores = nil
	info.pValues = nil

	if count > 0 {
		info.pSemaphores = (*C.VkSemaphore)(allocator.Malloc(count * int(unsafe.Sizeof([1]C.VkSemaphore{}))))
		info.pValues = (*C.uint64_t)(allocator.Malloc(count * int(unsafe.Sizeof(C.uint64_t(0)))))

		semaphoreSlice := unsafe.Slice(info.pSemaphores, count)
		valueSlice := unsafe.Slice(info.pValues, count)

		for i := 0; i < count; i++ {
			if o.Semaphores[i] == nil {
				return nil, errors.Errorf("the SemaphoreWaitInfo 'Semaphores' list has a nil semaphore at element %d- all elements must be non-nil", i)
			}

			semaphoreSlice[i] = C.VkSemaphore(unsafe.Pointer(o.Semaphores[i].Handle()))
			valueSlice[i] = C.uint64_t(o.Values[i])
		}
	}

	return preallocatedPointer, nil
}
