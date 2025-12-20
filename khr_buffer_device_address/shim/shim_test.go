package khr_buffer_device_address_shim

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/core1_1"
	"github.com/vkngwrapper/core/v2/core1_2"
	core_mocks "github.com/vkngwrapper/core/v2/mocks"
	"github.com/vkngwrapper/extensions/v3/khr_buffer_device_address"
	mock_buffer_device_address "github.com/vkngwrapper/extensions/v3/khr_buffer_device_address/mocks"
	"go.uber.org/mock/gomock"
)

func TestVulkanShim_GetBufferDeviceAddress(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extension := mock_buffer_device_address.NewMockExtension(ctrl)
	device := core_mocks.NewMockDevice(ctrl)
	shim := NewShim(extension, device)

	buffer := core_mocks.NewMockBuffer(ctrl)
	extension.EXPECT().GetBufferDeviceAddress(
		device,
		khr_buffer_device_address.BufferDeviceAddressInfo{
			Buffer: buffer,
			NextOptions: common.NextOptions{
				core1_1.MemoryDedicatedAllocateInfo{
					Buffer: buffer,
				},
			},
		}).Return(uint64(101), nil)

	address, err := shim.GetBufferDeviceAddress(core1_2.BufferDeviceAddressInfo{
		Buffer: buffer,
		NextOptions: common.NextOptions{
			core1_1.MemoryDedicatedAllocateInfo{
				Buffer: buffer,
			},
		},
	})
	require.Equal(t, uint64(101), address)
	require.NoError(t, err)
}

func TestVulkanShim_GetBufferOpaqueCaptureAddress(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extension := mock_buffer_device_address.NewMockExtension(ctrl)
	device := core_mocks.NewMockDevice(ctrl)
	shim := NewShim(extension, device)

	buffer := core_mocks.NewMockBuffer(ctrl)
	extension.EXPECT().GetBufferOpaqueCaptureAddress(
		device,
		khr_buffer_device_address.BufferDeviceAddressInfo{
			Buffer: buffer,
			NextOptions: common.NextOptions{
				core1_1.MemoryDedicatedAllocateInfo{
					Buffer: buffer,
				},
			},
		}).Return(uint64(101), nil)

	address, err := shim.GetBufferOpaqueCaptureAddress(core1_2.BufferDeviceAddressInfo{
		Buffer: buffer,
		NextOptions: common.NextOptions{
			core1_1.MemoryDedicatedAllocateInfo{
				Buffer: buffer,
			},
		},
	})
	require.Equal(t, uint64(101), address)
	require.NoError(t, err)
}

func TestVulkanShim_GetDeviceMemoryOpaqueCaptureAddress(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extension := mock_buffer_device_address.NewMockExtension(ctrl)
	device := core_mocks.NewMockDevice(ctrl)
	shim := NewShim(extension, device)

	buffer := core_mocks.NewMockBuffer(ctrl)
	memory := core_mocks.NewMockDeviceMemory(ctrl)

	extension.EXPECT().GetDeviceMemoryOpaqueCaptureAddress(
		device,
		khr_buffer_device_address.DeviceMemoryOpaqueCaptureAddressInfo{
			Memory: memory,
			NextOptions: common.NextOptions{
				core1_1.MemoryDedicatedAllocateInfo{
					Buffer: buffer,
				},
			},
		}).Return(uint64(101), nil)

	address, err := shim.GetDeviceMemoryOpaqueCaptureAddress(core1_2.DeviceMemoryOpaqueCaptureAddressInfo{
		Memory: memory,
		NextOptions: common.NextOptions{
			core1_1.MemoryDedicatedAllocateInfo{
				Buffer: buffer,
			},
		},
	})
	require.Equal(t, uint64(101), address)
	require.NoError(t, err)
}
