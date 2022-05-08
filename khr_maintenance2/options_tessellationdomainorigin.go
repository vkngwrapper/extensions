package khr_maintenance2

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type PipelineTessellationDomainOriginStateOptions struct {
	DomainOrigin TessellationDomainOrigin
	common.HaveNext
}

func (o PipelineTessellationDomainOriginStateOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPipelineTessellationDomainOriginStateCreateInfoKHR{})))
	}

	createInfo := (*C.VkPipelineTessellationDomainOriginStateCreateInfoKHR)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_PIPELINE_TESSELLATION_DOMAIN_ORIGIN_STATE_CREATE_INFO_KHR
	createInfo.pNext = next
	createInfo.domainOrigin = (C.VkTessellationDomainOriginKHR)(o.DomainOrigin)

	return preallocatedPointer, nil
}

func (o PipelineTessellationDomainOriginStateOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	createInfo := (*C.VkPipelineTessellationDomainOriginStateCreateInfoKHR)(cDataPointer)
	return createInfo.pNext, nil
}

var _ common.Options = PipelineTessellationDomainOriginStateOptions{}
