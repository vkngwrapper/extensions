package khr_get_memory_requirements2_test

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
	"github.com/vkngwrapper/core/v3/mocks"
	"github.com/vkngwrapper/extensions/v3/khr_get_memory_requirements2"
	khr_get_memory_requirements2_driver "github.com/vkngwrapper/extensions/v3/khr_get_memory_requirements2/loader"
	mock_get_memory_requirements2 "github.com/vkngwrapper/extensions/v3/khr_get_memory_requirements2/mocks"
	"go.uber.org/mock/gomock"
)

func TestVulkanExtension_BufferMemoryRequirements(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	buffer := mocks.NewDummyBuffer(device)

	extDriver := mock_get_memory_requirements2.NewMockLoader(ctrl)
	extension := khr_get_memory_requirements2.CreateExtensionDriverFromLoader(extDriver, device)

	extDriver.EXPECT().VkGetBufferMemoryRequirements2KHR(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device loader.VkDevice,
		pInfo *khr_get_memory_requirements2_driver.VkBufferMemoryRequirementsInfo2KHR,
		pMemoryRequirements *khr_get_memory_requirements2_driver.VkMemoryRequirements2KHR,
	) {
		val := reflect.ValueOf(pInfo).Elem()
		require.Equal(t, uint64(1000146000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_BUFFER_MEMORY_REQUIREMENTS_INFO_2_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, buffer.Handle(), loader.VkBuffer(val.FieldByName("buffer").UnsafePointer()))

		val = reflect.ValueOf(pMemoryRequirements).Elem()
		require.Equal(t, uint64(1000146003), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_MEMORY_REQUIREMENTS_2_KHR
		require.True(t, val.FieldByName("pNext").IsNil())

		val = val.FieldByName("memoryRequirements")
		*(*loader.VkDeviceSize)(unsafe.Pointer(val.FieldByName("size").UnsafeAddr())) = loader.VkDeviceSize(1)
		*(*loader.VkDeviceSize)(unsafe.Pointer(val.FieldByName("alignment").UnsafeAddr())) = loader.VkDeviceSize(3)
		*(*loader.Uint32)(unsafe.Pointer(val.FieldByName("memoryTypeBits").UnsafeAddr())) = loader.Uint32(5)
	})

	var outData khr_get_memory_requirements2.MemoryRequirements2
	err := extension.GetBufferMemoryRequirements2(
		khr_get_memory_requirements2.BufferMemoryRequirementsInfo2{
			Buffer: buffer,
		}, &outData)
	require.NoError(t, err)

	require.Equal(t, 1, outData.MemoryRequirements.Size)
	require.Equal(t, 3, outData.MemoryRequirements.Alignment)
	require.Equal(t, uint32(5), outData.MemoryRequirements.MemoryTypeBits)
}

func TestVulkanExtension_ImageMemoryRequirements(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	image := mocks.NewDummyImage(device)

	extDriver := mock_get_memory_requirements2.NewMockLoader(ctrl)
	extension := khr_get_memory_requirements2.CreateExtensionDriverFromLoader(extDriver, device)

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
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, image.Handle(), loader.VkImage(val.FieldByName("image").UnsafePointer()))

		val = reflect.ValueOf(pMemoryRequirements).Elem()
		require.Equal(t, uint64(1000146003), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_MEMORY_REQUIREMENTS_2_KHR
		require.True(t, val.FieldByName("pNext").IsNil())

		val = val.FieldByName("memoryRequirements")
		*(*loader.VkDeviceSize)(unsafe.Pointer(val.FieldByName("size").UnsafeAddr())) = loader.VkDeviceSize(1)
		*(*loader.VkDeviceSize)(unsafe.Pointer(val.FieldByName("alignment").UnsafeAddr())) = loader.VkDeviceSize(3)
		*(*loader.Uint32)(unsafe.Pointer(val.FieldByName("memoryTypeBits").UnsafeAddr())) = loader.Uint32(5)
	})

	var outData khr_get_memory_requirements2.MemoryRequirements2
	err := extension.GetImageMemoryRequirements2(
		khr_get_memory_requirements2.ImageMemoryRequirementsInfo2{
			Image: image,
		}, &outData)
	require.NoError(t, err)

	require.Equal(t, 1, outData.MemoryRequirements.Size)
	require.Equal(t, 3, outData.MemoryRequirements.Alignment)
	require.Equal(t, uint32(5), outData.MemoryRequirements.MemoryTypeBits)
}

func TestVulkanExtension_SparseImageMemoryRequirements(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	image := mocks.NewDummyImage(device)

	extDriver := mock_get_memory_requirements2.NewMockLoader(ctrl)
	extension := khr_get_memory_requirements2.CreateExtensionDriverFromLoader(extDriver, device)

	extDriver.EXPECT().VkGetImageSparseMemoryRequirements2KHR(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
	).DoAndReturn(
		func(device loader.VkDevice,
			pInfo *khr_get_memory_requirements2_driver.VkImageSparseMemoryRequirementsInfo2KHR,
			pSparseMemoryRequirementCount *loader.Uint32,
			pSparseMemoryRequirements *khr_get_memory_requirements2_driver.VkSparseImageMemoryRequirements2KHR) {

			options := reflect.ValueOf(pInfo).Elem()
			require.Equal(t, uint64(1000146002), options.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_IMAGE_SPARSE_MEMORY_REQUIREMENTS_INFO_2_KHR
			require.True(t, options.FieldByName("pNext").IsNil())
			require.Equal(t, image.Handle(), (loader.VkImage)(options.FieldByName("image").UnsafePointer()))

			*pSparseMemoryRequirementCount = loader.Uint32(2)
		})

	extDriver.EXPECT().VkGetImageSparseMemoryRequirements2KHR(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(device loader.VkDevice,
			pInfo *khr_get_memory_requirements2_driver.VkImageSparseMemoryRequirementsInfo2KHR,
			pSparseMemoryRequirementCount *loader.Uint32,
			pSparseMemoryRequirements *khr_get_memory_requirements2_driver.VkSparseImageMemoryRequirements2KHR) {

			require.Equal(t, loader.Uint32(2), *pSparseMemoryRequirementCount)

			options := reflect.ValueOf(pInfo).Elem()
			require.Equal(t, uint64(1000146002), options.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_IMAGE_SPARSE_MEMORY_REQUIREMENTS_INFO_2_KHR
			require.True(t, options.FieldByName("pNext").IsNil())
			require.Equal(t, image.Handle(), (loader.VkImage)(options.FieldByName("image").UnsafePointer()))

			requirementSlice := ([]khr_get_memory_requirements2_driver.VkSparseImageMemoryRequirements2KHR)(unsafe.Slice(pSparseMemoryRequirements, 2))
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

	outData, err := extension.GetImageSparseMemoryRequirements2(
		khr_get_memory_requirements2.ImageSparseMemoryRequirementsInfo2{
			Image: image,
		}, nil)
	require.NoError(t, err)
	require.Equal(t, []*khr_get_memory_requirements2.SparseImageMemoryRequirements2{
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
