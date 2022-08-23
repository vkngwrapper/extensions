package ext_debug_utils

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/core1_0"
)

// DebugUtilsMessageTypeFlags specifies which types of events cause a debug messenger callback
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsMessageTypeFlagBitsEXT.html
type DebugUtilsMessageTypeFlags int32

var messageTypesMapping = common.NewFlagStringMapping[DebugUtilsMessageTypeFlags]()

func (f DebugUtilsMessageTypeFlags) Register(str string) {
	messageTypesMapping.Register(f, str)
}
func (f DebugUtilsMessageTypeFlags) String() string {
	return messageTypesMapping.FlagsToString(f)
}

////

// DebugUtilsMessageSeverityFlags specifies which severities of events cause a debug messenger
// callback
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsMessageSeverityFlagBitsEXT.html
type DebugUtilsMessageSeverityFlags int32

var messageSeveritiesMapping = common.NewFlagStringMapping[DebugUtilsMessageSeverityFlags]()

func (f DebugUtilsMessageSeverityFlags) Register(str string) {
	messageSeveritiesMapping.Register(f, str)
}
func (f DebugUtilsMessageSeverityFlags) String() string {
	return messageSeveritiesMapping.FlagsToString(f)
}

////

// DebugUtilsMessengerCreateFlags is reserved for future use
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsMessengerCreateFlagsEXT.html
type DebugUtilsMessengerCreateFlags int32

var callbackDataFlagsMapping = common.NewFlagStringMapping[DebugUtilsMessengerCreateFlags]()

func (f DebugUtilsMessengerCreateFlags) Register(str string) {
	callbackDataFlagsMapping.Register(f, str)
}
func (f DebugUtilsMessengerCreateFlags) String() string {
	return callbackDataFlagsMapping.FlagsToString(f)
}

////

const (
	// ExtensionName is "VK_EXT_debug_utils"
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_debug_utils.html
	ExtensionName string = C.VK_EXT_DEBUG_UTILS_EXTENSION_NAME

	// TypeGeneral specifies that some general event has occurred
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsMessageTypeFlagBitsEXT.html
	TypeGeneral DebugUtilsMessageTypeFlags = C.VK_DEBUG_UTILS_MESSAGE_TYPE_GENERAL_BIT_EXT
	// TypeValidation specifies that something has occurred during validation against
	// the Vulkan specification that may indicate invalid behavior
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsMessageTypeFlagBitsEXT.html
	TypeValidation DebugUtilsMessageTypeFlags = C.VK_DEBUG_UTILS_MESSAGE_TYPE_VALIDATION_BIT_EXT
	// TypePerformance specifies a potentially non-optimal use of Vulkan, e.g. using
	// CommandBuffer.CmdClearColorImage when setting AttachmentDescription.LoadOp to
	// AttachmentLoadOpClear would have worked
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsMessageTypeFlagBitsEXT.html
	TypePerformance DebugUtilsMessageTypeFlags = C.VK_DEBUG_UTILS_MESSAGE_TYPE_PERFORMANCE_BIT_EXT

	// SeverityVerbose specifies the most verbose output indicating all diagnostic messages
	// from the Vulkan loader, layers, and drivers should be captured
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsMessageSeverityFlagBitsEXT.html
	SeverityVerbose DebugUtilsMessageSeverityFlags = C.VK_DEBUG_UTILS_MESSAGE_SEVERITY_VERBOSE_BIT_EXT
	// SeverityInfo specifies an informational message such as resource details that may be
	// handy when debugging an application
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsMessageSeverityFlagBitsEXT.html
	SeverityInfo DebugUtilsMessageSeverityFlags = C.VK_DEBUG_UTILS_MESSAGE_SEVERITY_INFO_BIT_EXT
	// SeverityWarning specifies use of Vulkan that may expose an app bug
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsMessageSeverityFlagBitsEXT.html
	SeverityWarning DebugUtilsMessageSeverityFlags = C.VK_DEBUG_UTILS_MESSAGE_SEVERITY_WARNING_BIT_EXT
	// SeverityError specifies that the application has violated a valid usage condition
	// of the specification
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsMessageSeverityFlagBitsEXT.html
	SeverityError DebugUtilsMessageSeverityFlags = C.VK_DEBUG_UTILS_MESSAGE_SEVERITY_ERROR_BIT_EXT

	// ObjectTypeDebugUtilsMessenger specifies a DebugUtilsMessenger handle
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkObjectType.html
	ObjectTypeDebugUtilsMessenger core1_0.ObjectType = C.VK_OBJECT_TYPE_DEBUG_UTILS_MESSENGER_EXT
)

func init() {
	TypeGeneral.Register("General")
	TypeValidation.Register("Validation")
	TypePerformance.Register("Performance")

	SeverityVerbose.Register("Verbose")
	SeverityInfo.Register("Info")
	SeverityWarning.Register("Warning")
	SeverityError.Register("Error")

	ObjectTypeDebugUtilsMessenger.Register("Debug Utils DebugUtilsMessenger")
}
