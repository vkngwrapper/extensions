package ext_shader_demote_to_helper_invocation

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

type PhysicalDeviceShaderDemoteToHelperInvocationFeatures struct {
	ShaderDemoteToHelperInvocation bool

	common.NextOptions
	common.NextOutData
}

func (o PhysicalDeviceShaderDemoteToHelperInvocationFeatures) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof([1]C.VkPhysicalDeviceShaderDemoteToHelperInvocationFeaturesEXT{})))
	}

	features := (*C.VkPhysicalDeviceShaderDemoteToHelperInvocationFeaturesEXT)(preallocatedPointer)
	features.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_DEMOTE_TO_HELPER_INVOCATION_FEATURES
	features.pNext = next
	features.shaderDemoteToHelperInvocation = C.VkBool32(0)
	if o.ShaderDemoteToHelperInvocation {
		features.shaderDemoteToHelperInvocation = C.VkBool32(1)
	}

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceShaderDemoteToHelperInvocationFeatures) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof([1]C.VkPhysicalDeviceShaderDemoteToHelperInvocationFeaturesEXT{})))
	}

	features := (*C.VkPhysicalDeviceShaderDemoteToHelperInvocationFeaturesEXT)(preallocatedPointer)
	features.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_DEMOTE_TO_HELPER_INVOCATION_FEATURES
	features.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceShaderDemoteToHelperInvocationFeatures) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	features := (*C.VkPhysicalDeviceShaderDemoteToHelperInvocationFeaturesEXT)(cDataPointer)
	o.ShaderDemoteToHelperInvocation = features.shaderDemoteToHelperInvocation != C.VkBool32(0)

	return features.pNext, nil
}
