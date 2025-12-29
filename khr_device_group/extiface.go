package khr_device_group

import (
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/extensions/v3/khr_surface"
)

//go:generate mockgen -source extiface.go -destination ./mocks/extension.go -package mock_device_group

// ExtensionDriver contains all the commands for the khr_device_group extension
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_device_group.html
type ExtensionDriver interface {
	// CmdDispatchBase dispatches compute work items with non-zero base values for the workgroup IDs
	//
	// commandBuffer - The CommandBuffer to dispatch the work items on
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
	CmdDispatchBase(commandBuffer core1_0.CommandBuffer, baseGroupX, baseGroupY, baseGroupZ, groupCountX, groupCountY, groupCountZ int)
	// CmdSetDeviceMask modifies the device mask of a CommandBuffer
	//
	// commandBuffer - The CommandBuffer to set the Device mask on
	//
	// deviceMask - The new value of the current Device mask
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDeviceMask.html
	CmdSetDeviceMask(commandBuffer core1_0.CommandBuffer, deviceMask uint32)
	// GetDeviceGroupPeerMemoryFeatures queries supported peer memory features of a Device
	//
	// device - The Device to query peer memory features on
	//
	// heapIndex - The index of the memory heap from which the memory is allocated
	//
	// localDeviceIndex - The device index of the PhysicalDevice that performs the memory access
	//
	// remoteDeviceIndex - The device index of the PhysicalDevice that the memory is allocated for
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDeviceGroupPeerMemoryFeatures.html
	GetDeviceGroupPeerMemoryFeatures(heapIndex, localDeviceIndex, remoteDeviceIndex int) PeerMemoryFeatureFlags

	// WithKHRSurface will return nil if the khr_surface extension is not active, or will return an
	// object with additional commands that are available when both khr_device_group and khr_surface are
	// both active
	WithKHRSurface() ExtensionDriverWithKHRSurface
	// WithKHRSwapchain will return nil if the khr_swapchain extension is not active, or will return an
	// object with additional commands that are available when both khr_device_group and khr_swapchain are
	// both active
	WithKHRSwapchain() ExtensionDriverWithKHRSwapchain
}

// ExtensionDriverWithKHRSurface contains commands available when both khr_device_group and khr_surface extensions
// are active
type ExtensionDriverWithKHRSurface interface {
	// GetDeviceGroupPresentCapabilities queries Device group present capabilities for a surface
	//
	// device - The Device being queried
	//
	// outData - A pre-allocated object in which the capabilities will be populated. It should include any desired
	// chained OutData objects
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDeviceGroupPresentCapabilitiesKHR.html
	GetDeviceGroupPresentCapabilities(outData *DeviceGroupPresentCapabilities) (common.VkResult, error)
	// GetDeviceGroupSurfacePresentModes queries present capabilities for a khr_surface.Surface
	//
	// device - The Device being queried
	//
	// surface - The Surface whose present capabilities are being requested
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDeviceGroupSurfacePresentModesKHR.html
	GetDeviceGroupSurfacePresentModes(surface khr_surface.Surface) (DeviceGroupPresentModeFlags, common.VkResult, error)
	// GetPhysicalDevicePresentRectangles queries present rectangles for a khr_surface.Surface on a PhysicalDevice
	//
	// physicalDevice - The PhysicalDevice being queried
	//
	// surface - The Surface whose present rectangles are being requested
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDevicePresentRectanglesKHR.html
	GetPhysicalDevicePresentRectangles(physicalDevice core1_0.PhysicalDevice, surface khr_surface.Surface) ([]core1_0.Rect2D, common.VkResult, error)
}

// ExtensionDriverWithKHRSwapchain contains commands available when both khr_device_group and khr_swapchain extensions
// are active
type ExtensionDriverWithKHRSwapchain interface {
	// AcquireNextImage2 retrieves the index of the next available presentable Image
	//
	// device - The Device which owns the requested khr_swapchain.Swapchain
	//
	// o - Contains parameters of the acquire operation
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkAcquireNextImage2KHR.html
	AcquireNextImage2(o AcquireNextImageInfo) (int, common.VkResult, error)
}
