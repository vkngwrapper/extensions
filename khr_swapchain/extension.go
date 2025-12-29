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
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
	"github.com/vkngwrapper/extensions/v3/khr_swapchain/loader"
)

// VulkanExtensionDriver is an implementation of the ExtensionDriver interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanExtensionDriver struct {
	loader  khr_swapchain_loader.Loader
	version common.APIVersion
	device  core1_0.Device
}

// CreateExtensionDriverFromCoreDriver produces an ExtensionDriver object from a Device with
// khr_swapchain loaded
func CreateExtensionDriverFromCoreDriver(driver core1_0.DeviceDriver) ExtensionDriver {
	device := driver.Device()

	if !device.IsDeviceExtensionActive(ExtensionName) {
		return nil
	}

	return &VulkanExtensionDriver{
		loader:  khr_swapchain_loader.CreateLoaderFromCore(driver.Loader()),
		version: device.APIVersion(),
		device:  device,
	}
}

// CreateExtensionDriverFromLoader generates an ExtensionDriver from a loader.Loader object- this is usually
// used in tests to build an ExtensionDriver from mock drivers
func CreateExtensionDriverFromLoader(loader khr_swapchain_loader.Loader, device core1_0.Device) *VulkanExtensionDriver {
	return &VulkanExtensionDriver{
		loader:  loader,
		version: common.APIVersion(math.MaxUint32),
		device:  device,
	}
}

func (s *VulkanExtensionDriver) Loader() khr_swapchain_loader.Loader {
	return s.loader
}

func (s *VulkanExtensionDriver) APIVersion() common.APIVersion {
	return s.version
}

func (s *VulkanExtensionDriver) Device() core1_0.Device { return s.device }

func (s *VulkanExtensionDriver) CreateSwapchain(allocation *loader.AllocationCallbacks, options SwapchainCreateInfo) (Swapchain, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, options)
	if err != nil {
		return Swapchain{}, core1_0.VKErrorUnknown, err
	}

	var swapchain khr_swapchain_loader.VkSwapchainKHR

	res, err := s.loader.VkCreateSwapchainKHR(s.device.Handle(), (*khr_swapchain_loader.VkSwapchainCreateInfoKHR)(createInfo), allocation.Handle(), &swapchain)
	if err != nil {
		return Swapchain{}, res, err
	}

	newSwapchain := Swapchain{
		handle:     swapchain,
		device:     s.device.Handle(),
		apiVersion: s.device.APIVersion(),
	}
	return newSwapchain, res, nil
}

func (s *VulkanExtensionDriver) QueuePresent(queue core1_0.Queue, o PresentInfo) (common.VkResult, error) {
	if !queue.Initialized() {
		panic("queue cannot be uninitialized")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return core1_0.VKErrorUnknown, err
	}

	createInfoPtr := (*khr_swapchain_loader.VkPresentInfoKHR)(createInfo)
	res, err := s.loader.VkQueuePresentKHR(queue.Handle(), createInfoPtr)
	popErr := o.PopulateOutData(createInfo)

	if popErr != nil {
		return core1_0.VKErrorUnknown, popErr
	} else if err != nil {
		return res, err
	}

	return res, res.ToError()
}

func (s *VulkanExtensionDriver) attemptImages(swapchain Swapchain) ([]core1_0.Image, common.VkResult, error) {
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
	var result []core1_0.Image
	for i := 0; i < imageCount; i++ {
		image := core1_0.InternalImage(swapchain.DeviceHandle(), imagesSlice[i], swapchain.APIVersion())
		result = append(result, image)
	}

	return result, res, nil
}

func (s *VulkanExtensionDriver) GetSwapchainImages(swapchain Swapchain) ([]core1_0.Image, common.VkResult, error) {
	var result []core1_0.Image
	var res common.VkResult
	var err error

	for doWhile := true; doWhile; doWhile = (res == core1_0.VKIncomplete) {
		result, res, err = s.attemptImages(swapchain)
	}

	return result, res, err
}

func (s *VulkanExtensionDriver) AcquireNextImage(swapchain Swapchain, timeout time.Duration, semaphore *core1_0.Semaphore, fence *core1_0.Fence) (int, common.VkResult, error) {
	var imageIndex loader.Uint32

	var semaphoreHandle loader.VkSemaphore
	var fenceHandle loader.VkFence

	if semaphore != nil {
		semaphoreHandle = semaphore.Handle()
	}
	if fence != nil {
		fenceHandle = fence.Handle()
	}

	res, err := s.loader.VkAcquireNextImageKHR(swapchain.DeviceHandle(), swapchain.Handle(), loader.Uint64(common.TimeoutNanoseconds(timeout)), semaphoreHandle, fenceHandle, &imageIndex)

	return int(imageIndex), res, err
}

func (s *VulkanExtensionDriver) DestroySwapchain(swapchain Swapchain, callbacks *loader.AllocationCallbacks) {
	if !swapchain.Initialized() {
		panic("swapchain was uninitialized")
	}
	s.loader.VkDestroySwapchainKHR(swapchain.DeviceHandle(), swapchain.Handle(), callbacks.Handle())
}
