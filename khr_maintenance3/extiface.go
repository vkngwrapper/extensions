package khr_maintenance3

import (
	"github.com/vkngwrapper/core/v3/core1_0"
)

//go:generate mockgen -source extiface.go -destination ./mocks/extension.go -package mock_maintenance3

// ExtensionDriver contains all commands for the khr_maintenance3 extension
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_maintenance3.html
type ExtensionDriver interface {
	// DescriptorSetLayoutSupport queries whether a DescriptorSetLayout can be created
	//
	// setLayoutOptions - Specifies the state of the DescriptorSetLayout object
	//
	// outData - A pre-allocated object in which information about support for the DescriptorSetLayout
	// object will be populated. It should include any desired chained OutData objects
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDescriptorSetLayoutSupport.html
	DescriptorSetLayoutSupport(setLayoutOptions core1_0.DescriptorSetLayoutCreateInfo, outData *DescriptorSetLayoutSupport) error
}
