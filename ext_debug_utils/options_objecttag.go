package ext_debug_utils

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/core1_0"
	"github.com/vkngwrapper/core/driver"
	"unsafe"
)

// DebugUtilsObjectTagInfo specifies parameters of a tag to attach to an object
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsObjectTagInfoEXT.html
type DebugUtilsObjectTagInfo struct {
	// ObjectType specifies the type of the object to be named
	ObjectType core1_0.ObjectType
	// ObjectHandle is the object to be tagged
	ObjectHandle driver.VulkanHandle

	// TagName is a numerical identifier of the tag
	TagName uint64
	// Tag is a slice of bytes containing the data to be associated with the object
	Tag []byte

	common.NextOptions
}

func (t DebugUtilsObjectTagInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkDebugUtilsObjectTagInfoEXT)
	}

	tagInfo := (*C.VkDebugUtilsObjectTagInfoEXT)(preallocatedPointer)
	tagInfo.sType = C.VK_STRUCTURE_TYPE_DEBUG_UTILS_OBJECT_TAG_INFO_EXT
	tagInfo.pNext = next
	tagInfo.objectType = C.VkObjectType(t.ObjectType)
	tagInfo.objectHandle = C.uint64_t(t.ObjectHandle)
	tagInfo.tagName = C.uint64_t(t.TagName)
	tagInfo.tagSize = C.size_t(len(t.Tag))
	tagInfo.pTag = allocator.CBytes(t.Tag)

	return preallocatedPointer, nil
}

func (t *DebugUtilsObjectTagInfo) PopulateFromCPointer(cPointer unsafe.Pointer) {
	tagInfo := (*C.VkDebugUtilsObjectTagInfoEXT)(cPointer)

	t.ObjectType = core1_0.ObjectType(tagInfo.objectType)
	t.ObjectHandle = driver.VulkanHandle(tagInfo.objectHandle)
	t.TagName = uint64(tagInfo.tagName)
	t.Tag = C.GoBytes(tagInfo.pTag, C.int(tagInfo.tagSize))
}
