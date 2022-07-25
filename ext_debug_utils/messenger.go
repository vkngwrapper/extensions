package ext_debug_utils

//go:generate mockgen -source messenger.go -destination ./mocks/messenger.go -package mock_debugutils

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"

*/
import "C"
import (
	"fmt"
	"github.com/vkngwrapper/core/driver"
	ext_driver "github.com/vkngwrapper/extensions/ext_debug_utils/driver"
	"runtime/cgo"
	"unsafe"
)

// DebugUtilsMessenger is a messenger object which handles passing along debug
// messages to a provided debug callback
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsMessengerEXT.html
type DebugUtilsMessenger interface {
	// Handle is the internal Vulkan object handle for this DebugUtilsMessenger
	Handle() ext_driver.VkDebugUtilsMessengerEXT

	// Destroy destroys the DebugUtilsMessenger object and the underlying structures. **Warning** after
	// destruction, this object will continue to exist, but the Vulkan object handle that backs it will
	// be invalid. Do not call further methods on this object.
	Destroy(callbacks *driver.AllocationCallbacks)
}

// VulkanDebugUtilsMessenger is an implementation of the DebugUtilsMessenger interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanDebugUtilsMessenger struct {
	instance   driver.VkInstance
	handle     ext_driver.VkDebugUtilsMessengerEXT
	coreDriver driver.Driver
	driver     ext_driver.Driver
}

func (m *VulkanDebugUtilsMessenger) Destroy(callbacks *driver.AllocationCallbacks) {
	m.driver.VkDestroyDebugUtilsMessengerEXT(m.instance, m.handle, callbacks.Handle())
	m.coreDriver.ObjectStore().Delete(driver.VulkanHandle(m.handle))
}

func (m *VulkanDebugUtilsMessenger) Handle() ext_driver.VkDebugUtilsMessengerEXT {
	return m.handle
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
