package khr_get_memory_requirements2

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/core1_0"
	"github.com/vkngwrapper/core/driver"
	khr_get_memory_requirements2_driver "github.com/vkngwrapper/extensions/khr_get_memory_requirements2/driver"
	"unsafe"
)

// VulkanExtension is an implementation of the Extension interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanExtension struct {
	driver khr_get_memory_requirements2_driver.Driver
}

// CreateExtensionFromDevice produces an Extension object from a Device with
// khr_get_memory_requirements2 loaded
func CreateExtensionFromDevice(device core1_0.Device) *VulkanExtension {
	if !device.IsDeviceExtensionActive(ExtensionName) {
		return nil
	}

	return &VulkanExtension{
		driver: khr_get_memory_requirements2_driver.CreateDriverFromCore(device.Driver()),
	}
}

// CreateExtensionFromDriver generates an Extension from a driver.Driver object- this is usually
// used in tests to build an Extension from mock drivers
func CreateExtensionFromDriver(driver khr_get_memory_requirements2_driver.Driver) *VulkanExtension {
	return &VulkanExtension{
		driver: driver,
	}
}

func (e *VulkanExtension) BufferMemoryRequirements2(device core1_0.Device, o BufferMemoryRequirementsInfo2, out *MemoryRequirements2) error {
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

	e.driver.VkGetBufferMemoryRequirements2KHR(device.Handle(),
		(*khr_get_memory_requirements2_driver.VkBufferMemoryRequirementsInfo2KHR)(optionPtr),
		(*khr_get_memory_requirements2_driver.VkMemoryRequirements2KHR)(outDataPtr),
	)

	return common.PopulateOutData(out, outDataPtr)
}

func (e *VulkanExtension) ImageMemoryRequirements2(device core1_0.Device, o ImageMemoryRequirementsInfo2, out *MemoryRequirements2) error {
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

	e.driver.VkGetImageMemoryRequirements2KHR(device.Handle(),
		(*khr_get_memory_requirements2_driver.VkImageMemoryRequirementsInfo2KHR)(optionPtr),
		(*khr_get_memory_requirements2_driver.VkMemoryRequirements2KHR)(outDataPtr),
	)

	return common.PopulateOutData(out, outDataPtr)
}

func (e *VulkanExtension) ImageSparseMemoryRequirements2(device core1_0.Device, o ImageSparseMemoryRequirementsInfo2, outDataFactory func() *SparseImageMemoryRequirements2) ([]*SparseImageMemoryRequirements2, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	optionPtr, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, err
	}

	requirementCountPtr := (*driver.Uint32)(arena.Malloc(int(unsafe.Sizeof(C.uint32_t(0)))))

	e.driver.VkGetImageSparseMemoryRequirements2KHR(device.Handle(),
		(*khr_get_memory_requirements2_driver.VkImageSparseMemoryRequirementsInfo2KHR)(optionPtr),
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

	e.driver.VkGetImageSparseMemoryRequirements2KHR(device.Handle(),
		(*khr_get_memory_requirements2_driver.VkImageSparseMemoryRequirementsInfo2KHR)(optionPtr),
		requirementCountPtr,
		(*khr_get_memory_requirements2_driver.VkSparseImageMemoryRequirements2KHR)(unsafe.Pointer(castOutDataPtr)),
	)

	err = common.PopulateOutDataSlice[C.VkSparseImageMemoryRequirements2KHR, *SparseImageMemoryRequirements2](outDataSlice, unsafe.Pointer(outDataPtr))
	if err != nil {
		return nil, err
	}

	return outDataSlice, nil
}
