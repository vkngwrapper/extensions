package mock_debugutils

import (
	"math/rand"
	"unsafe"

	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/extensions/v3/ext_debug_utils"
	ext_driver "github.com/vkngwrapper/extensions/v3/ext_debug_utils/loader"
)

func NewFakeMessenger() ext_driver.VkDebugUtilsMessengerEXT {
	return ext_driver.VkDebugUtilsMessengerEXT(unsafe.Pointer(uintptr(rand.Int())))
}

func NewDummyMessenger(instance core1_0.Instance) ext_debug_utils.DebugUtilsMessenger {
	return ext_debug_utils.InternalDebugUtilsMessenger(instance.Handle(), NewFakeMessenger(), instance.APIVersion())
}
