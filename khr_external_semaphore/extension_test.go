package khr_external_semaphore_test

import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
	mock_driver "github.com/CannibalVox/VKng/core/driver/mocks"
	"github.com/CannibalVox/VKng/core/mocks"
	"github.com/CannibalVox/VKng/extensions/khr_external_semaphore"
	khr_external_semaphore_driver "github.com/CannibalVox/VKng/extensions/khr_external_semaphore/driver"
	"github.com/CannibalVox/VKng/extensions/khr_external_semaphore_capabilities"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func TestExportSemaphoreOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	loader, err := core.CreateLoaderFromDriver(coreDriver)
	require.NoError(t, err)
	device := mocks.EasyMockDevice(ctrl, coreDriver)
	mockSemaphore := mocks.EasyMockSemaphore(ctrl)

	coreDriver.EXPECT().VkCreateSemaphore(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device driver.VkDevice,
		pCreateInfo *driver.VkSemaphoreCreateInfo,
		pAllocator *driver.VkAllocationCallbacks,
		pSemaphore *driver.VkSemaphore,
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

	semaphore, _, err := loader.CreateSemaphore(device, nil, core1_0.SemaphoreOptions{
		HaveNext: common.HaveNext{
			khr_external_semaphore.ExportSemaphoreOptions{
				HandleTypes: khr_external_semaphore_capabilities.ExternalSemaphoreHandleTypeOpaqueWin32KMT,
			},
		},
	})
	require.NoError(t, err)
	require.Equal(t, mockSemaphore.Handle(), semaphore.Handle())
}