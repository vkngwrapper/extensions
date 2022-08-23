package khr_maintenance2

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v2/common"
	"unsafe"
)

// PipelineTessellationDomainOriginStateCreateInfo specifies the origin of the tessellation domain
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineTessellationDomainOriginStateCreateInfo.html
type PipelineTessellationDomainOriginStateCreateInfo struct {
	// DomainOrigin controls the origin of the tessellation domain space
	DomainOrigin TessellationDomainOrigin

	common.NextOptions
}

func (o PipelineTessellationDomainOriginStateCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPipelineTessellationDomainOriginStateCreateInfoKHR{})))
	}

	createInfo := (*C.VkPipelineTessellationDomainOriginStateCreateInfoKHR)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_PIPELINE_TESSELLATION_DOMAIN_ORIGIN_STATE_CREATE_INFO_KHR
	createInfo.pNext = next
	createInfo.domainOrigin = (C.VkTessellationDomainOriginKHR)(o.DomainOrigin)

	return preallocatedPointer, nil
}
