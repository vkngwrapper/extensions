package khr_acceleration_structure

/*
   #include <stdlib.h>
   #include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/loader"
	khr_acceleration_structure_loader "github.com/vkngwrapper/extensions/v3/khr_acceleration_structure/loader"
)

type AccelerationStructure struct {
	handle khr_acceleration_structure_loader.VkAccelerationStructureKHR
	device loader.VkDevice

	apiVersion common.APIVersion
}

func (s AccelerationStructure) Handle() khr_acceleration_structure_loader.VkAccelerationStructureKHR {
	return s.handle
}

func (s AccelerationStructure) DeviceHandle() loader.VkDevice {
	return s.device
}

func (s AccelerationStructure) APIVersion() common.APIVersion {
	return s.apiVersion
}

func (s AccelerationStructure) Initialized() bool {
	return s.handle != 0
}

func InternalAccelerationStructure(device loader.VkDevice, handle khr_acceleration_structure_loader.VkAccelerationStructureKHR, version common.APIVersion) AccelerationStructure {
	return AccelerationStructure{
		device:     device,
		handle:     handle,
		apiVersion: version,
	}
}
