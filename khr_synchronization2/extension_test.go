package khr_synchronization2_test

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/core1_1"
	"github.com/vkngwrapper/core/v3/core1_2"
	"github.com/vkngwrapper/core/v3/loader"
	mock_loader "github.com/vkngwrapper/core/v3/loader/mocks"
	"github.com/vkngwrapper/core/v3/mocks"
	"github.com/vkngwrapper/core/v3/mocks/mocks1_1"
	"github.com/vkngwrapper/extensions/v3/khr_synchronization2"
	khr_synchronization2_loader "github.com/vkngwrapper/extensions/v3/khr_synchronization2/loader"
	mock_synchronization2 "github.com/vkngwrapper/extensions/v3/khr_synchronization2/mocks"
	"go.uber.org/mock/gomock"
)

func TestPhysicalDeviceSynchronization2Features_OutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	instance := mocks.NewDummyInstance(common.Vulkan1_1, []string{})
	mockLoader := mock_loader.NewMockLoader(ctrl)
	driver := mocks1_1.InternalCoreInstanceDriver(instance, mockLoader)

	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_1)

	mockLoader.EXPECT().VkGetPhysicalDeviceFeatures2(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice loader.VkPhysicalDevice,
		pFeatures *loader.VkPhysicalDeviceFeatures2) {

		info := reflect.ValueOf(pFeatures).Elem()
		require.Equal(t, uint64(1000059000), info.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FEATURES_2
		require.False(t, info.FieldByName("pNext").IsNil())

		pInnerFeatures := (*khr_synchronization2_loader.VkPhysicalDeviceSynchronization2Features)(info.FieldByName("pNext").UnsafePointer())
		features := reflect.ValueOf(pInnerFeatures).Elem()
		require.Equal(t, uint64(0x3b9f9497), features.FieldByName("sType").Uint())
		require.True(t, features.FieldByName("pNext").IsNil())
		*(*uint64)(unsafe.Pointer(features.FieldByName("synchronization2").UnsafeAddr())) = uint64(1)
	})

	var baseFeatures core1_1.PhysicalDeviceFeatures2
	var features khr_synchronization2.PhysicalDeviceSynchronization2Features

	baseFeatures.NextOutData.Next = &features

	err := driver.GetPhysicalDeviceFeatures2(physicalDevice, &baseFeatures)
	require.NoError(t, err)

	require.Equal(t, khr_synchronization2.PhysicalDeviceSynchronization2Features{
		Synchronization2: true,
	}, features)
}

func TestPhysicalDeviceSynchronization2Features_Options(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	instance := mocks.NewDummyInstance(common.Vulkan1_1, []string{})
	mockLoader := mock_loader.NewMockLoader(ctrl)
	driver := mocks1_1.InternalCoreInstanceDriver(instance, mockLoader)

	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_1)
	expectedDevice := mocks.NewDummyDevice(common.Vulkan1_1, []string{})

	mockLoader.EXPECT().VkCreateDevice(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice loader.VkPhysicalDevice,
		pCreateInfo *loader.VkDeviceCreateInfo,
		pAllocator *loader.VkAllocationCallbacks,
		pDevice *loader.VkDevice) (common.VkResult, error) {

		info := reflect.ValueOf(pCreateInfo).Elem()
		require.Equal(t, uint64(3), info.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_CREATE_INFO
		require.False(t, info.FieldByName("pNext").IsNil())

		pInnerFeatures := (*khr_synchronization2_loader.VkPhysicalDeviceSynchronization2Features)(info.FieldByName("pNext").UnsafePointer())
		features := reflect.ValueOf(pInnerFeatures).Elem()

		require.Equal(t, uint64(0x3b9f9497), features.FieldByName("sType").Uint())
		require.True(t, features.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1), features.FieldByName("synchronization2").Uint())

		*pDevice = expectedDevice.Handle()

		return core1_0.VKSuccess, nil
	})

	device, res, err := driver.CreateDevice(physicalDevice, nil, core1_0.DeviceCreateInfo{
		QueueCreateInfos: []core1_0.DeviceQueueCreateInfo{
			{
				QueueFamilyIndex: 0,
				QueuePriorities:  []float32{1},
			},
		},
		NextOptions: common.NextOptions{khr_synchronization2.PhysicalDeviceSynchronization2Features{
			Synchronization2: true,
		}},
	})
	require.NoError(t, err)
	require.Equal(t, core1_0.VKSuccess, res)
	require.NotNil(t, device)
	require.Equal(t, expectedDevice.Handle(), device.Device().Handle())
}

func TestVulkanExtension_QueueSubmit2(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_synchronization2.NewMockLoader(ctrl)
	extension := khr_synchronization2.CreateExtensionDriverFromLoader(extDriver)

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	queue := mocks.NewDummyQueue(device)
	fence := mocks.NewDummyFence(device)
	semaphore1 := mocks.NewDummySemaphore(device)
	semaphore2 := mocks.NewDummySemaphore(device)
	semaphore3 := mocks.NewDummySemaphore(device)
	commandPool := mocks.NewDummyCommandPool(device)
	commandBuffer := mocks.NewDummyCommandBuffer(commandPool, device)

	extDriver.EXPECT().VkQueueSubmit2(queue.Handle(), loader.Uint32(1), gomock.Not(gomock.Nil()), fence.Handle()).DoAndReturn(
		func(queue loader.VkQueue, submitCount loader.Uint32, pSubmits *khr_synchronization2_loader.VkSubmitInfo2, fence loader.VkFence) (common.VkResult, error) {
			submitSlice := ([]khr_synchronization2_loader.VkSubmitInfo2)(unsafe.Slice(pSubmits, 1))
			submitInfo := reflect.ValueOf(submitSlice).Index(0)

			require.Equal(t, uint64(0x3b9f9494), submitInfo.FieldByName("sType").Uint())
			require.True(t, submitInfo.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(1), submitInfo.FieldByName("flags").Uint())
			require.Equal(t, uint64(2), submitInfo.FieldByName("waitSemaphoreInfoCount").Uint())
			require.Equal(t, uint64(1), submitInfo.FieldByName("commandBufferInfoCount").Uint())
			require.Equal(t, uint64(1), submitInfo.FieldByName("signalSemaphoreInfoCount").Uint())

			waitSemaphorePtr := (*khr_synchronization2_loader.VkSemaphoreSubmitInfo)(submitInfo.FieldByName("pWaitSemaphoreInfos").UnsafePointer())
			waitSemaphoreSlice := unsafe.Slice(waitSemaphorePtr, 2)

			semaphoreInfo := reflect.ValueOf(waitSemaphoreSlice).Index(0)
			require.Equal(t, uint64(0x3b9f9495), semaphoreInfo.FieldByName("sType").Uint())
			require.True(t, semaphoreInfo.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(3), semaphoreInfo.FieldByName("value").Uint())
			require.Equal(t, uint64(0x400000000), semaphoreInfo.FieldByName("stageMask").Uint())
			require.Equal(t, uint64(7), semaphoreInfo.FieldByName("deviceIndex").Uint())
			require.Equal(t, semaphore1.Handle(), loader.VkSemaphore(unsafe.Pointer(semaphoreInfo.FieldByName("semaphore").UnsafePointer())))

			semaphoreInfo = reflect.ValueOf(waitSemaphoreSlice).Index(1)
			require.Equal(t, uint64(0x3b9f9495), semaphoreInfo.FieldByName("sType").Uint())
			require.True(t, semaphoreInfo.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(11), semaphoreInfo.FieldByName("value").Uint())
			require.Equal(t, uint64(0x00002000), semaphoreInfo.FieldByName("stageMask").Uint())
			require.Equal(t, uint64(13), semaphoreInfo.FieldByName("deviceIndex").Uint())
			require.Equal(t, semaphore2.Handle(), loader.VkSemaphore(unsafe.Pointer(semaphoreInfo.FieldByName("semaphore").UnsafePointer())))

			commandBufferPtr := (*khr_synchronization2_loader.VkCommandBufferSubmitInfo)(submitInfo.FieldByName("pCommandBufferInfos").UnsafePointer())
			commandBufferSlice := unsafe.Slice(commandBufferPtr, 1)

			commandBufferInfo := reflect.ValueOf(commandBufferSlice).Index(0)
			require.Equal(t, uint64(0x3b9f9496), commandBufferInfo.FieldByName("sType").Uint())
			require.True(t, commandBufferInfo.FieldByName("pNext").IsNil())
			require.Equal(t, commandBuffer.Handle(), loader.VkCommandBuffer(unsafe.Pointer(commandBufferInfo.FieldByName("commandBuffer").UnsafePointer())))
			require.Equal(t, uint64(17), commandBufferInfo.FieldByName("deviceMask").Uint())

			signalSemaphorePtr := (*khr_synchronization2_loader.VkSemaphoreSubmitInfo)(submitInfo.FieldByName("pSignalSemaphoreInfos").UnsafePointer())
			signalSemaphoreSlice := unsafe.Slice(signalSemaphorePtr, 1)

			semaphoreInfo = reflect.ValueOf(signalSemaphoreSlice).Index(0)
			require.Equal(t, uint64(0x3b9f9495), semaphoreInfo.FieldByName("sType").Uint())
			require.True(t, semaphoreInfo.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(19), semaphoreInfo.FieldByName("value").Uint())
			require.Equal(t, uint64(0x800000000), semaphoreInfo.FieldByName("stageMask").Uint())
			require.Equal(t, uint64(23), semaphoreInfo.FieldByName("deviceIndex").Uint())
			require.Equal(t, semaphore3.Handle(), loader.VkSemaphore(unsafe.Pointer(semaphoreInfo.FieldByName("semaphore").UnsafePointer())))

			return core1_0.VKSuccess, nil
		})

	res, err := extension.QueueSubmit2(queue, []khr_synchronization2.SubmitInfo2{
		{
			Flags: khr_synchronization2.SubmitProtected,
			WaitSemaphoreInfos: []khr_synchronization2.SemaphoreSubmitInfo{
				{
					Semaphore:   semaphore1,
					Value:       3,
					StageMask:   khr_synchronization2.PipelineStage2Blit,
					DeviceIndex: 7,
				},
				{
					Semaphore:   semaphore2,
					Value:       11,
					StageMask:   khr_synchronization2.PipelineStage2BottomOfPipe,
					DeviceIndex: 13,
				},
			},
			CommandBufferInfos: []khr_synchronization2.CommandBufferSubmitInfo{
				{
					CommandBuffer: commandBuffer,
					DeviceMask:    17,
				},
			},
			SignalSemaphoreInfos: []khr_synchronization2.SemaphoreSubmitInfo{
				{
					Semaphore:   semaphore3,
					Value:       19,
					StageMask:   khr_synchronization2.PipelineStage2Clear,
					DeviceIndex: 23,
				},
			},
		},
	}, fence)
	require.NoError(t, err)
	require.Equal(t, core1_0.VKSuccess, res)
}

func TestVulkanExtension_CmdWriteTimestamp2(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_synchronization2.NewMockLoader(ctrl)
	extension := khr_synchronization2.CreateExtensionDriverFromLoader(extDriver)

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	commandPool := mocks.NewDummyCommandPool(device)
	commandBuffer := mocks.NewDummyCommandBuffer(commandPool, device)
	queryPool := mocks.NewDummyQueryPool(device)

	extDriver.EXPECT().VkCmdWriteTimestamp2KHR(commandBuffer.Handle(), khr_synchronization2_loader.VkPipelineStageFlags2(0x800000000), queryPool.Handle(), loader.Uint32(7))

	extension.CmdWriteTimestamp2(commandBuffer, khr_synchronization2.PipelineStage2Clear, queryPool, 7)
}

func TestVulkanExtension_CmdWaitEvents2(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_synchronization2.NewMockLoader(ctrl)
	extension := khr_synchronization2.CreateExtensionDriverFromLoader(extDriver)

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	commandPool := mocks.NewDummyCommandPool(device)
	commandBuffer := mocks.NewDummyCommandBuffer(commandPool, device)
	event1 := mocks.NewDummyEvent(device)
	event2 := mocks.NewDummyEvent(device)
	buffer := mocks.NewDummyBuffer(device)

	extDriver.EXPECT().VkCmdWaitEvents2KHR(commandBuffer.Handle(), loader.Uint32(2), gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).DoAndReturn(
		func(commandBuffer loader.VkCommandBuffer, eventCount loader.Uint32, pEvents *loader.VkEvent, pDependencyInfos *khr_synchronization2_loader.VkDependencyInfo) {
			eventsSlice := ([]loader.VkEvent)(unsafe.Slice(pEvents, 2))
			require.Equal(t, event1.Handle(), eventsSlice[0])
			require.Equal(t, event2.Handle(), eventsSlice[1])

			dependencySlice := ([]khr_synchronization2_loader.VkDependencyInfo)(unsafe.Slice(pDependencyInfos, 2))

			depVal := reflect.ValueOf(dependencySlice).Index(0)
			require.Equal(t, uint64(0x3b9f9493), depVal.FieldByName("sType").Uint())
			require.True(t, depVal.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(4), depVal.FieldByName("dependencyFlags").Uint())
			require.Equal(t, uint64(1), depVal.FieldByName("memoryBarrierCount").Uint())
			require.Equal(t, uint64(0), depVal.FieldByName("bufferMemoryBarrierCount").Uint())
			require.True(t, depVal.FieldByName("pBufferMemoryBarriers").IsNil())
			require.Equal(t, uint64(0), depVal.FieldByName("imageMemoryBarrierCount").Uint())
			require.True(t, depVal.FieldByName("pImageMemoryBarriers").IsNil())

			barrierPtr := (*khr_synchronization2_loader.VkMemoryBarrier2)(depVal.FieldByName("pMemoryBarriers").UnsafePointer())
			barrierSlice := unsafe.Slice(barrierPtr, 1)

			barrierVal := reflect.ValueOf(barrierSlice).Index(0)
			require.Equal(t, uint64(0x3b9f9490), barrierVal.FieldByName("sType").Uint())
			require.True(t, barrierVal.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0x02000000), barrierVal.FieldByName("srcStageMask").Uint())
			require.Equal(t, uint64(0x00200000), barrierVal.FieldByName("srcAccessMask").Uint())
			require.Equal(t, uint64(0x400000000), barrierVal.FieldByName("dstStageMask").Uint())
			require.Equal(t, uint64(0x00400000), barrierVal.FieldByName("dstAccessMask").Uint())

			depVal = reflect.ValueOf(dependencySlice).Index(1)
			require.Equal(t, uint64(0x3b9f9493), depVal.FieldByName("sType").Uint())
			require.True(t, depVal.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0), depVal.FieldByName("dependencyFlags").Uint())
			require.Equal(t, uint64(0), depVal.FieldByName("memoryBarrierCount").Uint())
			require.True(t, depVal.FieldByName("pMemoryBarriers").IsNil())
			require.Equal(t, uint64(1), depVal.FieldByName("bufferMemoryBarrierCount").Uint())
			require.Equal(t, uint64(0), depVal.FieldByName("imageMemoryBarrierCount").Uint())
			require.True(t, depVal.FieldByName("pImageMemoryBarriers").IsNil())

			bufferBarrierPtr := (*khr_synchronization2_loader.VkBufferMemoryBarrier2)(depVal.FieldByName("pBufferMemoryBarriers").UnsafePointer())
			bufferBarrierSlice := unsafe.Slice(bufferBarrierPtr, 1)

			barrierVal = reflect.ValueOf(bufferBarrierSlice).Index(0)
			require.Equal(t, uint64(0x3b9f9491), barrierVal.FieldByName("sType").Uint())
			require.True(t, barrierVal.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0x00002000), barrierVal.FieldByName("srcStageMask").Uint())
			require.Equal(t, uint64(0x00000080), barrierVal.FieldByName("srcAccessMask").Uint())
			require.Equal(t, uint64(0x800000000), barrierVal.FieldByName("dstStageMask").Uint())
			require.Equal(t, uint64(0x00080000), barrierVal.FieldByName("dstAccessMask").Uint())
			require.Equal(t, uint64(3), barrierVal.FieldByName("srcQueueFamilyIndex").Uint())
			require.Equal(t, uint64(7), barrierVal.FieldByName("dstQueueFamilyIndex").Uint())
			require.Equal(t, uint64(11), barrierVal.FieldByName("offset").Uint())
			require.Equal(t, uint64(13), barrierVal.FieldByName("size").Uint())
			require.Equal(t, buffer.Handle(), (loader.VkBuffer)(barrierVal.FieldByName("buffer").UnsafePointer()))
		})

	err := extension.CmdWaitEvents2(
		commandBuffer,
		[]core1_0.Event{event1, event2},
		[]khr_synchronization2.DependencyInfo{
			{
				DependencyFlags: core1_1.DependencyDeviceGroup,
				MemoryBarriers: []khr_synchronization2.MemoryBarrier2{
					{
						SrcStageMask:  khr_synchronization2.PipelineStage2AccelerationStructureBuild,
						SrcAccessMask: khr_synchronization2.Access2AccelerationStructureRead,
						DstStageMask:  khr_synchronization2.PipelineStage2Blit,
						DstAccessMask: khr_synchronization2.Access2AccelerationStructureWrite,
					},
				},
			},
			{
				DependencyFlags: 0,
				BufferMemoryBarriers: []khr_synchronization2.BufferMemoryBarrier2{
					{
						SrcStageMask:        khr_synchronization2.PipelineStage2BottomOfPipe,
						SrcAccessMask:       khr_synchronization2.Access2ColorAttachmentRead,
						DstStageMask:        khr_synchronization2.PipelineStage2Clear,
						DstAccessMask:       khr_synchronization2.Access2ColorAttachmentReadNoncoherent,
						SrcQueueFamilyIndex: 3,
						DstQueueFamilyIndex: 7,
						Buffer:              buffer,
						Offset:              11,
						Size:                13,
					},
				},
			},
		},
	)
	require.NoError(t, err)
}

func TestVulkanExtension_CmdSetEvent2(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_synchronization2.NewMockLoader(ctrl)
	extension := khr_synchronization2.CreateExtensionDriverFromLoader(extDriver)

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	commandPool := mocks.NewDummyCommandPool(device)
	commandBuffer := mocks.NewDummyCommandBuffer(commandPool, device)
	event := mocks.NewDummyEvent(device)

	buffer := mocks.NewDummyBuffer(device)
	image1 := mocks.NewDummyImage(device)
	image2 := mocks.NewDummyImage(device)

	extDriver.EXPECT().VkCmdSetEvent2KHR(commandBuffer.Handle(), event.Handle(), gomock.Not(gomock.Nil())).DoAndReturn(
		func(commandBuffer loader.VkCommandBuffer, event loader.VkEvent, pDependencyInfo *khr_synchronization2_loader.VkDependencyInfo) {
			depVal := reflect.ValueOf(pDependencyInfo).Elem()

			require.Equal(t, uint64(0x3b9f9493), depVal.FieldByName("sType").Uint())
			require.True(t, depVal.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(2), depVal.FieldByName("dependencyFlags").Uint())
			require.Equal(t, uint64(1), depVal.FieldByName("memoryBarrierCount").Uint())
			require.Equal(t, uint64(1), depVal.FieldByName("bufferMemoryBarrierCount").Uint())
			require.Equal(t, uint64(2), depVal.FieldByName("imageMemoryBarrierCount").Uint())

			barrierPtr := (*khr_synchronization2_loader.VkMemoryBarrier2)(depVal.FieldByName("pMemoryBarriers").UnsafePointer())
			barrierSlice := unsafe.Slice(barrierPtr, 1)

			barrierVal := reflect.ValueOf(barrierSlice).Index(0)
			require.Equal(t, uint64(0x3b9f9490), barrierVal.FieldByName("sType").Uint())
			require.True(t, barrierVal.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0x00040000), barrierVal.FieldByName("srcStageMask").Uint())
			require.Equal(t, uint64(0x00000100), barrierVal.FieldByName("srcAccessMask").Uint())
			require.Equal(t, uint64(0x100000000), barrierVal.FieldByName("dstStageMask").Uint())
			require.Equal(t, uint64(0x00800000), barrierVal.FieldByName("dstAccessMask").Uint())

			bufferBarrierPtr := (*khr_synchronization2_loader.VkBufferMemoryBarrier2)(depVal.FieldByName("pBufferMemoryBarriers").UnsafePointer())
			bufferBarrierSlice := unsafe.Slice(bufferBarrierPtr, 1)

			barrierVal = reflect.ValueOf(bufferBarrierSlice).Index(0)
			require.Equal(t, uint64(0x3b9f9491), barrierVal.FieldByName("sType").Uint())
			require.True(t, barrierVal.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0x800000000), barrierVal.FieldByName("srcStageMask").Uint())
			require.Equal(t, uint64(0x00000080), barrierVal.FieldByName("srcAccessMask").Uint())
			require.Equal(t, uint64(0x00000400), barrierVal.FieldByName("dstStageMask").Uint())
			require.Equal(t, uint64(0x00020000), barrierVal.FieldByName("dstAccessMask").Uint())
			require.Equal(t, uint64(53), barrierVal.FieldByName("srcQueueFamilyIndex").Uint())
			require.Equal(t, uint64(59), barrierVal.FieldByName("dstQueueFamilyIndex").Uint())
			require.Equal(t, uint64(61), barrierVal.FieldByName("offset").Uint())
			require.Equal(t, uint64(67), barrierVal.FieldByName("size").Uint())
			require.Equal(t, buffer.Handle(), (loader.VkBuffer)(barrierVal.FieldByName("buffer").UnsafePointer()))

			imageBarrierPtr := (*khr_synchronization2_loader.VkImageMemoryBarrier2)(depVal.FieldByName("pImageMemoryBarriers").UnsafePointer())
			imageBarrierSlice := unsafe.Slice(imageBarrierPtr, 2)

			barrierVal = reflect.ValueOf(imageBarrierSlice).Index(0)
			require.Equal(t, uint64(0x3b9f9492), barrierVal.FieldByName("sType").Uint())
			require.True(t, barrierVal.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0x00000008), barrierVal.FieldByName("srcStageMask").Uint())
			require.Equal(t, uint64(0x00002000), barrierVal.FieldByName("srcAccessMask").Uint())
			require.Equal(t, uint64(0x400000000), barrierVal.FieldByName("dstStageMask").Uint())
			require.Equal(t, uint64(0x00400000), barrierVal.FieldByName("dstAccessMask").Uint())
			require.Equal(t, uint64(0x3b9c9308), barrierVal.FieldByName("oldLayout").Uint())
			require.Equal(t, uint64(0x3b9e7768), barrierVal.FieldByName("newLayout").Uint())
			require.Equal(t, uint64(7), barrierVal.FieldByName("srcQueueFamilyIndex").Uint())
			require.Equal(t, uint64(11), barrierVal.FieldByName("dstQueueFamilyIndex").Uint())
			require.Equal(t, uint64(0x10), barrierVal.FieldByName("subresourceRange").FieldByName("aspectMask").Uint())
			require.Equal(t, uint64(13), barrierVal.FieldByName("subresourceRange").FieldByName("baseMipLevel").Uint())
			require.Equal(t, uint64(17), barrierVal.FieldByName("subresourceRange").FieldByName("levelCount").Uint())
			require.Equal(t, uint64(19), barrierVal.FieldByName("subresourceRange").FieldByName("baseArrayLayer").Uint())
			require.Equal(t, uint64(23), barrierVal.FieldByName("subresourceRange").FieldByName("layerCount").Uint())
			require.Equal(t, image1.Handle(), (loader.VkImage)(barrierVal.FieldByName("image").UnsafePointer()))

			barrierVal = reflect.ValueOf(imageBarrierSlice).Index(1)
			require.Equal(t, uint64(0x3b9f9492), barrierVal.FieldByName("sType").Uint())
			require.True(t, barrierVal.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0x00000800), barrierVal.FieldByName("srcStageMask").Uint())
			require.Equal(t, uint64(0x01000000), barrierVal.FieldByName("srcAccessMask").Uint())
			require.Equal(t, uint64(0x00002000), barrierVal.FieldByName("dstStageMask").Uint())
			require.Equal(t, uint64(0x00200000), barrierVal.FieldByName("dstAccessMask").Uint())
			require.Equal(t, uint64(0x3b9e7769), barrierVal.FieldByName("oldLayout").Uint())
			require.Equal(t, uint64(0x3b9c9309), barrierVal.FieldByName("newLayout").Uint())
			require.Equal(t, uint64(29), barrierVal.FieldByName("srcQueueFamilyIndex").Uint())
			require.Equal(t, uint64(31), barrierVal.FieldByName("dstQueueFamilyIndex").Uint())
			require.Equal(t, uint64(2), barrierVal.FieldByName("subresourceRange").FieldByName("aspectMask").Uint())
			require.Equal(t, uint64(37), barrierVal.FieldByName("subresourceRange").FieldByName("baseMipLevel").Uint())
			require.Equal(t, uint64(41), barrierVal.FieldByName("subresourceRange").FieldByName("levelCount").Uint())
			require.Equal(t, uint64(43), barrierVal.FieldByName("subresourceRange").FieldByName("baseArrayLayer").Uint())
			require.Equal(t, uint64(47), barrierVal.FieldByName("subresourceRange").FieldByName("layerCount").Uint())
			require.Equal(t, image2.Handle(), (loader.VkImage)(barrierVal.FieldByName("image").UnsafePointer()))

		})

	err := extension.CmdSetEvent2(commandBuffer, event, khr_synchronization2.DependencyInfo{
		DependencyFlags: core1_1.DependencyViewLocal,
		ImageMemoryBarriers: []khr_synchronization2.ImageMemoryBarrier2{
			{
				SrcStageMask:        khr_synchronization2.PipelineStage2VertexShader,
				SrcAccessMask:       khr_synchronization2.Access2HostRead,
				DstStageMask:        khr_synchronization2.PipelineStage2Blit,
				DstAccessMask:       khr_synchronization2.Access2AccelerationStructureWrite,
				OldLayout:           core1_1.ImageLayoutDepthReadOnlyStencilAttachmentOptimal,
				NewLayout:           core1_2.ImageLayoutDepthAttachmentOptimal,
				SrcQueueFamilyIndex: 7,
				DstQueueFamilyIndex: 11,
				Image:               image1,
				SubresourceRange: core1_0.ImageSubresourceRange{
					AspectMask:     core1_1.ImageAspectPlane0,
					BaseMipLevel:   13,
					LevelCount:     17,
					BaseArrayLayer: 19,
					LayerCount:     23,
				},
			},
			{
				SrcStageMask:        khr_synchronization2.PipelineStage2ComputeShader,
				SrcAccessMask:       khr_synchronization2.Access2FragmentDensityMapRead,
				DstStageMask:        khr_synchronization2.PipelineStage2BottomOfPipe,
				DstAccessMask:       khr_synchronization2.Access2AccelerationStructureRead,
				OldLayout:           core1_2.ImageLayoutDepthReadOnlyOptimal,
				NewLayout:           core1_1.ImageLayoutDepthAttachmentStencilReadOnlyOptimal,
				SrcQueueFamilyIndex: 29,
				DstQueueFamilyIndex: 31,
				Image:               image2,
				SubresourceRange: core1_0.ImageSubresourceRange{
					AspectMask:     core1_0.ImageAspectDepth,
					BaseMipLevel:   37,
					LevelCount:     41,
					BaseArrayLayer: 43,
					LayerCount:     47,
				},
			},
		},
		BufferMemoryBarriers: []khr_synchronization2.BufferMemoryBarrier2{
			{
				SrcStageMask:        khr_synchronization2.PipelineStage2Clear,
				SrcAccessMask:       khr_synchronization2.Access2ColorAttachmentRead,
				DstStageMask:        khr_synchronization2.PipelineStage2ColorAttachmentOutput,
				DstAccessMask:       khr_synchronization2.Access2CommandPreprocessRead,
				SrcQueueFamilyIndex: 53,
				DstQueueFamilyIndex: 59,
				Buffer:              buffer,
				Offset:              61,
				Size:                67,
			},
		},
		MemoryBarriers: []khr_synchronization2.MemoryBarrier2{
			{
				SrcStageMask:  khr_synchronization2.PipelineStage2ConditionalRendering,
				SrcAccessMask: khr_synchronization2.Access2ColorAttachmentWrite,
				DstStageMask:  khr_synchronization2.PipelineStage2Copy,
				DstAccessMask: khr_synchronization2.Access2FragmentShadingRateAttachmentRead,
			},
		},
	})

	require.NoError(t, err)
}

func TestVulkanExtension_CmdResetEvent2(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_synchronization2.NewMockLoader(ctrl)
	extension := khr_synchronization2.CreateExtensionDriverFromLoader(extDriver)

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	commandPool := mocks.NewDummyCommandPool(device)
	commandBuffer := mocks.NewDummyCommandBuffer(commandPool, device)
	event := mocks.NewDummyEvent(device)

	extDriver.EXPECT().VkCmdResetEvent2KHR(commandBuffer.Handle(), event.Handle(), khr_synchronization2_loader.VkPipelineStageFlags2(0x00020000))

	extension.CmdResetEvent2(commandBuffer, event, khr_synchronization2.PipelineStage2CommandPreprocess)
}

func TestVulkanExtension_CmdPipelineBarrier2(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_synchronization2.NewMockLoader(ctrl)
	extension := khr_synchronization2.CreateExtensionDriverFromLoader(extDriver)

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	commandPool := mocks.NewDummyCommandPool(device)
	commandBuffer := mocks.NewDummyCommandBuffer(commandPool, device)
	buffer1 := mocks.NewDummyBuffer(device)
	buffer2 := mocks.NewDummyBuffer(device)
	image := mocks.NewDummyImage(device)

	extDriver.EXPECT().VkCmdPipelineBarrier2KHR(commandBuffer.Handle(), gomock.Not(gomock.Nil())).DoAndReturn(
		func(commandBuffer loader.VkCommandBuffer, pDependencyInfo *khr_synchronization2_loader.VkDependencyInfo) {
			depVal := reflect.ValueOf(pDependencyInfo).Elem()

			require.Equal(t, uint64(0x3b9f9493), depVal.FieldByName("sType").Uint())
			require.True(t, depVal.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(1), depVal.FieldByName("dependencyFlags").Uint())
			require.Equal(t, uint64(1), depVal.FieldByName("memoryBarrierCount").Uint())
			require.Equal(t, uint64(2), depVal.FieldByName("bufferMemoryBarrierCount").Uint())
			require.Equal(t, uint64(1), depVal.FieldByName("imageMemoryBarrierCount").Uint())

			barrierPtr := (*khr_synchronization2_loader.VkMemoryBarrier2)(depVal.FieldByName("pMemoryBarriers").UnsafePointer())
			barrierSlice := unsafe.Slice(barrierPtr, 1)

			barrierVal := reflect.ValueOf(barrierSlice).Index(0)
			require.Equal(t, uint64(0x3b9f9490), barrierVal.FieldByName("sType").Uint())
			require.True(t, barrierVal.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0x02000000), barrierVal.FieldByName("srcStageMask").Uint())
			require.Equal(t, uint64(0x00200000), barrierVal.FieldByName("srcAccessMask").Uint())
			require.Equal(t, uint64(0x400000000), barrierVal.FieldByName("dstStageMask").Uint())
			require.Equal(t, uint64(0x00100000), barrierVal.FieldByName("dstAccessMask").Uint())

			bufferBarrierPtr := (*khr_synchronization2_loader.VkBufferMemoryBarrier2)(depVal.FieldByName("pBufferMemoryBarriers").UnsafePointer())
			bufferBarrierSlice := unsafe.Slice(bufferBarrierPtr, 2)

			barrierVal = reflect.ValueOf(bufferBarrierSlice).Index(0)
			require.Equal(t, uint64(0x3b9f9491), barrierVal.FieldByName("sType").Uint())
			require.True(t, barrierVal.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0x00002000), barrierVal.FieldByName("srcStageMask").Uint())
			require.Equal(t, uint64(0x00000200), barrierVal.FieldByName("srcAccessMask").Uint())
			require.Equal(t, uint64(0x100000000), barrierVal.FieldByName("dstStageMask").Uint())
			require.Equal(t, uint64(0x00000800), barrierVal.FieldByName("dstAccessMask").Uint())
			require.Equal(t, uint64(13), barrierVal.FieldByName("srcQueueFamilyIndex").Uint())
			require.Equal(t, uint64(17), barrierVal.FieldByName("dstQueueFamilyIndex").Uint())
			require.Equal(t, uint64(19), barrierVal.FieldByName("offset").Uint())
			require.Equal(t, uint64(23), barrierVal.FieldByName("size").Uint())
			require.Equal(t, buffer1.Handle(), (loader.VkBuffer)(barrierVal.FieldByName("buffer").UnsafePointer()))

			barrierVal = reflect.ValueOf(bufferBarrierSlice).Index(1)
			require.Equal(t, uint64(0x3b9f9491), barrierVal.FieldByName("sType").Uint())
			require.True(t, barrierVal.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0x00000400), barrierVal.FieldByName("srcStageMask").Uint())
			require.Equal(t, uint64(0x02000000), barrierVal.FieldByName("srcAccessMask").Uint())
			require.Equal(t, uint64(0x800000000), barrierVal.FieldByName("dstStageMask").Uint())
			require.Equal(t, uint64(0x01000000), barrierVal.FieldByName("dstAccessMask").Uint())
			require.Equal(t, uint64(29), barrierVal.FieldByName("srcQueueFamilyIndex").Uint())
			require.Equal(t, uint64(31), barrierVal.FieldByName("dstQueueFamilyIndex").Uint())
			require.Equal(t, uint64(37), barrierVal.FieldByName("offset").Uint())
			require.Equal(t, uint64(41), barrierVal.FieldByName("size").Uint())
			require.Equal(t, buffer2.Handle(), (loader.VkBuffer)(barrierVal.FieldByName("buffer").UnsafePointer()))

			imageBarrierPtr := (*khr_synchronization2_loader.VkImageMemoryBarrier2)(depVal.FieldByName("pImageMemoryBarriers").UnsafePointer())
			imageBarrierSlice := unsafe.Slice(imageBarrierPtr, 1)

			barrierVal = reflect.ValueOf(imageBarrierSlice).Index(0)
			require.Equal(t, uint64(0x3b9f9492), barrierVal.FieldByName("sType").Uint())
			require.True(t, barrierVal.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0x00020000), barrierVal.FieldByName("srcStageMask").Uint())
			require.Equal(t, uint64(0x00400000), barrierVal.FieldByName("srcAccessMask").Uint())
			require.Equal(t, uint64(0x00000800), barrierVal.FieldByName("dstStageMask").Uint())
			require.Equal(t, uint64(0x00100000), barrierVal.FieldByName("dstAccessMask").Uint())
			require.Equal(t, uint64(3), barrierVal.FieldByName("oldLayout").Uint())
			require.Equal(t, uint64(8), barrierVal.FieldByName("newLayout").Uint())
			require.Equal(t, uint64(43), barrierVal.FieldByName("srcQueueFamilyIndex").Uint())
			require.Equal(t, uint64(47), barrierVal.FieldByName("dstQueueFamilyIndex").Uint())
			require.Equal(t, uint64(2), barrierVal.FieldByName("subresourceRange").FieldByName("aspectMask").Uint())
			require.Equal(t, uint64(53), barrierVal.FieldByName("subresourceRange").FieldByName("baseMipLevel").Uint())
			require.Equal(t, uint64(59), barrierVal.FieldByName("subresourceRange").FieldByName("levelCount").Uint())
			require.Equal(t, uint64(61), barrierVal.FieldByName("subresourceRange").FieldByName("baseArrayLayer").Uint())
			require.Equal(t, uint64(67), barrierVal.FieldByName("subresourceRange").FieldByName("layerCount").Uint())
			require.Equal(t, image.Handle(), (loader.VkImage)(barrierVal.FieldByName("image").UnsafePointer()))
		})

	err := extension.CmdPipelineBarrier2(commandBuffer, khr_synchronization2.DependencyInfo{
		DependencyFlags: core1_0.DependencyByRegion,
		MemoryBarriers: []khr_synchronization2.MemoryBarrier2{
			{
				SrcStageMask:  khr_synchronization2.PipelineStage2AccelerationStructureBuildNV,
				SrcAccessMask: khr_synchronization2.Access2AccelerationStructureReadNV,
				DstStageMask:  khr_synchronization2.PipelineStage2Blit,
				DstAccessMask: khr_synchronization2.Access2ConditionalRenderingRead,
			},
		},
		BufferMemoryBarriers: []khr_synchronization2.BufferMemoryBarrier2{
			{
				SrcStageMask:        khr_synchronization2.PipelineStage2BottomOfPipe,
				SrcAccessMask:       khr_synchronization2.Access2DepthStencilAttachmentRead,
				DstStageMask:        khr_synchronization2.PipelineStage2Copy,
				DstAccessMask:       khr_synchronization2.Access2TransferRead,
				SrcQueueFamilyIndex: 13,
				DstQueueFamilyIndex: 17,
				Buffer:              buffer1,
				Offset:              19,
				Size:                23,
			},
			{
				SrcStageMask:        khr_synchronization2.PipelineStage2ColorAttachmentOutput,
				SrcAccessMask:       khr_synchronization2.Access2TransformFeedbackWrite,
				DstStageMask:        khr_synchronization2.PipelineStage2Clear,
				DstAccessMask:       khr_synchronization2.Access2FragmentDensityMapRead,
				SrcQueueFamilyIndex: 29,
				DstQueueFamilyIndex: 31,
				Buffer:              buffer2,
				Offset:              37,
				Size:                41,
			},
		},
		ImageMemoryBarriers: []khr_synchronization2.ImageMemoryBarrier2{
			{
				SrcStageMask:        khr_synchronization2.PipelineStage2CommandPreprocess,
				SrcAccessMask:       khr_synchronization2.Access2AccelerationStructureWrite,
				DstStageMask:        khr_synchronization2.PipelineStage2ComputeShader,
				DstAccessMask:       khr_synchronization2.Access2ConditionalRenderingRead,
				OldLayout:           core1_0.ImageLayoutDepthStencilAttachmentOptimal,
				NewLayout:           core1_0.ImageLayoutPreInitialized,
				SrcQueueFamilyIndex: 43,
				DstQueueFamilyIndex: 47,
				Image:               image,
				SubresourceRange: core1_0.ImageSubresourceRange{
					AspectMask:     core1_0.ImageAspectDepth,
					BaseMipLevel:   53,
					LevelCount:     59,
					BaseArrayLayer: 61,
					LayerCount:     67,
				},
			},
		},
	})
	require.NoError(t, err)
}
