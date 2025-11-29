package khr_create_renderpass2_shim

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/core1_0"
	"github.com/vkngwrapper/core/v2/core1_1"
	"github.com/vkngwrapper/core/v2/core1_2"
	core_mocks "github.com/vkngwrapper/core/v2/mocks"
	"github.com/vkngwrapper/extensions/v2/khr_create_renderpass2"
	mock_create_renderpass2 "github.com/vkngwrapper/extensions/v2/khr_create_renderpass2/mocks"
	"go.uber.org/mock/gomock"
)

func TestVulkanCommandBufferShim_CmdBeginRenderPass2(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extension := mock_create_renderpass2.NewMockExtension(ctrl)
	commandBuffer := core_mocks.NewMockCommandBuffer(ctrl)
	shim := NewCommandBufferShim(extension, commandBuffer)

	renderPass := core_mocks.NewMockRenderPass(ctrl)
	framebuffer := core_mocks.NewMockFramebuffer(ctrl)

	extension.EXPECT().CmdBeginRenderPass2(
		commandBuffer,
		core1_0.RenderPassBeginInfo{
			RenderPass:  renderPass,
			Framebuffer: framebuffer,
			RenderArea: core1_0.Rect2D{
				Offset: core1_0.Offset2D{
					X: 3,
					Y: 7,
				},
				Extent: core1_0.Extent2D{
					Width:  11,
					Height: 13,
				},
			},
			ClearValues: []core1_0.ClearValue{
				core1_0.ClearValueDepthStencil{
					Depth:   5.0,
					Stencil: 17,
				},
			},
			NextOptions: common.NextOptions{
				Next: core1_1.DeviceGroupRenderPassBeginInfo{
					DeviceMask: 23,
				},
			},
		},
		khr_create_renderpass2.SubpassBeginInfo{
			Contents: core1_0.SubpassContentsInline,
		})

	err := shim.CmdBeginRenderPass2(
		core1_0.RenderPassBeginInfo{
			RenderPass:  renderPass,
			Framebuffer: framebuffer,
			RenderArea: core1_0.Rect2D{
				Offset: core1_0.Offset2D{
					X: 3,
					Y: 7,
				},
				Extent: core1_0.Extent2D{
					Width:  11,
					Height: 13,
				},
			},
			ClearValues: []core1_0.ClearValue{
				core1_0.ClearValueDepthStencil{
					Depth:   5.0,
					Stencil: 17,
				},
			},
			NextOptions: common.NextOptions{
				Next: core1_1.DeviceGroupRenderPassBeginInfo{
					DeviceMask: 23,
				},
			},
		},
		core1_2.SubpassBeginInfo{
			Contents: core1_0.SubpassContentsInline,
		})
	require.NoError(t, err)
}

func TestVulkanCommandBufferShim_CmdEndRenderPass2(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extension := mock_create_renderpass2.NewMockExtension(ctrl)
	commandBuffer := core_mocks.NewMockCommandBuffer(ctrl)
	shim := NewCommandBufferShim(extension, commandBuffer)

	extension.EXPECT().CmdEndRenderPass2(commandBuffer, khr_create_renderpass2.SubpassEndInfo{})

	err := shim.CmdEndRenderPass2(core1_2.SubpassEndInfo{})
	require.NoError(t, err)
}

func TestVulkanCommandBufferShim_CmdNextSubpass2(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extension := mock_create_renderpass2.NewMockExtension(ctrl)
	commandBuffer := core_mocks.NewMockCommandBuffer(ctrl)
	shim := NewCommandBufferShim(extension, commandBuffer)

	extension.EXPECT().CmdNextSubpass2(
		commandBuffer,
		khr_create_renderpass2.SubpassBeginInfo{
			Contents: core1_0.SubpassContentsInline,
		},
		khr_create_renderpass2.SubpassEndInfo{})

	err := shim.CmdNextSubpass2(
		core1_2.SubpassBeginInfo{
			Contents: core1_0.SubpassContentsInline,
		},
		core1_2.SubpassEndInfo{})
	require.NoError(t, err)
}
