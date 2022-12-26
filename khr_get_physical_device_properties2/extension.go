package khr_get_physical_device_properties2

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/core1_0"
	"github.com/vkngwrapper/core/v2/driver"
	ext_driver "github.com/vkngwrapper/extensions/v2/khr_get_physical_device_properties2/driver"
	"unsafe"
)

// VulkanExtension is an implementation of the Extension interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanExtension struct {
	driver ext_driver.Driver
}

// CreateExtensionFromInstance produces an Extension object from an Instance with
// khr_get_physical_device_properties2 loaded
func CreateExtensionFromInstance(instance core1_0.Instance) *VulkanExtension {
	if !instance.IsInstanceExtensionActive(ExtensionName) {
		return nil
	}

	return &VulkanExtension{
		driver: ext_driver.CreateDriverFromCore(instance.Driver()),
	}
}

// CreateExtensionFromDriver generates an Extension from a driver.Driver object- this is usually
// used in tests to build an Extension from mock drivers
func CreateExtensionFromDriver(driver ext_driver.Driver) *VulkanExtension {
	return &VulkanExtension{
		driver: driver,
	}
}

func (e *VulkanExtension) PhysicalDeviceFeatures2(physicalDevice core1_0.PhysicalDevice, out *PhysicalDeviceFeatures2) error {
	if physicalDevice == nil {
		panic("physicalDevice cannot be nil")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	outData, err := common.AllocOutDataHeader(arena, out)
	if err != nil {
		return err
	}

	e.driver.VkGetPhysicalDeviceFeatures2KHR(physicalDevice.Handle(), (*ext_driver.VkPhysicalDeviceFeatures2KHR)(outData))

	return common.PopulateOutData(out, outData)
}

func (e *VulkanExtension) PhysicalDeviceFormatProperties2(physicalDevice core1_0.PhysicalDevice, format core1_0.Format, out *FormatProperties2) error {
	if physicalDevice == nil {
		panic("physicalDevice cannot be nil")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	outData, err := common.AllocOutDataHeader(arena, out)
	if err != nil {
		return err
	}

	e.driver.VkGetPhysicalDeviceFormatProperties2KHR(physicalDevice.Handle(), driver.VkFormat(format), (*ext_driver.VkFormatProperties2KHR)(outData))

	return common.PopulateOutData(out, outData)
}

func (e *VulkanExtension) PhysicalDeviceImageFormatProperties2(physicalDevice core1_0.PhysicalDevice, options PhysicalDeviceImageFormatInfo2, out *ImageFormatProperties2) (common.VkResult, error) {
	if physicalDevice == nil {
		panic("physicalDevice cannot be nil")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	optionData, err := common.AllocOptions(arena, options)
	if err != nil {
		return core1_0.VKErrorUnknown, err
	}

	outData, err := common.AllocOutDataHeader(arena, out)
	if err != nil {
		return core1_0.VKErrorUnknown, err
	}

	res, err := e.driver.VkGetPhysicalDeviceImageFormatProperties2KHR(physicalDevice.Handle(), (*ext_driver.VkPhysicalDeviceImageFormatInfo2KHR)(optionData), (*ext_driver.VkImageFormatProperties2KHR)(outData))
	if err != nil {
		return res, err
	}

	err = common.PopulateOutData(out, outData)
	if err != nil {
		return core1_0.VKErrorUnknown, err
	}

	return res, nil
}

func (e *VulkanExtension) PhysicalDeviceMemoryProperties2(physicalDevice core1_0.PhysicalDevice, out *PhysicalDeviceMemoryProperties2) error {
	if physicalDevice == nil {
		panic("physicalDevice cannot be nil")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	outData, err := common.AllocOutDataHeader(arena, out)
	if err != nil {
		return err
	}

	e.driver.VkGetPhysicalDeviceMemoryProperties2KHR(physicalDevice.Handle(), (*ext_driver.VkPhysicalDeviceMemoryProperties2KHR)(outData))

	return common.PopulateOutData(out, outData)
}

func (e *VulkanExtension) PhysicalDeviceProperties2(physicalDevice core1_0.PhysicalDevice, out *PhysicalDeviceProperties2) error {
	if physicalDevice == nil {
		panic("physicalDevice cannot be nil")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	outData, err := common.AllocOutDataHeader(arena, out)
	if err != nil {
		return err
	}

	e.driver.VkGetPhysicalDeviceProperties2KHR(physicalDevice.Handle(), (*ext_driver.VkPhysicalDeviceProperties2KHR)(outData))

	return common.PopulateOutData(out, outData)
}

func (e *VulkanExtension) PhysicalDeviceQueueFamilyProperties2(physicalDevice core1_0.PhysicalDevice, outDataFactory func() *QueueFamilyProperties2) ([]*QueueFamilyProperties2, error) {
	if physicalDevice == nil {
		panic("physicalDevice cannot be nil")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	outDataCountPtr := (*driver.Uint32)(arena.Malloc(int(unsafe.Sizeof(C.uint32_t(0)))))

	e.driver.VkGetPhysicalDeviceQueueFamilyProperties2KHR(physicalDevice.Handle(), outDataCountPtr, nil)

	outDataCount := int(*outDataCountPtr)
	if outDataCount == 0 {
		return nil, nil
	}

	out := make([]*QueueFamilyProperties2, outDataCount)
	for i := 0; i < outDataCount; i++ {
		if outDataFactory != nil {
			out[i] = outDataFactory()
		} else {
			out[i] = &QueueFamilyProperties2{}
		}
	}

	outData, err := common.AllocOutDataHeaderSlice[C.VkQueueFamilyProperties2KHR, *QueueFamilyProperties2](arena, out)
	if err != nil {
		return nil, err
	}

	e.driver.VkGetPhysicalDeviceQueueFamilyProperties2KHR(physicalDevice.Handle(), outDataCountPtr, (*ext_driver.VkQueueFamilyProperties2KHR)(unsafe.Pointer(outData)))

	err = common.PopulateOutDataSlice[C.VkQueueFamilyProperties2KHR, *QueueFamilyProperties2](out, unsafe.Pointer(outData))
	return out, err
}

func (e *VulkanExtension) PhysicalDeviceSparseImageFormatProperties2(physicalDevice core1_0.PhysicalDevice, options PhysicalDeviceSparseImageFormatInfo2, outDataFactory func() *SparseImageFormatProperties2) ([]*SparseImageFormatProperties2, error) {
	if physicalDevice == nil {
		panic("physicalDevice cannot be nil")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	outDataCountPtr := (*driver.Uint32)(arena.Malloc(int(unsafe.Sizeof(C.uint32_t(0)))))
	optionData, err := common.AllocOptions(arena, options)
	if err != nil {
		return nil, err
	}

	e.driver.VkGetPhysicalDeviceSparseImageFormatProperties2KHR(physicalDevice.Handle(), (*ext_driver.VkPhysicalDeviceSparseImageFormatInfo2KHR)(optionData), outDataCountPtr, nil)

	outDataCount := int(*outDataCountPtr)
	if outDataCount == 0 {
		return nil, nil
	}

	out := make([]*SparseImageFormatProperties2, outDataCount)
	for i := 0; i < outDataCount; i++ {
		if outDataFactory != nil {
			out[i] = outDataFactory()
		} else {
			out[i] = &SparseImageFormatProperties2{}
		}
	}

	outData, err := common.AllocOutDataHeaderSlice[C.VkSparseImageFormatProperties2KHR, *SparseImageFormatProperties2](arena, out)
	if err != nil {
		return nil, err
	}

	e.driver.VkGetPhysicalDeviceSparseImageFormatProperties2KHR(physicalDevice.Handle(), (*ext_driver.VkPhysicalDeviceSparseImageFormatInfo2KHR)(optionData), outDataCountPtr, (*ext_driver.VkSparseImageFormatProperties2KHR)(unsafe.Pointer(outData)))

	err = common.PopulateOutDataSlice[C.VkSparseImageFormatProperties2KHR, *SparseImageFormatProperties2](out, unsafe.Pointer(outData))

	return out, err
}

var _ Extension = &VulkanExtension{}
