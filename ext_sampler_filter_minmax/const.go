package ext_sampler_filter_minmax

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/vkngwrapper/core/v2/core1_0"
	_ "github.com/vkngwrapper/extensions/v3/vulkan"
)

// SamplerReductionMode specifies reduction mode for texture filtering
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerReductionMode.html
type SamplerReductionMode int32

var samplerReductionModeMapping = make(map[SamplerReductionMode]string)

func (e SamplerReductionMode) Register(str string) {
	samplerReductionModeMapping[e] = str
}

func (e SamplerReductionMode) String() string {
	return samplerReductionModeMapping[e]
}

////

const (
	// ExtensionName is "VK_EXT_sampler_filter_minmax"
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_sampler_filter_minmax.html
	ExtensionName string = C.VK_EXT_SAMPLER_FILTER_MINMAX_EXTENSION_NAME

	// FormatFeatureSampledImageFilterMinmax specifies the Image can be used as a sampled Image
	// with a min or max SamplerReductionMode
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatFeatureFlagBits.html
	FormatFeatureSampledImageFilterMinmax core1_0.FormatFeatureFlags = C.VK_FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_MINMAX_BIT_EXT

	// SamplerReductionModeMax specifies that texel values are combined by taking
	// the component-wise maximum of values in the footprint with non-zero weights
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerReductionMode.html
	SamplerReductionModeMax SamplerReductionMode = C.VK_SAMPLER_REDUCTION_MODE_MAX_EXT
	// SamplerReductionModeMin specifies that texel values are combined by taking the
	// component-wise minimum of values in the footprint with non-zero weights
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerReductionMode.html
	SamplerReductionModeMin SamplerReductionMode = C.VK_SAMPLER_REDUCTION_MODE_MIN_EXT
	// SamplerReductionModeWeightedAverage specifies that texel values are combined by
	// computing a weighted average of values in the footprint
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerReductionMode.html
	SamplerReductionModeWeightedAverage SamplerReductionMode = C.VK_SAMPLER_REDUCTION_MODE_WEIGHTED_AVERAGE_EXT
)

func init() {
	FormatFeatureSampledImageFilterMinmax.Register("Sampled Image Filter Min-Max")

	SamplerReductionModeMin.Register("Min")
	SamplerReductionModeMax.Register("Max")
	SamplerReductionModeWeightedAverage.Register("Weighted Average")
}
