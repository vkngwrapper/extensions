package khr_swapchain_test

import (
	"testing"
	"time"
	"unsafe"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/driver"
	mock_driver "github.com/vkngwrapper/core/v3/driver/mocks"
	"github.com/vkngwrapper/core/v3/mocks/mocks1_0"
	mock_surface "github.com/vkngwrapper/extensions/v3/khr_surface/mocks"
	"github.com/vkngwrapper/extensions/v3/khr_swapchain"
	khr_swapchain_driver "github.com/vkngwrapper/extensions/v3/khr_swapchain/driver"
	mock_swapchain "github.com/vkngwrapper/extensions/v3/khr_swapchain/mocks"
	"go.uber.org/mock/gomock"
)

func TestVulkanSwapchain_AcquireNextImage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	builder := mocks1_0.NewMockDeviceObjectBuilder(ctrl)

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	swapchainDriver := mock_swapchain.NewMockDriver(ctrl)
	extension := khr_swapchain.CreateExtensionFromDriver(swapchainDriver, builder)
	device := mocks1_0.EasyMockDevice(ctrl, coreDriver)
	surface := mock_surface.EasyMockSurface(ctrl)

	swapchainDriver.EXPECT().VkCreateSwapchainKHR(device.Handle(), gomock.Not(gomock.Nil()), gomock.Nil(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(device driver.VkDevice, pCreateInfo *khr_swapchain_driver.VkSwapchainCreateInfoKHR, pAllocator *driver.VkAllocationCallbacks, pSwapchain *khr_swapchain_driver.VkSwapchainKHR) (common.VkResult, error) {
			*pSwapchain = mock_swapchain.NewFakeSwapchain()

			return core1_0.VKSuccess, nil
		})

	swapchain, _, err := extension.CreateSwapchain(device, nil, khr_swapchain.SwapchainCreateInfo{
		Surface: surface,
	})
	require.NoError(t, err)

	swapchainDriver.EXPECT().VkAcquireNextImageKHR(
		device.Handle(),
		swapchain.Handle(),
		driver.Uint64(60000000000), // 60 billion nanoseconds = 1 minute
		driver.VkSemaphore(unsafe.Pointer(nil)),
		driver.VkFence(unsafe.Pointer(nil)),
		gomock.Not(gomock.Nil),
	).DoAndReturn(
		func(device driver.VkDevice, swapchain khr_swapchain_driver.VkSwapchainKHR, timeout driver.Uint64, semaphore driver.VkSemaphore, fence driver.VkFence, pImageIndex *driver.Uint32) (common.VkResult, error) {
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

	builder := mocks1_0.NewMockDeviceObjectBuilder(ctrl)

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	swapchainDriver := mock_swapchain.NewMockDriver(ctrl)
	extension := khr_swapchain.CreateExtensionFromDriver(swapchainDriver, builder)
	device := mocks1_0.EasyMockDevice(ctrl, coreDriver)
	surface := mock_surface.EasyMockSurface(ctrl)

	swapchainDriver.EXPECT().VkCreateSwapchainKHR(device.Handle(), gomock.Not(gomock.Nil()), gomock.Nil(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(device driver.VkDevice, pCreateInfo *khr_swapchain_driver.VkSwapchainCreateInfoKHR, pAllocator *driver.VkAllocationCallbacks, pSwapchain *khr_swapchain_driver.VkSwapchainKHR) (common.VkResult, error) {
			*pSwapchain = mock_swapchain.NewFakeSwapchain()

			return core1_0.VKSuccess, nil
		})

	swapchain, _, err := extension.CreateSwapchain(device, nil, khr_swapchain.SwapchainCreateInfo{
		Surface: surface,
	})
	require.NoError(t, err)

	swapchainDriver.EXPECT().VkAcquireNextImageKHR(
		device.Handle(),
		swapchain.Handle(),
		driver.Uint64(^uint64(0)), // max uint64 = no timeout
		driver.VkSemaphore(unsafe.Pointer(nil)),
		driver.VkFence(unsafe.Pointer(nil)),
		gomock.Not(gomock.Nil),
	).DoAndReturn(
		func(device driver.VkDevice, swapchain khr_swapchain_driver.VkSwapchainKHR, timeout driver.Uint64, semaphore driver.VkSemaphore, fence driver.VkFence, pImageIndex *driver.Uint32) (common.VkResult, error) {
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

	builder := mocks1_0.NewMockDeviceObjectBuilder(ctrl)

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	swapchainDriver := mock_swapchain.NewMockDriver(ctrl)
	extension := khr_swapchain.CreateExtensionFromDriver(swapchainDriver, builder)
	device := mocks1_0.EasyMockDevice(ctrl, coreDriver)
	surface := mock_surface.EasyMockSurface(ctrl)

	swapchainDriver.EXPECT().VkCreateSwapchainKHR(device.Handle(), gomock.Not(gomock.Nil()), gomock.Nil(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(device driver.VkDevice, pCreateInfo *khr_swapchain_driver.VkSwapchainCreateInfoKHR, pAllocator *driver.VkAllocationCallbacks, pSwapchain *khr_swapchain_driver.VkSwapchainKHR) (common.VkResult, error) {
			*pSwapchain = mock_swapchain.NewFakeSwapchain()

			return core1_0.VKSuccess, nil
		})

	swapchain, _, err := extension.CreateSwapchain(device, nil, khr_swapchain.SwapchainCreateInfo{
		Surface: surface,
	})
	require.NoError(t, err)

	fence := mocks1_0.EasyMockFence(ctrl)
	semaphore := mocks1_0.EasyMockSemaphore(ctrl)

	swapchainDriver.EXPECT().VkAcquireNextImageKHR(
		device.Handle(),
		swapchain.Handle(),
		driver.Uint64(60000000000), // 60 billion nanoseconds = 1 minute
		semaphore.Handle(),
		fence.Handle(),
		gomock.Not(gomock.Nil),
	).DoAndReturn(
		func(device driver.VkDevice, swapchain khr_swapchain_driver.VkSwapchainKHR, timeout driver.Uint64, semaphore driver.VkSemaphore, fence driver.VkFence, pImageIndex *driver.Uint32) (common.VkResult, error) {
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
	device := mocks1_0.EasyMockDevice(ctrl, coreDriver)

	image1 := mocks1_0.EasyMockImage(ctrl)
	image2 := mocks1_0.EasyMockImage(ctrl)

	builder := mocks1_0.NewMockDeviceObjectBuilder(ctrl)
	builder.EXPECT().CreateImageObject(coreDriver, device.Handle(), image1.Handle(), common.Vulkan1_0).Return(
		image1,
	)
	builder.EXPECT().CreateImageObject(coreDriver, device.Handle(), image2.Handle(), common.Vulkan1_0).Return(
		image2,
	)

	swapchainDriver := mock_swapchain.NewMockDriver(ctrl)
	extension := khr_swapchain.CreateExtensionFromDriver(swapchainDriver, builder)
	surface := mock_surface.EasyMockSurface(ctrl)

	swapchainDriver.EXPECT().VkCreateSwapchainKHR(device.Handle(), gomock.Not(gomock.Nil()), gomock.Nil(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(device driver.VkDevice, pCreateInfo *khr_swapchain_driver.VkSwapchainCreateInfoKHR, pAllocator *driver.VkAllocationCallbacks, pSwapchain *khr_swapchain_driver.VkSwapchainKHR) (common.VkResult, error) {
			*pSwapchain = mock_swapchain.NewFakeSwapchain()

			return core1_0.VKSuccess, nil
		})

	swapchain, _, err := extension.CreateSwapchain(device, nil, khr_swapchain.SwapchainCreateInfo{
		Surface: surface,
	})
	require.NoError(t, err)

	swapchainDriver.EXPECT().VkGetSwapchainImagesKHR(
		device.Handle(),
		swapchain.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
	).DoAndReturn(
		func(device driver.VkDevice, swapchain khr_swapchain_driver.VkSwapchainKHR, pSwapchainImageCount *driver.Uint32, pSwapchainImages *driver.VkImage) (common.VkResult, error) {
			*pSwapchainImageCount = driver.Uint32(2)

			return core1_0.VKSuccess, nil
		})

	swapchainDriver.EXPECT().VkGetSwapchainImagesKHR(
		device.Handle(),
		swapchain.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(device driver.VkDevice, swapchain khr_swapchain_driver.VkSwapchainKHR, pSwapchainImageCount *driver.Uint32, pSwapchainImages *driver.VkImage) (common.VkResult, error) {
			require.Equal(t, driver.Uint32(2), *pSwapchainImageCount)

			imageSlice := ([]driver.VkImage)(unsafe.Slice(pSwapchainImages, 2))
			imageSlice[0] = image1.Handle()
			imageSlice[1] = image2.Handle()

			return core1_0.VKSuccess, nil
		})

	images, _, err := swapchain.SwapchainImages()
	require.NoError(t, err)
	require.Len(t, images, 2)
	require.Equal(t, image1.Handle(), images[0].Handle())
	require.Equal(t, image2.Handle(), images[1].Handle())
}

func TestVulkanSwapchain_Images_Incomplete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := mocks1_0.EasyMockDevice(ctrl, coreDriver)

	image1 := mocks1_0.EasyMockImage(ctrl)
	image2 := mocks1_0.EasyMockImage(ctrl)

	builder := mocks1_0.NewMockDeviceObjectBuilder(ctrl)
	builder.EXPECT().CreateImageObject(coreDriver, device.Handle(), image1.Handle(), common.Vulkan1_0).Return(
		image1,
	).Times(2)
	builder.EXPECT().CreateImageObject(coreDriver, device.Handle(), image2.Handle(), common.Vulkan1_0).Return(
		image2,
	)

	swapchainDriver := mock_swapchain.NewMockDriver(ctrl)
	extension := khr_swapchain.CreateExtensionFromDriver(swapchainDriver, builder)

	surface := mock_surface.EasyMockSurface(ctrl)

	swapchainDriver.EXPECT().VkCreateSwapchainKHR(device.Handle(), gomock.Not(gomock.Nil()), gomock.Nil(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(device driver.VkDevice, pCreateInfo *khr_swapchain_driver.VkSwapchainCreateInfoKHR, pAllocator *driver.VkAllocationCallbacks, pSwapchain *khr_swapchain_driver.VkSwapchainKHR) (common.VkResult, error) {
			*pSwapchain = mock_swapchain.NewFakeSwapchain()

			return core1_0.VKSuccess, nil
		})

	swapchain, _, err := extension.CreateSwapchain(device, nil, khr_swapchain.SwapchainCreateInfo{
		Surface: surface,
	})
	require.NoError(t, err)

	swapchainDriver.EXPECT().VkGetSwapchainImagesKHR(
		device.Handle(),
		swapchain.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
	).DoAndReturn(
		func(device driver.VkDevice, swapchain khr_swapchain_driver.VkSwapchainKHR, pSwapchainImageCount *driver.Uint32, pSwapchainImages *driver.VkImage) (common.VkResult, error) {
			*pSwapchainImageCount = driver.Uint32(1)

			return core1_0.VKSuccess, nil
		})

	swapchainDriver.EXPECT().VkGetSwapchainImagesKHR(
		device.Handle(),
		swapchain.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(device driver.VkDevice, swapchain khr_swapchain_driver.VkSwapchainKHR, pSwapchainImageCount *driver.Uint32, pSwapchainImages *driver.VkImage) (common.VkResult, error) {
			require.Equal(t, driver.Uint32(1), *pSwapchainImageCount)

			imageSlice := ([]driver.VkImage)(unsafe.Slice(pSwapchainImages, 1))
			imageSlice[0] = image1.Handle()

			return core1_0.VKIncomplete, nil
		})

	swapchainDriver.EXPECT().VkGetSwapchainImagesKHR(
		device.Handle(),
		swapchain.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
	).DoAndReturn(
		func(device driver.VkDevice, swapchain khr_swapchain_driver.VkSwapchainKHR, pSwapchainImageCount *driver.Uint32, pSwapchainImages *driver.VkImage) (common.VkResult, error) {
			*pSwapchainImageCount = driver.Uint32(2)

			return core1_0.VKSuccess, nil
		})

	swapchainDriver.EXPECT().VkGetSwapchainImagesKHR(
		device.Handle(),
		swapchain.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(device driver.VkDevice, swapchain khr_swapchain_driver.VkSwapchainKHR, pSwapchainImageCount *driver.Uint32, pSwapchainImages *driver.VkImage) (common.VkResult, error) {
			require.Equal(t, driver.Uint32(2), *pSwapchainImageCount)

			imageSlice := ([]driver.VkImage)(unsafe.Slice(pSwapchainImages, 2))
			imageSlice[0] = image1.Handle()
			imageSlice[1] = image2.Handle()

			return core1_0.VKSuccess, nil
		})

	images, _, err := swapchain.SwapchainImages()
	require.NoError(t, err)
	require.Len(t, images, 2)
	require.Equal(t, image1.Handle(), images[0].Handle())
	require.Equal(t, image2.Handle(), images[1].Handle())
}
