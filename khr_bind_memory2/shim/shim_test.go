package khr_bind_memory2_shim

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_1"
	core_mocks "github.com/vkngwrapper/core/v3/mocks/mocks1_0"
	"github.com/vkngwrapper/extensions/v3/khr_bind_memory2"
	mock_bind_memory2 "github.com/vkngwrapper/extensions/v3/khr_bind_memory2/mocks"
	"go.uber.org/mock/gomock"
)

func TestVulkanShim_BindBufferMemory2(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extension := mock_bind_memory2.NewMockExtension(ctrl)
	device := core_mocks.NewMockDevice(ctrl)
	shim := NewShim(device, extension)

	buffer1 := core_mocks.NewMockBuffer(ctrl)
	buffer2 := core_mocks.NewMockBuffer(ctrl)
	memory := core_mocks.NewMockDeviceMemory(ctrl)

	extension.EXPECT().BindBufferMemory2(
		device,
		[]khr_bind_memory2.BindBufferMemoryInfo{
			{
				Buffer:       buffer1,
				Memory:       memory,
				MemoryOffset: 3,
			},
			{
				Buffer:       buffer2,
				Memory:       memory,
				MemoryOffset: 7,
				NextOptions: common.NextOptions{
					core1_1.MemoryDedicatedAllocateInfo{
						Buffer: buffer2,
					},
				},
			},
		})
	_, err := shim.BindBufferMemory2([]core1_1.BindBufferMemoryInfo{
		{
			Buffer:       buffer1,
			Memory:       memory,
			MemoryOffset: 3,
		},
		{
			Buffer:       buffer2,
			Memory:       memory,
			MemoryOffset: 7,
			NextOptions: common.NextOptions{
				core1_1.MemoryDedicatedAllocateInfo{
					Buffer: buffer2,
				},
			},
		},
	})
	require.NoError(t, err)
}

func TestVulkanShim_BindImageMemory2(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extension := mock_bind_memory2.NewMockExtension(ctrl)
	device := core_mocks.NewMockDevice(ctrl)
	shim := NewShim(device, extension)

	image1 := core_mocks.NewMockImage(ctrl)
	image2 := core_mocks.NewMockImage(ctrl)
	memory := core_mocks.NewMockDeviceMemory(ctrl)
	buffer := core_mocks.NewMockBuffer(ctrl)

	extension.EXPECT().BindImageMemory2(
		device,
		[]khr_bind_memory2.BindImageMemoryInfo{
			{
				Image:        image1,
				Memory:       memory,
				MemoryOffset: 3,
			},
			{
				Image:        image2,
				Memory:       memory,
				MemoryOffset: 7,
				NextOptions: common.NextOptions{
					core1_1.MemoryDedicatedAllocateInfo{
						Buffer: buffer,
					},
				},
			},
		})
	_, err := shim.BindImageMemory2([]core1_1.BindImageMemoryInfo{
		{
			Image:        image1,
			Memory:       memory,
			MemoryOffset: 3,
		},
		{
			Image:        image2,
			Memory:       memory,
			MemoryOffset: 7,
			NextOptions: common.NextOptions{
				core1_1.MemoryDedicatedAllocateInfo{
					Buffer: buffer,
				},
			},
		},
	})
	require.NoError(t, err)
}
