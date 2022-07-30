package khr_multiview

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/common"
	"unsafe"
)

// RenderPassMultiviewCreateInfo contains multiview information for all subpasses
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassMultiviewCreateInfo.html
type RenderPassMultiviewCreateInfo struct {
	// ViewMasks is a slice of view masks, where each mask is a bitfield of view indices describing
	// which views rendering is broadcast to in each subpass, when multiview is enabled
	ViewMasks []uint32
	// ViewOffsets is a slice of view offsets, one for each subpass dependency. Each view offset
	// controls which view in the source subpass the views in the destination subpass depends on.
	ViewOffsets []int
	// CorrelationMasks is a slice of view masks indicating stes of views that may be
	// more efficient to render concurrently
	CorrelationMasks []uint32

	common.NextOptions
}

func (o RenderPassMultiviewCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkRenderPassMultiviewCreateInfoKHR{})))
	}

	info := (*C.VkRenderPassMultiviewCreateInfoKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_RENDER_PASS_MULTIVIEW_CREATE_INFO_KHR
	info.pNext = next

	count := len(o.ViewMasks)
	info.subpassCount = C.uint32_t(count)
	info.pViewMasks = nil
	if count > 0 {
		viewMasks := (*C.uint32_t)(allocator.Malloc(count * int(unsafe.Sizeof(C.uint32_t(0)))))
		viewMaskSlice := ([]C.uint32_t)(unsafe.Slice(viewMasks, count))

		for i := 0; i < count; i++ {
			viewMaskSlice[i] = C.uint32_t(o.ViewMasks[i])
		}
		info.pViewMasks = viewMasks
	}

	count = len(o.ViewOffsets)
	info.dependencyCount = C.uint32_t(count)
	info.pViewOffsets = nil
	if count > 0 {
		viewOffsets := (*C.int32_t)(allocator.Malloc(count * int(unsafe.Sizeof(C.int32_t(0)))))
		viewOffsetSlice := ([]C.int32_t)(unsafe.Slice(viewOffsets, count))

		for i := 0; i < count; i++ {
			viewOffsetSlice[i] = C.int32_t(o.ViewOffsets[i])
		}
		info.pViewOffsets = viewOffsets
	}

	count = len(o.CorrelationMasks)
	info.correlationMaskCount = C.uint32_t(count)
	info.pCorrelationMasks = nil
	if count > 0 {
		correlationMasks := (*C.uint32_t)(allocator.Malloc(count * int(unsafe.Sizeof(C.uint32_t(0)))))
		correlationMaskSlice := ([]C.uint32_t)(unsafe.Slice(correlationMasks, count))

		for i := 0; i < count; i++ {
			correlationMaskSlice[i] = C.uint32_t(o.CorrelationMasks[i])
		}
		info.pCorrelationMasks = correlationMasks
	}

	return preallocatedPointer, nil
}
