package khr_image_format_list

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v2/mocks"
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/common/extensions"
	"github.com/vkngwrapper/core/v2/core1_0"
	"github.com/vkngwrapper/core/v2/driver"
	mock_driver "github.com/vkngwrapper/core/v2/driver/mocks"
	khr_image_format_list_driver "github.com/vkngwrapper/extensions/v2/khr_image_format_list/driver"
	"reflect"
	"testing"
	"unsafe"
)

func TestImageFormatListCreateOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := extensions.CreateDeviceObject(coreDriver, mocks.NewFakeDeviceHandle(), common.Vulkan1_0)
	mockImage := mocks.EasyMockImage(ctrl)

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

		next := (*khr_image_format_list_driver.VkImageFormatListCreateInfoKHR)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000147000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_IMAGE_FORMAT_LIST_CREATE_INFO_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(3), val.FieldByName("viewFormatCount").Uint())

		formatPtr := (*driver.VkFormat)(val.FieldByName("pViewFormats").UnsafePointer())
		formatSlice := ([]driver.VkFormat)(unsafe.Slice(formatPtr, 3))
		require.Equal(t, []driver.VkFormat{64, 57, 52}, formatSlice)

		return core1_0.VKSuccess, nil
	})

	image, _, err := device.CreateImage(
		nil,
		core1_0.ImageCreateInfo{
			NextOptions: common.NextOptions{
				ImageFormatListCreateInfo{
					ViewFormats: []core1_0.Format{
						core1_0.FormatA2B10G10R10UnsignedNormalizedPacked,
						core1_0.FormatA8B8G8R8SRGBPacked,
						core1_0.FormatA8B8G8R8SignedNormalizedPacked,
					},
				},
			},
		})
	require.NoError(t, err)
	require.Equal(t, mockImage.Handle(), image.Handle())
}
