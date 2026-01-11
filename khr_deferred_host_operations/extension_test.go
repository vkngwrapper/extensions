package khr_deferred_host_operations_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
	"github.com/vkngwrapper/core/v3/mocks"
	"github.com/vkngwrapper/extensions/v3/khr_deferred_host_operations"
	khr_deferred_host_operations_loader "github.com/vkngwrapper/extensions/v3/khr_deferred_host_operations/loader"
	mock_deferred_host_operations "github.com/vkngwrapper/extensions/v3/khr_deferred_host_operations/mocks"
	"go.uber.org/mock/gomock"
)

func TestVulkanExtensionDriver_CreateDeferredOperation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	mockLoader := mock_deferred_host_operations.NewMockLoader(ctrl)
	driver := khr_deferred_host_operations.CreateExtensionDriverFromLoader(mockLoader, device)
	expectedOperation := mock_deferred_host_operations.NewDummyDeferredOperation(device)

	mockLoader.EXPECT().VkCreateDeferredOperationKHR(
		device.Handle(),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device loader.VkDevice, pAllocator *loader.VkAllocationCallbacks, pDeferredOperation *khr_deferred_host_operations_loader.VkDeferredOperationKHR) (common.VkResult, error) {
		*pDeferredOperation = expectedOperation.Handle()

		return core1_0.VKSuccess, nil
	})

	operation, res, err := driver.CreateDeferredOperation(nil)
	require.NoError(t, err)
	require.Equal(t, core1_0.VKSuccess, res)
	require.Equal(t, expectedOperation.Handle(), operation.Handle())
}

func TestVulkanExtensionDriver_DestroyDeferredOperation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	mockLoader := mock_deferred_host_operations.NewMockLoader(ctrl)
	driver := khr_deferred_host_operations.CreateExtensionDriverFromLoader(mockLoader, device)
	operation := mock_deferred_host_operations.NewDummyDeferredOperation(device)

	mockLoader.EXPECT().VkDestroyDeferredOperationKHR(
		device.Handle(),
		operation.Handle(),
		gomock.Nil(),
	)

	driver.DestroyDeferredOperation(operation, nil)
}

func TestVulkanExtensionDriver_DeferredOperationJoin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	mockLoader := mock_deferred_host_operations.NewMockLoader(ctrl)
	driver := khr_deferred_host_operations.CreateExtensionDriverFromLoader(mockLoader, device)
	operation := mock_deferred_host_operations.NewDummyDeferredOperation(device)

	mockLoader.EXPECT().VkDeferredOperationJoinKHR(
		device.Handle(),
		operation.Handle(),
	).Return(core1_0.VKSuccess, nil)

	res, err := driver.DeferredOperationJoin(operation)
	require.NoError(t, err)
	require.Equal(t, core1_0.VKSuccess, res)
}

func TestVulkanExtensionDriver_GetDeferredOperationResult(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	mockLoader := mock_deferred_host_operations.NewMockLoader(ctrl)
	driver := khr_deferred_host_operations.CreateExtensionDriverFromLoader(mockLoader, device)
	operation := mock_deferred_host_operations.NewDummyDeferredOperation(device)

	mockLoader.EXPECT().VkGetDeferredOperationResultKHR(
		device.Handle(),
		operation.Handle(),
	).Return(core1_0.VKSuccess, nil)

	res, err := driver.GetDeferredOperationResult(operation)
	require.NoError(t, err)
	require.Equal(t, core1_0.VKSuccess, res)
}

func TestVulkanExtensionDriver_GetDeferredOperationMaxConcurrency(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	mockLoader := mock_deferred_host_operations.NewMockLoader(ctrl)
	driver := khr_deferred_host_operations.CreateExtensionDriverFromLoader(mockLoader, device)
	operation := mock_deferred_host_operations.NewDummyDeferredOperation(device)

	mockLoader.EXPECT().VkGetDeferredOperationMaxConcurrencyKHR(
		device.Handle(),
		operation.Handle(),
	).Return(loader.Uint32(3))

	val := driver.GetDeferredOperationMaxConcurrency(operation)
	require.Equal(t, 3, val)
}
