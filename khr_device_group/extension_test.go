package khr_device_group_test

import (
	"reflect"
	"testing"
	"time"
	"unsafe"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
	mock_driver "github.com/vkngwrapper/core/v3/loader/mocks"
	"github.com/vkngwrapper/core/v3/mocks"
	"github.com/vkngwrapper/core/v3/mocks/mocks1_0"
	"github.com/vkngwrapper/extensions/v3/khr_bind_memory2"
	khr_bind_memory2_driver "github.com/vkngwrapper/extensions/v3/khr_bind_memory2/loader"
	mock_bind_memory2 "github.com/vkngwrapper/extensions/v3/khr_bind_memory2/mocks"
	"github.com/vkngwrapper/extensions/v3/khr_device_group"
	khr_device_group_driver "github.com/vkngwrapper/extensions/v3/khr_device_group/loader"
	mock_device_group "github.com/vkngwrapper/extensions/v3/khr_device_group/mocks"
	khr_surface_driver "github.com/vkngwrapper/extensions/v3/khr_surface/loader"
	mock_surface "github.com/vkngwrapper/extensions/v3/khr_surface/mocks"
	"github.com/vkngwrapper/extensions/v3/khr_swapchain"
	khr_swapchain_driver "github.com/vkngwrapper/extensions/v3/khr_swapchain/loader"
	mock_swapchain "github.com/vkngwrapper/extensions/v3/khr_swapchain/mocks"
	"go.uber.org/mock/gomock"
)

func TestVulkanExtension_CmdDispatchBase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	commandPool := mocks.NewDummyCommandPool(device)
	commandBuffer := mocks.NewDummyCommandBuffer(commandPool, device)

	extDriver := mock_device_group.NewMockLoader(ctrl)
	extension := khr_device_group.CreateExtensionDriverFromLoader(extDriver, device, false, false)

	extDriver.EXPECT().VkCmdDispatchBaseKHR(
		commandBuffer.Handle(),
		loader.Uint32(1),
		loader.Uint32(3),
		loader.Uint32(5),
		loader.Uint32(7),
		loader.Uint32(11),
		loader.Uint32(13),
	)

	extension.CmdDispatchBase(commandBuffer, 1, 3, 5, 7, 11, 13)
}

func TestVulkanExtension_CmdSetDeviceMask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	commandPool := mocks.NewDummyCommandPool(device)
	commandBuffer := mocks.NewDummyCommandBuffer(commandPool, device)

	extDriver := mock_device_group.NewMockLoader(ctrl)
	extension := khr_device_group.CreateExtensionDriverFromLoader(extDriver, device, false, false)

	extDriver.EXPECT().VkCmdSetDeviceMaskKHR(commandBuffer.Handle(), loader.Uint32(3))

	extension.CmdSetDeviceMask(commandBuffer, 3)
}

func TestVulkanExtension_GetDeviceGroupPeerMemoryFeatures(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})

	extDriver := mock_device_group.NewMockLoader(ctrl)
	extension := khr_device_group.CreateExtensionDriverFromLoader(extDriver, device, false, false)

	extDriver.EXPECT().VkGetDeviceGroupPeerMemoryFeaturesKHR(
		device.Handle(),
		loader.Uint32(1),
		loader.Uint32(3),
		loader.Uint32(5),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(
		device loader.VkDevice,
		heapIndex, localDeviceIndex, remoteDeviceIndex loader.Uint32,
		pPeerMemoryFeatures *khr_device_group_driver.VkPeerMemoryFeatureFlagsKHR,
	) {
		*pPeerMemoryFeatures = khr_device_group_driver.VkPeerMemoryFeatureFlagsKHR(1) // VK_PEER_MEMORY_FEATURE_COPY_SRC_BIT_KHR
	})

	features := extension.GetDeviceGroupPeerMemoryFeatures(
		1, 3, 5,
	)
	require.Equal(t, khr_device_group.PeerMemoryFeatureCopySrc, features)
}

func TestVulkanExtension_WithKHRSurface(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_device_group.NewMockLoader(ctrl)
	extension := khr_device_group.CreateExtensionDriverFromLoader(extDriver, mocks.NewDummyDevice(common.Vulkan1_0, []string{}), true, true)
	require.NotNil(t, extension.WithKHRSurface())
	require.NotNil(t, extension.WithKHRSwapchain())
}

func TestVulkanExtension_WithKHRSurface_None(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_device_group.NewMockLoader(ctrl)
	extension := khr_device_group.CreateExtensionDriverFromLoader(extDriver, mocks.NewDummyDevice(common.Vulkan1_0, []string{}), false, false)
	require.Nil(t, extension.WithKHRSurface())
	require.Nil(t, extension.WithKHRSwapchain())
}

func TestVulkanExtensionWithKHRSurface_GetDeviceGroupPresentCapabilities(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})

	extDriver := mock_device_group.NewMockLoader(ctrl)
	extension := khr_device_group.CreateExtensionDriverFromLoader(extDriver, device, true, false)

	extDriver.EXPECT().VkGetDeviceGroupPresentCapabilitiesKHR(
		device.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device loader.VkDevice,
		info *khr_device_group_driver.VkDeviceGroupPresentCapabilitiesKHR) (common.VkResult, error) {

		val := reflect.ValueOf(info).Elem()
		require.Equal(t, uint64(1000060007), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_GROUP_PRESENT_CAPABILITIES_KHR
		require.True(t, val.FieldByName("pNext").IsNil())

		mask := val.FieldByName("presentMask")
		*(*uint32)(unsafe.Pointer(mask.Index(0).UnsafeAddr())) = uint32(1)
		*(*uint32)(unsafe.Pointer(mask.Index(1).UnsafeAddr())) = uint32(2)
		*(*uint32)(unsafe.Pointer(mask.Index(2).UnsafeAddr())) = uint32(7)
		for i := 3; i < 32; i++ {
			*(*uint32)(unsafe.Pointer(mask.Index(i).UnsafeAddr())) = uint32(0)
		}
		*(*uint32)(unsafe.Pointer(val.FieldByName("modes").UnsafeAddr())) = 0

		return core1_0.VKSuccess, nil
	})

	var outData khr_device_group.DeviceGroupPresentCapabilities
	_, err := extension.WithKHRSurface().GetDeviceGroupPresentCapabilities(
		&outData,
	)
	require.NoError(t, err)
	require.Equal(t, khr_device_group.DeviceGroupPresentCapabilities{
		PresentMask: [32]uint32{1, 2, 7, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}, outData)
}

func TestVulkanExtensionWithKHRSurface_GetDeviceGroupSurfacePresentModes(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	instance := mocks.NewDummyInstance(common.Vulkan1_0, []string{})
	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	surface := mock_surface.NewDummySurface(instance)

	extDriver := mock_device_group.NewMockLoader(ctrl)
	extension := khr_device_group.CreateExtensionDriverFromLoader(extDriver, device, true, false)

	extDriver.EXPECT().VkGetDeviceGroupSurfacePresentModesKHR(
		device.Handle(),
		surface.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device loader.VkDevice,
		surface khr_surface_driver.VkSurfaceKHR,
		pModes *khr_device_group_driver.VkDeviceGroupPresentModeFlagsKHR) (common.VkResult, error) {

		*pModes = khr_device_group_driver.VkDeviceGroupPresentModeFlagsKHR(4) // VK_DEVICE_GROUP_PRESENT_MODE_SUM_BIT_KHR

		return core1_0.VKSuccess, nil
	})

	modes, _, err := extension.WithKHRSurface().GetDeviceGroupSurfacePresentModes(surface)
	require.NoError(t, err)
	require.Equal(t, khr_device_group.DeviceGroupPresentModeSum, modes)
}

func TestVulkanExtensionWithKHRSurface_GetPhysicalDevicePresentRectangles(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	instance := mocks.NewDummyInstance(common.Vulkan1_0, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_0)
	surface := mock_surface.NewDummySurface(instance)
	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})

	extDriver := mock_device_group.NewMockLoader(ctrl)
	extension := khr_device_group.CreateExtensionDriverFromLoader(extDriver, device, true, false)

	extDriver.EXPECT().VkGetPhysicalDevicePresentRectanglesKHR(
		physicalDevice.Handle(),
		surface.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
	).DoAndReturn(func(physicalDevice loader.VkPhysicalDevice,
		surface khr_surface_driver.VkSurfaceKHR,
		pRectCount *loader.Uint32,
		pRects *loader.VkRect2D,
	) (common.VkResult, error) {
		*pRectCount = loader.Uint32(3)

		return core1_0.VKSuccess, nil
	})

	extDriver.EXPECT().VkGetPhysicalDevicePresentRectanglesKHR(
		physicalDevice.Handle(),
		surface.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice loader.VkPhysicalDevice,
		surface khr_surface_driver.VkSurfaceKHR,
		pRectCount *loader.Uint32,
		pRects *loader.VkRect2D,
	) (common.VkResult, error) {
		require.Equal(t, loader.Uint32(3), *pRectCount)

		rectSlice := ([]loader.VkRect2D)(unsafe.Slice(pRects, 3))
		val := reflect.ValueOf(rectSlice)

		r := val.Index(0)
		*(*int32)(unsafe.Pointer(r.FieldByName("offset").FieldByName("x").UnsafeAddr())) = int32(1)
		*(*int32)(unsafe.Pointer(r.FieldByName("offset").FieldByName("y").UnsafeAddr())) = int32(3)
		*(*uint32)(unsafe.Pointer(r.FieldByName("extent").FieldByName("width").UnsafeAddr())) = uint32(5)
		*(*uint32)(unsafe.Pointer(r.FieldByName("extent").FieldByName("height").UnsafeAddr())) = uint32(7)

		r = val.Index(1)
		*(*int32)(unsafe.Pointer(r.FieldByName("offset").FieldByName("x").UnsafeAddr())) = int32(11)
		*(*int32)(unsafe.Pointer(r.FieldByName("offset").FieldByName("y").UnsafeAddr())) = int32(13)
		*(*uint32)(unsafe.Pointer(r.FieldByName("extent").FieldByName("width").UnsafeAddr())) = uint32(17)
		*(*uint32)(unsafe.Pointer(r.FieldByName("extent").FieldByName("height").UnsafeAddr())) = uint32(19)

		r = val.Index(2)
		*(*int32)(unsafe.Pointer(r.FieldByName("offset").FieldByName("x").UnsafeAddr())) = int32(23)
		*(*int32)(unsafe.Pointer(r.FieldByName("offset").FieldByName("y").UnsafeAddr())) = int32(29)
		*(*uint32)(unsafe.Pointer(r.FieldByName("extent").FieldByName("width").UnsafeAddr())) = uint32(31)
		*(*uint32)(unsafe.Pointer(r.FieldByName("extent").FieldByName("height").UnsafeAddr())) = uint32(37)

		return core1_0.VKSuccess, nil
	})

	rects, _, err := extension.WithKHRSurface().GetPhysicalDevicePresentRectangles(physicalDevice, surface)
	require.NoError(t, err)
	require.Equal(t, []core1_0.Rect2D{
		{
			Offset: core1_0.Offset2D{X: 1, Y: 3},
			Extent: core1_0.Extent2D{Width: 5, Height: 7},
		},
		{
			Offset: core1_0.Offset2D{X: 11, Y: 13},
			Extent: core1_0.Extent2D{Width: 17, Height: 19},
		},
		{
			Offset: core1_0.Offset2D{X: 23, Y: 29},
			Extent: core1_0.Extent2D{Width: 31, Height: 37},
		},
	}, rects)
}

func TestVulkanExtensionWithKHRSurface_GetPhysicalDevicePresentRectangles_Incomplete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	extDriver := mock_device_group.NewMockLoader(ctrl)
	extension := khr_device_group.CreateExtensionDriverFromLoader(extDriver, device, true, false)

	instance := mocks.NewDummyInstance(common.Vulkan1_0, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_0)
	surface := mock_surface.NewDummySurface(instance)

	extDriver.EXPECT().VkGetPhysicalDevicePresentRectanglesKHR(
		physicalDevice.Handle(),
		surface.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
	).DoAndReturn(func(physicalDevice loader.VkPhysicalDevice,
		surface khr_surface_driver.VkSurfaceKHR,
		pRectCount *loader.Uint32,
		pRects *loader.VkRect2D,
	) (common.VkResult, error) {
		*pRectCount = loader.Uint32(2)

		return core1_0.VKSuccess, nil
	})

	extDriver.EXPECT().VkGetPhysicalDevicePresentRectanglesKHR(
		physicalDevice.Handle(),
		surface.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice loader.VkPhysicalDevice,
		surface khr_surface_driver.VkSurfaceKHR,
		pRectCount *loader.Uint32,
		pRects *loader.VkRect2D,
	) (common.VkResult, error) {
		require.Equal(t, loader.Uint32(2), *pRectCount)

		rectSlice := ([]loader.VkRect2D)(unsafe.Slice(pRects, 2))
		val := reflect.ValueOf(rectSlice)

		r := val.Index(0)
		*(*int32)(unsafe.Pointer(r.FieldByName("offset").FieldByName("x").UnsafeAddr())) = int32(1)
		*(*int32)(unsafe.Pointer(r.FieldByName("offset").FieldByName("y").UnsafeAddr())) = int32(3)
		*(*uint32)(unsafe.Pointer(r.FieldByName("extent").FieldByName("width").UnsafeAddr())) = uint32(5)
		*(*uint32)(unsafe.Pointer(r.FieldByName("extent").FieldByName("height").UnsafeAddr())) = uint32(7)

		r = val.Index(1)
		*(*int32)(unsafe.Pointer(r.FieldByName("offset").FieldByName("x").UnsafeAddr())) = int32(11)
		*(*int32)(unsafe.Pointer(r.FieldByName("offset").FieldByName("y").UnsafeAddr())) = int32(13)
		*(*uint32)(unsafe.Pointer(r.FieldByName("extent").FieldByName("width").UnsafeAddr())) = uint32(17)
		*(*uint32)(unsafe.Pointer(r.FieldByName("extent").FieldByName("height").UnsafeAddr())) = uint32(19)

		return core1_0.VKIncomplete, nil
	})

	extDriver.EXPECT().VkGetPhysicalDevicePresentRectanglesKHR(
		physicalDevice.Handle(),
		surface.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
	).DoAndReturn(func(physicalDevice loader.VkPhysicalDevice,
		surface khr_surface_driver.VkSurfaceKHR,
		pRectCount *loader.Uint32,
		pRects *loader.VkRect2D,
	) (common.VkResult, error) {
		*pRectCount = loader.Uint32(3)

		return core1_0.VKSuccess, nil
	})

	extDriver.EXPECT().VkGetPhysicalDevicePresentRectanglesKHR(
		physicalDevice.Handle(),
		surface.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice loader.VkPhysicalDevice,
		surface khr_surface_driver.VkSurfaceKHR,
		pRectCount *loader.Uint32,
		pRects *loader.VkRect2D,
	) (common.VkResult, error) {
		require.Equal(t, loader.Uint32(3), *pRectCount)

		rectSlice := ([]loader.VkRect2D)(unsafe.Slice(pRects, 3))
		val := reflect.ValueOf(rectSlice)

		r := val.Index(0)
		*(*int32)(unsafe.Pointer(r.FieldByName("offset").FieldByName("x").UnsafeAddr())) = int32(1)
		*(*int32)(unsafe.Pointer(r.FieldByName("offset").FieldByName("y").UnsafeAddr())) = int32(3)
		*(*uint32)(unsafe.Pointer(r.FieldByName("extent").FieldByName("width").UnsafeAddr())) = uint32(5)
		*(*uint32)(unsafe.Pointer(r.FieldByName("extent").FieldByName("height").UnsafeAddr())) = uint32(7)

		r = val.Index(1)
		*(*int32)(unsafe.Pointer(r.FieldByName("offset").FieldByName("x").UnsafeAddr())) = int32(11)
		*(*int32)(unsafe.Pointer(r.FieldByName("offset").FieldByName("y").UnsafeAddr())) = int32(13)
		*(*uint32)(unsafe.Pointer(r.FieldByName("extent").FieldByName("width").UnsafeAddr())) = uint32(17)
		*(*uint32)(unsafe.Pointer(r.FieldByName("extent").FieldByName("height").UnsafeAddr())) = uint32(19)

		r = val.Index(2)
		*(*int32)(unsafe.Pointer(r.FieldByName("offset").FieldByName("x").UnsafeAddr())) = int32(23)
		*(*int32)(unsafe.Pointer(r.FieldByName("offset").FieldByName("y").UnsafeAddr())) = int32(29)
		*(*uint32)(unsafe.Pointer(r.FieldByName("extent").FieldByName("width").UnsafeAddr())) = uint32(31)
		*(*uint32)(unsafe.Pointer(r.FieldByName("extent").FieldByName("height").UnsafeAddr())) = uint32(37)

		return core1_0.VKSuccess, nil
	})

	rects, _, err := extension.WithKHRSurface().GetPhysicalDevicePresentRectangles(physicalDevice, surface)
	require.NoError(t, err)
	require.Equal(t, []core1_0.Rect2D{
		{
			Offset: core1_0.Offset2D{X: 1, Y: 3},
			Extent: core1_0.Extent2D{Width: 5, Height: 7},
		},
		{
			Offset: core1_0.Offset2D{X: 11, Y: 13},
			Extent: core1_0.Extent2D{Width: 17, Height: 19},
		},
		{
			Offset: core1_0.Offset2D{X: 23, Y: 29},
			Extent: core1_0.Extent2D{Width: 31, Height: 37},
		},
	}, rects)
}

func TestVulkanExtensionWithKHRSwapchain_AcquireNextImage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	swapchain := khr_swapchain.NewDummySwapchain(device)
	semaphore := mocks.NewDummySemaphore(device)

	extDriver := mock_device_group.NewMockLoader(ctrl)
	extension := khr_device_group.CreateExtensionDriverFromLoader(extDriver, device, false, true)

	extDriver.EXPECT().VkAcquireNextImage2KHR(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device loader.VkDevice,
		pAcquireInfo *khr_device_group_driver.VkAcquireNextImageInfoKHR,
		pImageIndex *loader.Uint32,
	) (common.VkResult, error) {
		*pImageIndex = loader.Uint32(2)

		val := reflect.ValueOf(pAcquireInfo).Elem()
		require.Equal(t, uint64(1000060010), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_ACQUIRE_NEXT_IMAGE_INFO_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, swapchain.Handle(), khr_swapchain_driver.VkSwapchainKHR(val.FieldByName("swapchain").UnsafePointer()))
		require.Equal(t, uint64(1000000000), val.FieldByName("timeout").Uint())
		require.Equal(t, semaphore.Handle(), loader.VkSemaphore(val.FieldByName("semaphore").UnsafePointer()))
		require.True(t, val.FieldByName("fence").IsNil())
		require.Equal(t, uint64(3), val.FieldByName("deviceMask").Uint())

		return core1_0.VKSuccess, nil
	})

	index, _, err := extension.WithKHRSwapchain().AcquireNextImage2(
		khr_device_group.AcquireNextImageInfo{
			Swapchain:  swapchain,
			Timeout:    time.Second,
			Semaphore:  semaphore,
			DeviceMask: 3,
		},
	)
	require.NoError(t, err)
	require.Equal(t, 2, index)
}

func TestMemoryAllocateFlagsOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	mockMemory := mocks.NewDummyDeviceMemory(device, 1)

	coreLoader := mock_driver.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := mocks1_0.InternalDeviceDriver(device, coreLoader)

	coreLoader.EXPECT().VkAllocateMemory(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(device loader.VkDevice,
			pAllocateInfo *loader.VkMemoryAllocateInfo,
			pAllocator *loader.VkAllocationCallbacks,
			pMemory *loader.VkDeviceMemory,
		) (common.VkResult, error) {
			*pMemory = mockMemory.Handle()

			val := reflect.ValueOf(pAllocateInfo).Elem()
			require.Equal(t, uint64(5), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_MEMORY_ALLOCATE_INFO
			require.Equal(t, uint64(1), val.FieldByName("allocationSize").Uint())
			require.Equal(t, uint64(3), val.FieldByName("memoryTypeIndex").Uint())

			next := (*khr_device_group_driver.VkMemoryAllocateFlagsInfoKHR)(val.FieldByName("pNext").UnsafePointer())
			val = reflect.ValueOf(next).Elem()
			require.Equal(t, uint64(1000060000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_MEMORY_ALLOCATE_FLAGS_INFO_KHR
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(1), val.FieldByName("flags").Uint()) //VK_MEMORY_ALLOCATE_DEVICE_MASK_BIT_KHR
			require.Equal(t, uint64(5), val.FieldByName("deviceMask").Uint())

			return core1_0.VKSuccess, nil
		})

	memory, _, err := driver.AllocateMemory(nil,
		core1_0.MemoryAllocateInfo{
			AllocationSize:  1,
			MemoryTypeIndex: 3,
			NextOptions: common.NextOptions{Next: khr_device_group.MemoryAllocateFlagsInfo{
				Flags:      khr_device_group.MemoryAllocateDeviceMask,
				DeviceMask: 5,
			}},
		})
	require.NoError(t, err)
	require.Equal(t, mockMemory.Handle(), memory.Handle())
}

func TestDeviceGroupCommandBufferBeginOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	commandPool := mocks.NewDummyCommandPool(device)

	coreLoader := mock_driver.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := mocks1_0.InternalDeviceDriver(device, coreLoader)

	commandBuffer := mocks.NewDummyCommandBuffer(commandPool, device)

	coreLoader.EXPECT().VkBeginCommandBuffer(
		commandBuffer.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(commandBuffer loader.VkCommandBuffer, pBeginInfo *loader.VkCommandBufferBeginInfo) (common.VkResult, error) {
		val := reflect.ValueOf(pBeginInfo).Elem()

		require.Equal(t, uint64(42), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_COMMAND_BUFFER_BEGIN_INFO
		require.Equal(t, uint64(1), val.FieldByName("flags").Uint())  // VK_COMMAND_BUFFER_USAGE_ONE_TIME_SUBMIT_BIT
		require.True(t, val.FieldByName("pInheritanceInfo").IsNil())

		next := (*khr_device_group_driver.VkDeviceGroupCommandBufferBeginInfoKHR)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000060004), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_GROUP_COMMAND_BUFFER_BEGIN_INFO_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(3), val.FieldByName("deviceMask").Uint())

		return core1_0.VKSuccess, nil
	})

	_, err := driver.BeginCommandBuffer(commandBuffer, core1_0.CommandBufferBeginInfo{
		Flags: core1_0.CommandBufferUsageOneTimeSubmit,
		NextOptions: common.NextOptions{Next: khr_device_group.DeviceGroupCommandBufferBeginInfo{
			DeviceMask: 3,
		}},
	})
	require.NoError(t, err)
}

func TestBindBufferMemoryDeviceGroupOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	buffer := mocks.NewDummyBuffer(device)
	memory := mocks.NewDummyDeviceMemory(device, 1)

	extDriver := mock_bind_memory2.NewMockLoader(ctrl)
	extension := khr_bind_memory2.CreateExtensionDriverFromLoader(extDriver, device)

	extDriver.EXPECT().VkBindBufferMemory2KHR(
		device.Handle(),
		loader.Uint32(1),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(
		device loader.VkDevice,
		infoCount loader.Uint32,
		pInfo *khr_bind_memory2_driver.VkBindBufferMemoryInfoKHR,
	) (common.VkResult, error) {
		val := reflect.ValueOf(pInfo).Elem()

		require.Equal(t, uint64(1000157000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_BIND_BUFFER_MEMORY_INFO_KHR
		require.Equal(t, buffer.Handle(), (loader.VkBuffer)(val.FieldByName("buffer").UnsafePointer()))
		require.Equal(t, memory.Handle(), (loader.VkDeviceMemory)(val.FieldByName("memory").UnsafePointer()))
		require.Equal(t, uint64(1), val.FieldByName("memoryOffset").Uint())

		next := (*khr_device_group_driver.VkBindBufferMemoryDeviceGroupInfoKHR)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()
		require.Equal(t, uint64(1000060013), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_BIND_BUFFER_MEMORY_DEVICE_GROUP_INFO_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(3), val.FieldByName("deviceIndexCount").Uint())

		indices := (*loader.Uint32)(val.FieldByName("pDeviceIndices").UnsafePointer())
		indexSlice := ([]loader.Uint32)(unsafe.Slice(indices, 3))
		val = reflect.ValueOf(indexSlice)

		require.Equal(t, uint64(1), val.Index(0).Uint())
		require.Equal(t, uint64(2), val.Index(1).Uint())
		require.Equal(t, uint64(7), val.Index(2).Uint())

		return core1_0.VKSuccess, nil
	})

	_, err := extension.BindBufferMemory2(
		khr_bind_memory2.BindBufferMemoryInfo{
			Buffer:       buffer,
			Memory:       memory,
			MemoryOffset: 1,

			NextOptions: common.NextOptions{
				khr_device_group.BindBufferMemoryDeviceGroupInfo{
					DeviceIndices: []int{1, 2, 7},
				},
			},
		},
	)
	require.NoError(t, err)
}

func TestBindImageMemoryDeviceGroupOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	image := mocks.NewDummyImage(device)
	memory := mocks.NewDummyDeviceMemory(device, 1)

	extDriver := mock_bind_memory2.NewMockLoader(ctrl)
	extension := khr_bind_memory2.CreateExtensionDriverFromLoader(extDriver, device)

	extDriver.EXPECT().VkBindImageMemory2KHR(
		device.Handle(),
		loader.Uint32(1),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(
		device loader.VkDevice,
		infoCount loader.Uint32,
		pInfo *khr_bind_memory2_driver.VkBindImageMemoryInfoKHR,
	) (common.VkResult, error) {
		val := reflect.ValueOf(pInfo).Elem()

		require.Equal(t, uint64(1000157001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_BIND_IMAGE_MEMORY_INFO_KHR
		require.Equal(t, image.Handle(), (loader.VkImage)(val.FieldByName("image").UnsafePointer()))
		require.Equal(t, memory.Handle(), (loader.VkDeviceMemory)(val.FieldByName("memory").UnsafePointer()))
		require.Equal(t, uint64(1), val.FieldByName("memoryOffset").Uint())

		next := (*khr_device_group_driver.VkBindImageMemoryDeviceGroupInfoKHR)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()
		require.Equal(t, uint64(1000060014), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_BIND_IMAGE_MEMORY_DEVICE_GROUP_INFO_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(3), val.FieldByName("deviceIndexCount").Uint())

		indices := (*loader.Uint32)(val.FieldByName("pDeviceIndices").UnsafePointer())
		indexSlice := ([]loader.Uint32)(unsafe.Slice(indices, 3))
		indexVal := reflect.ValueOf(indexSlice)

		require.Equal(t, uint64(1), indexVal.Index(0).Uint())
		require.Equal(t, uint64(2), indexVal.Index(1).Uint())
		require.Equal(t, uint64(7), indexVal.Index(2).Uint())

		require.Equal(t, uint64(2), val.FieldByName("splitInstanceBindRegionCount").Uint())

		regions := (*loader.VkRect2D)(val.FieldByName("pSplitInstanceBindRegions").UnsafePointer())
		regionSlice := ([]loader.VkRect2D)(unsafe.Slice(regions, 2))
		regionVal := reflect.ValueOf(regionSlice)

		oneRegion := regionVal.Index(0)
		require.Equal(t, int64(3), oneRegion.FieldByName("offset").FieldByName("x").Int())
		require.Equal(t, int64(5), oneRegion.FieldByName("offset").FieldByName("y").Int())
		require.Equal(t, uint64(7), oneRegion.FieldByName("extent").FieldByName("width").Uint())
		require.Equal(t, uint64(11), oneRegion.FieldByName("extent").FieldByName("height").Uint())

		oneRegion = regionVal.Index(1)
		require.Equal(t, int64(13), oneRegion.FieldByName("offset").FieldByName("x").Int())
		require.Equal(t, int64(17), oneRegion.FieldByName("offset").FieldByName("y").Int())
		require.Equal(t, uint64(19), oneRegion.FieldByName("extent").FieldByName("width").Uint())
		require.Equal(t, uint64(23), oneRegion.FieldByName("extent").FieldByName("height").Uint())

		return core1_0.VKSuccess, nil
	})

	_, err := extension.BindImageMemory2(
		khr_bind_memory2.BindImageMemoryInfo{
			Image:        image,
			Memory:       memory,
			MemoryOffset: 1,

			NextOptions: common.NextOptions{
				khr_device_group.BindImageMemoryDeviceGroupInfo{
					DeviceIndices: []int{1, 2, 7},
					SplitInstanceBindRegions: []core1_0.Rect2D{
						{
							Offset: core1_0.Offset2D{X: 3, Y: 5},
							Extent: core1_0.Extent2D{Width: 7, Height: 11},
						},
						{
							Offset: core1_0.Offset2D{X: 13, Y: 17},
							Extent: core1_0.Extent2D{Width: 19, Height: 23},
						},
					},
				},
			},
		},
	)
	require.NoError(t, err)
}

func TestBindImageMemorySwapchainOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	image := mocks.NewDummyImage(device)
	memory := mocks.NewDummyDeviceMemory(device, 1)
	swapchain := khr_swapchain.NewDummySwapchain(device)

	extDriver := mock_bind_memory2.NewMockLoader(ctrl)
	extension := khr_bind_memory2.CreateExtensionDriverFromLoader(extDriver, device)

	extDriver.EXPECT().VkBindImageMemory2KHR(
		device.Handle(),
		loader.Uint32(1),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(
		device loader.VkDevice,
		infoCount loader.Uint32,
		pInfo *khr_bind_memory2_driver.VkBindImageMemoryInfoKHR,
	) (common.VkResult, error) {
		val := reflect.ValueOf(pInfo).Elem()

		require.Equal(t, uint64(1000157001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_BIND_IMAGE_MEMORY_INFO_KHR
		require.Equal(t, image.Handle(), (loader.VkImage)(val.FieldByName("image").UnsafePointer()))
		require.Equal(t, memory.Handle(), (loader.VkDeviceMemory)(val.FieldByName("memory").UnsafePointer()))
		require.Equal(t, uint64(1), val.FieldByName("memoryOffset").Uint())

		next := (*khr_device_group_driver.VkBindImageMemorySwapchainInfoKHR)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()
		require.Equal(t, uint64(1000060009), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_BIND_IMAGE_MEMORY_SWAPCHAIN_INFO_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, swapchain.Handle(), (khr_swapchain_driver.VkSwapchainKHR)(val.FieldByName("swapchain").UnsafePointer()))
		require.Equal(t, uint64(3), val.FieldByName("imageIndex").Uint())

		return core1_0.VKSuccess, nil
	})

	_, err := extension.BindImageMemory2(
		khr_bind_memory2.BindImageMemoryInfo{
			Image:        image,
			Memory:       memory,
			MemoryOffset: 1,

			NextOptions: common.NextOptions{
				khr_device_group.BindImageMemorySwapchainInfo{
					Swapchain:  swapchain,
					ImageIndex: 3,
				},
			},
		},
	)
	require.NoError(t, err)
}

func TestDeviceGroupBindSparseOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	fence := mocks.NewDummyFence(device)

	coreLoader := mock_driver.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := mocks1_0.InternalDeviceDriver(device, coreLoader)

	semaphore1 := mocks.NewDummySemaphore(device)
	semaphore2 := mocks.NewDummySemaphore(device)
	semaphore3 := mocks.NewDummySemaphore(device)

	queue := mocks.NewDummyQueue(device)

	coreLoader.EXPECT().VkQueueBindSparse(
		queue.Handle(),
		loader.Uint32(1),
		gomock.Not(gomock.Nil()),
		fence.Handle(),
	).DoAndReturn(func(
		queue loader.VkQueue,
		optionCount loader.Uint32,
		pSparseOptions *loader.VkBindSparseInfo,
		fence loader.VkFence,
	) (common.VkResult, error) {
		val := reflect.ValueOf(pSparseOptions).Elem()

		require.Equal(t, uint64(7), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_BIND_SPARSE_INFO
		require.Equal(t, uint64(1), val.FieldByName("waitSemaphoreCount").Uint())
		require.Equal(t, semaphore1.Handle(), loader.VkSemaphore(val.FieldByName("pWaitSemaphores").Elem().UnsafePointer()))
		require.Equal(t, uint64(2), val.FieldByName("signalSemaphoreCount").Uint())

		semaphores := (*loader.VkSemaphore)(val.FieldByName("pSignalSemaphores").UnsafePointer())
		semaphoreSlice := ([]loader.VkSemaphore)(unsafe.Slice(semaphores, 2))
		require.Equal(t, semaphore2.Handle(), semaphoreSlice[0])
		require.Equal(t, semaphore3.Handle(), semaphoreSlice[1])

		next := (*khr_device_group_driver.VkDeviceGroupBindSparseInfoKHR)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000060006), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_GROUP_BIND_SPARSE_INFO_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1), val.FieldByName("resourceDeviceIndex").Uint())
		require.Equal(t, uint64(3), val.FieldByName("memoryDeviceIndex").Uint())

		return core1_0.VKSuccess, nil
	})

	_, err := driver.QueueBindSparse(queue, &fence,
		core1_0.BindSparseInfo{
			WaitSemaphores:   []core1_0.Semaphore{semaphore1},
			SignalSemaphores: []core1_0.Semaphore{semaphore2, semaphore3},
			NextOptions: common.NextOptions{
				khr_device_group.DeviceGroupBindSparseInfo{
					ResourceDeviceIndex: 1,
					MemoryDeviceIndex:   3,
				},
			},
		},
	)
	require.NoError(t, err)
}

func TestImageSwapchainCreateOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	mockImage := mocks.NewDummyImage(device)
	swapchain := khr_swapchain.NewDummySwapchain(device)

	coreLoader := mock_driver.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := mocks1_0.InternalDeviceDriver(device, coreLoader)

	coreLoader.EXPECT().VkCreateImage(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(
		device loader.VkDevice,
		pCreateInfo *loader.VkImageCreateInfo,
		pAllocator *loader.VkAllocationCallbacks,
		pImage *loader.VkImage,
	) (common.VkResult, error) {
		*pImage = mockImage.Handle()

		val := reflect.ValueOf(pCreateInfo).Elem()
		require.Equal(t, uint64(14), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_IMAGE_CREATE_INFO
		require.Equal(t, uint64(1), val.FieldByName("mipLevels").Uint())
		require.Equal(t, uint64(3), val.FieldByName("arrayLayers").Uint())

		next := (*khr_device_group_driver.VkImageSwapchainCreateInfoKHR)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()
		require.Equal(t, uint64(1000060008), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_IMAGE_SWAPCHAIN_CREATE_INFO_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, swapchain.Handle(), (khr_swapchain_driver.VkSwapchainKHR)(val.FieldByName("swapchain").UnsafePointer()))

		return core1_0.VKSuccess, nil
	})

	image, _, err := driver.CreateImage(nil, core1_0.ImageCreateInfo{
		MipLevels:   1,
		ArrayLayers: 3,
		NextOptions: common.NextOptions{
			khr_device_group.ImageSwapchainCreateInfo{
				Swapchain: swapchain,
			},
		},
	})
	require.Equal(t, mockImage.Handle(), image.Handle())
	require.NoError(t, err)
}

func TestDeviceGroupPresentOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})

	extDriver := mock_swapchain.NewMockLoader(ctrl)
	extension := khr_swapchain.CreateExtensionDriverFromLoader(extDriver, device)

	queue := mocks.NewDummyQueue(device)
	swapchain := khr_swapchain.NewDummySwapchain(device)

	extDriver.EXPECT().VkQueuePresentKHR(
		queue.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(
		queue loader.VkQueue,
		pPresentInfo *khr_swapchain_driver.VkPresentInfoKHR,
	) (common.VkResult, error) {
		val := reflect.ValueOf(pPresentInfo).Elem()
		require.Equal(t, uint64(1000001001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PRESENT_INFO_KHR
		require.Equal(t, uint64(1), val.FieldByName("swapchainCount").Uint())
		require.Equal(t, swapchain.Handle(), (khr_swapchain_driver.VkSwapchainKHR)(val.FieldByName("pSwapchains").Elem().UnsafePointer()))
		require.Equal(t, uint64(3), val.FieldByName("pImageIndices").Elem().Uint())

		next := (*khr_device_group_driver.VkDeviceGroupPresentInfoKHR)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000060011), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_GROUP_PRESENT_INFO_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1), val.FieldByName("swapchainCount").Uint())
		require.Equal(t, uint64(7), val.FieldByName("pDeviceMasks").Elem().Uint())
		require.Equal(t, uint64(4), val.FieldByName("mode").Uint()) // VK_DEVICE_GROUP_PRESENT_MODE_SUM_BIT_KHR

		return core1_0.VKSuccess, nil
	})

	_, err := extension.QueuePresent(queue, khr_swapchain.PresentInfo{
		Swapchains:   []khr_swapchain.Swapchain{swapchain},
		ImageIndices: []int{3},
		NextOptions: common.NextOptions{
			khr_device_group.DeviceGroupPresentInfo{
				DeviceMasks: []uint32{7},
				Mode:        khr_device_group.DeviceGroupPresentModeSum,
			},
		},
	})
	require.NoError(t, err)
}

func TestDeviceGroupRenderPassBeginOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreLoader := mock_driver.LoaderForVersion(ctrl, common.Vulkan1_0)
	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	commandPool := mocks.NewDummyCommandPool(device)

	driver := mocks1_0.InternalDeviceDriver(device, coreLoader)

	commandBuffer := mocks.NewDummyCommandBuffer(commandPool, device)
	renderPass := mocks.NewDummyRenderPass(device)
	framebuffer := mocks.NewDummyFramebuffer(device)

	coreLoader.EXPECT().VkCmdBeginRenderPass(
		commandBuffer.Handle(),
		gomock.Not(gomock.Nil()),
		loader.VkSubpassContents(1), // VK_SUBPASS_CONTENTS_SECONDARY_COMMAND_BUFFERS
	).DoAndReturn(func(
		commandBuffer loader.VkCommandBuffer,
		pRenderPassBegin *loader.VkRenderPassBeginInfo,
		contents loader.VkSubpassContents,
	) {
		val := reflect.ValueOf(pRenderPassBegin).Elem()
		require.Equal(t, uint64(43), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_RENDER_PASS_BEGIN_INFO
		require.Equal(t, renderPass.Handle(), (loader.VkRenderPass)(val.FieldByName("renderPass").UnsafePointer()))
		require.Equal(t, framebuffer.Handle(), (loader.VkFramebuffer)(val.FieldByName("framebuffer").UnsafePointer()))

		next := (*khr_device_group_driver.VkDeviceGroupRenderPassBeginInfoKHR)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000060003), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_GROUP_RENDER_PASS_BEGIN_INFO_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(7), val.FieldByName("deviceMask").Uint())
		require.Equal(t, uint64(2), val.FieldByName("deviceRenderAreaCount").Uint())

		areas := (*loader.VkRect2D)(val.FieldByName("pDeviceRenderAreas").UnsafePointer())
		areaSlice := ([]loader.VkRect2D)(unsafe.Slice(areas, 2))
		val = reflect.ValueOf(areaSlice)

		oneArea := val.Index(0)
		require.Equal(t, int64(1), oneArea.FieldByName("offset").FieldByName("x").Int())
		require.Equal(t, int64(3), oneArea.FieldByName("offset").FieldByName("y").Int())
		require.Equal(t, uint64(5), oneArea.FieldByName("extent").FieldByName("width").Uint())
		require.Equal(t, uint64(7), oneArea.FieldByName("extent").FieldByName("height").Uint())

		oneArea = val.Index(1)
		require.Equal(t, int64(11), oneArea.FieldByName("offset").FieldByName("x").Int())
		require.Equal(t, int64(13), oneArea.FieldByName("offset").FieldByName("y").Int())
		require.Equal(t, uint64(17), oneArea.FieldByName("extent").FieldByName("width").Uint())
		require.Equal(t, uint64(19), oneArea.FieldByName("extent").FieldByName("height").Uint())
	})

	err := driver.CmdBeginRenderPass(
		commandBuffer,
		core1_0.SubpassContentsSecondaryCommandBuffers,
		core1_0.RenderPassBeginInfo{
			RenderPass:  renderPass,
			Framebuffer: framebuffer,
			NextOptions: common.NextOptions{
				khr_device_group.DeviceGroupRenderPassBeginInfo{
					DeviceMask: 7,
					DeviceRenderAreas: []core1_0.Rect2D{
						{
							Offset: core1_0.Offset2D{X: 1, Y: 3},
							Extent: core1_0.Extent2D{Width: 5, Height: 7},
						},
						{
							Offset: core1_0.Offset2D{X: 11, Y: 13},
							Extent: core1_0.Extent2D{Width: 17, Height: 19},
						},
					},
				},
			},
		},
	)
	require.NoError(t, err)
}

func TestDeviceGroupSubmitOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreLoader := mock_driver.LoaderForVersion(ctrl, common.Vulkan1_0)
	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	driver := mocks1_0.InternalDeviceDriver(device, coreLoader)
	commandPool := mocks.NewDummyCommandPool(device)
	fence := mocks.NewDummyFence(device)
	commandBuffer := mocks.NewDummyCommandBuffer(commandPool, device)

	semaphore1 := mocks.NewDummySemaphore(device)
	semaphore2 := mocks.NewDummySemaphore(device)
	semaphore3 := mocks.NewDummySemaphore(device)

	queue := mocks.NewDummyQueue(device)

	coreLoader.EXPECT().VkQueueSubmit(
		queue.Handle(),
		loader.Uint32(1),
		gomock.Not(gomock.Nil()),
		fence.Handle(),
	).DoAndReturn(func(queue loader.VkQueue, submitCount loader.Uint32, pSubmits *loader.VkSubmitInfo, fence loader.VkFence) (common.VkResult, error) {
		val := reflect.ValueOf(pSubmits).Elem()

		require.Equal(t, uint64(4), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SUBMIT_INFO
		require.Equal(t, uint64(1), val.FieldByName("waitSemaphoreCount").Uint())
		require.Equal(t, uint64(1), val.FieldByName("commandBufferCount").Uint())
		require.Equal(t, uint64(2), val.FieldByName("signalSemaphoreCount").Uint())

		require.Equal(t, semaphore1.Handle(), loader.VkSemaphore(val.FieldByName("pWaitSemaphores").Elem().UnsafePointer()))
		require.Equal(t, uint64(0x00002000), val.FieldByName("pWaitDstStageMask").Elem().Uint()) // VK_PIPELINE_STAGE_BOTTOM_OF_PIPE_BIT
		require.Equal(t, commandBuffer.Handle(), loader.VkCommandBuffer(val.FieldByName("pCommandBuffers").Elem().UnsafePointer()))

		semaphores := (*loader.VkSemaphore)(val.FieldByName("pSignalSemaphores").UnsafePointer())
		semaphoreSlice := ([]loader.VkSemaphore)(unsafe.Slice(semaphores, 2))
		require.Equal(t, semaphore2.Handle(), semaphoreSlice[0])
		require.Equal(t, semaphore3.Handle(), semaphoreSlice[1])

		next := (*khr_device_group_driver.VkDeviceGroupSubmitInfoKHR)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000060005), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_GROUP_SUBMIT_INFO_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1), val.FieldByName("waitSemaphoreCount").Uint())
		require.Equal(t, uint64(1), val.FieldByName("commandBufferCount").Uint())
		require.Equal(t, uint64(2), val.FieldByName("signalSemaphoreCount").Uint())

		require.Equal(t, uint64(1), val.FieldByName("pWaitSemaphoreDeviceIndices").Elem().Uint())
		require.Equal(t, uint64(2), val.FieldByName("pCommandBufferDeviceMasks").Elem().Uint())

		indices := (*loader.Uint32)(val.FieldByName("pSignalSemaphoreDeviceIndices").UnsafePointer())
		indexSlice := ([]loader.Uint32)(unsafe.Slice(indices, 2))
		require.Equal(t, []loader.Uint32{3, 5}, indexSlice)

		return core1_0.VKSuccess, nil
	})

	_, err := driver.QueueSubmit(queue, &fence,
		core1_0.SubmitInfo{
			CommandBuffers:   []core1_0.CommandBuffer{commandBuffer},
			WaitSemaphores:   []core1_0.Semaphore{semaphore1},
			SignalSemaphores: []core1_0.Semaphore{semaphore2, semaphore3},
			WaitDstStageMask: []core1_0.PipelineStageFlags{core1_0.PipelineStageBottomOfPipe},

			NextOptions: common.NextOptions{
				khr_device_group.DeviceGroupSubmitInfo{
					WaitSemaphoreDeviceIndices:   []int{1},
					CommandBufferDeviceMasks:     []uint32{2},
					SignalSemaphoreDeviceIndices: []int{3, 5},
				},
			},
		},
	)
	require.NoError(t, err)
}

func TestDeviceGroupSwapchainCreateOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	instance := mocks.NewDummyInstance(common.Vulkan1_0, []string{})
	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	surface := mock_surface.NewDummySurface(instance)
	mockSwapchain := khr_swapchain.NewDummySwapchain(device)

	extDriver := mock_swapchain.NewMockLoader(ctrl)
	extension := khr_swapchain.CreateExtensionDriverFromLoader(extDriver, device)

	extDriver.EXPECT().VkCreateSwapchainKHR(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(device loader.VkDevice,
			pCreateInfo *khr_swapchain_driver.VkSwapchainCreateInfoKHR,
			pAllocator *loader.VkAllocationCallbacks,
			pSwapchain *khr_swapchain_driver.VkSwapchainKHR,
		) (common.VkResult, error) {
			*pSwapchain = mockSwapchain.Handle()

			val := reflect.ValueOf(pCreateInfo).Elem()
			require.Equal(t, uint64(1000001000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SWAPCHAIN_CREATE_INFO_KHR
			require.Equal(t, surface.Handle(), khr_surface_driver.VkSurfaceKHR(val.FieldByName("surface").UnsafePointer()))

			next := (*khr_device_group_driver.VkDeviceGroupSwapchainCreateInfoKHR)(val.FieldByName("pNext").UnsafePointer())
			val = reflect.ValueOf(next).Elem()

			require.Equal(t, uint64(1000060012), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_GROUP_SWAPCHAIN_CREATE_INFO_KHR
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(1), val.FieldByName("modes").Uint()) // VK_DEVICE_GROUP_PRESENT_MODE_LOCAL_BIT_KHR

			return core1_0.VKSuccess, nil
		})

	swapchain, _, err := extension.CreateSwapchain(
		nil,
		khr_swapchain.SwapchainCreateInfo{
			Surface: surface,
			NextOptions: common.NextOptions{
				khr_device_group.DeviceGroupSwapchainCreateInfo{
					Modes: khr_device_group.DeviceGroupPresentModeLocal,
				},
			},
		})
	require.NoError(t, err)
	require.Equal(t, mockSwapchain.Handle(), swapchain.Handle())
}
