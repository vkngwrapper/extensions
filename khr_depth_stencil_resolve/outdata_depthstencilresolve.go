package khr_depth_stencil_resolve

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/common"
	"unsafe"
)

// PhysicalDeviceDepthStencilResolveProperties describes depth/stencil resolve properties that can
// be supported by an implementation
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDepthStencilResolveProperties.html
type PhysicalDeviceDepthStencilResolveProperties struct {
	// SupportedDepthResolveModes indicates the set of supported depth resolve modes
	SupportedDepthResolveModes ResolveModeFlags
	// SupportedStencilResolveModes indicates the set of supported stencil resolve modes
	SupportedStencilResolveModes ResolveModeFlags
	// IndependentResolveNone is true if the implementation supports setting the depth
	// and stencil resolve modes to different values when one of those modes is ResolveModeNone
	IndependentResolveNone bool
	// IndependentResolve is true if the implementation supports all combinations of the supported
	// depth and stencil resolve modes, including setting either depth or stencil resolve mode to
	// ResolveModeNone
	IndependentResolve bool

	common.NextOutData
}

func (o *PhysicalDeviceDepthStencilResolveProperties) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceDepthStencilResolvePropertiesKHR{})))
	}

	info := (*C.VkPhysicalDeviceDepthStencilResolvePropertiesKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DEPTH_STENCIL_RESOLVE_PROPERTIES_KHR
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceDepthStencilResolveProperties) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceDepthStencilResolvePropertiesKHR)(cDataPointer)
	o.SupportedStencilResolveModes = ResolveModeFlags(info.supportedStencilResolveModes)
	o.SupportedDepthResolveModes = ResolveModeFlags(info.supportedDepthResolveModes)
	o.IndependentResolveNone = info.independentResolveNone != C.VkBool32(0)
	o.IndependentResolve = info.independentResolve != C.VkBool32(0)

	return info.pNext, nil
}
