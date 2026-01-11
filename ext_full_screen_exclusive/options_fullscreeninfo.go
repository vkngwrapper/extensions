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

type SurfaceFullScreenExclusiveInfo struct {
	FullScreenExclusive FullScreenExclusive

	common.Options
}

func (o SurfaceFullScreenExclusiveInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof([1]C.VkSurfaceFullScreenExclusiveInfoEXT{})))
	}
	createInfo := (*C.VkSurfaceFullScreenExclusiveInfoEXT)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_SURFACE_FULL_SCREEN_EXCLUSIVE_INFO_EXT
	createInfo.fullScreenExclusive = (C.VkFullScreenExclusiveEXT)(o.FullScreenExclusive)
	createInfo.pNext = next

	return preallocatedPointer, nil
}
