package ext_debug_utils

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
#include "debug_callback.h"

VKAPI_ATTR VkBool32 VKAPI_CALL debugCallback(
	VkDebugUtilsMessageSeverityFlagBitsEXT messageSeverity,
	VkDebugUtilsMessageTypeFlagsEXT messageType,
	const VkDebugUtilsMessengerCallbackDataEXT *pCallbackData,
	void *pUserData) {

	return goDebugCallback(messageSeverity, messageType, (VkDebugUtilsMessengerCallbackDataEXT*)pCallbackData, pUserData);
}
*/
import "C"
import (
	"runtime/cgo"
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3/common"
)

// DebugUtilsMessengerCreateInfo specifies parameters of a newly-created DebugUtilsMessenger
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsMessengerCreateInfoEXT.html
type DebugUtilsMessengerCreateInfo struct {
	// Flags is reserved for future use
	Flags DebugUtilsMessengerCreateFlags
	// MessageSeverity specifies which severity of event(s) will cause this callback
	// to be called
	MessageSeverity DebugUtilsMessageSeverityFlags
	// MessageType specifies which type of event(s) will cause this callback to be called
	MessageType DebugUtilsMessageTypeFlags
	// UserCallback is the application callback function to call
	UserCallback CallbackFunction

	common.NextOptions
}

func (o DebugUtilsMessengerCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof([1]C.VkDebugUtilsMessengerCreateInfoEXT{})))
	}
	createInfo := (*C.VkDebugUtilsMessengerCreateInfoEXT)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_DEBUG_UTILS_MESSENGER_CREATE_INFO_EXT
	createInfo.flags = C.VkDebugUtilsMessengerCreateFlagsEXT(o.Flags)
	createInfo.pNext = next

	createInfo.messageSeverity = C.VkDebugUtilsMessageSeverityFlagsEXT(o.MessageSeverity)
	createInfo.messageType = C.VkDebugUtilsMessageTypeFlagsEXT(o.MessageType)
	createInfo.pfnUserCallback = (C.PFN_vkDebugUtilsMessengerCallbackEXT)(C.debugCallback)
	createInfo.pUserData = unsafe.Pointer(cgo.NewHandle(o.UserCallback))

	return preallocatedPointer, nil
}
