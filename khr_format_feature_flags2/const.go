package khr_format_feature_flags2

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import "github.com/vkngwrapper/core/v3/common"

const (
	// ExtensionName is "VK_KHR_format_feature_flags2"
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VK_KHR_format_feature_flags2.html
	ExtensionName string = C.VK_KHR_FORMAT_FEATURE_FLAGS_2_EXTENSION_NAME
)

type FormatFeatureFlags2 uint32

var formatFeatureFlags2Mapping = common.NewFlagStringMapping[FormatFeatureFlags2]()

func (e FormatFeatureFlags2) Register(str string) {
	formatFeatureFlags2Mapping.Register(e, str)
}

func (e FormatFeatureFlags2) String() string {
	return formatFeatureFlags2Mapping.FlagsToString(e)
}

// This type (and seemingly ONLY this type) is defined as static const, which cgo doesn't play
// well with.  From go's point of view, this go code and the c preamble above are differen compilation
// units which means that static consts should not be visible from here.  As a result, these have to be
// manually defined.  May god have mercy on my soul.
const (
	FormatFeature2SampledImage                                                     = FormatFeatureFlags2(0x00000001)
	FormatFeature2StorageImage                                                     = FormatFeatureFlags2(0x00000002)
	FormatFeature2StorageImageAtomic                                               = FormatFeatureFlags2(0x00000004)
	FormatFeature2UniformTexelBuffer                                               = FormatFeatureFlags2(0x00000008)
	FormatFeature2StorageTexelBuffer                                               = FormatFeatureFlags2(0x00000010)
	FormatFeature2StorageTexelBufferAtomic                                         = FormatFeatureFlags2(0x00000020)
	FormatFeature2VertexBuffer                                                     = FormatFeatureFlags2(0x00000040)
	FormatFeature2ColorAttachment                                                  = FormatFeatureFlags2(0x00000080)
	FormatFeature2ColorAttachmentBlend                                             = FormatFeatureFlags2(0x00000100)
	FormatFeature2DepthStencilAttachment                                           = FormatFeatureFlags2(0x00000200)
	FormatFeature2BlitSrc                                                          = FormatFeatureFlags2(0x00000400)
	FormatFeature2BlitDst                                                          = FormatFeatureFlags2(0x00000800)
	FormatFeature2SampledImageFilterLinear                                         = FormatFeatureFlags2(0x00001000)
	FormatFeature2SampledImageFilterCubic                                          = FormatFeatureFlags2(0x00002000)
	FormatFeature2TransferSrc                                                      = FormatFeatureFlags2(0x00004000)
	FormatFeature2TransferDst                                                      = FormatFeatureFlags2(0x00008000)
	FormatFeature2SampledImageFilterMinmax                                         = FormatFeatureFlags2(0x00010000)
	FormatFeature2MidpointChromaSamples                                            = FormatFeatureFlags2(0x00020000)
	FormatFeature2SampledImageYcbcrConversionLinearFilter                          = FormatFeatureFlags2(0x00040000)
	FormatFeature2SampledImageYcbcrConversionSeparateReconstructionFilter          = FormatFeatureFlags2(0x00080000)
	FormatFeature2SampledImageYcbcrConversionChromaReconstructionExplicit          = FormatFeatureFlags2(0x00100000)
	FormatFeature2SampledImageYcbcrConversionChromaReconstructionExplicitForceable = FormatFeatureFlags2(0x00200000)
	FormatFeature2Disjoint                                                         = FormatFeatureFlags2(0x00400000)
	FormatFeature2CositedChromaSamples                                             = FormatFeatureFlags2(0x00800000)
	FormatFeature2StorageReadWithoutFormat                                         = FormatFeatureFlags2(0x80000000)
	FormatFeature2StorageWriteWithoutFormat                                        = FormatFeatureFlags2(0x10000000)
	FormatFeature2SampledImageDepthComparison                                      = FormatFeatureFlags2(0x20000000)
)

func init() {
	FormatFeature2SampledImage.Register("SampledImage")
	FormatFeature2StorageImage.Register("StorageImage")
	FormatFeature2StorageImageAtomic.Register("StorageImageAtomic")
	FormatFeature2UniformTexelBuffer.Register("UniformTexelBuffer")
	FormatFeature2StorageTexelBuffer.Register("StorageTexelBuffer")
	FormatFeature2StorageTexelBufferAtomic.Register("StorageTexelBufferAtomic")
	FormatFeature2VertexBuffer.Register("VertexBuffer")
	FormatFeature2ColorAttachment.Register("ColorAttachment")
	FormatFeature2ColorAttachmentBlend.Register("ColorAttachmentBlend")
	FormatFeature2DepthStencilAttachment.Register("DepthStencilAttachment")
	FormatFeature2BlitSrc.Register("BlitSrc")
	FormatFeature2BlitDst.Register("BlitDst")
	FormatFeature2SampledImageFilterLinear.Register("SampledImageFilterLinear")
	FormatFeature2TransferSrc.Register("TransferSrc")
	FormatFeature2TransferDst.Register("TransferDst")
	FormatFeature2MidpointChromaSamples.Register("MidpointChromaSamples")
	FormatFeature2SampledImageYcbcrConversionLinearFilter.Register("SampledImageYcbcrConversionLinearFilter")
	FormatFeature2SampledImageYcbcrConversionSeparateReconstructionFilter.Register("SampledImageYcbcrConversionSeparateReconstructionFilter")
	FormatFeature2SampledImageYcbcrConversionChromaReconstructionExplicit.Register("SampledImageYcbcrConversionChromaReconstructionExplicit")
	FormatFeature2SampledImageYcbcrConversionChromaReconstructionExplicitForceable.Register("SampledImageYcbcrConversionChromaReconstructionExplicitForceable")
	FormatFeature2Disjoint.Register("Disjoint")
	FormatFeature2CositedChromaSamples.Register("CositedChromaSamples")
	FormatFeature2StorageReadWithoutFormat.Register("StorageReadWithoutFormat")
	FormatFeature2StorageWriteWithoutFormat.Register("StorageWriteWithoutFormat")
	FormatFeature2SampledImageDepthComparison.Register("SampledImageDepthComparison")
	FormatFeature2SampledImageFilterMinmax.Register("SampledImageFilterMinmax")
	FormatFeature2SampledImageFilterCubic.Register("SampledImageFilterCubic")
}
