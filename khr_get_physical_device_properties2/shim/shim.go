package khr_get_physical_device_properties2_shim

//go:generate mockgen -source shim.go -destination ../mocks/shim.go -package mock_get_physical_device_properties2

import (
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/core1_0"
	"github.com/vkngwrapper/core/v2/core1_1"
	"github.com/vkngwrapper/extensions/v3/khr_get_physical_device_properties2"
)

// Shim contains all commands for the khr_get_physical_device_properties2 extension
type Shim interface {
	// Features2 reports capabilities of a PhysicalDevice
	//
	// out - A pre-allocated object in which the results will be populated. It should include any desired
	// chained OutData objects.
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html
	Features2(out *core1_1.PhysicalDeviceFeatures2) error
	// FormatProperties2 lists the PhysicalDevice object's format capabilities
	//
	// format - The format whose properties are queried
	//
	// out - A pre-allocated object in which the results will be populated. It should include any desired
	// chained OutData objects
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFormatProperties2.html
	FormatProperties2(format core1_0.Format, out *core1_1.FormatProperties2) error
	// ImageFormatProperties2 lists the PhysicalDevice object's format capabilities
	//
	// o - Describes the parameters that would be consumed by Device.CreateImage
	//
	// out - A pre-allocated object in which the results will be populated. It should include any desired
	// chained OutData objects
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceImageFormatProperties2.html
	ImageFormatProperties2(o core1_1.PhysicalDeviceImageFormatInfo2, out *core1_1.ImageFormatProperties2) (common.VkResult, error)
	// MemoryProperties2 reports memory information for this PhysicalDevice
	//
	// out - A pre-allocated object in which the results will be populated. It should include any desired
	// chained OutData objects
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceMemoryProperties2.html
	MemoryProperties2(out *core1_1.PhysicalDeviceMemoryProperties2) error
	// Properties2 returns properties of this PhysicalDevice
	//
	// out - A pre-allocated object in which the results will be populated. It should include any desired
	// chained OutData objects
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties2.html
	Properties2(out *core1_1.PhysicalDeviceProperties2) error
	// QueueFamilyProperties2 reports properties of the queues of this PhysicalDevice
	//
	// outDataFactory - This method can be provided to allocate each QueueFamilyProperties2 object
	// that is returned, along with any chained OutData structures. It can also be left nil, in which case
	// QueueFamilyProperties2 will be allocated with no chained structures.
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceQueueFamilyProperties2.html
	QueueFamilyProperties2(outDataFactory func() *core1_1.QueueFamilyProperties2) ([]*core1_1.QueueFamilyProperties2, error)
	// SparseImageFormatProperties2 retrieves properties of an Image format applied to sparse Image
	//
	// o - Contains input parameters
	//
	// outDataFactory - This method can be provided to allocate each SparseImageFormatProperties2 object
	// that is returned, along with any chained OutData structures. It can also be left nil, in which case
	// SparseImageFormatProperties2 will be allocated with no chained structures.
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSparseImageFormatProperties2KHR.html
	SparseImageFormatProperties2(o core1_1.PhysicalDeviceSparseImageFormatInfo2, outDataFactory func() *core1_1.SparseImageFormatProperties2) ([]*core1_1.SparseImageFormatProperties2, error)
}

type VulkanShim struct {
	extension      khr_get_physical_device_properties2.Extension
	physicalDevice core1_0.PhysicalDevice
}

// Compiler check for interface
var _ Shim = &VulkanShim{}

func NewShim(extension khr_get_physical_device_properties2.Extension, device core1_0.PhysicalDevice) *VulkanShim {
	if extension == nil {
		panic("extension cannot be nil")
	}
	if device == nil {
		panic("device cannot be nil")
	}
	return &VulkanShim{
		extension:      extension,
		physicalDevice: device,
	}
}

func (s *VulkanShim) Features2(out *core1_1.PhysicalDeviceFeatures2) error {
	return s.extension.PhysicalDeviceFeatures2(s.physicalDevice, (*khr_get_physical_device_properties2.PhysicalDeviceFeatures2)(out))
}

func (s *VulkanShim) FormatProperties2(format core1_0.Format, out *core1_1.FormatProperties2) error {
	return s.extension.PhysicalDeviceFormatProperties2(s.physicalDevice, format, (*khr_get_physical_device_properties2.FormatProperties2)(out))
}

func (s *VulkanShim) ImageFormatProperties2(o core1_1.PhysicalDeviceImageFormatInfo2, out *core1_1.ImageFormatProperties2) (common.VkResult, error) {
	return s.extension.PhysicalDeviceImageFormatProperties2(
		s.physicalDevice,
		khr_get_physical_device_properties2.PhysicalDeviceImageFormatInfo2(o),
		(*khr_get_physical_device_properties2.ImageFormatProperties2)(out),
	)
}

func (s *VulkanShim) MemoryProperties2(out *core1_1.PhysicalDeviceMemoryProperties2) error {
	return s.extension.PhysicalDeviceMemoryProperties2(s.physicalDevice, (*khr_get_physical_device_properties2.PhysicalDeviceMemoryProperties2)(out))
}

func (s *VulkanShim) Properties2(out *core1_1.PhysicalDeviceProperties2) error {
	return s.extension.PhysicalDeviceProperties2(s.physicalDevice, (*khr_get_physical_device_properties2.PhysicalDeviceProperties2)(out))
}

func (s *VulkanShim) QueueFamilyProperties2(outDataFactory func() *core1_1.QueueFamilyProperties2) ([]*core1_1.QueueFamilyProperties2, error) {
	factory := func() *khr_get_physical_device_properties2.QueueFamilyProperties2 {
		return (*khr_get_physical_device_properties2.QueueFamilyProperties2)(outDataFactory())
	}
	properties, err := s.extension.PhysicalDeviceQueueFamilyProperties2(s.physicalDevice, factory)
	if err != nil {
		return nil, err
	}

	var retVal []*core1_1.QueueFamilyProperties2
	for _, property := range properties {
		retVal = append(retVal, (*core1_1.QueueFamilyProperties2)(property))
	}

	return retVal, nil
}

func (s *VulkanShim) SparseImageFormatProperties2(o core1_1.PhysicalDeviceSparseImageFormatInfo2, outDataFactory func() *core1_1.SparseImageFormatProperties2) ([]*core1_1.SparseImageFormatProperties2, error) {
	factory := func() *khr_get_physical_device_properties2.SparseImageFormatProperties2 {
		return (*khr_get_physical_device_properties2.SparseImageFormatProperties2)(outDataFactory())
	}

	properties, err := s.extension.PhysicalDeviceSparseImageFormatProperties2(
		s.physicalDevice,
		khr_get_physical_device_properties2.PhysicalDeviceSparseImageFormatInfo2(o),
		factory,
	)
	if err != nil {
		return nil, err
	}

	var retVal []*core1_1.SparseImageFormatProperties2
	for _, property := range properties {
		retVal = append(retVal, (*core1_1.SparseImageFormatProperties2)(property))
	}

	return retVal, nil
}
