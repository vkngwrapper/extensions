package khr_depth_stencil_resolve

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/vkngwrapper/core/v3/common"
	_ "github.com/vkngwrapper/extensions/v3/vulkan"
)

// ResolveModeFlags indicates supported depth and stencil resolve modes
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResolveModeFlagBits.html
type ResolveModeFlags int32

var resolveModeFlagsMapping = common.NewFlagStringMapping[ResolveModeFlags]()

func (f ResolveModeFlags) Register(str string) {
	resolveModeFlagsMapping.Register(f, str)
}

func (f ResolveModeFlags) String() string {
	return resolveModeFlagsMapping.FlagsToString(f)
}

////

const (
	// ExtensionName is "VK_KHR_depth_stencil_resolve"
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_depth_stencil_resolve.html
	ExtensionName string = C.VK_KHR_DEPTH_STENCIL_RESOLVE_EXTENSION_NAME

	// ResolveModeAverage indicates that the result of the resolve operation is the average
	// of the sample values
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResolveModeFlagBits.html
	ResolveModeAverage ResolveModeFlags = C.VK_RESOLVE_MODE_AVERAGE_BIT_KHR
	// ResolveModeMax indicates that the result of the resolve operation is the maximum of the
	// sample values
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResolveModeFlagBits.html
	ResolveModeMax ResolveModeFlags = C.VK_RESOLVE_MODE_MAX_BIT_KHR
	// ResolveModeMin indicates that the result of the resolve operation is the minimum of the
	// sample values
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResolveModeFlagBits.html
	ResolveModeMin ResolveModeFlags = C.VK_RESOLVE_MODE_MIN_BIT_KHR
	// ResolveModeNone indicates that no resolve operation is performed
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResolveModeFlagBits.html
	ResolveModeNone ResolveModeFlags = C.VK_RESOLVE_MODE_NONE_KHR
	// ResolveModeSampleZero indicates that the result of the resolve operation is equal to
	// the value of sample 0
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResolveModeFlagBits.html
	ResolveModeSampleZero ResolveModeFlags = C.VK_RESOLVE_MODE_SAMPLE_ZERO_BIT_KHR
)

func init() {
	ResolveModeAverage.Register("Average")
	ResolveModeMax.Register("Max")
	ResolveModeMin.Register("Min")
	ResolveModeNone.Register("None")
	ResolveModeSampleZero.Register("Sample Zero")
}
