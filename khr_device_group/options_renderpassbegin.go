package khr_device_group

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
)

// DeviceGroupRenderPassBeginInfo sets the initial Device mask and render areas for a RenderPass
// instance
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupRenderPassBeginInfo.html
type DeviceGroupRenderPassBeginInfo struct {
	// DeviceMask is the deivce mask for the RenderPass instance
	DeviceMask uint32
	// DeviceRenderAreas is a slice of Rect2D structures defining the render area for each
	// PhysicalDevice
	DeviceRenderAreas []core1_0.Rect2D

	common.NextOptions
}

func (o DeviceGroupRenderPassBeginInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkDeviceGroupRenderPassBeginInfoKHR{})))
	}

	info := (*C.VkDeviceGroupRenderPassBeginInfoKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_DEVICE_GROUP_RENDER_PASS_BEGIN_INFO_KHR
	info.pNext = next
	info.deviceMask = C.uint32_t(o.DeviceMask)

	count := len(o.DeviceRenderAreas)
	info.deviceRenderAreaCount = C.uint32_t(count)
	info.pDeviceRenderAreas = nil

	if count > 0 {
		areas := (*C.VkRect2D)(allocator.Malloc(count * C.sizeof_struct_VkRect2D))
		areaSlice := ([]C.VkRect2D)(unsafe.Slice(areas, count))

		for i := 0; i < count; i++ {
			areaSlice[i].offset.x = C.int32_t(o.DeviceRenderAreas[i].Offset.X)
			areaSlice[i].offset.y = C.int32_t(o.DeviceRenderAreas[i].Offset.Y)
			areaSlice[i].extent.width = C.uint32_t(o.DeviceRenderAreas[i].Extent.Width)
			areaSlice[i].extent.height = C.uint32_t(o.DeviceRenderAreas[i].Extent.Height)
		}

		info.pDeviceRenderAreas = areas
	}

	return preallocatedPointer, nil
}
