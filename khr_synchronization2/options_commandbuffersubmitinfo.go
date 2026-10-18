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

type CommandBufferSubmitInfo struct {
	CommandBuffer core1_0.CommandBuffer
	DeviceMask    uint32

	common.NextOptions
}

func (o CommandBufferSubmitInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if !o.CommandBuffer.Initialized() {
		return nil, errors.New("khr_synchronization2.CommandBufferSubmitInfo.CommandBuffer cannot be uninitialized")
	}
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof([1]C.VkCommandBufferSubmitInfo{})))
	}

	submitInfo := (*C.VkCommandBufferSubmitInfo)(preallocatedPointer)
	submitInfo.sType = C.VK_STRUCTURE_TYPE_COMMAND_BUFFER_SUBMIT_INFO
	submitInfo.pNext = next
	submitInfo.commandBuffer = C.VkCommandBuffer(unsafe.Pointer(o.CommandBuffer.Handle()))
	submitInfo.deviceMask = C.uint32_t(o.DeviceMask)

	return preallocatedPointer, nil
}
