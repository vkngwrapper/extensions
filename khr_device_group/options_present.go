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

// DeviceGroupPresentInfo controls which PhysicalDevice objects' Image objects are presented
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupPresentInfoKHR.html
type DeviceGroupPresentInfo struct {
	// DeviceMasks is a slice of Device masks, one for each element of khr_swapchain.PresentInfo.Swapchains
	DeviceMasks []uint32
	// Mode specifies the Device group present mode that will be used for this present
	Mode DeviceGroupPresentModeFlags

	common.NextOptions
}

func (o DeviceGroupPresentInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkDeviceGroupPresentInfoKHR)
	}

	info := (*C.VkDeviceGroupPresentInfoKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_DEVICE_GROUP_PRESENT_INFO_KHR
	info.pNext = next
	info.mode = C.VkDeviceGroupPresentModeFlagBitsKHR(o.Mode)

	count := len(o.DeviceMasks)
	info.swapchainCount = C.uint32_t(count)
	info.pDeviceMasks = nil

	if count > 0 {
		masks := (*C.uint32_t)(allocator.Malloc(count * int(unsafe.Sizeof(C.uint32_t(0)))))
		maskSlice := ([]C.uint32_t)(unsafe.Slice(masks, count))

		for i := 0; i < count; i++ {
			maskSlice[i] = C.uint32_t(o.DeviceMasks[i])
		}
		info.pDeviceMasks = masks
	}

	return preallocatedPointer, nil
}
