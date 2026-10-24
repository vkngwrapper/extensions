package khr_shader_integer_dot_product

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

// PhysicalDeviceShaderIntegerDotProductFeatures describes integer dot product features
// that can be supported by an implementation
//
// https://docs.vulkan.org/refpages/latest/refpages/source/VkPhysicalDeviceShaderIntegerDotProductFeatures.html
type PhysicalDeviceShaderIntegerDotProductFeatures struct {
	// ShaderIntegerDotProduct specifies whether shader modules can declare the
	// DotProductInputAllKHR, DotProductInput4x8BitKHR, DotProductInput4x8BitPackedKHR and
	// DotProductKHR capabilities.
	ShaderIntegerDotProduct bool

	common.NextOptions
	common.NextOutData
}

func (o PhysicalDeviceShaderIntegerDotProductFeatures) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof([1]C.VkPhysicalDeviceShaderIntegerDotProductFeatures{})))
	}

	features := (*C.VkPhysicalDeviceShaderIntegerDotProductFeatures)(preallocatedPointer)
	features.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_INTEGER_DOT_PRODUCT_FEATURES
	features.pNext = next
	features.shaderIntegerDotProduct = C.VkBool32(0)
	if o.ShaderIntegerDotProduct {
		features.shaderIntegerDotProduct = C.VkBool32(1)
	}

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceShaderIntegerDotProductFeatures) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof([1]C.VkPhysicalDeviceShaderIntegerDotProductFeatures{})))
	}

	features := (*C.VkPhysicalDeviceShaderIntegerDotProductFeatures)(preallocatedPointer)
	features.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_INTEGER_DOT_PRODUCT_FEATURES
	features.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceShaderIntegerDotProductFeatures) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	features := (*C.VkPhysicalDeviceShaderIntegerDotProductFeatures)(cDataPointer)
	o.ShaderIntegerDotProduct = features.shaderIntegerDotProduct != C.VkBool32(0)

	return features.pNext, nil
}
