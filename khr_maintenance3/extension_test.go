package khr_maintenance3_test

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
	"github.com/vkngwrapper/core/v3/mocks"
	"github.com/vkngwrapper/extensions/v3/khr_get_physical_device_properties2"
	khr_get_physical_device_properties2_driver "github.com/vkngwrapper/extensions/v3/khr_get_physical_device_properties2/loader"
	mock_get_physical_device_properties2 "github.com/vkngwrapper/extensions/v3/khr_get_physical_device_properties2/mocks"
	"github.com/vkngwrapper/extensions/v3/khr_maintenance3"
	khr_maintenance3_driver "github.com/vkngwrapper/extensions/v3/khr_maintenance3/loader"
	mock_maintenance3 "github.com/vkngwrapper/extensions/v3/khr_maintenance3/mocks"
	"go.uber.org/mock/gomock"
)

func TestVulkanExtension_DescriptorSetLayoutSupport(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_maintenance3.NewMockLoader(ctrl)
	extension := khr_maintenance3.CreateExtensionFromDriver(extDriver)

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})

	extDriver.EXPECT().VkGetDescriptorSetLayoutSupportKHR(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil())).DoAndReturn(
		func(device loader.VkDevice, pCreateInfo *loader.VkDescriptorSetLayoutCreateInfo, pSupport *khr_maintenance3_driver.VkDescriptorSetLayoutSupportKHR) {
			optionVal := reflect.ValueOf(pCreateInfo).Elem()

			require.Equal(t, uint64(32), optionVal.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_CREATE_INFO
			require.True(t, optionVal.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(1), optionVal.FieldByName("bindingCount").Uint())

			bindingPtr := (*loader.VkDescriptorSetLayoutBinding)(optionVal.FieldByName("pBindings").UnsafePointer())
			binding := reflect.ValueOf(bindingPtr).Elem()
			require.Equal(t, uint64(1), binding.FieldByName("binding").Uint())
			require.Equal(t, uint64(3), binding.FieldByName("descriptorCount").Uint())

			outDataVal := reflect.ValueOf(pSupport).Elem()

			require.Equal(t, uint64(1000168001), outDataVal.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_SUPPORT_KHR
			require.True(t, outDataVal.FieldByName("pNext").IsNil())

			*(*loader.VkBool32)(unsafe.Pointer(outDataVal.FieldByName("supported").UnsafeAddr())) = loader.VkBool32(1)
		})

	outData := &khr_maintenance3.DescriptorSetLayoutSupport{}
	err := extension.DescriptorSetLayoutSupport(device, core1_0.DescriptorSetLayoutCreateInfo{
		Bindings: []core1_0.DescriptorSetLayoutBinding{
			{
				Binding:         1,
				DescriptorCount: 3,
			},
		},
	}, outData)
	require.NoError(t, err)
	require.True(t, outData.Supported)
}

func TestMaintenance3OutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_get_physical_device_properties2.NewMockLoader(ctrl)
	extension := khr_get_physical_device_properties2.CreateExtensionFromDriver(extDriver)

	instance := mocks.NewDummyInstance(common.Vulkan1_0, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_0)

	extDriver.EXPECT().VkGetPhysicalDeviceProperties2KHR(physicalDevice.Handle(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(physicalDevice loader.VkPhysicalDevice, pProperties *khr_get_physical_device_properties2_driver.VkPhysicalDeviceProperties2KHR) {
			val := reflect.ValueOf(pProperties).Elem()

			require.Equal(t, uint64(1000059001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROPERTIES_2_KHR

			props := val.FieldByName("properties")
			*(*loader.Uint32)(unsafe.Pointer(props.FieldByName("vendorID").UnsafeAddr())) = loader.Uint32(3)

			maintPtr := (*khr_maintenance3_driver.VkPhysicalDeviceMaintenance3PropertiesKHR)(val.FieldByName("pNext").UnsafePointer())
			maint := reflect.ValueOf(maintPtr).Elem()

			require.Equal(t, uint64(1000168000), maint.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MAINTENANCE_3_PROPERTIES_KHR
			require.True(t, maint.FieldByName("pNext").IsNil())

			*(*loader.Uint32)(unsafe.Pointer(maint.FieldByName("maxPerSetDescriptors").UnsafeAddr())) = loader.Uint32(5)
			*(*loader.Uint64)(unsafe.Pointer(maint.FieldByName("maxMemoryAllocationSize").UnsafeAddr())) = loader.Uint64(7)
		})

	maintOutData := &khr_maintenance3.PhysicalDeviceMaintenance3Properties{}
	outData := &khr_get_physical_device_properties2.PhysicalDeviceProperties2{
		NextOutData: common.NextOutData{Next: maintOutData},
	}
	err := extension.GetPhysicalDeviceProperties2(physicalDevice, outData)
	require.NoError(t, err)

	require.Equal(t, uint32(3), outData.Properties.VendorID)
	require.Equal(t, 5, maintOutData.MaxPerSetDescriptors)
	require.Equal(t, 7, maintOutData.MaxMemoryAllocationSize)
}
