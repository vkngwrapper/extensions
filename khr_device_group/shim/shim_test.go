package khr_device_group_shim

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3/core1_1"
	core_mocks "github.com/vkngwrapper/core/v3/mocks/mocks1_0"
	"github.com/vkngwrapper/extensions/v3/khr_device_group"
	mock_device_group "github.com/vkngwrapper/extensions/v3/khr_device_group/mocks"
	"go.uber.org/mock/gomock"
)

func TestVulkanCommandBufferShim_CmdDispatchBase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extension := mock_device_group.NewMockExtension(ctrl)
	commandBuffer := core_mocks.NewMockCommandBuffer(ctrl)
	shim := NewCommandBufferShim(extension, commandBuffer)

	extension.EXPECT().CmdDispatchBase(
		commandBuffer,
		1, 3, 5, 7, 11, 13,
	)

	shim.CmdDispatchBase(1, 3, 5, 7, 11, 13)
}

func TestVulkanCommandBufferShim_CmdSetDeviceMask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extension := mock_device_group.NewMockExtension(ctrl)
	commandBuffer := core_mocks.NewMockCommandBuffer(ctrl)
	shim := NewCommandBufferShim(extension, commandBuffer)

	extension.EXPECT().CmdSetDeviceMask(commandBuffer, uint32(7))

	shim.CmdSetDeviceMask(7)
}

func TestVulkanDeviceShim_DeviceGroupPeerMemoryFeatures(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extension := mock_device_group.NewMockExtension(ctrl)
	device := core_mocks.NewMockDevice(ctrl)
	shim := NewDeviceShim(extension, device)

	extension.EXPECT().DeviceGroupPeerMemoryFeatures(
		device,
		1, 3, 7,
	).Return(khr_device_group.PeerMemoryFeatureCopyDst)

	flags := shim.DeviceGroupPeerMemoryFeatures(
		1, 3, 7,
	)
	require.Equal(t, core1_1.PeerMemoryFeatureCopyDst, flags)
}
