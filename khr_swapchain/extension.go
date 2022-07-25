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

type Extension interface {
	Driver() khr_swapchain_driver.Driver
	APIVersion() common.APIVersion

	CreateSwapchain(device core1_0.Device, allocation *driver.AllocationCallbacks, options SwapchainCreateInfo) (Swapchain, common.VkResult, error)
	QueuePresent(queue core1_0.Queue, o PresentInfo) (common.VkResult, error)
}

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
		return &vulkanSwapchain{
			handle:            swapchain,
			device:            device.Handle(),
			driver:            e.driver,
			minimumAPIVersion: device.APIVersion(),
			coreDriver:        coreDriver,
		}
	}).(*vulkanSwapchain)
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
