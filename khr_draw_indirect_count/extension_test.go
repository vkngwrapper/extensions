package khr_draw_indirect_count

import (
	"testing"

	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/loader"
	"github.com/vkngwrapper/core/v3/mocks"
	mock_draw_indirect_count "github.com/vkngwrapper/extensions/v3/khr_draw_indirect_count/mocks"
	"go.uber.org/mock/gomock"
)

func TestVulkanExtension_CmdDrawIndexedIndirectCount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_draw_indirect_count.NewMockLoader(ctrl)
	extension := CreateExtensionDriverFromLoader(extDriver)

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	commandPool := mocks.NewDummyCommandPool(device)
	commandBuffer := mocks.NewDummyCommandBuffer(commandPool, device)
	buffer := mocks.NewDummyBuffer(device)
	countBuffer := mocks.NewDummyBuffer(device)

	extDriver.EXPECT().VkCmdDrawIndexedIndirectCountKHR(
		commandBuffer.Handle(),
		buffer.Handle(),
		loader.VkDeviceSize(1),
		countBuffer.Handle(),
		loader.VkDeviceSize(3),
		loader.Uint32(5),
		loader.Uint32(7),
	)

	extension.CmdDrawIndexedIndirectCount(
		commandBuffer,
		buffer,
		uint64(1),
		countBuffer,
		uint64(3),
		5,
		7,
	)
}

func TestVulkanExtension_CmdDrawIndirectCount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_draw_indirect_count.NewMockLoader(ctrl)
	extension := CreateExtensionDriverFromLoader(extDriver)

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	commandPool := mocks.NewDummyCommandPool(device)
	commandBuffer := mocks.NewDummyCommandBuffer(commandPool, device)
	buffer := mocks.NewDummyBuffer(device)
	countBuffer := mocks.NewDummyBuffer(device)

	extDriver.EXPECT().VkCmdDrawIndirectCountKHR(
		commandBuffer.Handle(),
		buffer.Handle(),
		loader.VkDeviceSize(11),
		countBuffer.Handle(),
		loader.VkDeviceSize(13),
		loader.Uint32(17),
		loader.Uint32(19),
	)

	extension.CmdDrawIndirectCount(
		commandBuffer,
		buffer,
		uint64(11),
		countBuffer,
		uint64(13),
		17,
		19,
	)
}
