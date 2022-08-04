package khr_device_group

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/common"
	"unsafe"
)

// DeviceGroupSubmitInfo indicates which PhysicalDevice objects execute Semaphore operations
// and CommandBuffer objects
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupSubmitInfo.html
type DeviceGroupSubmitInfo struct {
	// WaitSemaphoreDeviceIndices is a slice of Device indices indicating which PhysicalDevice
	// executes the Semaphore wait operation in the corresponding element of SubmitInfo.WaitSemaphores
	WaitSemaphoreDeviceIndices []int
	// CommandBufferDeviceMasks is a slice of Device masks indicating which PhysicalDevice objects
	// execute the CommandBuffer in teh corresponding element of SubmitInfo.CommandBuffers
	CommandBufferDeviceMasks []uint32
	// SignalSemaphoreDeviceIndices is a slice of Device indices indicating which PhysicalDevice
	// executes the Semaphore signal operation in the SubmitInfo.SignalSemaphores
	SignalSemaphoreDeviceIndices []int

	common.NextOptions
}

func (o DeviceGroupSubmitInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkDeviceGroupSubmitInfoKHR{})))
	}

	info := (*C.VkDeviceGroupSubmitInfoKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_DEVICE_GROUP_SUBMIT_INFO_KHR
	info.pNext = next

	count := len(o.WaitSemaphoreDeviceIndices)
	info.waitSemaphoreCount = C.uint32_t(count)
	info.pWaitSemaphoreDeviceIndices = nil

	if count > 0 {
		indices := (*C.uint32_t)(allocator.Malloc(count * int(unsafe.Sizeof(C.uint32_t(0)))))
		indexSlice := ([]C.uint32_t)(unsafe.Slice(indices, count))

		for i := 0; i < count; i++ {
			indexSlice[i] = C.uint32_t(o.WaitSemaphoreDeviceIndices[i])
		}
		info.pWaitSemaphoreDeviceIndices = indices
	}

	count = len(o.CommandBufferDeviceMasks)
	info.commandBufferCount = C.uint32_t(count)
	info.pCommandBufferDeviceMasks = nil

	if count > 0 {
		masks := (*C.uint32_t)(allocator.Malloc(count * int(unsafe.Sizeof(C.uint32_t(0)))))
		maskSlice := ([]C.uint32_t)(unsafe.Slice(masks, count))

		for i := 0; i < count; i++ {
			maskSlice[i] = C.uint32_t(o.CommandBufferDeviceMasks[i])
		}
		info.pCommandBufferDeviceMasks = masks
	}

	count = len(o.SignalSemaphoreDeviceIndices)
	info.signalSemaphoreCount = C.uint32_t(count)
	info.pSignalSemaphoreDeviceIndices = nil

	if count > 0 {
		indices := (*C.uint32_t)(allocator.Malloc(count * int(unsafe.Sizeof(C.uint32_t(0)))))
		indexSlice := ([]C.uint32_t)(unsafe.Slice(indices, count))

		for i := 0; i < count; i++ {
			indexSlice[i] = C.uint32_t(o.SignalSemaphoreDeviceIndices[i])
		}
		info.pSignalSemaphoreDeviceIndices = indices
	}

	return preallocatedPointer, nil
}
