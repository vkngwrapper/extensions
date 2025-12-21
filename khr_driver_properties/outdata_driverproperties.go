package khr_driver_properties

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

// PhysicalDeviceDriverProperties contains driver identification information
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDriverProperties.html
type PhysicalDeviceDriverProperties struct {
	// DriverID is a unique identifier for the driver of the PhysicalDevice
	DriverID DriverID
	// DriverName is a string which is the name of the driver
	DriverName string
	// DriverInfo is a string with additional information about the driver
	DriverInfo string
	// ConformanceVersion is the version of the Vulkan conformance test thsi driver is conformant
	// against
	ConformanceVersion ConformanceVersion

	common.NextOutData
}

func (o *PhysicalDeviceDriverProperties) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceDriverPropertiesKHR{})))
	}

	outData := (*C.VkPhysicalDeviceDriverPropertiesKHR)(preallocatedPointer)
	outData.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DRIVER_PROPERTIES_KHR
	outData.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceDriverProperties) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	outData := (*C.VkPhysicalDeviceDriverPropertiesKHR)(cDataPointer)
	o.DriverID = DriverID(outData.driverID)
	o.ConformanceVersion.Major = uint8(outData.conformanceVersion.major)
	o.ConformanceVersion.Minor = uint8(outData.conformanceVersion.minor)
	o.ConformanceVersion.Subminor = uint8(outData.conformanceVersion.subminor)
	o.ConformanceVersion.Patch = uint8(outData.conformanceVersion.patch)
	o.DriverName = C.GoString(&outData.driverName[0])
	o.DriverInfo = C.GoString(&outData.driverInfo[0])

	return outData.pNext, nil
}
