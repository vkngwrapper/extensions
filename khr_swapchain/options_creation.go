package khr_swapchain

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/pkg/errors"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	ext_surface "github.com/vkngwrapper/extensions/v3/khr_surface"
)

// SwapchainCreateInfo specifies parameters of a newly-created Swapchain object
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainCreateInfoKHR.html
type SwapchainCreateInfo struct {
	// Surface is the khr_surface.Surface onto which the Swapchain will present Image objects
	Surface ext_surface.Surface

	// Flags - Indicates parameters of the Swapchain creation
	Flags SwapchainCreateFlags

	// MinImageCount is the minimum number of presentable Image objects that the application needs.
	// The implementation will either create the Swapchain with at least that many Image objects, or it
	// will fail to create the Swapchain
	MinImageCount int

	// ImageFormat specifies the format the Swapchain Image objects will be created with
	ImageFormat core1_0.Format
	// ImageColorSpace specifies the way the Swapchain interprets Image data
	ImageColorSpace ext_surface.ColorSpace
	// ImageExtent is the size, in pixels, of the Swapchain Image objects. The behavior is platform-dependent
	// if the Image extent does not match the Surface object's CurrentExtent as returned by
	// khr_surface.Surface.PhysicalDeviceSurfaceCapabilities
	ImageExtent core1_0.Extent2D
	// ImageArrayLayers is the number of views in a multiview/stereo Surface
	ImageArrayLayers int
	// ImageUsage describes the intended usage of the (acquired) Swapchain Image objects
	ImageUsage core1_0.ImageUsageFlags

	// ImageSharingMode is the sharing mode used for the Image objects of the Swapchain
	ImageSharingMode core1_0.SharingMode
	// QueueFamilyIndices is a slice of queue family indices having access to the Image objects of
	// the Swapchain when ImageSharingMode is SharingModeConcurrent
	QueueFamilyIndices []int

	// PreTransform describes the transform, relative to the presentation engine's natural orientation,
	// applied to the Image content prior to presentation
	PreTransform ext_surface.SurfaceTransformFlags
	// CompositeAlpha indicates the alpha compositing mode to use when this Surface is composited together
	// with other surfaces on certain window systems
	CompositeAlpha ext_surface.CompositeAlphaFlags
	// PresentMode is the presentation mode the Swapchain will use
	PresentMode ext_surface.PresentMode

	// Clipped specifies whether the Vulkan implementation is allowed to discard rendering operations that
	// affect regions of the Surface that are not visible
	Clipped bool
	// OldSwapchain is, optionally, the existing non-retired Swapchain currently associated with Surface.
	// Providing a valid OldSwapchain may aid in resource reuse, and also allows the application to still
	// present any Image objects that are already acquired from it
	OldSwapchain Swapchain

	common.NextOptions
}

func (o SwapchainCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if o.Surface == nil {
		return nil, errors.New("khr_swapchain.SwapchainCreateInfo.Surface cannot be nil")
	}

	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof([1]C.VkSwapchainCreateInfoKHR{})))
	}
	createInfo := (*C.VkSwapchainCreateInfoKHR)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_SWAPCHAIN_CREATE_INFO_KHR
	createInfo.flags = C.VkSwapchainCreateFlagsKHR(o.Flags)
	createInfo.pNext = next

	createInfo.surface = C.VkSurfaceKHR(unsafe.Pointer(o.Surface.Handle()))
	createInfo.minImageCount = C.uint32_t(o.MinImageCount)

	createInfo.imageFormat = C.VkFormat(o.ImageFormat)
	createInfo.imageColorSpace = C.VkColorSpaceKHR(o.ImageColorSpace)
	createInfo.imageExtent.width = C.uint32_t(o.ImageExtent.Width)
	createInfo.imageExtent.height = C.uint32_t(o.ImageExtent.Height)
	createInfo.imageArrayLayers = C.uint32_t(o.ImageArrayLayers)
	createInfo.imageUsage = C.VkImageUsageFlags(o.ImageUsage)

	createInfo.imageSharingMode = C.VkSharingMode(o.ImageSharingMode)
	createInfo.queueFamilyIndexCount = C.uint32_t(len(o.QueueFamilyIndices))

	if len(o.QueueFamilyIndices) == 0 {
		createInfo.pQueueFamilyIndices = nil
	} else {
		familyIndexPtr := (*C.uint32_t)(allocator.Malloc(len(o.QueueFamilyIndices) * int(unsafe.Sizeof(C.uint32_t(0)))))
		createInfo.pQueueFamilyIndices = familyIndexPtr

		familyIndexSlice := ([]C.uint32_t)(unsafe.Slice(familyIndexPtr, len(o.QueueFamilyIndices)))
		for i, index := range o.QueueFamilyIndices {
			familyIndexSlice[i] = C.uint32_t(index)
		}
	}

	createInfo.preTransform = C.VkSurfaceTransformFlagBitsKHR(o.PreTransform)
	createInfo.compositeAlpha = C.VkCompositeAlphaFlagBitsKHR(o.CompositeAlpha)
	createInfo.presentMode = C.VkPresentModeKHR(o.PresentMode)

	createInfo.clipped = C.VK_FALSE
	if o.Clipped {
		createInfo.clipped = C.VK_TRUE
	}

	createInfo.oldSwapchain = nil
	if o.OldSwapchain != nil {
		createInfo.oldSwapchain = (C.VkSwapchainKHR)(unsafe.Pointer(o.OldSwapchain.Handle()))
	}

	return preallocatedPointer, nil
}
