package khr_shader_float_controls

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

// PhysicalDeviceFloatControlsProperties describes properties supported by khr_shader_float_controls
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFloatControlsProperties.html
type PhysicalDeviceFloatControlsProperties struct {
	// DenormBehaviorIndependence indicates whether, and how, denorm behavior can be
	// set independently for different bit widths
	DenormBehaviorIndependence ShaderFloatControlsIndependence
	// RoundingModeIndependence indicates whether, and how, rounding modes can be set indpendently
	// for different bit widths
	RoundingModeIndependence ShaderFloatControlsIndependence

	// ShaderSignedZeroInfNanPreserveFloat16 indicates whether the sign of zero, NaN, and +/- infinity
	// can be preserved in 16-bit floating-point computations
	ShaderSignedZeroInfNanPreserveFloat16 bool
	// ShaderSignedZeroInfNanPreserveFloat32 indicates whether the sign of zero, NaN, and +/- infinity
	// can be preserved in 32-bit floating-point computations
	ShaderSignedZeroInfNanPreserveFloat32 bool
	// ShaderSignedZeroInfNanPreserveFloat64 indicates whether the sign of zero, NaN, and +/- infinity
	// can be preserved in 64-bit floating-point computations
	ShaderSignedZeroInfNanPreserveFloat64 bool
	// ShaderDenormPreserveFloat16 indicates whether denormals can be preserved in 16-bit floating-point
	// computations
	ShaderDenormPreserveFloat16 bool
	// ShaderDenormPreserveFloat32 indicates whether denormals can be preserved in 32-bit floating-point
	// computations
	ShaderDenormPreserveFloat32 bool
	// ShaderDenormPreserveFloat64 indicates whether denormals can be preserved in 64-bit floating-point
	// computations
	ShaderDenormPreserveFloat64 bool
	// ShaderDenormFlushToZeroFloat16 indicates whether denormals can be flushed to zero in 16-bit
	// floating-point computations
	ShaderDenormFlushToZeroFloat16 bool
	// ShaderDenormFlushToZeroFloat32 indicates whether denormals can be flushed to zero in 32-bit
	// floating-point computations
	ShaderDenormFlushToZeroFloat32 bool
	// ShaderDenormFlushToZeroFloat64 indicates whether denormals can be flushed to zero in 64-bit
	// floating-point computations
	ShaderDenormFlushToZeroFloat64 bool
	// ShaderRoundingModeRTEFloat16 indicates whether an implementation supports the round-to-nearest-even
	// rounding mode for 16-bit floating-point arithmetic and conversion instructions
	ShaderRoundingModeRTEFloat16 bool
	// ShaderRoundingModeRTEFloat32 indicates whether an implementation supports the round-to-nearest-even
	// rounding mode for 32-bit floating-point arithmetic and conversion instructions
	ShaderRoundingModeRTEFloat32 bool
	// ShaderRoundingModeRTEFloat64 indicates whether an implementation supports the round-to-nearest-even
	// rounding mode for 64-bit floating-point arithmetic and conversion instructions
	ShaderRoundingModeRTEFloat64 bool
	// ShaderRoundingModeRTZFloat16 indicates whether an implementation supports the round-toward-zero
	// rounding mode for 16-bit floating-point arithmetic and conversion instructions
	ShaderRoundingModeRTZFloat16 bool
	// ShaderRoundingModeRTZFloat32 indicates whether an implementation supports the round-toward-zero
	// rounding mode for 32-bit floating-point arithmetic and conversion instructions
	ShaderRoundingModeRTZFloat32 bool
	// ShaderRoundingModeRTZFloat64 indicates whether an implementation supports the round-toward-zero
	// rounding mode for 64-bit floating-point arithmetic and conversion instructions
	ShaderRoundingModeRTZFloat64 bool

	common.NextOutData
}

func (o *PhysicalDeviceFloatControlsProperties) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceFloatControlsPropertiesKHR{})))
	}

	info := (*C.VkPhysicalDeviceFloatControlsPropertiesKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FLOAT_CONTROLS_PROPERTIES_KHR
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceFloatControlsProperties) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceFloatControlsPropertiesKHR)(cDataPointer)

	o.DenormBehaviorIndependence = ShaderFloatControlsIndependence(info.denormBehaviorIndependence)
	o.RoundingModeIndependence = ShaderFloatControlsIndependence(info.roundingModeIndependence)
	o.ShaderSignedZeroInfNanPreserveFloat16 = info.shaderSignedZeroInfNanPreserveFloat16 != C.VkBool32(0)
	o.ShaderSignedZeroInfNanPreserveFloat32 = info.shaderSignedZeroInfNanPreserveFloat32 != C.VkBool32(0)
	o.ShaderSignedZeroInfNanPreserveFloat64 = info.shaderSignedZeroInfNanPreserveFloat64 != C.VkBool32(0)
	o.ShaderDenormPreserveFloat16 = info.shaderDenormPreserveFloat16 != C.VkBool32(0)
	o.ShaderDenormPreserveFloat32 = info.shaderDenormPreserveFloat32 != C.VkBool32(0)
	o.ShaderDenormPreserveFloat64 = info.shaderDenormPreserveFloat64 != C.VkBool32(0)
	o.ShaderDenormFlushToZeroFloat16 = info.shaderDenormFlushToZeroFloat16 != C.VkBool32(0)
	o.ShaderDenormFlushToZeroFloat32 = info.shaderDenormFlushToZeroFloat32 != C.VkBool32(0)
	o.ShaderDenormFlushToZeroFloat64 = info.shaderDenormFlushToZeroFloat64 != C.VkBool32(0)
	o.ShaderRoundingModeRTEFloat16 = info.shaderRoundingModeRTEFloat16 != C.VkBool32(0)
	o.ShaderRoundingModeRTEFloat32 = info.shaderRoundingModeRTEFloat32 != C.VkBool32(0)
	o.ShaderRoundingModeRTEFloat64 = info.shaderRoundingModeRTEFloat64 != C.VkBool32(0)
	o.ShaderRoundingModeRTZFloat16 = info.shaderRoundingModeRTZFloat16 != C.VkBool32(0)
	o.ShaderRoundingModeRTZFloat32 = info.shaderRoundingModeRTZFloat32 != C.VkBool32(0)
	o.ShaderRoundingModeRTZFloat64 = info.shaderRoundingModeRTZFloat64 != C.VkBool32(0)

	return info.pNext, nil
}
