package khr_shader_float16_int8

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"

import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v2/common"
	"unsafe"
)

// PhysicalDeviceShaderFloat16Int8Features describes features supported by khr_shader_float16_int8
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderFloat16Int8Features.html
type PhysicalDeviceShaderFloat16Int8Features struct {
	// ShaderFloat16 indicates whether 16-bit floats (halfs) are supported in shader code
	ShaderFloat16 bool
	// ShaderInt8 indicates whether 8-bit integer (signed and unsigned) are supported in
	// shader code
	ShaderInt8 bool

	common.NextOptions
	common.NextOutData
}

func (o *PhysicalDeviceShaderFloat16Int8Features) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceShaderFloat16Int8FeaturesKHR{})))
	}

	info := (*C.VkPhysicalDeviceShaderFloat16Int8FeaturesKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_FLOAT16_INT8_FEATURES_KHR
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceShaderFloat16Int8Features) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceShaderFloat16Int8FeaturesKHR)(cDataPointer)

	o.ShaderFloat16 = info.shaderFloat16 != C.VkBool32(0)
	o.ShaderInt8 = info.shaderInt8 != C.VkBool32(0)

	return info.pNext, nil
}

func (o PhysicalDeviceShaderFloat16Int8Features) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceShaderFloat16Int8FeaturesKHR{})))
	}

	info := (*C.VkPhysicalDeviceShaderFloat16Int8FeaturesKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_FLOAT16_INT8_FEATURES_KHR
	info.pNext = next
	info.shaderFloat16 = C.VkBool32(0)
	info.shaderInt8 = C.VkBool32(0)

	if o.ShaderFloat16 {
		info.shaderFloat16 = C.VkBool32(1)
	}

	if o.ShaderInt8 {
		info.shaderInt8 = C.VkBool32(1)
	}

	return preallocatedPointer, nil
}
