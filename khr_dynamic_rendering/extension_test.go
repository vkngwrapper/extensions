package khr_dynamic_rendering_test

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
	"github.com/vkngwrapper/extensions/v3/khr_depth_stencil_resolve"
	"github.com/vkngwrapper/extensions/v3/khr_dynamic_rendering"
	khr_dynamic_rendering_loader "github.com/vkngwrapper/extensions/v3/khr_dynamic_rendering/loader"
	mock_dynamic_rendering "github.com/vkngwrapper/extensions/v3/khr_dynamic_rendering/mocks"
	"github.com/vkngwrapper/extensions/v3/khr_get_physical_device_properties2"
	khr_get_physical_device_properties2_loader "github.com/vkngwrapper/extensions/v3/khr_get_physical_device_properties2/loader"
	mock_get_physical_device_properties2 "github.com/vkngwrapper/extensions/v3/khr_get_physical_device_properties2/mocks"
	"go.uber.org/mock/gomock"
)

func TestCreateExtensionDriverFromCoreDriver_MissingExt(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	coreLoader := mock_loader.NewMockLoader(ctrl)
	driver := mocks1_0.InternalDeviceDriver(device, coreLoader)
	require.Nil(t, khr_dynamic_rendering.CreateExtensionDriverFromCoreDriver(driver))
}

func TestVulkanExtension_UninitializedCommandBuffer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	extDriver := mock_dynamic_rendering.NewMockLoader(ctrl)
	extension := khr_dynamic_rendering.CreateExtensionDriverFromLoader(extDriver, device)

	require.PanicsWithValue(t, "commandBuffer cannot be uninitialized", func() {
		_ = extension.CmdBeginRendering(core1_0.CommandBuffer{}, khr_dynamic_rendering.RenderingInfo{})
	})
	require.PanicsWithValue(t, "commandBuffer cannot be uninitialized", func() {
		_ = extension.CmdEndRendering(core1_0.CommandBuffer{})
	})
}

func TestVulkanExtension_CmdBeginRendering(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	commandPool := mocks.NewDummyCommandPool(device)
	commandBuffer := mocks.NewDummyCommandBuffer(commandPool, device)
	colorView := mocks.NewDummyImageView(device)
	resolveView := mocks.NewDummyImageView(device)
	depthView := mocks.NewDummyImageView(device)

	extDriver := mock_dynamic_rendering.NewMockLoader(ctrl)
	extension := khr_dynamic_rendering.CreateExtensionDriverFromLoader(extDriver, device)

	extDriver.EXPECT().VkCmdBeginRenderingKHR(
		commandBuffer.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(commandBuffer loader.VkCommandBuffer, pRenderingInfo *khr_dynamic_rendering_loader.VkRenderingInfoKHR) {
		val := reflect.ValueOf(pRenderingInfo).Elem()
		require.Equal(t, uint64(1000044000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_RENDERING_INFO_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(6), val.FieldByName("flags").Uint()) // VK_RENDERING_SUSPENDING_BIT_KHR | VK_RENDERING_RESUMING_BIT_KHR
		require.Equal(t, int64(1), val.FieldByName("renderArea").FieldByName("offset").FieldByName("x").Int())
		require.Equal(t, int64(3), val.FieldByName("renderArea").FieldByName("offset").FieldByName("y").Int())
		require.Equal(t, uint64(5), val.FieldByName("renderArea").FieldByName("extent").FieldByName("width").Uint())
		require.Equal(t, uint64(7), val.FieldByName("renderArea").FieldByName("extent").FieldByName("height").Uint())
		require.Equal(t, uint64(11), val.FieldByName("layerCount").Uint())
		require.Equal(t, uint64(13), val.FieldByName("viewMask").Uint())
		require.Equal(t, uint64(2), val.FieldByName("colorAttachmentCount").Uint())
		require.True(t, val.FieldByName("pStencilAttachment").IsNil())

		attachments := (*khr_dynamic_rendering_loader.VkRenderingAttachmentInfoKHR)(val.FieldByName("pColorAttachments").UnsafePointer())
		attachmentSlice := unsafe.Slice(attachments, 2)
		attachment := reflect.ValueOf(attachmentSlice).Index(0)
		require.Equal(t, uint64(1000044001), attachment.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_RENDERING_ATTACHMENT_INFO_KHR
		require.True(t, attachment.FieldByName("pNext").IsNil())
		require.Equal(t, colorView.Handle(), loader.VkImageView(attachment.FieldByName("imageView").UnsafePointer()))
		require.Equal(t, uint64(2), attachment.FieldByName("imageLayout").Uint()) // VK_IMAGE_LAYOUT_COLOR_ATTACHMENT_OPTIMAL
		require.Equal(t, uint64(2), attachment.FieldByName("resolveMode").Uint()) // VK_RESOLVE_MODE_AVERAGE_BIT_KHR
		require.Equal(t, resolveView.Handle(), loader.VkImageView(attachment.FieldByName("resolveImageView").UnsafePointer()))
		require.Equal(t, uint64(1), attachment.FieldByName("resolveImageLayout").Uint()) // VK_IMAGE_LAYOUT_GENERAL
		require.Equal(t, uint64(1), attachment.FieldByName("loadOp").Uint())             // VK_ATTACHMENT_LOAD_OP_CLEAR
		require.Equal(t, uint64(0), attachment.FieldByName("storeOp").Uint())            // VK_ATTACHMENT_STORE_OP_STORE
		clear := (*loader.Float)(unsafe.Pointer(attachment.FieldByName("clearValue").UnsafeAddr()))
		clearSlice := unsafe.Slice(clear, 4)
		require.InDelta(t, 0.25, float32(clearSlice[0]), 0.0001)
		require.InDelta(t, 0.5, float32(clearSlice[1]), 0.0001)
		require.InDelta(t, 0.75, float32(clearSlice[2]), 0.0001)
		require.InDelta(t, 1, float32(clearSlice[3]), 0.0001)

		attachment = reflect.ValueOf(attachmentSlice).Index(1)
		require.Equal(t, uint64(1000044001), attachment.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_RENDERING_ATTACHMENT_INFO_KHR
		require.True(t, attachment.FieldByName("pNext").IsNil())
		require.True(t, attachment.FieldByName("imageView").IsNil())
		require.Equal(t, uint64(0), attachment.FieldByName("imageLayout").Uint()) // VK_IMAGE_LAYOUT_UNDEFINED
		require.Equal(t, uint64(0), attachment.FieldByName("resolveMode").Uint()) // VK_RESOLVE_MODE_NONE_KHR
		require.True(t, attachment.FieldByName("resolveImageView").IsNil())
		require.Equal(t, uint64(0), attachment.FieldByName("resolveImageLayout").Uint()) // VK_IMAGE_LAYOUT_UNDEFINED
		require.Equal(t, uint64(2), attachment.FieldByName("loadOp").Uint())             // VK_ATTACHMENT_LOAD_OP_DONT_CARE
		require.Equal(t, uint64(1000301000), attachment.FieldByName("storeOp").Uint())   // VK_ATTACHMENT_STORE_OP_NONE_KHR
		clearBytes := unsafe.Slice((*byte)(unsafe.Pointer(attachment.FieldByName("clearValue").UnsafeAddr())), 16)
		require.Equal(t, make([]byte, 16), clearBytes)

		attachment = val.FieldByName("pDepthAttachment").Elem()
		require.Equal(t, uint64(1000044001), attachment.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_RENDERING_ATTACHMENT_INFO_KHR
		require.True(t, attachment.FieldByName("pNext").IsNil())
		require.Equal(t, depthView.Handle(), loader.VkImageView(attachment.FieldByName("imageView").UnsafePointer()))
		require.Equal(t, uint64(3), attachment.FieldByName("imageLayout").Uint()) // VK_IMAGE_LAYOUT_DEPTH_STENCIL_ATTACHMENT_OPTIMAL
		require.Equal(t, uint64(0), attachment.FieldByName("resolveMode").Uint()) // VK_RESOLVE_MODE_NONE_KHR
		require.True(t, attachment.FieldByName("resolveImageView").IsNil())
		require.Equal(t, uint64(0), attachment.FieldByName("resolveImageLayout").Uint()) // VK_IMAGE_LAYOUT_UNDEFINED
		require.Equal(t, uint64(1), attachment.FieldByName("loadOp").Uint())             // VK_ATTACHMENT_LOAD_OP_CLEAR
		require.Equal(t, uint64(1), attachment.FieldByName("storeOp").Uint())            // VK_ATTACHMENT_STORE_OP_DONT_CARE
		clearPointer := unsafe.Pointer(attachment.FieldByName("clearValue").UnsafeAddr())
		require.InDelta(t, 0.5, float32(*(*loader.Float)(clearPointer)), 0.0001)
		require.Equal(t, uint32(17), *(*uint32)(unsafe.Add(clearPointer, 4)))
	})

	err := extension.CmdBeginRendering(commandBuffer, khr_dynamic_rendering.RenderingInfo{
		Flags:      khr_dynamic_rendering.RenderingSuspending | khr_dynamic_rendering.RenderingResuming,
		RenderArea: core1_0.Rect2D{Offset: core1_0.Offset2D{X: 1, Y: 3}, Extent: core1_0.Extent2D{Width: 5, Height: 7}},
		LayerCount: 11,
		ViewMask:   13,
		ColorAttachments: []khr_dynamic_rendering.RenderingAttachmentInfo{
			{
				ImageView:          colorView,
				ImageLayout:        core1_0.ImageLayoutColorAttachmentOptimal,
				ResolveMode:        khr_depth_stencil_resolve.ResolveModeAverage,
				ResolveImageView:   resolveView,
				ResolveImageLayout: core1_0.ImageLayoutGeneral,
				LoadOp:             core1_0.AttachmentLoadOpClear,
				StoreOp:            core1_0.AttachmentStoreOpStore,
				ClearValue:         core1_0.ClearValueFloat{0.25, 0.5, 0.75, 1},
			},
			{
				LoadOp:  core1_0.AttachmentLoadOpDontCare,
				StoreOp: khr_dynamic_rendering.AttachmentStoreOpNone,
			},
		},
		DepthAttachment: &khr_dynamic_rendering.RenderingAttachmentInfo{
			ImageView:   depthView,
			ImageLayout: core1_0.ImageLayoutDepthStencilAttachmentOptimal,
			LoadOp:      core1_0.AttachmentLoadOpClear,
			StoreOp:     core1_0.AttachmentStoreOpDontCare,
			ClearValue:  core1_0.ClearValueDepthStencil{Depth: 0.5, Stencil: 17},
		},
	})
	require.NoError(t, err)
}

func TestVulkanExtension_CmdBeginRendering_Minimal(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	commandPool := mocks.NewDummyCommandPool(device)
	commandBuffer := mocks.NewDummyCommandBuffer(commandPool, device)
	extDriver := mock_dynamic_rendering.NewMockLoader(ctrl)
	extension := khr_dynamic_rendering.CreateExtensionDriverFromLoader(extDriver, device)

	extDriver.EXPECT().VkCmdBeginRenderingKHR(commandBuffer.Handle(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(commandBuffer loader.VkCommandBuffer, pRenderingInfo *khr_dynamic_rendering_loader.VkRenderingInfoKHR) {
			val := reflect.ValueOf(pRenderingInfo).Elem()
			require.Equal(t, uint64(1000044000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_RENDERING_INFO_KHR
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0), val.FieldByName("flags").Uint())
			require.Equal(t, int64(0), val.FieldByName("renderArea").FieldByName("offset").FieldByName("x").Int())
			require.Equal(t, int64(0), val.FieldByName("renderArea").FieldByName("offset").FieldByName("y").Int())
			require.Equal(t, uint64(1), val.FieldByName("renderArea").FieldByName("extent").FieldByName("width").Uint())
			require.Equal(t, uint64(1), val.FieldByName("renderArea").FieldByName("extent").FieldByName("height").Uint())
			require.Equal(t, uint64(1), val.FieldByName("layerCount").Uint())
			require.Equal(t, uint64(0), val.FieldByName("viewMask").Uint())
			require.Equal(t, uint64(0), val.FieldByName("colorAttachmentCount").Uint())
			require.True(t, val.FieldByName("pColorAttachments").IsNil())
			require.True(t, val.FieldByName("pDepthAttachment").IsNil())
			require.True(t, val.FieldByName("pStencilAttachment").IsNil())
		})

	err := extension.CmdBeginRendering(commandBuffer, khr_dynamic_rendering.RenderingInfo{
		RenderArea: core1_0.Rect2D{Extent: core1_0.Extent2D{Width: 1, Height: 1}},
		LayerCount: 1,
	})
	require.NoError(t, err)
}

func TestVulkanExtension_CmdEndRendering(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	commandPool := mocks.NewDummyCommandPool(device)
	commandBuffer := mocks.NewDummyCommandBuffer(commandPool, device)
	extDriver := mock_dynamic_rendering.NewMockLoader(ctrl)
	extension := khr_dynamic_rendering.CreateExtensionDriverFromLoader(extDriver, device)

	extDriver.EXPECT().VkCmdEndRenderingKHR(commandBuffer.Handle())
	require.NoError(t, extension.CmdEndRendering(commandBuffer))
}

func TestPipelineRenderingCreateInfo_AsPipelineNext(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	expectedPipeline := mocks.NewDummyPipeline(device)
	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := mocks1_0.InternalDeviceDriver(device, coreLoader)

	coreLoader.EXPECT().VkCreateGraphicsPipelines(device.Handle(), loader.VkPipelineCache(0), loader.Uint32(1), gomock.Not(gomock.Nil()), gomock.Nil(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(device loader.VkDevice, pipelineCache loader.VkPipelineCache, createInfoCount loader.Uint32, pCreateInfos *loader.VkGraphicsPipelineCreateInfo, pAllocator *loader.VkAllocationCallbacks, pPipelines *loader.VkPipeline) (common.VkResult, error) {
			*pPipelines = expectedPipeline.Handle()
			val := reflect.ValueOf(pCreateInfos).Elem()
			require.Equal(t, uint64(28), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_GRAPHICS_PIPELINE_CREATE_INFO
			require.True(t, val.FieldByName("renderPass").IsNil())

			next := (*khr_dynamic_rendering_loader.VkPipelineRenderingCreateInfoKHR)(val.FieldByName("pNext").UnsafePointer())
			val = reflect.ValueOf(next).Elem()
			require.Equal(t, uint64(1000044002), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PIPELINE_RENDERING_CREATE_INFO_KHR
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(3), val.FieldByName("viewMask").Uint())
			require.Equal(t, uint64(2), val.FieldByName("colorAttachmentCount").Uint())
			formats := (*loader.VkFormat)(val.FieldByName("pColorAttachmentFormats").UnsafePointer())
			require.Equal(t, []loader.VkFormat{37, 44}, unsafe.Slice(formats, 2))            // VK_FORMAT_R8G8B8A8_UNORM, VK_FORMAT_B8G8R8A8_UNORM
			require.Equal(t, uint64(126), val.FieldByName("depthAttachmentFormat").Uint())   // VK_FORMAT_D32_SFLOAT
			require.Equal(t, uint64(127), val.FieldByName("stencilAttachmentFormat").Uint()) // VK_FORMAT_S8_UINT
			return core1_0.VKSuccess, nil
		})

	pipelines, _, err := driver.CreateGraphicsPipelines(nil, nil, core1_0.GraphicsPipelineCreateInfo{
		NextOptions: common.NextOptions{Next: khr_dynamic_rendering.PipelineRenderingCreateInfo{
			ViewMask:                3,
			ColorAttachmentFormats:  []core1_0.Format{core1_0.FormatR8G8B8A8UnsignedNormalized, core1_0.FormatB8G8R8A8UnsignedNormalized},
			DepthAttachmentFormat:   core1_0.FormatD32SignedFloat,
			StencilAttachmentFormat: core1_0.FormatS8UnsignedInt,
		}},
	})
	require.NoError(t, err)
	require.Len(t, pipelines, 1)
	require.Equal(t, expectedPipeline.Handle(), pipelines[0].Handle())
}

func TestCommandBufferInheritanceRenderingInfo_AsNext(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	commandPool := mocks.NewDummyCommandPool(device)
	commandBuffer := mocks.NewDummyCommandBuffer(commandPool, device)
	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := mocks1_0.InternalDeviceDriver(device, coreLoader)

	coreLoader.EXPECT().VkBeginCommandBuffer(commandBuffer.Handle(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(commandBuffer loader.VkCommandBuffer, pBeginInfo *loader.VkCommandBufferBeginInfo) (common.VkResult, error) {
			val := reflect.ValueOf(pBeginInfo).Elem()
			require.Equal(t, uint64(42), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_COMMAND_BUFFER_BEGIN_INFO
			require.True(t, val.FieldByName("pNext").IsNil())
			val = val.FieldByName("pInheritanceInfo").Elem()
			require.Equal(t, uint64(41), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_COMMAND_BUFFER_INHERITANCE_INFO
			require.True(t, val.FieldByName("renderPass").IsNil())
			next := (*khr_dynamic_rendering_loader.VkCommandBufferInheritanceRenderingInfoKHR)(val.FieldByName("pNext").UnsafePointer())
			val = reflect.ValueOf(next).Elem()
			require.Equal(t, uint64(1000044004), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_COMMAND_BUFFER_INHERITANCE_RENDERING_INFO_KHR
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(2), val.FieldByName("flags").Uint()) // VK_RENDERING_SUSPENDING_BIT_KHR
			require.Equal(t, uint64(5), val.FieldByName("viewMask").Uint())
			require.Equal(t, uint64(2), val.FieldByName("colorAttachmentCount").Uint())
			formats := (*loader.VkFormat)(val.FieldByName("pColorAttachmentFormats").UnsafePointer())
			require.Equal(t, []loader.VkFormat{44, 37}, unsafe.Slice(formats, 2))            // VK_FORMAT_B8G8R8A8_UNORM, VK_FORMAT_R8G8B8A8_UNORM
			require.Equal(t, uint64(126), val.FieldByName("depthAttachmentFormat").Uint())   // VK_FORMAT_D32_SFLOAT
			require.Equal(t, uint64(127), val.FieldByName("stencilAttachmentFormat").Uint()) // VK_FORMAT_S8_UINT
			require.Equal(t, uint64(4), val.FieldByName("rasterizationSamples").Uint())      // VK_SAMPLE_COUNT_4_BIT
			return core1_0.VKSuccess, nil
		})

	_, err := driver.BeginCommandBuffer(commandBuffer, core1_0.CommandBufferBeginInfo{
		Flags: core1_0.CommandBufferUsageRenderPassContinue,
		InheritanceInfo: &core1_0.CommandBufferInheritanceInfo{
			NextOptions: common.NextOptions{Next: khr_dynamic_rendering.CommandBufferInheritanceRenderingInfo{
				Flags:                   khr_dynamic_rendering.RenderingSuspending,
				ViewMask:                5,
				ColorAttachmentFormats:  []core1_0.Format{core1_0.FormatB8G8R8A8UnsignedNormalized, core1_0.FormatR8G8B8A8UnsignedNormalized},
				DepthAttachmentFormat:   core1_0.FormatD32SignedFloat,
				StencilAttachmentFormat: core1_0.FormatS8UnsignedInt,
				RasterizationSamples:    core1_0.Samples4,
			}},
		},
	})
	require.NoError(t, err)
}

func TestPhysicalDeviceDynamicRenderingFeatures_Options(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	instance := mocks.NewDummyInstance(common.Vulkan1_0, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_0)
	mockDevice := mocks.NewDummyDevice(common.Vulkan1_0, []string{})

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
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

		next := (*khr_dynamic_rendering_loader.VkPhysicalDeviceDynamicRenderingFeaturesKHR)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000044003), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DYNAMIC_RENDERING_FEATURES_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1), val.FieldByName("dynamicRendering").Uint())

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
				Next: khr_dynamic_rendering.PhysicalDeviceDynamicRenderingFeatures{
					DynamicRendering: true,
				},
			},
		})
	require.NoError(t, err)
	require.Equal(t, mockDevice.Handle(), device.Device().Handle())
}

func TestPhysicalDeviceDynamicRenderingFeatures_OutData(t *testing.T) {
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
		pFeatures *khr_get_physical_device_properties2_loader.VkPhysicalDeviceFeatures2KHR) {

		val := reflect.ValueOf(pFeatures).Elem()
		require.Equal(t, uint64(1000059000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FEATURES_2_KHR

		next := (*khr_dynamic_rendering_loader.VkPhysicalDeviceDynamicRenderingFeaturesKHR)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000044003), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DYNAMIC_RENDERING_FEATURES_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("dynamicRendering").UnsafeAddr())) = loader.VkBool32(1)
	})

	var outData khr_dynamic_rendering.PhysicalDeviceDynamicRenderingFeatures
	err := extension.GetPhysicalDeviceFeatures2(
		physicalDevice,
		&khr_get_physical_device_properties2.PhysicalDeviceFeatures2{
			NextOutData: common.NextOutData{Next: &outData},
		},
	)
	require.NoError(t, err)
	require.Equal(t, khr_dynamic_rendering.PhysicalDeviceDynamicRenderingFeatures{
		DynamicRendering: true,
	}, outData)
}
