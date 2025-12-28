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

// VulkanExtensionDriver is an implementation of the ExtensionDriver interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanExtensionDriver struct {
	loader   khr_device_group_creation_driver.Loader
	instance core.Instance
}

// CreateExtensionDriverFromCoreDriver produces an ExtensionDriver object from an Instance with
// khr_device_group_creation loaded
func CreateExtensionDriverFromCoreDriver(coreDriver core1_0.CoreInstanceDriver) *VulkanExtensionDriver {
	instance := coreDriver.Instance()
	if !instance.IsInstanceExtensionActive(ExtensionName) {
		return nil
	}

	return &VulkanExtensionDriver{
		loader:   khr_device_group_creation_driver.CreateLoaderFromCore(coreDriver.Loader()),
		instance: instance,
	}
}

// CreateExtensionDriverFromLoader generates an ExtensionDriver from a loader.Loader object- this is usually
// used in tests to build an ExtensionDriver from mock drivers
func CreateExtensionDriverFromLoader(driver khr_device_group_creation_driver.Loader, instance core.Instance) *VulkanExtensionDriver {
	return &VulkanExtensionDriver{
		loader:   driver,
		instance: instance,
	}
}

func (e *VulkanExtensionDriver) attemptEnumeratePhysicalDeviceGroups(outDataFactory func() *PhysicalDeviceGroupProperties) ([]*PhysicalDeviceGroupProperties, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	countPtr := (*loader.Uint32)(arena.Malloc(int(unsafe.Sizeof(C.uint32_t(0)))))

	res, err := e.loader.VkEnumeratePhysicalDeviceGroupsKHR(
		e.instance.Handle(),
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

	res, err = e.loader.VkEnumeratePhysicalDeviceGroupsKHR(
		e.instance.Handle(),
		countPtr,
		(*khr_device_group_creation_driver.VkPhysicalDeviceGroupPropertiesKHR)(unsafe.Pointer(outData)),
	)
	if err != nil {
		return nil, res, err
	}

	err = common.PopulateOutDataSlice[C.VkPhysicalDeviceGroupPropertiesKHR, *PhysicalDeviceGroupProperties](outDataSlice, unsafe.Pointer(outData), e.instance, e.instance.APIVersion(), e.loader.CoreLoader())
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	return outDataSlice, res, nil
}

func (e *VulkanExtensionDriver) EnumeratePhysicalDeviceGroups(outDataFactory func() *PhysicalDeviceGroupProperties) ([]*PhysicalDeviceGroupProperties, common.VkResult, error) {
	var outData []*PhysicalDeviceGroupProperties
	var result common.VkResult
	var err error

	for doWhile := true; doWhile; doWhile = (result == core1_0.VKIncomplete) {
		outData, result, err = e.attemptEnumeratePhysicalDeviceGroups(outDataFactory)
	}
	return outData, result, err
}
