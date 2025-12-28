package khr_vulkan_memory_model_test

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
	"github.com/vkngwrapper/extensions/v3/khr_vulkan_memory_model"
	khr_vulkan_memory_model_driver "github.com/vkngwrapper/extensions/v3/khr_vulkan_memory_model/loader"
	"go.uber.org/mock/gomock"
)

func TestPhysicalDeviceVulkanMemoryModelFeaturesOptions(t *testing.T) {
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

			featuresPtr := (*khr_vulkan_memory_model_driver.VkPhysicalDeviceVulkanMemoryModelFeaturesKHR)(val.FieldByName("pNext").UnsafePointer())
			val = reflect.ValueOf(featuresPtr).Elem()

			require.Equal(t, uint64(1000211000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VULKAN_MEMORY_MODEL_FEATURES_KHR
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(1), val.FieldByName("vulkanMemoryModel").Uint())
			require.Equal(t, uint64(0), val.FieldByName("vulkanMemoryModelDeviceScope").Uint())
			require.Equal(t, uint64(1), val.FieldByName("vulkanMemoryModelAvailabilityVisibilityChains").Uint())

			return core1_0.VKSuccess, nil
		})

	device, _, err := driver.CreateDevice(physicalDevice, nil, core1_0.DeviceCreateInfo{
		QueueCreateInfos: []core1_0.DeviceQueueCreateInfo{
			{
				QueuePriorities: []float32{0},
			},
		},

		NextOptions: common.NextOptions{Next: khr_vulkan_memory_model.PhysicalDeviceVulkanMemoryModelFeatures{
			VulkanMemoryModel:                             true,
			VulkanMemoryModelDeviceScope:                  false,
			VulkanMemoryModelAvailabilityVisibilityChains: true,
		}},
	})
	require.NoError(t, err)
	require.NotNil(t, device)
	require.Equal(t, mockDevice.Handle(), device.Handle())
}

func TestPhysicalDeviceVulkanMemoryModelFeaturesOutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	instance := mocks.NewDummyInstance(common.Vulkan1_0, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_0)

	extDriver := mock_get_physical_device_properties2.NewMockLoader(ctrl)
	extension := khr_get_physical_device_properties2.CreateExtensionDriverFromLoader(extDriver)

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

			outData := (*khr_vulkan_memory_model_driver.VkPhysicalDeviceVulkanMemoryModelFeaturesKHR)(val.FieldByName("pNext").UnsafePointer())
			val = reflect.ValueOf(outData).Elem()

			require.Equal(t, uint64(1000211000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VULKAN_MEMORY_MODEL_FEATURES_KHR
			require.True(t, val.FieldByName("pNext").IsNil())

			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("vulkanMemoryModel").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("vulkanMemoryModelDeviceScope").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("vulkanMemoryModelAvailabilityVisibilityChains").UnsafeAddr())) = loader.VkBool32(1)
		})

	var outData khr_vulkan_memory_model.PhysicalDeviceVulkanMemoryModelFeatures
	err := extension.GetPhysicalDeviceFeatures2(physicalDevice, &khr_get_physical_device_properties2.PhysicalDeviceFeatures2{
		NextOutData: common.NextOutData{Next: &outData},
	})
	require.NoError(t, err)
	require.Equal(t, khr_vulkan_memory_model.PhysicalDeviceVulkanMemoryModelFeatures{
		VulkanMemoryModel:                             true,
		VulkanMemoryModelDeviceScope:                  false,
		VulkanMemoryModelAvailabilityVisibilityChains: true,
	}, outData)
}
