package khr_synchronization2

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"errors"
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
)

type SemaphoreSubmitInfo struct {
	Semaphore   core1_0.Semaphore
	Value       uint64
	StageMask   PipelineStageFlags2
	DeviceIndex int

	common.NextOptions
}

func (o SemaphoreSubmitInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if !o.Semaphore.Initialized() {
		return nil, errors.New("khr_synchronization2.SemaphoreSubmitInfo.Semaphore cannot be uninitialized")
	}
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof([1]C.VkSemaphoreSubmitInfo{})))
	}

	submitInfo := (*C.VkSemaphoreSubmitInfo)(preallocatedPointer)
	submitInfo.sType = C.VK_STRUCTURE_TYPE_SEMAPHORE_SUBMIT_INFO
	submitInfo.pNext = next
	submitInfo.semaphore = C.VkSemaphore(unsafe.Pointer(o.Semaphore.Handle()))
	submitInfo.value = C.uint64_t(o.Value)
	submitInfo.stageMask = C.VkPipelineStageFlags2(o.StageMask)
	submitInfo.deviceIndex = C.uint32_t(o.DeviceIndex)

	return preallocatedPointer, nil
}
