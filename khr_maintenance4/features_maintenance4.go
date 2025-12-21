package khr_maintenance4

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3/common"
)

// PhysicalDeviceMaintenance4Features describes support for maintenance4 functionality
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMaintenance4Features.html
type PhysicalDeviceMaintenance4Features struct {
	Maintenance4 bool

	common.NextOptions
	common.NextOutData
}

func (o PhysicalDeviceMaintenance4Features) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(C.sizeof_struct_VkPhysicalDeviceMaintenance4Features))
	}
	info := (*C.VkPhysicalDeviceMaintenance4FeaturesKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MAINTENANCE_4_FEATURES
	info.pNext = next
	info.maintenance4 = C.VkBool32(0)

	if o.Maintenance4 {
		info.maintenance4 = C.VkBool32(1)
	}

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceMaintenance4Features) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(C.sizeof_struct_VkPhysicalDeviceMaintenance4Features))
	}
	info := (*C.VkPhysicalDeviceMaintenance4FeaturesKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MAINTENANCE_4_FEATURES
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceMaintenance4Features) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	outData := (*C.VkPhysicalDeviceMaintenance4FeaturesKHR)(cDataPointer)
	o.Maintenance4 = outData.maintenance4 != C.VkBool32(0)

	return outData.pNext, nil
}
