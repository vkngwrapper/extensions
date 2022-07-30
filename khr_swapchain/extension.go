package khr_swapchain

//go:generate mockgen -source extension.go -destination ./mocks/extension.go -package mock_swapchain

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/core1_0"
	"github.com/vkngwrapper/core/driver"
	khr_swapchain_driver "github.com/vkngwrapper/extensions/khr_swapchain/driver"
	"math"
)

// VulkanExtension is an implementation of the Extension interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanExtension struct {
	driver  khr_swapchain_driver.Driver
	version common.APIVersion
}

// Extension contains all commands for the khr_swapchain extension (that were not added in core 1.1)
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_swapchain.html
type Extension interface {
	// Driver is the Vulkan wrapper driver used by this Extension
	Driver() khr_swapchain_driver.Driver
	// APIVersion is the maximum Vulkan API version supported by this Extension. If it is at least Vulkan 1.1,
	// khr_swapchain1_1.PromoteExtension can be used to promote this to a khr_swapchain1_1.Extension
	APIVersion() common.APIVersion

	// CreateSwapchain creates a Swapchain
	//
	// device - The Device to create the Swapchain for
	//
	// allocation - Controls host memory allocation behavior
	//
	// options - Specifies the parameters of the created Swapchain
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateSwapchainKHR.html
	CreateSwapchain(device core1_0.Device, allocation *driver.AllocationCallbacks, options SwapchainCreateInfo) (Swapchain, common.VkResult, error)
	// QueuePresent queues an Image for presentation
	// queue - A core1_0.Queue that is capable of presentation to the target khr_surface.Surface object's
	// platform on the same Device as the Image object's Swapchain
	//
	// o - Specifies parameters of the presentation
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkQueuePresentKHR.html
	QueuePresent(queue core1_0.Queue, o PresentInfo) (common.VkResult, error)
}

// CreateExtensionFromDevice produces an Extension object from a Device with
// khr_swapchain loaded
func CreateExtensionFromDevice(device core1_0.Device) *VulkanExtension {
	if !device.IsDeviceExtensionActive(ExtensionName) {
		return nil
	}

	return &VulkanExtension{
		driver:  khr_swapchain_driver.CreateDriverFromCore(device.Driver()),
		version: device.APIVersion(),
	}
}

// CreateExtensionFromDriver generates an Extension from a driver.Driver object- this is usually
// used in tests to build an Extension from mock drivers
func CreateExtensionFromDriver(driver khr_swapchain_driver.Driver) *VulkanExtension {
	return &VulkanExtension{
		driver:  driver,
		version: common.APIVersion(math.MaxUint32),
	}
}

func (e *VulkanExtension) Driver() khr_swapchain_driver.Driver {
	return e.driver
}

func (e *VulkanExtension) APIVersion() common.APIVersion {
	return e.version
}

func (e *VulkanExtension) CreateSwapchain(device core1_0.Device, allocation *driver.AllocationCallbacks, options SwapchainCreateInfo) (Swapchain, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, options)
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	var swapchain khr_swapchain_driver.VkSwapchainKHR

	res, err := e.driver.VkCreateSwapchainKHR(device.Handle(), (*khr_swapchain_driver.VkSwapchainCreateInfoKHR)(createInfo), allocation.Handle(), &swapchain)
	if err != nil {
		return nil, res, err
	}

	coreDriver := device.Driver()
	newSwapchain := coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(swapchain), driver.Core1_0, func() any {
		return &VulkanSwapchain{
			handle:            swapchain,
			device:            device.Handle(),
			driver:            e.driver,
			minimumAPIVersion: device.APIVersion(),
			coreDriver:        coreDriver,
		}
	}).(*VulkanSwapchain)
	return newSwapchain, res, nil
}

func (e *VulkanExtension) QueuePresent(queue core1_0.Queue, o PresentInfo) (common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return core1_0.VKErrorUnknown, err
	}

	createInfoPtr := (*khr_swapchain_driver.VkPresentInfoKHR)(createInfo)
	res, err := e.driver.VkQueuePresentKHR(queue.Handle(), createInfoPtr)
	popErr := o.PopulateOutData(createInfo)

	if popErr != nil {
		return core1_0.VKErrorUnknown, popErr
	} else if err != nil {
		return res, err
	}

	return res, res.ToError()
}
