package khr_get_physical_device_properties2

//go:generate mockgen -source extiface.go -destination ./mocks/extension.go -package mock_get_physical_device_properties2

import (
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/core1_0"
)

// Extension contains all commands for the khr_get_physical_device_properties2 extension
type Extension interface {
	// PhysicalDeviceFeatures2 reports capabilities of a PhysicalDevice
	//
	// physicalDevice - The PhysicalDevice whose features are being queried
	//
	// out - A pre-allocated object in which the results will be populated. It should include any desired
	// chained OutData objects.
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html
	PhysicalDeviceFeatures2(physicalDevice core1_0.PhysicalDevice, out *PhysicalDeviceFeatures2) error
	// PhysicalDeviceFormatProperties2 lists the PhysicalDevice object's format capabilities
	//
	// physicalDevice - The PhysicalDevice whose format properties are being queried
	//
	// format - The format whose properties are queried
	//
	// out - A pre-allocated object in which the results will be populated. It should include any desired
	// chained OutData objects
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFormatProperties2.html
	PhysicalDeviceFormatProperties2(physicalDevice core1_0.PhysicalDevice, format core1_0.Format, out *FormatProperties2) error
	// PhysicalDeviceImageFormatProperties2 lists the PhysicalDevice object's format capabilities
	//
	// physicalDevice - The PhysicalDevice whose image format properties are being queried
	//
	// options - Describes the parameters that would be consumed by Device.CreateImage
	//
	// out - A pre-allocated object in which the results will be populated. It should include any desired
	// chained OutData objects
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceImageFormatProperties2.html
	PhysicalDeviceImageFormatProperties2(physicalDevice core1_0.PhysicalDevice, options PhysicalDeviceImageFormatInfo2, out *ImageFormatProperties2) (common.VkResult, error)
	// PhysicalDeviceMemoryProperties2 reports memory information for this PhysicalDevice
	//
	// physicalDevice - The PhysicalDevice whose memory properties are being queried
	//
	// out - A pre-allocated object in which the results will be populated. It should include any desired
	// chained OutData objects
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceMemoryProperties2.html
	PhysicalDeviceMemoryProperties2(physicalDevice core1_0.PhysicalDevice, out *PhysicalDeviceMemoryProperties2) error
	// PhysicalDeviceProperties2 returns properties of this PhysicalDevice
	//
	// physicalDevice - The PhysicalDevice whose properties are being queried
	//
	// out - A pre-allocated object in which the results will be populated. It should include any desired
	// chained OutData objects
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties2.html
	PhysicalDeviceProperties2(physicalDevice core1_0.PhysicalDevice, out *PhysicalDeviceProperties2) error
	// PhysicalDeviceQueueFamilyProperties2 reports properties of the queues of this PhysicalDevice
	//
	// physicalDevice - The PhysicalDevice whose queue family properties are being queried
	//
	// outDataFactory - This method can be provided to allocate each QueueFamilyProperties2 object
	// that is returned, along with any chained OutData structures. It can also be left nil, in which case
	// QueueFamilyProperties2 will be allocated with no chained structures.
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceQueueFamilyProperties2.html
	PhysicalDeviceQueueFamilyProperties2(physicalDevice core1_0.PhysicalDevice, outDataFactory func() *QueueFamilyProperties2) ([]*QueueFamilyProperties2, error)
	// PhysicalDeviceSparseImageFormatProperties2 retrieves properties of an Image format applied to sparse Image
	//
	// physicalDevice - The PhysicalDevice whose sparse image format properties are being queried
	//
	// options - Contains input parameters
	//
	// outDataFactory - This method can be provided to allocate each SparseImageFormatProperties2 object
	// that is returned, along with any chained OutData structures. It can also be left nil, in which case
	// SparseImageFormatProperties2 will be allocated with no chained structures.
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSparseImageFormatProperties2KHR.html
	PhysicalDeviceSparseImageFormatProperties2(physicalDevice core1_0.PhysicalDevice, options PhysicalDeviceSparseImageFormatInfo2, outDataFactory func() *SparseImageFormatProperties2) ([]*SparseImageFormatProperties2, error)
}
