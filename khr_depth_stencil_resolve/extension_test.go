package khr_depth_stencil_resolve

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
	"github.com/vkngwrapper/core/v3/mocks"
	"github.com/vkngwrapper/extensions/v3/khr_create_renderpass2"
	khr_create_renderpass2_driver "github.com/vkngwrapper/extensions/v3/khr_create_renderpass2/loader"
	mock_create_renderpass2 "github.com/vkngwrapper/extensions/v3/khr_create_renderpass2/mocks"
	khr_depth_stencil_resolve_driver "github.com/vkngwrapper/extensions/v3/khr_depth_stencil_resolve/loader"
	"github.com/vkngwrapper/extensions/v3/khr_get_physical_device_properties2"
	khr_get_physical_device_properties2_driver "github.com/vkngwrapper/extensions/v3/khr_get_physical_device_properties2/loader"
	mock_get_physical_device_properties2 "github.com/vkngwrapper/extensions/v3/khr_get_physical_device_properties2/mocks"
	"go.uber.org/mock/gomock"
)

func TestPhysicalDeviceDepthStencilResolveOutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_get_physical_device_properties2.NewMockDriver(ctrl)
	extension := khr_get_physical_device_properties2.CreateExtensionFromDriver(extDriver)

	instance := mocks.NewDummyInstance(common.Vulkan1_0, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_0)

	extDriver.EXPECT().VkGetPhysicalDeviceProperties2KHR(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice loader.VkPhysicalDevice,
		pProperties *khr_get_physical_device_properties2_driver.VkPhysicalDeviceProperties2KHR) {

		val := reflect.ValueOf(pProperties).Elem()
		require.Equal(t, uint64(1000059001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROPERTIES_2_KHR

		next := (*khr_depth_stencil_resolve_driver.VkPhysicalDeviceDepthStencilResolvePropertiesKHR)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000199000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DEPTH_STENCIL_RESOLVE_PROPERTIES_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		depthResolveModePtr := (*khr_depth_stencil_resolve_driver.VkResolveModeFlagsKHR)(unsafe.Pointer(val.FieldByName("supportedDepthResolveModes").UnsafeAddr()))
		*depthResolveModePtr = khr_depth_stencil_resolve_driver.VkResolveModeFlagsKHR(2) // VK_RESOLVE_MODE_AVERAGE_BIT_KHR
		stencilResolveModePtr := (*khr_depth_stencil_resolve_driver.VkResolveModeFlagsKHR)(unsafe.Pointer(val.FieldByName("supportedStencilResolveModes").UnsafeAddr()))
		*stencilResolveModePtr = khr_depth_stencil_resolve_driver.VkResolveModeFlagsKHR(8)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("independentResolveNone").UnsafeAddr())) = loader.VkBool32(0)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("independentResolve").UnsafeAddr())) = loader.VkBool32(1)
	})

	var outData PhysicalDeviceDepthStencilResolveProperties
	err := extension.PhysicalDeviceProperties2(
		physicalDevice,
		&khr_get_physical_device_properties2.PhysicalDeviceProperties2{
			NextOutData: common.NextOutData{&outData},
		})
	require.NoError(t, err)
	require.Equal(t, PhysicalDeviceDepthStencilResolveProperties{
		SupportedDepthResolveModes:   ResolveModeAverage,
		SupportedStencilResolveModes: ResolveModeMax,
		IndependentResolve:           true,
		IndependentResolveNone:       false,
	}, outData)
}

func TestSubpassDescriptionDepthStencilResolveOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	mockRenderPass := mocks.NewDummyRenderPass(device)

	extDriver := mock_create_renderpass2.NewMockDriver(ctrl)
	extension := khr_create_renderpass2.CreateExtensionFromDriver(extDriver)

	extDriver.EXPECT().VkCreateRenderPass2KHR(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device loader.VkDevice,
		pCreateInfo *khr_create_renderpass2_driver.VkRenderPassCreateInfo2KHR,
		pAllocator *loader.VkAllocationCallbacks,
		pRenderPass *loader.VkRenderPass) (common.VkResult, error) {

		*pRenderPass = mockRenderPass.Handle()

		val := reflect.ValueOf(pCreateInfo).Elem()
		require.Equal(t, uint64(1000109004), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_RENDER_PASS_CREATE_INFO_2_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1), val.FieldByName("subpassCount").Uint())

		subpass := val.FieldByName("pSubpasses").Elem()
		require.Equal(t, uint64(1000109002), subpass.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SUBPASS_DESCRIPTION_2_KHR

		next := (*khr_depth_stencil_resolve_driver.VkSubpassDescriptionDepthStencilResolveKHR)(subpass.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000199001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SUBPASS_DESCRIPTION_DEPTH_STENCIL_RESOLVE_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(4), val.FieldByName("depthResolveMode").Uint())   // VK_RESOLVE_MODE_MIN_BIT_KHR
		require.Equal(t, uint64(1), val.FieldByName("stencilResolveMode").Uint()) // VK_RESOLVE_MODE_SAMPLE_ZERO_BIT_KHR

		val = val.FieldByName("pDepthStencilResolveAttachment").Elem()

		require.Equal(t, uint64(1000109001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_ATTACHMENT_REFERENCE_2_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(3), val.FieldByName("attachment").Uint())
		require.Equal(t, uint64(7), val.FieldByName("layout").Uint())     // VK_IMAGE_LAYOUT_TRANSFER_DST_OPTIMAL
		require.Equal(t, uint64(4), val.FieldByName("aspectMask").Uint()) // VK_IMAGE_ASPECT_STENCIL_BIT

		return core1_0.VKSuccess, nil
	})

	renderPass, _, err := extension.CreateRenderPass2(device, nil,
		khr_create_renderpass2.RenderPassCreateInfo2{
			Subpasses: []khr_create_renderpass2.SubpassDescription2{
				{
					NextOptions: common.NextOptions{
						SubpassDescriptionDepthStencilResolve{
							DepthResolveMode:   ResolveModeMin,
							StencilResolveMode: ResolveModeSampleZero,
							DepthStencilResolveAttachment: &khr_create_renderpass2.AttachmentReference2{
								Attachment: 3,
								Layout:     core1_0.ImageLayoutTransferDstOptimal,
								AspectMask: core1_0.ImageAspectStencil,
							},
						},
					},
				},
			},
		})
	require.NoError(t, err)
	require.Equal(t, mockRenderPass.Handle(), renderPass.Handle())

}
