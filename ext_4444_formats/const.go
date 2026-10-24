package ext_4444_formats

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import "github.com/vkngwrapper/core/v3/core1_0"

const (
	// ExtensionName is "VK_EXT_4444_formats"
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VK_EXT_4444_formats.html
	ExtensionName string = C.VK_EXT_4444_FORMATS_EXTENSION_NAME

	// FormatA4B4G4R4UnsignedNormalizedPacked16 specifies a four-component, 16-bit packed
	// unsigned normalized format that has a 4-bit A component in bits 12..15, a 4-bit B
	// component in bits 8..11, a 4-bit G component in bits 4..7, and a 4-bit R component in
	// bits 0..3.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkFormat.html
	FormatA4B4G4R4UnsignedNormalizedPacked16 core1_0.Format = C.VK_FORMAT_A4B4G4R4_UNORM_PACK16_EXT

	// FormatA4R4G4B4UnsignedNormalizedPacked16 specifies a four-component, 16-bit packed
	// unsigned normalized format that has a 4-bit A component in bits 12..15, a 4-bit R
	// component in bits 8..11, a 4-bit G component in bits 4..7, and a 4-bit B component in
	// bits 0..3.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkFormat.html
	FormatA4R4G4B4UnsignedNormalizedPacked16 core1_0.Format = C.VK_FORMAT_A4R4G4B4_UNORM_PACK16_EXT
)
