package ext_extended_dynamic_state2_test

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
	"github.com/vkngwrapper/extensions/v3/ext_extended_dynamic_state2"
	ext_extended_dynamic_state2_loader "github.com/vkngwrapper/extensions/v3/ext_extended_dynamic_state2/loader"
	mock_extended_dynamic_state2 "github.com/vkngwrapper/extensions/v3/ext_extended_dynamic_state2/mocks"
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

		pInnerFeatures := (*ext_extended_dynamic_state2_loader.VkPhysicalDeviceExtendedDynamicState2FeaturesEXT)(info.FieldByName("pNext").UnsafePointer())
		features := reflect.ValueOf(pInnerFeatures).Elem()
		require.Equal(t, uint64(0x3ba08aa8), features.FieldByName("sType").Uint())
		require.True(t, features.FieldByName("pNext").IsNil())
		*(*uint64)(unsafe.Pointer(features.FieldByName("extendedDynamicState2").UnsafeAddr())) = uint64(1)
		*(*uint64)(unsafe.Pointer(features.FieldByName("extendedDynamicState2LogicOp").UnsafeAddr())) = uint64(1)
		*(*uint64)(unsafe.Pointer(features.FieldByName("extendedDynamicState2PatchControlPoints").UnsafeAddr())) = uint64(1)
	})

	var baseFeatures core1_1.PhysicalDeviceFeatures2
	var features ext_extended_dynamic_state2.PhysicalDeviceExtendedDynamicState2Features

	baseFeatures.NextOutData.Next = &features

	err := driver.GetPhysicalDeviceFeatures2(physicalDevice, &baseFeatures)
	require.NoError(t, err)

	require.Equal(t, ext_extended_dynamic_state2.PhysicalDeviceExtendedDynamicState2Features{
		ExtendedDynamicState2:                   true,
		ExtendedDynamicState2LogicOp:            true,
		ExtendedDynamicState2PatchControlPoints: true,
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

		pInnerFeatures := (*ext_extended_dynamic_state2_loader.VkPhysicalDeviceExtendedDynamicState2FeaturesEXT)(info.FieldByName("pNext").UnsafePointer())
		features := reflect.ValueOf(pInnerFeatures).Elem()

		require.Equal(t, uint64(0x3ba08aa8), features.FieldByName("sType").Uint())
		require.True(t, features.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1), features.FieldByName("extendedDynamicState2").Uint())
		require.Equal(t, uint64(1), features.FieldByName("extendedDynamicState2LogicOp").Uint())
		require.Equal(t, uint64(1), features.FieldByName("extendedDynamicState2PatchControlPoints").Uint())

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
		NextOptions: common.NextOptions{Next: ext_extended_dynamic_state2.PhysicalDeviceExtendedDynamicState2Features{
			ExtendedDynamicState2:                   true,
			ExtendedDynamicState2LogicOp:            true,
			ExtendedDynamicState2PatchControlPoints: true,
		}},
	})
	require.NoError(t, err)
	require.Equal(t, core1_0.VKSuccess, res)
	require.NotNil(t, device)
	require.Equal(t, expectedDevice.Handle(), device.Device().Handle())
}

func TestVulkanExtension_CmdSetDepthBiasEnable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_extended_dynamic_state2.NewMockLoader(ctrl)
	extension := ext_extended_dynamic_state2.CreateExtensionDriverFromLoader(extDriver)

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	commandPool := mocks.NewDummyCommandPool(device)
	commandBuffer := mocks.NewDummyCommandBuffer(commandPool, device)

	extDriver.EXPECT().VkCmdSetDepthBiasEnableEXT(commandBuffer.Handle(), loader.VkBool32(1))

	extension.CmdSetDepthBiasEnable(commandBuffer, true)
}

func TestVulkanExtension_CmdSetLogicOp(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_extended_dynamic_state2.NewMockLoader(ctrl)
	extension := ext_extended_dynamic_state2.CreateExtensionDriverFromLoader(extDriver)

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	commandPool := mocks.NewDummyCommandPool(device)
	commandBuffer := mocks.NewDummyCommandBuffer(commandPool, device)

	extDriver.EXPECT().VkCmdSetLogicOpEXT(commandBuffer.Handle(), loader.VkLogicOp(4))

	extension.CmdSetLogicOp(commandBuffer, core1_0.LogicOpAndInverted)
}

func TestVulkanExtension_CmdSetPatchControlPoints(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_extended_dynamic_state2.NewMockLoader(ctrl)
	extension := ext_extended_dynamic_state2.CreateExtensionDriverFromLoader(extDriver)

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	commandPool := mocks.NewDummyCommandPool(device)
	commandBuffer := mocks.NewDummyCommandBuffer(commandPool, device)

	extDriver.EXPECT().VkCmdSetPatchControlPointsEXT(commandBuffer.Handle(), loader.Uint32(7))

	extension.CmdSetPatchControlPoints(commandBuffer, 7)
}

func TestVulkanExtension_CmdSetPrimitiveRestartEnable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_extended_dynamic_state2.NewMockLoader(ctrl)
	extension := ext_extended_dynamic_state2.CreateExtensionDriverFromLoader(extDriver)

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	commandPool := mocks.NewDummyCommandPool(device)
	commandBuffer := mocks.NewDummyCommandBuffer(commandPool, device)

	extDriver.EXPECT().VkCmdSetPrimitiveRestartEnableEXT(commandBuffer.Handle(), loader.VkBool32(1))

	extension.CmdSetPrimitiveRestartEnable(commandBuffer, true)
}

func TestVulkanExtension_CmdSetRasterizerDiscardEnable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_extended_dynamic_state2.NewMockLoader(ctrl)
	extension := ext_extended_dynamic_state2.CreateExtensionDriverFromLoader(extDriver)

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	commandPool := mocks.NewDummyCommandPool(device)
	commandBuffer := mocks.NewDummyCommandBuffer(commandPool, device)

	extDriver.EXPECT().VkCmdSetRasterizerDiscardEnableEXT(commandBuffer.Handle(), loader.VkBool32(1))

	extension.CmdSetRasterizerDiscardEnable(commandBuffer, true)
}
