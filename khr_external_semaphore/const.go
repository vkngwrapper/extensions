package khr_external_semaphore

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import "github.com/vkngwrapper/core/common"
import _ "github.com/vkngwrapper/extensions/vulkan"

// SemaphoreImportFlags specifies additional parameters of Semaphore payload import
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreImportFlagBits.html
type SemaphoreImportFlags int32

var semaphoreImportFlagsMapping = common.NewFlagStringMapping[SemaphoreImportFlags]()

func (f SemaphoreImportFlags) Register(str string) {
	semaphoreImportFlagsMapping.Register(f, str)
}

func (f SemaphoreImportFlags) String() string {
	return semaphoreImportFlagsMapping.FlagsToString(f)
}

////

const (
	// ExtensionName is "VK_KHR_external_semaphore"
	ExtensionName string = C.VK_KHR_EXTERNAL_SEMAPHORE_EXTENSION_NAME

	// SemaphoreImportTemporary specifies that the Semaphore payload will be imported only
	// temporarily, regardless of the permanence of the handle type
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreImportFlagBits.html
	SemaphoreImportTemporary SemaphoreImportFlags = C.VK_SEMAPHORE_IMPORT_TEMPORARY_BIT_KHR
)

func init() {
	SemaphoreImportTemporary.Register("Temporary")
}
