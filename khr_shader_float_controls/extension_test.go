package khr_shader_float_controls_test

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/loader"
	"github.com/vkngwrapper/core/v3/mocks"
	"github.com/vkngwrapper/extensions/v3/khr_get_physical_device_properties2"
	khr_get_physical_device_properties2_driver "github.com/vkngwrapper/extensions/v3/khr_get_physical_device_properties2/loader"
	mock_get_physical_device_properties2 "github.com/vkngwrapper/extensions/v3/khr_get_physical_device_properties2/mocks"
	"github.com/vkngwrapper/extensions/v3/khr_shader_float_controls"
	khr_shader_float_controls_driver "github.com/vkngwrapper/extensions/v3/khr_shader_float_controls/loader"
	"go.uber.org/mock/gomock"
)

func TestPhysicalDeviceFloatControlsOutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_get_physical_device_properties2.NewMockLoader(ctrl)
	extension := khr_get_physical_device_properties2.CreateExtensionFromDriver(extDriver)

	instance := mocks.NewDummyInstance(common.Vulkan1_0, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_0)

	extDriver.EXPECT().VkGetPhysicalDeviceProperties2KHR(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice loader.VkPhysicalDevice,
		pProperties *khr_get_physical_device_properties2_driver.VkPhysicalDeviceProperties2KHR) {

		val := reflect.ValueOf(pProperties).Elem()
		require.Equal(t, uint64(1000059001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROPERTIES_2_KHR

		next := (*khr_shader_float_controls_driver.VkPhysicalDeviceFloatControlsPropertiesKHR)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()
		require.Equal(t, uint64(1000197000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FLOAT_CONTROLS_PROPERTIES_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		denormBehavior := (*khr_shader_float_controls_driver.VkShaderFloatControlsIndependenceKHR)(unsafe.Pointer(val.FieldByName("denormBehaviorIndependence").UnsafeAddr()))
		*denormBehavior = khr_shader_float_controls_driver.VkShaderFloatControlsIndependenceKHR(0) // VK_SHADER_FLOAT_CONTROLS_INDEPENDENCE_32_BIT_ONLY_KHR
		roundingMode := (*khr_shader_float_controls_driver.VkShaderFloatControlsIndependenceKHR)(unsafe.Pointer(val.FieldByName("roundingModeIndependence").UnsafeAddr()))
		*roundingMode = khr_shader_float_controls_driver.VkShaderFloatControlsIndependenceKHR(1) // VK_SHADER_FLOAT_CONTROLS_INDEPENDENCE_ALL_KHR

		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderSignedZeroInfNanPreserveFloat16").UnsafeAddr())) = loader.VkBool32(1)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderSignedZeroInfNanPreserveFloat32").UnsafeAddr())) = loader.VkBool32(0)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderSignedZeroInfNanPreserveFloat64").UnsafeAddr())) = loader.VkBool32(1)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderDenormPreserveFloat16").UnsafeAddr())) = loader.VkBool32(0)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderDenormPreserveFloat32").UnsafeAddr())) = loader.VkBool32(1)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderDenormPreserveFloat64").UnsafeAddr())) = loader.VkBool32(0)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderDenormFlushToZeroFloat16").UnsafeAddr())) = loader.VkBool32(1)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderDenormFlushToZeroFloat32").UnsafeAddr())) = loader.VkBool32(0)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderDenormFlushToZeroFloat64").UnsafeAddr())) = loader.VkBool32(1)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderRoundingModeRTEFloat16").UnsafeAddr())) = loader.VkBool32(0)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderRoundingModeRTEFloat32").UnsafeAddr())) = loader.VkBool32(1)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderRoundingModeRTEFloat64").UnsafeAddr())) = loader.VkBool32(0)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderRoundingModeRTZFloat16").UnsafeAddr())) = loader.VkBool32(1)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderRoundingModeRTZFloat32").UnsafeAddr())) = loader.VkBool32(0)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderRoundingModeRTZFloat64").UnsafeAddr())) = loader.VkBool32(1)
	})

	var outData khr_shader_float_controls.PhysicalDeviceFloatControlsProperties
	err := extension.GetPhysicalDeviceProperties2(
		physicalDevice,
		&khr_get_physical_device_properties2.PhysicalDeviceProperties2{
			NextOutData: common.NextOutData{&outData},
		})
	require.NoError(t, err)
	require.Equal(t, khr_shader_float_controls.PhysicalDeviceFloatControlsProperties{
		DenormBehaviorIndependence: khr_shader_float_controls.ShaderFloatControlsIndependence32BitOnly,
		RoundingModeIndependence:   khr_shader_float_controls.ShaderFloatControlsIndependenceAll,

		ShaderSignedZeroInfNanPreserveFloat16: true,
		ShaderSignedZeroInfNanPreserveFloat32: false,
		ShaderSignedZeroInfNanPreserveFloat64: true,
		ShaderDenormPreserveFloat16:           false,
		ShaderDenormPreserveFloat32:           true,
		ShaderDenormPreserveFloat64:           false,
		ShaderDenormFlushToZeroFloat16:        true,
		ShaderDenormFlushToZeroFloat32:        false,
		ShaderDenormFlushToZeroFloat64:        true,
		ShaderRoundingModeRTEFloat16:          false,
		ShaderRoundingModeRTEFloat32:          true,
		ShaderRoundingModeRTEFloat64:          false,
		ShaderRoundingModeRTZFloat16:          true,
		ShaderRoundingModeRTZFloat32:          false,
		ShaderRoundingModeRTZFloat64:          true,
	}, outData)
}
