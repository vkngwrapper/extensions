package mock_sampler_ycbcr_conversion

import (
	"math/rand"
	"unsafe"

	khr_sampler_ycbcr_conversion_driver "github.com/vkngwrapper/extensions/v3/khr_sampler_ycbcr_conversion/driver"
	gomock "go.uber.org/mock/gomock"
)

func NewFakeSamplerYcbcrConversion() khr_sampler_ycbcr_conversion_driver.VkSamplerYcbcrConversionKHR {
	return khr_sampler_ycbcr_conversion_driver.VkSamplerYcbcrConversionKHR(unsafe.Pointer(uintptr(rand.Int())))
}

func EasyMockSamplerYcbcrConversion(ctrl *gomock.Controller) *MockSamplerYcbcrConversion {
	sampler := NewMockSamplerYcbcrConversion(ctrl)
	sampler.EXPECT().Handle().Return(NewFakeSamplerYcbcrConversion()).AnyTimes()

	return sampler
}
