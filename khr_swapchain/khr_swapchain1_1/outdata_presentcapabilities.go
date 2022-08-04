package khr_swapchain1_1

/*
#include <stdlib.h>
#include "../../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/core1_1"
	"unsafe"
)

// DeviceGroupPresentCapabilities returns present capabilities from other PhysicalDevice objects
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupPresentCapabilitiesKHR.html
type DeviceGroupPresentCapabilities struct {
	// PresentMask is an array of masks, where the mask at element i is non-zero if PhysicalDevice i
	// has a presentation engine, and where bit j is set in element i if PhysicalDevice i can present
	// swapchain Image objects from PhysicalDevice j
	PresentMask [core1_1.MaxGroupSize]uint32
	// Modes indicates which Device group presentation modes are supported
	Modes DeviceGroupPresentModeFlags

	common.NextOutData
}

func (o *DeviceGroupPresentCapabilities) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkDeviceGroupPresentCapabilitiesKHR{})))
	}

	info := (*C.VkDeviceGroupPresentCapabilitiesKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_DEVICE_GROUP_PRESENT_CAPABILITIES_KHR
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *DeviceGroupPresentCapabilities) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkDeviceGroupPresentCapabilitiesKHR)(cDataPointer)

	for i := 0; i < core1_1.MaxGroupSize; i++ {
		o.PresentMask[i] = uint32(info.presentMask[i])
	}
	o.Modes = DeviceGroupPresentModeFlags(info.modes)

	return info.pNext, nil
}
