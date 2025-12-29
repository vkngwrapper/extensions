package ext_debug_utils

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"

*/
import "C"
import (
	"fmt"
	"runtime/cgo"
	"unsafe"

	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/loader"
	ext_driver "github.com/vkngwrapper/extensions/v3/ext_debug_utils/loader"
)

// DebugUtilsMessenger is an implementation of the DebugUtilsMessenger interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type DebugUtilsMessenger struct {
	instance   loader.VkInstance
	handle     ext_driver.VkDebugUtilsMessengerEXT
	apiVersion common.APIVersion
}

func (m DebugUtilsMessenger) InstanceHandle() loader.VkInstance {
	return m.instance
}

func (m DebugUtilsMessenger) Handle() ext_driver.VkDebugUtilsMessengerEXT {
	return m.handle
}

func (m DebugUtilsMessenger) APIVersion() common.APIVersion {
	return m.apiVersion
}

func (m DebugUtilsMessenger) Initialized() bool {
	return m.handle != 0
}

func InternalDebugUtilsMessenger(instance loader.VkInstance, handle ext_driver.VkDebugUtilsMessengerEXT, apiVersion common.APIVersion) DebugUtilsMessenger {
	return DebugUtilsMessenger{
		instance:   instance,
		handle:     handle,
		apiVersion: apiVersion,
	}
}

// CallbackFunction is the application callback function type
type CallbackFunction func(msgType DebugUtilsMessageTypeFlags, severity DebugUtilsMessageSeverityFlags, data *DebugUtilsMessengerCallbackData) bool

//export goDebugCallback
func goDebugCallback(messageSeverity C.VkDebugUtilsMessageSeverityFlagBitsEXT, messageType C.VkDebugUtilsMessageTypeFlagsEXT, data *C.VkDebugUtilsMessengerCallbackDataEXT, userData unsafe.Pointer) C.VkBool32 {
	severity := DebugUtilsMessageSeverityFlags(messageSeverity)
	msgType := DebugUtilsMessageTypeFlags(messageType)

	callbackData := &DebugUtilsMessengerCallbackData{}

	err := callbackData.PopulateFromCPointer(unsafe.Pointer(data))
	if err != nil {
		callbackData = &DebugUtilsMessengerCallbackData{
			MessageIDName: "vkng-internal",
			Message:       fmt.Sprintf("error loading debug callback data from C: %v+", err),
		}
	}

	f := cgo.Handle(userData).Value().(CallbackFunction)
	if f(msgType, severity, callbackData) {
		return C.VK_TRUE
	}

	return C.VK_FALSE
}
