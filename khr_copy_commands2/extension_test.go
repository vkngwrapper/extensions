package khr_copy_commands2_test

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
	"github.com/vkngwrapper/core/v3/mocks"
	"github.com/vkngwrapper/extensions/v3/khr_copy_commands2"
	khr_copy_commands2_loader "github.com/vkngwrapper/extensions/v3/khr_copy_commands2/loader"
	mock_copy_commands2 "github.com/vkngwrapper/extensions/v3/khr_copy_commands2/mocks"
	"go.uber.org/mock/gomock"
)

func TestVulkanExtension_CmdResolveImage2(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	commandPool := mocks.NewDummyCommandPool(device)
	commandBuffer := mocks.NewDummyCommandBuffer(commandPool, device)

	extLoader := mock_copy_commands2.NewMockLoader(ctrl)
	extension := khr_copy_commands2.CreateExtensionDriverFromLoader(extLoader)

	srcImage := mocks.NewDummyImage(device)
	dstImage := mocks.NewDummyImage(device)

	extLoader.EXPECT().VkCmdResolveImage2KHR(commandBuffer.Handle(), gomock.Not(gomock.Nil())).DoAndReturn(
		func(commandBuffer loader.VkCommandBuffer, pResolveImage2 *khr_copy_commands2_loader.VkResolveImageInfo2KHR) {
			val := reflect.ValueOf(pResolveImage2).Elem()

			require.Equal(t, uint64(1000337005), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_RESOLVE_IMAGE_INFO_2
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, srcImage.Handle(), loader.VkImage(val.FieldByName("srcImage").UnsafePointer()))
			require.Equal(t, dstImage.Handle(), loader.VkImage(val.FieldByName("dstImage").UnsafePointer()))
			require.Equal(t, uint64(2), val.FieldByName("srcImageLayout").Uint()) // VK_IMAGE_LAYOUT_COLOR_ATTACHMENT_OPTIMAL
			require.Equal(t, uint64(4), val.FieldByName("dstImageLayout").Uint()) // VK_IMAGE_LAYOUT_DEPTH_STENCIL_READ_ONLY_OPTIMAL
			require.Equal(t, uint64(1), val.FieldByName("regionCount").Uint())
			require.False(t, val.FieldByName("pRegions").IsNil())

			regions := (*khr_copy_commands2_loader.VkImageResolve2KHR)(unsafe.Pointer(val.FieldByName("pRegions").Elem().UnsafeAddr()))
			regionSlice := ([]khr_copy_commands2_loader.VkImageResolve2KHR)(unsafe.Slice(regions, 1))

			val = reflect.ValueOf(regionSlice).Index(0)

			require.Equal(t, uint64(2), val.FieldByName("srcSubresource").FieldByName("aspectMask").Uint()) // VK_IMAGE_ASPECT_DEPTH_BIT
			require.Equal(t, uint64(3), val.FieldByName("srcSubresource").FieldByName("mipLevel").Uint())
			require.Equal(t, uint64(5), val.FieldByName("srcSubresource").FieldByName("baseArrayLayer").Uint())
			require.Equal(t, uint64(7), val.FieldByName("srcSubresource").FieldByName("layerCount").Uint())

			require.Equal(t, uint64(8), val.FieldByName("dstSubresource").FieldByName("aspectMask").Uint()) // VK_IMAGE_ASPECT_METADATA_BIT
			require.Equal(t, uint64(11), val.FieldByName("dstSubresource").FieldByName("mipLevel").Uint())
			require.Equal(t, uint64(13), val.FieldByName("dstSubresource").FieldByName("baseArrayLayer").Uint())
			require.Equal(t, uint64(17), val.FieldByName("dstSubresource").FieldByName("layerCount").Uint())

			require.Equal(t, int64(19), val.FieldByName("srcOffset").FieldByName("x").Int())
			require.Equal(t, int64(23), val.FieldByName("srcOffset").FieldByName("y").Int())
			require.Equal(t, int64(29), val.FieldByName("srcOffset").FieldByName("z").Int())

			require.Equal(t, int64(31), val.FieldByName("dstOffset").FieldByName("x").Int())
			require.Equal(t, int64(37), val.FieldByName("dstOffset").FieldByName("y").Int())
			require.Equal(t, int64(41), val.FieldByName("dstOffset").FieldByName("z").Int())

			require.Equal(t, uint64(43), val.FieldByName("extent").FieldByName("width").Uint())
			require.Equal(t, uint64(47), val.FieldByName("extent").FieldByName("height").Uint())
			require.Equal(t, uint64(49), val.FieldByName("extent").FieldByName("depth").Uint())
		})

	err := extension.CmdResolveImage2(commandBuffer, khr_copy_commands2.ResolveImageInfo2{
		SrcImage:       srcImage,
		SrcImageLayout: core1_0.ImageLayoutColorAttachmentOptimal,
		DstImage:       dstImage,
		DstImageLayout: core1_0.ImageLayoutDepthStencilReadOnlyOptimal,
		Regions: []khr_copy_commands2.ImageResolve2{
			{
				SrcSubresource: core1_0.ImageSubresourceLayers{
					AspectMask:     core1_0.ImageAspectDepth,
					MipLevel:       3,
					BaseArrayLayer: 5,
					LayerCount:     7,
				},
				DstSubresource: core1_0.ImageSubresourceLayers{
					AspectMask:     core1_0.ImageAspectMetadata,
					MipLevel:       11,
					BaseArrayLayer: 13,
					LayerCount:     17,
				},
				SrcOffset: core1_0.Offset3D{
					X: 19,
					Y: 23,
					Z: 29,
				},
				DstOffset: core1_0.Offset3D{
					X: 31,
					Y: 37,
					Z: 41,
				},
				Extent: core1_0.Extent3D{
					Width:  43,
					Height: 47,
					Depth:  49,
				},
			},
		},
	})
	require.NoError(t, err)
}

func TestVulkanExtension_CmdCopyImageToBuffer2(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	commandPool := mocks.NewDummyCommandPool(device)
	commandBuffer := mocks.NewDummyCommandBuffer(commandPool, device)

	extLoader := mock_copy_commands2.NewMockLoader(ctrl)
	extension := khr_copy_commands2.CreateExtensionDriverFromLoader(extLoader)

	srcImage := mocks.NewDummyImage(device)
	dstBuffer := mocks.NewDummyBuffer(device)

	extLoader.EXPECT().VkCmdCopyImageToBuffer2KHR(commandBuffer.Handle(), gomock.Not(gomock.Nil())).DoAndReturn(
		func(commandBuffer loader.VkCommandBuffer, pCopyImageToBuffer2 *khr_copy_commands2_loader.VkCopyImageToBufferInfo2KHR) {
			val := reflect.ValueOf(pCopyImageToBuffer2).Elem()

			require.Equal(t, uint64(1000337003), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_COPY_IMAGE_TO_BUFFER_INFO_2
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, srcImage.Handle(), loader.VkImage(val.FieldByName("srcImage").UnsafePointer()))
			require.Equal(t, dstBuffer.Handle(), loader.VkBuffer(val.FieldByName("dstBuffer").UnsafePointer()))
			require.Equal(t, uint64(3), val.FieldByName("srcImageLayout").Uint()) // VK_IMAGE_LAYOUT_DEPTH_STENCIL_ATTACHMENT_OPTIMAL
			require.Equal(t, uint64(1), val.FieldByName("regionCount").Uint())
			require.False(t, val.FieldByName("pRegions").IsNil())

			regions := (*khr_copy_commands2_loader.VkBufferImageCopy2KHR)(unsafe.Pointer(val.FieldByName("pRegions").Elem().UnsafeAddr()))
			regionSlice := ([]khr_copy_commands2_loader.VkBufferImageCopy2KHR)(unsafe.Slice(regions, 1))

			val = reflect.ValueOf(regionSlice).Index(0)

			require.Equal(t, uint64(1000337009), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_BUFFER_IMAGE_COPY_2
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(3), val.FieldByName("bufferOffset").Uint())
			require.Equal(t, uint64(5), val.FieldByName("bufferRowLength").Uint())
			require.Equal(t, uint64(7), val.FieldByName("bufferImageHeight").Uint())
			require.Equal(t, uint64(8), val.FieldByName("imageSubresource").FieldByName("aspectMask").Uint()) // VK_IMAGE_ASPECT_METADATA_BIT
			require.Equal(t, uint64(11), val.FieldByName("imageSubresource").FieldByName("mipLevel").Uint())
			require.Equal(t, uint64(13), val.FieldByName("imageSubresource").FieldByName("layerCount").Uint())
			require.Equal(t, uint64(17), val.FieldByName("imageSubresource").FieldByName("baseArrayLayer").Uint())
			require.Equal(t, int64(19), val.FieldByName("imageOffset").FieldByName("x").Int())
			require.Equal(t, int64(23), val.FieldByName("imageOffset").FieldByName("y").Int())
			require.Equal(t, int64(29), val.FieldByName("imageOffset").FieldByName("z").Int())
			require.Equal(t, uint64(31), val.FieldByName("imageExtent").FieldByName("width").Uint())
			require.Equal(t, uint64(37), val.FieldByName("imageExtent").FieldByName("height").Uint())
			require.Equal(t, uint64(41), val.FieldByName("imageExtent").FieldByName("depth").Uint())
		})

	err := extension.CmdCopyImageToBuffer2(commandBuffer, khr_copy_commands2.CopyImageToBufferInfo2{
		SrcImage:       srcImage,
		SrcImageLayout: core1_0.ImageLayoutDepthStencilAttachmentOptimal,
		DstBuffer:      dstBuffer,

		Regions: []khr_copy_commands2.BufferImageCopy2{
			{
				BufferOffset:      3,
				BufferRowLength:   5,
				BufferImageHeight: 7,

				ImageSubresource: core1_0.ImageSubresourceLayers{
					AspectMask:     core1_0.ImageAspectMetadata,
					MipLevel:       11,
					LayerCount:     13,
					BaseArrayLayer: 17,
				},

				ImageOffset: core1_0.Offset3D{
					X: 19,
					Y: 23,
					Z: 29,
				},
				ImageExtent: core1_0.Extent3D{
					Width:  31,
					Height: 37,
					Depth:  41,
				},
			},
		},
	})
	require.NoError(t, err)
}

func TestVulkanExtension_CmdCopyImage2(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	commandPool := mocks.NewDummyCommandPool(device)
	commandBuffer := mocks.NewDummyCommandBuffer(commandPool, device)

	extLoader := mock_copy_commands2.NewMockLoader(ctrl)
	extension := khr_copy_commands2.CreateExtensionDriverFromLoader(extLoader)

	srcImage := mocks.NewDummyImage(device)
	dstImage := mocks.NewDummyImage(device)

	extLoader.EXPECT().VkCmdCopyImage2KHR(commandBuffer.Handle(), gomock.Not(gomock.Nil())).DoAndReturn(
		func(commandBuffer loader.VkCommandBuffer, pCopyImage2 *khr_copy_commands2_loader.VkCopyImageInfo2KHR) {
			val := reflect.ValueOf(pCopyImage2).Elem()

			require.Equal(t, uint64(1000337001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_COPY_IMAGE_INFO_2
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, srcImage.Handle(), loader.VkImage(val.FieldByName("srcImage").UnsafePointer()))
			require.Equal(t, dstImage.Handle(), loader.VkImage(val.FieldByName("dstImage").UnsafePointer()))
			require.Equal(t, uint64(2), val.FieldByName("srcImageLayout").Uint()) // VK_IMAGE_LAYOUT_COLOR_ATTACHMENT_OPTIMAL
			require.Equal(t, uint64(3), val.FieldByName("dstImageLayout").Uint()) // VK_IMAGE_LAYOUT_DEPTH_STENCIL_ATTACHMENT_OPTIMAL
			require.Equal(t, uint64(1), val.FieldByName("regionCount").Uint())
			require.False(t, val.FieldByName("pRegions").IsNil())

			regions := (*khr_copy_commands2_loader.VkImageCopy2KHR)(unsafe.Pointer(val.FieldByName("pRegions").Elem().UnsafeAddr()))
			regionSlice := ([]khr_copy_commands2_loader.VkImageCopy2KHR)(unsafe.Slice(regions, 1))

			val = reflect.ValueOf(regionSlice).Index(0)

			require.Equal(t, uint64(1000337007), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_IMAGE_COPY_2
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(8), val.FieldByName("srcSubresource").FieldByName("aspectMask").Uint()) // VK_IMAGE_ASPECT_METADATA_BIT
			require.Equal(t, uint64(3), val.FieldByName("srcSubresource").FieldByName("mipLevel").Uint())
			require.Equal(t, uint64(5), val.FieldByName("srcSubresource").FieldByName("baseArrayLayer").Uint())
			require.Equal(t, uint64(7), val.FieldByName("srcSubresource").FieldByName("layerCount").Uint())

			require.Equal(t, uint64(1), val.FieldByName("dstSubresource").FieldByName("aspectMask").Uint()) // VK_IMAGE_ASPECT_COLOR_BIT
			require.Equal(t, uint64(11), val.FieldByName("dstSubresource").FieldByName("mipLevel").Uint())
			require.Equal(t, uint64(13), val.FieldByName("dstSubresource").FieldByName("baseArrayLayer").Uint())
			require.Equal(t, uint64(17), val.FieldByName("dstSubresource").FieldByName("layerCount").Uint())

			require.Equal(t, int64(19), val.FieldByName("srcOffset").FieldByName("x").Int())
			require.Equal(t, int64(23), val.FieldByName("srcOffset").FieldByName("y").Int())
			require.Equal(t, int64(29), val.FieldByName("srcOffset").FieldByName("z").Int())

			require.Equal(t, int64(31), val.FieldByName("dstOffset").FieldByName("x").Int())
			require.Equal(t, int64(37), val.FieldByName("dstOffset").FieldByName("y").Int())
			require.Equal(t, int64(41), val.FieldByName("dstOffset").FieldByName("z").Int())

			require.Equal(t, uint64(43), val.FieldByName("extent").FieldByName("width").Uint())
			require.Equal(t, uint64(47), val.FieldByName("extent").FieldByName("height").Uint())
			require.Equal(t, uint64(53), val.FieldByName("extent").FieldByName("depth").Uint())
		})

	err := extension.CmdCopyImage2(commandBuffer, khr_copy_commands2.CopyImageInfo2{
		SrcImage:       srcImage,
		SrcImageLayout: core1_0.ImageLayoutColorAttachmentOptimal,
		DstImage:       dstImage,
		DstImageLayout: core1_0.ImageLayoutDepthStencilAttachmentOptimal,
		Regions: []khr_copy_commands2.ImageCopy2{
			{
				SrcSubresource: core1_0.ImageSubresourceLayers{
					AspectMask:     core1_0.ImageAspectMetadata,
					MipLevel:       3,
					BaseArrayLayer: 5,
					LayerCount:     7,
				},
				DstSubresource: core1_0.ImageSubresourceLayers{
					AspectMask:     core1_0.ImageAspectColor,
					MipLevel:       11,
					BaseArrayLayer: 13,
					LayerCount:     17,
				},
				SrcOffset: core1_0.Offset3D{
					X: 19,
					Y: 23,
					Z: 29,
				},
				DstOffset: core1_0.Offset3D{
					X: 31,
					Y: 37,
					Z: 41,
				},
				Extent: core1_0.Extent3D{
					Width:  43,
					Height: 47,
					Depth:  53,
				},
			},
		},
	})
	require.NoError(t, err)
}

func TestVulkanExtension_CmdCopyBufferToImage2(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	commandPool := mocks.NewDummyCommandPool(device)
	commandBuffer := mocks.NewDummyCommandBuffer(commandPool, device)

	extLoader := mock_copy_commands2.NewMockLoader(ctrl)
	extension := khr_copy_commands2.CreateExtensionDriverFromLoader(extLoader)

	srcBuffer := mocks.NewDummyBuffer(device)
	dstImage := mocks.NewDummyImage(device)

	extLoader.EXPECT().VkCmdCopyBufferToImage2KHR(commandBuffer.Handle(), gomock.Not(gomock.Nil())).DoAndReturn(
		func(commandBuffer loader.VkCommandBuffer, pCopyBufferToImage2 *khr_copy_commands2_loader.VkCopyBufferToImageInfo2KHR) {
			val := reflect.ValueOf(pCopyBufferToImage2).Elem()

			require.Equal(t, uint64(1000337002), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_COPY_BUFFER_TO_IMAGE_INFO_2
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, srcBuffer.Handle(), loader.VkBuffer(val.FieldByName("srcBuffer").UnsafePointer()))
			require.Equal(t, dstImage.Handle(), loader.VkImage(val.FieldByName("dstImage").UnsafePointer()))
			require.Equal(t, uint64(4), val.FieldByName("dstImageLayout").Uint()) // VK_IMAGE_LAYOUT_DEPTH_STENCIL_READ_ONLY_OPTIMAL
			require.Equal(t, uint64(1), val.FieldByName("regionCount").Uint())
			require.False(t, val.FieldByName("pRegions").IsNil())

			regions := (*khr_copy_commands2_loader.VkBufferImageCopy2KHR)(unsafe.Pointer(val.FieldByName("pRegions").Elem().UnsafeAddr()))
			regionSlice := ([]khr_copy_commands2_loader.VkBufferImageCopy2KHR)(unsafe.Slice(regions, 1))

			val = reflect.ValueOf(regionSlice).Index(0)

			require.Equal(t, uint64(1000337009), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_BUFFER_IMAGE_COPY_2
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(3), val.FieldByName("bufferOffset").Uint())
			require.Equal(t, uint64(5), val.FieldByName("bufferRowLength").Uint())
			require.Equal(t, uint64(7), val.FieldByName("bufferImageHeight").Uint())
			require.Equal(t, uint64(8), val.FieldByName("imageSubresource").FieldByName("aspectMask").Uint()) // VK_IMAGE_ASPECT_METADATA_BIT
			require.Equal(t, uint64(11), val.FieldByName("imageSubresource").FieldByName("mipLevel").Uint())
			require.Equal(t, uint64(13), val.FieldByName("imageSubresource").FieldByName("layerCount").Uint())
			require.Equal(t, uint64(17), val.FieldByName("imageSubresource").FieldByName("baseArrayLayer").Uint())
			require.Equal(t, int64(19), val.FieldByName("imageOffset").FieldByName("x").Int())
			require.Equal(t, int64(23), val.FieldByName("imageOffset").FieldByName("y").Int())
			require.Equal(t, int64(29), val.FieldByName("imageOffset").FieldByName("z").Int())
			require.Equal(t, uint64(31), val.FieldByName("imageExtent").FieldByName("width").Uint())
			require.Equal(t, uint64(37), val.FieldByName("imageExtent").FieldByName("height").Uint())
			require.Equal(t, uint64(41), val.FieldByName("imageExtent").FieldByName("depth").Uint())
		})

	err := extension.CmdCopyBufferToImage2(commandBuffer, khr_copy_commands2.CopyBufferToImageInfo2{
		SrcBuffer:      srcBuffer,
		DstImage:       dstImage,
		DstImageLayout: core1_0.ImageLayoutDepthStencilReadOnlyOptimal,
		Regions: []khr_copy_commands2.BufferImageCopy2{
			{
				BufferOffset:      3,
				BufferRowLength:   5,
				BufferImageHeight: 7,

				ImageSubresource: core1_0.ImageSubresourceLayers{
					AspectMask:     core1_0.ImageAspectMetadata,
					MipLevel:       11,
					LayerCount:     13,
					BaseArrayLayer: 17,
				},

				ImageOffset: core1_0.Offset3D{
					X: 19,
					Y: 23,
					Z: 29,
				},
				ImageExtent: core1_0.Extent3D{
					Width:  31,
					Height: 37,
					Depth:  41,
				},
			},
		},
	})
	require.NoError(t, err)
}

func TestVulkanExtension_CopyBuffer2(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	commandPool := mocks.NewDummyCommandPool(device)
	commandBuffer := mocks.NewDummyCommandBuffer(commandPool, device)

	extLoader := mock_copy_commands2.NewMockLoader(ctrl)
	extension := khr_copy_commands2.CreateExtensionDriverFromLoader(extLoader)

	srcBuffer := mocks.NewDummyBuffer(device)
	dstBuffer := mocks.NewDummyBuffer(device)

	extLoader.EXPECT().VkCmdCopyBuffer2KHR(commandBuffer.Handle(), gomock.Not(gomock.Nil())).DoAndReturn(
		func(commandBuffer loader.VkCommandBuffer, pCopyBufferInfo2 *khr_copy_commands2_loader.VkCopyBufferInfo2KHR) {
			val := reflect.ValueOf(pCopyBufferInfo2).Elem()

			require.Equal(t, uint64(1000337000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_COPY_BUFFER_INFO_2
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, srcBuffer.Handle(), loader.VkBuffer(val.FieldByName("srcBuffer").UnsafePointer()))
			require.Equal(t, dstBuffer.Handle(), loader.VkBuffer(val.FieldByName("dstBuffer").UnsafePointer()))
			require.Equal(t, uint64(2), val.FieldByName("regionCount").Uint())
			require.False(t, val.FieldByName("pRegions").IsNil())

			regions := (*khr_copy_commands2_loader.VkBufferCopy2KHR)(unsafe.Pointer(val.FieldByName("pRegions").Elem().UnsafeAddr()))
			regionSlice := ([]khr_copy_commands2_loader.VkBufferCopy2KHR)(unsafe.Slice(regions, 2))

			val = reflect.ValueOf(regionSlice)

			require.Equal(t, uint64(5), val.Index(0).FieldByName("srcOffset").Uint())
			require.Equal(t, uint64(7), val.Index(0).FieldByName("dstOffset").Uint())
			require.Equal(t, uint64(11), val.Index(0).FieldByName("size").Uint())
			require.Equal(t, uint64(13), val.Index(1).FieldByName("srcOffset").Uint())
			require.Equal(t, uint64(17), val.Index(1).FieldByName("dstOffset").Uint())
			require.Equal(t, uint64(19), val.Index(1).FieldByName("size").Uint())
		})

	err := extension.CmdCopyBuffer2(commandBuffer, khr_copy_commands2.CopyBufferInfo2{
		SrcBuffer: srcBuffer,
		DstBuffer: dstBuffer,
		Regions: []khr_copy_commands2.BufferCopy2{
			{
				SrcOffset: 5,
				DstOffset: 7,
				Size:      11,
			},
			{
				SrcOffset: 13,
				DstOffset: 17,
				Size:      19,
			},
		},
	})
	require.NoError(t, err)
}

func TestVulkanExtension_CmdBlitImage2(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	commandPool := mocks.NewDummyCommandPool(device)
	commandBuffer := mocks.NewDummyCommandBuffer(commandPool, device)

	extLoader := mock_copy_commands2.NewMockLoader(ctrl)
	extension := khr_copy_commands2.CreateExtensionDriverFromLoader(extLoader)

	srcImage := mocks.NewDummyImage(device)
	dstImage := mocks.NewDummyImage(device)

	extLoader.EXPECT().VkCmdBlitImage2KHR(commandBuffer.Handle(), gomock.Not(gomock.Nil())).DoAndReturn(
		func(commandBuffer loader.VkCommandBuffer, pBlitImageInfo2 *khr_copy_commands2_loader.VkBlitImageInfo2KHR) {
			val := reflect.ValueOf(pBlitImageInfo2).Elem()

			require.Equal(t, uint64(1000337004), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_BLIT_IMAGE_INFO_2
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, srcImage.Handle(), loader.VkImage(val.FieldByName("srcImage").UnsafePointer()))
			require.Equal(t, uint64(5), val.FieldByName("srcImageLayout").Uint()) // VK_IMAGE_LAYOUT_SHADER_READ_ONLY_OPTIMAL
			require.Equal(t, dstImage.Handle(), loader.VkImage(val.FieldByName("dstImage").UnsafePointer()))
			require.Equal(t, uint64(2), val.FieldByName("dstImageLayout").Uint()) // VK_IMAGE_LAYOUT_COLOR_ATTACHMENT_OPTIMAL
			require.Equal(t, uint64(1), val.FieldByName("regionCount").Uint())
			require.Equal(t, uint64(1), val.FieldByName("filter").Uint()) // VK_FILTER_LINEAR
			require.False(t, val.FieldByName("pRegions").IsNil())

			regions := (*khr_copy_commands2_loader.VkImageBlit2KHR)(unsafe.Pointer(val.FieldByName("pRegions").Elem().UnsafeAddr()))
			regionSlice := ([]khr_copy_commands2_loader.VkImageBlit2KHR)(unsafe.Slice(regions, 1))

			val = reflect.ValueOf(regionSlice).Index(0)

			require.Equal(t, uint64(1000337008), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_IMAGE_BLIT_2
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(1), val.FieldByName("srcSubresource").FieldByName("aspectMask").Uint()) // VK_IMAGE_ASPECT_COLOR_BIT
			require.Equal(t, uint64(5), val.FieldByName("srcSubresource").FieldByName("mipLevel").Uint())
			require.Equal(t, uint64(7), val.FieldByName("srcSubresource").FieldByName("baseArrayLayer").Uint())
			require.Equal(t, uint64(11), val.FieldByName("srcSubresource").FieldByName("layerCount").Uint())
			require.Equal(t, uint64(2), val.FieldByName("dstSubresource").FieldByName("aspectMask").Uint()) // VK_IMAGE_ASPECT_DEPTH_BIT
			require.Equal(t, uint64(37), val.FieldByName("dstSubresource").FieldByName("mipLevel").Uint())
			require.Equal(t, uint64(41), val.FieldByName("dstSubresource").FieldByName("baseArrayLayer").Uint())
			require.Equal(t, uint64(43), val.FieldByName("dstSubresource").FieldByName("layerCount").Uint())

			offsets := val.FieldByName("srcOffsets")
			require.Equal(t, int64(13), offsets.Index(0).FieldByName("x").Int())
			require.Equal(t, int64(17), offsets.Index(0).FieldByName("y").Int())
			require.Equal(t, int64(19), offsets.Index(0).FieldByName("z").Int())
			require.Equal(t, int64(23), offsets.Index(1).FieldByName("x").Int())
			require.Equal(t, int64(29), offsets.Index(1).FieldByName("y").Int())
			require.Equal(t, int64(31), offsets.Index(1).FieldByName("z").Int())

			offsets = val.FieldByName("dstOffsets")
			require.Equal(t, int64(47), offsets.Index(0).FieldByName("x").Int())
			require.Equal(t, int64(51), offsets.Index(0).FieldByName("y").Int())
			require.Equal(t, int64(53), offsets.Index(0).FieldByName("z").Int())
			require.Equal(t, int64(59), offsets.Index(1).FieldByName("x").Int())
			require.Equal(t, int64(61), offsets.Index(1).FieldByName("y").Int())
			require.Equal(t, int64(67), offsets.Index(1).FieldByName("z").Int())
		})

	err := extension.CmdBlitImage2(commandBuffer, khr_copy_commands2.BlitImageInfo2{
		SrcImage:       srcImage,
		SrcImageLayout: core1_0.ImageLayoutShaderReadOnlyOptimal,
		DstImage:       dstImage,
		DstImageLayout: core1_0.ImageLayoutColorAttachmentOptimal,
		Regions: []khr_copy_commands2.ImageBlit2{
			{
				SrcSubresource: core1_0.ImageSubresourceLayers{
					AspectMask:     core1_0.ImageAspectColor,
					MipLevel:       5,
					BaseArrayLayer: 7,
					LayerCount:     11,
				},
				SrcOffsets: [2]core1_0.Offset3D{
					{
						X: 13,
						Y: 17,
						Z: 19,
					},
					{
						X: 23,
						Y: 29,
						Z: 31,
					},
				},
				DstSubresource: core1_0.ImageSubresourceLayers{
					AspectMask:     core1_0.ImageAspectDepth,
					MipLevel:       37,
					BaseArrayLayer: 41,
					LayerCount:     43,
				},
				DstOffsets: [2]core1_0.Offset3D{
					{
						X: 47,
						Y: 51,
						Z: 53,
					},
					{
						X: 59,
						Y: 61,
						Z: 67,
					},
				},
			},
		},
		Filter: core1_0.FilterLinear,
	})
	require.NoError(t, err)
}
