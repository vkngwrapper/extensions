package khr_portability_enumeration

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
	// ExtensionName is "VK_KHR_portability_enumeration"
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_portability_enumeration.html
	ExtensionName string = C.VK_KHR_PORTABILITY_ENUMERATION_EXTENSION_NAME

	// InstanceCreateEnumeratePortability specifies that the Instance will enumerate
	// available Vulkan Portability-compliant PhysicalDevice objects and groups in
	// addition to the Vulkan PhysicalDevice objects and groups that are enumerated
	// by default
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstanceCreateFlagBits.html
	InstanceCreateEnumeratePortability core1_0.InstanceCreateFlags = C.VK_INSTANCE_CREATE_ENUMERATE_PORTABILITY_BIT_KHR
)

func init() {
	InstanceCreateEnumeratePortability.Register("Enumerate Portability")
}
