package khr_device_group_creation

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
	khr_device_group_creation_driver "github.com/vkngwrapper/extensions/v3/khr_device_group_creation/loader"
)

// VulkanExtension is an implementation of the Extension interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanExtension struct {
	driver khr_device_group_creation_driver.Loader
}

// CreateExtensionFromInstance produces an Extension object from an Instance with
// khr_device_group_creation loaded
func CreateExtensionFromInstance(instance core.Instance) *VulkanExtension {
	if !instance.IsInstanceExtensionActive(ExtensionName) {
		return nil
	}

	return &VulkanExtension{
		driver: khr_device_group_creation_driver.CreateLoaderFromCore(instance.Driver()),
	}
}

// CreateExtensionFromDriver generates an Extension from a loader.Loader object- this is usually
// used in tests to build an Extension from mock drivers
func CreateExtensionFromDriver(driver khr_device_group_creation_driver.Loader) *VulkanExtension {
	return &VulkanExtension{
		driver: driver,
	}
}

func (e *VulkanExtension) attemptEnumeratePhysicalDeviceGroups(instance core.Instance, outDataFactory func() *PhysicalDeviceGroupProperties) ([]*PhysicalDeviceGroupProperties, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	countPtr := (*loader.Uint32)(arena.Malloc(int(unsafe.Sizeof(C.uint32_t(0)))))

	res, err := e.driver.VkEnumeratePhysicalDeviceGroupsKHR(
		instance.Handle(),
		countPtr,
		nil,
	)
	if err != nil {
		return nil, res, err
	}

	count := int(*countPtr)
	if count == 0 {
		return nil, core1_0.VKSuccess, nil
	}

	outDataSlice := make([]*PhysicalDeviceGroupProperties, count)
	for i := 0; i < count; i++ {
		if outDataFactory != nil {
			outDataSlice[i] = outDataFactory()
		} else {
			outDataSlice[i] = &PhysicalDeviceGroupProperties{}
		}
	}

	outData, err := common.AllocOutDataHeaderSlice[C.VkPhysicalDeviceGroupPropertiesKHR, *PhysicalDeviceGroupProperties](arena, outDataSlice)
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	res, err = e.driver.VkEnumeratePhysicalDeviceGroupsKHR(
		instance.Handle(),
		countPtr,
		(*khr_device_group_creation_driver.VkPhysicalDeviceGroupPropertiesKHR)(unsafe.Pointer(outData)),
	)
	if err != nil {
		return nil, res, err
	}

	err = common.PopulateOutDataSlice[C.VkPhysicalDeviceGroupPropertiesKHR, *PhysicalDeviceGroupProperties](outDataSlice, unsafe.Pointer(outData), instance)
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	return outDataSlice, res, nil
}

func (e *VulkanExtension) EnumeratePhysicalDeviceGroups(instance core.Instance, outDataFactory func() *PhysicalDeviceGroupProperties) ([]*PhysicalDeviceGroupProperties, common.VkResult, error) {
	if instance.Handle() == 0 {
		panic("instance cannot be uninitialized")
	}
	var outData []*PhysicalDeviceGroupProperties
	var result common.VkResult
	var err error

	for doWhile := true; doWhile; doWhile = (result == core1_0.VKIncomplete) {
		outData, result, err = e.attemptEnumeratePhysicalDeviceGroups(instance, outDataFactory)
	}
	return outData, result, err
}
