package khr_device_group_creation_shim

import (
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/core1_0"
	"github.com/vkngwrapper/core/v2/core1_1"
	"github.com/vkngwrapper/extensions/v2/khr_device_group_creation"
)

//go:generate mockgen -source shim.go -destination ../mocks/shim.go -package mock_device_group_creation

// Shim contains all the commands for the khr_device_group_creation extension
type Shim interface {
	// EnumeratePhysicalDeviceGroups enumerates groups of PhysicalDevice objects that can be used to
	// create a single logical Device
	//
	// outDataFactory - This method can be provided to allocate each PhysicalDeviceGroupProperties object
	// that is returned, along with any chained OutData structures. It can also be left nil, in which case
	// PhysicalDeviceGroupProperties will be allocated with no chained structures.
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkEnumeratePhysicalDeviceGroups.html
	EnumeratePhysicalDeviceGroups(outDataFactory func() *core1_1.PhysicalDeviceGroupProperties) ([]*core1_1.PhysicalDeviceGroupProperties, common.VkResult, error)
}

type VulkanShim struct {
	extension khr_device_group_creation.Extension
	instance  core1_0.Instance
}

// Compiler check that VulkanShim satisfies Shim
var _ Shim = &VulkanShim{}

func NewShim(extension khr_device_group_creation.Extension, instance core1_0.Instance) *VulkanShim {
	if extension == nil {
		panic("extension cannot be nil")
	}
	if instance == nil {
		panic("instance cannot be nil")
	}
	return &VulkanShim{
		extension: extension,
		instance:  instance,
	}
}

func (s *VulkanShim) EnumeratePhysicalDeviceGroups(outDataFactory func() *core1_1.PhysicalDeviceGroupProperties) ([]*core1_1.PhysicalDeviceGroupProperties, common.VkResult, error) {
	factory := func() *khr_device_group_creation.PhysicalDeviceGroupProperties {
		return (*khr_device_group_creation.PhysicalDeviceGroupProperties)(outDataFactory())
	}

	retVal, result, err := s.extension.EnumeratePhysicalDeviceGroups(s.instance, factory)
	castRetVal := make([]*core1_1.PhysicalDeviceGroupProperties, 0, len(retVal))
	for _, prop := range retVal {
		castRetVal = append(castRetVal, (*core1_1.PhysicalDeviceGroupProperties)(prop))
	}

	return castRetVal, result, err
}
