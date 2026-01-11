package khr_format_feature_flags2

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/core1_1"
	"github.com/vkngwrapper/core/v3/loader"
	mock_loader "github.com/vkngwrapper/core/v3/loader/mocks"
	"github.com/vkngwrapper/core/v3/mocks"
	"github.com/vkngwrapper/core/v3/mocks/mocks1_1"
	khr_format_feature_flags2_loader "github.com/vkngwrapper/extensions/v3/khr_format_feature_flags2/loader"
	"go.uber.org/mock/gomock"
)

func TestFormatProperties3(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	instance := mocks.NewDummyInstance(common.Vulkan1_0, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_0)
	mockLoader := mock_loader.NewMockLoader(ctrl)
	instanceDriver := mocks1_1.InternalCoreInstanceDriver(instance, mockLoader)

	var formatProps3 FormatProperties3
	formatProperties := core1_1.FormatProperties2{
		NextOutData: common.NextOutData{
			Next: &formatProps3,
		},
	}

	mockLoader.EXPECT().VkGetPhysicalDeviceFormatProperties2(
		physicalDevice.Handle(),
		loader.VkFormat(52),
		gomock.Not(gomock.Nil()),
	).Do(func(physicalDevice loader.VkPhysicalDevice, format loader.VkFormat, pFormatProperties *loader.VkFormatProperties2) {
		props := reflect.ValueOf(pFormatProperties).Elem()
		props3 := reflect.ValueOf((*khr_format_feature_flags2_loader.VkFormatProperties3KHR)(props.FieldByName("pNext").UnsafePointer())).Elem()

		*(*loader.Uint32)(unsafe.Pointer(props3.FieldByName("linearTilingFeatures").UnsafeAddr())) = loader.Uint32(0x00400000)
		*(*loader.Uint32)(unsafe.Pointer(props3.FieldByName("optimalTilingFeatures").UnsafeAddr())) = loader.Uint32(0x00200000)
		*(*loader.Uint32)(unsafe.Pointer(props3.FieldByName("bufferFeatures").UnsafeAddr())) = loader.Uint32(0x00000400)
	})

	err := instanceDriver.GetPhysicalDeviceFormatProperties2(
		physicalDevice,
		core1_0.FormatA8B8G8R8SignedNormalizedPacked,
		&formatProperties,
	)
	require.NoError(t, err)
	require.Equal(t, FormatProperties3{
		LinearTilingFeatures:  FormatFeature2Disjoint,
		OptimalTilingFeatures: FormatFeature2SampledImageYcbcrConversionChromaReconstructionExplicitForceable,
		BufferFeatures:        FormatFeature2BlitSrc,
	}, formatProps3)
}
