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

// PhysicalDeviceShaderIntegerDotProductProperties contains information about integer dot
// product support for a physical device
//
// https://docs.vulkan.org/refpages/latest/refpages/source/VkPhysicalDeviceShaderIntegerDotProductProperties.html
type PhysicalDeviceShaderIntegerDotProductProperties struct {
	// IntegerDotProduct8BitUnsignedAccelerated is true if the support for 8-bit unsigned
	// dot product operations using the OpUDotKHR SPIR-V instruction is accelerated
	IntegerDotProduct8BitUnsignedAccelerated bool
	// IntegerDotProduct8BitSignedAccelerated is true if the support for 8-bit signed dot
	// product operations using the OpSDotKHR SPIR-V instruction is accelerated
	IntegerDotProduct8BitSignedAccelerated bool
	// IntegerDotProduct8BitMixedSignednessAccelerated is true if the support for 8-bit
	// mixed signedness dot product operations using the OpSUDotKHR SPIR-V instruction is
	// accelerated
	IntegerDotProduct8BitMixedSignednessAccelerated bool
	// IntegerDotProduct4x8BitPackedUnsignedAccelerated is true if the support for 8-bit
	// unsigned dot product operations from operands packed into 32-bit integers using the
	// OpUDotKHR SPIR-V instruction is accelerated
	IntegerDotProduct4x8BitPackedUnsignedAccelerated bool
	// IntegerDotProduct4x8BitPackedSignedAccelerated is true if the support for 8-bit
	// signed dot product operations from operands packed into 32-bit integers using the
	// OpSDotKHR SPIR-V instruction is accelerated
	IntegerDotProduct4x8BitPackedSignedAccelerated bool
	// IntegerDotProduct4x8BitPackedMixedSignednessAccelerated is true if the support for
	// 8-bit mixed signedness dot product operations from operands packed into 32-bit
	// integers using the OpSUDotKHR SPIR-V instruction is accelerated
	IntegerDotProduct4x8BitPackedMixedSignednessAccelerated bool
	// IntegerDotProduct16BitUnsignedAccelerated is true if the support for 16-bit unsigned
	// dot product operations using the OpUDotKHR SPIR-V instruction is accelerated
	IntegerDotProduct16BitUnsignedAccelerated bool
	// IntegerDotProduct16BitSignedAccelerated is true if the support for 16-bit signed
	// dot product operations using the OpSDotKHR SPIR-V instruction is accelerated
	IntegerDotProduct16BitSignedAccelerated bool
	// IntegerDotProduct16BitMixedSignednessAccelerated is true if the support for 16-bit
	// mixed signedness dot product operations using the OpSUDotKHR SPIR-V instruction is
	// accelerated
	IntegerDotProduct16BitMixedSignednessAccelerated bool
	// IntegerDotProduct32BitUnsignedAccelerated is true if the support for 32-bit unsigned
	// dot product operations using the OpUDotKHR SPIR-V instruction is accelerated
	IntegerDotProduct32BitUnsignedAccelerated bool
	// IntegerDotProduct32BitSignedAccelerated is true if the support for 32-bit signed
	// dot product operations using the OpSDotKHR SPIR-V instruction is accelerated
	IntegerDotProduct32BitSignedAccelerated bool
	// IntegerDotProduct32BitMixedSignednessAccelerated is true if the support for 32-bit
	// mixed signedness dot product operations using the OpSUDotKHR SPIR-V instruction is
	// accelerated
	IntegerDotProduct32BitMixedSignednessAccelerated bool
	// IntegerDotProduct64BitUnsignedAccelerated is true if if the support for 64-bit
	// unsigned dot product operations using the OpUDotKHR SPIR-V instruction is
	// accelerated
	IntegerDotProduct64BitUnsignedAccelerated bool
	// IntegerDotProduct64BitSignedAccelerated is true if the support for 64-bit signed
	// dot product operations using the OpSDotKHR SPIR-V instruction is accelerated
	IntegerDotProduct64BitSignedAccelerated bool
	// IntegerDotProduct64BitMixedSignednessAccelerated is true if the support for 64-bit
	// mixed signedness dot product operations using the OpSUDotKHR SPIR-V instruction is
	// accelerated
	IntegerDotProduct64BitMixedSignednessAccelerated bool
	// IntegerDotProductAccumulatingSaturating8BitUnsignedAccelerated is true if the
	// support for 8-bit unsigned accumulating saturating dot product operations using the
	// OpUDotAccSatKHR SPIR-V instruction is accelerated
	IntegerDotProductAccumulatingSaturating8BitUnsignedAccelerated bool
	// IntegerDotProductAccumulatingSaturating8BitSignedAccelerated is true if the support
	// for 8-bit signed accumulating saturating dot product operations using the
	// OpSDotAccSatKHR SPIR-V instruction is accelerated
	IntegerDotProductAccumulatingSaturating8BitSignedAccelerated bool
	// IntegerDotProductAccumulatingSaturating8BitMixedSignednessAccelerated is true if
	// the support for 8-bit mixed signedness accumulating saturating dot product
	// operations using the OpSUDotAccSatKHR SPIR-V instruction is accelerated
	IntegerDotProductAccumulatingSaturating8BitMixedSignednessAccelerated bool
	// IntegerDotProductAccumulatingSaturating4x8BitPackedUnsignedAccelerated is true if
	// the support for 8-bit unsigned accumulating saturating dot product operations from
	// operands packed into 32-bit integers using the OpUDotAccSatKHR SPIR-V instruction is
	// accelerated
	IntegerDotProductAccumulatingSaturating4x8BitPackedUnsignedAccelerated bool
	// IntegerDotProductAccumulatingSaturating4x8BitPackedSignedAccelerated is true if the
	// support for 8-bit signed accumulating saturating dot product operations from
	// operands packed into 32-bit integers using the OpSDotAccSatKHR SPIR-V instruction is
	// accelerated
	IntegerDotProductAccumulatingSaturating4x8BitPackedSignedAccelerated bool
	// IntegerDotProductAccumulatingSaturating4x8BitPackedMixedSignednessAccelerated is
	// true if the support for 8-bit mixed signedness accumulating saturating dot product
	// operations from operands packed into 32-bit integers using the OpSUDotAccSatKHR
	// SPIR-V instruction is accelerated
	IntegerDotProductAccumulatingSaturating4x8BitPackedMixedSignednessAccelerated bool
	// IntegerDotProductAccumulatingSaturating16BitUnsignedAccelerated is true if the
	// support for 16-bit unsigned accumulating saturating dot product operations using
	// the OpUDotAccSatKHR SPIR-V instruction is accelerated
	IntegerDotProductAccumulatingSaturating16BitUnsignedAccelerated bool
	// IntegerDotProductAccumulatingSaturating16BitSignedAccelerated is true if the support
	// for 16-bit signed accumulating saturating dot product operations using the
	// OpSDotAccSatKHR SPIR-V instruction is accelerated
	IntegerDotProductAccumulatingSaturating16BitSignedAccelerated bool
	// IntegerDotProductAccumulatingSaturating16BitMixedSignednessAccelerated is true if
	// the support for 16-bit mixed signedness accumulating saturating dot product
	// operations using the OpSUDotAccSatKHR SPIR-V instruction is accelerated
	IntegerDotProductAccumulatingSaturating16BitMixedSignednessAccelerated bool
	// IntegerDotProductAccumulatingSaturating32BitUnsignedAccelerated is true if the
	// support for 32-bit unsigned accumulating saturating dot product operations using the
	// OpUDotAccSatKHR SPIR-V instruction is accelerated
	IntegerDotProductAccumulatingSaturating32BitUnsignedAccelerated bool
	// IntegerDotProductAccumulatingSaturating32BitSignedAccelerated is true if the support
	// for 32-bit signed accumulating saturating dot product operations using the
	// OpSDotAccSatKHR SPIR-V instruction is accelerated
	IntegerDotProductAccumulatingSaturating32BitSignedAccelerated bool
	// IntegerDotProductAccumulatingSaturating32BitMixedSignednessAccelerated is true if
	// the support for 32-bit mixed signedness accumulating saturating dot product
	// operations using the OpSUDotAccSatKHR SPIR-V instruction is accelerated
	IntegerDotProductAccumulatingSaturating32BitMixedSignednessAccelerated bool
	// IntegerDotProductAccumulatingSaturating64BitUnsignedAccelerated is true if the
	// support for 64-bit unsigned accumulating saturating dot product operations using
	// the OpUDotAccSatKHR SPIR-V instruction is accelerated
	IntegerDotProductAccumulatingSaturating64BitUnsignedAccelerated bool
	// IntegerDotProductAccumulatingSaturating64BitSignedAccelerated is true if the support
	// for 64-bit signed accumulating saturating dot product operations using the
	// OpSDotAccSatKHR SPIR-V instruction is accelerated
	IntegerDotProductAccumulatingSaturating64BitSignedAccelerated bool
	// IntegerDotProductAccumulatingSaturating64BitMixedSignednessAccelerated is true if
	// the support for 64-bit mixed signedness accumulating saturating dot product
	// operations using the OpSUDotAccSatKHR SPIR-V instruction is accelerated
	IntegerDotProductAccumulatingSaturating64BitMixedSignednessAccelerated bool

	common.NextOutData
}

func (o *PhysicalDeviceShaderIntegerDotProductProperties) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof([1]C.VkPhysicalDeviceShaderIntegerDotProductProperties{})))
	}

	features := (*C.VkPhysicalDeviceShaderIntegerDotProductProperties)(preallocatedPointer)
	features.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_INTEGER_DOT_PRODUCT_FEATURES
	features.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceShaderIntegerDotProductProperties) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	features := (*C.VkPhysicalDeviceShaderIntegerDotProductProperties)(cDataPointer)
	o.IntegerDotProduct8BitUnsignedAccelerated = features.integerDotProduct8BitUnsignedAccelerated != C.VkBool32(0)
	o.IntegerDotProduct8BitSignedAccelerated = features.integerDotProduct8BitSignedAccelerated != C.VkBool32(0)
	o.IntegerDotProduct8BitMixedSignednessAccelerated = features.integerDotProduct8BitMixedSignednessAccelerated != C.VkBool32(0)
	o.IntegerDotProduct4x8BitPackedUnsignedAccelerated = features.integerDotProduct4x8BitPackedUnsignedAccelerated != C.VkBool32(0)
	o.IntegerDotProduct4x8BitPackedSignedAccelerated = features.integerDotProduct4x8BitPackedSignedAccelerated != C.VkBool32(0)
	o.IntegerDotProduct4x8BitPackedMixedSignednessAccelerated = features.integerDotProduct4x8BitPackedMixedSignednessAccelerated != C.VkBool32(0)
	o.IntegerDotProduct16BitUnsignedAccelerated = features.integerDotProduct16BitUnsignedAccelerated != C.VkBool32(0)
	o.IntegerDotProduct16BitSignedAccelerated = features.integerDotProduct16BitSignedAccelerated != C.VkBool32(0)
	o.IntegerDotProduct16BitMixedSignednessAccelerated = features.integerDotProduct16BitMixedSignednessAccelerated != C.VkBool32(0)
	o.IntegerDotProduct32BitUnsignedAccelerated = features.integerDotProduct32BitUnsignedAccelerated != C.VkBool32(0)
	o.IntegerDotProduct32BitSignedAccelerated = features.integerDotProduct32BitSignedAccelerated != C.VkBool32(0)
	o.IntegerDotProduct32BitMixedSignednessAccelerated = features.integerDotProduct32BitMixedSignednessAccelerated != C.VkBool32(0)
	o.IntegerDotProduct64BitUnsignedAccelerated = features.integerDotProduct64BitUnsignedAccelerated != C.VkBool32(0)
	o.IntegerDotProduct64BitSignedAccelerated = features.integerDotProduct64BitSignedAccelerated != C.VkBool32(0)
	o.IntegerDotProduct64BitMixedSignednessAccelerated = features.integerDotProduct64BitMixedSignednessAccelerated != C.VkBool32(0)
	o.IntegerDotProductAccumulatingSaturating8BitUnsignedAccelerated = features.integerDotProductAccumulatingSaturating8BitUnsignedAccelerated != C.VkBool32(0)
	o.IntegerDotProductAccumulatingSaturating8BitSignedAccelerated = features.integerDotProductAccumulatingSaturating8BitSignedAccelerated != C.VkBool32(0)
	o.IntegerDotProductAccumulatingSaturating8BitMixedSignednessAccelerated = features.integerDotProductAccumulatingSaturating8BitMixedSignednessAccelerated != C.VkBool32(0)
	o.IntegerDotProductAccumulatingSaturating4x8BitPackedUnsignedAccelerated = features.integerDotProductAccumulatingSaturating4x8BitPackedUnsignedAccelerated != C.VkBool32(0)
	o.IntegerDotProductAccumulatingSaturating4x8BitPackedSignedAccelerated = features.integerDotProductAccumulatingSaturating4x8BitPackedSignedAccelerated != C.VkBool32(0)
	o.IntegerDotProductAccumulatingSaturating4x8BitPackedMixedSignednessAccelerated = features.integerDotProductAccumulatingSaturating4x8BitPackedMixedSignednessAccelerated != C.VkBool32(0)
	o.IntegerDotProductAccumulatingSaturating16BitUnsignedAccelerated = features.integerDotProductAccumulatingSaturating16BitUnsignedAccelerated != C.VkBool32(0)
	o.IntegerDotProductAccumulatingSaturating16BitSignedAccelerated = features.integerDotProductAccumulatingSaturating16BitSignedAccelerated != C.VkBool32(0)
	o.IntegerDotProductAccumulatingSaturating16BitMixedSignednessAccelerated = features.integerDotProductAccumulatingSaturating16BitMixedSignednessAccelerated != C.VkBool32(0)
	o.IntegerDotProductAccumulatingSaturating32BitUnsignedAccelerated = features.integerDotProductAccumulatingSaturating32BitUnsignedAccelerated != C.VkBool32(0)
	o.IntegerDotProductAccumulatingSaturating32BitSignedAccelerated = features.integerDotProductAccumulatingSaturating32BitSignedAccelerated != C.VkBool32(0)
	o.IntegerDotProductAccumulatingSaturating32BitMixedSignednessAccelerated = features.integerDotProductAccumulatingSaturating32BitMixedSignednessAccelerated != C.VkBool32(0)
	o.IntegerDotProductAccumulatingSaturating64BitUnsignedAccelerated = features.integerDotProductAccumulatingSaturating64BitUnsignedAccelerated != C.VkBool32(0)
	o.IntegerDotProductAccumulatingSaturating64BitSignedAccelerated = features.integerDotProductAccumulatingSaturating64BitSignedAccelerated != C.VkBool32(0)
	o.IntegerDotProductAccumulatingSaturating64BitMixedSignednessAccelerated = features.integerDotProductAccumulatingSaturating64BitMixedSignednessAccelerated != C.VkBool32(0)

	return features.pNext, nil
}
