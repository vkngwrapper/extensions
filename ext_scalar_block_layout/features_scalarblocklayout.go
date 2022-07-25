package ext_scalar_block_layout

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

// PhysicalDeviceScalarBlockLayoutFeatures indicates support for scalar block layouts
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceScalarBlockLayoutFeatures.html
type PhysicalDeviceScalarBlockLayoutFeatures struct {
	// ScalarBlockLayout indicates that the implementation supports the layout of resource blocks
	// in shaders using scalar alignment
	ScalarBlockLayout bool

	common.NextOptions
	common.NextOutData
}

func (o *PhysicalDeviceScalarBlockLayoutFeatures) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceScalarBlockLayoutFeaturesEXT{})))
	}

	info := (*C.VkPhysicalDeviceScalarBlockLayoutFeaturesEXT)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SCALAR_BLOCK_LAYOUT_FEATURES_EXT
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceScalarBlockLayoutFeatures) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceScalarBlockLayoutFeaturesEXT)(cDataPointer)

	o.ScalarBlockLayout = info.scalarBlockLayout != C.VkBool32(0)

	return info.pNext, nil
}

func (o PhysicalDeviceScalarBlockLayoutFeatures) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceScalarBlockLayoutFeaturesEXT{})))
	}

	info := (*C.VkPhysicalDeviceScalarBlockLayoutFeaturesEXT)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SCALAR_BLOCK_LAYOUT_FEATURES_EXT
	info.pNext = next
	info.scalarBlockLayout = C.VkBool32(0)

	if o.ScalarBlockLayout {
		info.scalarBlockLayout = C.VkBool32(1)
	}

	return preallocatedPointer, nil
}
