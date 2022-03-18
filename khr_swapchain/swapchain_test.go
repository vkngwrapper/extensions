package khr_swapchain_test

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
	mock_driver "github.com/CannibalVox/VKng/core/driver/mocks"
	"github.com/CannibalVox/VKng/core/mocks"
	mock_surface "github.com/CannibalVox/VKng/extensions/khr_surface/mocks"
	"github.com/CannibalVox/VKng/extensions/khr_swapchain"
	mock_swapchain "github.com/CannibalVox/VKng/extensions/khr_swapchain/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
	"unsafe"
)

func TestVulkanSwapchain_AcquireNextImage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	swapchainDriver := mock_swapchain.NewMockDriver(ctrl)
	extension := khr_swapchain.CreateExtensionFromDriver(swapchainDriver)
	device := mocks.EasyMockDevice(ctrl, coreDriver)
	surface := mock_surface.EasyMockSurface(ctrl)

	swapchainDriver.EXPECT().VkCreateSwapchainKHR(mocks.Exactly(device.Handle()), gomock.Not(gomock.Nil()), gomock.Nil(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(device driver.VkDevice, pCreateInfo *khr_swapchain.VkSwapchainCreateInfoKHR, pAllocator *driver.VkAllocationCallbacks, pSwapchain *khr_swapchain.VkSwapchainKHR) (common.VkResult, error) {
			*pSwapchain = mock_swapchain.NewFakeSwapchain()

			return core1_0.VKSuccess, nil
		})

	swapchain, _, err := extension.CreateSwapchain(device, nil, &khr_swapchain.CreationOptions{
		Surface: surface,
	})
	require.NoError(t, err)

	swapchainDriver.EXPECT().VkAcquireNextImageKHR(
		mocks.Exactly(device.Handle()),
		mocks.Exactly(swapchain.Handle()),
		driver.Uint64(60000000000), // 60 billion nanoseconds = 1 minute
		driver.VkSemaphore(unsafe.Pointer(nil)),
		driver.VkFence(unsafe.Pointer(nil)),
		gomock.Not(gomock.Nil),
	).DoAndReturn(
		func(device driver.VkDevice, swapchain khr_swapchain.VkSwapchainKHR, timeout driver.Uint64, semaphore driver.VkSemaphore, fence driver.VkFence, pImageIndex *driver.Uint32) (common.VkResult, error) {
			*pImageIndex = driver.Uint32(3)

			return core1_0.VKSuccess, nil
		})

	index, _, err := swapchain.AcquireNextImage(time.Minute, nil, nil)
	require.NoError(t, err)
	require.Equal(t, 3, index)
}

func TestVulkanSwapchain_AcquireNextImage_NoTimeout(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	swapchainDriver := mock_swapchain.NewMockDriver(ctrl)
	extension := khr_swapchain.CreateExtensionFromDriver(swapchainDriver)
	device := mocks.EasyMockDevice(ctrl, coreDriver)
	surface := mock_surface.EasyMockSurface(ctrl)

	swapchainDriver.EXPECT().VkCreateSwapchainKHR(mocks.Exactly(device.Handle()), gomock.Not(gomock.Nil()), gomock.Nil(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(device driver.VkDevice, pCreateInfo *khr_swapchain.VkSwapchainCreateInfoKHR, pAllocator *driver.VkAllocationCallbacks, pSwapchain *khr_swapchain.VkSwapchainKHR) (common.VkResult, error) {
			*pSwapchain = mock_swapchain.NewFakeSwapchain()

			return core1_0.VKSuccess, nil
		})

	swapchain, _, err := extension.CreateSwapchain(device, nil, &khr_swapchain.CreationOptions{
		Surface: surface,
	})
	require.NoError(t, err)

	swapchainDriver.EXPECT().VkAcquireNextImageKHR(
		mocks.Exactly(device.Handle()),
		mocks.Exactly(swapchain.Handle()),
		driver.Uint64(^uint64(0)), // max uint64 = no timeout
		driver.VkSemaphore(unsafe.Pointer(nil)),
		driver.VkFence(unsafe.Pointer(nil)),
		gomock.Not(gomock.Nil),
	).DoAndReturn(
		func(device driver.VkDevice, swapchain khr_swapchain.VkSwapchainKHR, timeout driver.Uint64, semaphore driver.VkSemaphore, fence driver.VkFence, pImageIndex *driver.Uint32) (common.VkResult, error) {
			*pImageIndex = driver.Uint32(3)

			return core1_0.VKSuccess, nil
		})

	index, _, err := swapchain.AcquireNextImage(common.NoTimeout, nil, nil)
	require.NoError(t, err)
	require.Equal(t, 3, index)
}

func TestVulkanSwapchain_AcquireNextImage_FenceAndSemaphore(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	swapchainDriver := mock_swapchain.NewMockDriver(ctrl)
	extension := khr_swapchain.CreateExtensionFromDriver(swapchainDriver)
	device := mocks.EasyMockDevice(ctrl, coreDriver)
	surface := mock_surface.EasyMockSurface(ctrl)

	swapchainDriver.EXPECT().VkCreateSwapchainKHR(mocks.Exactly(device.Handle()), gomock.Not(gomock.Nil()), gomock.Nil(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(device driver.VkDevice, pCreateInfo *khr_swapchain.VkSwapchainCreateInfoKHR, pAllocator *driver.VkAllocationCallbacks, pSwapchain *khr_swapchain.VkSwapchainKHR) (common.VkResult, error) {
			*pSwapchain = mock_swapchain.NewFakeSwapchain()

			return core1_0.VKSuccess, nil
		})

	swapchain, _, err := extension.CreateSwapchain(device, nil, &khr_swapchain.CreationOptions{
		Surface: surface,
	})
	require.NoError(t, err)

	fence := mocks.EasyMockFence(ctrl)
	semaphore := mocks.EasyMockSemaphore(ctrl)

	swapchainDriver.EXPECT().VkAcquireNextImageKHR(
		mocks.Exactly(device.Handle()),
		mocks.Exactly(swapchain.Handle()),
		driver.Uint64(60000000000), // 60 billion nanoseconds = 1 minute
		mocks.Exactly(semaphore.Handle()),
		mocks.Exactly(fence.Handle()),
		gomock.Not(gomock.Nil),
	).DoAndReturn(
		func(device driver.VkDevice, swapchain khr_swapchain.VkSwapchainKHR, timeout driver.Uint64, semaphore driver.VkSemaphore, fence driver.VkFence, pImageIndex *driver.Uint32) (common.VkResult, error) {
			*pImageIndex = driver.Uint32(3)

			return core1_0.VKSuccess, nil
		})

	index, _, err := swapchain.AcquireNextImage(time.Minute, semaphore, fence)
	require.NoError(t, err)
	require.Equal(t, 3, index)
}

func TestVulkanSwapchain_Images(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	swapchainDriver := mock_swapchain.NewMockDriver(ctrl)
	extension := khr_swapchain.CreateExtensionFromDriver(swapchainDriver)
	device := mocks.EasyMockDevice(ctrl, coreDriver)
	surface := mock_surface.EasyMockSurface(ctrl)

	swapchainDriver.EXPECT().VkCreateSwapchainKHR(mocks.Exactly(device.Handle()), gomock.Not(gomock.Nil()), gomock.Nil(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(device driver.VkDevice, pCreateInfo *khr_swapchain.VkSwapchainCreateInfoKHR, pAllocator *driver.VkAllocationCallbacks, pSwapchain *khr_swapchain.VkSwapchainKHR) (common.VkResult, error) {
			*pSwapchain = mock_swapchain.NewFakeSwapchain()

			return core1_0.VKSuccess, nil
		})

	swapchain, _, err := extension.CreateSwapchain(device, nil, &khr_swapchain.CreationOptions{
		Surface: surface,
	})
	require.NoError(t, err)

	image1 := mocks.EasyMockImage(ctrl)
	image2 := mocks.EasyMockImage(ctrl)

	swapchainDriver.EXPECT().VkGetSwapchainImagesKHR(
		mocks.Exactly(device.Handle()),
		mocks.Exactly(swapchain.Handle()),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
	).DoAndReturn(
		func(device driver.VkDevice, swapchain khr_swapchain.VkSwapchainKHR, pSwapchainImageCount *driver.Uint32, pSwapchainImages *driver.VkImage) (common.VkResult, error) {
			*pSwapchainImageCount = driver.Uint32(2)

			return core1_0.VKSuccess, nil
		})

	swapchainDriver.EXPECT().VkGetSwapchainImagesKHR(
		mocks.Exactly(device.Handle()),
		mocks.Exactly(swapchain.Handle()),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(device driver.VkDevice, swapchain khr_swapchain.VkSwapchainKHR, pSwapchainImageCount *driver.Uint32, pSwapchainImages *driver.VkImage) (common.VkResult, error) {
			require.Equal(t, driver.Uint32(2), *pSwapchainImageCount)

			imageSlice := ([]driver.VkImage)(unsafe.Slice(pSwapchainImages, 2))
			imageSlice[0] = image1.Handle()
			imageSlice[1] = image2.Handle()

			return core1_0.VKSuccess, nil
		})

	swapchainDriver.EXPECT().CreateImage(mocks.Exactly(image1.Handle()), mocks.Exactly(device.Handle())).Return(image1)
	swapchainDriver.EXPECT().CreateImage(mocks.Exactly(image2.Handle()), mocks.Exactly(device.Handle())).Return(image2)

	images, _, err := swapchain.Images()
	require.NoError(t, err)
	require.Len(t, images, 2)
	require.Same(t, image1.Handle(), images[0].Handle())
	require.Same(t, image2.Handle(), images[1].Handle())
}

func TestVulkanSwapchain_Images_Incomplete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	swapchainDriver := mock_swapchain.NewMockDriver(ctrl)
	extension := khr_swapchain.CreateExtensionFromDriver(swapchainDriver)
	device := mocks.EasyMockDevice(ctrl, coreDriver)
	surface := mock_surface.EasyMockSurface(ctrl)

	swapchainDriver.EXPECT().VkCreateSwapchainKHR(mocks.Exactly(device.Handle()), gomock.Not(gomock.Nil()), gomock.Nil(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(device driver.VkDevice, pCreateInfo *khr_swapchain.VkSwapchainCreateInfoKHR, pAllocator *driver.VkAllocationCallbacks, pSwapchain *khr_swapchain.VkSwapchainKHR) (common.VkResult, error) {
			*pSwapchain = mock_swapchain.NewFakeSwapchain()

			return core1_0.VKSuccess, nil
		})

	swapchain, _, err := extension.CreateSwapchain(device, nil, &khr_swapchain.CreationOptions{
		Surface: surface,
	})
	require.NoError(t, err)

	image1 := mocks.EasyMockImage(ctrl)
	image2 := mocks.EasyMockImage(ctrl)

	swapchainDriver.EXPECT().VkGetSwapchainImagesKHR(
		mocks.Exactly(device.Handle()),
		mocks.Exactly(swapchain.Handle()),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
	).DoAndReturn(
		func(device driver.VkDevice, swapchain khr_swapchain.VkSwapchainKHR, pSwapchainImageCount *driver.Uint32, pSwapchainImages *driver.VkImage) (common.VkResult, error) {
			*pSwapchainImageCount = driver.Uint32(1)

			return core1_0.VKSuccess, nil
		})

	swapchainDriver.EXPECT().VkGetSwapchainImagesKHR(
		mocks.Exactly(device.Handle()),
		mocks.Exactly(swapchain.Handle()),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(device driver.VkDevice, swapchain khr_swapchain.VkSwapchainKHR, pSwapchainImageCount *driver.Uint32, pSwapchainImages *driver.VkImage) (common.VkResult, error) {
			require.Equal(t, driver.Uint32(1), *pSwapchainImageCount)

			imageSlice := ([]driver.VkImage)(unsafe.Slice(pSwapchainImages, 1))
			imageSlice[0] = image1.Handle()

			return core1_0.VKIncomplete, nil
		})

	swapchainDriver.EXPECT().VkGetSwapchainImagesKHR(
		mocks.Exactly(device.Handle()),
		mocks.Exactly(swapchain.Handle()),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
	).DoAndReturn(
		func(device driver.VkDevice, swapchain khr_swapchain.VkSwapchainKHR, pSwapchainImageCount *driver.Uint32, pSwapchainImages *driver.VkImage) (common.VkResult, error) {
			*pSwapchainImageCount = driver.Uint32(2)

			return core1_0.VKSuccess, nil
		})

	swapchainDriver.EXPECT().VkGetSwapchainImagesKHR(
		mocks.Exactly(device.Handle()),
		mocks.Exactly(swapchain.Handle()),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(device driver.VkDevice, swapchain khr_swapchain.VkSwapchainKHR, pSwapchainImageCount *driver.Uint32, pSwapchainImages *driver.VkImage) (common.VkResult, error) {
			require.Equal(t, driver.Uint32(2), *pSwapchainImageCount)

			imageSlice := ([]driver.VkImage)(unsafe.Slice(pSwapchainImages, 2))
			imageSlice[0] = image1.Handle()
			imageSlice[1] = image2.Handle()

			return core1_0.VKSuccess, nil
		})

	swapchainDriver.EXPECT().CreateImage(mocks.Exactly(image1.Handle()), mocks.Exactly(device.Handle())).Return(image1)

	swapchainDriver.EXPECT().CreateImage(mocks.Exactly(image1.Handle()), mocks.Exactly(device.Handle())).Return(image1)
	swapchainDriver.EXPECT().CreateImage(mocks.Exactly(image2.Handle()), mocks.Exactly(device.Handle())).Return(image2)

	images, _, err := swapchain.Images()
	require.NoError(t, err)
	require.Len(t, images, 2)
	require.Same(t, image1.Handle(), images[0].Handle())
	require.Same(t, image2.Handle(), images[1].Handle())
}
