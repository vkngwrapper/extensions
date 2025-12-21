package khr_swapchain1_1

/*
#include <stdlib.h>
#include "../../vulkan/vulkan.h"
*/
import "C"
import "github.com/vkngwrapper/core/v3/common"

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
	// DeviceGroupPresentModeLocal specifies that any PhysicalDevice with a presentation engine can
	// present its own Swapchain Image objects
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupPresentModeFlagBitsKHR.html
	DeviceGroupPresentModeLocal DeviceGroupPresentModeFlags = C.VK_DEVICE_GROUP_PRESENT_MODE_LOCAL_BIT_KHR
	// DeviceGroupPresentModeRemote specifies that any PhysicalDevice with a presentation engine can
	// present Swapchain Image objects from any PhysicalDevice in its PresentMask
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupPresentModeFlagBitsKHR.html
	DeviceGroupPresentModeRemote DeviceGroupPresentModeFlags = C.VK_DEVICE_GROUP_PRESENT_MODE_REMOTE_BIT_KHR
	// DeviceGroupPresentModeSum specifies that any PhysicalDevice with a presentation engine can present
	// the sum of Swapchain Image objects from any PhysicalDevice in its PresentMask
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupPresentModeFlagBitsKHR.html
	DeviceGroupPresentModeSum DeviceGroupPresentModeFlags = C.VK_DEVICE_GROUP_PRESENT_MODE_SUM_BIT_KHR
	// DeviceGroupPresentModeLocalMultiDevice specifies that multiple PhysicalDevice objects with a presentation
	// engine can each present their own Swapchain Image objects
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupPresentModeFlagBitsKHR.html
	DeviceGroupPresentModeLocalMultiDevice DeviceGroupPresentModeFlags = C.VK_DEVICE_GROUP_PRESENT_MODE_LOCAL_MULTI_DEVICE_BIT_KHR
)

func init() {
	DeviceGroupPresentModeLocal.Register("Local")
	DeviceGroupPresentModeRemote.Register("Remote")
	DeviceGroupPresentModeSum.Register("Sum")
	DeviceGroupPresentModeLocalMultiDevice.Register("Local Multi-Device")
}
