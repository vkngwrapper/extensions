package khr_shader_atomic_int64

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

// PhysicalDeviceShaderAtomicInt64Features describes features supported by khr_shader_atomic_int64
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderAtomicInt64Features.html
type PhysicalDeviceShaderAtomicInt64Features struct {
	// ShaderBufferInt64Atomics indicates whether shaders can perform 64-bit unsigned and signed
	// integer atomic operations on Buffer objects
	ShaderBufferInt64Atomics bool
	// ShaderSharedInt64Atomics indicates whether shaders can 64-bit unsigned and signed integer
	// atomic operations on shared memory
	ShaderSharedInt64Atomics bool

	common.NextOptions
	common.NextOutData
}

func (o *PhysicalDeviceShaderAtomicInt64Features) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceShaderAtomicInt64FeaturesKHR{})))
	}

	info := (*C.VkPhysicalDeviceShaderAtomicInt64FeaturesKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_ATOMIC_INT64_FEATURES_KHR
	info.pNext = next
	info.shaderBufferInt64Atomics = C.VkBool32(0)
	info.shaderSharedInt64Atomics = C.VkBool32(0)

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceShaderAtomicInt64Features) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceShaderAtomicInt64FeaturesKHR)(cDataPointer)

	o.ShaderBufferInt64Atomics = info.shaderBufferInt64Atomics != C.VkBool32(0)
	o.ShaderSharedInt64Atomics = info.shaderSharedInt64Atomics != C.VkBool32(0)

	return info.pNext, nil
}

func (o PhysicalDeviceShaderAtomicInt64Features) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceShaderAtomicInt64FeaturesKHR{})))
	}

	info := (*C.VkPhysicalDeviceShaderAtomicInt64FeaturesKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_ATOMIC_INT64_FEATURES_KHR
	info.pNext = next
	info.shaderBufferInt64Atomics = C.VkBool32(0)
	info.shaderSharedInt64Atomics = C.VkBool32(0)

	if o.ShaderBufferInt64Atomics {
		info.shaderBufferInt64Atomics = C.VkBool32(1)
	}

	if o.ShaderSharedInt64Atomics {
		info.shaderSharedInt64Atomics = C.VkBool32(1)
	}

	return preallocatedPointer, nil
}
