package khr_get_surface_capabilities2

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
	"github.com/vkngwrapper/extensions/v3/khr_surface"
)

type PhysicalDeviceSurfaceInfo2 struct {
	Surface khr_surface.Surface

	common.NextOptions
}

func (o PhysicalDeviceSurfaceInfo2) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if !o.Surface.Initialized() {
		return nil, errors.New("khr_swapchain.SwapchainCreateInfo.Surface cannot be uninitialized")
	}

	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof([1]C.VkPhysicalDeviceSurfaceInfo2KHR{})))
	}

	createInfo := (*C.VkPhysicalDeviceSurfaceInfo2KHR)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SURFACE_INFO_2_KHR
	createInfo.surface = C.VkSurfaceKHR(unsafe.Pointer(o.Surface.Handle()))
	createInfo.pNext = next

	return preallocatedPointer, nil
}
