package khr_maintenance3_shim

import (
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/core1_1"
	"github.com/vkngwrapper/extensions/v3/khr_maintenance3"
)

//go:generate mockgen -source shim.go -destination ../mocks/shim.go -package mock_maintenance3

// Shim contains all commands for the khr_maintenance3 extension
type Shim interface {
	// DescriptorSetLayoutSupport queries whether a DescriptorSetLayout can be created
	//
	// o - Specifies the state of the DescriptorSetLayout object
	//
	// outData - A pre-allocated object in which information about support for the DescriptorSetLayout
	// object will be populated. It should include any desired chained OutData objects
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDescriptorSetLayoutSupport.html
	DescriptorSetLayoutSupport(o core1_0.DescriptorSetLayoutCreateInfo, outData *core1_1.DescriptorSetLayoutSupport) error
}

type VulkanShim struct {
	extension khr_maintenance3.Extension
	device    core1_0.Device
}

// Compiler check that VulkanShim satisfies Shim
var _ Shim = &VulkanShim{}

func NewShim(extension khr_maintenance3.Extension, device core1_0.Device) *VulkanShim {
	if extension == nil {
		panic("extension cannot be nil")
	}
	if device == nil {
		panic("device cannot be nil")
	}
	return &VulkanShim{
		extension: extension,
		device:    device,
	}
}

func (s *VulkanShim) DescriptorSetLayoutSupport(o core1_0.DescriptorSetLayoutCreateInfo, outData *core1_1.DescriptorSetLayoutSupport) error {
	return s.extension.DescriptorSetLayoutSupport(s.device, o, (*khr_maintenance3.DescriptorSetLayoutSupport)(outData))
}
