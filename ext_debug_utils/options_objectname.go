package ext_debug_utils

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/core1_0"
	"github.com/vkngwrapper/core/v2/driver"
	"unsafe"
)

// DebugUtilsObjectNameInfo specifies parameters of a name to give to an object
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsObjectNameInfoEXT.html
type DebugUtilsObjectNameInfo struct {
	// ObjectName is a string specifying the name to apply to ObjectHandle
	ObjectName string
	// ObjectHandle is the handle of the object to be named
	ObjectHandle driver.VulkanHandle
	// ObjectType specifies the type of the object to be named
	ObjectType core1_0.ObjectType

	common.NextOptions
}

func (i DebugUtilsObjectNameInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkDebugUtilsObjectNameInfoEXT)
	}

	nameInfo := (*C.VkDebugUtilsObjectNameInfoEXT)(preallocatedPointer)
	nameInfo.sType = C.VK_STRUCTURE_TYPE_DEBUG_UTILS_OBJECT_NAME_INFO_EXT
	nameInfo.pNext = next
	nameInfo.objectType = C.VkObjectType(i.ObjectType)
	nameInfo.objectHandle = C.uint64_t(i.ObjectHandle)
	nameInfo.pObjectName = (*C.char)(allocator.CString(i.ObjectName))

	return preallocatedPointer, nil
}

func (i *DebugUtilsObjectNameInfo) PopulateFromCPointer(cDataPointer unsafe.Pointer) {
	objectNameInfo := (*C.VkDebugUtilsObjectNameInfoEXT)(cDataPointer)
	i.ObjectType = core1_0.ObjectType(objectNameInfo.objectType)
	i.ObjectHandle = driver.VulkanHandle(objectNameInfo.objectHandle)
	i.ObjectName = ""

	if objectNameInfo.pObjectName != nil {
		i.ObjectName = C.GoString(objectNameInfo.pObjectName)
	}
}
