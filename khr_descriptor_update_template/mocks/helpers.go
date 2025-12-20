package mock_descriptor_update_template

import (
	"math/rand"
	"unsafe"

	khr_descriptor_update_template_driver "github.com/vkngwrapper/extensions/v3/khr_descriptor_update_template/driver"
	gomock "go.uber.org/mock/gomock"
)

func NewFakeDescriptorTemplate() khr_descriptor_update_template_driver.VkDescriptorUpdateTemplateKHR {
	return khr_descriptor_update_template_driver.VkDescriptorUpdateTemplateKHR(unsafe.Pointer(uintptr(rand.Int())))
}

func EasyMockDescriptorTemplate(ctrl *gomock.Controller) *MockDescriptorUpdateTemplate {
	swapchain := NewMockDescriptorUpdateTemplate(ctrl)
	swapchain.EXPECT().Handle().Return(NewFakeDescriptorTemplate()).AnyTimes()

	return swapchain
}
