package khr_16bit_storage_test

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
	"github.com/vkngwrapper/extensions/v3/khr_16bit_storage"
	khr_16bit_storage_driver "github.com/vkngwrapper/extensions/v3/khr_16bit_storage/loader"
	"github.com/vkngwrapper/extensions/v3/khr_get_physical_device_properties2"
	khr_get_physical_device_properties2_driver "github.com/vkngwrapper/extensions/v3/khr_get_physical_device_properties2/loader"
	mock_get_physical_device_properties2 "github.com/vkngwrapper/extensions/v3/khr_get_physical_device_properties2/mocks"
	"go.uber.org/mock/gomock"
)

func TestDevice16BitStorageOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	instance := mocks.NewDummyInstance(common.Vulkan1_0, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_0)
	expectedDevice := mocks.NewDummyDevice(common.Vulkan1_0, []string{})

	coreLoader := mock_driver.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := mocks1_0.InternalCoreInstanceDriver(instance, coreLoader)

	coreLoader.EXPECT().VkCreateDevice(physicalDevice.Handle(), gomock.Not(gomock.Nil()), gomock.Nil(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(physicalDevice loader.VkPhysicalDevice, pCreateInfo *loader.VkDeviceCreateInfo, pAllocator *loader.VkAllocationCallbacks, pDevice *loader.VkDevice) (common.VkResult, error) {
			*pDevice = expectedDevice.Handle()

			val := reflect.ValueOf(pCreateInfo).Elem()
			require.Equal(t, uint64(3), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_CREATE_INFO
			require.Equal(t, uint64(0), val.FieldByName("flags").Uint())

			storageFeatures := (*khr_16bit_storage_driver.VkPhysicalDevice16BitStorageFeaturesKHR)(val.FieldByName("pNext").UnsafePointer())
			storageVal := reflect.ValueOf(storageFeatures).Elem()

			require.Equal(t, uint64(1000083000), storageVal.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_16BIT_STORAGE_FEATURES_KHR
			require.True(t, storageVal.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0), storageVal.FieldByName("storageBuffer16BitAccess").Uint())
			require.Equal(t, uint64(1), storageVal.FieldByName("uniformAndStorageBuffer16BitAccess").Uint())
			require.Equal(t, uint64(0), storageVal.FieldByName("storagePushConstant16").Uint())
			require.Equal(t, uint64(1), storageVal.FieldByName("storageInputOutput16").Uint())

			return core1_0.VKSuccess, nil
		})

	storage := khr_16bit_storage.PhysicalDevice16BitStorageFeatures{
		StorageInputOutput16:               true,
		UniformAndStorageBuffer16BitAccess: true,
		StoragePushConstant16:              false,
		StorageBuffer16BitAccess:           false,
	}
	device, _, err := driver.CreateDevice(physicalDevice, nil, core1_0.DeviceCreateInfo{
		QueueCreateInfos: []core1_0.DeviceQueueCreateInfo{
			{
				QueuePriorities: []float32{0},
			},
		},

		NextOptions: common.NextOptions{Next: storage},
	})

	require.NoError(t, err)
	require.NotNil(t, device)
	require.Equal(t, expectedDevice.Handle(), device.Handle())
}

func TestDevice16BitStorageOutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	instance := mocks.NewDummyInstance(common.Vulkan1_0, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_0)

	extDriver := mock_get_physical_device_properties2.NewMockLoader(ctrl)
	extension := khr_get_physical_device_properties2.CreateExtensionDriverFromLoader(extDriver)

	extDriver.EXPECT().VkGetPhysicalDeviceFeatures2KHR(physicalDevice.Handle(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(physicalDevice loader.VkPhysicalDevice, pFeatures *khr_get_physical_device_properties2_driver.VkPhysicalDeviceFeatures2KHR) {
			val := reflect.ValueOf(pFeatures).Elem()

			require.Equal(t, uint64(1000059000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FEATURES_2_KHR

			featureVal := val.FieldByName("features")
			*(*loader.VkBool32)(unsafe.Pointer(featureVal.FieldByName("fillModeNonSolid").UnsafeAddr())) = loader.VkBool32(1)

			outDataPtr := (*khr_16bit_storage_driver.VkPhysicalDevice16BitStorageFeaturesKHR)(val.FieldByName("pNext").UnsafePointer())
			outDataVal := reflect.ValueOf(outDataPtr).Elem()
			*(*loader.VkBool32)(unsafe.Pointer(outDataVal.FieldByName("storageBuffer16BitAccess").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(outDataVal.FieldByName("uniformAndStorageBuffer16BitAccess").UnsafeAddr())) = loader.VkBool32(0)
			*(*loader.VkBool32)(unsafe.Pointer(outDataVal.FieldByName("storagePushConstant16").UnsafeAddr())) = loader.VkBool32(1)
			*(*loader.VkBool32)(unsafe.Pointer(outDataVal.FieldByName("storageInputOutput16").UnsafeAddr())) = loader.VkBool32(1)
		})

	outData := &khr_16bit_storage.PhysicalDevice16BitStorageFeatures{}
	features := &khr_get_physical_device_properties2.PhysicalDeviceFeatures2{
		NextOutData: common.NextOutData{Next: outData},
	}

	err := extension.GetPhysicalDeviceFeatures2(physicalDevice, features)
	require.NoError(t, err)

	require.True(t, outData.StoragePushConstant16)
	require.False(t, outData.UniformAndStorageBuffer16BitAccess)
	require.True(t, outData.StorageInputOutput16)
	require.False(t, outData.StorageBuffer16BitAccess)

	require.True(t, features.Features.FillModeNonSolid)
}
