package khr_device_group_shim

import (
	"github.com/vkngwrapper/core/v2/core1_0"
	"github.com/vkngwrapper/core/v2/core1_1"
	"github.com/vkngwrapper/extensions/v3/khr_device_group"
)

//go:generate mockgen -source shim.go -destination ../mocks/shim.go -package mock_device_group

// CommandBufferShim contains all the commands for the khr_device_group extension that uses CommandBuffer
type CommandBufferShim interface {
	// CmdDispatchBase dispatches compute work items with non-zero base values for the workgroup IDs
	//
	// baseGroupX - The start value of the X component of WorkgroupId
	//
	// baseGroupY - The start value of the Y component of WorkGroupId
	//
	// baseGroupZ - The start value of the Z component of WorkGroupId
	//
	// groupCountX - The number of local workgroups to dispatch in the X dimension
	//
	// groupCountY - The number of local workgroups to dispatch in the Y dimension
	//
	// groupCountZ - The number of local workgroups to dispatch in the Z dimension
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDispatchBase.html
	CmdDispatchBase(baseGroupX, baseGroupY, baseGroupZ, groupCountX, groupCountY, groupCountZ int)
	// CmdSetDeviceMask modifies the device mask of a CommandBuffer
	//
	// deviceMask - The new value of the current Device mask
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDeviceMask.html
	CmdSetDeviceMask(deviceMask uint32)
}

type VulkanCommandBufferShim struct {
	extension     khr_device_group.Extension
	commandBuffer core1_0.CommandBuffer
}

func NewCommandBufferShim(extension khr_device_group.Extension, commandBuffer core1_0.CommandBuffer) *VulkanCommandBufferShim {
	if extension == nil {
		panic("extension cannot be nil")
	}
	if commandBuffer == nil {
		panic("commandBuffer cannot be nil")
	}
	return &VulkanCommandBufferShim{
		extension:     extension,
		commandBuffer: commandBuffer,
	}
}

func (s *VulkanCommandBufferShim) CmdDispatchBase(baseGroupX, baseGroupY, baseGroupZ, groupCountX, groupCountY, groupCountZ int) {
	s.extension.CmdDispatchBase(s.commandBuffer, baseGroupX, baseGroupY, baseGroupZ, groupCountX, groupCountY, groupCountZ)
}

func (s *VulkanCommandBufferShim) CmdSetDeviceMask(deviceMask uint32) {
	s.extension.CmdSetDeviceMask(s.commandBuffer, deviceMask)
}

// DeviceShim contains all the commands for the khr_device_group extension that uses Device
type DeviceShim interface {
	// DeviceGroupPeerMemoryFeatures queries supported peer memory features of a Device
	//
	// heapIndex - The index of the memory heap from which the memory is allocated
	//
	// localDeviceIndex - The device index of the PhysicalDevice that performs the memory access
	//
	// remoteDeviceIndex - The device index of the PhysicalDevice that the memory is allocated for
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDeviceGroupPeerMemoryFeatures.html
	DeviceGroupPeerMemoryFeatures(heapIndex, localDeviceIndex, remoteDeviceIndex int) core1_1.PeerMemoryFeatureFlags
}

type VulkanDeviceShim struct {
	extension khr_device_group.Extension
	device    core1_0.Device
}

func NewDeviceShim(extension khr_device_group.Extension, device core1_0.Device) *VulkanDeviceShim {
	if extension == nil {
		panic("extension cannot be nil")
	}
	if device == nil {
		panic("device cannot be nil")
	}
	return &VulkanDeviceShim{
		extension: extension,
		device:    device,
	}
}

func (s *VulkanDeviceShim) DeviceGroupPeerMemoryFeatures(heapIndex, localDeviceIndex, remoteDeviceIndex int) core1_1.PeerMemoryFeatureFlags {
	flags := s.extension.DeviceGroupPeerMemoryFeatures(s.device, heapIndex, localDeviceIndex, remoteDeviceIndex)
	return core1_1.PeerMemoryFeatureFlags(flags)
}
