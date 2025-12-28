package khr_surface_test

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
	"github.com/vkngwrapper/core/v3/mocks"
	"github.com/vkngwrapper/extensions/v3/khr_surface"
	khr_surface_driver "github.com/vkngwrapper/extensions/v3/khr_surface/loader"
	mock_surface "github.com/vkngwrapper/extensions/v3/khr_surface/mocks"
	"go.uber.org/mock/gomock"
)

func TestVulkanSurface_PresentModes(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	instance := mocks.NewDummyInstance(common.Vulkan1_0, []string{})
	surfaceDriver := mock_surface.NewMockLoader(ctrl)
	device := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_0)

	extension := khr_surface.CreateExtensionDriverFromLoader(surfaceDriver, instance)

	surface := mock_surface.NewDummySurface(instance)

	surfaceDriver.EXPECT().VkGetPhysicalDeviceSurfacePresentModesKHR(
		device.Handle(),
		surface.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil()).DoAndReturn(
		func(device loader.VkPhysicalDevice, surface khr_surface_driver.VkSurfaceKHR, pPresentModeCount *loader.Uint32, pPresentModes *khr_surface_driver.VkPresentModeKHR) (common.VkResult, error) {
			*pPresentModeCount = 2

			return core1_0.VKSuccess, nil
		})

	surfaceDriver.EXPECT().VkGetPhysicalDeviceSurfacePresentModesKHR(
		device.Handle(),
		surface.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil())).DoAndReturn(
		func(device loader.VkPhysicalDevice, surface khr_surface_driver.VkSurfaceKHR, pPresentModeCount *loader.Uint32, pPresentModes *khr_surface_driver.VkPresentModeKHR) (common.VkResult, error) {
			require.Equal(t, loader.Uint32(2), *pPresentModeCount)

			presentModeSlice := ([]khr_surface_driver.VkPresentModeKHR)(unsafe.Slice(pPresentModes, 2))
			presentModeSlice[0] = khr_surface_driver.VkPresentModeKHR(0) // VK_PRESENT_MODE_IMMEDIATE_KHR
			presentModeSlice[1] = khr_surface_driver.VkPresentModeKHR(3) // VK_PRESENT_MODE_FIFO_RELAXED_KHR

			return core1_0.VKSuccess, nil
		})

	presentModes, res, err := extension.GetPhysicalDeviceSurfacePresentModes(surface, device)
	require.Equal(t, core1_0.VKSuccess, res)
	require.NoError(t, err)
	require.Len(t, presentModes, 2)
	require.Equal(t, khr_surface.PresentModeImmediate, presentModes[0])
	require.Equal(t, khr_surface.PresentModeFIFORelaxed, presentModes[1])
}

func TestVulkanSurface_PresentModes_Incomplete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	instance := mocks.NewDummyInstance(common.Vulkan1_0, []string{})
	surfaceDriver := mock_surface.NewMockLoader(ctrl)
	device := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_0)

	extension := khr_surface.CreateExtensionDriverFromLoader(surfaceDriver, instance)

	surface := mock_surface.NewDummySurface(instance)

	surfaceDriver.EXPECT().VkGetPhysicalDeviceSurfacePresentModesKHR(
		device.Handle(),
		surface.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil()).DoAndReturn(
		func(device loader.VkPhysicalDevice, surface khr_surface_driver.VkSurfaceKHR, pPresentModeCount *loader.Uint32, pPresentModes *khr_surface_driver.VkPresentModeKHR) (common.VkResult, error) {
			*pPresentModeCount = 1

			return core1_0.VKSuccess, nil
		})

	surfaceDriver.EXPECT().VkGetPhysicalDeviceSurfacePresentModesKHR(
		device.Handle(),
		surface.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil())).DoAndReturn(
		func(device loader.VkPhysicalDevice, surface khr_surface_driver.VkSurfaceKHR, pPresentModeCount *loader.Uint32, pPresentModes *khr_surface_driver.VkPresentModeKHR) (common.VkResult, error) {
			require.Equal(t, loader.Uint32(1), *pPresentModeCount)

			presentModeSlice := ([]khr_surface_driver.VkPresentModeKHR)(unsafe.Slice(pPresentModes, 1))
			presentModeSlice[0] = khr_surface_driver.VkPresentModeKHR(0) // VK_PRESENT_MODE_IMMEDIATE_KHR

			return core1_0.VKIncomplete, nil
		})

	surfaceDriver.EXPECT().VkGetPhysicalDeviceSurfacePresentModesKHR(
		device.Handle(),
		surface.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil()).DoAndReturn(
		func(device loader.VkPhysicalDevice, surface khr_surface_driver.VkSurfaceKHR, pPresentModeCount *loader.Uint32, pPresentModes *khr_surface_driver.VkPresentModeKHR) (common.VkResult, error) {
			*pPresentModeCount = 2

			return core1_0.VKSuccess, nil
		})

	surfaceDriver.EXPECT().VkGetPhysicalDeviceSurfacePresentModesKHR(
		device.Handle(),
		surface.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil())).DoAndReturn(
		func(device loader.VkPhysicalDevice, surface khr_surface_driver.VkSurfaceKHR, pPresentModeCount *loader.Uint32, pPresentModes *khr_surface_driver.VkPresentModeKHR) (common.VkResult, error) {
			require.Equal(t, loader.Uint32(2), *pPresentModeCount)

			presentModeSlice := ([]khr_surface_driver.VkPresentModeKHR)(unsafe.Slice(pPresentModes, 2))
			presentModeSlice[0] = khr_surface_driver.VkPresentModeKHR(0) // VK_PRESENT_MODE_IMMEDIATE_KHR
			presentModeSlice[1] = khr_surface_driver.VkPresentModeKHR(3) // VK_PRESENT_MODE_FIFO_RELAXED_KHR

			return core1_0.VKSuccess, nil
		})

	presentModes, res, err := extension.GetPhysicalDeviceSurfacePresentModes(surface, device)
	require.Equal(t, core1_0.VKSuccess, res)
	require.NoError(t, err)
	require.Len(t, presentModes, 2)
	require.Equal(t, khr_surface.PresentModeImmediate, presentModes[0])
	require.Equal(t, khr_surface.PresentModeFIFORelaxed, presentModes[1])
}

func TestVulkanSurface_SupportsDevice(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	instance := mocks.NewDummyInstance(common.Vulkan1_0, []string{})
	surfaceDriver := mock_surface.NewMockLoader(ctrl)
	device := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_0)

	extension := khr_surface.CreateExtensionDriverFromLoader(surfaceDriver, instance)

	surface := mock_surface.NewDummySurface(instance)
	surfaceDriver.EXPECT().VkGetPhysicalDeviceSurfaceSupportKHR(
		device.Handle(),
		loader.Uint32(3),
		surface.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(device loader.VkPhysicalDevice, queueFamilyIndex loader.Uint32, surface khr_surface_driver.VkSurfaceKHR, pSupport *loader.VkBool32) (common.VkResult, error) {
			*pSupport = loader.VkBool32(1)

			return core1_0.VKSuccess, nil
		})

	supports, _, err := extension.GetPhysicalDeviceSurfaceSupport(surface, device, 3)
	require.NoError(t, err)
	require.True(t, supports)
}

func TestVulkanSurface_Capabilities(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	instance := mocks.NewDummyInstance(common.Vulkan1_0, []string{})
	surfaceDriver := mock_surface.NewMockLoader(ctrl)
	device := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_0)

	extension := khr_surface.CreateExtensionDriverFromLoader(surfaceDriver, instance)

	surface := mock_surface.NewDummySurface(instance)

	surfaceDriver.EXPECT().VkGetPhysicalDeviceSurfaceCapabilitiesKHR(
		device.Handle(),
		surface.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(device loader.VkPhysicalDevice, surface khr_surface_driver.VkSurfaceKHR, pCapabilities *khr_surface_driver.VkSurfaceCapabilitiesKHR) (common.VkResult, error) {
			val := reflect.ValueOf(pCapabilities).Elem()

			*(*uint32)(unsafe.Pointer(val.FieldByName("currentTransform").UnsafeAddr())) = uint32(0x00000002) // VK_SURFACE_TRANSFORM_ROTATE_90_BIT_KHR
			*(*uint32)(unsafe.Pointer(val.FieldByName("maxImageCount").UnsafeAddr())) = uint32(7)
			*(*uint32)(unsafe.Pointer(val.FieldByName("minImageCount").UnsafeAddr())) = uint32(11)
			*(*uint32)(unsafe.Pointer(val.FieldByName("maxImageArrayLayers").UnsafeAddr())) = uint32(5)
			*(*uint32)(unsafe.Pointer(val.FieldByName("supportedTransforms").UnsafeAddr())) = uint32(0x00000010)     // VK_SURFACE_TRANSFORM_HORIZONTAL_MIRROR_BIT_KHR
			*(*uint32)(unsafe.Pointer(val.FieldByName("supportedCompositeAlpha").UnsafeAddr())) = uint32(0x00000002) // VK_COMPOSITE_ALPHA_PRE_MULTIPLIED_BIT_KHR
			*(*uint32)(unsafe.Pointer(val.FieldByName("supportedUsageFlags").UnsafeAddr())) = uint32(0x00000002)     // VK_IMAGE_USAGE_TRANSFER_DST_BIT

			extent := val.FieldByName("currentExtent")

			*(*uint32)(unsafe.Pointer(extent.FieldByName("width").UnsafeAddr())) = uint32(1)
			*(*uint32)(unsafe.Pointer(extent.FieldByName("height").UnsafeAddr())) = uint32(3)

			extent = val.FieldByName("maxImageExtent")

			*(*uint32)(unsafe.Pointer(extent.FieldByName("width").UnsafeAddr())) = uint32(13)
			*(*uint32)(unsafe.Pointer(extent.FieldByName("height").UnsafeAddr())) = uint32(17)

			extent = val.FieldByName("minImageExtent")

			*(*uint32)(unsafe.Pointer(extent.FieldByName("width").UnsafeAddr())) = uint32(19)
			*(*uint32)(unsafe.Pointer(extent.FieldByName("height").UnsafeAddr())) = uint32(23)

			return core1_0.VKSuccess, nil
		})

	capabilities, _, err := extension.GetPhysicalDeviceSurfaceCapabilities(surface, device)
	require.NoError(t, err)
	require.Equal(t, 1, capabilities.CurrentExtent.Width)
	require.Equal(t, 3, capabilities.CurrentExtent.Height)
	require.Equal(t, khr_surface.TransformRotate90, capabilities.CurrentTransform)
	require.Equal(t, 5, capabilities.MaxImageArrayLayers)
	require.Equal(t, 7, capabilities.MaxImageCount)
	require.Equal(t, 11, capabilities.MinImageCount)
	require.Equal(t, 13, capabilities.MaxImageExtent.Width)
	require.Equal(t, 17, capabilities.MaxImageExtent.Height)
	require.Equal(t, 19, capabilities.MinImageExtent.Width)
	require.Equal(t, 23, capabilities.MinImageExtent.Height)
	require.Equal(t, khr_surface.TransformHorizontalMirror, capabilities.SupportedTransforms)
	require.Equal(t, khr_surface.CompositeAlphaPreMultiplied, capabilities.SupportedCompositeAlpha)
	require.Equal(t, core1_0.ImageUsageTransferDst, capabilities.SupportedUsageFlags)
}

func TestVulkanSurface_Formats(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	instance := mocks.NewDummyInstance(common.Vulkan1_0, []string{})
	surfaceDriver := mock_surface.NewMockLoader(ctrl)
	device := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_0)

	extension := khr_surface.CreateExtensionDriverFromLoader(surfaceDriver, instance)

	surface := mock_surface.NewDummySurface(instance)

	surfaceDriver.EXPECT().VkGetPhysicalDeviceSurfaceFormatsKHR(
		device.Handle(),
		surface.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
	).DoAndReturn(
		func(device loader.VkPhysicalDevice, surface khr_surface_driver.VkSurfaceKHR, pFormatCount *loader.Uint32, pFormats *khr_surface_driver.VkSurfaceFormatKHR) (common.VkResult, error) {
			*pFormatCount = loader.Uint32(2)

			return core1_0.VKSuccess, nil
		})

	surfaceDriver.EXPECT().VkGetPhysicalDeviceSurfaceFormatsKHR(
		device.Handle(),
		surface.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(device loader.VkPhysicalDevice, surface khr_surface_driver.VkSurfaceKHR, pFormatCount *loader.Uint32, pFormats *khr_surface_driver.VkSurfaceFormatKHR) (common.VkResult, error) {
			require.Equal(t, loader.Uint32(2), *pFormatCount)

			formatSlice := ([]khr_surface_driver.VkSurfaceFormatKHR)(unsafe.Slice(pFormats, 2))
			val := reflect.ValueOf(formatSlice)

			format := val.Index(0)
			*(*uint32)(unsafe.Pointer(format.FieldByName("format").UnsafeAddr())) = uint32(64)    // VK_FORMAT_A2B10G10R10_UNORM_PACK32
			*(*uint32)(unsafe.Pointer(format.FieldByName("colorSpace").UnsafeAddr())) = uint32(0) // VK_COLOR_SPACE_SRGB_NONLINEAR_KHR

			format = val.Index(1)
			*(*uint32)(unsafe.Pointer(format.FieldByName("format").UnsafeAddr())) = uint32(162)   // VK_FORMAT_ASTC_5x5_SRGB_BLOCK
			*(*uint32)(unsafe.Pointer(format.FieldByName("colorSpace").UnsafeAddr())) = uint32(0) // VK_COLOR_SPACE_SRGB_NONLINEAR_KHR

			return core1_0.VKSuccess, nil
		})

	formats, _, err := extension.GetPhysicalDeviceSurfaceFormats(surface, device)
	require.NoError(t, err)
	require.Len(t, formats, 2)

	require.Equal(t, core1_0.FormatA2B10G10R10UnsignedNormalizedPacked, formats[0].Format)
	require.Equal(t, khr_surface.ColorSpaceSRGBNonlinear, formats[0].ColorSpace)

	require.Equal(t, core1_0.FormatASTC5x5_sRGB, formats[1].Format)
	require.Equal(t, khr_surface.ColorSpaceSRGBNonlinear, formats[1].ColorSpace)
}

func TestVulkanSurface_Formats_Incomplete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	instance := mocks.NewDummyInstance(common.Vulkan1_0, []string{})
	surfaceDriver := mock_surface.NewMockLoader(ctrl)
	device := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_0)

	extension := khr_surface.CreateExtensionDriverFromLoader(surfaceDriver, instance)

	surface := mock_surface.NewDummySurface(instance)

	surfaceDriver.EXPECT().VkGetPhysicalDeviceSurfaceFormatsKHR(
		device.Handle(),
		surface.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
	).DoAndReturn(
		func(device loader.VkPhysicalDevice, surface khr_surface_driver.VkSurfaceKHR, pFormatCount *loader.Uint32, pFormats *khr_surface_driver.VkSurfaceFormatKHR) (common.VkResult, error) {
			*pFormatCount = loader.Uint32(1)

			return core1_0.VKSuccess, nil
		})

	surfaceDriver.EXPECT().VkGetPhysicalDeviceSurfaceFormatsKHR(
		device.Handle(),
		surface.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(device loader.VkPhysicalDevice, surface khr_surface_driver.VkSurfaceKHR, pFormatCount *loader.Uint32, pFormats *khr_surface_driver.VkSurfaceFormatKHR) (common.VkResult, error) {
			require.Equal(t, loader.Uint32(1), *pFormatCount)

			formatSlice := ([]khr_surface_driver.VkSurfaceFormatKHR)(unsafe.Slice(pFormats, 1))
			val := reflect.ValueOf(formatSlice)

			format := val.Index(0)
			*(*uint32)(unsafe.Pointer(format.FieldByName("format").UnsafeAddr())) = uint32(64)    // VK_FORMAT_A2B10G10R10_UNORM_PACK32
			*(*uint32)(unsafe.Pointer(format.FieldByName("colorSpace").UnsafeAddr())) = uint32(0) // VK_COLOR_SPACE_SRGB_NONLINEAR_KHR

			return core1_0.VKIncomplete, nil
		})

	surfaceDriver.EXPECT().VkGetPhysicalDeviceSurfaceFormatsKHR(
		device.Handle(),
		surface.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
	).DoAndReturn(
		func(device loader.VkPhysicalDevice, surface khr_surface_driver.VkSurfaceKHR, pFormatCount *loader.Uint32, pFormats *khr_surface_driver.VkSurfaceFormatKHR) (common.VkResult, error) {
			*pFormatCount = loader.Uint32(2)

			return core1_0.VKSuccess, nil
		})

	surfaceDriver.EXPECT().VkGetPhysicalDeviceSurfaceFormatsKHR(
		device.Handle(),
		surface.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(device loader.VkPhysicalDevice, surface khr_surface_driver.VkSurfaceKHR, pFormatCount *loader.Uint32, pFormats *khr_surface_driver.VkSurfaceFormatKHR) (common.VkResult, error) {
			require.Equal(t, loader.Uint32(2), *pFormatCount)

			formatSlice := ([]khr_surface_driver.VkSurfaceFormatKHR)(unsafe.Slice(pFormats, 2))
			val := reflect.ValueOf(formatSlice)

			format := val.Index(0)
			*(*uint32)(unsafe.Pointer(format.FieldByName("format").UnsafeAddr())) = uint32(64)    // VK_FORMAT_A2B10G10R10_UNORM_PACK32
			*(*uint32)(unsafe.Pointer(format.FieldByName("colorSpace").UnsafeAddr())) = uint32(0) // VK_COLOR_SPACE_SRGB_NONLINEAR_KHR

			format = val.Index(1)
			*(*uint32)(unsafe.Pointer(format.FieldByName("format").UnsafeAddr())) = uint32(162)   // VK_FORMAT_ASTC_5x5_SRGB_BLOCK
			*(*uint32)(unsafe.Pointer(format.FieldByName("colorSpace").UnsafeAddr())) = uint32(0) // VK_COLOR_SPACE_SRGB_NONLINEAR_KHR

			return core1_0.VKSuccess, nil
		})

	formats, _, err := extension.GetPhysicalDeviceSurfaceFormats(surface, device)
	require.NoError(t, err)
	require.Len(t, formats, 2)

	require.Equal(t, core1_0.FormatA2B10G10R10UnsignedNormalizedPacked, formats[0].Format)
	require.Equal(t, khr_surface.ColorSpaceSRGBNonlinear, formats[0].ColorSpace)

	require.Equal(t, core1_0.FormatASTC5x5_sRGB, formats[1].Format)
	require.Equal(t, khr_surface.ColorSpaceSRGBNonlinear, formats[1].ColorSpace)
}
