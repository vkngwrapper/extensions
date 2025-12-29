package khr_swapchain

import (
	"time"

	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
	khr_swapchain_driver "github.com/vkngwrapper/extensions/v3/khr_swapchain/loader"
)

// ExtensionDriver contains all commands for the khr_swapchain extension (that were not added in core 1.1)
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_swapchain.html
type ExtensionDriver interface {
	// Loader is the Vulkan wrapper loader used by this ExtensionDriver
	Loader() khr_swapchain_driver.Loader
	// APIVersion is the maximum Vulkan API version supported by this ExtensionDriver. If it is at least Vulkan 1.1,
	// khr_swapchain1_1.PromoteExtension can be used to promote this to a khr_swapchain1_1.ExtensionDriver
	APIVersion() common.APIVersion
	// Device is the vulkan Device object that backs this device extension
	Device() core1_0.Device

	// CreateSwapchain creates a Swapchain
	//
	// allocation - Controls host memory allocation behavior
	//
	// options - Specifies the parameters of the created Swapchain
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateSwapchainKHR.html
	CreateSwapchain(allocation *loader.AllocationCallbacks, options SwapchainCreateInfo) (Swapchain, common.VkResult, error)
	// QueuePresent queues an Image for presentation
	// queue - A core1_0.Queue that is capable of presentation to the target khr_surface.Surface object's
	// platform on the same Device as the Image object's Swapchain
	//
	// o - Specifies parameters of the presentation
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkQueuePresentKHR.html
	QueuePresent(queue core1_0.Queue, o PresentInfo) (common.VkResult, error)

	// DestroySwapchain deletes a Swapchain and underlying structures from the device. **Warning**
	// after destruction, the object will still exist, but the Vulkan object handle
	// that backs it will be invalid. Do not call further methods with the Swapchain
	//
	// callbacks - A set of allocation callbacks to control the memory free behavior of this command
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroySwapchainKHR.html
	DestroySwapchain(swapchain Swapchain, callbacks *loader.AllocationCallbacks)
	// GetSwapchainImages obtains a slice of the presentable Image objects associated with a Swapchain
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetSwapchainImagesKHR.html
	GetSwapchainImages(swapchain Swapchain) ([]core1_0.Image, common.VkResult, error)
	// AcquireNextImage retrieves the index of the next available presentable Image in a Swapchain
	//
	// timeout - Specifies how long the function waits, in nanoseconds, if no Image is available, before
	// returning core1_0.VKTimeout. May be common.NoTimeout to wait indefinitely. The timeout is adjusted
	// to the closest value allowed by the implementation timeout accuracy, which may be substantially
	// longer than the requested timeout.
	//
	// semaphore - Optionally, a Semaphore to signal when the Image is acquired
	//
	// fence - Optionally, a Fence to signal when the Image is acquired
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkAcquireNextImageKHR.html
	AcquireNextImage(swapchain Swapchain, timeout time.Duration, semaphore *core1_0.Semaphore, fence *core1_0.Fence) (int, common.VkResult, error)
}
