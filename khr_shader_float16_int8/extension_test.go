package khr_shader_float16_int8_test

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
	mock_driver "github.com/vkngwrapper/core/v3/loader/mocks"
	"github.com/vkngwrapper/core/v3/mocks"
	"github.com/vkngwrapper/core/v3/mocks/mocks1_0"
	"github.com/vkngwrapper/extensions/v3/khr_get_physical_device_properties2"
	khr_get_physical_device_properties2_driver "github.com/vkngwrapper/extensions/v3/khr_get_physical_device_properties2/loader"
	mock_get_physical_device_properties2 "github.com/vkngwrapper/extensions/v3/khr_get_physical_device_properties2/mocks"
	"github.com/vkngwrapper/extensions/v3/khr_shader_float16_int8"
	khr_shader_float16_int8_driver "github.com/vkngwrapper/extensions/v3/khr_shader_float16_int8/loader"
	"go.uber.org/mock/gomock"
)

func TestPhysicalDeviceShaderFloat16Int8FeaturesOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreLoader := mock_driver.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := mocks1_0.InternalCoreInstanceDriver(coreLoader)
	instance := mocks.NewDummyInstance(common.Vulkan1_0, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_0)
	mockDevice := mocks.NewDummyDevice(common.Vulkan1_0, []string{})

	coreLoader.EXPECT().VkCreateDevice(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice loader.VkPhysicalDevice,
		pCreateInfo *loader.VkDeviceCreateInfo,
		pAllocator *loader.VkAllocationCallbacks,
		pDevice *loader.VkDevice) (common.VkResult, error) {
		*pDevice = mockDevice.Handle()

		val := reflect.ValueOf(pCreateInfo).Elem()
		require.Equal(t, uint64(3), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_CREATE_INFO

		next := (*khr_shader_float16_int8_driver.VkPhysicalDeviceShaderFloat16Int8FeaturesKHR)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000082000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_FLOAT16_INT8_FEATURES_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1), val.FieldByName("shaderInt8").Uint())
		require.Equal(t, uint64(1), val.FieldByName("shaderFloat16").Uint())

		return core1_0.VKSuccess, nil
	})

	device, _, err := driver.CreateDevice(
		physicalDevice,
		nil,
		core1_0.DeviceCreateInfo{
			QueueCreateInfos: []core1_0.DeviceQueueCreateInfo{
				{
					QueuePriorities: []float32{0},
				},
			},
			NextOptions: common.NextOptions{
				khr_shader_float16_int8.PhysicalDeviceShaderFloat16Int8Features{
					ShaderInt8:    true,
					ShaderFloat16: true,
				},
			},
		})
	require.NoError(t, err)
	require.Equal(t, mockDevice.Handle(), device.Handle())
}

func TestPhysicalDeviceShaderFloat16Int8FeaturesOutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_get_physical_device_properties2.NewMockLoader(ctrl)
	extension := khr_get_physical_device_properties2.CreateExtensionFromDriver(extDriver)

	instance := mocks.NewDummyInstance(common.Vulkan1_0, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_0)

	extDriver.EXPECT().VkGetPhysicalDeviceFeatures2KHR(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice loader.VkPhysicalDevice,
		pFeatures *khr_get_physical_device_properties2_driver.VkPhysicalDeviceFeatures2KHR) {

		val := reflect.ValueOf(pFeatures).Elem()
		require.Equal(t, uint64(1000059000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FEATURES_2_KHR

		next := (*khr_shader_float16_int8_driver.VkPhysicalDeviceShaderFloat16Int8FeaturesKHR)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000082000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_FLOAT16_INT8_FEATURES_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderInt8").UnsafeAddr())) = loader.VkBool32(1)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderFloat16").UnsafeAddr())) = loader.VkBool32(1)
	})

	var outData khr_shader_float16_int8.PhysicalDeviceShaderFloat16Int8Features
	err := extension.GetPhysicalDeviceFeatures2(
		physicalDevice,
		&khr_get_physical_device_properties2.PhysicalDeviceFeatures2{
			NextOutData: common.NextOutData{&outData},
		},
	)
	require.NoError(t, err)
	require.Equal(t, khr_shader_float16_int8.PhysicalDeviceShaderFloat16Int8Features{
		ShaderFloat16: true,
		ShaderInt8:    true,
	}, outData)
}
