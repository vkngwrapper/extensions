package khr_maintenance4

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/core1_0"
	"github.com/vkngwrapper/core/v2/core1_1"
	"github.com/vkngwrapper/core/v2/driver"
	"github.com/vkngwrapper/extensions/v3/khr_maintenance4/driver"
)

// VulkanExtension is an implementation of the Extension interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanExtension struct {
	driver khr_maintenance4_driver.Driver
}

// CreateExtensionFromDevice produces an Extension object from a Device with
// khr_maintenance4 loaded
func CreateExtensionFromDevice(device core1_0.Device) *VulkanExtension {
	if !device.IsDeviceExtensionActive(ExtensionName) {
		return nil
	}

	return &VulkanExtension{
		driver: khr_maintenance4_driver.CreateDriverFromCore(device.Driver()),
	}
}

// CreateExtensionFromDriver generates an Extension from a driver.Driver object- this is usually
// used in tests to build an Extension from mock drivers
func CreateExtensionFromDriver(driver khr_maintenance4_driver.Driver) *VulkanExtension {
	return &VulkanExtension{
		driver: driver,
	}
}

func (e *VulkanExtension) DeviceBufferMemoryRequirements(device core1_0.Device, options DeviceBufferMemoryRequirements, outData *core1_1.MemoryRequirements2) error {
	if device == nil {
		panic("device cannot be nil")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	optionsPtr, err := common.AllocOptions(arena, options)
	if err != nil {
		return err
	}

	outDataPtr, err := common.AllocOutDataHeader(arena, outData)
	if err != nil {
		return err
	}

	e.driver.VkGetDeviceBufferMemoryRequirementsKHR(device.Handle(), (*khr_maintenance4_driver.VkDeviceBufferMemoryRequirementsKHR)(optionsPtr), (*driver.VkMemoryRequirements2)(outDataPtr))

	return common.PopulateOutData(outData, outDataPtr)
}

func (e *VulkanExtension) DeviceImageMemoryRequirements(device core1_0.Device, options DeviceImageMemoryRequirements, outData *core1_1.MemoryRequirements2) error {
	if device == nil {
		panic("device cannot be nil")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	optionsPtr, err := common.AllocOptions(arena, options)
	if err != nil {
		return err
	}

	outDataPtr, err := common.AllocOutDataHeader(arena, outData)
	if err != nil {
		return err
	}

	e.driver.VkGetDeviceImageMemoryRequirementsKHR(device.Handle(), (*khr_maintenance4_driver.VkDeviceImageMemoryRequirementsKHR)(optionsPtr), (*driver.VkMemoryRequirements2)(outDataPtr))

	return common.PopulateOutData(outData, outDataPtr)
}

func (e *VulkanExtension) DeviceImageSparseMemoryRequirements(device core1_0.Device, options DeviceImageMemoryRequirements, outDataFactory func() *core1_1.SparseImageMemoryRequirements2) ([]*core1_1.SparseImageMemoryRequirements2, error) {
	if device == nil {
		panic("device cannot be nil")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	optionsPtr, err := common.AllocOptions(arena, options)
	if err != nil {
		return nil, err
	}
	sparseCount := (*driver.Uint32)(arena.Malloc(int(unsafe.Sizeof(C.uint32_t(0)))))

	e.driver.VkGetDeviceImageSparseMemoryRequirementsKHR(device.Handle(), (*khr_maintenance4_driver.VkDeviceImageMemoryRequirementsKHR)(optionsPtr), sparseCount, nil)

	count := int(*sparseCount)
	if count == 0 {
		return nil, nil
	}

	outDataSlice := make([]*core1_1.SparseImageMemoryRequirements2, count)
	for i := 0; i < count; i++ {
		if outDataFactory != nil {
			outDataSlice[i] = outDataFactory()
		} else {
			outDataSlice[i] = &core1_1.SparseImageMemoryRequirements2{}
		}
	}

	outDataPtr, err := common.AllocOutDataHeaderSlice[C.VkSparseImageMemoryRequirements2, *core1_1.SparseImageMemoryRequirements2](arena, outDataSlice)
	if err != nil {
		return nil, err
	}

	e.driver.VkGetDeviceImageSparseMemoryRequirementsKHR(device.Handle(), (*khr_maintenance4_driver.VkDeviceImageMemoryRequirementsKHR)(optionsPtr), sparseCount, (*driver.VkSparseImageMemoryRequirements2)(unsafe.Pointer(outDataPtr)))

	err = common.PopulateOutDataSlice[C.VkSparseImageMemoryRequirements2, *core1_1.SparseImageMemoryRequirements2](outDataSlice, unsafe.Pointer(outDataPtr))
	if err != nil {
		return nil, err
	}

	return outDataSlice, nil
}

var _ Extension = &VulkanExtension{}
