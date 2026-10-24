package ext_4444_formats

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

type PhysicalDevice4444FormatsFeatures struct {
	FormatA4R4G4B4 bool
	FormatA4B4G4R4 bool

	common.NextOutData
	common.NextOptions
}

func (o PhysicalDevice4444FormatsFeatures) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof([1]C.VkPhysicalDevice4444FormatsFeaturesEXT{})))
	}

	features := (*C.VkPhysicalDevice4444FormatsFeaturesEXT)(preallocatedPointer)
	features.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_4444_FORMATS_FEATURES_EXT
	features.pNext = next
	features.formatA4R4G4B4 = C.VkBool32(0)
	if o.FormatA4R4G4B4 {
		features.formatA4R4G4B4 = C.VkBool32(1)
	}
	features.formatA4B4G4R4 = C.VkBool32(0)
	if o.FormatA4B4G4R4 {
		features.formatA4B4G4R4 = C.VkBool32(1)
	}

	return preallocatedPointer, nil
}

func (o *PhysicalDevice4444FormatsFeatures) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof([1]C.VkPhysicalDevice4444FormatsFeaturesEXT{})))
	}

	features := (*C.VkPhysicalDevice4444FormatsFeaturesEXT)(preallocatedPointer)
	features.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_4444_FORMATS_FEATURES_EXT
	features.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDevice4444FormatsFeatures) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	features := (*C.VkPhysicalDevice4444FormatsFeaturesEXT)(cDataPointer)
	o.FormatA4R4G4B4 = features.formatA4R4G4B4 != C.VkBool32(0)
	o.FormatA4B4G4R4 = features.formatA4B4G4R4 != C.VkBool32(0)

	return features.pNext, nil
}
