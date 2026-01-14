package khr_ray_query

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

type PhysicalDeviceRayQueryFeatures struct {
	RayQuery bool

	common.NextOutData
	common.NextOptions
}

func (o PhysicalDeviceRayQueryFeatures) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceRayQueryFeaturesKHR{})))
	}

	info := (*C.VkPhysicalDeviceRayQueryFeaturesKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_RAY_QUERY_FEATURES_KHR
	info.pNext = next
	info.rayQuery = C.VkBool32(0)

	if o.RayQuery {
		info.rayQuery = C.VkBool32(1)
	}

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceRayQueryFeatures) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceRayQueryFeaturesKHR{})))
	}

	info := (*C.VkPhysicalDeviceRayQueryFeaturesKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_RAY_QUERY_FEATURES_KHR
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceRayQueryFeatures) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	outData := (*C.VkPhysicalDeviceRayQueryFeaturesKHR)(cDataPointer)

	o.RayQuery = outData.rayQuery != C.VkBool32(0)

	return outData.pNext, nil
}
