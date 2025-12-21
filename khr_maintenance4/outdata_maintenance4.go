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

// PhysicalDeviceMaintenance4Properties describes properties introduced with khr_maintenance4
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMaintenance4Properties.html
type PhysicalDeviceMaintenance4Properties struct {
	MaxBufferSize int

	common.NextOutData
}

func (o *PhysicalDeviceMaintenance4Properties) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(C.sizeof_struct_VkPhysicalDeviceMaintenance4Properties))
	}
	info := (*C.VkPhysicalDeviceMaintenance4PropertiesKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MAINTENANCE_4_PROPERTIES
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceMaintenance4Properties) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	outData := (*C.VkPhysicalDeviceMaintenance4PropertiesKHR)(cDataPointer)
	o.MaxBufferSize = int(outData.maxBufferSize)

	return outData.pNext, nil
}
