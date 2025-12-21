package ext_descriptor_indexing

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	_ "github.com/vkngwrapper/extensions/v3/vulkan"
)

// DescriptorBindingFlags specifies DescriptorSetLayout binding properties
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorBindingFlagBits.html
type DescriptorBindingFlags int32

var descriptorBindingFlagsMapping = common.NewFlagStringMapping[DescriptorBindingFlags]()

func (f DescriptorBindingFlags) Register(str string) {
	descriptorBindingFlagsMapping.Register(f, str)
}
func (f DescriptorBindingFlags) String() string {
	return descriptorBindingFlagsMapping.FlagsToString(f)
}

////

const (
	// ExtensionName is "VK_EXT_descriptor_indexing"
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_descriptor_indexing.html
	ExtensionName string = C.VK_EXT_DESCRIPTOR_INDEXING_EXTENSION_NAME

	// DescriptorBindingPartiallyBound indicates that descriptors in this binding that are
	// not dynamically used need not contain valid descriptors at the time the descriptors
	// are consumed
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorBindingFlagBits.html
	DescriptorBindingPartiallyBound DescriptorBindingFlags = C.VK_DESCRIPTOR_BINDING_PARTIALLY_BOUND_BIT_EXT
	// DescriptorBindingUpdateAfterBind indicates that if descriptors in this binding are updated
	// between when the DescriptorSet is bound in a CommandBuffer and when that CommandBuffer is
	// submitted to a Queue, then the submission will use the most recently-set descriptors
	// for this binding and the updates do not invalidate the CommandBuffer
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorBindingFlagBits.html
	DescriptorBindingUpdateAfterBind DescriptorBindingFlags = C.VK_DESCRIPTOR_BINDING_UPDATE_AFTER_BIND_BIT_EXT
	// DescriptorBindingUpdateUnusedWhilePending indicates that descriptors in this binding can be
	// updated after a CommandBuffer has bound this DescriptorSet, or while a CommandBuffer that
	// uses this DescriptorSet is pending execution, as long as the descriptors that are updated
	// are not used by those CommandBuffer objects
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorBindingFlagBits.html
	DescriptorBindingUpdateUnusedWhilePending DescriptorBindingFlags = C.VK_DESCRIPTOR_BINDING_UPDATE_UNUSED_WHILE_PENDING_BIT_EXT
	// DescriptorBindingVariableDescriptorCount indicates that this is a variable-sized descriptor
	// binding whose size will be specified when a DescriptorSet is allocated using this layout
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorBindingFlagBits.html
	DescriptorBindingVariableDescriptorCount DescriptorBindingFlags = C.VK_DESCRIPTOR_BINDING_VARIABLE_DESCRIPTOR_COUNT_BIT_EXT

	// DescriptorPoolCreateUpdateAfterBind specifies that DescriptorSet objects allocated from this
	// pool can include bindings with DescriptorBindingUpdateAfterBind
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorPoolCreateFlagBits.html
	DescriptorPoolCreateUpdateAfterBind core1_0.DescriptorPoolCreateFlags = C.VK_DESCRIPTOR_POOL_CREATE_UPDATE_AFTER_BIND_BIT_EXT

	// DescriptorSetLayoutCreateUpdateAfterBindPool specifies that DescriptorSet objects using this
	// layout must be allocated from a DescriptorPool created with DescriptorPoolCreateUpdateAfterBind
	// set
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayoutCreateFlagBits.html
	DescriptorSetLayoutCreateUpdateAfterBindPool core1_0.DescriptorSetLayoutCreateFlags = C.VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT_EXT

	// VkErrorFragmentation indicates a DescriptorPool creation has failed due to fragmentation
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResult.html
	VkErrorFragmentation common.VkResult = C.VK_ERROR_FRAGMENTATION_EXT
)

func init() {
	DescriptorBindingPartiallyBound.Register("Partially-Bound")
	DescriptorBindingUpdateAfterBind.Register("Update After Bind")
	DescriptorBindingUpdateUnusedWhilePending.Register("Update Unused While Pending")
	DescriptorBindingVariableDescriptorCount.Register("Variable Descriptor Count")

	DescriptorPoolCreateUpdateAfterBind.Register("Update After Bind")

	DescriptorSetLayoutCreateUpdateAfterBindPool.Register("Update After Bind Pool")

	VkErrorFragmentation.Register("fragmentation")
}
