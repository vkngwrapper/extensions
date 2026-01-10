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

type SurfaceFormat2 struct {
	SurfaceFormat khr_surface.SurfaceFormat

	common.NextOutData
}

func (o *SurfaceFormat2) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkSurfaceFormat2KHR{})))
	}

	info := (*C.VkSurfaceFormat2KHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_SURFACE_FORMAT_2_KHR
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *SurfaceFormat2) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkSurfaceFormat2KHR)(cDataPointer)

	o.SurfaceFormat = khr_surface.SurfaceFormat{
		Format:     core1_0.Format(info.surfaceFormat.format),
		ColorSpace: khr_surface.ColorSpace(info.surfaceFormat.colorSpace),
	}

	return info.pNext, nil
}
