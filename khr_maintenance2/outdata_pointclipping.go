package khr_maintenance2

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

// PhysicalDevicePointClippingProperties describes the point clipping behavior supported
// by an implementation
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevicePointClippingProperties.html
type PhysicalDevicePointClippingProperties struct {
	// PointClippingBehavior specifies the point clipping behavior supported by the implementation
	PointClippingBehavior PointClippingBehavior

	common.NextOutData
}

func (o *PhysicalDevicePointClippingProperties) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDevicePointClippingPropertiesKHR{})))
	}

	properties := (*C.VkPhysicalDevicePointClippingPropertiesKHR)(preallocatedPointer)
	properties.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_POINT_CLIPPING_PROPERTIES_KHR
	properties.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDevicePointClippingProperties) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	properties := (*C.VkPhysicalDevicePointClippingPropertiesKHR)(cDataPointer)
	o.PointClippingBehavior = PointClippingBehavior(properties.pointClippingBehavior)

	return properties.pNext, nil
}
