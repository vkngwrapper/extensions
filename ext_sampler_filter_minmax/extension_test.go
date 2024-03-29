package ext_sampler_filter_minmax

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v2/mocks"
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/common/extensions"
	"github.com/vkngwrapper/core/v2/core1_0"
	"github.com/vkngwrapper/core/v2/driver"
	mock_driver "github.com/vkngwrapper/core/v2/driver/mocks"
	ext_sampler_filter_minmax_driver "github.com/vkngwrapper/extensions/v2/ext_sampler_filter_minmax/driver"
	"github.com/vkngwrapper/extensions/v2/khr_get_physical_device_properties2"
	khr_get_physical_device_properties2_driver "github.com/vkngwrapper/extensions/v2/khr_get_physical_device_properties2/driver"
	mock_get_physical_device_properties2 "github.com/vkngwrapper/extensions/v2/khr_get_physical_device_properties2/mocks"
	"reflect"
	"testing"
	"unsafe"
)

func TestPhysicalDeviceSamplerFilterMinmaxOutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_get_physical_device_properties2.NewMockDriver(ctrl)
	extension := khr_get_physical_device_properties2.CreateExtensionFromDriver(extDriver)

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	physicalDevice := mocks.EasyMockPhysicalDevice(ctrl, coreDriver)

	extDriver.EXPECT().VkGetPhysicalDeviceProperties2KHR(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice driver.VkPhysicalDevice,
		pProperties *khr_get_physical_device_properties2_driver.VkPhysicalDeviceProperties2KHR) {
		val := reflect.ValueOf(pProperties).Elem()
		require.Equal(t, uint64(1000059001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROPERTIES_2_KHR

		next := (*ext_sampler_filter_minmax_driver.VkPhysicalDeviceSamplerFilterMinmaxPropertiesEXT)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000130000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SAMPLER_FILTER_MINMAX_PROPERTIES_EXT
		require.True(t, val.FieldByName("pNext").IsNil())
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("filterMinmaxSingleComponentFormats").UnsafeAddr())) = driver.VkBool32(1)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("filterMinmaxImageComponentMapping").UnsafeAddr())) = driver.VkBool32(1)
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

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := extensions.CreateDeviceObject(coreDriver, mocks.NewFakeDeviceHandle(), common.Vulkan1_0)
	mockSampler := mocks.EasyMockSampler(ctrl)

	coreDriver.EXPECT().VkCreateSampler(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device driver.VkDevice,
		pCreateInfo *driver.VkSamplerCreateInfo,
		pAllocator *driver.VkAllocationCallbacks,
		pSampler *driver.VkSampler) (common.VkResult, error) {
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

	sampler, _, err := device.CreateSampler(
		nil,
		core1_0.SamplerCreateInfo{
			NextOptions: common.NextOptions{SamplerReductionModeCreateInfo{
				ReductionMode: SamplerReductionModeMax,
			}},
		})
	require.NoError(t, err)
	require.Equal(t, mockSampler.Handle(), sampler.Handle())
}
