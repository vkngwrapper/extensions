package ext_separate_stencil_usage

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/driver"
	mock_driver "github.com/vkngwrapper/core/v3/driver/mocks"
	"github.com/vkngwrapper/core/v3/mocks/mocks1_0"
	ext_separate_stencil_usage_driver "github.com/vkngwrapper/extensions/v3/ext_separate_stencil_usage/driver"
	"go.uber.org/mock/gomock"
)

func TestImageStencilUsageCreateOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := mocks1_0.NewDummyDevice(coreDriver, common.Vulkan1_0, []string{})
	mockImage := mocks1_0.EasyMockImage(ctrl)

	coreDriver.EXPECT().VkCreateImage(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device driver.VkDevice,
		pCreateInfo *driver.VkImageCreateInfo,
		pAllocator *driver.VkAllocationCallbacks,
		pImage *driver.VkImage) (common.VkResult, error) {

		*pImage = mockImage.Handle()

		val := reflect.ValueOf(pCreateInfo).Elem()
		require.Equal(t, uint64(14), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_IMAGE_CREATE_INFO

		next := (*ext_separate_stencil_usage_driver.VkImageStencilUsageCreateInfoEXT)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000246000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_IMAGE_STENCIL_USAGE_CREATE_INFO_EXT
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(0x10), val.FieldByName("stencilUsage").Uint()) // VK_IMAGE_USAGE_COLOR_ATTACHMENT_BIT

		return core1_0.VKSuccess, nil
	})

	image, _, err := device.CreateImage(
		nil,
		core1_0.ImageCreateInfo{
			NextOptions: common.NextOptions{ImageStencilUsageCreateInfo{
				StencilUsage: core1_0.ImageUsageColorAttachment,
			}},
		})
	require.NoError(t, err)
	require.Equal(t, mockImage.Handle(), image.Handle())
}
