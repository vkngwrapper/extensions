package khr_buffer_device_address_test

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v2/mocks"
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/common/extensions"
	"github.com/vkngwrapper/core/v2/core1_0"
	"github.com/vkngwrapper/core/v2/driver"
	mock_driver "github.com/vkngwrapper/core/v2/driver/mocks"
	"github.com/vkngwrapper/extensions/v2/khr_buffer_device_address"
	khr_buffer_device_address_driver "github.com/vkngwrapper/extensions/v2/khr_buffer_device_address/driver"
	mock_buffer_device_address "github.com/vkngwrapper/extensions/v2/khr_buffer_device_address/mocks"
	"github.com/vkngwrapper/extensions/v2/khr_get_physical_device_properties2"
	khr_get_physical_device_properties2_driver "github.com/vkngwrapper/extensions/v2/khr_get_physical_device_properties2/driver"
	mock_get_physical_device_properties2 "github.com/vkngwrapper/extensions/v2/khr_get_physical_device_properties2/mocks"
	"reflect"
	"testing"
	"unsafe"
)

func TestVulkanExtension_GetBufferDeviceAddress(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_buffer_device_address.NewMockDriver(ctrl)
	extension := khr_buffer_device_address.CreateExtensionFromDriver(extDriver)

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := mocks.EasyMockDevice(ctrl, coreDriver)
	buffer := mocks.EasyMockBuffer(ctrl)

	extDriver.EXPECT().VkGetBufferDeviceAddressKHR(
		device.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device driver.VkDevice, pInfo *khr_buffer_device_address_driver.VkBufferDeviceAddressInfoKHR) driver.VkDeviceAddress {
		val := reflect.ValueOf(pInfo).Elem()

		require.Equal(t, uint64(1000244001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_BUFFER_DEVICE_ADDRESS_INFO_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, buffer.Handle(), driver.VkBuffer(val.FieldByName("buffer").UnsafePointer()))

		return 5
	})

	address, err := extension.GetBufferDeviceAddress(
		device,
		khr_buffer_device_address.BufferDeviceAddressInfo{
			Buffer: buffer,
		})
	require.NoError(t, err)
	require.Equal(t, uint64(5), address)
}

func TestVulkanExtension_GetBufferOpaqueCaptureAddress(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_buffer_device_address.NewMockDriver(ctrl)
	extension := khr_buffer_device_address.CreateExtensionFromDriver(extDriver)

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := mocks.EasyMockDevice(ctrl, coreDriver)
	buffer := mocks.EasyMockBuffer(ctrl)

	extDriver.EXPECT().VkGetBufferOpaqueCaptureAddressKHR(
		device.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device driver.VkDevice, pInfo *khr_buffer_device_address_driver.VkBufferDeviceAddressInfoKHR) driver.Uint64 {
		val := reflect.ValueOf(pInfo).Elem()

		require.Equal(t, uint64(1000244001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_BUFFER_DEVICE_ADDRESS_INFO_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, buffer.Handle(), driver.VkBuffer(val.FieldByName("buffer").UnsafePointer()))

		return 7
	})

	address, err := extension.GetBufferOpaqueCaptureAddress(
		device,
		khr_buffer_device_address.BufferDeviceAddressInfo{
			Buffer: buffer,
		})
	require.NoError(t, err)
	require.Equal(t, uint64(7), address)
}

func TestVulkanExtension_GetDeviceMemoryOpaqueCaptureAddress(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_buffer_device_address.NewMockDriver(ctrl)
	extension := khr_buffer_device_address.CreateExtensionFromDriver(extDriver)

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := mocks.EasyMockDevice(ctrl, coreDriver)
	deviceMemory := mocks.EasyMockDeviceMemory(ctrl)

	extDriver.EXPECT().VkGetDeviceMemoryOpaqueCaptureAddressKHR(
		device.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device driver.VkDevice, pInfo *khr_buffer_device_address_driver.VkDeviceMemoryOpaqueCaptureAddressInfoKHR) driver.Uint64 {
		val := reflect.ValueOf(pInfo).Elem()

		require.Equal(t, uint64(1000257004), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_MEMORY_OPAQUE_CAPTURE_ADDRESS_INFO_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, deviceMemory.Handle(), driver.VkDeviceMemory(val.FieldByName("memory").UnsafePointer()))

		return 11
	})

	address, err := extension.GetDeviceMemoryOpaqueCaptureAddress(
		device,
		khr_buffer_device_address.DeviceMemoryOpaqueCaptureAddressInfo{
			Memory: deviceMemory,
		})
	require.NoError(t, err)
	require.Equal(t, uint64(11), address)
}

func TestBufferOpaqueCaptureAddressCreateOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := extensions.CreateDeviceObject(coreDriver, mocks.NewFakeDeviceHandle(), common.Vulkan1_0)
	mockBuffer := mocks.EasyMockBuffer(ctrl)

	coreDriver.EXPECT().VkCreateBuffer(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device driver.VkDevice,
		pCreateInfo *driver.VkBufferCreateInfo,
		pAllocator *driver.VkAllocationCallbacks,
		pBuffer *driver.VkBuffer) (common.VkResult, error) {

		*pBuffer = mockBuffer.Handle()
		val := reflect.ValueOf(pCreateInfo).Elem()

		require.Equal(t, uint64(12), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_BUFFER_CREATE_INFO

		next := (*khr_buffer_device_address_driver.VkBufferOpaqueCaptureAddressCreateInfoKHR)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000257002), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_BUFFER_OPAQUE_CAPTURE_ADDRESS_CREATE_INFO_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(13), val.FieldByName("opaqueCaptureAddress").Uint())

		return core1_0.VKSuccess, nil
	})

	buffer, _, err := device.CreateBuffer(
		nil,
		core1_0.BufferCreateInfo{
			NextOptions: common.NextOptions{
				khr_buffer_device_address.BufferOpaqueCaptureAddressCreateInfo{
					OpaqueCaptureAddress: 13,
				},
			},
		})
	require.NoError(t, err)
	require.Equal(t, mockBuffer.Handle(), buffer.Handle())
}

func TestMemoryOpaqueCaptureAddressAllocateOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := extensions.CreateDeviceObject(coreDriver, mocks.NewFakeDeviceHandle(), common.Vulkan1_0)
	mockMemory := mocks.EasyMockDeviceMemory(ctrl)

	coreDriver.EXPECT().VkAllocateMemory(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device driver.VkDevice,
		pAllocateInfo *driver.VkMemoryAllocateInfo,
		pAllocator *driver.VkAllocationCallbacks,
		pMemory *driver.VkDeviceMemory) (common.VkResult, error) {

		*pMemory = mockMemory.Handle()
		val := reflect.ValueOf(pAllocateInfo).Elem()

		require.Equal(t, uint64(5), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_MEMORY_ALLOCATE_INFO

		next := (*khr_buffer_device_address_driver.VkMemoryOpaqueCaptureAddressAllocateInfoKHR)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000257003), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_MEMORY_OPAQUE_CAPTURE_ADDRESS_ALLOCATE_INFO_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(17), val.FieldByName("opaqueCaptureAddress").Uint())

		return core1_0.VKSuccess, nil
	})

	memory, _, err := device.AllocateMemory(
		nil,
		core1_0.MemoryAllocateInfo{
			NextOptions: common.NextOptions{
				khr_buffer_device_address.MemoryOpaqueCaptureAddressAllocateInfo{
					OpaqueCaptureAddress: 17,
				},
			},
		})
	require.NoError(t, err)
	require.Equal(t, mockMemory.Handle(), memory.Handle())
}

func TestPhysicalDeviceBufferAddressFeaturesOptions(t *testing.T) {
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

		next := (*khr_buffer_device_address_driver.VkPhysicalDeviceBufferDeviceAddressFeaturesKHR)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000257000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_BUFFER_DEVICE_ADDRESS_FEATURES_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1), val.FieldByName("bufferDeviceAddress").Uint())
		require.Equal(t, uint64(0), val.FieldByName("bufferDeviceAddressCaptureReplay").Uint())
		require.Equal(t, uint64(1), val.FieldByName("bufferDeviceAddressMultiDevice").Uint())

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

			NextOptions: common.NextOptions{khr_buffer_device_address.PhysicalDeviceBufferDeviceAddressFeatures{
				BufferDeviceAddress:            true,
				BufferDeviceAddressMultiDevice: true,
			}},
		})
	require.NoError(t, err)
	require.Equal(t, mockDevice.Handle(), device.Handle())
}

func TestPhysicalDeviceBufferAddressFeaturesOutData(t *testing.T) {
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

		next := (*khr_buffer_device_address_driver.VkPhysicalDeviceBufferDeviceAddressFeaturesKHR)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000257000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_BUFFER_DEVICE_ADDRESS_FEATURES_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("bufferDeviceAddress").UnsafeAddr())) = driver.VkBool32(0)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("bufferDeviceAddressCaptureReplay").UnsafeAddr())) = driver.VkBool32(1)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("bufferDeviceAddressMultiDevice").UnsafeAddr())) = driver.VkBool32(0)
	})

	var outData khr_buffer_device_address.PhysicalDeviceBufferDeviceAddressFeatures
	err := extension.PhysicalDeviceFeatures2(
		physicalDevice,
		&khr_get_physical_device_properties2.PhysicalDeviceFeatures2{
			NextOutData: common.NextOutData{&outData},
		})
	require.NoError(t, err)
	require.Equal(t, khr_buffer_device_address.PhysicalDeviceBufferDeviceAddressFeatures{
		BufferDeviceAddressCaptureReplay: true,
	}, outData)
}
