package khr_portability_subset

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v2/mocks"
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/common/extensions"
	"github.com/vkngwrapper/core/v2/core1_0"
	"github.com/vkngwrapper/core/v2/driver"
	mock_driver "github.com/vkngwrapper/core/v2/driver/mocks"
	"github.com/vkngwrapper/extensions/v2/khr_get_physical_device_properties2"
	khr_get_physical_device_properties2_driver "github.com/vkngwrapper/extensions/v2/khr_get_physical_device_properties2/driver"
	mock_get_physical_device_properties2 "github.com/vkngwrapper/extensions/v2/khr_get_physical_device_properties2/mocks"
	khr_portability_subset_driver "github.com/vkngwrapper/extensions/v2/khr_portability_subset/driver"
	"reflect"
	"testing"
	"unsafe"
)

func TestPhysicalDevicePortabilitySubsetFeaturesOutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extensionDriver := mock_get_physical_device_properties2.NewMockDriver(ctrl)
	extension := khr_get_physical_device_properties2.CreateExtensionFromDriver(extensionDriver)

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	physicalDevice := mocks.EasyMockPhysicalDevice(ctrl, coreDriver)

	extensionDriver.EXPECT().VkGetPhysicalDeviceFeatures2KHR(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(physicalDevice driver.VkPhysicalDevice,
			pFeatures *khr_get_physical_device_properties2_driver.VkPhysicalDeviceFeatures2KHR) {

			val := reflect.ValueOf(pFeatures).Elem()
			require.Equal(t, uint64(1000059000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FEATURES_2_KHR

			next := (*khr_portability_subset_driver.VkPhysicalDevicePortabilitySubsetFeaturesKHR)(val.FieldByName("pNext").UnsafePointer())
			val = reflect.ValueOf(next).Elem()
			require.Equal(t, uint64(1000163000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PORTABILITY_SUBSET_FEATURES_KHR
			require.True(t, val.FieldByName("pNext").IsNil())
			*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("constantAlphaColorBlendFactors").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("events").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("imageViewFormatReinterpretation").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("imageViewFormatSwizzle").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("imageView2DOn3DImage").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("multisampleArrayImage").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("mutableComparisonSamplers").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("pointPolygons").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("samplerMipLodBias").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("separateStencilMaskRef").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("shaderSampleRateInterpolationFunctions").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("tessellationIsolines").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("tessellationPointMode").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("triangleFans").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("vertexAttributeAccessBeyondStride").UnsafeAddr())) = driver.VkBool32(1)
		})

	var subsetFeatures PhysicalDevicePortabilitySubsetFeatures
	err := extension.PhysicalDeviceFeatures2(
		physicalDevice,
		&khr_get_physical_device_properties2.PhysicalDeviceFeatures2{
			NextOutData: common.NextOutData{&subsetFeatures},
		})
	require.NoError(t, err)
	require.Equal(t, PhysicalDevicePortabilitySubsetFeatures{
		ConstantAlphaColorBlendFactors:          true,
		ImageViewFormatReinterpretation:         true,
		ImageView2DOn3DImage:                    true,
		MutableComparisonSamplers:               true,
		SamplerMipLodBias:                       true,
		ShaderSamplerRateInterpolationFunctions: true,
		TessellationPointMode:                   true,
		VertexAttributeAccessBeyondStride:       true,
	}, subsetFeatures)
}

func TestPhysicalDevicePortabilitySubsetOutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extensionDriver := mock_get_physical_device_properties2.NewMockDriver(ctrl)
	extension := khr_get_physical_device_properties2.CreateExtensionFromDriver(extensionDriver)

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	physicalDevice := mocks.EasyMockPhysicalDevice(ctrl, coreDriver)

	extensionDriver.EXPECT().VkGetPhysicalDeviceProperties2KHR(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(physicalDevice driver.VkPhysicalDevice, pProperties *khr_get_physical_device_properties2_driver.VkPhysicalDeviceProperties2KHR) {
			val := reflect.ValueOf(pProperties).Elem()
			require.Equal(t, uint64(1000059001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROPERTIES_2_KHR

			next := (*khr_portability_subset_driver.VkPhysicalDevicePortabilitySubsetPropertiesKHR)(val.FieldByName("pNext").UnsafePointer())
			val = reflect.ValueOf(next).Elem()

			require.Equal(t, uint64(1000163001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PORTABILITY_SUBSET_PROPERTIES_KHR
			require.True(t, val.FieldByName("pNext").IsNil())
			*(*driver.Uint32)(unsafe.Pointer(val.FieldByName("minVertexInputBindingStrideAlignment").UnsafeAddr())) = driver.Uint32(3)
		})

	var subsetProperties PhysicalDevicePortabilitySubsetProperties
	err := extension.PhysicalDeviceProperties2(
		physicalDevice,
		&khr_get_physical_device_properties2.PhysicalDeviceProperties2{
			NextOutData: common.NextOutData{&subsetProperties},
		})
	require.NoError(t, err)
	require.Equal(t, PhysicalDevicePortabilitySubsetProperties{
		MinVertexInputBindingStrideAlignment: 3,
	}, subsetProperties)
}

func TestPhysicalDevicePortabilitySubsetFeaturesOptions(t *testing.T) {
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

		next := (*khr_portability_subset_driver.VkPhysicalDevicePortabilitySubsetFeaturesKHR)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000163000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PORTABILITY_SUBSET_FEATURES_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(0), val.FieldByName("constantAlphaColorBlendFactors").Uint())
		require.Equal(t, uint64(1), val.FieldByName("events").Uint())
		require.Equal(t, uint64(0), val.FieldByName("imageViewFormatReinterpretation").Uint())
		require.Equal(t, uint64(1), val.FieldByName("imageViewFormatSwizzle").Uint())
		require.Equal(t, uint64(0), val.FieldByName("imageView2DOn3DImage").Uint())
		require.Equal(t, uint64(1), val.FieldByName("multisampleArrayImage").Uint())
		require.Equal(t, uint64(0), val.FieldByName("mutableComparisonSamplers").Uint())
		require.Equal(t, uint64(1), val.FieldByName("pointPolygons").Uint())
		require.Equal(t, uint64(0), val.FieldByName("samplerMipLodBias").Uint())
		require.Equal(t, uint64(1), val.FieldByName("separateStencilMaskRef").Uint())
		require.Equal(t, uint64(0), val.FieldByName("shaderSampleRateInterpolationFunctions").Uint())
		require.Equal(t, uint64(1), val.FieldByName("tessellationIsolines").Uint())
		require.Equal(t, uint64(0), val.FieldByName("tessellationPointMode").Uint())
		require.Equal(t, uint64(1), val.FieldByName("triangleFans").Uint())
		require.Equal(t, uint64(0), val.FieldByName("vertexAttributeAccessBeyondStride").Uint())

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

			NextOptions: common.NextOptions{PhysicalDevicePortabilitySubsetFeatures{
				Events:                 true,
				ImageViewFormatSwizzle: true,
				MultisampleArrayImage:  true,
				PointPolygons:          true,
				SeparateStencilMaskRef: true,
				TessellationIsolines:   true,
				TriangleFans:           true,
			}},
		})
	require.NoError(t, err)
	require.Equal(t, mockDevice.Handle(), device.Handle())
}
