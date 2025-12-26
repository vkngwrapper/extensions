package khr_descriptor_update_template_test

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
	mock_driver "github.com/vkngwrapper/core/v3/loader/mocks"
	"github.com/vkngwrapper/core/v3/mocks/mocks1_0"
	"github.com/vkngwrapper/extensions/v3/khr_descriptor_update_template"
	khr_descriptor_update_template_driver "github.com/vkngwrapper/extensions/v3/khr_descriptor_update_template/loader"
	mock_descriptor_update_template "github.com/vkngwrapper/extensions/v3/khr_descriptor_update_template/mocks"
	"go.uber.org/mock/gomock"
)

func TestVulkanDescriptorTemplate_UpdateDescriptorSetFromBuffer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_descriptor_update_template.NewMockDriver(ctrl)
	extension := khr_descriptor_update_template.CreateExtensionFromDriver(extDriver)

	coreLoader := mock_driver.LoaderForVersion(ctrl, common.Vulkan1_0)
	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	descriptorSet := mocks.NewDummyDescriptorSet(descriptorPool, device)
	buffer := mocks.NewDummyBuffer(device)

	handle := mock_descriptor_update_template.NewFakeDescriptorTemplate()

	extDriver.EXPECT().VkCreateDescriptorUpdateTemplateKHR(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(
		device loader.VkDevice,
		pCreateInfo *khr_descriptor_update_template_driver.VkDescriptorUpdateTemplateCreateInfoKHR,
		pAllocator *loader.VkAllocationCallbacks,
		pDescriptorTemplate *khr_descriptor_update_template_driver.VkDescriptorUpdateTemplateKHR,
	) (common.VkResult, error) {
		*pDescriptorTemplate = handle

		return core1_0.VKSuccess, nil
	})

	extDriver.EXPECT().VkUpdateDescriptorSetWithTemplateKHR(
		device.Handle(),
		descriptorSet.Handle(),
		handle,
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(
		device loader.VkDevice,
		descriptorSet loader.VkDescriptorSet,
		template khr_descriptor_update_template_driver.VkDescriptorUpdateTemplateKHR,
		pData unsafe.Pointer,
	) {
		infoPtr := (*loader.VkDescriptorBufferInfo)(pData)
		info := reflect.ValueOf(infoPtr).Elem()
		require.Equal(t, buffer.Handle(), (loader.VkBuffer)(info.FieldByName("buffer").UnsafePointer()))
		require.Equal(t, uint64(1), info.FieldByName("offset").Uint())
		require.Equal(t, uint64(3), info.FieldByName("_range").Uint())
	})

	template, _, err := extension.CreateDescriptorUpdateTemplate(device, khr_descriptor_update_template.DescriptorUpdateTemplateCreateInfo{}, nil)
	require.NoError(t, err)
	require.NotNil(t, template)

	template.UpdateDescriptorSetFromBuffer(descriptorSet, core1_0.DescriptorBufferInfo{
		Buffer: buffer,
		Offset: 1,
		Range:  3,
	})
}

func TestVulkanDescriptorTemplate_UpdateDescriptorSetFromImage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_descriptor_update_template.NewMockDriver(ctrl)
	extension := khr_descriptor_update_template.CreateExtensionFromDriver(extDriver)

	coreLoader := mock_driver.LoaderForVersion(ctrl, common.Vulkan1_0)
	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	descriptorSet := mocks.NewDummyDescriptorSet(descriptorPool, device)
	sampler := mocks.NewDummySampler(device)
	imageView := mocks.NewDummyImageView(device)

	handle := mock_descriptor_update_template.NewFakeDescriptorTemplate()

	extDriver.EXPECT().VkCreateDescriptorUpdateTemplateKHR(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(
		device loader.VkDevice,
		pCreateInfo *khr_descriptor_update_template_driver.VkDescriptorUpdateTemplateCreateInfoKHR,
		pAllocator *loader.VkAllocationCallbacks,
		pDescriptorTemplate *khr_descriptor_update_template_driver.VkDescriptorUpdateTemplateKHR,
	) (common.VkResult, error) {
		*pDescriptorTemplate = handle

		return core1_0.VKSuccess, nil
	})

	extDriver.EXPECT().VkUpdateDescriptorSetWithTemplateKHR(
		device.Handle(),
		descriptorSet.Handle(),
		handle,
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(
		device loader.VkDevice,
		descriptorSet loader.VkDescriptorSet,
		template khr_descriptor_update_template_driver.VkDescriptorUpdateTemplateKHR,
		pData unsafe.Pointer,
	) {
		infoPtr := (*loader.VkDescriptorImageInfo)(pData)
		info := reflect.ValueOf(infoPtr).Elem()
		require.Equal(t, sampler.Handle(), (loader.VkSampler)(info.FieldByName("sampler").UnsafePointer()))
		require.Equal(t, imageView.Handle(), (loader.VkImageView)(info.FieldByName("imageView").UnsafePointer()))
		require.Equal(t, uint64(7), info.FieldByName("imageLayout").Uint()) // VK_IMAGE_LAYOUT_TRANSFER_DST_OPTIMAL
	})

	template, _, err := extension.CreateDescriptorUpdateTemplate(device, khr_descriptor_update_template.DescriptorUpdateTemplateCreateInfo{}, nil)
	require.NoError(t, err)
	require.NotNil(t, template)

	template.UpdateDescriptorSetFromImage(descriptorSet, core1_0.DescriptorImageInfo{
		Sampler:     sampler,
		ImageView:   imageView,
		ImageLayout: core1_0.ImageLayoutTransferDstOptimal,
	})
}

func TestVulkanDescriptorTemplate_UpdateDescriptorSetFromObjectHandle(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_descriptor_update_template.NewMockDriver(ctrl)
	extension := khr_descriptor_update_template.CreateExtensionFromDriver(extDriver)

	coreLoader := mock_driver.LoaderForVersion(ctrl, common.Vulkan1_0)
	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	descriptorSet := mocks.NewDummyDescriptorSet(descriptorPool, device)
	bufferView := mocks1_0.EasyMockBufferView(ctrl)

	handle := mock_descriptor_update_template.NewFakeDescriptorTemplate()

	extDriver.EXPECT().VkCreateDescriptorUpdateTemplateKHR(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(
		device loader.VkDevice,
		pCreateInfo *khr_descriptor_update_template_driver.VkDescriptorUpdateTemplateCreateInfoKHR,
		pAllocator *loader.VkAllocationCallbacks,
		pDescriptorTemplate *khr_descriptor_update_template_driver.VkDescriptorUpdateTemplateKHR,
	) (common.VkResult, error) {
		*pDescriptorTemplate = handle

		return core1_0.VKSuccess, nil
	})

	extDriver.EXPECT().VkUpdateDescriptorSetWithTemplateKHR(
		device.Handle(),
		descriptorSet.Handle(),
		handle,
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(
		device loader.VkDevice,
		descriptorSet loader.VkDescriptorSet,
		template khr_descriptor_update_template_driver.VkDescriptorUpdateTemplateKHR,
		pData unsafe.Pointer,
	) {
		info := (loader.VkBufferView)(pData)
		require.Equal(t, bufferView.Handle(), info)
	})

	template, _, err := extension.CreateDescriptorUpdateTemplate(device, khr_descriptor_update_template.DescriptorUpdateTemplateCreateInfo{}, nil)
	require.NoError(t, err)
	require.NotNil(t, template)

	template.UpdateDescriptorSetFromObjectHandle(descriptorSet, loader.VulkanHandle(bufferView.Handle()))
}
