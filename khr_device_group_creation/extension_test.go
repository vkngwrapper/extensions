package khr_device_group_creation_test

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
	mock_driver "github.com/vkngwrapper/core/v3/loader/mocks"
	"github.com/vkngwrapper/core/v3/mocks"
	"github.com/vkngwrapper/core/v3/mocks/mocks1_0"
	"github.com/vkngwrapper/extensions/v3/khr_device_group_creation"
	khr_device_group_creation_driver "github.com/vkngwrapper/extensions/v3/khr_device_group_creation/loader"
	mock_device_group_creation "github.com/vkngwrapper/extensions/v3/khr_device_group_creation/mocks"
	"go.uber.org/mock/gomock"
)

func TestVulkanExtension_EnumeratePhysicalDeviceGroups(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreLoader := mock_driver.LoaderForVersion(ctrl, common.Vulkan1_0)
	instance := mocks.NewDummyInstance(common.Vulkan1_0, []string{})

	physicalDevice1 := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_0)
	physicalDevice2 := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_0)
	physicalDevice3 := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_0)
	physicalDevice4 := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_0)
	physicalDevice5 := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_0)
	physicalDevice6 := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_0)

	extDriver := mock_device_group_creation.NewMockLoader(ctrl)
	extension := khr_device_group_creation.CreateExtensionFromDriver(extDriver)

	extDriver.EXPECT().VkEnumeratePhysicalDeviceGroupsKHR(
		instance.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
	).DoAndReturn(func(instance loader.VkInstance, pCount *loader.Uint32, pProperties *khr_device_group_creation_driver.VkPhysicalDeviceGroupPropertiesKHR) (common.VkResult, error) {
		*pCount = loader.Uint32(3)

		return core1_0.VKSuccess, nil
	})

	extDriver.EXPECT().VkEnumeratePhysicalDeviceGroupsKHR(
		instance.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(instance loader.VkInstance, pCount *loader.Uint32, pProperties *khr_device_group_creation_driver.VkPhysicalDeviceGroupPropertiesKHR) (common.VkResult, error) {
		require.Equal(t, loader.Uint32(3), *pCount)

		propertySlice := ([]khr_device_group_creation_driver.VkPhysicalDeviceGroupPropertiesKHR)(unsafe.Slice(pProperties, 3))
		props := reflect.ValueOf(propertySlice)
		prop := props.Index(0)
		require.Equal(t, uint64(1000070000), prop.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_GROUP_PROPERTIES_KHR
		require.True(t, prop.FieldByName("pNext").IsNil())
		*(*loader.Uint32)(unsafe.Pointer(prop.FieldByName("physicalDeviceCount").UnsafeAddr())) = loader.Uint32(1)
		*(*loader.VkBool32)(unsafe.Pointer(prop.FieldByName("subsetAllocation").UnsafeAddr())) = loader.VkBool32(1)

		propDevices := prop.FieldByName("physicalDevices")
		*(*loader.VkPhysicalDevice)(unsafe.Pointer(propDevices.Index(0).UnsafeAddr())) = physicalDevice1.Handle()

		prop = props.Index(1)
		require.Equal(t, uint64(1000070000), prop.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_GROUP_PROPERTIES_KHR
		require.True(t, prop.FieldByName("pNext").IsNil())
		*(*loader.Uint32)(unsafe.Pointer(prop.FieldByName("physicalDeviceCount").UnsafeAddr())) = loader.Uint32(2)
		*(*loader.VkBool32)(unsafe.Pointer(prop.FieldByName("subsetAllocation").UnsafeAddr())) = loader.VkBool32(0)

		propDevices = prop.FieldByName("physicalDevices")
		*(*loader.VkPhysicalDevice)(unsafe.Pointer(propDevices.Index(0).UnsafeAddr())) = physicalDevice2.Handle()
		*(*loader.VkPhysicalDevice)(unsafe.Pointer(propDevices.Index(1).UnsafeAddr())) = physicalDevice3.Handle()

		prop = props.Index(2)
		require.Equal(t, uint64(1000070000), prop.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_GROUP_PROPERTIES_KHR
		require.True(t, prop.FieldByName("pNext").IsNil())
		*(*loader.Uint32)(unsafe.Pointer(prop.FieldByName("physicalDeviceCount").UnsafeAddr())) = loader.Uint32(3)
		*(*loader.VkBool32)(unsafe.Pointer(prop.FieldByName("subsetAllocation").UnsafeAddr())) = loader.VkBool32(1)

		propDevices = prop.FieldByName("physicalDevices")
		*(*loader.VkPhysicalDevice)(unsafe.Pointer(propDevices.Index(0).UnsafeAddr())) = physicalDevice4.Handle()
		*(*loader.VkPhysicalDevice)(unsafe.Pointer(propDevices.Index(1).UnsafeAddr())) = physicalDevice5.Handle()
		*(*loader.VkPhysicalDevice)(unsafe.Pointer(propDevices.Index(2).UnsafeAddr())) = physicalDevice6.Handle()

		return core1_0.VKSuccess, nil
	})

	coreLoader.EXPECT().VkGetPhysicalDeviceProperties(physicalDevice1.Handle(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(device loader.VkPhysicalDevice, pProperties *loader.VkPhysicalDeviceProperties) {
			value := reflect.ValueOf(pProperties).Elem()
			*(*loader.Uint32)(unsafe.Pointer(value.FieldByName("apiVersion").UnsafeAddr())) = loader.Uint32(common.Vulkan1_0)
		})

	coreLoader.EXPECT().VkGetPhysicalDeviceProperties(physicalDevice2.Handle(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(device loader.VkPhysicalDevice, pProperties *loader.VkPhysicalDeviceProperties) {
			value := reflect.ValueOf(pProperties).Elem()
			*(*loader.Uint32)(unsafe.Pointer(value.FieldByName("apiVersion").UnsafeAddr())) = loader.Uint32(common.Vulkan1_0)
		})

	coreLoader.EXPECT().VkGetPhysicalDeviceProperties(physicalDevice3.Handle(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(device loader.VkPhysicalDevice, pProperties *loader.VkPhysicalDeviceProperties) {
			value := reflect.ValueOf(pProperties).Elem()
			*(*loader.Uint32)(unsafe.Pointer(value.FieldByName("apiVersion").UnsafeAddr())) = loader.Uint32(common.Vulkan1_0)
		})

	coreLoader.EXPECT().VkGetPhysicalDeviceProperties(physicalDevice4.Handle(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(device loader.VkPhysicalDevice, pProperties *loader.VkPhysicalDeviceProperties) {
			value := reflect.ValueOf(pProperties).Elem()
			*(*loader.Uint32)(unsafe.Pointer(value.FieldByName("apiVersion").UnsafeAddr())) = loader.Uint32(common.Vulkan1_0)
		})

	coreLoader.EXPECT().VkGetPhysicalDeviceProperties(physicalDevice5.Handle(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(device loader.VkPhysicalDevice, pProperties *loader.VkPhysicalDeviceProperties) {
			value := reflect.ValueOf(pProperties).Elem()
			*(*loader.Uint32)(unsafe.Pointer(value.FieldByName("apiVersion").UnsafeAddr())) = loader.Uint32(common.Vulkan1_0)
		})

	coreLoader.EXPECT().VkGetPhysicalDeviceProperties(physicalDevice6.Handle(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(device loader.VkPhysicalDevice, pProperties *loader.VkPhysicalDeviceProperties) {
			value := reflect.ValueOf(pProperties).Elem()
			*(*loader.Uint32)(unsafe.Pointer(value.FieldByName("apiVersion").UnsafeAddr())) = loader.Uint32(common.Vulkan1_0)
		})

	groups, _, err := extension.EnumeratePhysicalDeviceGroups(instance, nil)
	require.NoError(t, err)
	require.Len(t, groups, 3)
	require.True(t, groups[0].SubsetAllocation)
	require.Len(t, groups[0].PhysicalDevices, 1)
	require.Equal(t, physicalDevice1.Handle(), groups[0].PhysicalDevices[0].Handle())

	require.False(t, groups[1].SubsetAllocation)
	require.Len(t, groups[1].PhysicalDevices, 2)
	require.Equal(t, physicalDevice2.Handle(), groups[1].PhysicalDevices[0].Handle())
	require.Equal(t, physicalDevice3.Handle(), groups[1].PhysicalDevices[1].Handle())

	require.True(t, groups[2].SubsetAllocation)
	require.Len(t, groups[2].PhysicalDevices, 3)
	require.Equal(t, physicalDevice4.Handle(), groups[2].PhysicalDevices[0].Handle())
	require.Equal(t, physicalDevice5.Handle(), groups[2].PhysicalDevices[1].Handle())
	require.Equal(t, physicalDevice6.Handle(), groups[2].PhysicalDevices[2].Handle())
}

func TestVulkanExtension_EnumeratePhysicalDeviceGroups_Incomplete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreLoader := mock_driver.LoaderForVersion(ctrl, common.Vulkan1_0)
	instance := mocks.NewDummyInstance(common.Vulkan1_0, []string{})

	physicalDevice1 := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_0)
	physicalDevice2 := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_0)
	physicalDevice3 := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_0)
	physicalDevice4 := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_0)
	physicalDevice5 := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_0)
	physicalDevice6 := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_0)

	extDriver := mock_device_group_creation.NewMockLoader(ctrl)
	extension := khr_device_group_creation.CreateExtensionFromDriver(extDriver)

	extDriver.EXPECT().VkEnumeratePhysicalDeviceGroupsKHR(
		instance.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
	).DoAndReturn(func(instance loader.VkInstance, pCount *loader.Uint32, pProperties *khr_device_group_creation_driver.VkPhysicalDeviceGroupPropertiesKHR) (common.VkResult, error) {
		*pCount = loader.Uint32(2)

		return core1_0.VKSuccess, nil
	})

	extDriver.EXPECT().VkEnumeratePhysicalDeviceGroupsKHR(
		instance.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(instance loader.VkInstance, pCount *loader.Uint32, pProperties *khr_device_group_creation_driver.VkPhysicalDeviceGroupPropertiesKHR) (common.VkResult, error) {
		require.Equal(t, loader.Uint32(2), *pCount)

		propertySlice := ([]khr_device_group_creation_driver.VkPhysicalDeviceGroupPropertiesKHR)(unsafe.Slice(pProperties, 3))
		props := reflect.ValueOf(propertySlice)
		prop := props.Index(0)
		require.Equal(t, uint64(1000070000), prop.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_GROUP_PROPERTIES_KHR
		require.True(t, prop.FieldByName("pNext").IsNil())
		*(*loader.Uint32)(unsafe.Pointer(prop.FieldByName("physicalDeviceCount").UnsafeAddr())) = loader.Uint32(1)
		*(*loader.VkBool32)(unsafe.Pointer(prop.FieldByName("subsetAllocation").UnsafeAddr())) = loader.VkBool32(1)

		propDevices := prop.FieldByName("physicalDevices")
		*(*loader.VkPhysicalDevice)(unsafe.Pointer(propDevices.Index(0).UnsafeAddr())) = physicalDevice1.Handle()

		prop = props.Index(1)
		require.Equal(t, uint64(1000070000), prop.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_GROUP_PROPERTIES_KHR
		require.True(t, prop.FieldByName("pNext").IsNil())
		*(*loader.Uint32)(unsafe.Pointer(prop.FieldByName("physicalDeviceCount").UnsafeAddr())) = loader.Uint32(2)
		*(*loader.VkBool32)(unsafe.Pointer(prop.FieldByName("subsetAllocation").UnsafeAddr())) = loader.VkBool32(0)

		propDevices = prop.FieldByName("physicalDevices")
		*(*loader.VkPhysicalDevice)(unsafe.Pointer(propDevices.Index(0).UnsafeAddr())) = physicalDevice2.Handle()
		*(*loader.VkPhysicalDevice)(unsafe.Pointer(propDevices.Index(1).UnsafeAddr())) = physicalDevice3.Handle()

		return core1_0.VKIncomplete, nil
	})

	extDriver.EXPECT().VkEnumeratePhysicalDeviceGroupsKHR(
		instance.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
	).DoAndReturn(func(instance loader.VkInstance, pCount *loader.Uint32, pProperties *khr_device_group_creation_driver.VkPhysicalDeviceGroupPropertiesKHR) (common.VkResult, error) {
		*pCount = loader.Uint32(3)

		return core1_0.VKSuccess, nil
	})

	extDriver.EXPECT().VkEnumeratePhysicalDeviceGroupsKHR(
		instance.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(instance loader.VkInstance, pCount *loader.Uint32, pProperties *khr_device_group_creation_driver.VkPhysicalDeviceGroupPropertiesKHR) (common.VkResult, error) {
		require.Equal(t, loader.Uint32(3), *pCount)

		propertySlice := ([]khr_device_group_creation_driver.VkPhysicalDeviceGroupPropertiesKHR)(unsafe.Slice(pProperties, 3))
		props := reflect.ValueOf(propertySlice)
		prop := props.Index(0)
		require.Equal(t, uint64(1000070000), prop.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_GROUP_PROPERTIES_KHR
		require.True(t, prop.FieldByName("pNext").IsNil())
		*(*loader.Uint32)(unsafe.Pointer(prop.FieldByName("physicalDeviceCount").UnsafeAddr())) = loader.Uint32(1)
		*(*loader.VkBool32)(unsafe.Pointer(prop.FieldByName("subsetAllocation").UnsafeAddr())) = loader.VkBool32(1)

		propDevices := prop.FieldByName("physicalDevices")
		*(*loader.VkPhysicalDevice)(unsafe.Pointer(propDevices.Index(0).UnsafeAddr())) = physicalDevice1.Handle()

		prop = props.Index(1)
		require.Equal(t, uint64(1000070000), prop.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_GROUP_PROPERTIES_KHR
		require.True(t, prop.FieldByName("pNext").IsNil())
		*(*loader.Uint32)(unsafe.Pointer(prop.FieldByName("physicalDeviceCount").UnsafeAddr())) = loader.Uint32(2)
		*(*loader.VkBool32)(unsafe.Pointer(prop.FieldByName("subsetAllocation").UnsafeAddr())) = loader.VkBool32(0)

		propDevices = prop.FieldByName("physicalDevices")
		*(*loader.VkPhysicalDevice)(unsafe.Pointer(propDevices.Index(0).UnsafeAddr())) = physicalDevice2.Handle()
		*(*loader.VkPhysicalDevice)(unsafe.Pointer(propDevices.Index(1).UnsafeAddr())) = physicalDevice3.Handle()

		prop = props.Index(2)
		require.Equal(t, uint64(1000070000), prop.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_GROUP_PROPERTIES_KHR
		require.True(t, prop.FieldByName("pNext").IsNil())
		*(*loader.Uint32)(unsafe.Pointer(prop.FieldByName("physicalDeviceCount").UnsafeAddr())) = loader.Uint32(3)
		*(*loader.VkBool32)(unsafe.Pointer(prop.FieldByName("subsetAllocation").UnsafeAddr())) = loader.VkBool32(1)

		propDevices = prop.FieldByName("physicalDevices")
		*(*loader.VkPhysicalDevice)(unsafe.Pointer(propDevices.Index(0).UnsafeAddr())) = physicalDevice4.Handle()
		*(*loader.VkPhysicalDevice)(unsafe.Pointer(propDevices.Index(1).UnsafeAddr())) = physicalDevice5.Handle()
		*(*loader.VkPhysicalDevice)(unsafe.Pointer(propDevices.Index(2).UnsafeAddr())) = physicalDevice6.Handle()

		return core1_0.VKSuccess, nil
	})

	coreLoader.EXPECT().VkGetPhysicalDeviceProperties(physicalDevice1.Handle(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(device loader.VkPhysicalDevice, pProperties *loader.VkPhysicalDeviceProperties) {
			value := reflect.ValueOf(pProperties).Elem()
			*(*loader.Uint32)(unsafe.Pointer(value.FieldByName("apiVersion").UnsafeAddr())) = loader.Uint32(common.Vulkan1_0)
		}).Times(2)

	coreLoader.EXPECT().VkGetPhysicalDeviceProperties(physicalDevice2.Handle(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(device loader.VkPhysicalDevice, pProperties *loader.VkPhysicalDeviceProperties) {
			value := reflect.ValueOf(pProperties).Elem()
			*(*loader.Uint32)(unsafe.Pointer(value.FieldByName("apiVersion").UnsafeAddr())) = loader.Uint32(common.Vulkan1_0)
		}).Times(2)

	coreLoader.EXPECT().VkGetPhysicalDeviceProperties(physicalDevice3.Handle(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(device loader.VkPhysicalDevice, pProperties *loader.VkPhysicalDeviceProperties) {
			value := reflect.ValueOf(pProperties).Elem()
			*(*loader.Uint32)(unsafe.Pointer(value.FieldByName("apiVersion").UnsafeAddr())) = loader.Uint32(common.Vulkan1_0)
		}).Times(2)

	coreLoader.EXPECT().VkGetPhysicalDeviceProperties(physicalDevice4.Handle(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(device loader.VkPhysicalDevice, pProperties *loader.VkPhysicalDeviceProperties) {
			value := reflect.ValueOf(pProperties).Elem()
			*(*loader.Uint32)(unsafe.Pointer(value.FieldByName("apiVersion").UnsafeAddr())) = loader.Uint32(common.Vulkan1_0)
		})

	coreLoader.EXPECT().VkGetPhysicalDeviceProperties(physicalDevice5.Handle(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(device loader.VkPhysicalDevice, pProperties *loader.VkPhysicalDeviceProperties) {
			value := reflect.ValueOf(pProperties).Elem()
			*(*loader.Uint32)(unsafe.Pointer(value.FieldByName("apiVersion").UnsafeAddr())) = loader.Uint32(common.Vulkan1_0)
		})

	coreLoader.EXPECT().VkGetPhysicalDeviceProperties(physicalDevice6.Handle(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(device loader.VkPhysicalDevice, pProperties *loader.VkPhysicalDeviceProperties) {
			value := reflect.ValueOf(pProperties).Elem()
			*(*loader.Uint32)(unsafe.Pointer(value.FieldByName("apiVersion").UnsafeAddr())) = loader.Uint32(common.Vulkan1_0)
		})

	groups, _, err := extension.EnumeratePhysicalDeviceGroups(instance, nil)
	require.NoError(t, err)
	require.Len(t, groups, 3)
	require.True(t, groups[0].SubsetAllocation)
	require.Len(t, groups[0].PhysicalDevices, 1)
	require.Equal(t, physicalDevice1.Handle(), groups[0].PhysicalDevices[0].Handle())

	require.False(t, groups[1].SubsetAllocation)
	require.Len(t, groups[1].PhysicalDevices, 2)
	require.Equal(t, physicalDevice2.Handle(), groups[1].PhysicalDevices[0].Handle())
	require.Equal(t, physicalDevice3.Handle(), groups[1].PhysicalDevices[1].Handle())

	require.True(t, groups[2].SubsetAllocation)
	require.Len(t, groups[2].PhysicalDevices, 3)
	require.Equal(t, physicalDevice4.Handle(), groups[2].PhysicalDevices[0].Handle())
	require.Equal(t, physicalDevice5.Handle(), groups[2].PhysicalDevices[1].Handle())
	require.Equal(t, physicalDevice6.Handle(), groups[2].PhysicalDevices[2].Handle())
}

func TestDeviceGroupOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreLoader := mock_driver.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := mocks1_0.InternalCoreInstanceDriver(coreLoader)
	instance := mocks.NewDummyInstance(common.Vulkan1_0, []string{})
	physicalDevice1 := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_0)

	physicalDevice2 := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_0)
	physicalDevice3 := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_0)

	handle := mocks.NewFakeDeviceHandle()

	coreLoader.EXPECT().VkCreateDevice(
		physicalDevice1.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice loader.VkPhysicalDevice, pCreateInfo *loader.VkDeviceCreateInfo, pAllocator *loader.VkAllocationCallbacks, pDevice *loader.VkDevice) (common.VkResult, error) {
		*pDevice = handle

		val := reflect.ValueOf(pCreateInfo).Elem()

		require.Equal(t, uint64(3), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_CREATE_INFO

		optionsPtr := (*khr_device_group_creation_driver.VkDeviceGroupDeviceCreateInfoKHR)(val.FieldByName("pNext").UnsafePointer())
		options := reflect.ValueOf(optionsPtr).Elem()

		require.Equal(t, uint64(1000070001), options.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_GROUP_DEVICE_CREATE_INFO_KHR
		require.True(t, options.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(3), options.FieldByName("physicalDeviceCount").Uint())

		devicePtr := (*loader.VkPhysicalDevice)(options.FieldByName("pPhysicalDevices").UnsafePointer())
		deviceSlice := ([]loader.VkPhysicalDevice)(unsafe.Slice(devicePtr, 3))
		require.Equal(t, physicalDevice1.Handle(), deviceSlice[0])
		require.Equal(t, physicalDevice2.Handle(), deviceSlice[1])
		require.Equal(t, physicalDevice3.Handle(), deviceSlice[2])

		return core1_0.VKSuccess, nil
	})

	device, _, err := driver.CreateDevice(physicalDevice1, nil, core1_0.DeviceCreateInfo{
		QueueCreateInfos: []core1_0.DeviceQueueCreateInfo{
			{
				QueuePriorities: []float32{0},
			},
		},
		NextOptions: common.NextOptions{Next: khr_device_group_creation.DeviceGroupDeviceCreateInfo{
			PhysicalDevices: []core.PhysicalDevice{physicalDevice1, physicalDevice2, physicalDevice3},
		}},
	})
	require.NoError(t, err)
	require.Equal(t, handle, device.Handle())
}
