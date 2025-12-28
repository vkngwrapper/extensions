package khr_external_semaphore_test

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
	"github.com/vkngwrapper/extensions/v3/khr_external_semaphore"
	khr_external_semaphore_driver "github.com/vkngwrapper/extensions/v3/khr_external_semaphore/loader"
	"github.com/vkngwrapper/extensions/v3/khr_external_semaphore_capabilities"
	"go.uber.org/mock/gomock"
)

func TestExportSemaphoreOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	mockSemaphore := mocks.NewDummySemaphore(device)

	coreLoader := mock_driver.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := mocks1_0.InternalDeviceDriver(device, coreLoader)

	coreLoader.EXPECT().VkCreateSemaphore(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device loader.VkDevice,
		pCreateInfo *loader.VkSemaphoreCreateInfo,
		pAllocator *loader.VkAllocationCallbacks,
		pSemaphore *loader.VkSemaphore,
	) (common.VkResult, error) {
		*pSemaphore = mockSemaphore.Handle()

		val := reflect.ValueOf(pCreateInfo).Elem()
		require.Equal(t, uint64(9), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SEMAPHORE_CREATE_INFO

		next := (*khr_external_semaphore_driver.VkExportSemaphoreCreateInfoKHR)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000077000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_EXPORT_SEMAPHORE_CREATE_INFO_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(4), val.FieldByName("handleTypes").Uint()) // VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT_KHR

		return core1_0.VKSuccess, nil
	})

	semaphore, _, err := driver.CreateSemaphore(nil, core1_0.SemaphoreCreateInfo{
		NextOptions: common.NextOptions{
			khr_external_semaphore.ExportSemaphoreCreateInfo{
				HandleTypes: khr_external_semaphore_capabilities.ExternalSemaphoreHandleTypeOpaqueWin32KMT,
			},
		},
	})
	require.NoError(t, err)
	require.Equal(t, mockSemaphore.Handle(), semaphore.Handle())
}
