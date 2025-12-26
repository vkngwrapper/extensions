package khr_external_fence_capabilities_test

import (
	"reflect"
	"testing"
	"unsafe"

	uuid2 "github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/loader"
	"github.com/vkngwrapper/core/v3/mocks"
	"github.com/vkngwrapper/extensions/v3/khr_external_fence_capabilities"
	khr_external_fence_capabilities_driver "github.com/vkngwrapper/extensions/v3/khr_external_fence_capabilities/loader"
	mock_external_fence_capabilities "github.com/vkngwrapper/extensions/v3/khr_external_fence_capabilities/mocks"
	"github.com/vkngwrapper/extensions/v3/khr_get_physical_device_properties2"
	khr_get_physical_device_properties2_driver "github.com/vkngwrapper/extensions/v3/khr_get_physical_device_properties2/loader"
	mock_get_physical_device_properties2 "github.com/vkngwrapper/extensions/v3/khr_get_physical_device_properties2/mocks"
	"go.uber.org/mock/gomock"
)

func TestVulkanExtension_PhysicalDeviceExternalFenceProperties(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_external_fence_capabilities.NewMockLoader(ctrl)
	extension := khr_external_fence_capabilities.CreateExtensionFromDriver(extDriver)

	instance := mocks.NewDummyInstance(common.Vulkan1_0, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_0)

	extDriver.EXPECT().VkGetPhysicalDeviceExternalFencePropertiesKHR(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice loader.VkPhysicalDevice,
		pExternalFenceInfo *khr_external_fence_capabilities_driver.VkPhysicalDeviceExternalFenceInfoKHR,
		pExternalFenceProperties *khr_external_fence_capabilities_driver.VkExternalFencePropertiesKHR,
	) {
		val := reflect.ValueOf(pExternalFenceInfo).Elem()
		require.Equal(t, uint64(1000112000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_FENCE_INFO_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(4), val.FieldByName("handleType").Uint()) // VK_EXTERNAL_FENCE_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT_KHR

		val = reflect.ValueOf(pExternalFenceProperties).Elem()
		*(*uint32)(unsafe.Pointer(val.FieldByName("exportFromImportedHandleTypes").UnsafeAddr())) = uint32(8) // VK_EXTERNAL_FENCE_HANDLE_TYPE_SYNC_FD_BIT_KHR
		*(*uint32)(unsafe.Pointer(val.FieldByName("compatibleHandleTypes").UnsafeAddr())) = uint32(4)         // VK_EXTERNAL_FENCE_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT_KHR
		*(*uint32)(unsafe.Pointer(val.FieldByName("externalFenceFeatures").UnsafeAddr())) = uint32(1)         // VK_EXTERNAL_FENCE_FEATURE_EXPORTABLE_BIT_KHR
	})

	var outData khr_external_fence_capabilities.ExternalFenceProperties
	err := extension.PhysicalDeviceExternalFenceProperties(
		physicalDevice,
		khr_external_fence_capabilities.PhysicalDeviceExternalFenceInfo{
			HandleType: khr_external_fence_capabilities.ExternalFenceHandleTypeOpaqueWin32KMT,
		},
		&outData,
	)
	require.NoError(t, err)
	require.Equal(t, khr_external_fence_capabilities.ExternalFenceProperties{
		ExportFromImportedHandleTypes: khr_external_fence_capabilities.ExternalFenceHandleTypeSyncFD,
		CompatibleHandleTypes:         khr_external_fence_capabilities.ExternalFenceHandleTypeOpaqueWin32KMT,
		ExternalFenceFeatures:         khr_external_fence_capabilities.ExternalFenceFeatureExportable,
	}, outData)
}

func TestPhysicalDeviceIDOutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_get_physical_device_properties2.NewMockLoader(ctrl)
	extension := khr_get_physical_device_properties2.CreateExtensionFromDriver(extDriver)

	instance := mocks.NewDummyInstance(common.Vulkan1_0, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_0)

	deviceUUID, err := uuid2.NewRandom()
	require.NoError(t, err)

	driverUUID, err := uuid2.NewRandom()
	require.NoError(t, err)

	extDriver.EXPECT().VkGetPhysicalDeviceProperties2KHR(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(
			physicalDevice loader.VkPhysicalDevice,
			pProperties *khr_get_physical_device_properties2_driver.VkPhysicalDeviceProperties2KHR,
		) {
			val := reflect.ValueOf(pProperties).Elem()
			require.Equal(t, uint64(1000059001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROPERTIES_2_KHR

			next := (*khr_external_fence_capabilities_driver.VkPhysicalDeviceIDPropertiesKHR)(val.FieldByName("pNext").UnsafePointer())
			val = reflect.ValueOf(next).Elem()
			require.Equal(t, uint64(1000071004), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_ID_PROPERTIES_KHR
			require.True(t, val.FieldByName("pNext").IsNil())

			for i := 0; i < len(deviceUUID); i++ {
				*(*byte)(unsafe.Pointer(val.FieldByName("deviceUUID").Index(i).UnsafeAddr())) = deviceUUID[i]
				*(*byte)(unsafe.Pointer(val.FieldByName("driverUUID").Index(i).UnsafeAddr())) = driverUUID[i]
			}

			*(*byte)(unsafe.Pointer(val.FieldByName("deviceLUID").Index(0).UnsafeAddr())) = byte(0xef)
			*(*byte)(unsafe.Pointer(val.FieldByName("deviceLUID").Index(1).UnsafeAddr())) = byte(0xbe)
			*(*byte)(unsafe.Pointer(val.FieldByName("deviceLUID").Index(2).UnsafeAddr())) = byte(0xad)
			*(*byte)(unsafe.Pointer(val.FieldByName("deviceLUID").Index(3).UnsafeAddr())) = byte(0xde)
			*(*byte)(unsafe.Pointer(val.FieldByName("deviceLUID").Index(4).UnsafeAddr())) = byte(0xef)
			*(*byte)(unsafe.Pointer(val.FieldByName("deviceLUID").Index(5).UnsafeAddr())) = byte(0xbe)
			*(*byte)(unsafe.Pointer(val.FieldByName("deviceLUID").Index(6).UnsafeAddr())) = byte(0xad)
			*(*byte)(unsafe.Pointer(val.FieldByName("deviceLUID").Index(7).UnsafeAddr())) = byte(0xde)

			*(*uint32)(unsafe.Pointer(val.FieldByName("deviceNodeMask").UnsafeAddr())) = uint32(7)
			*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("deviceLUIDValid").UnsafeAddr())) = loader.VkBool32(1)
		})

	var properties khr_get_physical_device_properties2.PhysicalDeviceProperties2
	var outData khr_external_fence_capabilities.PhysicalDeviceIDProperties
	properties.NextOutData = common.NextOutData{&outData}

	err = extension.GetPhysicalDeviceProperties2(
		physicalDevice,
		&properties,
	)
	require.NoError(t, err)
	require.Equal(t, khr_external_fence_capabilities.PhysicalDeviceIDProperties{
		DeviceUUID:      deviceUUID,
		DriverUUID:      driverUUID,
		DeviceLUID:      0xdeadbeefdeadbeef,
		DeviceNodeMask:  7,
		DeviceLUIDValid: true,
	}, outData)
}
