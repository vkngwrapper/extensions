package khr_acceleration_structure

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

type PhysicalDeviceAccelerationStructureFeatures struct {
	AccelerationStructure                                 bool
	AccelerationStructureCaptureReplay                    bool
	AccelerationStructureIndirectBuild                    bool
	AccelerationStructureHostCommands                     bool
	DescriptorBindingAccelerationStructureUpdateAfterBind bool

	common.NextOptions
	common.NextOutData
}

func (o PhysicalDeviceAccelerationStructureFeatures) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceAccelerationStructureFeaturesKHR{})))
	}

	info := (*C.VkPhysicalDeviceAccelerationStructureFeaturesKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_ACCELERATION_STRUCTURE_FEATURES_KHR
	info.pNext = next
	info.accelerationStructure = C.VkBool32(0)
	info.accelerationStructureCaptureReplay = C.VkBool32(0)
	info.accelerationStructureIndirectBuild = C.VkBool32(0)
	info.accelerationStructureHostCommands = C.VkBool32(0)
	info.descriptorBindingAccelerationStructureUpdateAfterBind = C.VkBool32(0)

	if o.AccelerationStructure {
		info.accelerationStructure = C.VkBool32(1)
	}

	if o.AccelerationStructureCaptureReplay {
		info.accelerationStructureCaptureReplay = C.VkBool32(1)
	}

	if o.AccelerationStructureIndirectBuild {
		info.accelerationStructureIndirectBuild = C.VkBool32(1)
	}

	if o.AccelerationStructureHostCommands {
		info.accelerationStructureHostCommands = C.VkBool32(1)
	}

	if o.DescriptorBindingAccelerationStructureUpdateAfterBind {
		info.descriptorBindingAccelerationStructureUpdateAfterBind = C.VkBool32(1)
	}

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceAccelerationStructureFeatures) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceAccelerationStructureFeaturesKHR{})))
	}

	info := (*C.VkPhysicalDeviceAccelerationStructureFeaturesKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_ACCELERATION_STRUCTURE_FEATURES_KHR
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceAccelerationStructureFeatures) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	outData := (*C.VkPhysicalDeviceAccelerationStructureFeaturesKHR)(cDataPointer)

	o.AccelerationStructure = outData.accelerationStructure != C.VkBool32(0)
	o.AccelerationStructureCaptureReplay = outData.accelerationStructureCaptureReplay != C.VkBool32(0)
	o.AccelerationStructureIndirectBuild = outData.accelerationStructureIndirectBuild != C.VkBool32(0)
	o.AccelerationStructureHostCommands = outData.accelerationStructureHostCommands != C.VkBool32(0)
	o.DescriptorBindingAccelerationStructureUpdateAfterBind = outData.descriptorBindingAccelerationStructureUpdateAfterBind != C.VkBool32(0)

	return outData.pNext, nil
}
