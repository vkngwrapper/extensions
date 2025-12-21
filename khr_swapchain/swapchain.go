package khr_swapchain

//go:generate mockgen -source swapchain.go -destination ./mocks/swapchain.go -package mock_swapchain

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"time"
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/driver"
	khr_swapchain_driver "github.com/vkngwrapper/extensions/v3/khr_swapchain/driver"
)

// VulkanSwapchain is an implementation of the Swapchain interface that actually communicates
// with Vulkan. This is the default implementation. See the interface for more documentation.
type VulkanSwapchain struct {
	handle     khr_swapchain_driver.VkSwapchainKHR
	device     driver.VkDevice
	driver     khr_swapchain_driver.Driver
	coreDriver driver.Driver

	minimumAPIVersion common.APIVersion

	builder core1_0.DeviceObjectBuilder
}

// Swapchain provides the ability to present rendering results to a Surface
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainKHR.html
type Swapchain interface {
	// Handle is the internal Vulkan object handle for this Swapchain
	Handle() khr_swapchain_driver.VkSwapchainKHR

	// Destroy deletes this Swapchain and underlying structures from the device. **Warning**
	// after destruction, this object will still exist, but the Vulkan object handle
	// that backs it will be invalid. Do not call further methods on this object.
	//
	// callbacks - A set of allocation callbacks to control the memory free behavior of this command
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroySwapchainKHR.html
	Destroy(callbacks *driver.AllocationCallbacks)
	// SwapchainImages obtains a slice of the presentable Image objects associated with this Swapchain
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetSwapchainImagesKHR.html
	SwapchainImages() ([]core1_0.Image, common.VkResult, error)
	// AcquireNextImage retrieves the index of the next available presentable Image
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
	AcquireNextImage(timeout time.Duration, semaphore core1_0.Semaphore, fence core1_0.Fence) (int, common.VkResult, error)
}

func (s *VulkanSwapchain) Handle() khr_swapchain_driver.VkSwapchainKHR {
	return s.handle
}

func (s *VulkanSwapchain) Destroy(callbacks *driver.AllocationCallbacks) {
	s.driver.VkDestroySwapchainKHR(s.device, s.handle, callbacks.Handle())
}

func (s *VulkanSwapchain) attemptImages() ([]core1_0.Image, common.VkResult, error) {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	imageCountPtr := allocator.Malloc(int(unsafe.Sizeof(C.uint32_t(0))))
	imageCountRef := (*driver.Uint32)(imageCountPtr)

	res, err := s.driver.VkGetSwapchainImagesKHR(s.device, s.handle, imageCountRef, nil)
	if err != nil {
		return nil, res, err
	}

	imageCount := int(*imageCountRef)
	if imageCount == 0 {
		return nil, res, nil
	}

	imagesPtr := (*driver.VkImage)(allocator.Malloc(imageCount * int(unsafe.Sizeof([1]C.VkImage{}))))

	res, err = s.driver.VkGetSwapchainImagesKHR(s.device, s.handle, imageCountRef, imagesPtr)
	if err != nil {
		return nil, res, err
	}

	imagesSlice := ([]driver.VkImage)(unsafe.Slice(imagesPtr, imageCount))
	var result []core1_0.Image
	for i := 0; i < imageCount; i++ {
		image := s.builder.CreateImageObject(s.coreDriver, s.device, imagesSlice[i], s.minimumAPIVersion)
		result = append(result, image)
	}

	return result, res, nil
}

func (s *VulkanSwapchain) SwapchainImages() ([]core1_0.Image, common.VkResult, error) {
	var result []core1_0.Image
	var res common.VkResult
	var err error

	for doWhile := true; doWhile; doWhile = (res == core1_0.VKIncomplete) {
		result, res, err = s.attemptImages()
	}

	return result, res, err
}

func (s *VulkanSwapchain) AcquireNextImage(timeout time.Duration, semaphore core1_0.Semaphore, fence core1_0.Fence) (int, common.VkResult, error) {
	var imageIndex driver.Uint32

	var semaphoreHandle driver.VkSemaphore
	var fenceHandle driver.VkFence

	if semaphore != nil {
		semaphoreHandle = semaphore.Handle()
	}
	if fence != nil {
		fenceHandle = fence.Handle()
	}

	res, err := s.driver.VkAcquireNextImageKHR(s.device, s.handle, driver.Uint64(common.TimeoutNanoseconds(timeout)), semaphoreHandle, fenceHandle, &imageIndex)

	return int(imageIndex), res, err
}
