package khr_swapchain_test

import (
	"testing"
	"time"
	"unsafe"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
	"github.com/vkngwrapper/core/v3/mocks"
	mock_surface "github.com/vkngwrapper/extensions/v3/khr_surface/mocks"
	"github.com/vkngwrapper/extensions/v3/khr_swapchain"
	khr_swapchain_driver "github.com/vkngwrapper/extensions/v3/khr_swapchain/loader"
	mock_swapchain "github.com/vkngwrapper/extensions/v3/khr_swapchain/mocks"
	"go.uber.org/mock/gomock"
)

func TestVulkanSwapchain_AcquireNextImage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	instance := mocks.NewDummyInstance(common.Vulkan1_0, []string{})
	swapchainDriver := mock_swapchain.NewMockLoader(ctrl)
	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	surface := mock_surface.NewDummySurface(instance)
	extension := khr_swapchain.CreateExtensionDriverFromLoader(swapchainDriver, device)

	swapchainDriver.EXPECT().VkCreateSwapchainKHR(device.Handle(), gomock.Not(gomock.Nil()), gomock.Nil(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(deviceIn loader.VkDevice, pCreateInfo *khr_swapchain_driver.VkSwapchainCreateInfoKHR, pAllocator *loader.VkAllocationCallbacks, pSwapchain *khr_swapchain_driver.VkSwapchainKHR) (common.VkResult, error) {
			*pSwapchain = khr_swapchain.NewDummySwapchain(device).Handle()

			return core1_0.VKSuccess, nil
		})

	swapchain, _, err := extension.CreateSwapchain(nil, khr_swapchain.SwapchainCreateInfo{
		Surface: surface,
	})
	require.NoError(t, err)

	swapchainDriver.EXPECT().VkAcquireNextImageKHR(
		device.Handle(),
		swapchain.Handle(),
		loader.Uint64(60000000000), // 60 billion nanoseconds = 1 minute
		loader.VkSemaphore(unsafe.Pointer(nil)),
		loader.VkFence(unsafe.Pointer(nil)),
		gomock.Not(gomock.Nil),
	).DoAndReturn(
		func(device loader.VkDevice, swapchain khr_swapchain_driver.VkSwapchainKHR, timeout loader.Uint64, semaphore loader.VkSemaphore, fence loader.VkFence, pImageIndex *loader.Uint32) (common.VkResult, error) {
			*pImageIndex = loader.Uint32(3)

			return core1_0.VKSuccess, nil
		})

	index, _, err := extension.AcquireNextImage(swapchain, time.Minute, nil, nil)
	require.NoError(t, err)
	require.Equal(t, 3, index)
}

func TestVulkanSwapchain_AcquireNextImage_NoTimeout(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	instance := mocks.NewDummyInstance(common.Vulkan1_0, []string{})
	swapchainDriver := mock_swapchain.NewMockLoader(ctrl)
	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	surface := mock_surface.NewDummySurface(instance)
	extension := khr_swapchain.CreateExtensionDriverFromLoader(swapchainDriver, device)

	swapchainDriver.EXPECT().VkCreateSwapchainKHR(device.Handle(), gomock.Not(gomock.Nil()), gomock.Nil(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(deviceIn loader.VkDevice, pCreateInfo *khr_swapchain_driver.VkSwapchainCreateInfoKHR, pAllocator *loader.VkAllocationCallbacks, pSwapchain *khr_swapchain_driver.VkSwapchainKHR) (common.VkResult, error) {
			*pSwapchain = khr_swapchain.NewDummySwapchain(device).Handle()

			return core1_0.VKSuccess, nil
		})

	swapchain, _, err := extension.CreateSwapchain(nil, khr_swapchain.SwapchainCreateInfo{
		Surface: surface,
	})
	require.NoError(t, err)

	swapchainDriver.EXPECT().VkAcquireNextImageKHR(
		device.Handle(),
		swapchain.Handle(),
		loader.Uint64(^uint64(0)), // max uint64 = no timeout
		loader.VkSemaphore(unsafe.Pointer(nil)),
		loader.VkFence(unsafe.Pointer(nil)),
		gomock.Not(gomock.Nil),
	).DoAndReturn(
		func(device loader.VkDevice, swapchain khr_swapchain_driver.VkSwapchainKHR, timeout loader.Uint64, semaphore loader.VkSemaphore, fence loader.VkFence, pImageIndex *loader.Uint32) (common.VkResult, error) {
			*pImageIndex = loader.Uint32(3)

			return core1_0.VKSuccess, nil
		})

	index, _, err := extension.AcquireNextImage(swapchain, common.NoTimeout, nil, nil)
	require.NoError(t, err)
	require.Equal(t, 3, index)
}

func TestVulkanSwapchain_AcquireNextImage_FenceAndSemaphore(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	instance := mocks.NewDummyInstance(common.Vulkan1_0, []string{})
	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	surface := mock_surface.NewDummySurface(instance)
	swapchainDriver := mock_swapchain.NewMockLoader(ctrl)
	extension := khr_swapchain.CreateExtensionDriverFromLoader(swapchainDriver, device)

	swapchainDriver.EXPECT().VkCreateSwapchainKHR(device.Handle(), gomock.Not(gomock.Nil()), gomock.Nil(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(deviceIn loader.VkDevice, pCreateInfo *khr_swapchain_driver.VkSwapchainCreateInfoKHR, pAllocator *loader.VkAllocationCallbacks, pSwapchain *khr_swapchain_driver.VkSwapchainKHR) (common.VkResult, error) {
			*pSwapchain = khr_swapchain.NewDummySwapchain(device).Handle()

			return core1_0.VKSuccess, nil
		})

	swapchain, _, err := extension.CreateSwapchain(nil, khr_swapchain.SwapchainCreateInfo{
		Surface: surface,
	})
	require.NoError(t, err)

	fence := mocks.NewDummyFence(device)
	semaphore := mocks.NewDummySemaphore(device)

	swapchainDriver.EXPECT().VkAcquireNextImageKHR(
		device.Handle(),
		swapchain.Handle(),
		loader.Uint64(60000000000), // 60 billion nanoseconds = 1 minute
		semaphore.Handle(),
		fence.Handle(),
		gomock.Not(gomock.Nil),
	).DoAndReturn(
		func(device loader.VkDevice, swapchain khr_swapchain_driver.VkSwapchainKHR, timeout loader.Uint64, semaphore loader.VkSemaphore, fence loader.VkFence, pImageIndex *loader.Uint32) (common.VkResult, error) {
			*pImageIndex = loader.Uint32(3)

			return core1_0.VKSuccess, nil
		})

	index, _, err := extension.AcquireNextImage(swapchain, time.Minute, &semaphore, &fence)
	require.NoError(t, err)
	require.Equal(t, 3, index)
}

func TestVulkanSwapchain_Images(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	instance := mocks.NewDummyInstance(common.Vulkan1_0, []string{})
	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})

	image1 := mocks.NewDummyImage(device)
	image2 := mocks.NewDummyImage(device)

	swapchainDriver := mock_swapchain.NewMockLoader(ctrl)
	extension := khr_swapchain.CreateExtensionDriverFromLoader(swapchainDriver, device)
	surface := mock_surface.NewDummySurface(instance)

	swapchainDriver.EXPECT().VkCreateSwapchainKHR(device.Handle(), gomock.Not(gomock.Nil()), gomock.Nil(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(deviceIn loader.VkDevice, pCreateInfo *khr_swapchain_driver.VkSwapchainCreateInfoKHR, pAllocator *loader.VkAllocationCallbacks, pSwapchain *khr_swapchain_driver.VkSwapchainKHR) (common.VkResult, error) {
			*pSwapchain = khr_swapchain.NewDummySwapchain(device).Handle()

			return core1_0.VKSuccess, nil
		})

	swapchain, _, err := extension.CreateSwapchain(nil, khr_swapchain.SwapchainCreateInfo{
		Surface: surface,
	})
	require.NoError(t, err)

	swapchainDriver.EXPECT().VkGetSwapchainImagesKHR(
		device.Handle(),
		swapchain.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
	).DoAndReturn(
		func(device loader.VkDevice, swapchain khr_swapchain_driver.VkSwapchainKHR, pSwapchainImageCount *loader.Uint32, pSwapchainImages *loader.VkImage) (common.VkResult, error) {
			*pSwapchainImageCount = loader.Uint32(2)

			return core1_0.VKSuccess, nil
		})

	swapchainDriver.EXPECT().VkGetSwapchainImagesKHR(
		device.Handle(),
		swapchain.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(device loader.VkDevice, swapchain khr_swapchain_driver.VkSwapchainKHR, pSwapchainImageCount *loader.Uint32, pSwapchainImages *loader.VkImage) (common.VkResult, error) {
			require.Equal(t, loader.Uint32(2), *pSwapchainImageCount)

			imageSlice := ([]loader.VkImage)(unsafe.Slice(pSwapchainImages, 2))
			imageSlice[0] = image1.Handle()
			imageSlice[1] = image2.Handle()

			return core1_0.VKSuccess, nil
		})

	images, _, err := extension.GetSwapchainImages(swapchain)
	require.NoError(t, err)
	require.Len(t, images, 2)
	require.Equal(t, image1.Handle(), images[0].Handle())
	require.Equal(t, image2.Handle(), images[1].Handle())
}

func TestVulkanSwapchain_Images_Incomplete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	instance := mocks.NewDummyInstance(common.Vulkan1_0, []string{})
	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})

	image1 := mocks.NewDummyImage(device)
	image2 := mocks.NewDummyImage(device)

	swapchainDriver := mock_swapchain.NewMockLoader(ctrl)
	extension := khr_swapchain.CreateExtensionDriverFromLoader(swapchainDriver, device)

	surface := mock_surface.NewDummySurface(instance)

	swapchainDriver.EXPECT().VkCreateSwapchainKHR(device.Handle(), gomock.Not(gomock.Nil()), gomock.Nil(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(deviceIn loader.VkDevice, pCreateInfo *khr_swapchain_driver.VkSwapchainCreateInfoKHR, pAllocator *loader.VkAllocationCallbacks, pSwapchain *khr_swapchain_driver.VkSwapchainKHR) (common.VkResult, error) {
			*pSwapchain = khr_swapchain.NewDummySwapchain(device).Handle()

			return core1_0.VKSuccess, nil
		})

	swapchain, _, err := extension.CreateSwapchain(nil, khr_swapchain.SwapchainCreateInfo{
		Surface: surface,
	})
	require.NoError(t, err)

	swapchainDriver.EXPECT().VkGetSwapchainImagesKHR(
		device.Handle(),
		swapchain.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
	).DoAndReturn(
		func(device loader.VkDevice, swapchain khr_swapchain_driver.VkSwapchainKHR, pSwapchainImageCount *loader.Uint32, pSwapchainImages *loader.VkImage) (common.VkResult, error) {
			*pSwapchainImageCount = loader.Uint32(1)

			return core1_0.VKSuccess, nil
		})

	swapchainDriver.EXPECT().VkGetSwapchainImagesKHR(
		device.Handle(),
		swapchain.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(device loader.VkDevice, swapchain khr_swapchain_driver.VkSwapchainKHR, pSwapchainImageCount *loader.Uint32, pSwapchainImages *loader.VkImage) (common.VkResult, error) {
			require.Equal(t, loader.Uint32(1), *pSwapchainImageCount)

			imageSlice := ([]loader.VkImage)(unsafe.Slice(pSwapchainImages, 1))
			imageSlice[0] = image1.Handle()

			return core1_0.VKIncomplete, nil
		})

	swapchainDriver.EXPECT().VkGetSwapchainImagesKHR(
		device.Handle(),
		swapchain.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
	).DoAndReturn(
		func(device loader.VkDevice, swapchain khr_swapchain_driver.VkSwapchainKHR, pSwapchainImageCount *loader.Uint32, pSwapchainImages *loader.VkImage) (common.VkResult, error) {
			*pSwapchainImageCount = loader.Uint32(2)

			return core1_0.VKSuccess, nil
		})

	swapchainDriver.EXPECT().VkGetSwapchainImagesKHR(
		device.Handle(),
		swapchain.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(device loader.VkDevice, swapchain khr_swapchain_driver.VkSwapchainKHR, pSwapchainImageCount *loader.Uint32, pSwapchainImages *loader.VkImage) (common.VkResult, error) {
			require.Equal(t, loader.Uint32(2), *pSwapchainImageCount)

			imageSlice := ([]loader.VkImage)(unsafe.Slice(pSwapchainImages, 2))
			imageSlice[0] = image1.Handle()
			imageSlice[1] = image2.Handle()

			return core1_0.VKSuccess, nil
		})

	images, _, err := extension.GetSwapchainImages(swapchain)
	require.NoError(t, err)
	require.Len(t, images, 2)
	require.Equal(t, image1.Handle(), images[0].Handle())
	require.Equal(t, image2.Handle(), images[1].Handle())
}
