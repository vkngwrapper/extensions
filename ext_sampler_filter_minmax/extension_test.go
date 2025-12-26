package ext_sampler_filter_minmax

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
	ext_sampler_filter_minmax_driver "github.com/vkngwrapper/extensions/v3/ext_sampler_filter_minmax/loader"
	"github.com/vkngwrapper/extensions/v3/khr_get_physical_device_properties2"
	khr_get_physical_device_properties2_driver "github.com/vkngwrapper/extensions/v3/khr_get_physical_device_properties2/loader"
	mock_get_physical_device_properties2 "github.com/vkngwrapper/extensions/v3/khr_get_physical_device_properties2/mocks"
	"go.uber.org/mock/gomock"
)

func TestPhysicalDeviceSamplerFilterMinmaxOutData(t *testing.T) {
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

		next := (*ext_sampler_filter_minmax_driver.VkPhysicalDeviceSamplerFilterMinmaxPropertiesEXT)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000130000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SAMPLER_FILTER_MINMAX_PROPERTIES_EXT
		require.True(t, val.FieldByName("pNext").IsNil())
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("filterMinmaxSingleComponentFormats").UnsafeAddr())) = loader.VkBool32(1)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("filterMinmaxImageComponentMapping").UnsafeAddr())) = loader.VkBool32(1)
	})

	var outData PhysicalDeviceSamplerFilterMinmaxProperties
	err := extension.PhysicalDeviceProperties2(
		physicalDevice,
		&khr_get_physical_device_properties2.PhysicalDeviceProperties2{
			NextOutData: common.NextOutData{&outData},
		})
	require.NoError(t, err)
	require.Equal(t, PhysicalDeviceSamplerFilterMinmaxProperties{
		FilterMinmaxImageComponentMapping:  true,
		FilterMinmaxSingleComponentFormats: true,
	}, outData)
}

func TestSamplerReductionModeCreateOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreLoader := mock_driver.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := mocks1_0.InternalDeviceDriver(coreLoader)
	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	mockSampler := mocks.NewDummySampler(device)

	coreLoader.EXPECT().VkCreateSampler(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device loader.VkDevice,
		pCreateInfo *loader.VkSamplerCreateInfo,
		pAllocator *loader.VkAllocationCallbacks,
		pSampler *loader.VkSampler) (common.VkResult, error) {
		*pSampler = mockSampler.Handle()

		val := reflect.ValueOf(pCreateInfo).Elem()
		require.Equal(t, uint64(31), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SAMPLER_CREATE_INFO

		next := (*ext_sampler_filter_minmax_driver.VkSamplerReductionModeCreateInfoEXT)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000130001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SAMPLER_REDUCTION_MODE_CREATE_INFO_EXT
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(2), val.FieldByName("reductionMode").Uint()) // VK_SAMPLER_REDUCTION_MODE_MAX_EXT

		return core1_0.VKSuccess, nil
	})

	sampler, _, err := driver.CreateSampler(
		device,
		nil,
		core1_0.SamplerCreateInfo{
			NextOptions: common.NextOptions{SamplerReductionModeCreateInfo{
				ReductionMode: SamplerReductionModeMax,
			}},
		})
	require.NoError(t, err)
	require.Equal(t, mockSampler.Handle(), sampler.Handle())
}
