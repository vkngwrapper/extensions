//go:build windows

package ext_full_screen_exclusive

/*
#define VK_USE_PLATFORM_WIN32_KHR
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3/common"
)

type SurfaceCapabilitiesFullScreenExclusive struct {
	FullScreenExclusiveSupported bool

	common.OutData
}

func (o *SurfaceCapabilitiesFullScreenExclusive) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkSurfaceCapabilitiesFullScreenExclusiveEXT{})))
	}

	info := (*C.VkSurfaceCapabilitiesFullScreenExclusiveEXT)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_SURFACE_CAPABILITIES_FULL_SCREEN_EXCLUSIVE_EXT
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *SurfaceCapabilitiesFullScreenExclusive) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkSurfaceCapabilitiesFullScreenExclusiveEXT)(cDataPointer)

	o.FullScreenExclusiveSupported = (info.fullScreenExclusiveSupported != 0)

	return info.pNext, nil
}
