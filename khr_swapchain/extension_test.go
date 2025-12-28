package khr_swapchain_test

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
	"github.com/vkngwrapper/core/v3/mocks"
	"github.com/vkngwrapper/extensions/v3/khr_surface"
	khr_surface_driver "github.com/vkngwrapper/extensions/v3/khr_surface/loader"
	mock_surface "github.com/vkngwrapper/extensions/v3/khr_surface/mocks"
	"github.com/vkngwrapper/extensions/v3/khr_swapchain"
	khr_swapchain_driver "github.com/vkngwrapper/extensions/v3/khr_swapchain/loader"
	mock_swapchain "github.com/vkngwrapper/extensions/v3/khr_swapchain/mocks"
	"go.uber.org/mock/gomock"
)

func TestVulkanExtension_CreateSwapchain(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	instance := mocks.NewDummyInstance(common.Vulkan1_0, []string{})
	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	expectedSwapchain := khr_swapchain.NewDummySwapchain(device)
	surface := mock_surface.NewDummySurface(instance)
	oldSwapchain := khr_swapchain.NewDummySwapchain(device)

	swapchainDriver := mock_swapchain.NewMockLoader(ctrl)
	extension := khr_swapchain.CreateExtensionDriverFromLoader(swapchainDriver, device)

	swapchainDriver.EXPECT().VkCreateSwapchainKHR(device.Handle(), gomock.Not(gomock.Nil()), gomock.Nil(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(device loader.VkDevice, pCreateInfo *khr_swapchain_driver.VkSwapchainCreateInfoKHR, pAllocator *loader.VkAllocationCallbacks, pSwapchain *khr_swapchain_driver.VkSwapchainKHR) (common.VkResult, error) {
			*pSwapchain = expectedSwapchain.Handle()

			val := reflect.ValueOf(*pCreateInfo)
			require.Equal(t, uint64(1000001000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SWAPCHAIN_CREATE_INFO_KHR
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0), val.FieldByName("flags").Uint())

			surfaceHandle := (khr_surface_driver.VkSurfaceKHR)(unsafe.Pointer(val.FieldByName("surface").Elem().UnsafeAddr()))
			require.Equal(t, surface.Handle(), surfaceHandle)

			require.Equal(t, uint64(1), val.FieldByName("minImageCount").Uint())
			require.Equal(t, uint64(67), val.FieldByName("imageFormat").Uint())    // VK_FORMAT_A2B10G10R10_SSCALED_PACK32
			require.Equal(t, uint64(0), val.FieldByName("imageColorSpace").Uint()) // VK_COLOR_SPACE_SRGB_NONLINEAR_KHR

			extent := val.FieldByName("imageExtent")
			require.Equal(t, uint64(3), extent.FieldByName("width").Uint())
			require.Equal(t, uint64(5), extent.FieldByName("height").Uint())

			require.Equal(t, uint64(7), val.FieldByName("imageArrayLayers").Uint())
			require.Equal(t, uint64(0x00000010), val.FieldByName("imageUsage").Uint()) // VK_IMAGE_USAGE_COLOR_ATTACHMENT_BIT
			require.Equal(t, uint64(1), val.FieldByName("imageSharingMode").Uint())    // VK_SHARING_MODE_CONCURRENT
			require.Equal(t, uint64(3), val.FieldByName("queueFamilyIndexCount").Uint())

			queueFamiliesPtr := (*uint32)(unsafe.Pointer(val.FieldByName("pQueueFamilyIndices").Elem().UnsafeAddr()))
			queueFamilies := ([]uint32)(unsafe.Slice(queueFamiliesPtr, 3))
			require.Equal(t, []uint32{11, 13, 17}, queueFamilies)

			require.Equal(t, uint64(0x00000020), val.FieldByName("preTransform").Uint())   // VK_SURFACE_TRANSFORM_HORIZONTAL_MIRROR_ROTATE_90_BIT_KHR
			require.Equal(t, uint64(0x00000001), val.FieldByName("compositeAlpha").Uint()) // VK_COMPOSITE_ALPHA_OPAQUE_BIT_KHR
			require.Equal(t, uint64(1), val.FieldByName("presentMode").Uint())             // VK_PRESENT_MODE_MAILBOX_KHR
			require.Equal(t, uint64(1), val.FieldByName("clipped").Uint())

			oldSwapchainHandle := (khr_swapchain_driver.VkSwapchainKHR)(unsafe.Pointer(val.FieldByName("oldSwapchain").Elem().UnsafeAddr()))
			require.Equal(t, oldSwapchain.Handle(), oldSwapchainHandle)

			return core1_0.VKSuccess, nil
		})

	swapchain, _, err := extension.CreateSwapchain(nil, khr_swapchain.SwapchainCreateInfo{
		Surface:            surface,
		MinImageCount:      1,
		ImageFormat:        core1_0.FormatA2B10G10R10SignedScaledPacked,
		ImageColorSpace:    khr_surface.ColorSpaceSRGBNonlinear,
		ImageExtent:        core1_0.Extent2D{Width: 3, Height: 5},
		ImageArrayLayers:   7,
		ImageUsage:         core1_0.ImageUsageColorAttachment,
		ImageSharingMode:   core1_0.SharingModeConcurrent,
		QueueFamilyIndices: []int{11, 13, 17},
		PreTransform:       khr_surface.TransformHorizontalMirrorRotate90,
		CompositeAlpha:     khr_surface.CompositeAlphaOpaque,
		PresentMode:        khr_surface.PresentModeMailbox,
		Clipped:            true,
		OldSwapchain:       oldSwapchain,
	})

	require.NoError(t, err)
	require.Equal(t, expectedSwapchain.Handle(), swapchain.Handle())
}

func TestVulkanExtension_PresentToQueue_NullOutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	swapchain := khr_swapchain.NewDummySwapchain(device)
	queue := mocks.NewDummyQueue(device)

	swapchainDriver := mock_swapchain.NewMockLoader(ctrl)
	extension := khr_swapchain.CreateExtensionDriverFromLoader(swapchainDriver, device)

	semaphore1 := mocks.NewDummySemaphore(device)
	semaphore2 := mocks.NewDummySemaphore(device)

	swapchainDriver.EXPECT().VkQueuePresentKHR(
		queue.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(queue loader.VkQueue, pPresentInfo *khr_swapchain_driver.VkPresentInfoKHR) (common.VkResult, error) {
			val := reflect.ValueOf(*pPresentInfo)

			require.Equal(t, uint64(1000001001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PRESENT_INFO_KHR
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(2), val.FieldByName("waitSemaphoreCount").Uint())
			require.Equal(t, uint64(1), val.FieldByName("swapchainCount").Uint())

			semaphorePtr := (*loader.VkSemaphore)(unsafe.Pointer(val.FieldByName("pWaitSemaphores").Elem().UnsafeAddr()))
			semaphores := ([]loader.VkSemaphore)(unsafe.Slice(semaphorePtr, 2))
			require.Equal(t, semaphore1.Handle(), semaphores[0])
			require.Equal(t, semaphore2.Handle(), semaphores[1])

			swapchainPtr := (*khr_swapchain_driver.VkSwapchainKHR)(unsafe.Pointer(val.FieldByName("pSwapchains").Elem().UnsafeAddr()))
			swapchains := ([]khr_swapchain_driver.VkSwapchainKHR)(unsafe.Slice(swapchainPtr, 1))
			require.Equal(t, swapchain.Handle(), swapchains[0])

			imageIndicesPtr := (*loader.Uint32)(unsafe.Pointer(val.FieldByName("pImageIndices").Elem().UnsafeAddr()))
			imageIndices := ([]loader.Uint32)(unsafe.Slice(imageIndicesPtr, 1))
			require.Equal(t, loader.Uint32(2), imageIndices[0])

			resultsPtr := (*loader.VkResult)(unsafe.Pointer(val.FieldByName("pResults").Elem().UnsafeAddr()))
			results := ([]loader.VkResult)(unsafe.Slice(resultsPtr, 1))
			results[0] = loader.VkResult(core1_0.VKSuccess)

			return core1_0.VKSuccess, nil
		})

	options := khr_swapchain.PresentInfo{
		WaitSemaphores: []core.Semaphore{semaphore1, semaphore2},
		Swapchains:     []khr_swapchain.Swapchain{swapchain},
		ImageIndices:   []int{2},
	}
	_, err := extension.QueuePresent(queue, options)
	require.NoError(t, err)
	require.Nil(t, options.OutData)
}

func TestVulkanExtension_PresentToQueue_RealOutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	swapchain := khr_swapchain.NewDummySwapchain(device)
	queue := mocks.NewDummyQueue(device)

	swapchainDriver := mock_swapchain.NewMockLoader(ctrl)
	extension := khr_swapchain.CreateExtensionDriverFromLoader(swapchainDriver, device)

	semaphore1 := mocks.NewDummySemaphore(device)
	semaphore2 := mocks.NewDummySemaphore(device)

	swapchainDriver.EXPECT().VkQueuePresentKHR(
		queue.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(queue loader.VkQueue, pPresentInfo *khr_swapchain_driver.VkPresentInfoKHR) (common.VkResult, error) {
			val := reflect.ValueOf(*pPresentInfo)

			require.Equal(t, uint64(1000001001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PRESENT_INFO_KHR
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(2), val.FieldByName("waitSemaphoreCount").Uint())
			require.Equal(t, uint64(1), val.FieldByName("swapchainCount").Uint())

			semaphorePtr := (*loader.VkSemaphore)(unsafe.Pointer(val.FieldByName("pWaitSemaphores").Elem().UnsafeAddr()))
			semaphores := ([]loader.VkSemaphore)(unsafe.Slice(semaphorePtr, 2))
			require.Equal(t, semaphore1.Handle(), semaphores[0])
			require.Equal(t, semaphore2.Handle(), semaphores[1])

			swapchainPtr := (*khr_swapchain_driver.VkSwapchainKHR)(unsafe.Pointer(val.FieldByName("pSwapchains").Elem().UnsafeAddr()))
			swapchains := ([]khr_swapchain_driver.VkSwapchainKHR)(unsafe.Slice(swapchainPtr, 1))
			require.Equal(t, swapchain.Handle(), swapchains[0])

			imageIndicesPtr := (*loader.Uint32)(unsafe.Pointer(val.FieldByName("pImageIndices").Elem().UnsafeAddr()))
			imageIndices := ([]loader.Uint32)(unsafe.Slice(imageIndicesPtr, 1))
			require.Equal(t, loader.Uint32(2), imageIndices[0])

			resultsPtr := (*loader.VkResult)(unsafe.Pointer(val.FieldByName("pResults").Elem().UnsafeAddr()))
			results := ([]loader.VkResult)(unsafe.Slice(resultsPtr, 1))
			results[0] = loader.VkResult(core1_0.VKTimeout)

			return core1_0.VKSuccess, nil
		})

	outData := khr_swapchain.PresentOutData{}
	_, err := extension.QueuePresent(queue, khr_swapchain.PresentInfo{
		WaitSemaphores: []core.Semaphore{semaphore1, semaphore2},
		Swapchains:     []khr_swapchain.Swapchain{swapchain},
		ImageIndices:   []int{2},
		OutData:        &outData,
	})
	require.NoError(t, err)

	require.Len(t, outData.Results, 1)
	require.Equal(t, core1_0.VKTimeout, outData.Results[0])
}
