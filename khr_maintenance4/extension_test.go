package khr_maintenance4_test

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/core1_1"
	"github.com/vkngwrapper/core/v3/loader"
	mock_driver "github.com/vkngwrapper/core/v3/loader/mocks"
	"github.com/vkngwrapper/core/v3/mocks"
	"github.com/vkngwrapper/core/v3/mocks/mocks1_0"
	"github.com/vkngwrapper/extensions/v3/khr_get_physical_device_properties2"
	khr_get_physical_device_properties2_driver "github.com/vkngwrapper/extensions/v3/khr_get_physical_device_properties2/loader"
	mock_get_physical_device_properties2 "github.com/vkngwrapper/extensions/v3/khr_get_physical_device_properties2/mocks"
	"github.com/vkngwrapper/extensions/v3/khr_maintenance4"
	khr_maintenance4_loader "github.com/vkngwrapper/extensions/v3/khr_maintenance4/loader"
	mock_maintenance4 "github.com/vkngwrapper/extensions/v3/khr_maintenance4/mocks"
	"go.uber.org/mock/gomock"
)

func TestPhysicalDeviceMaintenance4Options(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreLoader := mock_driver.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := mocks1_0.InternalCoreInstanceDriver(coreLoader)
	instance := mocks.NewDummyInstance(common.Vulkan1_0, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_0)
	mockDevice := mocks.NewDummyDevice(common.Vulkan1_0, []string{})

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

		next := (*khr_maintenance4_loader.VkPhysicalDeviceMaintenance4FeaturesKHR)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000413000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MAINTENANCE_4_FEATURES
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1), val.FieldByName("maintenance4").Uint())

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

	extDriver := mock_get_physical_device_properties2.NewMockLoader(ctrl)
	extension := khr_get_physical_device_properties2.CreateExtensionFromDriver(extDriver)

	instance := mocks.NewDummyInstance(common.Vulkan1_0, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_0)

	extDriver.EXPECT().VkGetPhysicalDeviceFeatures2KHR(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice loader.VkPhysicalDevice,
		pFeatures *khr_get_physical_device_properties2_driver.VkPhysicalDeviceFeatures2KHR) {

		val := reflect.ValueOf(pFeatures).Elem()
		require.Equal(t, uint64(1000059000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FEATURES_2_KHR

		next := (*khr_maintenance4_loader.VkPhysicalDeviceMaintenance4FeaturesKHR)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000413000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MAINTENANCE_4_FEATURES
		require.True(t, val.FieldByName("pNext").IsNil())
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("maintenance4").UnsafeAddr())) = loader.VkBool32(1)
	})

	var outData khr_maintenance4.PhysicalDeviceMaintenance4Features
	err := extension.GetPhysicalDeviceFeatures2(
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

	extDriver := mock_get_physical_device_properties2.NewMockLoader(ctrl)
	extension := khr_get_physical_device_properties2.CreateExtensionFromDriver(extDriver)

	instance := mocks.NewDummyInstance(common.Vulkan1_0, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_0)

	extDriver.EXPECT().VkGetPhysicalDeviceProperties2KHR(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice loader.VkPhysicalDevice,
		pProperties *khr_get_physical_device_properties2_driver.VkPhysicalDeviceProperties2KHR,
	) {
		val := reflect.ValueOf(pProperties).Elem()
		require.Equal(t, uint64(1000059001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROPERTIES_2_KHR

		next := (*khr_maintenance4_loader.VkPhysicalDeviceMaintenance4PropertiesKHR)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000413001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MAINTENANCE_4_PROPERTIES
		require.True(t, val.FieldByName("pNext").IsNil())
		*(*loader.VkDeviceSize)(unsafe.Pointer(val.FieldByName("maxBufferSize").UnsafeAddr())) = loader.VkDeviceSize(997)
	})

	var outData khr_maintenance4.PhysicalDeviceMaintenance4Properties
	properties := khr_get_physical_device_properties2.PhysicalDeviceProperties2{
		NextOutData: common.NextOutData{&outData},
	}

	err := extension.GetPhysicalDeviceProperties2(physicalDevice, &properties)
	require.NoError(t, err)
	require.Equal(t, khr_maintenance4.PhysicalDeviceMaintenance4Properties{
		MaxBufferSize: 997,
	}, outData)
}

func TestVulkanExtension_DeviceBufferMemoryRequirements(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_maintenance4.NewMockLoader(ctrl)
	extension := khr_maintenance4.CreateExtensionFromDriver(extDriver)

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})

	extDriver.EXPECT().VkGetDeviceBufferMemoryRequirementsKHR(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil())).DoAndReturn(
		func(device loader.VkDevice, pInfo *khr_maintenance4_loader.VkDeviceBufferMemoryRequirementsKHR, pMemoryRequirements *loader.VkMemoryRequirements2) {
			optionVal := reflect.ValueOf(pInfo).Elem()

			require.Equal(t, uint64(1000413002), optionVal.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_BUFFER_MEMORY_REQUIREMENTS
			require.True(t, optionVal.FieldByName("pNext").IsNil())
			require.False(t, optionVal.FieldByName("pCreateInfo").IsNil())

			createInfoPtr := (*loader.VkBufferCreateInfo)(optionVal.FieldByName("pCreateInfo").UnsafePointer())
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
			*(*loader.VkDeviceSize)(unsafe.Pointer(memReqsVal.FieldByName("size").UnsafeAddr())) = loader.VkDeviceSize(11)
			*(*loader.VkDeviceSize)(unsafe.Pointer(memReqsVal.FieldByName("alignment").UnsafeAddr())) = loader.VkDeviceSize(13)
			*(*loader.Uint32)(unsafe.Pointer(memReqsVal.FieldByName("memoryTypeBits").UnsafeAddr())) = loader.Uint32(17)
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

	extDriver := mock_maintenance4.NewMockLoader(ctrl)
	extension := khr_maintenance4.CreateExtensionFromDriver(extDriver)

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})

	extDriver.EXPECT().VkGetDeviceImageMemoryRequirementsKHR(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil())).DoAndReturn(
		func(device loader.VkDevice, pInfo *khr_maintenance4_loader.VkDeviceImageMemoryRequirementsKHR, pMemoryRequirements *loader.VkMemoryRequirements2) {
			optionVal := reflect.ValueOf(pInfo).Elem()

			require.Equal(t, uint64(1000413003), optionVal.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_IMAGE_MEMORY_REQUIREMENTS
			require.True(t, optionVal.FieldByName("pNext").IsNil())
			require.False(t, optionVal.FieldByName("pCreateInfo").IsNil())

			createInfoPtr := (*loader.VkImageCreateInfo)(optionVal.FieldByName("pCreateInfo").UnsafePointer())
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
			*(*loader.VkDeviceSize)(unsafe.Pointer(memReqsVal.FieldByName("size").UnsafeAddr())) = loader.VkDeviceSize(17)
			*(*loader.VkDeviceSize)(unsafe.Pointer(memReqsVal.FieldByName("alignment").UnsafeAddr())) = loader.VkDeviceSize(19)
			*(*loader.Uint32)(unsafe.Pointer(memReqsVal.FieldByName("memoryTypeBits").UnsafeAddr())) = loader.Uint32(23)
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

	extDriver := mock_maintenance4.NewMockLoader(ctrl)
	extension := khr_maintenance4.CreateExtensionFromDriver(extDriver)
	
	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})

	extDriver.EXPECT().VkGetDeviceImageSparseMemoryRequirementsKHR(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
	).DoAndReturn(
		func(device loader.VkDevice,
			pInfo *khr_maintenance4_loader.VkDeviceImageMemoryRequirementsKHR,
			pSparseMemoryRequirementCount *loader.Uint32,
			pSparseMemoryRequirements *loader.VkSparseImageMemoryRequirements2) {

			options := reflect.ValueOf(pInfo).Elem()
			require.Equal(t, uint64(1000413003), options.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_IMAGE_MEMORY_REQUIREMENTS
			require.True(t, options.FieldByName("pNext").IsNil())
			require.False(t, options.FieldByName("pCreateInfo").IsNil())

			createInfoPtr := (*loader.VkImageCreateInfo)(options.FieldByName("pCreateInfo").UnsafePointer())
			createInfo := reflect.ValueOf(createInfoPtr).Elem()

			require.Equal(t, uint64(14), createInfo.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_IMAGE_CREATE_INFO
			require.True(t, createInfo.FieldByName("pNext").IsNil())

			require.Equal(t, uint64(7), createInfo.FieldByName("extent").FieldByName("width").Uint())
			require.Equal(t, uint64(11), createInfo.FieldByName("extent").FieldByName("height").Uint())
			require.Equal(t, uint64(13), createInfo.FieldByName("extent").FieldByName("depth").Uint())
			require.Equal(t, uint64(0), createInfo.FieldByName("flags").Uint())
			require.Equal(t, uint64(4), createInfo.FieldByName("usage").Uint()) // VK_IMAGE_USAGE_SAMPLED_BIT

			*pSparseMemoryRequirementCount = loader.Uint32(2)
		})

	extDriver.EXPECT().VkGetDeviceImageSparseMemoryRequirementsKHR(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(device loader.VkDevice,
			pInfo *khr_maintenance4_loader.VkDeviceImageMemoryRequirementsKHR,
			pSparseMemoryRequirementCount *loader.Uint32,
			pSparseMemoryRequirements *loader.VkSparseImageMemoryRequirements2) {

			require.Equal(t, loader.Uint32(2), *pSparseMemoryRequirementCount)

			options := reflect.ValueOf(pInfo).Elem()
			require.Equal(t, uint64(1000413003), options.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_IMAGE_MEMORY_REQUIREMENTS
			require.True(t, options.FieldByName("pNext").IsNil())
			require.False(t, options.FieldByName("pCreateInfo").IsNil())

			createInfoPtr := (*loader.VkImageCreateInfo)(options.FieldByName("pCreateInfo").UnsafePointer())
			createInfo := reflect.ValueOf(createInfoPtr).Elem()

			require.Equal(t, uint64(14), createInfo.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_IMAGE_CREATE_INFO
			require.True(t, createInfo.FieldByName("pNext").IsNil())

			require.Equal(t, uint64(7), createInfo.FieldByName("extent").FieldByName("width").Uint())
			require.Equal(t, uint64(11), createInfo.FieldByName("extent").FieldByName("height").Uint())
			require.Equal(t, uint64(13), createInfo.FieldByName("extent").FieldByName("depth").Uint())
			require.Equal(t, uint64(0), createInfo.FieldByName("flags").Uint())
			require.Equal(t, uint64(4), createInfo.FieldByName("usage").Uint()) // VK_IMAGE_USAGE_SAMPLED_BIT

			requirementSlice := ([]loader.VkSparseImageMemoryRequirements2)(unsafe.Slice(pSparseMemoryRequirements, 2))
			outData := reflect.ValueOf(requirementSlice)
			element := outData.Index(0)
			require.Equal(t, uint64(1000146004), element.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SPARSE_IMAGE_MEMORY_REQUIREMENTS_2_KHR
			require.True(t, element.FieldByName("pNext").IsNil())

			memReqs := element.FieldByName("memoryRequirements")
			imageAspectFlags := (*loader.VkImageAspectFlags)(unsafe.Pointer(memReqs.FieldByName("formatProperties").FieldByName("aspectMask").UnsafeAddr()))
			*imageAspectFlags = loader.VkImageAspectFlags(0x00000008) // VK_IMAGE_ASPECT_METADATA_BIT
			width := (*loader.Uint32)(unsafe.Pointer(memReqs.FieldByName("formatProperties").FieldByName("imageGranularity").FieldByName("width").UnsafeAddr()))
			*width = loader.Uint32(1)
			height := (*loader.Uint32)(unsafe.Pointer(memReqs.FieldByName("formatProperties").FieldByName("imageGranularity").FieldByName("height").UnsafeAddr()))
			*height = loader.Uint32(3)
			depth := (*loader.Uint32)(unsafe.Pointer(memReqs.FieldByName("formatProperties").FieldByName("imageGranularity").FieldByName("depth").UnsafeAddr()))
			*depth = loader.Uint32(5)
			flags := (*loader.VkSparseImageFormatFlags)(unsafe.Pointer(memReqs.FieldByName("formatProperties").FieldByName("flags").UnsafeAddr()))
			*flags = loader.VkSparseImageFormatFlags(0x00000004) // VK_SPARSE_IMAGE_FORMAT_NONSTANDARD_BLOCK_SIZE_BIT
			*(*loader.Uint32)(unsafe.Pointer(memReqs.FieldByName("imageMipTailFirstLod").UnsafeAddr())) = loader.Uint32(7)
			*(*loader.VkDeviceSize)(unsafe.Pointer(memReqs.FieldByName("imageMipTailSize").UnsafeAddr())) = loader.VkDeviceSize(17)
			*(*loader.VkDeviceSize)(unsafe.Pointer(memReqs.FieldByName("imageMipTailOffset").UnsafeAddr())) = loader.VkDeviceSize(11)
			*(*loader.VkDeviceSize)(unsafe.Pointer(memReqs.FieldByName("imageMipTailStride").UnsafeAddr())) = loader.VkDeviceSize(13)

			element = outData.Index(1)
			require.Equal(t, uint64(1000146004), element.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SPARSE_IMAGE_MEMORY_REQUIREMENTS_2_KHR
			require.True(t, element.FieldByName("pNext").IsNil())

			memReqs = element.FieldByName("memoryRequirements")
			imageAspectFlags = (*loader.VkImageAspectFlags)(unsafe.Pointer(memReqs.FieldByName("formatProperties").FieldByName("aspectMask").UnsafeAddr()))
			*imageAspectFlags = loader.VkImageAspectFlags(0x00000004) // VK_IMAGE_ASPECT_STENCIL_BIT
			width = (*loader.Uint32)(unsafe.Pointer(memReqs.FieldByName("formatProperties").FieldByName("imageGranularity").FieldByName("width").UnsafeAddr()))
			*width = loader.Uint32(19)
			height = (*loader.Uint32)(unsafe.Pointer(memReqs.FieldByName("formatProperties").FieldByName("imageGranularity").FieldByName("height").UnsafeAddr()))
			*height = loader.Uint32(23)
			depth = (*loader.Uint32)(unsafe.Pointer(memReqs.FieldByName("formatProperties").FieldByName("imageGranularity").FieldByName("depth").UnsafeAddr()))
			*depth = loader.Uint32(29)
			flags = (*loader.VkSparseImageFormatFlags)(unsafe.Pointer(memReqs.FieldByName("formatProperties").FieldByName("flags").UnsafeAddr()))
			*flags = loader.VkSparseImageFormatFlags(0)
			*(*loader.Uint32)(unsafe.Pointer(memReqs.FieldByName("imageMipTailFirstLod").UnsafeAddr())) = loader.Uint32(43)
			*(*loader.VkDeviceSize)(unsafe.Pointer(memReqs.FieldByName("imageMipTailSize").UnsafeAddr())) = loader.VkDeviceSize(31)
			*(*loader.VkDeviceSize)(unsafe.Pointer(memReqs.FieldByName("imageMipTailOffset").UnsafeAddr())) = loader.VkDeviceSize(41)
			*(*loader.VkDeviceSize)(unsafe.Pointer(memReqs.FieldByName("imageMipTailStride").UnsafeAddr())) = loader.VkDeviceSize(37)
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
