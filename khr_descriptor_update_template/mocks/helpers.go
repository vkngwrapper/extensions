package mock_descriptor_update_template

import (
	"github.com/golang/mock/gomock"
	khr_descriptor_update_template_driver "github.com/vkngwrapper/extensions/khr_descriptor_update_template/driver"
	"math/rand"
	"unsafe"
)

func NewFakeDescriptorTemplate() khr_descriptor_update_template_driver.VkDescriptorUpdateTemplateKHR {
	return khr_descriptor_update_template_driver.VkDescriptorUpdateTemplateKHR(unsafe.Pointer(uintptr(rand.Int())))
}

func EasyMockDescriptorTemplate(ctrl *gomock.Controller) *MockDescriptorUpdateTemplate {
	swapchain := NewMockDescriptorUpdateTemplate(ctrl)
	swapchain.EXPECT().Handle().Return(NewFakeDescriptorTemplate()).AnyTimes()

	return swapchain
}
