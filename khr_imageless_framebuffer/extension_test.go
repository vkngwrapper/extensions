package khr_imageless_framebuffer_test

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
	"github.com/vkngwrapper/extensions/v3/khr_get_physical_device_properties2"
	khr_get_physical_device_properties2_driver "github.com/vkngwrapper/extensions/v3/khr_get_physical_device_properties2/loader"
	mock_get_physical_device_properties2 "github.com/vkngwrapper/extensions/v3/khr_get_physical_device_properties2/mocks"
	"github.com/vkngwrapper/extensions/v3/khr_imageless_framebuffer"
	khr_imageless_framebuffer_driver "github.com/vkngwrapper/extensions/v3/khr_imageless_framebuffer/loader"
	"go.uber.org/mock/gomock"
)

func TestFramebufferAttachmentsCreateOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	mockFramebuffer := mocks.NewDummyFramebuffer(device)

	coreLoader := mock_driver.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := mocks1_0.InternalDeviceDriver(device, coreLoader)

	coreLoader.EXPECT().VkCreateFramebuffer(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device loader.VkDevice,
		pCreateInfo *loader.VkFramebufferCreateInfo,
		pAllocator *loader.VkAllocationCallbacks,
		pFramebuffer *loader.VkFramebuffer) (common.VkResult, error) {

		*pFramebuffer = mockFramebuffer.Handle()

		val := reflect.ValueOf(pCreateInfo).Elem()
		require.Equal(t, uint64(37), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_FRAMEBUFFER_CREATE_INFO

		next := (*khr_imageless_framebuffer_driver.VkFramebufferAttachmentsCreateInfoKHR)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000108001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_FRAMEBUFFER_ATTACHMENTS_CREATE_INFO_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(2), val.FieldByName("attachmentImageInfoCount").Uint())

		imageInfos := (*khr_imageless_framebuffer_driver.VkFramebufferAttachmentImageInfoKHR)(val.FieldByName("pAttachmentImageInfos").UnsafePointer())
		imageInfoSlice := unsafe.Slice(imageInfos, 2)
		val = reflect.ValueOf(imageInfoSlice)

		info := val.Index(0)
		require.Equal(t, uint64(1000108002), info.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_FRAMEBUFFER_ATTACHMENT_IMAGE_INFO_KHR
		require.True(t, info.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(0x10), info.FieldByName("flags").Uint()) // VK_IMAGE_CREATE_CUBE_COMPATIBLE_BIT
		require.Equal(t, uint64(4), info.FieldByName("usage").Uint())    // VK_IMAGE_USAGE_SAMPLED_BIT
		require.Equal(t, uint64(1), info.FieldByName("width").Uint())
		require.Equal(t, uint64(3), info.FieldByName("height").Uint())
		require.Equal(t, uint64(5), info.FieldByName("layerCount").Uint())
		require.Equal(t, uint64(2), info.FieldByName("viewFormatCount").Uint())

		viewFormats := (*loader.VkFormat)(info.FieldByName("pViewFormats").UnsafePointer())
		viewFormatSlice := unsafe.Slice(viewFormats, 2)

		require.Equal(t, []loader.VkFormat{68, 53}, viewFormatSlice)

		info = val.Index(1)
		require.Equal(t, uint64(1000108002), info.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_FRAMEBUFFER_ATTACHMENT_IMAGE_INFO_KHR
		require.True(t, info.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1), info.FieldByName("flags").Uint())    // VK_IMAGE_CREATE_SPARSE_BINDING_BIT
		require.Equal(t, uint64(0x10), info.FieldByName("usage").Uint()) // VK_IMAGE_USAGE_COLOR_ATTACHMENT_BIT
		require.Equal(t, uint64(7), info.FieldByName("width").Uint())
		require.Equal(t, uint64(11), info.FieldByName("height").Uint())
		require.Equal(t, uint64(13), info.FieldByName("layerCount").Uint())
		require.Equal(t, uint64(3), info.FieldByName("viewFormatCount").Uint())

		viewFormats = (*loader.VkFormat)(info.FieldByName("pViewFormats").UnsafePointer())
		viewFormatSlice = unsafe.Slice(viewFormats, 3)

		require.Equal(t, []loader.VkFormat{161, 164, 163}, viewFormatSlice)

		return core1_0.VKSuccess, nil
	})

	framebuffer, _, err := driver.CreateFramebuffer(
		nil,
		core1_0.FramebufferCreateInfo{
			NextOptions: common.NextOptions{
				khr_imageless_framebuffer.FramebufferAttachmentsCreateInfo{
					AttachmentImageInfos: []khr_imageless_framebuffer.FramebufferAttachmentImageInfo{
						{
							Flags:      core1_0.ImageCreateCubeCompatible,
							Usage:      core1_0.ImageUsageSampled,
							Width:      1,
							Height:     3,
							LayerCount: 5,
							ViewFormats: []core1_0.Format{
								core1_0.FormatA2B10G10R10UnsignedIntPacked,
								core1_0.FormatA8B8G8R8UnsignedScaledPacked,
							},
						},
						{
							Flags:      core1_0.ImageCreateSparseBinding,
							Usage:      core1_0.ImageUsageColorAttachment,
							Width:      7,
							Height:     11,
							LayerCount: 13,
							ViewFormats: []core1_0.Format{
								core1_0.FormatASTC5x5_UnsignedNormalized,
								core1_0.FormatASTC6x5_sRGB,
								core1_0.FormatASTC6x5_UnsignedNormalized,
							},
						},
					},
				},
			},
		})
	require.NoError(t, err)
	require.Equal(t, mockFramebuffer.Handle(), framebuffer.Handle())
}

func TestPhysicalDeviceImagelessFramebufferFeaturesOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	instance := mocks.NewDummyInstance(common.Vulkan1_0, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_0)
	mockDevice := mocks.NewDummyDevice(common.Vulkan1_0, []string{})

	coreLoader := mock_driver.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := mocks1_0.InternalCoreInstanceDriver(instance, coreLoader)

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

		next := (*khr_imageless_framebuffer_driver.VkPhysicalDeviceImagelessFramebufferFeaturesKHR)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000108000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGELESS_FRAMEBUFFER_FEATURES_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1), val.FieldByName("imagelessFramebuffer").Uint())

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
				khr_imageless_framebuffer.PhysicalDeviceImagelessFramebufferFeatures{
					ImagelessFramebuffer: true,
				},
			},
		})
	require.NoError(t, err)
	require.Equal(t, mockDevice.Handle(), device.Device().Handle())
}

func TestPhysicalDeviceImagelessFramebufferFeaturesOutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_get_physical_device_properties2.NewMockLoader(ctrl)
	extension := khr_get_physical_device_properties2.CreateExtensionDriverFromLoader(extDriver)

	instance := mocks.NewDummyInstance(common.Vulkan1_0, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_0)

	extDriver.EXPECT().VkGetPhysicalDeviceFeatures2KHR(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice loader.VkPhysicalDevice,
		pFeatures *khr_get_physical_device_properties2_driver.VkPhysicalDeviceFeatures2KHR) {

		val := reflect.ValueOf(pFeatures).Elem()
		require.Equal(t, uint64(1000059000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FEATURES_2_KHR

		next := (*khr_imageless_framebuffer_driver.VkPhysicalDeviceImagelessFramebufferFeaturesKHR)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000108000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGELESS_FRAMEBUFFER_FEATURES_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("imagelessFramebuffer").UnsafeAddr())) = loader.VkBool32(1)
	})

	var outData khr_imageless_framebuffer.PhysicalDeviceImagelessFramebufferFeatures
	err := extension.GetPhysicalDeviceFeatures2(
		physicalDevice,
		&khr_get_physical_device_properties2.PhysicalDeviceFeatures2{
			NextOutData: common.NextOutData{&outData},
		},
	)
	require.NoError(t, err)
	require.Equal(t, khr_imageless_framebuffer.PhysicalDeviceImagelessFramebufferFeatures{
		ImagelessFramebuffer: true,
	}, outData)
}

func TestRenderPassAttachmentBeginInfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	commandPool := mocks.NewDummyCommandPool(device)
	commandBuffer := mocks.NewDummyCommandBuffer(commandPool, device)

	coreLoader := mock_driver.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := mocks1_0.InternalDeviceDriver(device, coreLoader)

	imageView1 := mocks.NewDummyImageView(device)
	imageView2 := mocks.NewDummyImageView(device)

	coreLoader.EXPECT().VkCmdBeginRenderPass(
		commandBuffer.Handle(),
		gomock.Not(gomock.Nil()),
		loader.VkSubpassContents(0), // VK_SUBPASS_CONTENTS_INLINE
	).DoAndReturn(func(commandBuffer loader.VkCommandBuffer,
		pRenderPassBegin *loader.VkRenderPassBeginInfo,
		contents loader.VkSubpassContents) {

		val := reflect.ValueOf(pRenderPassBegin).Elem()
		require.Equal(t, uint64(43), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_RENDER_PASS_BEGIN_INFO

		next := (*khr_imageless_framebuffer_driver.VkRenderPassAttachmentBeginInfoKHR)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000108003), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_RENDER_PASS_ATTACHMENT_BEGIN_INFO_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(2), val.FieldByName("attachmentCount").Uint())

		firstImageView := val.FieldByName("pAttachments").UnsafePointer()
		require.Equal(t, imageView1.Handle(), *(*loader.VkImageView)(firstImageView))

		secondImageView := unsafe.Add(firstImageView, unsafe.Sizeof(uintptr(0)))
		require.Equal(t, imageView2.Handle(), *(*loader.VkImageView)(secondImageView))
	})

	err := driver.CmdBeginRenderPass(commandBuffer, core1_0.SubpassContentsInline, core1_0.RenderPassBeginInfo{
		NextOptions: common.NextOptions{khr_imageless_framebuffer.RenderPassAttachmentBeginInfo{
			Attachments: []core1_0.ImageView{imageView1, imageView2},
		}},
	})
	require.NoError(t, err)
}
