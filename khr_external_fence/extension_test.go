package khr_external_fence_test

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
	mock_driver "github.com/vkngwrapper/core/v3/loader/mocks"
	"github.com/vkngwrapper/core/v3/mocks"
	"github.com/vkngwrapper/core/v3/mocks/mocks1_0"
	"github.com/vkngwrapper/extensions/v3/khr_external_fence"
	khr_external_fence_driver "github.com/vkngwrapper/extensions/v3/khr_external_fence/loader"
	"github.com/vkngwrapper/extensions/v3/khr_external_fence_capabilities"
	"go.uber.org/mock/gomock"
)

func TestExportFenceOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreLoader := mock_driver.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := mocks1_0.InternalDeviceDriver(coreLoader)
	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	mockFence := mocks.NewDummyFence(device)

	coreLoader.EXPECT().VkCreateFence(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device loader.VkDevice, pCreateInfo *loader.VkFenceCreateInfo, pAllocator *loader.VkAllocationCallbacks, pFence *loader.VkFence) (common.VkResult, error) {
		*pFence = mockFence.Handle()

		val := reflect.ValueOf(pCreateInfo).Elem()
		require.Equal(t, uint64(8), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_FENCE_CREATE_INFO
		require.Equal(t, uint64(1), val.FieldByName("flags").Uint()) // VK_FENCE_CREATE_SIGNALED_BIT

		next := (*khr_external_fence_driver.VkExportFenceCreateInfoKHR)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000113000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_EXPORT_FENCE_CREATE_INFO_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(2), val.FieldByName("handleTypes").Uint()) // VK_EXTERNAL_FENCE_HANDLE_TYPE_OPAQUE_WIN32_BIT_KHR

		return core1_0.VKSuccess, nil
	})

	fence, _, err := driver.CreateFence(
		device,
		nil,
		core1_0.FenceCreateInfo{
			Flags: core1_0.FenceCreateSignaled,

			NextOptions: common.NextOptions{
				khr_external_fence.ExportFenceCreateInfo{
					HandleTypes: khr_external_fence_capabilities.ExternalFenceHandleTypeOpaqueWin32,
				},
			},
		})
	require.NoError(t, err)
	require.Equal(t, mockFence.Handle(), fence.Handle())
}
