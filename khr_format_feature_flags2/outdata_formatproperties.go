package khr_format_feature_flags2

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

type FormatProperties3 struct {
	LinearTilingFeatures  FormatFeatureFlags2
	OptimalTilingFeatures FormatFeatureFlags2
	BufferFeatures        FormatFeatureFlags2

	common.NextOutData
}

func (o *FormatProperties3) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkFormatProperties3KHR{})))
	}

	info := (*C.VkFormatProperties3KHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_FORMAT_PROPERTIES_3_KHR
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *FormatProperties3) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkFormatProperties3KHR)(cDataPointer)

	o.LinearTilingFeatures = FormatFeatureFlags2(info.linearTilingFeatures)
	o.OptimalTilingFeatures = FormatFeatureFlags2(info.optimalTilingFeatures)
	o.BufferFeatures = FormatFeatureFlags2(info.bufferFeatures)

	return info.pNext, nil
}
