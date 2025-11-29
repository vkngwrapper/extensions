package khr_get_memory_requirements2_shim

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/core1_0"
	"github.com/vkngwrapper/core/v2/core1_1"
	core_mocks "github.com/vkngwrapper/core/v2/mocks"
	"github.com/vkngwrapper/extensions/v2/khr_get_memory_requirements2"
	mock_get_memory_requirements2 "github.com/vkngwrapper/extensions/v2/khr_get_memory_requirements2/mocks"
	"go.uber.org/mock/gomock"
)

func TestVulkanShim_BufferMemoryRequirements2(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extension := mock_get_memory_requirements2.NewMockExtension(ctrl)
	device := core_mocks.NewMockDevice(ctrl)
	shim := NewShim(extension, device)

	buffer := core_mocks.NewMockBuffer(ctrl)

	extension.EXPECT().BufferMemoryRequirements2(
		device,
		khr_get_memory_requirements2.BufferMemoryRequirementsInfo2{
			Buffer: buffer,
		},
		gomock.Any(),
	).DoAndReturn(func(device core1_0.Device,
		o khr_get_memory_requirements2.BufferMemoryRequirementsInfo2,
		out *khr_get_memory_requirements2.MemoryRequirements2) error {

		out.MemoryRequirements.MemoryTypeBits = 11
		out.MemoryRequirements.Size = 3
		out.MemoryRequirements.Alignment = 5

		next := out.Next.(*core1_1.MemoryDedicatedRequirements)
		next.PrefersDedicatedAllocation = true

		return nil
	})

	var outData core1_1.MemoryRequirements2
	outData.Next = &core1_1.MemoryDedicatedRequirements{}
	err := shim.BufferMemoryRequirements2(
		core1_1.BufferMemoryRequirementsInfo2{
			Buffer: buffer,
		},
		&outData)
	require.NoError(t, err)
	require.Equal(t,
		core1_1.MemoryRequirements2{
			MemoryRequirements: core1_0.MemoryRequirements{
				Size:           3,
				Alignment:      5,
				MemoryTypeBits: 11,
			},
			NextOutData: common.NextOutData{
				&core1_1.MemoryDedicatedRequirements{
					PrefersDedicatedAllocation: true,
				},
			},
		}, outData)
}

func TestVulkanShim_ImageMemoryRequirements2(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extension := mock_get_memory_requirements2.NewMockExtension(ctrl)
	device := core_mocks.NewMockDevice(ctrl)
	shim := NewShim(extension, device)

	image := core_mocks.NewMockImage(ctrl)

	extension.EXPECT().ImageMemoryRequirements2(
		device,
		khr_get_memory_requirements2.ImageMemoryRequirementsInfo2{
			Image: image,
		}, gomock.Any()).DoAndReturn(
		func(device core1_0.Device, o khr_get_memory_requirements2.ImageMemoryRequirementsInfo2, out *khr_get_memory_requirements2.MemoryRequirements2) error {
			out.MemoryRequirements.Size = 5
			out.MemoryRequirements.Alignment = 7
			out.MemoryRequirements.MemoryTypeBits = 13

			next := out.Next.(*core1_1.MemoryDedicatedRequirements)
			next.RequiresDedicatedAllocation = true

			return nil
		})

	var outData core1_1.MemoryRequirements2
	outData.Next = &core1_1.MemoryDedicatedRequirements{}
	err := shim.ImageMemoryRequirements2(
		core1_1.ImageMemoryRequirementsInfo2{
			Image: image,
		},
		&outData)
	require.NoError(t, err)
	require.Equal(t, core1_1.MemoryRequirements2{
		MemoryRequirements: core1_0.MemoryRequirements{
			Size:           5,
			Alignment:      7,
			MemoryTypeBits: 13,
		},
		NextOutData: common.NextOutData{
			&core1_1.MemoryDedicatedRequirements{
				RequiresDedicatedAllocation: true,
			},
		},
	}, outData)
}

func TestVulkanShim_ImageSparseMemoryRequirements2(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extension := mock_get_memory_requirements2.NewMockExtension(ctrl)
	device := core_mocks.NewMockDevice(ctrl)
	shim := NewShim(extension, device)

	image := core_mocks.NewMockImage(ctrl)

	extension.EXPECT().ImageSparseMemoryRequirements2(
		device,
		khr_get_memory_requirements2.ImageSparseMemoryRequirementsInfo2{
			Image: image,
		},
		gomock.Any()).DoAndReturn(func(
		device core1_0.Device,
		o khr_get_memory_requirements2.ImageSparseMemoryRequirementsInfo2,
		outDataFactory func() *khr_get_memory_requirements2.SparseImageMemoryRequirements2) ([]*khr_get_memory_requirements2.SparseImageMemoryRequirements2, error) {

		retVal := make([]*khr_get_memory_requirements2.SparseImageMemoryRequirements2, 2)
		retVal[0] = outDataFactory()
		retVal[1] = outDataFactory()

		retVal[0].MemoryRequirements = core1_0.SparseImageMemoryRequirements{
			FormatProperties: core1_0.SparseImageFormatProperties{
				AspectMask: core1_0.ImageAspectDepth,
				ImageGranularity: core1_0.Extent3D{
					Width:  1,
					Height: 3,
					Depth:  5,
				},
				Flags: core1_0.SparseImageFormatSingleMipTail,
			},
			ImageMipTailFirstLod: 7,
			ImageMipTailOffset:   11,
			ImageMipTailSize:     13,
			ImageMipTailStride:   17,
		}
		next := retVal[0].Next.(*core1_1.MemoryDedicatedRequirements)
		next.RequiresDedicatedAllocation = true

		retVal[1].MemoryRequirements = core1_0.SparseImageMemoryRequirements{
			FormatProperties: core1_0.SparseImageFormatProperties{
				AspectMask: core1_0.ImageAspectMetadata,
				ImageGranularity: core1_0.Extent3D{
					Width:  19,
					Height: 23,
					Depth:  29,
				},
				Flags: core1_0.SparseImageFormatNonstandardBlockSize,
			},
			ImageMipTailFirstLod: 31,
			ImageMipTailOffset:   37,
			ImageMipTailSize:     41,
			ImageMipTailStride:   43,
		}
		next = retVal[1].Next.(*core1_1.MemoryDedicatedRequirements)
		next.PrefersDedicatedAllocation = true

		return retVal, nil
	})

	reqsList, err := shim.ImageSparseMemoryRequirements2(
		core1_1.ImageSparseMemoryRequirementsInfo2{
			Image: image,
		},
		func() *core1_1.SparseImageMemoryRequirements2 {
			return &core1_1.SparseImageMemoryRequirements2{
				NextOutData: common.NextOutData{
					Next: &core1_1.MemoryDedicatedRequirements{},
				},
			}
		})

	require.NoError(t, err)
	require.Equal(t, []*core1_1.SparseImageMemoryRequirements2{
		{
			MemoryRequirements: core1_0.SparseImageMemoryRequirements{
				FormatProperties: core1_0.SparseImageFormatProperties{
					AspectMask: core1_0.ImageAspectDepth,
					ImageGranularity: core1_0.Extent3D{
						Width:  1,
						Height: 3,
						Depth:  5,
					},
					Flags: core1_0.SparseImageFormatSingleMipTail,
				},
				ImageMipTailFirstLod: 7,
				ImageMipTailOffset:   11,
				ImageMipTailSize:     13,
				ImageMipTailStride:   17,
			},
			NextOutData: common.NextOutData{
				Next: &core1_1.MemoryDedicatedRequirements{
					RequiresDedicatedAllocation: true,
				},
			},
		},
		{
			MemoryRequirements: core1_0.SparseImageMemoryRequirements{
				FormatProperties: core1_0.SparseImageFormatProperties{
					AspectMask: core1_0.ImageAspectMetadata,
					ImageGranularity: core1_0.Extent3D{
						Width:  19,
						Height: 23,
						Depth:  29,
					},
					Flags: core1_0.SparseImageFormatNonstandardBlockSize,
				},
				ImageMipTailFirstLod: 31,
				ImageMipTailOffset:   37,
				ImageMipTailSize:     41,
				ImageMipTailStride:   43,
			},
			NextOutData: common.NextOutData{
				Next: &core1_1.MemoryDedicatedRequirements{
					PrefersDedicatedAllocation: true,
				},
			},
		},
	}, reqsList)
}
