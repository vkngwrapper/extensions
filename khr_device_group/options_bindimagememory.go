package khr_device_group

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/core1_0"
	"unsafe"
)

// BindImageMemoryDeviceGroupInfo specifies Device within a group to bind to
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindImageMemoryDeviceGroupInfo.html
type BindImageMemoryDeviceGroupInfo struct {
	// DeviceIndices is a slice of Device indices
	DeviceIndices []int
	// SplitInstanceBindRegions is a slice of Rect2D structures describing which regions of
	// the Image are attached to each instance of DeviceMemory
	SplitInstanceBindRegions []core1_0.Rect2D

	common.NextOptions
}

func (o BindImageMemoryDeviceGroupInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkBindImageMemoryDeviceGroupInfoKHR{})))
	}

	info := (*C.VkBindImageMemoryDeviceGroupInfoKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_BIND_IMAGE_MEMORY_DEVICE_GROUP_INFO_KHR
	info.pNext = next

	count := len(o.DeviceIndices)
	info.deviceIndexCount = C.uint32_t(count)
	info.pDeviceIndices = nil
	if count > 0 {
		indices := (*C.uint32_t)(allocator.Malloc(count * int(unsafe.Sizeof(C.uint32_t(0)))))
		indexSlice := ([]C.uint32_t)(unsafe.Slice(indices, count))

		for i := 0; i < count; i++ {
			indexSlice[i] = C.uint32_t(o.DeviceIndices[i])
		}

		info.pDeviceIndices = indices
	}

	count = len(o.SplitInstanceBindRegions)
	info.splitInstanceBindRegionCount = C.uint32_t(count)
	info.pSplitInstanceBindRegions = nil
	if count > 0 {
		regions := (*C.VkRect2D)(allocator.Malloc(count * C.sizeof_struct_VkRect2D))
		regionSlice := ([]C.VkRect2D)(unsafe.Slice(regions, count))

		for i := 0; i < count; i++ {
			regionSlice[i].offset.x = C.int32_t(o.SplitInstanceBindRegions[i].Offset.X)
			regionSlice[i].offset.y = C.int32_t(o.SplitInstanceBindRegions[i].Offset.Y)
			regionSlice[i].extent.width = C.uint32_t(o.SplitInstanceBindRegions[i].Extent.Width)
			regionSlice[i].extent.height = C.uint32_t(o.SplitInstanceBindRegions[i].Extent.Height)
		}

		info.pSplitInstanceBindRegions = regions
	}

	return preallocatedPointer, nil
}
