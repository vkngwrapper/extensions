package khr_synchronization2

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

type DependencyInfo struct {
	DependencyFlags      core1_0.DependencyFlags
	MemoryBarriers       []MemoryBarrier2
	BufferMemoryBarriers []BufferMemoryBarrier2
	ImageMemoryBarriers  []ImageMemoryBarrier2

	common.NextOptions
}

func (o DependencyInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof([1]C.VkDependencyInfo{})))
	}
	createInfo := (*C.VkDependencyInfo)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_DEPENDENCY_INFO
	createInfo.pNext = next

	createInfo.dependencyFlags = C.VkDependencyFlags(o.DependencyFlags)

	barrierCount := len(o.MemoryBarriers)
	bufferBarrierCount := len(o.BufferMemoryBarriers)
	imageBarrierCount := len(o.ImageMemoryBarriers)
	createInfo.memoryBarrierCount = C.uint32_t(barrierCount)
	createInfo.pMemoryBarriers = nil
	createInfo.bufferMemoryBarrierCount = C.uint32_t(bufferBarrierCount)
	createInfo.pBufferMemoryBarriers = nil
	createInfo.imageMemoryBarrierCount = C.uint32_t(imageBarrierCount)
	createInfo.pImageMemoryBarriers = nil

	var err error
	if barrierCount > 0 {
		createInfo.pMemoryBarriers, err = common.AllocOptionSlice[C.VkMemoryBarrier2](allocator, o.MemoryBarriers)
		if err != nil {
			return nil, err
		}
	}

	if bufferBarrierCount > 0 {
		createInfo.pBufferMemoryBarriers, err = common.AllocOptionSlice[C.VkBufferMemoryBarrier2](allocator, o.BufferMemoryBarriers)
		if err != nil {
			return nil, err
		}
	}

	if imageBarrierCount > 0 {
		createInfo.pImageMemoryBarriers, err = common.AllocOptionSlice[C.VkImageMemoryBarrier2](allocator, o.ImageMemoryBarriers)
		if err != nil {
			return nil, err
		}
	}

	return preallocatedPointer, nil
}
