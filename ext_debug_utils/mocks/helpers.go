package mock_debugutils

import (
	"github.com/golang/mock/gomock"
	ext_driver "github.com/vkngwrapper/extensions/ext_debug_utils/driver"
	"math/rand"
	"unsafe"
)

func NewFakeMessenger() ext_driver.VkDebugUtilsMessengerEXT {
	return ext_driver.VkDebugUtilsMessengerEXT(unsafe.Pointer(uintptr(rand.Int())))
}

func EasyMockMessenger(ctrl *gomock.Controller) *MockDebugUtilsMessenger {
	messenger := NewMockDebugUtilsMessenger(ctrl)
	messenger.EXPECT().Handle().Return(NewFakeMessenger()).AnyTimes()

	return messenger
}
