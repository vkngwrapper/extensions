package ext_pipeline_creation_cache_control

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

type PhysicalDevicePipelineCreationCacheControlFeatures struct {
	PipelineCreationCacheControl bool

	common.NextOutData
	common.NextOptions
}

func (o PhysicalDevicePipelineCreationCacheControlFeatures) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof([1]C.VkPhysicalDevicePipelineCreationCacheControlFeaturesEXT{})))
	}

	features := (*C.VkPhysicalDevicePipelineCreationCacheControlFeaturesEXT)(preallocatedPointer)
	features.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PIPELINE_CREATION_CACHE_CONTROL_FEATURES
	features.pNext = next
	features.pipelineCreationCacheControl = C.VkBool32(0)
	if o.PipelineCreationCacheControl {
		features.pipelineCreationCacheControl = C.VkBool32(1)
	}

	return preallocatedPointer, nil
}

func (o *PhysicalDevicePipelineCreationCacheControlFeatures) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof([1]C.VkPhysicalDevicePipelineCreationCacheControlFeaturesEXT{})))
	}

	features := (*C.VkPhysicalDevicePipelineCreationCacheControlFeaturesEXT)(preallocatedPointer)
	features.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PIPELINE_CREATION_CACHE_CONTROL_FEATURES
	features.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDevicePipelineCreationCacheControlFeatures) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	features := (*C.VkPhysicalDevicePipelineCreationCacheControlFeaturesEXT)(cDataPointer)
	o.PipelineCreationCacheControl = features.pipelineCreationCacheControl != C.VkBool32(0)

	return features.pNext, nil
}
