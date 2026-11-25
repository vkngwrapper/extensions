package ext_private_data_test

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
	"github.com/vkngwrapper/extensions/v3/ext_private_data"
	ext_private_data_loader "github.com/vkngwrapper/extensions/v3/ext_private_data/loader"
	mock_private_data "github.com/vkngwrapper/extensions/v3/ext_private_data/mocks"
	"go.uber.org/mock/gomock"
)

func TestCreatePrivateDataSlot(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})

	extDriver := mock_private_data.NewMockLoader(ctrl)
	extension := ext_private_data.CreateExtensionDriverFromLoader(extDriver, device)

	privateDataHandle := ext_private_data_loader.VkPrivateDataSlot(unsafe.Pointer(mocks.NewFakeInstanceHandle()))

	extDriver.EXPECT().VkCreatePrivateDataSlotEXT(device.Handle(), gomock.Not(gomock.Nil()), gomock.Nil(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(device loader.VkDevice, pCreateInfo *ext_private_data_loader.VkPrivateDataSlotCreateInfoEXT, pAllocator *loader.VkAllocationCallbacks, pPrivateDataSlot *ext_private_data_loader.VkPrivateDataSlot) (common.VkResult, error) {
			info := reflect.ValueOf(pCreateInfo).Elem()
			require.Equal(t, uint64(0x3b9f4a5a), info.FieldByName("sType").Uint())
			require.True(t, info.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(1), info.FieldByName("flags").Uint())

			*pPrivateDataSlot = ext_private_data_loader.VkPrivateDataSlot(privateDataHandle)
			return core1_0.VKSuccess, nil
		})

	actualPrivateDataSlot, res, err := extension.CreatePrivateDataSlot(nil, ext_private_data.PrivateDataSlotCreateInfo{
		Flags: 1,
	})
	require.NoError(t, err)
	require.Equal(t, core1_0.VKSuccess, res)
	require.Equal(t, privateDataHandle, actualPrivateDataSlot.Handle())
}

func TestDestroyPrivateDataSlot(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})

	extDriver := mock_private_data.NewMockLoader(ctrl)
	extension := ext_private_data.CreateExtensionDriverFromLoader(extDriver, device)

	privateDataHandle := ext_private_data_loader.VkPrivateDataSlot(unsafe.Pointer(mocks.NewFakeInstanceHandle()))
	privateDataSlot := ext_private_data.InternalPrivateDataSlot(device.Handle(), privateDataHandle, common.Vulkan1_0)

	extDriver.EXPECT().VkDestroyPrivateDataSlotEXT(device.Handle(), privateDataHandle, nil)

	extension.DestroyPrivateDataSlot(privateDataSlot, nil)
}

func TestGetPrivateData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	buffer := mocks.NewDummyBuffer(device)

	extDriver := mock_private_data.NewMockLoader(ctrl)
	extension := ext_private_data.CreateExtensionDriverFromLoader(extDriver, device)

	privateDataHandle := ext_private_data_loader.VkPrivateDataSlot(unsafe.Pointer(mocks.NewFakeInstanceHandle()))
	privateDataSlot := ext_private_data.InternalPrivateDataSlot(device.Handle(), privateDataHandle, common.Vulkan1_0)

	var expectedData uint64 = 0xdeadbeef

	extDriver.EXPECT().VkGetPrivateDataEXT(device.Handle(), loader.VkObjectType(0x9), loader.Uint64(buffer.Handle()), privateDataHandle, gomock.Not(gomock.Nil())).
		DoAndReturn(func(device loader.VkDevice, objectType loader.VkObjectType, objectHandle loader.Uint64, privateDataSlot ext_private_data_loader.VkPrivateDataSlot, data *loader.Uint64) {
			*data = loader.Uint64(expectedData)
		})

	data := extension.GetPrivateData(core1_0.ObjectTypeBuffer, loader.VulkanHandle(buffer.Handle()), privateDataSlot)
	require.Equal(t, expectedData, data)
}

func TestSetPrivateData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	buffer := mocks.NewDummyBuffer(device)

	extDriver := mock_private_data.NewMockLoader(ctrl)
	extension := ext_private_data.CreateExtensionDriverFromLoader(extDriver, device)

	privateDataHandle := ext_private_data_loader.VkPrivateDataSlot(unsafe.Pointer(mocks.NewFakeInstanceHandle()))
	privateDataSlot := ext_private_data.InternalPrivateDataSlot(device.Handle(), privateDataHandle, common.Vulkan1_0)

	var expectedData uint64 = 0xdeadbeef

	extDriver.EXPECT().VkSetPrivateDataEXT(device.Handle(), loader.VkObjectType(0x9), loader.Uint64(buffer.Handle()), privateDataHandle, loader.Uint64(expectedData))

	res, err := extension.SetPrivateData(core1_0.ObjectTypeBuffer, loader.VulkanHandle(buffer.Handle()), privateDataSlot, expectedData)
	require.NoError(t, err)
	require.Equal(t, res, core1_0.VKSuccess)
}

func TestPhysicalDevicePrivateDataFeatures_OutData(t *testing.T) {
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

		pInnerFeatures := (*ext_private_data_loader.VkPhysicalDevicePrivateDataFeaturesEXT)(info.FieldByName("pNext").UnsafePointer())
		features := reflect.ValueOf(pInnerFeatures).Elem()
		require.Equal(t, uint64(0x3b9f4a58), features.FieldByName("sType").Uint())
		require.True(t, features.FieldByName("pNext").IsNil())
		*(*uint64)(unsafe.Pointer(features.FieldByName("privateData").UnsafeAddr())) = uint64(1)
	})

	var baseFeatures core1_1.PhysicalDeviceFeatures2
	var features ext_private_data.PhysicalDevicePrivateDataFeatures

	baseFeatures.NextOutData.Next = &features

	err := driver.GetPhysicalDeviceFeatures2(physicalDevice, &baseFeatures)
	require.NoError(t, err)

	require.Equal(t, ext_private_data.PhysicalDevicePrivateDataFeatures{
		PrivateData: true,
	}, features)
}

func TestPhysicalDevicePrivateDataFeatures_Options(t *testing.T) {
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

		pInnerFeatures := (*ext_private_data_loader.VkPhysicalDevicePrivateDataFeaturesEXT)(info.FieldByName("pNext").UnsafePointer())
		features := reflect.ValueOf(pInnerFeatures).Elem()

		require.Equal(t, uint64(0x3b9f4a58), features.FieldByName("sType").Uint())
		require.True(t, features.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1), features.FieldByName("privateData").Uint())

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
		NextOptions: common.NextOptions{Next: ext_private_data.PhysicalDevicePrivateDataFeatures{
			PrivateData: true,
		}},
	})
	require.NoError(t, err)
	require.Equal(t, core1_0.VKSuccess, res)
	require.NotNil(t, device)
	require.Equal(t, expectedDevice.Handle(), device.Device().Handle())
}

func TestDevicePrivateDataCreateInfo_Options(t *testing.T) {
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

		pInnerFeatures := (*ext_private_data_loader.VkDevicePrivateDataCreateInfoEXT)(info.FieldByName("pNext").UnsafePointer())
		features := reflect.ValueOf(pInnerFeatures).Elem()

		require.Equal(t, uint64(0x3b9f4a59), features.FieldByName("sType").Uint())
		require.True(t, features.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(7), features.FieldByName("privateDataSlotRequestCount").Uint())

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
		NextOptions: common.NextOptions{Next: ext_private_data.DevicePrivateDataCreateInfo{
			PrivateDataSlotRequestCount: 7,
		}},
	})
	require.NoError(t, err)
	require.Equal(t, core1_0.VKSuccess, res)
	require.NotNil(t, device)
	require.Equal(t, expectedDevice.Handle(), device.Device().Handle())
}
