package khr_get_physical_device_properties2_shim

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/core1_0"
	"github.com/vkngwrapper/core/v2/core1_1"
	core_mocks "github.com/vkngwrapper/core/v2/mocks"
	"github.com/vkngwrapper/extensions/v2/khr_get_physical_device_properties2"
	mock_get_physical_device_properties2 "github.com/vkngwrapper/extensions/v2/khr_get_physical_device_properties2/mocks"
	"go.uber.org/mock/gomock"
)

func TestVulkanShim_Features2(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extension := mock_get_physical_device_properties2.NewMockExtension(ctrl)
	physicalDevice := core_mocks.NewMockPhysicalDevice(ctrl)
	shim := NewShim(extension, physicalDevice)

	extension.EXPECT().PhysicalDeviceFeatures2(
		physicalDevice,
		gomock.Any(),
	).DoAndReturn(func(physicalDevice core1_0.PhysicalDevice, out *khr_get_physical_device_properties2.PhysicalDeviceFeatures2) error {
		out.Features.SampleRateShading = true

		return nil
	})

	var outData core1_1.PhysicalDeviceFeatures2
	err := shim.Features2(&outData)
	require.NoError(t, err)
	require.Equal(t,
		core1_1.PhysicalDeviceFeatures2{
			Features: core1_0.PhysicalDeviceFeatures{
				SampleRateShading: true,
			},
		}, outData)
}

func TestVulkanShim_FormatProperties2(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extension := mock_get_physical_device_properties2.NewMockExtension(ctrl)
	physicalDevice := core_mocks.NewMockPhysicalDevice(ctrl)
	shim := NewShim(extension, physicalDevice)

	extension.EXPECT().PhysicalDeviceFormatProperties2(
		physicalDevice,
		core1_0.FormatA2R10G10B10UnsignedNormalizedPacked,
		gomock.Any(),
	).DoAndReturn(func(physicalDevice core1_0.PhysicalDevice, format core1_0.Format, out *khr_get_physical_device_properties2.FormatProperties2) error {
		out.FormatProperties = core1_0.FormatProperties{
			LinearTilingFeatures:  core1_0.FormatFeatureBlitDestination,
			OptimalTilingFeatures: core1_0.FormatFeatureColorAttachmentBlend,
			BufferFeatures:        core1_0.FormatFeatureDepthStencilAttachment,
		}

		return nil
	})

	var properties core1_1.FormatProperties2
	err := shim.FormatProperties2(
		core1_0.FormatA2R10G10B10UnsignedNormalizedPacked,
		&properties,
	)
	require.NoError(t, err)
	require.Equal(t, core1_1.FormatProperties2{
		FormatProperties: core1_0.FormatProperties{
			LinearTilingFeatures:  core1_0.FormatFeatureBlitDestination,
			OptimalTilingFeatures: core1_0.FormatFeatureColorAttachmentBlend,
			BufferFeatures:        core1_0.FormatFeatureDepthStencilAttachment,
		},
	}, properties)
}

func TestVulkanShim_ImageFormatProperties2(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extension := mock_get_physical_device_properties2.NewMockExtension(ctrl)
	physicalDevice := core_mocks.NewMockPhysicalDevice(ctrl)
	shim := NewShim(extension, physicalDevice)

	extension.EXPECT().PhysicalDeviceImageFormatProperties2(
		physicalDevice,
		khr_get_physical_device_properties2.PhysicalDeviceImageFormatInfo2{
			Format: core1_0.FormatA2R10G10B10SignedScaledPacked,
			Type:   core1_0.ImageType2D,
			Tiling: core1_0.ImageTilingOptimal,
			Usage:  core1_0.ImageUsageInputAttachment,
			Flags:  core1_0.ImageCreateSparseAliased,
		},
		gomock.Any()).DoAndReturn(
		func(physicalDevice core1_0.PhysicalDevice, options khr_get_physical_device_properties2.PhysicalDeviceImageFormatInfo2, out *khr_get_physical_device_properties2.ImageFormatProperties2) (common.VkResult, error) {
			out.ImageFormatProperties = core1_0.ImageFormatProperties{
				MaxExtent: core1_0.Extent3D{
					Width:  1,
					Height: 3,
					Depth:  5,
				},
				MaxMipLevels:    7,
				MaxArrayLayers:  11,
				SampleCounts:    core1_0.Samples32,
				MaxResourceSize: 13,
			}

			return core1_0.VKSuccess, nil
		})

	var outData core1_1.ImageFormatProperties2
	_, err := shim.ImageFormatProperties2(
		core1_1.PhysicalDeviceImageFormatInfo2{
			Format: core1_0.FormatA2R10G10B10SignedScaledPacked,
			Type:   core1_0.ImageType2D,
			Tiling: core1_0.ImageTilingOptimal,
			Usage:  core1_0.ImageUsageInputAttachment,
			Flags:  core1_0.ImageCreateSparseAliased,
		},
		&outData)
	require.NoError(t, err)
	require.Equal(t, core1_1.ImageFormatProperties2{
		ImageFormatProperties: core1_0.ImageFormatProperties{
			MaxExtent: core1_0.Extent3D{
				Width:  1,
				Height: 3,
				Depth:  5,
			},
			MaxMipLevels:    7,
			MaxArrayLayers:  11,
			SampleCounts:    core1_0.Samples32,
			MaxResourceSize: 13,
		},
	}, outData)
}

func TestVulkanShim_MemoryProperties2(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extension := mock_get_physical_device_properties2.NewMockExtension(ctrl)
	physicalDevice := core_mocks.NewMockPhysicalDevice(ctrl)
	shim := NewShim(extension, physicalDevice)

	extension.EXPECT().PhysicalDeviceMemoryProperties2(
		physicalDevice,
		gomock.Any(),
	).DoAndReturn(func(physicalDevice core1_0.PhysicalDevice, out *khr_get_physical_device_properties2.PhysicalDeviceMemoryProperties2) error {
		out.MemoryProperties.MemoryTypes = []core1_0.MemoryType{
			{
				PropertyFlags: core1_0.MemoryPropertyHostCoherent,
				HeapIndex:     3,
			},
		}
		out.MemoryProperties.MemoryHeaps = []core1_0.MemoryHeap{
			{
				Size:  5,
				Flags: core1_0.MemoryHeapDeviceLocal,
			},
			{
				Size: 7,
			},
		}
		return nil
	})

	var out core1_1.PhysicalDeviceMemoryProperties2
	err := shim.MemoryProperties2(&out)
	require.NoError(t, err)
	require.Equal(t, core1_1.PhysicalDeviceMemoryProperties2{
		MemoryProperties: core1_0.PhysicalDeviceMemoryProperties{
			MemoryTypes: []core1_0.MemoryType{
				{
					PropertyFlags: core1_0.MemoryPropertyHostCoherent,
					HeapIndex:     3,
				},
			},
			MemoryHeaps: []core1_0.MemoryHeap{
				{
					Size:  5,
					Flags: core1_0.MemoryHeapDeviceLocal,
				},
				{
					Size: 7,
				},
			},
		},
	}, out)
}

func TestVulkanShim_QueueFamilyProperties2(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extension := mock_get_physical_device_properties2.NewMockExtension(ctrl)
	physicalDevice := core_mocks.NewMockPhysicalDevice(ctrl)
	shim := NewShim(extension, physicalDevice)

	extension.EXPECT().PhysicalDeviceQueueFamilyProperties2(
		physicalDevice,
		gomock.Any(),
	).DoAndReturn(
		func(physicalDevice core1_0.PhysicalDevice, outDataFactory func() *khr_get_physical_device_properties2.QueueFamilyProperties2) ([]*khr_get_physical_device_properties2.QueueFamilyProperties2, error) {
			retVal := make([]*khr_get_physical_device_properties2.QueueFamilyProperties2, 2)
			retVal[0] = outDataFactory()
			retVal[1] = outDataFactory()

			retVal[0].QueueFamilyProperties = core1_0.QueueFamilyProperties{
				QueueCount:         1,
				QueueFlags:         core1_0.QueueGraphics,
				TimestampValidBits: 3,
				MinImageTransferGranularity: core1_0.Extent3D{
					Width:  5,
					Height: 7,
					Depth:  11,
				},
			}
			retVal[1].QueueFamilyProperties = core1_0.QueueFamilyProperties{
				QueueCount:         13,
				QueueFlags:         core1_0.QueueCompute,
				TimestampValidBits: 17,
				MinImageTransferGranularity: core1_0.Extent3D{
					Width:  19,
					Height: 23,
					Depth:  29,
				},
			}

			return retVal, nil
		})

	retVal, err := shim.QueueFamilyProperties2(func() *core1_1.QueueFamilyProperties2 {
		return &core1_1.QueueFamilyProperties2{}
	})
	require.NoError(t, err)
	require.Equal(t, []*core1_1.QueueFamilyProperties2{
		{
			QueueFamilyProperties: core1_0.QueueFamilyProperties{
				QueueCount:         1,
				QueueFlags:         core1_0.QueueGraphics,
				TimestampValidBits: 3,
				MinImageTransferGranularity: core1_0.Extent3D{
					Width:  5,
					Height: 7,
					Depth:  11,
				},
			},
		},
		{
			QueueFamilyProperties: core1_0.QueueFamilyProperties{
				QueueCount:         13,
				QueueFlags:         core1_0.QueueCompute,
				TimestampValidBits: 17,
				MinImageTransferGranularity: core1_0.Extent3D{
					Width:  19,
					Height: 23,
					Depth:  29,
				},
			},
		},
	}, retVal)
}

func TestVulkanShim_Properties2(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extension := mock_get_physical_device_properties2.NewMockExtension(ctrl)
	physicalDevice := core_mocks.NewMockPhysicalDevice(ctrl)
	shim := NewShim(extension, physicalDevice)

	extension.EXPECT().PhysicalDeviceProperties2(
		physicalDevice,
		gomock.Any(),
	).DoAndReturn(func(physicalDevice core1_0.PhysicalDevice, out *khr_get_physical_device_properties2.PhysicalDeviceProperties2) error {
		out.Properties.DriverName = "wow"
		return nil
	})

	var out core1_1.PhysicalDeviceProperties2
	err := shim.Properties2(&out)
	require.NoError(t, err)
	require.Equal(t, core1_1.PhysicalDeviceProperties2{
		Properties: core1_0.PhysicalDeviceProperties{
			DriverName: "wow",
		},
	}, out)
}

func TestVulkanShim_SparseImageFormatProperties2(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extension := mock_get_physical_device_properties2.NewMockExtension(ctrl)
	physicalDevice := core_mocks.NewMockPhysicalDevice(ctrl)
	shim := NewShim(extension, physicalDevice)

	extension.EXPECT().PhysicalDeviceSparseImageFormatProperties2(
		physicalDevice,
		khr_get_physical_device_properties2.PhysicalDeviceSparseImageFormatInfo2{
			Format:  core1_0.FormatA2B10G10R10SignedScaledPacked,
			Type:    core1_0.ImageType3D,
			Samples: core1_0.Samples8,
			Usage:   core1_0.ImageUsageTransferSrc,
			Tiling:  core1_0.ImageTilingOptimal,
		}, gomock.Any()).DoAndReturn(
		func(physicalDevice core1_0.PhysicalDevice, options khr_get_physical_device_properties2.PhysicalDeviceSparseImageFormatInfo2, outDataFactory func() *khr_get_physical_device_properties2.SparseImageFormatProperties2) ([]*khr_get_physical_device_properties2.SparseImageFormatProperties2, error) {
			retVal := make([]*khr_get_physical_device_properties2.SparseImageFormatProperties2, 2)
			retVal[0] = outDataFactory()
			retVal[1] = outDataFactory()

			retVal[0].Properties = core1_0.SparseImageFormatProperties{
				AspectMask: core1_0.ImageAspectMetadata,
				ImageGranularity: core1_0.Extent3D{
					Width:  1,
					Height: 3,
					Depth:  5,
				},
				Flags: core1_0.SparseImageFormatNonstandardBlockSize,
			}
			retVal[1].Properties = core1_0.SparseImageFormatProperties{
				AspectMask: core1_0.ImageAspectStencil,
				ImageGranularity: core1_0.Extent3D{
					Width:  7,
					Height: 11,
					Depth:  13,
				},
				Flags: core1_0.SparseImageFormatAlignedMipSize,
			}

			return retVal, nil
		})

	retVal, err := shim.SparseImageFormatProperties2(
		core1_1.PhysicalDeviceSparseImageFormatInfo2{
			Format:  core1_0.FormatA2B10G10R10SignedScaledPacked,
			Type:    core1_0.ImageType3D,
			Samples: core1_0.Samples8,
			Usage:   core1_0.ImageUsageTransferSrc,
			Tiling:  core1_0.ImageTilingOptimal,
		}, func() *core1_1.SparseImageFormatProperties2 {
			return &core1_1.SparseImageFormatProperties2{}
		})
	require.NoError(t, err)
	require.Equal(t, []*core1_1.SparseImageFormatProperties2{
		{
			Properties: core1_0.SparseImageFormatProperties{
				AspectMask: core1_0.ImageAspectMetadata,
				ImageGranularity: core1_0.Extent3D{
					Width:  1,
					Height: 3,
					Depth:  5,
				},
				Flags: core1_0.SparseImageFormatNonstandardBlockSize,
			},
		},
		{
			Properties: core1_0.SparseImageFormatProperties{
				AspectMask: core1_0.ImageAspectStencil,
				ImageGranularity: core1_0.Extent3D{
					Width:  7,
					Height: 11,
					Depth:  13,
				},
				Flags: core1_0.SparseImageFormatAlignedMipSize,
			},
		},
	}, retVal)
}
