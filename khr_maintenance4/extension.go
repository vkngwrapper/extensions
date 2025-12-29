package khr_maintenance4

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
	"github.com/vkngwrapper/core/v3/core1_1"
	"github.com/vkngwrapper/core/v3/loader"
	"github.com/vkngwrapper/extensions/v3/khr_maintenance4/loader"
)

// VulkanExtensionDriver is an implementation of the ExtensionDriver interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanExtensionDriver struct {
	driver khr_maintenance4_loader.Loader
	device core.Device
}

// CreateExtensionDriverFromCoreDriver produces an ExtensionDriver object from a Device with
// khr_maintenance4 loaded
func CreateExtensionDriverFromCoreDriver(driver core1_0.DeviceDriver) ExtensionDriver {
	device := driver.Device()
	if !device.IsDeviceExtensionActive(ExtensionName) {
		return nil
	}

	return &VulkanExtensionDriver{
		driver: khr_maintenance4_loader.CreateLoaderFromCore(driver.Loader()),
		device: device,
	}
}

// CreateExtensionDriverFromLoader generates an ExtensionDriver from a loader.Loader object- this is usually
// used in tests to build an ExtensionDriver from mock drivers
func CreateExtensionDriverFromLoader(driver khr_maintenance4_loader.Loader, device core.Device) *VulkanExtensionDriver {
	return &VulkanExtensionDriver{
		driver: driver,
		device: device,
	}
}

func (e *VulkanExtensionDriver) GetDeviceBufferMemoryRequirements(options DeviceBufferMemoryRequirements, outData *core1_1.MemoryRequirements2) error {
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

	e.driver.VkGetDeviceBufferMemoryRequirementsKHR(e.device.Handle(), (*khr_maintenance4_loader.VkDeviceBufferMemoryRequirementsKHR)(optionsPtr), (*loader.VkMemoryRequirements2)(outDataPtr))

	return common.PopulateOutData(outData, outDataPtr)
}

func (e *VulkanExtensionDriver) GetDeviceImageMemoryRequirements(options DeviceImageMemoryRequirements, outData *core1_1.MemoryRequirements2) error {
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

	e.driver.VkGetDeviceImageMemoryRequirementsKHR(e.device.Handle(), (*khr_maintenance4_loader.VkDeviceImageMemoryRequirementsKHR)(optionsPtr), (*loader.VkMemoryRequirements2)(outDataPtr))

	return common.PopulateOutData(outData, outDataPtr)
}

func (e *VulkanExtensionDriver) GetDeviceImageSparseMemoryRequirements(options DeviceImageMemoryRequirements, outDataFactory func() *core1_1.SparseImageMemoryRequirements2) ([]*core1_1.SparseImageMemoryRequirements2, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	optionsPtr, err := common.AllocOptions(arena, options)
	if err != nil {
		return nil, err
	}
	sparseCount := (*loader.Uint32)(arena.Malloc(int(unsafe.Sizeof(C.uint32_t(0)))))

	e.driver.VkGetDeviceImageSparseMemoryRequirementsKHR(e.device.Handle(), (*khr_maintenance4_loader.VkDeviceImageMemoryRequirementsKHR)(optionsPtr), sparseCount, nil)

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

	e.driver.VkGetDeviceImageSparseMemoryRequirementsKHR(e.device.Handle(), (*khr_maintenance4_loader.VkDeviceImageMemoryRequirementsKHR)(optionsPtr), sparseCount, (*loader.VkSparseImageMemoryRequirements2)(unsafe.Pointer(outDataPtr)))

	err = common.PopulateOutDataSlice[C.VkSparseImageMemoryRequirements2, *core1_1.SparseImageMemoryRequirements2](outDataSlice, unsafe.Pointer(outDataPtr))
	if err != nil {
		return nil, err
	}

	return outDataSlice, nil
}

var _ ExtensionDriver = &VulkanExtensionDriver{}
