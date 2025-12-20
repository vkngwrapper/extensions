package khr_external_memory_capabilities_shim

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v2/core1_0"
	"github.com/vkngwrapper/core/v2/core1_1"
	core_mocks "github.com/vkngwrapper/core/v2/mocks"
	"github.com/vkngwrapper/extensions/v3/khr_external_memory_capabilities"
	mock_external_memory_capabilities "github.com/vkngwrapper/extensions/v3/khr_external_memory_capabilities/mocks"
	"go.uber.org/mock/gomock"
)

func TestVulkanShim_ExternalBufferProperties(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extension := mock_external_memory_capabilities.NewMockExtension(ctrl)
	physicalDevice := core_mocks.NewMockPhysicalDevice(ctrl)
	shim := NewShim(extension, physicalDevice)

	extension.EXPECT().PhysicalDeviceExternalBufferProperties(
		physicalDevice,
		khr_external_memory_capabilities.PhysicalDeviceExternalBufferInfo{
			Flags:      core1_0.BufferCreateSparseBinding,
			Usage:      core1_0.BufferUsageTransferDst,
			HandleType: khr_external_memory_capabilities.ExternalMemoryHandleTypeOpaqueFD,
		},
		gomock.Any()).
		DoAndReturn(func(
			physicalDevice core1_0.PhysicalDevice,
			o khr_external_memory_capabilities.PhysicalDeviceExternalBufferInfo,
			outData *khr_external_memory_capabilities.ExternalBufferProperties,
		) error {
			outData.ExternalMemoryProperties = khr_external_memory_capabilities.ExternalMemoryProperties{
				ExternalMemoryFeatures:        khr_external_memory_capabilities.ExternalMemoryFeatureDedicatedOnly,
				ExportFromImportedHandleTypes: khr_external_memory_capabilities.ExternalMemoryHandleTypeD3D12Resource,
				CompatibleHandleTypes:         khr_external_memory_capabilities.ExternalMemoryHandleTypeD3D11Texture,
			}

			return nil
		})

	var outData core1_1.ExternalBufferProperties
	err := shim.ExternalBufferProperties(
		core1_1.PhysicalDeviceExternalBufferInfo{
			Flags:      core1_0.BufferCreateSparseBinding,
			Usage:      core1_0.BufferUsageTransferDst,
			HandleType: core1_1.ExternalMemoryHandleTypeOpaqueFD,
		},
		&outData)
	require.NoError(t, err)
	require.Equal(t,
		core1_1.ExternalBufferProperties{
			ExternalMemoryProperties: core1_1.ExternalMemoryProperties{
				ExternalMemoryFeatures:        core1_1.ExternalMemoryFeatureDedicatedOnly,
				ExportFromImportedHandleTypes: core1_1.ExternalMemoryHandleTypeD3D12Resource,
				CompatibleHandleTypes:         core1_1.ExternalMemoryHandleTypeD3D11Texture,
			},
		},
		outData)
}
