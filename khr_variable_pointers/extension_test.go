package khr_variable_pointers_test

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
	"github.com/vkngwrapper/extensions/v3/khr_variable_pointers"
	khr_variable_pointers_driver "github.com/vkngwrapper/extensions/v3/khr_variable_pointers/loader"
	"go.uber.org/mock/gomock"
)

func TestVariablePointersFeaturesOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	instance := mocks.NewDummyInstance(common.Vulkan1_0, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_0)
	mockDevice := mocks.NewDummyDevice(common.Vulkan1_0, []string{})

	coreLoader := mock_driver.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := mocks1_0.InternalCoreInstanceDriver(instance, coreLoader)

	coreLoader.EXPECT().VkCreateDevice(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(physicalDevice loader.VkPhysicalDevice,
			pCreateInfo *loader.VkDeviceCreateInfo,
			pAllocator *loader.VkAllocationCallbacks,
			pDevice *loader.VkDevice) (common.VkResult, error) {

			*pDevice = mockDevice.Handle()

			val := reflect.ValueOf(pCreateInfo).Elem()
			require.Equal(t, uint64(3), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_CREATE_INFO

			featuresPtr := (*khr_variable_pointers_driver.VkPhysicalDeviceVariablePointersFeaturesKHR)(val.FieldByName("pNext").UnsafePointer())
			val = reflect.ValueOf(featuresPtr).Elem()

			require.Equal(t, uint64(1000120000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VARIABLE_POINTERS_FEATURES_KHR
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(1), val.FieldByName("variablePointers").Uint())
			require.Equal(t, uint64(0), val.FieldByName("variablePointersStorageBuffer").Uint())

			return core1_0.VKSuccess, nil
		})

	device, _, err := driver.CreateDevice(physicalDevice, nil, core1_0.DeviceCreateInfo{
		QueueCreateInfos: []core1_0.DeviceQueueCreateInfo{
			{
				QueuePriorities: []float32{0},
			},
		},

		NextOptions: common.NextOptions{Next: khr_variable_pointers.PhysicalDeviceVariablePointersFeatures{
			VariablePointers:              true,
			VariablePointersStorageBuffer: false,
		}},
	})
	require.NoError(t, err)
	require.NotNil(t, device)
	require.Equal(t, mockDevice.Handle(), device.Handle())
}

func TestVariablePointersFeaturesOutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	instance := mocks.NewDummyInstance(common.Vulkan1_0, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_0)

	extDriver := mock_get_physical_device_properties2.NewMockLoader(ctrl)
	extension := khr_get_physical_device_properties2.CreateExtensionDriverFromLoader(extDriver)

	var pointersOutData khr_variable_pointers.PhysicalDeviceVariablePointersFeatures

	extDriver.EXPECT().VkGetPhysicalDeviceFeatures2KHR(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(
			physicalDevice loader.VkPhysicalDevice,
			pFeatures *khr_get_physical_device_properties2_driver.VkPhysicalDeviceFeatures2KHR,
		) {
			val := reflect.ValueOf(pFeatures).Elem()

			require.Equal(t, uint64(1000059000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FEATURES_2_KHR

			outData := (*khr_variable_pointers_driver.VkPhysicalDeviceVariablePointersFeaturesKHR)(val.FieldByName("pNext").UnsafePointer())
			val = reflect.ValueOf(outData).Elem()

			require.Equal(t, uint64(1000120000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VARIABLE_POINTERS_FEATURES_KHR
			require.True(t, val.FieldByName("pNext").IsNil())

			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("variablePointers").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("variablePointersStorageBuffer").UnsafeAddr())) = loader.VkBool32(1)
		})

	err := extension.GetPhysicalDeviceFeatures2(physicalDevice, &khr_get_physical_device_properties2.PhysicalDeviceFeatures2{
		NextOutData: common.NextOutData{Next: &pointersOutData},
	})
	require.NoError(t, err)
	require.True(t, pointersOutData.VariablePointersStorageBuffer)
	require.False(t, pointersOutData.VariablePointers)
}
