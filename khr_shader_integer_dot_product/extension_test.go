package khr_shader_integer_dot_product_test

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
	"github.com/vkngwrapper/extensions/v3/khr_shader_integer_dot_product"
	khr_shader_integer_dot_product_loader "github.com/vkngwrapper/extensions/v3/khr_shader_integer_dot_product/loader"
	"go.uber.org/mock/gomock"
)

func TestPhysicalDeviceShaderIntegerDotProductProperties_OutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	instance := mocks.NewDummyInstance(common.Vulkan1_1, []string{})
	mockLoader := mock_loader.NewMockLoader(ctrl)
	driver := mocks1_1.InternalCoreInstanceDriver(instance, mockLoader)

	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_1)

	mockLoader.EXPECT().VkGetPhysicalDeviceProperties2(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice loader.VkPhysicalDevice, pBaseProperties *loader.VkPhysicalDeviceProperties2) {
		baseProperties := reflect.ValueOf(pBaseProperties).Elem()

		require.Equal(t, uint64(1000059001), baseProperties.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROPERTIES_2
		require.False(t, baseProperties.FieldByName("pNext").IsNil())

		pProperties := (*khr_shader_integer_dot_product_loader.VkPhysicalDeviceShaderIntegerDotProductProperties)(baseProperties.FieldByName("pNext").UnsafePointer())
		properties := reflect.ValueOf(pProperties).Elem()

		require.Equal(t, uint64(0x3b9f0fc0), properties.FieldByName("sType").Uint())
		require.True(t, properties.FieldByName("pNext").IsNil())
		*(*uint64)(unsafe.Pointer(properties.FieldByName("integerDotProduct8BitUnsignedAccelerated").UnsafeAddr())) = uint64(1)
		*(*uint64)(unsafe.Pointer(properties.FieldByName("integerDotProduct8BitSignedAccelerated").UnsafeAddr())) = uint64(0)
		*(*uint64)(unsafe.Pointer(properties.FieldByName("integerDotProduct8BitMixedSignednessAccelerated").UnsafeAddr())) = uint64(1)
		*(*uint64)(unsafe.Pointer(properties.FieldByName("integerDotProduct4x8BitPackedUnsignedAccelerated").UnsafeAddr())) = uint64(0)
		*(*uint64)(unsafe.Pointer(properties.FieldByName("integerDotProduct4x8BitPackedSignedAccelerated").UnsafeAddr())) = uint64(1)
		*(*uint64)(unsafe.Pointer(properties.FieldByName("integerDotProduct4x8BitPackedMixedSignednessAccelerated").UnsafeAddr())) = uint64(0)
		*(*uint64)(unsafe.Pointer(properties.FieldByName("integerDotProduct16BitUnsignedAccelerated").UnsafeAddr())) = uint64(1)
		*(*uint64)(unsafe.Pointer(properties.FieldByName("integerDotProduct16BitSignedAccelerated").UnsafeAddr())) = uint64(0)
		*(*uint64)(unsafe.Pointer(properties.FieldByName("integerDotProduct16BitMixedSignednessAccelerated").UnsafeAddr())) = uint64(1)
		*(*uint64)(unsafe.Pointer(properties.FieldByName("integerDotProduct32BitUnsignedAccelerated").UnsafeAddr())) = uint64(0)
		*(*uint64)(unsafe.Pointer(properties.FieldByName("integerDotProduct32BitSignedAccelerated").UnsafeAddr())) = uint64(1)
		*(*uint64)(unsafe.Pointer(properties.FieldByName("integerDotProduct32BitMixedSignednessAccelerated").UnsafeAddr())) = uint64(0)
		*(*uint64)(unsafe.Pointer(properties.FieldByName("integerDotProduct64BitUnsignedAccelerated").UnsafeAddr())) = uint64(1)
		*(*uint64)(unsafe.Pointer(properties.FieldByName("integerDotProduct64BitSignedAccelerated").UnsafeAddr())) = uint64(0)
		*(*uint64)(unsafe.Pointer(properties.FieldByName("integerDotProduct64BitMixedSignednessAccelerated").UnsafeAddr())) = uint64(1)
		*(*uint64)(unsafe.Pointer(properties.FieldByName("integerDotProductAccumulatingSaturating8BitUnsignedAccelerated").UnsafeAddr())) = uint64(0)
		*(*uint64)(unsafe.Pointer(properties.FieldByName("integerDotProductAccumulatingSaturating8BitSignedAccelerated").UnsafeAddr())) = uint64(1)
		*(*uint64)(unsafe.Pointer(properties.FieldByName("integerDotProductAccumulatingSaturating8BitMixedSignednessAccelerated").UnsafeAddr())) = uint64(0)
		*(*uint64)(unsafe.Pointer(properties.FieldByName("integerDotProductAccumulatingSaturating4x8BitPackedUnsignedAccelerated").UnsafeAddr())) = uint64(1)
		*(*uint64)(unsafe.Pointer(properties.FieldByName("integerDotProductAccumulatingSaturating4x8BitPackedSignedAccelerated").UnsafeAddr())) = uint64(0)
		*(*uint64)(unsafe.Pointer(properties.FieldByName("integerDotProductAccumulatingSaturating4x8BitPackedMixedSignednessAccelerated").UnsafeAddr())) = uint64(1)
		*(*uint64)(unsafe.Pointer(properties.FieldByName("integerDotProductAccumulatingSaturating16BitUnsignedAccelerated").UnsafeAddr())) = uint64(0)
		*(*uint64)(unsafe.Pointer(properties.FieldByName("integerDotProductAccumulatingSaturating16BitSignedAccelerated").UnsafeAddr())) = uint64(1)
		*(*uint64)(unsafe.Pointer(properties.FieldByName("integerDotProductAccumulatingSaturating16BitMixedSignednessAccelerated").UnsafeAddr())) = uint64(0)
		*(*uint64)(unsafe.Pointer(properties.FieldByName("integerDotProductAccumulatingSaturating32BitUnsignedAccelerated").UnsafeAddr())) = uint64(1)
		*(*uint64)(unsafe.Pointer(properties.FieldByName("integerDotProductAccumulatingSaturating32BitSignedAccelerated").UnsafeAddr())) = uint64(0)
		*(*uint64)(unsafe.Pointer(properties.FieldByName("integerDotProductAccumulatingSaturating32BitMixedSignednessAccelerated").UnsafeAddr())) = uint64(1)
		*(*uint64)(unsafe.Pointer(properties.FieldByName("integerDotProductAccumulatingSaturating64BitUnsignedAccelerated").UnsafeAddr())) = uint64(0)
		*(*uint64)(unsafe.Pointer(properties.FieldByName("integerDotProductAccumulatingSaturating64BitSignedAccelerated").UnsafeAddr())) = uint64(1)
		*(*uint64)(unsafe.Pointer(properties.FieldByName("integerDotProductAccumulatingSaturating64BitMixedSignednessAccelerated").UnsafeAddr())) = uint64(0)
	})

	var baseProperties core1_1.PhysicalDeviceProperties2
	var properties khr_shader_integer_dot_product.PhysicalDeviceShaderIntegerDotProductProperties

	baseProperties.Next = &properties

	err := driver.GetPhysicalDeviceProperties2(physicalDevice, &baseProperties)
	require.NoError(t, err)

	require.Equal(t, khr_shader_integer_dot_product.PhysicalDeviceShaderIntegerDotProductProperties{
		IntegerDotProduct8BitUnsignedAccelerated:                                      true,
		IntegerDotProduct8BitSignedAccelerated:                                        false,
		IntegerDotProduct8BitMixedSignednessAccelerated:                               true,
		IntegerDotProduct4x8BitPackedUnsignedAccelerated:                              false,
		IntegerDotProduct4x8BitPackedSignedAccelerated:                                true,
		IntegerDotProduct4x8BitPackedMixedSignednessAccelerated:                       false,
		IntegerDotProduct16BitUnsignedAccelerated:                                     true,
		IntegerDotProduct16BitSignedAccelerated:                                       false,
		IntegerDotProduct16BitMixedSignednessAccelerated:                              true,
		IntegerDotProduct32BitUnsignedAccelerated:                                     false,
		IntegerDotProduct32BitSignedAccelerated:                                       true,
		IntegerDotProduct32BitMixedSignednessAccelerated:                              false,
		IntegerDotProduct64BitUnsignedAccelerated:                                     true,
		IntegerDotProduct64BitSignedAccelerated:                                       false,
		IntegerDotProduct64BitMixedSignednessAccelerated:                              true,
		IntegerDotProductAccumulatingSaturating8BitUnsignedAccelerated:                false,
		IntegerDotProductAccumulatingSaturating8BitSignedAccelerated:                  true,
		IntegerDotProductAccumulatingSaturating8BitMixedSignednessAccelerated:         false,
		IntegerDotProductAccumulatingSaturating4x8BitPackedUnsignedAccelerated:        true,
		IntegerDotProductAccumulatingSaturating4x8BitPackedSignedAccelerated:          false,
		IntegerDotProductAccumulatingSaturating4x8BitPackedMixedSignednessAccelerated: true,
		IntegerDotProductAccumulatingSaturating16BitUnsignedAccelerated:               false,
		IntegerDotProductAccumulatingSaturating16BitSignedAccelerated:                 true,
		IntegerDotProductAccumulatingSaturating16BitMixedSignednessAccelerated:        false,
		IntegerDotProductAccumulatingSaturating32BitUnsignedAccelerated:               true,
		IntegerDotProductAccumulatingSaturating32BitSignedAccelerated:                 false,
		IntegerDotProductAccumulatingSaturating32BitMixedSignednessAccelerated:        true,
		IntegerDotProductAccumulatingSaturating64BitUnsignedAccelerated:               false,
		IntegerDotProductAccumulatingSaturating64BitSignedAccelerated:                 true,
		IntegerDotProductAccumulatingSaturating64BitMixedSignednessAccelerated:        false,
	}, properties)
}

func TestPhysicalDeviceShaderIntegerDotProductFeatures_Options(t *testing.T) {
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

		pInnerFeatures := (*khr_shader_integer_dot_product_loader.VkPhysicalDeviceShaderIntegerDotProductFeatures)(info.FieldByName("pNext").UnsafePointer())
		features := reflect.ValueOf(pInnerFeatures).Elem()

		require.Equal(t, uint64(0x3b9f0fc0), features.FieldByName("sType").Uint())
		require.True(t, features.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1), features.FieldByName("shaderIntegerDotProduct").Uint())

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
		NextOptions: common.NextOptions{Next: khr_shader_integer_dot_product.PhysicalDeviceShaderIntegerDotProductFeatures{
			ShaderIntegerDotProduct: true,
		}},
	})
	require.NoError(t, err)
	require.Equal(t, core1_0.VKSuccess, res)
	require.NotNil(t, device)
	require.Equal(t, expectedDevice.Handle(), device.Device().Handle())
}

func TestPhysicalDeviceShaderIntegerDotProductFeatures_OutData(t *testing.T) {
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

		pInnerFeatures := (*khr_shader_integer_dot_product_loader.VkPhysicalDeviceShaderIntegerDotProductFeatures)(info.FieldByName("pNext").UnsafePointer())
		features := reflect.ValueOf(pInnerFeatures).Elem()
		require.Equal(t, uint64(0x3b9f0fc0), features.FieldByName("sType").Uint())
		require.True(t, features.FieldByName("pNext").IsNil())
		*(*uint64)(unsafe.Pointer(features.FieldByName("shaderIntegerDotProduct").UnsafeAddr())) = uint64(1)
	})

	var baseFeatures core1_1.PhysicalDeviceFeatures2
	var features khr_shader_integer_dot_product.PhysicalDeviceShaderIntegerDotProductFeatures

	baseFeatures.NextOutData.Next = &features

	err := driver.GetPhysicalDeviceFeatures2(physicalDevice, &baseFeatures)
	require.NoError(t, err)

	require.Equal(t, khr_shader_integer_dot_product.PhysicalDeviceShaderIntegerDotProductFeatures{
		ShaderIntegerDotProduct: true,
	}, features)
}
