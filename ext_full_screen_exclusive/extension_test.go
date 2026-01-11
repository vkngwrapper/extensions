//go:build windows

package ext_full_screen_exclusive_test

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
	mock_loader "github.com/vkngwrapper/core/v3/loader/mocks"
	"github.com/vkngwrapper/core/v3/mocks"
	"github.com/vkngwrapper/core/v3/mocks/mocks1_0"
	"github.com/vkngwrapper/extensions/v3/ext_full_screen_exclusive"
	mock_full_screen_exclusive "github.com/vkngwrapper/extensions/v3/ext_full_screen_exclusive/mocks"
	"github.com/vkngwrapper/extensions/v3/khr_device_group"
	khr_device_group_loader "github.com/vkngwrapper/extensions/v3/khr_device_group/loader"
	"github.com/vkngwrapper/extensions/v3/khr_get_physical_device_properties2"
	"github.com/vkngwrapper/extensions/v3/khr_get_surface_capabilities2"
	khr_get_surface_capabilities2_loader "github.com/vkngwrapper/extensions/v3/khr_get_surface_capabilities2/loader"
	"github.com/vkngwrapper/extensions/v3/khr_surface"
	khr_surface_loader "github.com/vkngwrapper/extensions/v3/khr_surface/loader"
	mock_surface "github.com/vkngwrapper/extensions/v3/khr_surface/mocks"
	"github.com/vkngwrapper/extensions/v3/khr_swapchain"
	"go.uber.org/mock/gomock"
)

func TestVulkanExtensionDriver_AcquireFullScreenExclusiveMode(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	swapchain := khr_swapchain.NewDummySwapchain(device)

	mockLoader := mock_full_screen_exclusive.NewMockLoader(ctrl)
	driver := ext_full_screen_exclusive.CreateExtensionDriverFromLoader(mockLoader, device, false)

	mockLoader.EXPECT().VkAcquireFullScreenExclusiveModeEXT(
		device.Handle(),
		swapchain.Handle(),
	).Return(core1_0.VKSuccess, nil)

	res, err := driver.AcquireFullScreenExclusiveMode(swapchain)
	require.NoError(t, err)
	require.Equal(t, core1_0.VKSuccess, res)
}

func TestVulkanExtensionDriver_ReleaseFullScreenExclusiveMode(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	swapchain := khr_swapchain.NewDummySwapchain(device)

	mockLoader := mock_full_screen_exclusive.NewMockLoader(ctrl)
	driver := ext_full_screen_exclusive.CreateExtensionDriverFromLoader(mockLoader, device, false)

	mockLoader.EXPECT().VkReleaseFullScreenExclusiveModeEXT(
		device.Handle(),
		swapchain.Handle(),
	).Return(core1_0.VKSuccess, nil)

	res, err := driver.ReleaseFullScreenExclusiveMode(swapchain)
	require.NoError(t, err)
	require.Equal(t, core1_0.VKSuccess, res)
}

func TestVulkanExtensionDriver_GetPhysicalDeviceSurfacePresentModes2(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	instance := mocks.NewDummyInstance(common.Vulkan1_0, []string{})
	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	surface := mock_surface.NewDummySurface(instance)
	mockLoader := mock_full_screen_exclusive.NewMockLoader(ctrl)
	driver := ext_full_screen_exclusive.CreateExtensionDriverFromLoader(mockLoader, device, false)

	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_0)

	mockLoader.EXPECT().VkGetPhysicalDeviceSurfacePresentModes2EXT(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
	).DoAndReturn(func(
		physicalDevice loader.VkPhysicalDevice,
		pSurfaceInfo *khr_get_surface_capabilities2_loader.VkPhysicalDeviceSurfaceInfo2KHR,
		pPresentModeCount *loader.Uint32,
		pPresentModes *khr_surface_loader.VkPresentModeKHR,
	) (common.VkResult, error) {
		input := reflect.ValueOf(*pSurfaceInfo)

		require.Equal(t, uint64(1000119000), input.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SURFACE_INFO_2_KHR
		require.Equal(t, surface.Handle(), khr_surface_loader.VkSurfaceKHR(input.FieldByName("surface").UnsafePointer()))
		require.True(t, input.FieldByName("pNext").IsNil())

		*pPresentModeCount = 2

		return core1_0.VKSuccess, nil
	})

	mockLoader.EXPECT().VkGetPhysicalDeviceSurfacePresentModes2EXT(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(
		physicalDevice loader.VkPhysicalDevice,
		pSurfaceInfo *khr_get_surface_capabilities2_loader.VkPhysicalDeviceSurfaceInfo2KHR,
		pPresentModeCount *loader.Uint32,
		pPresentModes *khr_surface_loader.VkPresentModeKHR,
	) (common.VkResult, error) {
		input := reflect.ValueOf(*pSurfaceInfo)

		require.Equal(t, uint64(1000119000), input.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SURFACE_INFO_2_KHR
		require.Equal(t, surface.Handle(), khr_surface_loader.VkSurfaceKHR(input.FieldByName("surface").UnsafePointer()))
		require.True(t, input.FieldByName("pNext").IsNil())

		require.Equal(t, loader.Uint32(2), *pPresentModeCount)

		presentModeSlice := unsafe.Slice(pPresentModes, 2)
		presentModeSlice[0] = khr_surface_loader.VkPresentModeKHR(0)
		presentModeSlice[1] = khr_surface_loader.VkPresentModeKHR(3)

		return core1_0.VKSuccess, nil
	})

	modes, res, err := driver.GetPhysicalDeviceSurfacePresentModes2(
		physicalDevice,
		khr_get_surface_capabilities2.PhysicalDeviceSurfaceInfo2{
			Surface: surface,
		})
	require.NoError(t, err)
	require.Equal(t, core1_0.VKSuccess, res)
	require.Equal(t, []khr_surface.PresentMode{
		khr_surface.PresentModeImmediate,
		khr_surface.PresentModeFIFORelaxed,
	}, modes)
}

func TestVulkanExtensionDriverWithDeviceGroups_GetDeviceGroupSurfacePresentModes2(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	instance := mocks.NewDummyInstance(common.Vulkan1_0, []string{})
	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	surface := mock_surface.NewDummySurface(instance)
	mockLoader := mock_full_screen_exclusive.NewMockLoader(ctrl)
	driver := ext_full_screen_exclusive.CreateExtensionDriverFromLoader(mockLoader, device, true)
	devGroupDriver, ok := driver.(ext_full_screen_exclusive.ExtensionDriverWithDeviceGroups)
	require.True(t, ok)

	mockLoader.EXPECT().VkGetDeviceGroupSurfacePresentModes2EXT(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device loader.VkDevice, pSurfaceInfo *khr_get_surface_capabilities2_loader.VkPhysicalDeviceSurfaceInfo2KHR, pModes *khr_device_group_loader.VkDeviceGroupPresentModeFlagsKHR) (common.VkResult, error) {
		input := reflect.ValueOf(*pSurfaceInfo)

		require.Equal(t, uint64(1000119000), input.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SURFACE_INFO_2_KHR
		require.Equal(t, surface.Handle(), khr_surface_loader.VkSurfaceKHR(input.FieldByName("surface").UnsafePointer()))
		require.True(t, input.FieldByName("pNext").IsNil())

		*pModes = khr_device_group_loader.VkDeviceGroupPresentModeFlagsKHR(4)
		return core1_0.VKSuccess, nil
	})

	var flags khr_device_group.DeviceGroupPresentModeFlags
	res, err := devGroupDriver.GetDeviceGroupSurfacePresentModes2(
		khr_get_surface_capabilities2.PhysicalDeviceSurfaceInfo2{
			Surface: surface,
		},
		&flags)
	require.NoError(t, err)
	require.Equal(t, core1_0.VKSuccess, res)
	require.Equal(t, khr_device_group.DeviceGroupPresentModeSum, flags)
}

func TestCreateExtensionDriverFromCoreDriver_SuccessWithDeviceGroups(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	instance := mocks.NewDummyInstance(common.Vulkan1_0, []string{
		khr_get_physical_device_properties2.ExtensionName,
		khr_surface.ExtensionName,
		khr_get_surface_capabilities2.ExtensionName,
	})
	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{
		ext_full_screen_exclusive.ExtensionName,
		khr_swapchain.ExtensionName,
		khr_device_group.ExtensionName,
	})
	mockLoader := mock_loader.NewMockLoader(ctrl)
	mockLoader.EXPECT().LoadProcAddr(gomock.Any()).Return(unsafe.Pointer(nil)).AnyTimes()

	driver := mocks1_0.InternalCoreDriver(instance, device, mockLoader)
	extDriver := ext_full_screen_exclusive.CreateExtensionDriverFromCoreDriver(driver, instance)
	require.NotNil(t, extDriver)

	extDriverDeviceGroup, ok := extDriver.(ext_full_screen_exclusive.ExtensionDriverWithDeviceGroups)
	require.True(t, ok)
	require.NotNil(t, extDriverDeviceGroup)
}

func TestCreateExtensionDriverFromCoreDriver_SuccessWithoutDeviceGroups(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	instance := mocks.NewDummyInstance(common.Vulkan1_0, []string{
		khr_get_physical_device_properties2.ExtensionName,
		khr_surface.ExtensionName,
		khr_get_surface_capabilities2.ExtensionName,
	})
	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{
		ext_full_screen_exclusive.ExtensionName,
		khr_swapchain.ExtensionName,
	})
	mockLoader := mock_loader.NewMockLoader(ctrl)
	mockLoader.EXPECT().LoadProcAddr(gomock.Any()).Return(unsafe.Pointer(nil)).AnyTimes()

	driver := mocks1_0.InternalCoreDriver(instance, device, mockLoader)
	extDriver := ext_full_screen_exclusive.CreateExtensionDriverFromCoreDriver(driver, instance)
	require.NotNil(t, extDriver)

	_, ok := extDriver.(ext_full_screen_exclusive.ExtensionDriverWithDeviceGroups)
	require.False(t, ok)
}

func TestCreateExtensionDriverFromCoreDriver_MissingDeps(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	instance := mocks.NewDummyInstance(common.Vulkan1_0, []string{
		khr_get_physical_device_properties2.ExtensionName,
		khr_surface.ExtensionName,
	})
	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{
		ext_full_screen_exclusive.ExtensionName,
		khr_swapchain.ExtensionName,
		khr_device_group.ExtensionName,
	})
	mockLoader := mock_loader.NewMockLoader(ctrl)
	mockLoader.EXPECT().LoadProcAddr(gomock.Any()).Return(unsafe.Pointer(nil)).AnyTimes()

	driver := mocks1_0.InternalCoreDriver(instance, device, mockLoader)
	extDriver := ext_full_screen_exclusive.CreateExtensionDriverFromCoreDriver(driver, instance)
	require.Nil(t, extDriver)
}

func TestCreateExtensionDriverFromCoreDriver_MissingExt(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	instance := mocks.NewDummyInstance(common.Vulkan1_0, []string{
		khr_get_physical_device_properties2.ExtensionName,
		khr_surface.ExtensionName,
		khr_get_surface_capabilities2.ExtensionName,
	})
	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{
		khr_swapchain.ExtensionName,
		khr_device_group.ExtensionName,
	})
	mockLoader := mock_loader.NewMockLoader(ctrl)
	mockLoader.EXPECT().LoadProcAddr(gomock.Any()).Return(unsafe.Pointer(nil)).AnyTimes()

	driver := mocks1_0.InternalCoreDriver(instance, device, mockLoader)
	extDriver := ext_full_screen_exclusive.CreateExtensionDriverFromCoreDriver(driver, instance)
	require.Nil(t, extDriver)
}

func TestCreateExtensionDriverFromCoreDriver_HasVulkan1_1(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	instance := mocks.NewDummyInstance(common.Vulkan1_1, []string{
		khr_surface.ExtensionName,
		khr_get_surface_capabilities2.ExtensionName,
	})
	device := mocks.NewDummyDevice(common.Vulkan1_1, []string{
		ext_full_screen_exclusive.ExtensionName,
		khr_swapchain.ExtensionName,
		khr_device_group.ExtensionName,
	})
	mockLoader := mock_loader.NewMockLoader(ctrl)
	mockLoader.EXPECT().LoadProcAddr(gomock.Any()).Return(unsafe.Pointer(nil)).AnyTimes()

	driver := mocks1_0.InternalCoreDriver(instance, device, mockLoader)
	extDriver := ext_full_screen_exclusive.CreateExtensionDriverFromCoreDriver(driver, instance)
	require.NotNil(t, extDriver)

	extDriverDeviceGroup, ok := extDriver.(ext_full_screen_exclusive.ExtensionDriverWithDeviceGroups)
	require.True(t, ok)
	require.NotNil(t, extDriverDeviceGroup)
}
