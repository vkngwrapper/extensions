package khr_timeline_semaphore

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

// TimelineSemaphoreSubmitInfo specifies signal and wait values for timeline Semaphore objects
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkTimelineSemaphoreSubmitInfoKHR.html
type TimelineSemaphoreSubmitInfo struct {
	// WaitSemaphoreValues is a slice of values for the corresponding Semaphore objects in
	// SubmitInfo.WaitSemaphores to wait for
	WaitSemaphoreValues []uint64
	// SignalSemaphoreValues is a slice of values for the corresponding Semaphore objects in
	// SubmitInfo.SignalSemaphores to set when signaled
	SignalSemaphoreValues []uint64

	common.NextOptions
}

func (o TimelineSemaphoreSubmitInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkTimelineSemaphoreSubmitInfoKHR{})))
	}

	info := (*C.VkTimelineSemaphoreSubmitInfoKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_TIMELINE_SEMAPHORE_SUBMIT_INFO_KHR
	info.pNext = next

	count := len(o.WaitSemaphoreValues)
	info.waitSemaphoreValueCount = C.uint32_t(count)
	info.pWaitSemaphoreValues = nil

	if count > 0 {
		info.pWaitSemaphoreValues = (*C.uint64_t)(allocator.Malloc(count * int(unsafe.Sizeof(C.uint64_t(0)))))
		waitSlice := unsafe.Slice(info.pWaitSemaphoreValues, count)

		for i := 0; i < count; i++ {
			waitSlice[i] = C.uint64_t(o.WaitSemaphoreValues[i])
		}
	}

	count = len(o.SignalSemaphoreValues)
	info.signalSemaphoreValueCount = C.uint32_t(count)
	info.pSignalSemaphoreValues = nil

	if count > 0 {
		info.pSignalSemaphoreValues = (*C.uint64_t)(allocator.Malloc(count * int(unsafe.Sizeof(C.uint64_t(0)))))
		signalSlice := unsafe.Slice(info.pSignalSemaphoreValues, count)

		for i := 0; i < count; i++ {
			signalSlice[i] = C.uint64_t(o.SignalSemaphoreValues[i])
		}
	}

	return preallocatedPointer, nil
}
