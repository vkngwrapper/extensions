package khr_synchronization2

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3/common"
)

type SubmitInfo2 struct {
	Flags                SubmitFlags
	WaitSemaphoreInfos   []SemaphoreSubmitInfo
	CommandBufferInfos   []CommandBufferSubmitInfo
	SignalSemaphoreInfos []SemaphoreSubmitInfo

	common.NextOptions
}

func (o SubmitInfo2) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof([1]C.VkSubmitInfo2{})))
	}
	submitInfo := (*C.VkSubmitInfo2)(preallocatedPointer)
	submitInfo.sType = C.VK_STRUCTURE_TYPE_SUBMIT_INFO_2
	submitInfo.pNext = next
	submitInfo.flags = C.VkSubmitFlags(o.Flags)
	waitSemaphoreCount := len(o.WaitSemaphoreInfos)
	commandBufferCount := len(o.CommandBufferInfos)
	signalSemaphoreCount := len(o.SignalSemaphoreInfos)
	submitInfo.waitSemaphoreInfoCount = C.uint32_t(waitSemaphoreCount)
	submitInfo.pWaitSemaphoreInfos = nil
	submitInfo.commandBufferInfoCount = C.uint32_t(commandBufferCount)
	submitInfo.pCommandBufferInfos = nil
	submitInfo.signalSemaphoreInfoCount = C.uint32_t(signalSemaphoreCount)
	submitInfo.pSignalSemaphoreInfos = nil

	var err error
	if waitSemaphoreCount > 0 {
		submitInfo.pWaitSemaphoreInfos, err = common.AllocOptionSlice[C.VkSemaphoreSubmitInfo](allocator, o.WaitSemaphoreInfos)
		if err != nil {
			return nil, err
		}
	}

	if commandBufferCount > 0 {
		submitInfo.pCommandBufferInfos, err = common.AllocOptionSlice[C.VkCommandBufferSubmitInfo](allocator, o.CommandBufferInfos)
		if err != nil {
			return nil, err
		}
	}

	if signalSemaphoreCount > 0 {
		submitInfo.pSignalSemaphoreInfos, err = common.AllocOptionSlice[C.VkSemaphoreSubmitInfo](allocator, o.SignalSemaphoreInfos)
		if err != nil {
			return nil, err
		}
	}

	return preallocatedPointer, nil
}
