package khr_image_format_list

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
	khr_image_format_list_driver "github.com/vkngwrapper/extensions/v3/khr_image_format_list/loader"
	"go.uber.org/mock/gomock"
)

func TestImageFormatListCreateOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreLoader := mock_driver.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := mocks1_0.InternalDeviceDriver(coreLoader)
	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	mockImage := mocks.NewDummyImage(device)

	coreLoader.EXPECT().VkCreateImage(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device loader.VkDevice,
		pCreateInfo *loader.VkImageCreateInfo,
		pAllocator *loader.VkAllocationCallbacks,
		pImage *loader.VkImage) (common.VkResult, error) {

		*pImage = mockImage.Handle()

		val := reflect.ValueOf(pCreateInfo).Elem()
		require.Equal(t, uint64(14), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_IMAGE_CREATE_INFO

		next := (*khr_image_format_list_driver.VkImageFormatListCreateInfoKHR)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000147000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_IMAGE_FORMAT_LIST_CREATE_INFO_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(3), val.FieldByName("viewFormatCount").Uint())

		formatPtr := (*loader.VkFormat)(val.FieldByName("pViewFormats").UnsafePointer())
		formatSlice := ([]loader.VkFormat)(unsafe.Slice(formatPtr, 3))
		require.Equal(t, []loader.VkFormat{64, 57, 52}, formatSlice)

		return core1_0.VKSuccess, nil
	})

	image, _, err := driver.CreateImage(
		device,
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
