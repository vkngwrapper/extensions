package khr_shader_terminate_invocation

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

// PhysicalDeviceShaderTerminateInvocationFeatures describes support for the SPIR-V
// SPV_KHR_terminate_invocation extension
//
// https://docs.vulkan.org/refpages/latest/refpages/source/VkPhysicalDeviceShaderTerminateInvocationFeatures.html
type PhysicalDeviceShaderTerminateInvocationFeatures struct {
	// ShaderTerminateInvocation specifies whether the implementation supports SPIR-V
	// modules that use the SPV_KHR_terminate_invocation extension
	ShaderTerminateInvocation bool

	common.NextOutData
	common.NextOptions
}

func (o PhysicalDeviceShaderTerminateInvocationFeatures) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof([1]C.VkPhysicalDeviceShaderTerminateInvocationFeatures{})))
	}

	features := (*C.VkPhysicalDeviceShaderTerminateInvocationFeatures)(preallocatedPointer)
	features.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_TERMINATE_INVOCATION_FEATURES
	features.pNext = next
	features.shaderTerminateInvocation = C.VkBool32(0)
	if o.ShaderTerminateInvocation {
		features.shaderTerminateInvocation = C.VkBool32(1)
	}

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceShaderTerminateInvocationFeatures) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof([1]C.VkPhysicalDeviceShaderTerminateInvocationFeatures{})))
	}

	features := (*C.VkPhysicalDeviceShaderTerminateInvocationFeatures)(preallocatedPointer)
	features.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_TERMINATE_INVOCATION_FEATURES
	features.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceShaderTerminateInvocationFeatures) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	features := (*C.VkPhysicalDeviceShaderTerminateInvocationFeatures)(cDataPointer)
	o.ShaderTerminateInvocation = features.shaderTerminateInvocation != C.VkBool32(0)

	return features.pNext, nil
}
