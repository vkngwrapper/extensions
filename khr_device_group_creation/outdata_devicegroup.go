package khr_device_group_creation

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/pkg/errors"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/driver"
)

// PhysicalDeviceGroupProperties specifies PhysicalDevice group properties
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceGroupProperties.html
type PhysicalDeviceGroupProperties struct {
	// PhysicalDevices is a slice of PhysicalDevice objects that represent all PhysicalDevice
	// objects in the group
	PhysicalDevices []core1_0.PhysicalDevice
	// SubsetAllocation specifies whether logical Device objects created from the group support
	// allocating DeviceMemory on a subset of Device objects, via MemoryAllocateFlagsInfo
	SubsetAllocation bool

	common.NextOutData
}

func (o *PhysicalDeviceGroupProperties) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceGroupPropertiesKHR{})))
	}

	createInfo := (*C.VkPhysicalDeviceGroupPropertiesKHR)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_GROUP_PROPERTIES_KHR
	createInfo.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceGroupProperties) PopulateOutData(cPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo := (*C.VkPhysicalDeviceGroupPropertiesKHR)(cPointer)
	o.SubsetAllocation = createInfo.subsetAllocation != C.VkBool32(0)

	instance, ok := common.OfType[core1_0.Instance](helpers)
	if !ok {
		return nil, errors.New("outdata population requires an Instance passed to populate helpers")
	}
	builder, ok := common.OfType[core1_0.InstanceObjectBuilder](helpers)
	if !ok {
		return nil, errors.New("outdata population requires an InstanceObjectBuilder passed to populate helpers")
	}

	count := int(createInfo.physicalDeviceCount)
	o.PhysicalDevices = make([]core1_0.PhysicalDevice, count)

	propertiesUnsafe := arena.Malloc(int(unsafe.Sizeof([1]C.VkPhysicalDeviceProperties{})))

	for i := 0; i < count; i++ {
		handle := driver.VkPhysicalDevice(unsafe.Pointer(createInfo.physicalDevices[i]))
		instance.Driver().VkGetPhysicalDeviceProperties(handle, (*driver.VkPhysicalDeviceProperties)(propertiesUnsafe))

		var properties core1_0.PhysicalDeviceProperties
		err = (&properties).PopulateFromCPointer(propertiesUnsafe)
		if err != nil {
			return nil, err
		}

		deviceVersion := instance.APIVersion().Min(properties.APIVersion)

		o.PhysicalDevices[i] = builder.CreatePhysicalDeviceObject(instance.Driver(), instance.Handle(), handle, instance.APIVersion(), deviceVersion)
	}

	return createInfo.pNext, nil
}
