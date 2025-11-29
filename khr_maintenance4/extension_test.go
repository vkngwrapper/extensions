package khr_maintenance4_test

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/common/extensions"
	"github.com/vkngwrapper/core/v2/core1_0"
	"github.com/vkngwrapper/core/v2/core1_1"
	"github.com/vkngwrapper/core/v2/driver"
	mock_driver "github.com/vkngwrapper/core/v2/driver/mocks"
	"github.com/vkngwrapper/core/v2/mocks"
	"github.com/vkngwrapper/extensions/v2/khr_get_physical_device_properties2"
	khr_get_physical_device_properties2_driver "github.com/vkngwrapper/extensions/v2/khr_get_physical_device_properties2/driver"
	mock_get_physical_device_properties2 "github.com/vkngwrapper/extensions/v2/khr_get_physical_device_properties2/mocks"
	"github.com/vkngwrapper/extensions/v2/khr_maintenance4"
	khr_maintenance4_driver "github.com/vkngwrapper/extensions/v2/khr_maintenance4/driver"
	mock_maintenance4 "github.com/vkngwrapper/extensions/v2/khr_maintenance4/dummies"
	"go.uber.org/mock/gomock"
)

func TestPhysicalDeviceMaintenance4Options(t *testing.T) {
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

		next := (*khr_maintenance4_driver.VkPhysicalDeviceMaintenance4FeaturesKHR)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000413000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MAINTENANCE_4_FEATURES
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1), val.FieldByName("maintenance4").Uint())

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
				khr_maintenance4.PhysicalDeviceMaintenance4Features{
					Maintenance4: true,
				},
			},
		})
	require.NoError(t, err)
	require.Equal(t, mockDevice.Handle(), device.Handle())
}

func TestPhysicalDeviceMaintenance4OutData(t *testing.T) {
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

		next := (*khr_maintenance4_driver.VkPhysicalDeviceMaintenance4FeaturesKHR)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000413000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MAINTENANCE_4_FEATURES
		require.True(t, val.FieldByName("pNext").IsNil())
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("maintenance4").UnsafeAddr())) = driver.VkBool32(1)
	})

	var outData khr_maintenance4.PhysicalDeviceMaintenance4Features
	err := extension.PhysicalDeviceFeatures2(
		physicalDevice,
		&khr_get_physical_device_properties2.PhysicalDeviceFeatures2{
			NextOutData: common.NextOutData{&outData},
		},
	)
	require.NoError(t, err)
	require.Equal(t, khr_maintenance4.PhysicalDeviceMaintenance4Features{
		Maintenance4: true,
	}, outData)
}

func TestMaintenance4PropertiesOutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_get_physical_device_properties2.NewMockDriver(ctrl)
	extension := khr_get_physical_device_properties2.CreateExtensionFromDriver(extDriver)

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	physicalDevice := mocks.EasyMockPhysicalDevice(ctrl, coreDriver)

	extDriver.EXPECT().VkGetPhysicalDeviceProperties2KHR(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice driver.VkPhysicalDevice,
		pProperties *khr_get_physical_device_properties2_driver.VkPhysicalDeviceProperties2KHR,
	) {
		val := reflect.ValueOf(pProperties).Elem()
		require.Equal(t, uint64(1000059001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROPERTIES_2_KHR

		next := (*khr_maintenance4_driver.VkPhysicalDeviceMaintenance4PropertiesKHR)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000413001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MAINTENANCE_4_PROPERTIES
		require.True(t, val.FieldByName("pNext").IsNil())
		*(*driver.VkDeviceSize)(unsafe.Pointer(val.FieldByName("maxBufferSize").UnsafeAddr())) = driver.VkDeviceSize(997)
	})

	var outData khr_maintenance4.PhysicalDeviceMaintenance4Properties
	properties := khr_get_physical_device_properties2.PhysicalDeviceProperties2{
		NextOutData: common.NextOutData{&outData},
	}

	err := extension.PhysicalDeviceProperties2(physicalDevice, &properties)
	require.NoError(t, err)
	require.Equal(t, khr_maintenance4.PhysicalDeviceMaintenance4Properties{
		MaxBufferSize: 997,
	}, outData)
}

func TestVulkanExtension_DeviceBufferMemoryRequirements(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_maintenance4.NewMockDriver(ctrl)
	extension := khr_maintenance4.CreateExtensionFromDriver(extDriver)

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := mocks.EasyMockDevice(ctrl, coreDriver)

	extDriver.EXPECT().VkGetDeviceBufferMemoryRequirementsKHR(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil())).DoAndReturn(
		func(device driver.VkDevice, pInfo *khr_maintenance4_driver.VkDeviceBufferMemoryRequirementsKHR, pMemoryRequirements *driver.VkMemoryRequirements2) {
			optionVal := reflect.ValueOf(pInfo).Elem()

			require.Equal(t, uint64(1000413002), optionVal.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_BUFFER_MEMORY_REQUIREMENTS
			require.True(t, optionVal.FieldByName("pNext").IsNil())
			require.False(t, optionVal.FieldByName("pCreateInfo").IsNil())

			createInfoPtr := (*driver.VkBufferCreateInfo)(optionVal.FieldByName("pCreateInfo").UnsafePointer())
			createInfo := reflect.ValueOf(createInfoPtr).Elem()

			require.Equal(t, uint64(12), createInfo.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_BUFFER_CREATE_INFO
			require.True(t, createInfo.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(7), createInfo.FieldByName("size").Uint())
			require.Equal(t, uint64(0), createInfo.FieldByName("flags").Uint())
			require.Equal(t, uint64(0x100), createInfo.FieldByName("usage").Uint()) // VK_BUFFER_USAGE_INDIRECT_BUFFER_BIT

			outDataVal := reflect.ValueOf(pMemoryRequirements).Elem()
			require.Equal(t, uint64(1000146003), outDataVal.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_MEMORY_REQUIREMENTS_2
			require.True(t, outDataVal.FieldByName("pNext").IsNil())

			memReqsVal := outDataVal.FieldByName("memoryRequirements")
			*(*driver.VkDeviceSize)(unsafe.Pointer(memReqsVal.FieldByName("size").UnsafeAddr())) = driver.VkDeviceSize(11)
			*(*driver.VkDeviceSize)(unsafe.Pointer(memReqsVal.FieldByName("alignment").UnsafeAddr())) = driver.VkDeviceSize(13)
			*(*driver.Uint32)(unsafe.Pointer(memReqsVal.FieldByName("memoryTypeBits").UnsafeAddr())) = driver.Uint32(17)
		})

	outData := &core1_1.MemoryRequirements2{}
	err := extension.DeviceBufferMemoryRequirements(device, khr_maintenance4.DeviceBufferMemoryRequirements{
		CreateInfo: core1_0.BufferCreateInfo{
			Size:  7,
			Usage: core1_0.BufferUsageIndirectBuffer,
		},
	}, outData)
	require.NoError(t, err)

	require.Equal(t, &core1_1.MemoryRequirements2{
		MemoryRequirements: core1_0.MemoryRequirements{
			Size:           11,
			Alignment:      13,
			MemoryTypeBits: 17,
		},
	}, outData)
}

func TestVulkanExtension_DeviceImageMemoryRequirements(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_maintenance4.NewMockDriver(ctrl)
	extension := khr_maintenance4.CreateExtensionFromDriver(extDriver)

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := mocks.EasyMockDevice(ctrl, coreDriver)

	extDriver.EXPECT().VkGetDeviceImageMemoryRequirementsKHR(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil())).DoAndReturn(
		func(device driver.VkDevice, pInfo *khr_maintenance4_driver.VkDeviceImageMemoryRequirementsKHR, pMemoryRequirements *driver.VkMemoryRequirements2) {
			optionVal := reflect.ValueOf(pInfo).Elem()

			require.Equal(t, uint64(1000413003), optionVal.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_IMAGE_MEMORY_REQUIREMENTS
			require.True(t, optionVal.FieldByName("pNext").IsNil())
			require.False(t, optionVal.FieldByName("pCreateInfo").IsNil())

			createInfoPtr := (*driver.VkImageCreateInfo)(optionVal.FieldByName("pCreateInfo").UnsafePointer())
			createInfo := reflect.ValueOf(createInfoPtr).Elem()

			require.Equal(t, uint64(14), createInfo.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_IMAGE_CREATE_INFO
			require.True(t, createInfo.FieldByName("pNext").IsNil())

			require.Equal(t, uint64(7), createInfo.FieldByName("extent").FieldByName("width").Uint())
			require.Equal(t, uint64(11), createInfo.FieldByName("extent").FieldByName("height").Uint())
			require.Equal(t, uint64(13), createInfo.FieldByName("extent").FieldByName("depth").Uint())
			require.Equal(t, uint64(0), createInfo.FieldByName("flags").Uint())
			require.Equal(t, uint64(4), createInfo.FieldByName("usage").Uint()) // VK_IMAGE_USAGE_SAMPLED_BIT

			outDataVal := reflect.ValueOf(pMemoryRequirements).Elem()
			require.Equal(t, uint64(1000146003), outDataVal.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_MEMORY_REQUIREMENTS_2
			require.True(t, outDataVal.FieldByName("pNext").IsNil())

			memReqsVal := outDataVal.FieldByName("memoryRequirements")
			*(*driver.VkDeviceSize)(unsafe.Pointer(memReqsVal.FieldByName("size").UnsafeAddr())) = driver.VkDeviceSize(17)
			*(*driver.VkDeviceSize)(unsafe.Pointer(memReqsVal.FieldByName("alignment").UnsafeAddr())) = driver.VkDeviceSize(19)
			*(*driver.Uint32)(unsafe.Pointer(memReqsVal.FieldByName("memoryTypeBits").UnsafeAddr())) = driver.Uint32(23)
		})

	outData := &core1_1.MemoryRequirements2{}
	err := extension.DeviceImageMemoryRequirements(device, khr_maintenance4.DeviceImageMemoryRequirements{
		CreateInfo: core1_0.ImageCreateInfo{
			Extent: core1_0.Extent3D{
				Width:  7,
				Height: 11,
				Depth:  13,
			},
			Usage: core1_0.ImageUsageSampled,
		},
	}, outData)
	require.NoError(t, err)

	require.Equal(t, &core1_1.MemoryRequirements2{
		MemoryRequirements: core1_0.MemoryRequirements{
			Size:           17,
			Alignment:      19,
			MemoryTypeBits: 23,
		},
	}, outData)
}

func TestVulkanExtension_SparseImageMemoryRequirements(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_maintenance4.NewMockDriver(ctrl)
	extension := khr_maintenance4.CreateExtensionFromDriver(extDriver)

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := mocks.EasyMockDevice(ctrl, coreDriver)

	extDriver.EXPECT().VkGetDeviceImageSparseMemoryRequirementsKHR(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
	).DoAndReturn(
		func(device driver.VkDevice,
			pInfo *khr_maintenance4_driver.VkDeviceImageMemoryRequirementsKHR,
			pSparseMemoryRequirementCount *driver.Uint32,
			pSparseMemoryRequirements *driver.VkSparseImageMemoryRequirements2) {

			options := reflect.ValueOf(pInfo).Elem()
			require.Equal(t, uint64(1000413003), options.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_IMAGE_MEMORY_REQUIREMENTS
			require.True(t, options.FieldByName("pNext").IsNil())
			require.False(t, options.FieldByName("pCreateInfo").IsNil())

			createInfoPtr := (*driver.VkImageCreateInfo)(options.FieldByName("pCreateInfo").UnsafePointer())
			createInfo := reflect.ValueOf(createInfoPtr).Elem()

			require.Equal(t, uint64(14), createInfo.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_IMAGE_CREATE_INFO
			require.True(t, createInfo.FieldByName("pNext").IsNil())

			require.Equal(t, uint64(7), createInfo.FieldByName("extent").FieldByName("width").Uint())
			require.Equal(t, uint64(11), createInfo.FieldByName("extent").FieldByName("height").Uint())
			require.Equal(t, uint64(13), createInfo.FieldByName("extent").FieldByName("depth").Uint())
			require.Equal(t, uint64(0), createInfo.FieldByName("flags").Uint())
			require.Equal(t, uint64(4), createInfo.FieldByName("usage").Uint()) // VK_IMAGE_USAGE_SAMPLED_BIT

			*pSparseMemoryRequirementCount = driver.Uint32(2)
		})

	extDriver.EXPECT().VkGetDeviceImageSparseMemoryRequirementsKHR(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(device driver.VkDevice,
			pInfo *khr_maintenance4_driver.VkDeviceImageMemoryRequirementsKHR,
			pSparseMemoryRequirementCount *driver.Uint32,
			pSparseMemoryRequirements *driver.VkSparseImageMemoryRequirements2) {

			require.Equal(t, driver.Uint32(2), *pSparseMemoryRequirementCount)

			options := reflect.ValueOf(pInfo).Elem()
			require.Equal(t, uint64(1000413003), options.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_IMAGE_MEMORY_REQUIREMENTS
			require.True(t, options.FieldByName("pNext").IsNil())
			require.False(t, options.FieldByName("pCreateInfo").IsNil())

			createInfoPtr := (*driver.VkImageCreateInfo)(options.FieldByName("pCreateInfo").UnsafePointer())
			createInfo := reflect.ValueOf(createInfoPtr).Elem()

			require.Equal(t, uint64(14), createInfo.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_IMAGE_CREATE_INFO
			require.True(t, createInfo.FieldByName("pNext").IsNil())

			require.Equal(t, uint64(7), createInfo.FieldByName("extent").FieldByName("width").Uint())
			require.Equal(t, uint64(11), createInfo.FieldByName("extent").FieldByName("height").Uint())
			require.Equal(t, uint64(13), createInfo.FieldByName("extent").FieldByName("depth").Uint())
			require.Equal(t, uint64(0), createInfo.FieldByName("flags").Uint())
			require.Equal(t, uint64(4), createInfo.FieldByName("usage").Uint()) // VK_IMAGE_USAGE_SAMPLED_BIT

			requirementSlice := ([]driver.VkSparseImageMemoryRequirements2)(unsafe.Slice(pSparseMemoryRequirements, 2))
			outData := reflect.ValueOf(requirementSlice)
			element := outData.Index(0)
			require.Equal(t, uint64(1000146004), element.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SPARSE_IMAGE_MEMORY_REQUIREMENTS_2_KHR
			require.True(t, element.FieldByName("pNext").IsNil())

			memReqs := element.FieldByName("memoryRequirements")
			imageAspectFlags := (*driver.VkImageAspectFlags)(unsafe.Pointer(memReqs.FieldByName("formatProperties").FieldByName("aspectMask").UnsafeAddr()))
			*imageAspectFlags = driver.VkImageAspectFlags(0x00000008) // VK_IMAGE_ASPECT_METADATA_BIT
			width := (*driver.Uint32)(unsafe.Pointer(memReqs.FieldByName("formatProperties").FieldByName("imageGranularity").FieldByName("width").UnsafeAddr()))
			*width = driver.Uint32(1)
			height := (*driver.Uint32)(unsafe.Pointer(memReqs.FieldByName("formatProperties").FieldByName("imageGranularity").FieldByName("height").UnsafeAddr()))
			*height = driver.Uint32(3)
			depth := (*driver.Uint32)(unsafe.Pointer(memReqs.FieldByName("formatProperties").FieldByName("imageGranularity").FieldByName("depth").UnsafeAddr()))
			*depth = driver.Uint32(5)
			flags := (*driver.VkSparseImageFormatFlags)(unsafe.Pointer(memReqs.FieldByName("formatProperties").FieldByName("flags").UnsafeAddr()))
			*flags = driver.VkSparseImageFormatFlags(0x00000004) // VK_SPARSE_IMAGE_FORMAT_NONSTANDARD_BLOCK_SIZE_BIT
			*(*driver.Uint32)(unsafe.Pointer(memReqs.FieldByName("imageMipTailFirstLod").UnsafeAddr())) = driver.Uint32(7)
			*(*driver.VkDeviceSize)(unsafe.Pointer(memReqs.FieldByName("imageMipTailSize").UnsafeAddr())) = driver.VkDeviceSize(17)
			*(*driver.VkDeviceSize)(unsafe.Pointer(memReqs.FieldByName("imageMipTailOffset").UnsafeAddr())) = driver.VkDeviceSize(11)
			*(*driver.VkDeviceSize)(unsafe.Pointer(memReqs.FieldByName("imageMipTailStride").UnsafeAddr())) = driver.VkDeviceSize(13)

			element = outData.Index(1)
			require.Equal(t, uint64(1000146004), element.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SPARSE_IMAGE_MEMORY_REQUIREMENTS_2_KHR
			require.True(t, element.FieldByName("pNext").IsNil())

			memReqs = element.FieldByName("memoryRequirements")
			imageAspectFlags = (*driver.VkImageAspectFlags)(unsafe.Pointer(memReqs.FieldByName("formatProperties").FieldByName("aspectMask").UnsafeAddr()))
			*imageAspectFlags = driver.VkImageAspectFlags(0x00000004) // VK_IMAGE_ASPECT_STENCIL_BIT
			width = (*driver.Uint32)(unsafe.Pointer(memReqs.FieldByName("formatProperties").FieldByName("imageGranularity").FieldByName("width").UnsafeAddr()))
			*width = driver.Uint32(19)
			height = (*driver.Uint32)(unsafe.Pointer(memReqs.FieldByName("formatProperties").FieldByName("imageGranularity").FieldByName("height").UnsafeAddr()))
			*height = driver.Uint32(23)
			depth = (*driver.Uint32)(unsafe.Pointer(memReqs.FieldByName("formatProperties").FieldByName("imageGranularity").FieldByName("depth").UnsafeAddr()))
			*depth = driver.Uint32(29)
			flags = (*driver.VkSparseImageFormatFlags)(unsafe.Pointer(memReqs.FieldByName("formatProperties").FieldByName("flags").UnsafeAddr()))
			*flags = driver.VkSparseImageFormatFlags(0)
			*(*driver.Uint32)(unsafe.Pointer(memReqs.FieldByName("imageMipTailFirstLod").UnsafeAddr())) = driver.Uint32(43)
			*(*driver.VkDeviceSize)(unsafe.Pointer(memReqs.FieldByName("imageMipTailSize").UnsafeAddr())) = driver.VkDeviceSize(31)
			*(*driver.VkDeviceSize)(unsafe.Pointer(memReqs.FieldByName("imageMipTailOffset").UnsafeAddr())) = driver.VkDeviceSize(41)
			*(*driver.VkDeviceSize)(unsafe.Pointer(memReqs.FieldByName("imageMipTailStride").UnsafeAddr())) = driver.VkDeviceSize(37)
		})

	outData, err := extension.DeviceImageSparseMemoryRequirements(device,
		khr_maintenance4.DeviceImageMemoryRequirements{
			CreateInfo: core1_0.ImageCreateInfo{
				Extent: core1_0.Extent3D{
					Width:  7,
					Height: 11,
					Depth:  13,
				},
				Usage: core1_0.ImageUsageSampled,
			},
		}, nil)
	require.NoError(t, err)
	require.Equal(t, []*core1_1.SparseImageMemoryRequirements2{
		{
			MemoryRequirements: core1_0.SparseImageMemoryRequirements{
				FormatProperties: core1_0.SparseImageFormatProperties{
					AspectMask: core1_0.ImageAspectMetadata,
					ImageGranularity: core1_0.Extent3D{
						Width:  1,
						Height: 3,
						Depth:  5,
					},
					Flags: core1_0.SparseImageFormatNonstandardBlockSize,
				},
				ImageMipTailFirstLod: 7,
				ImageMipTailOffset:   11,
				ImageMipTailStride:   13,
				ImageMipTailSize:     17,
			},
		},
		{
			MemoryRequirements: core1_0.SparseImageMemoryRequirements{
				FormatProperties: core1_0.SparseImageFormatProperties{
					AspectMask: core1_0.ImageAspectStencil,
					ImageGranularity: core1_0.Extent3D{
						Width:  19,
						Height: 23,
						Depth:  29,
					},
					Flags: 0,
				},
				ImageMipTailSize:     31,
				ImageMipTailStride:   37,
				ImageMipTailOffset:   41,
				ImageMipTailFirstLod: 43,
			},
		},
	}, outData)
}
