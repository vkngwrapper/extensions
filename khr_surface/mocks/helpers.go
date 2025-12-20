package mock_surface

import (
	"math/rand"
	"unsafe"

	ext_driver "github.com/vkngwrapper/extensions/v3/khr_surface/driver"
	gomock "go.uber.org/mock/gomock"
)

func NewFakeSurface() ext_driver.VkSurfaceKHR {
	return ext_driver.VkSurfaceKHR(unsafe.Pointer(uintptr(rand.Int())))
}

func EasyMockSurface(ctrl *gomock.Controller) *MockSurface {
	surface := NewMockSurface(ctrl)
	surface.EXPECT().Handle().Return(NewFakeSurface()).AnyTimes()

	return surface
}
