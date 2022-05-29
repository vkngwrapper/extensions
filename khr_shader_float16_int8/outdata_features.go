package khr_shader_float16_int8

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type PhysicalDeviceShaderFloat16Int8FeaturesOutData struct {
	ShaderFloat16 bool
	ShaderInt8    bool

	common.HaveNext
}

func (o *PhysicalDeviceShaderFloat16Int8FeaturesOutData) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceShaderFloat16Int8FeaturesKHR{})))
	}

	info := (*C.VkPhysicalDeviceShaderFloat16Int8FeaturesKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_FLOAT16_INT8_FEATURES_KHR
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceShaderFloat16Int8FeaturesOutData) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceShaderFloat16Int8FeaturesKHR)(cDataPointer)

	o.ShaderFloat16 = info.shaderFloat16 != C.VkBool32(0)
	o.ShaderInt8 = info.shaderInt8 != C.VkBool32(0)

	return info.pNext, nil
}
