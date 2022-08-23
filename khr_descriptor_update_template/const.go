package khr_descriptor_update_template

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/core1_0"
)

// DescriptorUpdateTemplateType indicates the valid usage of the DescriptorUpdateTemplate
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorUpdateTemplateType.html
type DescriptorUpdateTemplateType int32

var descriptorTemplateTypeMapping = make(map[DescriptorUpdateTemplateType]string)

func (e DescriptorUpdateTemplateType) Register(str string) {
	descriptorTemplateTypeMapping[e] = str
}

func (e DescriptorUpdateTemplateType) String() string {
	return descriptorTemplateTypeMapping[e]
}

////

// DescriptorUpdateTemplateCreateFlags is reserved for future use
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorUpdateTemplateCreateFlags.html
type DescriptorUpdateTemplateCreateFlags int32

var descriptorTemplateFlagsMapping = common.NewFlagStringMapping[DescriptorUpdateTemplateCreateFlags]()

func (f DescriptorUpdateTemplateCreateFlags) Register(str string) {
	descriptorTemplateFlagsMapping.Register(f, str)
}
func (f DescriptorUpdateTemplateCreateFlags) String() string {
	return descriptorTemplateFlagsMapping.FlagsToString(f)
}

////

const (
	// ExtensionName is "VK_KHR_descriptor_update_template"
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_descriptor_update_template.html
	ExtensionName string = C.VK_KHR_DESCRIPTOR_UPDATE_TEMPLATE_EXTENSION_NAME

	// DescriptorUpdateTemplateTypeDescriptorSet indicates the valid usage of the DescriptorUpdateTemplate
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorUpdateTemplateType.html
	DescriptorUpdateTemplateTypeDescriptorSet DescriptorUpdateTemplateType = C.VK_DESCRIPTOR_UPDATE_TEMPLATE_TYPE_DESCRIPTOR_SET_KHR

	// ObjectTypeDescriptorUpdateTemplate specifies a DescriptorUpdateTemplate handle
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkObjectType.html
	ObjectTypeDescriptorUpdateTemplate core1_0.ObjectType = C.VK_OBJECT_TYPE_DESCRIPTOR_UPDATE_TEMPLATE_KHR
)

func init() {
	DescriptorUpdateTemplateTypeDescriptorSet.Register("Descriptor Set")

	ObjectTypeDescriptorUpdateTemplate.Register("Descriptor Template")
}
