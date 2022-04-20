package khr_device_group

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/extensions/khr_device_group_creation"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type DeviceGroupPresentCapabilitiesOutData struct {
	PresentMask [khr_device_group_creation.MaxGroupSize]uint32
	Modes       DeviceGroupPresentModeFlags

	common.HaveNext
}

func (o *DeviceGroupPresentCapabilitiesOutData) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkDeviceGroupPresentCapabilitiesKHR{})))
	}

	info := (*C.VkDeviceGroupPresentCapabilitiesKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_DEVICE_GROUP_PRESENT_CAPABILITIES_KHR
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *DeviceGroupPresentCapabilitiesOutData) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkDeviceGroupPresentCapabilitiesKHR)(cDataPointer)

	for i := 0; i < khr_device_group_creation.MaxGroupSize; i++ {
		o.PresentMask[i] = uint32(info.presentMask[i])
	}
	o.Modes = DeviceGroupPresentModeFlags(info.modes)

	return info.pNext, nil
}
