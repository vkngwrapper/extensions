package mock_debugutils

import (
	"math/rand"
	"unsafe"

	ext_driver "github.com/vkngwrapper/extensions/v2/ext_debug_utils/driver"
	gomock "go.uber.org/mock/gomock"
)

func NewFakeMessenger() ext_driver.VkDebugUtilsMessengerEXT {
	return ext_driver.VkDebugUtilsMessengerEXT(unsafe.Pointer(uintptr(rand.Int())))
}

func EasyMockMessenger(ctrl *gomock.Controller) *MockDebugUtilsMessenger {
	messenger := NewMockDebugUtilsMessenger(ctrl)
	messenger.EXPECT().Handle().Return(NewFakeMessenger()).AnyTimes()

	return messenger
}
