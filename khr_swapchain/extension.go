package khr_swapchain

//go:generate mockgen -source extension.go -destination ./mocks/extension.go -package mock_swapchain

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"math"
	"time"
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
	khr_swapchain_driver "github.com/vkngwrapper/extensions/v3/khr_swapchain/loader"
)

// VulkanExtension is an implementation of the Extension interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanExtension struct {
	loader  khr_swapchain_driver.Loader
	version common.APIVersion
}

// CreateExtensionFromDevice produces an Extension object from a Device with
// khr_swapchain loaded
func CreateExtensionFromDevice(device core.Device) *VulkanExtension {
	if !device.IsDeviceExtensionActive(ExtensionName) {
		return nil
	}

	return &VulkanExtension{
		loader:  khr_swapchain_driver.CreateDriverFromCore(device.Driver()),
		version: device.APIVersion(),
	}
}

// CreateExtensionFromDriver generates an Extension from a loader.Loader object- this is usually
// used in tests to build an Extension from mock drivers
func CreateExtensionFromDriver(loader khr_swapchain_driver.Loader) *VulkanExtension {
	return &VulkanExtension{
		loader:  loader,
		version: common.APIVersion(math.MaxUint32),
	}
}

func (e *VulkanExtension) Loader() khr_swapchain_driver.Loader {
	return e.loader
}

func (e *VulkanExtension) APIVersion() common.APIVersion {
	return e.version
}

func (e *VulkanExtension) CreateSwapchain(device core.Device, allocation *loader.AllocationCallbacks, options SwapchainCreateInfo) (Swapchain, common.VkResult, error) {
	if device.Handle() == 0 {
		panic("device cannot be uninitialized")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, options)
	if err != nil {
		return Swapchain{}, core1_0.VKErrorUnknown, err
	}

	var swapchain khr_swapchain_driver.VkSwapchainKHR

	res, err := e.loader.VkCreateSwapchainKHR(device.Handle(), (*khr_swapchain_driver.VkSwapchainCreateInfoKHR)(createInfo), allocation.Handle(), &swapchain)
	if err != nil {
		return Swapchain{}, res, err
	}

	newSwapchain := Swapchain{
		handle:     swapchain,
		device:     device.Handle(),
		apiVersion: device.APIVersion(),
	}
	return newSwapchain, res, nil
}

func (e *VulkanExtension) QueuePresent(queue core.Queue, o PresentInfo) (common.VkResult, error) {
	if queue.Handle() == 0 {
		panic("queue cannot be uninitialized")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return core1_0.VKErrorUnknown, err
	}

	createInfoPtr := (*khr_swapchain_driver.VkPresentInfoKHR)(createInfo)
	res, err := e.loader.VkQueuePresentKHR(queue.Handle(), createInfoPtr)
	popErr := o.PopulateOutData(createInfo)

	if popErr != nil {
		return core1_0.VKErrorUnknown, popErr
	} else if err != nil {
		return res, err
	}

	return res, res.ToError()
}

func (s *VulkanExtension) attemptImages(swapchain Swapchain) ([]core.Image, common.VkResult, error) {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	imageCountPtr := allocator.Malloc(int(unsafe.Sizeof(C.uint32_t(0))))
	imageCountRef := (*loader.Uint32)(imageCountPtr)

	res, err := s.loader.VkGetSwapchainImagesKHR(swapchain.DeviceHandle(), swapchain.Handle(), imageCountRef, nil)
	if err != nil {
		return nil, res, err
	}

	imageCount := int(*imageCountRef)
	if imageCount == 0 {
		return nil, res, nil
	}

	imagesPtr := (*loader.VkImage)(allocator.Malloc(imageCount * int(unsafe.Sizeof([1]C.VkImage{}))))

	res, err = s.loader.VkGetSwapchainImagesKHR(swapchain.DeviceHandle(), swapchain.Handle(), imageCountRef, imagesPtr)
	if err != nil {
		return nil, res, err
	}

	imagesSlice := ([]loader.VkImage)(unsafe.Slice(imagesPtr, imageCount))
	var result []core.Image
	for i := 0; i < imageCount; i++ {
		image := core.InternalImage(swapchain.DeviceHandle(), imagesSlice[i], swapchain.APIVersion())
		result = append(result, image)
	}

	return result, res, nil
}

func (s *VulkanExtension) SwapchainImages(swapchain Swapchain) ([]core.Image, common.VkResult, error) {
	var result []core.Image
	var res common.VkResult
	var err error

	for doWhile := true; doWhile; doWhile = (res == core1_0.VKIncomplete) {
		result, res, err = s.attemptImages(swapchain)
	}

	return result, res, err
}

func (s *VulkanExtension) AcquireNextImage(swapchain Swapchain, timeout time.Duration, semaphore core.Semaphore, fence core.Fence) (int, common.VkResult, error) {
	var imageIndex loader.Uint32

	var semaphoreHandle loader.VkSemaphore
	var fenceHandle loader.VkFence

	if semaphore.Handle() != 0 {
		semaphoreHandle = semaphore.Handle()
	}
	if fence.Handle() != 0 {
		fenceHandle = fence.Handle()
	}

	res, err := s.loader.VkAcquireNextImageKHR(swapchain.DeviceHandle(), swapchain.Handle(), loader.Uint64(common.TimeoutNanoseconds(timeout)), semaphoreHandle, fenceHandle, &imageIndex)

	return int(imageIndex), res, err
}

func (v *VulkanExtension) DestroySwapchain(swapchain Swapchain, callbacks *loader.AllocationCallbacks) {
	if swapchain.Handle() == 0 {
		panic("swapchain was uninitialized")
	}
	v.loader.VkDestroySwapchainKHR(swapchain.DeviceHandle(), swapchain.Handle(), callbacks.Handle())
}
