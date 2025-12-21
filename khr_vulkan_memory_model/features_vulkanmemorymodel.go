package khr_vulkan_memory_model

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3/common"
)

// PhysicalDeviceVulkanMemoryModelFeatures describes features supported by the memory model
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceVulkanMemoryModelFeatures.html
type PhysicalDeviceVulkanMemoryModelFeatures struct {
	// VulkanMemoryModel indicates whether the Vulkan Memory Model is supported
	VulkanMemoryModel bool
	// VulkanMemoryModelDeviceScope indicates whether the Vulkan Memory Model can use Device
	// scope synchronization
	VulkanMemoryModelDeviceScope bool
	// VulkanMemoryModelAvailabilityVisibilityChains indicates whether the Vulkan Memory Model
	// can use available and visibility chains with more than one element
	VulkanMemoryModelAvailabilityVisibilityChains bool

	common.NextOptions
	common.NextOutData
}

func (o *PhysicalDeviceVulkanMemoryModelFeatures) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceVulkanMemoryModelFeaturesKHR{})))
	}

	info := (*C.VkPhysicalDeviceVulkanMemoryModelFeaturesKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VULKAN_MEMORY_MODEL_FEATURES_KHR
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceVulkanMemoryModelFeatures) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceVulkanMemoryModelFeaturesKHR)(cDataPointer)

	o.VulkanMemoryModel = info.vulkanMemoryModel != C.VkBool32(0)
	o.VulkanMemoryModelDeviceScope = info.vulkanMemoryModelDeviceScope != C.VkBool32(0)
	o.VulkanMemoryModelAvailabilityVisibilityChains = info.vulkanMemoryModelAvailabilityVisibilityChains != C.VkBool32(0)

	return info.pNext, nil
}

func (o PhysicalDeviceVulkanMemoryModelFeatures) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceVulkanMemoryModelFeaturesKHR{})))
	}

	info := (*C.VkPhysicalDeviceVulkanMemoryModelFeaturesKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VULKAN_MEMORY_MODEL_FEATURES_KHR
	info.pNext = next
	info.vulkanMemoryModel = C.VkBool32(0)
	info.vulkanMemoryModelDeviceScope = C.VkBool32(0)
	info.vulkanMemoryModelAvailabilityVisibilityChains = C.VkBool32(0)

	if o.VulkanMemoryModel {
		info.vulkanMemoryModel = C.VkBool32(1)
	}

	if o.VulkanMemoryModelDeviceScope {
		info.vulkanMemoryModelDeviceScope = C.VkBool32(1)
	}

	if o.VulkanMemoryModelAvailabilityVisibilityChains {
		info.vulkanMemoryModelAvailabilityVisibilityChains = C.VkBool32(1)
	}

	return preallocatedPointer, nil
}
