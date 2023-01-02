package khr_draw_indirect_count_shim

import (
	"github.com/golang/mock/gomock"
	core_mocks "github.com/vkngwrapper/core/v2/mocks"
	mock_draw_indirect_count "github.com/vkngwrapper/extensions/v2/khr_draw_indirect_count/mocks"
	"testing"
)

func TestVulkanShim_CmdDrawIndexedIndirectCount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extension := mock_draw_indirect_count.NewMockExtension(ctrl)
	commandBuffer := core_mocks.NewMockCommandBuffer(ctrl)
	shim := NewShim(extension, commandBuffer)

	buffer := core_mocks.NewMockBuffer(ctrl)
	countBUffer := core_mocks.NewMockBuffer(ctrl)

	extension.EXPECT().CmdDrawIndexedIndirectCount(
		commandBuffer,
		buffer,
		uint64(1),
		countBUffer,
		uint64(3), 5, 7,
	)

	shim.CmdDrawIndexedIndirectCount(buffer, 1, countBUffer, 3, 5, 7)
}

func TestVulkanShim_CmdDrawIndirectCount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extension := mock_draw_indirect_count.NewMockExtension(ctrl)
	commandBuffer := core_mocks.NewMockCommandBuffer(ctrl)
	shim := NewShim(extension, commandBuffer)

	buffer := core_mocks.NewMockBuffer(ctrl)
	countBUffer := core_mocks.NewMockBuffer(ctrl)

	extension.EXPECT().CmdDrawIndirectCount(
		commandBuffer,
		buffer,
		uint64(1),
		countBUffer,
		uint64(3), 5, 7,
	)

	shim.CmdDrawIndirectCount(buffer, 1, countBUffer, 3, 5, 7)
}
