package khr_get_memory_requirements2

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
	"github.com/vkngwrapper/extensions/v3/khr_get_memory_requirements2/loader"
)

// VulkanExtensionDriver is an implementation of the ExtensionDriver interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanExtensionDriver struct {
	driver khr_get_memory_requirements2_loader.Loader
	device core.Device
}

// CreateExtensionDriverFromCoreDriver produces an ExtensionDriver object from a Device with
// khr_get_memory_requirements2 loaded
func CreateExtensionDriverFromCoreDriver(driver core1_0.DeviceDriver) *VulkanExtensionDriver {
	device := driver.Device()
	if !device.IsDeviceExtensionActive(ExtensionName) {
		return nil
	}

	return &VulkanExtensionDriver{
		driver: khr_get_memory_requirements2_loader.CreateLoaderFromCore(driver.Loader()),
		device: device,
	}
}

// CreateExtensionDriverFromLoader generates an ExtensionDriver from a loader.Loader object- this is usually
// used in tests to build an ExtensionDriver from mock drivers
func CreateExtensionDriverFromLoader(driver khr_get_memory_requirements2_loader.Loader, device core.Device) *VulkanExtensionDriver {
	return &VulkanExtensionDriver{
		driver: driver,
		device: device,
	}
}

func (e *VulkanExtensionDriver) GetBufferMemoryRequirements2(o BufferMemoryRequirementsInfo2, out *MemoryRequirements2) error {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	optionPtr, err := common.AllocOptions(arena, o)
	if err != nil {
		return err
	}

	outDataPtr, err := common.AllocOutDataHeader(arena, out)
	if err != nil {
		return err
	}

	e.driver.VkGetBufferMemoryRequirements2KHR(e.device.Handle(),
		(*khr_get_memory_requirements2_loader.VkBufferMemoryRequirementsInfo2KHR)(optionPtr),
		(*khr_get_memory_requirements2_loader.VkMemoryRequirements2KHR)(outDataPtr),
	)

	return common.PopulateOutData(out, outDataPtr)
}

func (e *VulkanExtensionDriver) GetImageMemoryRequirements2(o ImageMemoryRequirementsInfo2, out *MemoryRequirements2) error {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	optionPtr, err := common.AllocOptions(arena, o)
	if err != nil {
		return err
	}

	outDataPtr, err := common.AllocOutDataHeader(arena, out)
	if err != nil {
		return err
	}

	e.driver.VkGetImageMemoryRequirements2KHR(e.device.Handle(),
		(*khr_get_memory_requirements2_loader.VkImageMemoryRequirementsInfo2KHR)(optionPtr),
		(*khr_get_memory_requirements2_loader.VkMemoryRequirements2KHR)(outDataPtr),
	)

	return common.PopulateOutData(out, outDataPtr)
}

func (e *VulkanExtensionDriver) GetImageSparseMemoryRequirements2(o ImageSparseMemoryRequirementsInfo2, outDataFactory func() *SparseImageMemoryRequirements2) ([]*SparseImageMemoryRequirements2, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	optionPtr, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, err
	}

	requirementCountPtr := (*loader.Uint32)(arena.Malloc(int(unsafe.Sizeof(C.uint32_t(0)))))

	e.driver.VkGetImageSparseMemoryRequirements2KHR(e.device.Handle(),
		(*khr_get_memory_requirements2_loader.VkImageSparseMemoryRequirementsInfo2KHR)(optionPtr),
		requirementCountPtr,
		nil,
	)

	count := int(*requirementCountPtr)
	if count == 0 {
		return nil, nil
	}

	outDataSlice := make([]*SparseImageMemoryRequirements2, count)
	for i := 0; i < count; i++ {
		if outDataFactory != nil {
			outDataSlice[i] = outDataFactory()
		} else {
			outDataSlice[i] = &SparseImageMemoryRequirements2{}
		}
	}

	outDataPtr, err := common.AllocOutDataHeaderSlice[C.VkSparseImageMemoryRequirements2KHR, *SparseImageMemoryRequirements2](arena, outDataSlice)
	if err != nil {
		return nil, err
	}

	castOutDataPtr := (*C.VkSparseImageMemoryRequirements2KHR)(outDataPtr)

	e.driver.VkGetImageSparseMemoryRequirements2KHR(e.device.Handle(),
		(*khr_get_memory_requirements2_loader.VkImageSparseMemoryRequirementsInfo2KHR)(optionPtr),
		requirementCountPtr,
		(*khr_get_memory_requirements2_loader.VkSparseImageMemoryRequirements2KHR)(unsafe.Pointer(castOutDataPtr)),
	)

	err = common.PopulateOutDataSlice[C.VkSparseImageMemoryRequirements2KHR, *SparseImageMemoryRequirements2](outDataSlice, unsafe.Pointer(outDataPtr))
	if err != nil {
		return nil, err
	}

	return outDataSlice, nil
}
