package khr_buffer_device_address_test

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
	"github.com/vkngwrapper/extensions/v3/khr_buffer_device_address"
	khr_buffer_device_address_driver "github.com/vkngwrapper/extensions/v3/khr_buffer_device_address/loader"
	mock_buffer_device_address "github.com/vkngwrapper/extensions/v3/khr_buffer_device_address/mocks"
	"github.com/vkngwrapper/extensions/v3/khr_get_physical_device_properties2"
	khr_get_physical_device_properties2_driver "github.com/vkngwrapper/extensions/v3/khr_get_physical_device_properties2/loader"
	mock_get_physical_device_properties2 "github.com/vkngwrapper/extensions/v3/khr_get_physical_device_properties2/mocks"
	"go.uber.org/mock/gomock"
)

func TestVulkanExtension_GetBufferDeviceAddress(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	buffer := mocks.NewDummyBuffer(device)

	extDriver := mock_buffer_device_address.NewMockDriver(ctrl)
	extension := khr_buffer_device_address.CreateExtensionDriverFromLoader(extDriver, device)

	extDriver.EXPECT().VkGetBufferDeviceAddressKHR(
		device.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device loader.VkDevice, pInfo *khr_buffer_device_address_driver.VkBufferDeviceAddressInfoKHR) loader.VkDeviceAddress {
		val := reflect.ValueOf(pInfo).Elem()

		require.Equal(t, uint64(1000244001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_BUFFER_DEVICE_ADDRESS_INFO_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, buffer.Handle(), loader.VkBuffer(val.FieldByName("buffer").UnsafePointer()))

		return 5
	})

	address, err := extension.GetBufferDeviceAddress(
		khr_buffer_device_address.BufferDeviceAddressInfo{
			Buffer: buffer,
		})
	require.NoError(t, err)
	require.Equal(t, uint64(5), address)
}

func TestVulkanExtension_GetBufferOpaqueCaptureAddress(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	buffer := mocks.NewDummyBuffer(device)

	extDriver := mock_buffer_device_address.NewMockDriver(ctrl)
	extension := khr_buffer_device_address.CreateExtensionDriverFromLoader(extDriver, device)

	extDriver.EXPECT().VkGetBufferOpaqueCaptureAddressKHR(
		device.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device loader.VkDevice, pInfo *khr_buffer_device_address_driver.VkBufferDeviceAddressInfoKHR) loader.Uint64 {
		val := reflect.ValueOf(pInfo).Elem()

		require.Equal(t, uint64(1000244001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_BUFFER_DEVICE_ADDRESS_INFO_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, buffer.Handle(), loader.VkBuffer(val.FieldByName("buffer").UnsafePointer()))

		return 7
	})

	address, err := extension.GetBufferOpaqueCaptureAddress(
		khr_buffer_device_address.BufferDeviceAddressInfo{
			Buffer: buffer,
		})
	require.NoError(t, err)
	require.Equal(t, uint64(7), address)
}

func TestVulkanExtension_GetDeviceMemoryOpaqueCaptureAddress(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	deviceMemory := mocks.NewDummyDeviceMemory(device, 1)

	extDriver := mock_buffer_device_address.NewMockDriver(ctrl)
	extension := khr_buffer_device_address.CreateExtensionDriverFromLoader(extDriver, device)

	extDriver.EXPECT().VkGetDeviceMemoryOpaqueCaptureAddressKHR(
		device.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device loader.VkDevice, pInfo *khr_buffer_device_address_driver.VkDeviceMemoryOpaqueCaptureAddressInfoKHR) loader.Uint64 {
		val := reflect.ValueOf(pInfo).Elem()

		require.Equal(t, uint64(1000257004), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_MEMORY_OPAQUE_CAPTURE_ADDRESS_INFO_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, deviceMemory.Handle(), loader.VkDeviceMemory(val.FieldByName("memory").UnsafePointer()))

		return 11
	})

	address, err := extension.GetDeviceMemoryOpaqueCaptureAddress(
		khr_buffer_device_address.DeviceMemoryOpaqueCaptureAddressInfo{
			Memory: deviceMemory,
		})
	require.NoError(t, err)
	require.Equal(t, uint64(11), address)
}

func TestBufferOpaqueCaptureAddressCreateOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	mockBuffer := mocks.NewDummyBuffer(device)

	coreLoader := mock_driver.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := mocks1_0.InternalDeviceDriver(device, coreLoader)

	coreLoader.EXPECT().VkCreateBuffer(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device loader.VkDevice,
		pCreateInfo *loader.VkBufferCreateInfo,
		pAllocator *loader.VkAllocationCallbacks,
		pBuffer *loader.VkBuffer) (common.VkResult, error) {

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

	buffer, _, err := driver.CreateBuffer(
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

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	mockMemory := mocks.NewDummyDeviceMemory(device, 1)

	coreLoader := mock_driver.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := mocks1_0.InternalDeviceDriver(device, coreLoader)

	coreLoader.EXPECT().VkAllocateMemory(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device loader.VkDevice,
		pAllocateInfo *loader.VkMemoryAllocateInfo,
		pAllocator *loader.VkAllocationCallbacks,
		pMemory *loader.VkDeviceMemory) (common.VkResult, error) {

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

	memory, _, err := driver.AllocateMemory(
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
	).DoAndReturn(func(physicalDevice loader.VkPhysicalDevice,
		pCreateInfo *loader.VkDeviceCreateInfo,
		pAllocator *loader.VkAllocationCallbacks,
		pDevice *loader.VkDevice) (common.VkResult, error) {

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

	device, _, err := driver.CreateDevice(
		physicalDevice,
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

	extDriver := mock_get_physical_device_properties2.NewMockLoader(ctrl)
	extension := khr_get_physical_device_properties2.CreateExtensionDriverFromLoader(extDriver)

	instance := mocks.NewDummyInstance(common.Vulkan1_0, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_0)

	extDriver.EXPECT().VkGetPhysicalDeviceFeatures2KHR(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice loader.VkPhysicalDevice,
		pFeatures *khr_get_physical_device_properties2_driver.VkPhysicalDeviceFeatures2KHR) {

		val := reflect.ValueOf(pFeatures).Elem()
		require.Equal(t, uint64(1000059000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FEATURES_2_KHR

		next := (*khr_buffer_device_address_driver.VkPhysicalDeviceBufferDeviceAddressFeaturesKHR)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000257000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_BUFFER_DEVICE_ADDRESS_FEATURES_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("bufferDeviceAddress").UnsafeAddr())) = loader.VkBool32(0)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("bufferDeviceAddressCaptureReplay").UnsafeAddr())) = loader.VkBool32(1)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("bufferDeviceAddressMultiDevice").UnsafeAddr())) = loader.VkBool32(0)
	})

	var outData khr_buffer_device_address.PhysicalDeviceBufferDeviceAddressFeatures
	err := extension.GetPhysicalDeviceFeatures2(
		physicalDevice,
		&khr_get_physical_device_properties2.PhysicalDeviceFeatures2{
			NextOutData: common.NextOutData{&outData},
		})
	require.NoError(t, err)
	require.Equal(t, khr_buffer_device_address.PhysicalDeviceBufferDeviceAddressFeatures{
		BufferDeviceAddressCaptureReplay: true,
	}, outData)
}
