package khr_descriptor_update_template_test

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v2/mocks"
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/core1_0"
	"github.com/vkngwrapper/core/v2/driver"
	mock_driver "github.com/vkngwrapper/core/v2/driver/mocks"
	"github.com/vkngwrapper/extensions/v2/khr_descriptor_update_template"
	khr_descriptor_update_template_driver "github.com/vkngwrapper/extensions/v2/khr_descriptor_update_template/driver"
	mock_descriptor_update_template "github.com/vkngwrapper/extensions/v2/khr_descriptor_update_template/mocks"
	"reflect"
	"testing"
	"unsafe"
)

func TestVulkanExtension_CreateDescriptorUpdateTemplate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_descriptor_update_template.NewMockDriver(ctrl)
	extension := khr_descriptor_update_template.CreateExtensionFromDriver(extDriver)

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := mocks.EasyMockDevice(ctrl, coreDriver)
	descriptorLayout := mocks.EasyMockDescriptorSetLayout(ctrl)
	pipelineLayout := mocks.EasyMockPipelineLayout(ctrl)

	handle := mock_descriptor_update_template.NewFakeDescriptorTemplate()

	extDriver.EXPECT().VkCreateDescriptorUpdateTemplateKHR(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(
		device driver.VkDevice,
		pCreateInfo *khr_descriptor_update_template_driver.VkDescriptorUpdateTemplateCreateInfoKHR,
		pAllocator *driver.VkAllocationCallbacks,
		pDescriptorTemplate *khr_descriptor_update_template_driver.VkDescriptorUpdateTemplateKHR,
	) (common.VkResult, error) {
		*pDescriptorTemplate = handle

		val := reflect.ValueOf(pCreateInfo).Elem()
		require.Equal(t, uint64(1000085000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DESCRIPTOR_UPDATE_TEMPLATE_CREATE_INFO_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(0), val.FieldByName("flags").Uint())
		require.Equal(t, uint64(2), val.FieldByName("descriptorUpdateEntryCount").Uint())
		require.Equal(t, uint64(0), val.FieldByName("templateType").Uint())
		require.Equal(t, descriptorLayout.Handle(), driver.VkDescriptorSetLayout(val.FieldByName("descriptorSetLayout").UnsafePointer()))
		require.Equal(t, pipelineLayout.Handle(), driver.VkPipelineLayout(val.FieldByName("pipelineLayout").UnsafePointer()))
		require.Equal(t, uint64(0), val.FieldByName("pipelineBindPoint").Uint())
		require.Equal(t, uint64(31), val.FieldByName("set").Uint())

		entriesPtr := (*khr_descriptor_update_template_driver.VkDescriptorUpdateTemplateEntryKHR)(val.FieldByName("pDescriptorUpdateEntries").UnsafePointer())
		entriesSlice := ([]khr_descriptor_update_template_driver.VkDescriptorUpdateTemplateEntryKHR)(unsafe.Slice(entriesPtr, 2))
		entries := reflect.ValueOf(entriesSlice)

		entry := entries.Index(0)
		require.Equal(t, uint64(1), entry.FieldByName("dstBinding").Uint())
		require.Equal(t, uint64(3), entry.FieldByName("dstArrayElement").Uint())
		require.Equal(t, uint64(5), entry.FieldByName("descriptorCount").Uint())
		require.Equal(t, uint64(1), entry.FieldByName("descriptorType").Uint()) // VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER
		require.Equal(t, uint64(7), entry.FieldByName("offset").Uint())
		require.Equal(t, uint64(11), entry.FieldByName("stride").Uint())

		entry = entries.Index(1)
		require.Equal(t, uint64(13), entry.FieldByName("dstBinding").Uint())
		require.Equal(t, uint64(17), entry.FieldByName("dstArrayElement").Uint())
		require.Equal(t, uint64(19), entry.FieldByName("descriptorCount").Uint())
		require.Equal(t, uint64(7), entry.FieldByName("descriptorType").Uint()) // VK_DESCRIPTOR_TYPE_STORAGE_BUFFER
		require.Equal(t, uint64(23), entry.FieldByName("offset").Uint())
		require.Equal(t, uint64(29), entry.FieldByName("stride").Uint())

		return core1_0.VKSuccess, nil
	})
	extDriver.EXPECT().VkDestroyDescriptorUpdateTemplateKHR(
		device.Handle(),
		handle,
		gomock.Nil(),
	)

	template, _, err := extension.CreateDescriptorUpdateTemplate(device, khr_descriptor_update_template.DescriptorUpdateTemplateCreateInfo{
		DescriptorUpdateEntries: []khr_descriptor_update_template.DescriptorUpdateTemplateEntry{
			{
				DstBinding:      1,
				DstArrayElement: 3,
				DescriptorCount: 5,
				DescriptorType:  core1_0.DescriptorTypeCombinedImageSampler,
				Offset:          7,
				Stride:          11,
			},
			{
				DstBinding:      13,
				DstArrayElement: 17,
				DescriptorCount: 19,
				DescriptorType:  core1_0.DescriptorTypeStorageBuffer,
				Offset:          23,
				Stride:          29,
			},
		},
		TemplateType:        khr_descriptor_update_template.DescriptorUpdateTemplateTypeDescriptorSet,
		DescriptorSetLayout: descriptorLayout,
		PipelineBindPoint:   core1_0.PipelineBindPointGraphics,
		PipelineLayout:      pipelineLayout,
		Set:                 31,
	}, nil)
	require.NoError(t, err)
	require.NotNil(t, template)
	require.Equal(t, handle, template.Handle())

	template.Destroy(nil)
}
