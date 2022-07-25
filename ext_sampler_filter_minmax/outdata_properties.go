package ext_sampler_filter_minmax

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/common"
	"unsafe"
)

// PhysicalDeviceSamplerFilterMinmaxProperties describes Sampler filter minmax limits that can
// be supported by an implementation
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSamplerFilterMinmaxProperties.html
type PhysicalDeviceSamplerFilterMinmaxProperties struct {
	// FilterMinmaxSingleComponentFormats indicates whether a minimum set of required formats
	// support min/max filtering
	FilterMinmaxSingleComponentFormats bool
	// FilterMinmaxImageComponentMapping indicates whether the implementation support non-identity
	// component mapping of the Image when doing min/max filtering
	FilterMinmaxImageComponentMapping bool

	common.NextOutData
}

func (o *PhysicalDeviceSamplerFilterMinmaxProperties) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceSamplerFilterMinmaxPropertiesEXT{})))
	}

	info := (*C.VkPhysicalDeviceSamplerFilterMinmaxPropertiesEXT)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SAMPLER_FILTER_MINMAX_PROPERTIES_EXT
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceSamplerFilterMinmaxProperties) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceSamplerFilterMinmaxPropertiesEXT)(cDataPointer)

	o.FilterMinmaxSingleComponentFormats = info.filterMinmaxSingleComponentFormats != C.VkBool32(0)
	o.FilterMinmaxImageComponentMapping = info.filterMinmaxImageComponentMapping != C.VkBool32(0)

	return info.pNext, nil
}
