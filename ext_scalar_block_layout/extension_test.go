package ext_scalar_block_layout_test

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v2/mocks"
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/common/extensions"
	"github.com/vkngwrapper/core/v2/core1_0"
	"github.com/vkngwrapper/core/v2/driver"
	mock_driver "github.com/vkngwrapper/core/v2/driver/mocks"
	"github.com/vkngwrapper/extensions/v2/ext_scalar_block_layout"
	ext_scalar_block_layout_driver "github.com/vkngwrapper/extensions/v2/ext_scalar_block_layout/driver"
	"github.com/vkngwrapper/extensions/v2/khr_get_physical_device_properties2"
	khr_get_physical_device_properties2_driver "github.com/vkngwrapper/extensions/v2/khr_get_physical_device_properties2/driver"
	mock_get_physical_device_properties2 "github.com/vkngwrapper/extensions/v2/khr_get_physical_device_properties2/mocks"
	"reflect"
	"testing"
	"unsafe"
)

func TestPhysicalDeviceScalarBlockLayoutFeaturesOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	coreDriver.EXPECT().CreateDeviceDriver(gomock.Any()).Return(coreDriver, nil)
	instance := mocks.EasyMockInstance(ctrl, coreDriver)
	physicalDevice := extensions.CreatePhysicalDeviceObject(coreDriver, instance.Handle(), mocks.NewFakePhysicalDeviceHandle(), common.Vulkan1_0, common.Vulkan1_0)
	mockDevice := mocks.EasyMockDevice(ctrl, coreDriver)

	coreDriver.EXPECT().VkCreateDevice(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice driver.VkPhysicalDevice,
		pCreateInfo *driver.VkDeviceCreateInfo,
		pAllocator *driver.VkAllocationCallbacks,
		pDevice *driver.VkDevice) (common.VkResult, error) {
		*pDevice = mockDevice.Handle()

		val := reflect.ValueOf(pCreateInfo).Elem()
		require.Equal(t, uint64(3), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_CREATE_INFO

		next := (*ext_scalar_block_layout_driver.VkPhysicalDeviceScalarBlockLayoutFeaturesEXT)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000221000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SCALAR_BLOCK_LAYOUT_FEATURES_EXT
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1), val.FieldByName("scalarBlockLayout").Uint())

		return core1_0.VKSuccess, nil
	})

	device, _, err := physicalDevice.CreateDevice(
		nil,
		core1_0.DeviceCreateInfo{
			QueueCreateInfos: []core1_0.DeviceQueueCreateInfo{
				{
					QueuePriorities: []float32{0},
				},
			},
			NextOptions: common.NextOptions{
				ext_scalar_block_layout.PhysicalDeviceScalarBlockLayoutFeatures{
					ScalarBlockLayout: true,
				},
			},
		})
	require.NoError(t, err)
	require.Equal(t, mockDevice.Handle(), device.Handle())
}

func TestPhysicalDeviceScalarBlockLayoutFeaturesOutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_get_physical_device_properties2.NewMockDriver(ctrl)
	extension := khr_get_physical_device_properties2.CreateExtensionFromDriver(extDriver)

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	physicalDevice := mocks.EasyMockPhysicalDevice(ctrl, coreDriver)

	extDriver.EXPECT().VkGetPhysicalDeviceFeatures2KHR(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice driver.VkPhysicalDevice,
		pFeatures *khr_get_physical_device_properties2_driver.VkPhysicalDeviceFeatures2KHR) {

		val := reflect.ValueOf(pFeatures).Elem()
		require.Equal(t, uint64(1000059000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FEATURES_2_KHR

		next := (*ext_scalar_block_layout_driver.VkPhysicalDeviceScalarBlockLayoutFeaturesEXT)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000221000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SCALAR_BLOCK_LAYOUT_FEATURES_EXT
		require.True(t, val.FieldByName("pNext").IsNil())
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("scalarBlockLayout").UnsafeAddr())) = driver.VkBool32(1)
	})

	var outData ext_scalar_block_layout.PhysicalDeviceScalarBlockLayoutFeatures
	err := extension.PhysicalDeviceFeatures2(
		physicalDevice,
		&khr_get_physical_device_properties2.PhysicalDeviceFeatures2{
			NextOutData: common.NextOutData{&outData},
		},
	)
	require.NoError(t, err)
	require.Equal(t, ext_scalar_block_layout.PhysicalDeviceScalarBlockLayoutFeatures{
		ScalarBlockLayout: true,
	}, outData)
}
