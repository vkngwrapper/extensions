package ext_extended_dynamic_state_test

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/core1_1"
	"github.com/vkngwrapper/core/v3/loader"
	mock_loader "github.com/vkngwrapper/core/v3/loader/mocks"
	"github.com/vkngwrapper/core/v3/mocks"
	"github.com/vkngwrapper/core/v3/mocks/mocks1_1"
	"github.com/vkngwrapper/extensions/v3/ext_extended_dynamic_state"
	ext_extended_dynamic_state_loader "github.com/vkngwrapper/extensions/v3/ext_extended_dynamic_state/loader"
	mock_extended_dynamic_state "github.com/vkngwrapper/extensions/v3/ext_extended_dynamic_state/mocks"
	"go.uber.org/mock/gomock"
)

func TestPhysicalDeviceExtendedDynamicStateFeatures_OutData(t *testing.T) {
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

		pInnerFeatures := (*ext_extended_dynamic_state_loader.VkPhysicalDeviceExtendedDynamicStateFeaturesEXT)(info.FieldByName("pNext").UnsafePointer())
		features := reflect.ValueOf(pInnerFeatures).Elem()
		require.Equal(t, uint64(0x3b9edcf8), features.FieldByName("sType").Uint())
		require.True(t, features.FieldByName("pNext").IsNil())
		*(*uint64)(unsafe.Pointer(features.FieldByName("extendedDynamicState").UnsafeAddr())) = uint64(1)
	})

	var baseFeatures core1_1.PhysicalDeviceFeatures2
	var features ext_extended_dynamic_state.PhysicalDeviceExtendedDynamicStateFeatures

	baseFeatures.NextOutData.Next = &features

	err := driver.GetPhysicalDeviceFeatures2(physicalDevice, &baseFeatures)
	require.NoError(t, err)

	require.Equal(t, ext_extended_dynamic_state.PhysicalDeviceExtendedDynamicStateFeatures{
		ExtendedDynamicState: true,
	}, features)
}

func TestPhysicalDeviceExtendedDynamicStateFeatures_Options(t *testing.T) {
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

		pInnerFeatures := (*ext_extended_dynamic_state_loader.VkPhysicalDeviceExtendedDynamicStateFeaturesEXT)(info.FieldByName("pNext").UnsafePointer())
		features := reflect.ValueOf(pInnerFeatures).Elem()

		require.Equal(t, uint64(0x3b9edcf8), features.FieldByName("sType").Uint())
		require.True(t, features.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1), features.FieldByName("extendedDynamicState").Uint())

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
		NextOptions: common.NextOptions{Next: ext_extended_dynamic_state.PhysicalDeviceExtendedDynamicStateFeatures{
			ExtendedDynamicState: true,
		}},
	})
	require.NoError(t, err)
	require.Equal(t, core1_0.VKSuccess, res)
	require.NotNil(t, device)
	require.Equal(t, expectedDevice.Handle(), device.Device().Handle())
}

func TestVulkanExtension_CmdBindVertexBuffers2(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_extended_dynamic_state.NewMockLoader(ctrl)
	extension := ext_extended_dynamic_state.CreateExtensionDriverFromLoader(extDriver)

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	commandPool := mocks.NewDummyCommandPool(device)
	commandBuffer := mocks.NewDummyCommandBuffer(commandPool, device)

	buffers := []core1_0.Buffer{
		mocks.NewDummyBuffer(device),
		mocks.NewDummyBuffer(device),
	}

	extDriver.EXPECT().VkCmdBindVertexBuffers2(commandBuffer.Handle(), loader.Uint32(1), loader.Uint32(2), gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).DoAndReturn(
		func(commandBuffer loader.VkCommandBuffer, firstBinding loader.Uint32, bindingCount loader.Uint32, pBuffers *loader.VkBuffer, pOffsets *loader.VkDeviceSize, pSizes *loader.VkDeviceSize, pStrides *loader.VkDeviceSize) {
			bufferSlice := ([]loader.VkBuffer)(unsafe.Slice(pBuffers, 2))
			offsetSlice := ([]loader.VkDeviceSize)(unsafe.Slice(pOffsets, 2))
			sizeSlice := ([]loader.VkDeviceSize)(unsafe.Slice(pSizes, 2))
			strideSlice := ([]loader.VkDeviceSize)(unsafe.Slice(pStrides, 2))

			require.Equal(t, []loader.VkBuffer{buffers[0].Handle(), buffers[1].Handle()}, bufferSlice)
			require.Equal(t, []loader.VkDeviceSize{loader.VkDeviceSize(3), loader.VkDeviceSize(5)}, offsetSlice)
			require.Equal(t, []loader.VkDeviceSize{loader.VkDeviceSize(7), loader.VkDeviceSize(11)}, sizeSlice)
			require.Equal(t, []loader.VkDeviceSize{loader.VkDeviceSize(13), loader.VkDeviceSize(17)}, strideSlice)
		})

	extension.CmdBindVertexBuffers2(commandBuffer, 1, buffers, []int{3, 5}, []int{7, 11}, []int{13, 17})
}

func TestVulkanExtension_CmdBindVertexBuffers2_OmitOptional(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_extended_dynamic_state.NewMockLoader(ctrl)
	extension := ext_extended_dynamic_state.CreateExtensionDriverFromLoader(extDriver)

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	commandPool := mocks.NewDummyCommandPool(device)
	commandBuffer := mocks.NewDummyCommandBuffer(commandPool, device)

	buffers := []core1_0.Buffer{
		mocks.NewDummyBuffer(device),
		mocks.NewDummyBuffer(device),
	}

	extDriver.EXPECT().VkCmdBindVertexBuffers2(commandBuffer.Handle(), loader.Uint32(1), loader.Uint32(2), gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil()), gomock.Nil(), gomock.Nil()).DoAndReturn(
		func(commandBuffer loader.VkCommandBuffer, firstBinding loader.Uint32, bindingCount loader.Uint32, pBuffers *loader.VkBuffer, pOffsets *loader.VkDeviceSize, pSizes *loader.VkDeviceSize, pStrides *loader.VkDeviceSize) {
			bufferSlice := ([]loader.VkBuffer)(unsafe.Slice(pBuffers, 2))
			offsetSlice := ([]loader.VkDeviceSize)(unsafe.Slice(pOffsets, 2))

			require.Equal(t, []loader.VkBuffer{buffers[0].Handle(), buffers[1].Handle()}, bufferSlice)
			require.Equal(t, []loader.VkDeviceSize{loader.VkDeviceSize(17), loader.VkDeviceSize(19)}, offsetSlice)
		})

	extension.CmdBindVertexBuffers2(commandBuffer, 1, buffers, []int{17, 19}, nil, nil)
}

func TestVulkanExtension_CmdSetCullMode(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_extended_dynamic_state.NewMockLoader(ctrl)
	extension := ext_extended_dynamic_state.CreateExtensionDriverFromLoader(extDriver)

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	commandPool := mocks.NewDummyCommandPool(device)
	commandBuffer := mocks.NewDummyCommandBuffer(commandPool, device)

	extDriver.EXPECT().VkCmdSetCullMode(commandBuffer.Handle(), loader.VkCullModeFlags(2))

	extension.CmdSetCullMode(commandBuffer, core1_0.CullModeBack)
}

func TestVulkanExtension_CmdSetDepthBoundsTestEnable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_extended_dynamic_state.NewMockLoader(ctrl)
	extension := ext_extended_dynamic_state.CreateExtensionDriverFromLoader(extDriver)

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	commandPool := mocks.NewDummyCommandPool(device)
	commandBuffer := mocks.NewDummyCommandBuffer(commandPool, device)

	extDriver.EXPECT().VkCmdSetDepthBoundsTestEnable(commandBuffer.Handle(), loader.VkBool32(1))

	extension.CmdSetDepthBoundsTestEnable(commandBuffer, true)
}

func TestVulkanExtension_CmdSetDepthCompareOp(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_extended_dynamic_state.NewMockLoader(ctrl)
	extension := ext_extended_dynamic_state.CreateExtensionDriverFromLoader(extDriver)

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	commandPool := mocks.NewDummyCommandPool(device)
	commandBuffer := mocks.NewDummyCommandBuffer(commandPool, device)

	extDriver.EXPECT().VkCmdSetDepthCompareOp(commandBuffer.Handle(), loader.VkCompareOp(4))

	extension.CmdSetDepthCompareOp(commandBuffer, core1_0.CompareOpGreater)
}

func TestVulkanExtension_CmdSetDepthTestEnable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_extended_dynamic_state.NewMockLoader(ctrl)
	extension := ext_extended_dynamic_state.CreateExtensionDriverFromLoader(extDriver)

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	commandPool := mocks.NewDummyCommandPool(device)
	commandBuffer := mocks.NewDummyCommandBuffer(commandPool, device)

	extDriver.EXPECT().VkCmdSetDepthTestEnable(commandBuffer.Handle(), loader.VkBool32(1))

	extension.CmdSetDepthTestEnable(commandBuffer, true)
}

func TestVulkanExtension_CmdSetDepthWriteEnable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_extended_dynamic_state.NewMockLoader(ctrl)
	extension := ext_extended_dynamic_state.CreateExtensionDriverFromLoader(extDriver)

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	commandPool := mocks.NewDummyCommandPool(device)
	commandBuffer := mocks.NewDummyCommandBuffer(commandPool, device)

	extDriver.EXPECT().VkCmdSetDepthWriteEnable(commandBuffer.Handle(), loader.VkBool32(1))

	extension.CmdSetDepthWriteEnable(commandBuffer, true)
}

func TestVulkanExtension_CmdSetFrontFace(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_extended_dynamic_state.NewMockLoader(ctrl)
	extension := ext_extended_dynamic_state.CreateExtensionDriverFromLoader(extDriver)

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	commandPool := mocks.NewDummyCommandPool(device)
	commandBuffer := mocks.NewDummyCommandBuffer(commandPool, device)

	extDriver.EXPECT().VkCmdSetFrontFace(commandBuffer.Handle(), loader.VkFrontFace(1))

	extension.CmdSetFrontFace(commandBuffer, core1_0.FrontFaceClockwise)
}

func TestVulkanExtension_CmdSetPrimitiveTopology(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_extended_dynamic_state.NewMockLoader(ctrl)
	extension := ext_extended_dynamic_state.CreateExtensionDriverFromLoader(extDriver)

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	commandPool := mocks.NewDummyCommandPool(device)
	commandBuffer := mocks.NewDummyCommandBuffer(commandPool, device)

	extDriver.EXPECT().VkCmdSetPrimitiveTopology(commandBuffer.Handle(), loader.VkPrimitiveTopology(10))

	extension.CmdSetPrimitiveTopology(commandBuffer, core1_0.PrimitiveTopologyPatchList)
}

func TestVulkanExtension_CmdSetStencilOp(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_extended_dynamic_state.NewMockLoader(ctrl)
	extension := ext_extended_dynamic_state.CreateExtensionDriverFromLoader(extDriver)

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	commandPool := mocks.NewDummyCommandPool(device)
	commandBuffer := mocks.NewDummyCommandBuffer(commandPool, device)

	extDriver.EXPECT().VkCmdSetStencilOp(commandBuffer.Handle(), loader.VkStencilFaceFlags(1), loader.VkStencilOp(7), loader.VkStencilOp(5), loader.VkStencilOp(2), loader.VkCompareOp(5))

	extension.CmdSetStencilOp(commandBuffer, core1_0.StencilFaceFront, core1_0.StencilDecrementAndWrap, core1_0.StencilInvert, core1_0.StencilReplace, core1_0.CompareOpNotEqual)
}

func TestVulkanExtension_CmdSetStencilTestEnable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_extended_dynamic_state.NewMockLoader(ctrl)
	extension := ext_extended_dynamic_state.CreateExtensionDriverFromLoader(extDriver)

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	commandPool := mocks.NewDummyCommandPool(device)
	commandBuffer := mocks.NewDummyCommandBuffer(commandPool, device)

	extDriver.EXPECT().VkCmdSetStencilTestEnable(commandBuffer.Handle(), loader.VkBool32(1))

	extension.CmdSetStencilTestEnable(commandBuffer, true)
}

func TestVulkanExtension_CmdSetScissorWithCount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_extended_dynamic_state.NewMockLoader(ctrl)
	extension := ext_extended_dynamic_state.CreateExtensionDriverFromLoader(extDriver)

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	commandPool := mocks.NewDummyCommandPool(device)
	commandBuffer := mocks.NewDummyCommandBuffer(commandPool, device)

	extDriver.EXPECT().VkCmdSetScissorWithCount(commandBuffer.Handle(), loader.Uint32(2), gomock.Not(gomock.Nil())).
		DoAndReturn(func(commandBuffer loader.VkCommandBuffer, scissorCount loader.Uint32, pScissors *loader.VkRect2D) {
			scissorSlice := unsafe.Slice(pScissors, 2)
			scissorVal := reflect.ValueOf(scissorSlice).Index(0)

			require.Equal(t, int64(1), scissorVal.FieldByName("offset").FieldByName("x").Int())
			require.Equal(t, int64(3), scissorVal.FieldByName("offset").FieldByName("y").Int())
			require.Equal(t, uint64(5), scissorVal.FieldByName("extent").FieldByName("width").Uint())
			require.Equal(t, uint64(7), scissorVal.FieldByName("extent").FieldByName("height").Uint())

			scissorVal = reflect.ValueOf(scissorSlice).Index(1)

			require.Equal(t, int64(11), scissorVal.FieldByName("offset").FieldByName("x").Int())
			require.Equal(t, int64(13), scissorVal.FieldByName("offset").FieldByName("y").Int())
			require.Equal(t, uint64(17), scissorVal.FieldByName("extent").FieldByName("width").Uint())
			require.Equal(t, uint64(19), scissorVal.FieldByName("extent").FieldByName("height").Uint())
		})

	extension.CmdSetScissorWithCount(commandBuffer, []core1_0.Rect2D{
		{
			Offset: core1_0.Offset2D{
				X: 1,
				Y: 3,
			},
			Extent: core1_0.Extent2D{
				Width:  5,
				Height: 7,
			},
		},
		{
			Offset: core1_0.Offset2D{
				X: 11,
				Y: 13,
			},
			Extent: core1_0.Extent2D{
				Width:  17,
				Height: 19,
			},
		},
	})
}

func TestVulkanExtension_CmdSetViewportWithCount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_extended_dynamic_state.NewMockLoader(ctrl)
	extension := ext_extended_dynamic_state.CreateExtensionDriverFromLoader(extDriver)

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	commandPool := mocks.NewDummyCommandPool(device)
	commandBuffer := mocks.NewDummyCommandBuffer(commandPool, device)

	extDriver.EXPECT().VkCmdSetViewportWithCount(commandBuffer.Handle(), loader.Uint32(1), gomock.Not(gomock.Nil())).
		DoAndReturn(func(commandBuffer loader.VkCommandBuffer, viewportCount loader.Uint32, pViewports *loader.VkViewport) {
			viewportSlice := unsafe.Slice(pViewports, 1)
			viewportVal := reflect.ValueOf(viewportSlice).Index(0)

			require.InDelta(t, float64(1.5), viewportVal.FieldByName("x").Float(), 0.0001)
			require.InDelta(t, float64(3.5), viewportVal.FieldByName("y").Float(), 0.0001)
			require.InDelta(t, float64(5.5), viewportVal.FieldByName("width").Float(), 0.0001)
			require.InDelta(t, float64(7.5), viewportVal.FieldByName("height").Float(), 0.0001)
			require.InDelta(t, float64(11.5), viewportVal.FieldByName("minDepth").Float(), 0.0001)
			require.InDelta(t, float64(13.5), viewportVal.FieldByName("maxDepth").Float(), 0.0001)
		})

	extension.CmdSetViewportWithCount(commandBuffer, []core1_0.Viewport{
		{
			X:        1.5,
			Y:        3.5,
			Width:    5.5,
			Height:   7.5,
			MinDepth: 11.5,
			MaxDepth: 13.5,
		},
	})
}
