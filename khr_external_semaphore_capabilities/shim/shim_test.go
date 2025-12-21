package khr_external_semaphore_capabilities_shim

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/core1_1"
	core_mocks "github.com/vkngwrapper/core/v3/mocks/mocks1_0"
	"github.com/vkngwrapper/extensions/v3/khr_external_semaphore_capabilities"
	mock_external_semaphore_capabilities "github.com/vkngwrapper/extensions/v3/khr_external_semaphore_capabilities/mocks"
	"go.uber.org/mock/gomock"
)

func TestVulkanShim_ExternalSemaphoreProperties(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extension := mock_external_semaphore_capabilities.NewMockExtension(ctrl)
	physicalDevice := core_mocks.NewMockPhysicalDevice(ctrl)
	shim := NewShim(extension, physicalDevice)

	extension.EXPECT().PhysicalDeviceExternalSemaphoreProperties(
		physicalDevice,
		khr_external_semaphore_capabilities.PhysicalDeviceExternalSemaphoreInfo{
			HandleType: khr_external_semaphore_capabilities.ExternalSemaphoreHandleTypeOpaqueWin32,
		},
		gomock.Any()).DoAndReturn(func(
		physicalDevice core1_0.PhysicalDevice,
		o khr_external_semaphore_capabilities.PhysicalDeviceExternalSemaphoreInfo,
		outData *khr_external_semaphore_capabilities.ExternalSemaphoreProperties,
	) error {
		outData.ExportFromImportedHandleTypes = khr_external_semaphore_capabilities.ExternalSemaphoreHandleTypeD3D12Fence
		outData.CompatibleHandleTypes = khr_external_semaphore_capabilities.ExternalSemaphoreHandleTypeSyncFD
		outData.ExternalSemaphoreFeatures = khr_external_semaphore_capabilities.ExternalSemaphoreFeatureImportable
		return nil
	})

	var outData core1_1.ExternalSemaphoreProperties
	err := shim.ExternalSemaphoreProperties(
		core1_1.PhysicalDeviceExternalSemaphoreInfo{
			HandleType: core1_1.ExternalSemaphoreHandleTypeOpaqueWin32,
		},
		&outData)
	require.NoError(t, err)
	require.Equal(t, core1_1.ExternalSemaphoreProperties{
		ExportFromImportedHandleTypes: core1_1.ExternalSemaphoreHandleTypeD3D12Fence,
		CompatibleHandleTypes:         core1_1.ExternalSemaphoreHandleTypeSyncFD,
		ExternalSemaphoreFeatures:     core1_1.ExternalSemaphoreFeatureImportable,
	}, outData)
}
