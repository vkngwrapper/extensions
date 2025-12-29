package khr_device_group

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"time"
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/pkg/errors"
	"github.com/vkngwrapper/core/v3"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/extensions/v3/khr_swapchain"
)

// AcquireNextImageInfo specifies parameters of the acquire operation
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAcquireNextImageInfoKHR.html
type AcquireNextImageInfo struct {
	// Swapchain is the khr_swapchain.Swapchain from which an Image is acquired
	Swapchain khr_swapchain.Swapchain
	// Timeout is how long to wait, in nanoseconds, if no Image is available, before returning core1_0.VKTimeout.
	// May be common.NoTimeout to wait indefinitely. The timeout is adjusted to the closest value allowed by the
	// implementation timeout accuracy, which may be substantially longer than the requested timeout.
	Timeout time.Duration
	// Semaphore is, optionally, a Semaphore to signal when acquisition is complete
	Semaphore core.Semaphore
	// Fence is, optionally, a Fence to signal when acquisition is complete
	Fence core.Fence
	// DeviceMask is a mask of PhysicalDevice objects for which the khr_swapchain.Swapchain Image will be
	// ready to use when the Semaphore or Fence is signaled
	DeviceMask uint32

	common.NextOptions
}

func (o AcquireNextImageInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if !o.Swapchain.Initialized() {
		return nil, errors.New("field Swapchain of AcquireNextImageInfo must contain a valid swapchain")
	}

	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkAcquireNextImageInfoKHR{})))
	}

	info := (*C.VkAcquireNextImageInfoKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_ACQUIRE_NEXT_IMAGE_INFO_KHR
	info.pNext = next
	info.swapchain = C.VkSwapchainKHR(unsafe.Pointer(o.Swapchain.Handle()))
	info.semaphore = nil
	info.fence = nil
	info.timeout = C.uint64_t(common.TimeoutNanoseconds(o.Timeout))
	info.deviceMask = C.uint32_t(o.DeviceMask)

	if o.Semaphore.Initialized() {
		info.semaphore = C.VkSemaphore(unsafe.Pointer(o.Semaphore.Handle()))
	}
	if o.Fence.Initialized() {
		info.fence = C.VkFence(unsafe.Pointer(o.Fence.Handle()))
	}

	return preallocatedPointer, nil
}
