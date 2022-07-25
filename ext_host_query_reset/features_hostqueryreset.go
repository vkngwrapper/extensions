package ext_host_query_reset

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

// PhysicalDeviceHostQueryResetFeatures describes whether queries can be reset from the host
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceHostQueryResetFeatures.html
type PhysicalDeviceHostQueryResetFeatures struct {
	// HostQueryReset indicates that hte implementation supports resetting queries from the host
	// with QueryPool.Reset
	HostQueryReset bool

	common.NextOptions
	common.NextOutData
}

func (o *PhysicalDeviceHostQueryResetFeatures) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceHostQueryResetFeaturesEXT{})))
	}

	info := (*C.VkPhysicalDeviceHostQueryResetFeaturesEXT)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_HOST_QUERY_RESET_FEATURES_EXT
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceHostQueryResetFeatures) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceHostQueryResetFeaturesEXT)(cDataPointer)
	o.HostQueryReset = info.hostQueryReset != C.VkBool32(0)

	return info.pNext, nil
}

func (o PhysicalDeviceHostQueryResetFeatures) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceHostQueryResetFeaturesEXT{})))
	}

	info := (*C.VkPhysicalDeviceHostQueryResetFeaturesEXT)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_HOST_QUERY_RESET_FEATURES_EXT
	info.pNext = next
	info.hostQueryReset = C.VkBool32(0)

	if o.HostQueryReset {
		info.hostQueryReset = C.VkBool32(1)
	}

	return preallocatedPointer, nil
}
