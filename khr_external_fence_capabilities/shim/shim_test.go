package khr_external_fence_capabilities_shim

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v2/core1_0"
	"github.com/vkngwrapper/core/v2/core1_1"
	core_mocks "github.com/vkngwrapper/core/v2/mocks"
	"github.com/vkngwrapper/extensions/v2/khr_external_fence_capabilities"
	mock_external_fence_capabilities "github.com/vkngwrapper/extensions/v2/khr_external_fence_capabilities/mocks"
	"testing"
)

func TestVulkanShim_ExternalFenceProperties(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extension := mock_external_fence_capabilities.NewMockExtension(ctrl)
	physicalDevice := core_mocks.NewMockPhysicalDevice(ctrl)
	shim := NewShim(extension, physicalDevice)

	extension.EXPECT().PhysicalDeviceExternalFenceProperties(
		physicalDevice,
		khr_external_fence_capabilities.PhysicalDeviceExternalFenceInfo{
			HandleType: khr_external_fence_capabilities.ExternalFenceHandleTypeOpaqueWin32,
		},
		gomock.Any()).
		DoAndReturn(func(
			physicalDevice core1_0.PhysicalDevice,
			o khr_external_fence_capabilities.PhysicalDeviceExternalFenceInfo,
			outData *khr_external_fence_capabilities.ExternalFenceProperties,
		) error {
			outData.ExportFromImportedHandleTypes = khr_external_fence_capabilities.ExternalFenceHandleTypeOpaqueFD
			outData.CompatibleHandleTypes = khr_external_fence_capabilities.ExternalFenceHandleTypeSyncFD
			outData.ExternalFenceFeatures = khr_external_fence_capabilities.ExternalFenceFeatureImportable

			return nil
		})

	var outData core1_1.ExternalFenceProperties
	err := shim.ExternalFenceProperties(
		core1_1.PhysicalDeviceExternalFenceInfo{
			HandleType: core1_1.ExternalFenceHandleTypeOpaqueWin32,
		},
		&outData)
	require.NoError(t, err)
	require.Equal(t,
		core1_1.ExternalFenceProperties{
			ExportFromImportedHandleTypes: core1_1.ExternalFenceHandleTypeOpaqueFD,
			CompatibleHandleTypes:         core1_1.ExternalFenceHandleTypeSyncFD,
			ExternalFenceFeatures:         core1_1.ExternalFenceFeatureImportable,
		},
		outData)
}
