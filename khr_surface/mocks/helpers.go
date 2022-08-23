package mock_surface

import (
	"github.com/golang/mock/gomock"
	ext_driver "github.com/vkngwrapper/extensions/v2/khr_surface/driver"
	"math/rand"
	"unsafe"
)

func NewFakeSurface() ext_driver.VkSurfaceKHR {
	return ext_driver.VkSurfaceKHR(unsafe.Pointer(uintptr(rand.Int())))
}

func EasyMockSurface(ctrl *gomock.Controller) *MockSurface {
	surface := NewMockSurface(ctrl)
	surface.EXPECT().Handle().Return(NewFakeSurface()).AnyTimes()

	return surface
}
