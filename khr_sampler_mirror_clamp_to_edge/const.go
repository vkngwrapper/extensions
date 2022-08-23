package khr_sampler_mirror_clamp_to_edge

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/vkngwrapper/core/v2/core1_0"
	_ "github.com/vkngwrapper/extensions/v2/vulkan"
)

const (
	// ExtensionName is "VK_KHR_sampler_mirror_clamp_to_edge"
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_sampler_mirror_clamp_to_edge.html
	ExtensionName string = C.VK_KHR_SAMPLER_MIRROR_CLAMP_TO_EDGE_EXTENSION_NAME

	// SamplerAddressModeMirrorClampToEdge specifies that the mirror clamp to edge wrap mode will
	// be used
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerAddressMode.html
	SamplerAddressModeMirrorClampToEdge core1_0.SamplerAddressMode = C.VK_SAMPLER_ADDRESS_MODE_MIRROR_CLAMP_TO_EDGE_KHR
)

func init() {
	SamplerAddressModeMirrorClampToEdge.Register("Mirror Clamp To Edge")
}
