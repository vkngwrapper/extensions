package khr_imageless_framebuffer

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3/common"
)

// PhysicalDeviceImagelessFramebufferFeatures indicates supports for imageless Framebuffer objects
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceImagelessFramebufferFeatures.html
type PhysicalDeviceImagelessFramebufferFeatures struct {
	// ImagelessFramebuffer indicates that the implementation supports specifying the ImageView for
	// attachments at RenderPass begin time via RenderPassAttachmentBeginInfo
	ImagelessFramebuffer bool

	common.NextOptions
	common.NextOutData
}

func (o *PhysicalDeviceImagelessFramebufferFeatures) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceImagelessFramebufferFeaturesKHR{})))
	}

	info := (*C.VkPhysicalDeviceImagelessFramebufferFeaturesKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGELESS_FRAMEBUFFER_FEATURES_KHR
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceImagelessFramebufferFeatures) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceImagelessFramebufferFeaturesKHR)(cDataPointer)

	o.ImagelessFramebuffer = info.imagelessFramebuffer != C.VkBool32(0)

	return info.pNext, nil
}

func (o PhysicalDeviceImagelessFramebufferFeatures) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceImagelessFramebufferFeaturesKHR{})))
	}

	info := (*C.VkPhysicalDeviceImagelessFramebufferFeaturesKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGELESS_FRAMEBUFFER_FEATURES_KHR
	info.pNext = next
	info.imagelessFramebuffer = C.VkBool32(0)

	if o.ImagelessFramebuffer {
		info.imagelessFramebuffer = C.VkBool32(1)
	}

	return preallocatedPointer, nil
}
