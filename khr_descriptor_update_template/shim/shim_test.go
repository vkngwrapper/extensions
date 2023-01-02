package khr_descriptor_update_template_shim

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v2/core1_0"
	"github.com/vkngwrapper/core/v2/core1_1"
	"github.com/vkngwrapper/core/v2/driver"
	core_mocks "github.com/vkngwrapper/core/v2/mocks"
	"github.com/vkngwrapper/extensions/v2/khr_descriptor_update_template"
	mock_descriptor_update_template "github.com/vkngwrapper/extensions/v2/khr_descriptor_update_template/mocks"
	"testing"
)

func TestVulkanShim_CreateDescriptorUpdateTemplate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extension := mock_descriptor_update_template.NewMockExtension(ctrl)
	device := core_mocks.NewMockDevice(ctrl)
	shim := NewShim(extension, device)

	layout := core_mocks.NewMockDescriptorSetLayout(ctrl)
	pipelineLayout := core_mocks.NewMockPipelineLayout(ctrl)
	descriptorUpdateTemplate := mock_descriptor_update_template.EasyMockDescriptorTemplate(ctrl)

	extension.EXPECT().CreateDescriptorUpdateTemplate(
		device,
		khr_descriptor_update_template.DescriptorUpdateTemplateCreateInfo{
			DescriptorUpdateEntries: []khr_descriptor_update_template.DescriptorUpdateTemplateEntry{
				{
					DstBinding:      1,
					DstArrayElement: 3,
					DescriptorCount: 5,
					DescriptorType:  core1_0.DescriptorTypeSampledImage,
					Offset:          7,
					Stride:          11,
				},
				{
					DstBinding:      2,
					DstArrayElement: 13,
					DescriptorCount: 17,
					DescriptorType:  core1_0.DescriptorTypeStorageBuffer,
					Offset:          19,
					Stride:          23,
				},
			},
			TemplateType:        khr_descriptor_update_template.DescriptorUpdateTemplateTypeDescriptorSet,
			DescriptorSetLayout: layout,
			PipelineBindPoint:   core1_0.PipelineBindPointGraphics,
			PipelineLayout:      pipelineLayout,
			Set:                 29,
		}, nil).Return(descriptorUpdateTemplate, core1_0.VKSuccess, nil)

	template, _, err := shim.CreateDescriptorUpdateTemplate(
		core1_1.DescriptorUpdateTemplateCreateInfo{
			DescriptorUpdateEntries: []core1_1.DescriptorUpdateTemplateEntry{
				{
					DstBinding:      1,
					DstArrayElement: 3,
					DescriptorCount: 5,
					DescriptorType:  core1_0.DescriptorTypeSampledImage,
					Offset:          7,
					Stride:          11,
				},
				{
					DstBinding:      2,
					DstArrayElement: 13,
					DescriptorCount: 17,
					DescriptorType:  core1_0.DescriptorTypeStorageBuffer,
					Offset:          19,
					Stride:          23,
				},
			},
			TemplateType:        core1_1.DescriptorUpdateTemplateTypeDescriptorSet,
			DescriptorSetLayout: layout,
			PipelineBindPoint:   core1_0.PipelineBindPointGraphics,
			PipelineLayout:      pipelineLayout,
			Set:                 29,
		}, nil)
	require.NoError(t, err)

	require.Equal(t, driver.VkDescriptorUpdateTemplate(descriptorUpdateTemplate.Handle()), template.Handle())
}
