package ext_inline_uniform_block

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import "github.com/vkngwrapper/core/v3/core1_0"

const (
	// ExtensionName is "VK_EXT_inline_uniform_block"
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VK_EXT_inline_uniform_block.html
	ExtensionName string = C.VK_EXT_INLINE_UNIFORM_BLOCK_EXTENSION_NAME

	// DescriptorTypeInlineUniformBlock specifies an inline uniform block
	DescriptorTypeInlineUniformBlock core1_0.DescriptorType = C.VK_DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK_EXT
)
