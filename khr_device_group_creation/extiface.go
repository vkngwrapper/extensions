package khr_device_group_creation

import (
	"github.com/vkngwrapper/core/v3/common"
)

//go:generate mockgen -source extiface.go -destination ./mocks/extension.go -package mock_device_group_creation

// ExtensionDriver contains all the commands for the khr_device_group_creation extension
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_device_group_creation.html
type ExtensionDriver interface {
	// EnumeratePhysicalDeviceGroups enumerates groups of PhysicalDevice objects that can be used to
	// create a single logical Device
	//
	// outDataFactory - This method can be provided to allocate each PhysicalDeviceGroupProperties object
	// that is returned, along with any chained OutData structures. It can also be left nil, in which case
	// PhysicalDeviceGroupProperties will be allocated with no chained structures.
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkEnumeratePhysicalDeviceGroups.html
	EnumeratePhysicalDeviceGroups(outDataFactory func() *PhysicalDeviceGroupProperties) ([]*PhysicalDeviceGroupProperties, common.VkResult, error)
}
