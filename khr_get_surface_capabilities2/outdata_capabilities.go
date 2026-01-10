package khr_get_surface_capabilities2

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
	"github.com/vkngwrapper/extensions/v3/khr_surface"
)

type SurfaceCapabilities2 struct {
	SurfaceCapabilities khr_surface.SurfaceCapabilities

	common.NextOutData
}

func (o *SurfaceCapabilities2) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkSurfaceCapabilities2KHR{})))
	}

	info := (*C.VkSurfaceCapabilities2KHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_SURFACE_CAPABILITIES_2_KHR
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *SurfaceCapabilities2) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkSurfaceCapabilities2KHR)(cDataPointer)

	o.SurfaceCapabilities = khr_surface.SurfaceCapabilities{
		MinImageCount: int(info.surfaceCapabilities.minImageCount),
		MaxImageCount: int(info.surfaceCapabilities.maxImageCount),
		CurrentExtent: core1_0.Extent2D{
			Width:  int(info.surfaceCapabilities.currentExtent.width),
			Height: int(info.surfaceCapabilities.currentExtent.height),
		},
		MinImageExtent: core1_0.Extent2D{
			Width:  int(info.surfaceCapabilities.minImageExtent.width),
			Height: int(info.surfaceCapabilities.minImageExtent.height),
		},
		MaxImageExtent: core1_0.Extent2D{
			Width:  int(info.surfaceCapabilities.maxImageExtent.width),
			Height: int(info.surfaceCapabilities.maxImageExtent.height),
		},
		MaxImageArrayLayers: int(info.surfaceCapabilities.maxImageArrayLayers),

		SupportedTransforms: khr_surface.SurfaceTransformFlags(info.surfaceCapabilities.supportedTransforms),
		CurrentTransform:    khr_surface.SurfaceTransformFlags(info.surfaceCapabilities.currentTransform),

		SupportedCompositeAlpha: khr_surface.CompositeAlphaFlags(info.surfaceCapabilities.supportedCompositeAlpha),
		SupportedUsageFlags:     core1_0.ImageUsageFlags(info.surfaceCapabilities.supportedUsageFlags),
	}

	return info.pNext, nil
}
