package khr_portability_subset

/*
#define VK_ENABLE_BETA_EXTENSIONS 1
#include <stdlib.h>
#include "../vulkan/vulkan.h"
#include "../vulkan/vulkan_beta.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v2/common"
	"unsafe"
)

// PhysicalDevicePortabilitySubsetFeatures describes the features that may not be supported by
// an implementation of the Vulkan 1.0 Portability Subset
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevicePortabilitySubsetFeaturesKHR.html
type PhysicalDevicePortabilitySubsetFeatures struct {
	// ConstantAlphaColorBlendFactors indicates whether this implementation supports constant alpha
	// blend factors used as a source or destination color
	ConstantAlphaColorBlendFactors bool
	// Events indicates whether this implementation supports synchronization using Event objects
	Events bool
	// ImageViewFormatReinterpretation indicates whether this implementation supports an ImageView being
	// created with a texel format containing a different number of components, or a different number
	// of bits in each component, than the texel format of the underlying Image
	ImageViewFormatReinterpretation bool
	// ImageViewFormatSwizzle indicates whether this implementation supports remapping format components
	// using ImageViewCreateInfo.Components
	ImageViewFormatSwizzle bool
	// ImageView2DOn3DImage indicates whether this implementation supports an Image being created with
	// ImageCreate2DArrayCompatible set, permitting a 2D or 2D array ImageView to be created on a 3D
	// Image
	ImageView2DOn3DImage bool
	// MultisampleArrayImage indicates whether this implementation supports an Image being created as
	// a 2D array with multiple samples per texel
	MultisampleArrayImage bool
	// MutableComparisonSamplers indicates whether this implementation allows descriptors with comparison
	// samplers to be updated
	MutableComparisonSamplers bool
	// PointPolygons indicates whether this implementation supports rasterization using a point
	// polygon mode
	PointPolygons bool
	// SamplerMipLodBias indicates whether this implementation supports setting a mipmap LOD bias value
	// when creating a Sampler
	SamplerMipLodBias bool
	// SeparateStencilMaskRef indicates whether this implementation supports separate front and back
	// stencil ops reference values
	SeparateStencilMaskRef bool
	// ShaderSamplerRateInterpolationFunctions indicates whether this implementation supports fragment
	// shaders which use the InterpolationFunction capability and the extended instructions
	// InterpolateAtCentroid, InterpolateAtOffset, and InterpolateAtSample from the GLSL.std.450
	// extended instruction set
	ShaderSamplerRateInterpolationFunctions bool
	// TessellationIsolines indicates whether this implementation supports isoline output from the
	// tessellation stage of a graphics Pipeline
	TessellationIsolines bool
	// TessellationPointMode indicates whether this implementation supports point output from the tessellation
	// stage of a graphics Pipeline
	TessellationPointMode bool
	// TriangleFans indicates whether this implementation supports triangle fans primitive topology
	TriangleFans bool
	// VertexAttributeAccessBeyondStride indicates whether this implementation supports accessing a vertex
	// input attribute beyond the stride of the corresponding vertex input binding
	VertexAttributeAccessBeyondStride bool

	common.NextOptions
	common.NextOutData
}

func (o *PhysicalDevicePortabilitySubsetFeatures) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkPhysicalDevicePortabilitySubsetFeaturesKHR)
	}

	outData := (*C.VkPhysicalDevicePortabilitySubsetFeaturesKHR)(preallocatedPointer)
	outData.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PORTABILITY_SUBSET_FEATURES_KHR
	outData.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDevicePortabilitySubsetFeatures) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	outData := (*C.VkPhysicalDevicePortabilitySubsetFeaturesKHR)(cDataPointer)
	o.ConstantAlphaColorBlendFactors = outData.constantAlphaColorBlendFactors != C.VkBool32(0)
	o.Events = outData.events != C.VkBool32(0)
	o.ImageViewFormatReinterpretation = outData.imageViewFormatReinterpretation != C.VkBool32(0)
	o.ImageViewFormatSwizzle = outData.imageViewFormatSwizzle != C.VkBool32(0)
	o.ImageView2DOn3DImage = outData.imageView2DOn3DImage != C.VkBool32(0)
	o.MultisampleArrayImage = outData.multisampleArrayImage != C.VkBool32(0)
	o.MutableComparisonSamplers = outData.mutableComparisonSamplers != C.VkBool32(0)
	o.PointPolygons = outData.pointPolygons != C.VkBool32(0)
	o.SamplerMipLodBias = outData.samplerMipLodBias != C.VkBool32(0)
	o.SeparateStencilMaskRef = outData.separateStencilMaskRef != C.VkBool32(0)
	o.ShaderSamplerRateInterpolationFunctions = outData.shaderSampleRateInterpolationFunctions != C.VkBool32(0)
	o.TessellationIsolines = outData.tessellationIsolines != C.VkBool32(0)
	o.TessellationPointMode = outData.tessellationPointMode != C.VkBool32(0)
	o.TriangleFans = outData.triangleFans != C.VkBool32(0)
	o.VertexAttributeAccessBeyondStride = outData.vertexAttributeAccessBeyondStride != C.VkBool32(0)

	return outData.pNext, nil
}

func (o PhysicalDevicePortabilitySubsetFeatures) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkPhysicalDevicePortabilitySubsetFeaturesKHR)
	}

	outData := (*C.VkPhysicalDevicePortabilitySubsetFeaturesKHR)(preallocatedPointer)
	outData.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PORTABILITY_SUBSET_FEATURES_KHR
	outData.pNext = next
	outData.constantAlphaColorBlendFactors = C.VkBool32(0)
	outData.events = C.VkBool32(0)
	outData.imageViewFormatReinterpretation = C.VkBool32(0)
	outData.imageViewFormatSwizzle = C.VkBool32(0)
	outData.imageView2DOn3DImage = C.VkBool32(0)
	outData.multisampleArrayImage = C.VkBool32(0)
	outData.mutableComparisonSamplers = C.VkBool32(0)
	outData.pointPolygons = C.VkBool32(0)
	outData.samplerMipLodBias = C.VkBool32(0)
	outData.separateStencilMaskRef = C.VkBool32(0)
	outData.shaderSampleRateInterpolationFunctions = C.VkBool32(0)
	outData.tessellationIsolines = C.VkBool32(0)
	outData.tessellationPointMode = C.VkBool32(0)
	outData.triangleFans = C.VkBool32(0)
	outData.vertexAttributeAccessBeyondStride = C.VkBool32(0)

	if o.ConstantAlphaColorBlendFactors {
		outData.constantAlphaColorBlendFactors = C.VkBool32(1)
	}

	if o.Events {
		outData.events = C.VkBool32(1)
	}

	if o.ImageViewFormatReinterpretation {
		outData.imageViewFormatReinterpretation = C.VkBool32(1)
	}

	if o.ImageViewFormatSwizzle {
		outData.imageViewFormatSwizzle = C.VkBool32(1)
	}

	if o.ImageView2DOn3DImage {
		outData.imageView2DOn3DImage = C.VkBool32(1)
	}

	if o.MultisampleArrayImage {
		outData.multisampleArrayImage = C.VkBool32(1)
	}

	if o.MutableComparisonSamplers {
		outData.mutableComparisonSamplers = C.VkBool32(1)
	}

	if o.PointPolygons {
		outData.pointPolygons = C.VkBool32(1)
	}

	if o.SamplerMipLodBias {
		outData.samplerMipLodBias = C.VkBool32(1)
	}

	if o.SeparateStencilMaskRef {
		outData.separateStencilMaskRef = C.VkBool32(1)
	}

	if o.ShaderSamplerRateInterpolationFunctions {
		outData.shaderSampleRateInterpolationFunctions = C.VkBool32(1)
	}

	if o.TessellationIsolines {
		outData.tessellationIsolines = C.VkBool32(1)
	}

	if o.TessellationPointMode {
		outData.tessellationPointMode = C.VkBool32(1)
	}

	if o.TriangleFans {
		outData.triangleFans = C.VkBool32(1)
	}

	if o.VertexAttributeAccessBeyondStride {
		outData.vertexAttributeAccessBeyondStride = C.VkBool32(1)
	}

	return preallocatedPointer, nil
}
