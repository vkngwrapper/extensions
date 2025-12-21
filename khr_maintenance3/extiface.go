package khr_maintenance3

import "github.com/vkngwrapper/core/v3/core1_0"

//go:generate mockgen -source extiface.go -destination ./mocks/extension.go -package mock_maintenance3

// Extension contains all commands for the khr_maintenance3 extension
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_maintenance3.html
type Extension interface {
	// DescriptorSetLayoutSupport queries whether a DescriptorSetLayout can be created
	//
	// device - The Device which will be used to create the DescriptorSetLayout
	//
	// setLayoutOptions - Specifies the state of the DescriptorSetLayout object
	//
	// outData - A pre-allocated object in which information about support for the DescriptorSetLayout
	// object will be populated. It should include any desired chained OutData objects
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDescriptorSetLayoutSupport.html
	DescriptorSetLayoutSupport(device core1_0.Device, setLayoutOptions core1_0.DescriptorSetLayoutCreateInfo, outData *DescriptorSetLayoutSupport) error
}
