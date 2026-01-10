package khr_get_surface_capabilities2_test

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
	"github.com/vkngwrapper/core/v3/mocks"
	"github.com/vkngwrapper/extensions/v3/khr_get_surface_capabilities2"
	khr_get_surface_capabilities2_loader "github.com/vkngwrapper/extensions/v3/khr_get_surface_capabilities2/loader"
	mock_get_surface_capabilities2 "github.com/vkngwrapper/extensions/v3/khr_get_surface_capabilities2/mocks"
	"github.com/vkngwrapper/extensions/v3/khr_surface"
	khr_surface_loader "github.com/vkngwrapper/extensions/v3/khr_surface/loader"
	mock_surface "github.com/vkngwrapper/extensions/v3/khr_surface/mocks"
	"go.uber.org/mock/gomock"
)

func TestVulkanExtension_GetPhysicalDeviceSurfaceFormats2(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader := mock_get_surface_capabilities2.NewMockLoader(ctrl)
	driver := khr_get_surface_capabilities2.CreateExtensionDriverFromLoader(mockLoader)

	instance := mocks.NewDummyInstance(common.Vulkan1_0, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_0)
	surface := mock_surface.NewDummySurface(instance)

	mockLoader.EXPECT().VkGetPhysicalDeviceSurfaceFormats2KHR(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
	).DoAndReturn(
		func(handle loader.VkPhysicalDevice,
			o *khr_get_surface_capabilities2_loader.VkPhysicalDeviceSurfaceInfo2KHR,
			outCount *loader.Uint32,
			out *khr_get_surface_capabilities2_loader.VkSurfaceFormat2KHR,
		) (common.VkResult, error) {
			input := reflect.ValueOf(*o)

			require.Equal(t, uint64(1000119000), input.FieldByName("sType").Uint())
			require.Equal(t, surface.Handle(), khr_surface_loader.VkSurfaceKHR(input.FieldByName("surface").UnsafePointer()))
			require.True(t, input.FieldByName("pNext").IsNil())

			*outCount = 2

			return core1_0.VKSuccess, nil
		})

	mockLoader.EXPECT().VkGetPhysicalDeviceSurfaceFormats2KHR(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(handle loader.VkPhysicalDevice,
			o *khr_get_surface_capabilities2_loader.VkPhysicalDeviceSurfaceInfo2KHR,
			outCount *loader.Uint32,
			out *khr_get_surface_capabilities2_loader.VkSurfaceFormat2KHR,
		) (common.VkResult, error) {
			input := reflect.ValueOf(*o)

			require.Equal(t, uint64(1000119000), input.FieldByName("sType").Uint())
			require.Equal(t, surface.Handle(), khr_surface_loader.VkSurfaceKHR(input.FieldByName("surface").UnsafePointer()))
			require.True(t, input.FieldByName("pNext").IsNil())

			require.Equal(t, loader.Uint32(2), *outCount)

			formatSlice := ([]khr_get_surface_capabilities2_loader.VkSurfaceFormat2KHR)(unsafe.Slice(out, 2))
			outData := reflect.ValueOf(formatSlice)
			format := outData.Index(0)
			require.Equal(t, uint64(1000119002), format.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SURFACE_FORMAT_2_KHR
			require.True(t, format.FieldByName("pNext").IsNil())

			surfaceFormat := format.FieldByName("surfaceFormat")
			*(*loader.Uint32)(unsafe.Pointer(surfaceFormat.FieldByName("format").UnsafeAddr())) = loader.Uint32(64)
			*(*loader.Uint32)(unsafe.Pointer(surfaceFormat.FieldByName("colorSpace").UnsafeAddr())) = loader.Uint32(0)

			format = outData.Index(1)
			require.Equal(t, uint64(1000119002), format.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SURFACE_FORMAT_2_KHR
			require.True(t, format.FieldByName("pNext").IsNil())

			surfaceFormat = format.FieldByName("surfaceFormat")
			*(*loader.Uint32)(unsafe.Pointer(surfaceFormat.FieldByName("format").UnsafeAddr())) = loader.Uint32(52)
			*(*loader.Uint32)(unsafe.Pointer(surfaceFormat.FieldByName("colorSpace").UnsafeAddr())) = loader.Uint32(0)

			return core1_0.VKSuccess, nil
		})

	formats, res, err := driver.GetPhysicalDeviceSurfaceFormats2(physicalDevice, khr_get_surface_capabilities2.PhysicalDeviceSurfaceInfo2{
		Surface: surface,
	}, nil)
	require.NoError(t, err)
	require.Equal(t, core1_0.VKSuccess, res)
	require.Equal(t, []*khr_get_surface_capabilities2.SurfaceFormat2{
		{
			SurfaceFormat: khr_surface.SurfaceFormat{
				Format:     core1_0.FormatA2B10G10R10UnsignedNormalizedPacked,
				ColorSpace: khr_surface.ColorSpaceSRGBNonlinear,
			},
		},
		{
			SurfaceFormat: khr_surface.SurfaceFormat{
				Format:     core1_0.FormatA8B8G8R8SignedNormalizedPacked,
				ColorSpace: khr_surface.ColorSpaceSRGBNonlinear,
			},
		},
	}, formats)
}

func TestVulkanExtension_GetPhysicalDeviceSurfaceCapabilities2(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader := mock_get_surface_capabilities2.NewMockLoader(ctrl)
	driver := khr_get_surface_capabilities2.CreateExtensionDriverFromLoader(mockLoader)

	instance := mocks.NewDummyInstance(common.Vulkan1_0, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_0)
	surface := mock_surface.NewDummySurface(instance)

	mockLoader.EXPECT().VkGetPhysicalDeviceSurfaceCapabilities2KHR(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(handle loader.VkPhysicalDevice,
			o *khr_get_surface_capabilities2_loader.VkPhysicalDeviceSurfaceInfo2KHR,
			out *khr_get_surface_capabilities2_loader.VkSurfaceCapabilities2KHR,
		) (common.VkResult, error) {
			input := reflect.ValueOf(*o)

			require.Equal(t, uint64(1000119000), input.FieldByName("sType").Uint())
			require.Equal(t, surface.Handle(), khr_surface_loader.VkSurfaceKHR(input.FieldByName("surface").UnsafePointer()))
			require.True(t, input.FieldByName("pNext").IsNil())

			output := reflect.ValueOf(out).Elem()
			require.Equal(t, uint64(1000119001), output.FieldByName("sType").Uint())
			require.True(t, output.FieldByName("pNext").IsNil())

			caps := output.FieldByName("surfaceCapabilities")

			*(*loader.Uint32)(unsafe.Pointer(caps.FieldByName("minImageCount").UnsafeAddr())) = loader.Uint32(1)
			*(*loader.Uint32)(unsafe.Pointer(caps.FieldByName("maxImageCount").UnsafeAddr())) = loader.Uint32(3)
			*(*loader.Uint32)(unsafe.Pointer(caps.FieldByName("currentExtent").FieldByName("width").UnsafeAddr())) = loader.Uint32(5)
			*(*loader.Uint32)(unsafe.Pointer(caps.FieldByName("currentExtent").FieldByName("height").UnsafeAddr())) = loader.Uint32(7)
			*(*loader.Uint32)(unsafe.Pointer(caps.FieldByName("minImageExtent").FieldByName("width").UnsafeAddr())) = loader.Uint32(11)
			*(*loader.Uint32)(unsafe.Pointer(caps.FieldByName("minImageExtent").FieldByName("height").UnsafeAddr())) = loader.Uint32(13)
			*(*loader.Uint32)(unsafe.Pointer(caps.FieldByName("maxImageExtent").FieldByName("width").UnsafeAddr())) = loader.Uint32(17)
			*(*loader.Uint32)(unsafe.Pointer(caps.FieldByName("maxImageExtent").FieldByName("height").UnsafeAddr())) = loader.Uint32(19)
			*(*loader.Uint32)(unsafe.Pointer(caps.FieldByName("maxImageArrayLayers").UnsafeAddr())) = loader.Uint32(23)
			*(*loader.Uint32)(unsafe.Pointer(caps.FieldByName("supportedTransforms").UnsafeAddr())) = loader.Uint32(1)
			*(*loader.Uint32)(unsafe.Pointer(caps.FieldByName("currentTransform").UnsafeAddr())) = loader.Uint32(2)
			*(*loader.Uint32)(unsafe.Pointer(caps.FieldByName("supportedCompositeAlpha").UnsafeAddr())) = loader.Uint32(4)
			*(*loader.Uint32)(unsafe.Pointer(caps.FieldByName("supportedUsageFlags").UnsafeAddr())) = loader.Uint32(8)

			return core1_0.VKSuccess, nil
		})

	var outData khr_get_surface_capabilities2.SurfaceCapabilities2

	res, err := driver.GetPhysicalDeviceSurfaceCapabilities2(
		physicalDevice,
		khr_get_surface_capabilities2.PhysicalDeviceSurfaceInfo2{
			Surface: surface,
		},
		&outData,
	)
	require.NoError(t, err)
	require.Equal(t, core1_0.VKSuccess, res)
	require.Equal(t, khr_get_surface_capabilities2.SurfaceCapabilities2{
		SurfaceCapabilities: khr_surface.SurfaceCapabilities{
			MinImageCount: 1,
			MaxImageCount: 3,
			CurrentExtent: core1_0.Extent2D{
				Width:  5,
				Height: 7,
			},
			MinImageExtent: core1_0.Extent2D{
				Width:  11,
				Height: 13,
			},
			MaxImageExtent: core1_0.Extent2D{
				Width:  17,
				Height: 19,
			},
			MaxImageArrayLayers:     23,
			SupportedTransforms:     khr_surface.TransformIdentity,
			CurrentTransform:        khr_surface.TransformRotate90,
			SupportedCompositeAlpha: khr_surface.CompositeAlphaPostMultiplied,
			SupportedUsageFlags:     core1_0.ImageUsageStorage,
		},
	}, outData)
}
