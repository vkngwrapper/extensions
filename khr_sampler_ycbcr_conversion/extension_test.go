package khr_sampler_ycbcr_conversion_test

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
	"github.com/vkngwrapper/extensions/v3/khr_bind_memory2"
	khr_bind_memory2_driver "github.com/vkngwrapper/extensions/v3/khr_bind_memory2/loader"
	mock_bind_memory2 "github.com/vkngwrapper/extensions/v3/khr_bind_memory2/mocks"
	"github.com/vkngwrapper/extensions/v3/khr_get_memory_requirements2"
	khr_get_memory_requirements2_driver "github.com/vkngwrapper/extensions/v3/khr_get_memory_requirements2/loader"
	mock_get_memory_requirements2 "github.com/vkngwrapper/extensions/v3/khr_get_memory_requirements2/mocks"
	"github.com/vkngwrapper/extensions/v3/khr_get_physical_device_properties2"
	khr_get_physical_device_properties2_driver "github.com/vkngwrapper/extensions/v3/khr_get_physical_device_properties2/loader"
	mock_get_physical_device_properties2 "github.com/vkngwrapper/extensions/v3/khr_get_physical_device_properties2/mocks"
	"github.com/vkngwrapper/extensions/v3/khr_sampler_ycbcr_conversion"
	khr_sampler_ycbcr_conversion_driver "github.com/vkngwrapper/extensions/v3/khr_sampler_ycbcr_conversion/loader"
	mock_sampler_ycbcr_conversion "github.com/vkngwrapper/extensions/v3/khr_sampler_ycbcr_conversion/mocks"
	"go.uber.org/mock/gomock"
)

func TestVulkanExtension_CreateSamplerYcbcrConversion(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	extDriver := mock_sampler_ycbcr_conversion.NewMockLoader(ctrl)
	extension := khr_sampler_ycbcr_conversion.CreateExtensionDriverFromLoader(extDriver, device)

	mockYcbcr := mocks.NewDummySamplerYcbcrConversion(device)

	extDriver.EXPECT().VkCreateSamplerYcbcrConversionKHR(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(device loader.VkDevice,
			pCreateInfo *khr_sampler_ycbcr_conversion_driver.VkSamplerYcbcrConversionCreateInfoKHR,
			pAllocator *loader.VkAllocationCallbacks,
			pYcbcrConversion *khr_sampler_ycbcr_conversion_driver.VkSamplerYcbcrConversionKHR,
		) (common.VkResult, error) {
			*pYcbcrConversion = khr_sampler_ycbcr_conversion_driver.VkSamplerYcbcrConversionKHR(mockYcbcr.Handle())

			val := reflect.ValueOf(pCreateInfo).Elem()
			require.Equal(t, uint64(1000156000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_CREATE_INFO_KHR
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(1000156021), val.FieldByName("format").Uint())             // VK_FORMAT_B12X4G12X4R12X4G12X4_422_UNORM_4PACK16_KHR
			require.Equal(t, uint64(2), val.FieldByName("ycbcrModel").Uint())                  // VK_SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_709_KHR
			require.Equal(t, uint64(1), val.FieldByName("ycbcrRange").Uint())                  // VK_SAMPLER_YCBCR_RANGE_ITU_NARROW_KHR
			require.Equal(t, uint64(4), val.FieldByName("components").FieldByName("r").Uint()) // VK_COMPONENT_SWIZZLE_G
			require.Equal(t, uint64(6), val.FieldByName("components").FieldByName("g").Uint()) // VK_COMPONENT_SWIZZLE_A
			require.Equal(t, uint64(0), val.FieldByName("components").FieldByName("b").Uint()) // VK_COMPONENT_SWIZZLE_IDENTITY
			require.Equal(t, uint64(2), val.FieldByName("components").FieldByName("a").Uint()) // VK_COMPONENT_SWIZZLE_ONE
			require.Equal(t, uint64(0), val.FieldByName("yChromaOffset").Uint())               // VK_CHROMA_LOCATION_COSITED_EVEN_KHR
			require.Equal(t, uint64(1), val.FieldByName("xChromaOffset").Uint())               // VK_CHROMA_LOCATION_MIDPOINT_KHR
			require.Equal(t, uint64(1), val.FieldByName("forceExplicitReconstruction").Uint())

			return core1_0.VKSuccess, nil
		})

	ycbcr, _, err := extension.CreateSamplerYcbcrConversion(
		khr_sampler_ycbcr_conversion.SamplerYcbcrConversionCreateInfo{
			Format:     khr_sampler_ycbcr_conversion.FormatB12X4G12X4R12X4G12X4HorizontalChromaComponentPacked,
			YcbcrModel: khr_sampler_ycbcr_conversion.SamplerYcbcrModelConversionYcbcr709,
			YcbcrRange: khr_sampler_ycbcr_conversion.SamplerYcbcrRangeITUNarrow,
			Components: core1_0.ComponentMapping{
				R: core1_0.ComponentSwizzleGreen,
				G: core1_0.ComponentSwizzleAlpha,
				B: core1_0.ComponentSwizzleIdentity,
				A: core1_0.ComponentSwizzleOne,
			},
			YChromaOffset:               khr_sampler_ycbcr_conversion.ChromaLocationCositedEven,
			XChromaOffset:               khr_sampler_ycbcr_conversion.ChromaLocationMidpoint,
			ChromaFilter:                core1_0.FilterLinear,
			ForceExplicitReconstruction: true,
		},
		nil,
	)
	require.NoError(t, err)
	require.Equal(t, mockYcbcr.Handle(), ycbcr.Handle())

	extDriver.EXPECT().VkDestroySamplerYcbcrConversionKHR(
		device.Handle(),
		khr_sampler_ycbcr_conversion_driver.VkSamplerYcbcrConversionKHR(ycbcr.Handle()),
		gomock.Nil(),
	)

	extension.DestroySamplerYcbcrConversion(ycbcr, nil)
}
func TestBindImagePlaneMemoryOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})

	extDriver := mock_bind_memory2.NewMockLoader(ctrl)
	extension := khr_bind_memory2.CreateExtensionDriverFromLoader(extDriver, device)

	image := mocks.NewDummyImage(device)
	memory := mocks.NewDummyDeviceMemory(device, 1)

	extDriver.EXPECT().VkBindImageMemory2KHR(
		device.Handle(),
		loader.Uint32(1),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device loader.VkDevice,
		bindInfoCount loader.Uint32,
		pBindInfos *khr_bind_memory2_driver.VkBindImageMemoryInfoKHR,
	) (common.VkResult, error) {
		val := reflect.ValueOf(pBindInfos).Elem()

		require.Equal(t, uint64(1000157001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_BIND_IMAGE_MEMORY_INFO_KHR
		require.Equal(t, image.Handle(), loader.VkImage(val.FieldByName("image").UnsafePointer()))
		require.Equal(t, memory.Handle(), loader.VkDeviceMemory(val.FieldByName("memory").UnsafePointer()))

		next := (*khr_sampler_ycbcr_conversion_driver.VkBindImagePlaneMemoryInfoKHR)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000156002), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_BIND_IMAGE_PLANE_MEMORY_INFO_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(0x40), val.FieldByName("planeAspect").Uint()) // VK_IMAGE_ASPECT_PLANE_2_BIT_KHR

		return core1_0.VKSuccess, nil
	})

	_, err := extension.BindImageMemory2(
		khr_bind_memory2.BindImageMemoryInfo{
			Image:  image,
			Memory: memory,

			NextOptions: common.NextOptions{
				khr_sampler_ycbcr_conversion.BindImagePlaneMemoryInfo{
					PlaneAspect: khr_sampler_ycbcr_conversion.ImageAspectPlane2,
				},
			},
		},
	)
	require.NoError(t, err)
}

func TestImagePlaneMemoryRequirementsOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})

	extDriver := mock_get_memory_requirements2.NewMockLoader(ctrl)
	extension := khr_get_memory_requirements2.CreateExtensionDriverFromLoader(extDriver, device)

	image := mocks.NewDummyImage(device)

	extDriver.EXPECT().VkGetImageMemoryRequirements2KHR(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device loader.VkDevice,
		pInfo *khr_get_memory_requirements2_driver.VkImageMemoryRequirementsInfo2KHR,
		pMemoryRequirements *khr_get_memory_requirements2_driver.VkMemoryRequirements2KHR,
	) {
		val := reflect.ValueOf(pInfo).Elem()
		require.Equal(t, uint64(1000146001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_IMAGE_MEMORY_REQUIREMENTS_INFO_2_KHR
		require.Equal(t, image.Handle(), loader.VkImage(val.FieldByName("image").UnsafePointer()))

		next := (*khr_sampler_ycbcr_conversion_driver.VkImagePlaneMemoryRequirementsInfoKHR)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()
		require.Equal(t, uint64(1000156003), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_IMAGE_PLANE_MEMORY_REQUIREMENTS_INFO_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(0x20), val.FieldByName("planeAspect").Uint()) // VK_IMAGE_ASPECT_PLANE_1_BIT_KHR

		val = reflect.ValueOf(pMemoryRequirements).Elem()
		require.Equal(t, uint64(1000146003), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_MEMORY_REQUIREMENTS_2_KHR
		require.True(t, val.FieldByName("pNext").IsNil())

		*(*uint32)(unsafe.Pointer(val.FieldByName("memoryRequirements").FieldByName("size").UnsafeAddr())) = uint32(17)
		*(*uint32)(unsafe.Pointer(val.FieldByName("memoryRequirements").FieldByName("alignment").UnsafeAddr())) = uint32(19)
		*(*uint32)(unsafe.Pointer(val.FieldByName("memoryRequirements").FieldByName("memoryTypeBits").UnsafeAddr())) = uint32(7)
	})

	var outData khr_get_memory_requirements2.MemoryRequirements2
	err := extension.GetImageMemoryRequirements2(
		khr_get_memory_requirements2.ImageMemoryRequirementsInfo2{
			Image: image,
			NextOptions: common.NextOptions{
				khr_sampler_ycbcr_conversion.ImagePlaneMemoryRequirementsInfo{
					PlaneAspect: khr_sampler_ycbcr_conversion.ImageAspectPlane1,
				},
			},
		},
		&outData)
	require.NoError(t, err)
	require.Equal(t, khr_get_memory_requirements2.MemoryRequirements2{
		MemoryRequirements: core1_0.MemoryRequirements{
			Size:           17,
			Alignment:      19,
			MemoryTypeBits: 7,
		},
	}, outData)
}

func TestSamplerYcbcrConversionOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	image := mocks.NewDummyImage(device)
	ycbcr := mocks.NewDummySamplerYcbcrConversion(device)
	mockImageView := mocks.NewDummyImageView(device)

	coreLoader := mock_driver.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := mocks1_0.InternalDeviceDriver(device, coreLoader)

	coreLoader.EXPECT().VkCreateImageView(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device loader.VkDevice,
		pCreateInfo *loader.VkImageViewCreateInfo,
		pAllocator *loader.VkAllocationCallbacks,
		pView *loader.VkImageView,
	) (common.VkResult, error) {
		*pView = mockImageView.Handle()

		val := reflect.ValueOf(pCreateInfo).Elem()
		require.Equal(t, uint64(15), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_IMAGE_VIEW_CREATE_INFO
		require.Equal(t, image.Handle(), loader.VkImage(val.FieldByName("image").UnsafePointer()))
		require.Equal(t, uint64(1000156028), val.FieldByName("format").Uint()) // VK_FORMAT_B16G16R16G16_422_UNORM_KHR

		next := (*khr_sampler_ycbcr_conversion_driver.VkSamplerYcbcrConversionInfoKHR)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()
		require.Equal(t, uint64(1000156001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_INFO_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, ycbcr.Handle(), loader.VkSamplerYcbcrConversion(val.FieldByName("conversion").UnsafePointer()))

		return core1_0.VKSuccess, nil
	})

	imageView, _, err := driver.CreateImageView(
		nil,
		core1_0.ImageViewCreateInfo{
			Image:  image,
			Format: khr_sampler_ycbcr_conversion.FormatB16G16R16G16HorizontalChroma,

			NextOptions: common.NextOptions{
				khr_sampler_ycbcr_conversion.SamplerYcbcrConversionInfo{
					Conversion: ycbcr,
				},
			},
		})
	require.NoError(t, err)
	require.Equal(t, mockImageView.Handle(), imageView.Handle())
}

func TestSamplerYcbcrFeaturesOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreLoader := mock_driver.LoaderForVersion(ctrl, common.Vulkan1_0)
	instance := mocks.NewDummyInstance(common.Vulkan1_0, []string{})
	driver := mocks1_0.InternalCoreInstanceDriver(instance, coreLoader)

	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_0)
	mockDevice := mocks.NewDummyDevice(common.Vulkan1_0, []string{})

	coreLoader.EXPECT().VkCreateDevice(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(physicalDevice loader.VkPhysicalDevice,
			pCreateInfo *loader.VkDeviceCreateInfo,
			pAllocator *loader.VkAllocationCallbacks,
			pDevice *loader.VkDevice,
		) (common.VkResult, error) {
			*pDevice = mockDevice.Handle()

			val := reflect.ValueOf(pCreateInfo).Elem()
			require.Equal(t, uint64(3), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_CREATE_INFO

			next := (*khr_sampler_ycbcr_conversion_driver.VkPhysicalDeviceSamplerYcbcrConversionFeaturesKHR)(val.FieldByName("pNext").UnsafePointer())
			val = reflect.ValueOf(next).Elem()
			require.Equal(t, uint64(1000156004), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SAMPLER_YCBCR_CONVERSION_FEATURES_KHR
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(1), val.FieldByName("samplerYcbcrConversion").Uint())

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

			NextOptions: common.NextOptions{
				khr_sampler_ycbcr_conversion.PhysicalDeviceSamplerYcbcrConversionFeatures{
					SamplerYcbcrConversion: true,
				},
			},
		})
	require.NoError(t, err)
	require.Equal(t, mockDevice.Handle(), device.Handle())
}

func TestSamplerYcbcrFeaturesOutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_get_physical_device_properties2.NewMockLoader(ctrl)
	extension := khr_get_physical_device_properties2.CreateExtensionDriverFromLoader(extDriver)

	instance := mocks.NewDummyInstance(common.Vulkan1_0, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_0)

	extDriver.EXPECT().VkGetPhysicalDeviceFeatures2KHR(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(physicalDevice loader.VkPhysicalDevice,
			pFeatures *khr_get_physical_device_properties2_driver.VkPhysicalDeviceFeatures2KHR,
		) {
			val := reflect.ValueOf(pFeatures).Elem()
			require.Equal(t, uint64(1000059000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FEATURES_2_KHR

			next := (*khr_sampler_ycbcr_conversion_driver.VkPhysicalDeviceSamplerYcbcrConversionFeaturesKHR)(val.FieldByName("pNext").UnsafePointer())
			val = reflect.ValueOf(next).Elem()
			require.Equal(t, uint64(1000156004), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SAMPLER_YCBCR_CONVERSION_FEATURES_KHR
			require.True(t, val.FieldByName("pNext").IsNil())
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("samplerYcbcrConversion").UnsafeAddr())) = loader.VkBool32(1)
		})

	var outData khr_sampler_ycbcr_conversion.PhysicalDeviceSamplerYcbcrConversionFeatures

	err := extension.GetPhysicalDeviceFeatures2(
		physicalDevice,
		&khr_get_physical_device_properties2.PhysicalDeviceFeatures2{
			NextOutData: common.NextOutData{
				&outData,
			},
		})
	require.NoError(t, err)
	require.Equal(t, khr_sampler_ycbcr_conversion.PhysicalDeviceSamplerYcbcrConversionFeatures{
		SamplerYcbcrConversion: true,
	}, outData)
}

func TestSamplerYcbcrImageFormatOutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_get_physical_device_properties2.NewMockLoader(ctrl)
	extension := khr_get_physical_device_properties2.CreateExtensionDriverFromLoader(extDriver)

	instance := mocks.NewDummyInstance(common.Vulkan1_0, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_0)

	extDriver.EXPECT().VkGetPhysicalDeviceImageFormatProperties2KHR(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(physicalDevice loader.VkPhysicalDevice,
			pImageFormatInfo *khr_get_physical_device_properties2_driver.VkPhysicalDeviceImageFormatInfo2KHR,
			pImageFormatProperties *khr_get_physical_device_properties2_driver.VkImageFormatProperties2KHR,
		) (common.VkResult, error) {
			val := reflect.ValueOf(pImageFormatInfo).Elem()
			require.Equal(t, uint64(1000059004), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_FORMAT_INFO_2_KHR
			require.True(t, val.FieldByName("pNext").IsNil())

			val = reflect.ValueOf(pImageFormatProperties).Elem()
			require.Equal(t, uint64(1000059003), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_IMAGE_FORMAT_PROPERTIES_2_KHR

			next := (*khr_sampler_ycbcr_conversion_driver.VkSamplerYcbcrConversionImageFormatPropertiesKHR)(val.FieldByName("pNext").UnsafePointer())
			val = reflect.ValueOf(next).Elem()
			require.Equal(t, uint64(1000156005), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_IMAGE_FORMAT_PROPERTIES_KHR
			require.True(t, val.FieldByName("pNext").IsNil())
			*(*uint32)(unsafe.Pointer(val.FieldByName("combinedImageSamplerDescriptorCount").UnsafeAddr())) = uint32(7)

			return core1_0.VKSuccess, nil
		})

	var outData khr_sampler_ycbcr_conversion.SamplerYcbcrConversionImageFormatProperties
	_, err := extension.GetPhysicalDeviceImageFormatProperties2(
		physicalDevice,
		khr_get_physical_device_properties2.PhysicalDeviceImageFormatInfo2{},
		&khr_get_physical_device_properties2.ImageFormatProperties2{
			NextOutData: common.NextOutData{&outData},
		})
	require.NoError(t, err)
	require.Equal(t, khr_sampler_ycbcr_conversion.SamplerYcbcrConversionImageFormatProperties{
		CombinedImageSamplerDescriptorCount: 7,
	}, outData)
}
