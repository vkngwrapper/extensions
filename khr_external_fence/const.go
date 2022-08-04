package khr_external_fence

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import "github.com/vkngwrapper/core/common"

// FenceImportFlags specifies additional parameters of a Fence payload import
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFenceImportFlagBits.html
type FenceImportFlags int32

var fenceImportFlagsMapping = common.NewFlagStringMapping[FenceImportFlags]()

func (f FenceImportFlags) Register(str string) {
	fenceImportFlagsMapping.Register(f, str)
}

func (f FenceImportFlags) String() string {
	return fenceImportFlagsMapping.FlagsToString(f)
}

////

const (
	// ExtensionName is "VK_KHR_external_fence"
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_external_fence.html
	ExtensionName string = C.VK_KHR_EXTERNAL_FENCE_EXTENSION_NAME

	// FenceImportTemporary specifies that the Fence payload will be imported only temporarily,
	// regardless of the permanence of HandleType
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFenceImportFlagBits.html
	FenceImportTemporary FenceImportFlags = C.VK_FENCE_IMPORT_TEMPORARY_BIT_KHR
)

func init() {
	FenceImportTemporary.Register("Temporary")
}
