package khr_device_group

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/core1_0"
	"github.com/vkngwrapper/extensions/khr_swapchain"
)

// MemoryAllocateFlags specifies flags for a DeviceMemory allocation
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryAllocateFlagBits.html
type MemoryAllocateFlags int32

var memoryAllocateFlagsMapping = common.NewFlagStringMapping[MemoryAllocateFlags]()

func (f MemoryAllocateFlags) Register(str string) {
	memoryAllocateFlagsMapping.Register(f, str)
}

func (f MemoryAllocateFlags) String() string {
	return memoryAllocateFlagsMapping.FlagsToString(f)
}

////

// PeerMemoryFeatureFlags specifies supported peer memory features
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPeerMemoryFeatureFlagBits.html
type PeerMemoryFeatureFlags int32

var peerMemoryFeaturesMapping = common.NewFlagStringMapping[PeerMemoryFeatureFlags]()

func (f PeerMemoryFeatureFlags) Register(str string) {
	peerMemoryFeaturesMapping.Register(f, str)
}
func (f PeerMemoryFeatureFlags) String() string {
	return peerMemoryFeaturesMapping.FlagsToString(f)
}

///

// DeviceGroupPresentModeFlags specifies supported Device group present modes
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupPresentModeFlagBitsKHR.html
type DeviceGroupPresentModeFlags int32

var deviceGroupPresentModeFlagsMapping = common.NewFlagStringMapping[DeviceGroupPresentModeFlags]()

func (f DeviceGroupPresentModeFlags) Register(str string) {
	deviceGroupPresentModeFlagsMapping.Register(f, str)
}

func (f DeviceGroupPresentModeFlags) String() string {
	return deviceGroupPresentModeFlagsMapping.FlagsToString(f)
}

////

const (
	// ExtensionName is "VK_KHR_device_group"
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_device_group.html
	ExtensionName string = C.VK_KHR_DEVICE_GROUP_EXTENSION_NAME

	// DependencyDeviceGroup specifies that dependencies are non-device-local
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDependencyFlagBits.html
	DependencyDeviceGroup core1_0.DependencyFlags = C.VK_DEPENDENCY_DEVICE_GROUP_BIT_KHR

	// MemoryAllocateDeviceMask specifies that memory will be allocated for the devices
	// in MemoryAllocateFlagsInfo.DeviceMask
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryAllocateFlagBitsKHR.html
	MemoryAllocateDeviceMask MemoryAllocateFlags = C.VK_MEMORY_ALLOCATE_DEVICE_MASK_BIT_KHR

	// DeviceGroupPresentModeLocal specifies that any PhysicalDevice with a presentation engine can
	// present its own swapchain Image objects
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupPresentModeFlagBitsKHR.html
	DeviceGroupPresentModeLocal DeviceGroupPresentModeFlags = C.VK_DEVICE_GROUP_PRESENT_MODE_LOCAL_BIT_KHR
	// DeviceGroupPresentModeRemote specifies that any PhysicalDevice with a presentation engine can
	// present swapchain Image objects from any PhysicalDevice in its PresentMask
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupPresentModeFlagBitsKHR.html
	DeviceGroupPresentModeRemote DeviceGroupPresentModeFlags = C.VK_DEVICE_GROUP_PRESENT_MODE_REMOTE_BIT_KHR
	// DeviceGroupPresentModeSum specifies that any PhysicalDevice with a presentation engine can present
	// the sum of swapchain Image objects from any PhysicalDevice in its PresentMask
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupPresentModeFlagBitsKHR.html
	DeviceGroupPresentModeSum DeviceGroupPresentModeFlags = C.VK_DEVICE_GROUP_PRESENT_MODE_SUM_BIT_KHR
	// DeviceGroupPresentModeLocalMultiDevice specifies that multiple PhysicalDevice objects with a presentation
	// engine can each present their own swapchain Image objects
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupPresentModeFlagBitsKHR.html
	DeviceGroupPresentModeLocalMultiDevice DeviceGroupPresentModeFlags = C.VK_DEVICE_GROUP_PRESENT_MODE_LOCAL_MULTI_DEVICE_BIT_KHR

	// PeerMemoryFeatureCopyDst specifies that the memory can be accessed as the destination of
	// any CommandBuffer.CmdCopy... command
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPeerMemoryFeatureFlagBits.html
	PeerMemoryFeatureCopyDst PeerMemoryFeatureFlags = C.VK_PEER_MEMORY_FEATURE_COPY_DST_BIT_KHR
	// PeerMemoryFeatureCopySrc specifies that the memory can be accessed as the source of any
	// CommandBuffer.CmdCopy... command
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPeerMemoryFeatureFlagBits.html
	PeerMemoryFeatureCopySrc PeerMemoryFeatureFlags = C.VK_PEER_MEMORY_FEATURE_COPY_SRC_BIT_KHR
	// PeerMemoryFeatureGenericDst specifies that the memory can be written as any memory access type
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPeerMemoryFeatureFlagBits.html
	PeerMemoryFeatureGenericDst PeerMemoryFeatureFlags = C.VK_PEER_MEMORY_FEATURE_GENERIC_DST_BIT_KHR
	// PeerMemoryFeatureGenericSrc specifies that the memory can be read as any memory access type
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPeerMemoryFeatureFlagBits.html
	PeerMemoryFeatureGenericSrc PeerMemoryFeatureFlags = C.VK_PEER_MEMORY_FEATURE_GENERIC_SRC_BIT_KHR

	// PipelineCreateDispatchBase specifies that a compute pipeline can be used with
	// CommandBuffer.CmdDispatchBase with a non-zero base workgroup
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreateFlagBits.html
	PipelineCreateDispatchBase core1_0.PipelineCreateFlags = C.VK_PIPELINE_CREATE_DISPATCH_BASE_KHR
	// PipelineCreateViewIndexFromDeviceIndex specifies that any shader input variables
	// decorated as ViewIndex will be assigned values as if they were decorated as DeviceIndex
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreateFlagBits.html
	PipelineCreateViewIndexFromDeviceIndex core1_0.PipelineCreateFlags = C.VK_PIPELINE_CREATE_VIEW_INDEX_FROM_DEVICE_INDEX_BIT_KHR

	// ImageCreateSplitInstanceBindRegions specifies that the Image can be used with a non-empty
	// BindImageMemoryDeviceGroupInfo.SplitInstanceBindRegions
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateFlagBits.html
	ImageCreateSplitInstanceBindRegions core1_0.ImageCreateFlags = C.VK_IMAGE_CREATE_SPLIT_INSTANCE_BIND_REGIONS_BIT_KHR

	// SwapchainCreateSplitInstanceBindRegions specifies that Image objects created from the swpachain
	// must use ImageCreateSplitInstanceBindRegions
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainCreateFlagBitsKHR.html
	SwapchainCreateSplitInstanceBindRegions khr_swapchain.SwapchainCreateFlags = C.VK_SWAPCHAIN_CREATE_SPLIT_INSTANCE_BIND_REGIONS_BIT_KHR
)

func init() {
	DependencyDeviceGroup.Register("Device Group")

	MemoryAllocateDeviceMask.Register("Device Mask")

	DeviceGroupPresentModeLocal.Register("Local")
	DeviceGroupPresentModeRemote.Register("Remote")
	DeviceGroupPresentModeSum.Register("Sum")
	DeviceGroupPresentModeLocalMultiDevice.Register("Local Multi-Device")

	PeerMemoryFeatureCopyDst.Register("Copy Dst")
	PeerMemoryFeatureCopySrc.Register("Copy Src")
	PeerMemoryFeatureGenericDst.Register("Generic Dst")
	PeerMemoryFeatureGenericSrc.Register("Generic Src")

	PipelineCreateDispatchBase.Register("Dispatch Base")
	PipelineCreateViewIndexFromDeviceIndex.Register("View Index From Device Index")

	ImageCreateSplitInstanceBindRegions.Register("Split Instance Bind Regions")
	SwapchainCreateSplitInstanceBindRegions.Register("Split Instance Bind Regions")
}
