package khr_driver_properties

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/loader"
	"github.com/vkngwrapper/core/v3/mocks"
	khr_driver_properties_driver "github.com/vkngwrapper/extensions/v3/khr_driver_properties/loader"
	"github.com/vkngwrapper/extensions/v3/khr_get_physical_device_properties2"
	khr_get_physical_device_properties2_driver "github.com/vkngwrapper/extensions/v3/khr_get_physical_device_properties2/loader"
	mock_get_physical_device_properties2 "github.com/vkngwrapper/extensions/v3/khr_get_physical_device_properties2/mocks"
	"go.uber.org/mock/gomock"
)

func TestPhysicalDeviceDriverOutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_get_physical_device_properties2.NewMockDriver(ctrl)
	extension := khr_get_physical_device_properties2.CreateExtensionFromDriver(extDriver)

	instance := mocks.NewDummyInstance(common.Vulkan1_0, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_0)

	extDriver.EXPECT().VkGetPhysicalDeviceProperties2KHR(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(physicalDevice loader.VkPhysicalDevice,
			pProperties *khr_get_physical_device_properties2_driver.VkPhysicalDeviceProperties2KHR) {

			val := reflect.ValueOf(pProperties).Elem()
			require.Equal(t, uint64(1000059001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROPERTIES_2_KHR

			next := (*khr_driver_properties_driver.VkPhysicalDeviceDriverPropertiesKHR)(val.FieldByName("pNext").UnsafePointer())
			val = reflect.ValueOf(next).Elem()

			require.Equal(t, uint64(1000196000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DRIVER_PROPERTIES_KHR
			require.True(t, val.FieldByName("pNext").IsNil())

			*(*uint32)(unsafe.Pointer(val.FieldByName("driverID").UnsafeAddr())) = uint32(10) // VK_DRIVER_ID_GOOGLE_SWIFTSHADER_KHR
			*(*uint8)(unsafe.Pointer(val.FieldByName("conformanceVersion").FieldByName("major").UnsafeAddr())) = uint8(1)
			*(*uint8)(unsafe.Pointer(val.FieldByName("conformanceVersion").FieldByName("minor").UnsafeAddr())) = uint8(3)
			*(*uint8)(unsafe.Pointer(val.FieldByName("conformanceVersion").FieldByName("subminor").UnsafeAddr())) = uint8(5)
			*(*uint8)(unsafe.Pointer(val.FieldByName("conformanceVersion").FieldByName("patch").UnsafeAddr())) = uint8(7)

			driverNamePtr := (*loader.Char)(unsafe.Pointer(val.FieldByName("driverName").UnsafeAddr()))
			driverNameSlice := ([]loader.Char)(unsafe.Slice(driverNamePtr, 256))
			driverName := "Some Loader"
			for i, r := range []byte(driverName) {
				driverNameSlice[i] = loader.Char(r)
			}
			driverNameSlice[len(driverName)] = 0

			driverInfoPtr := (*loader.Char)(unsafe.Pointer(val.FieldByName("driverInfo").UnsafeAddr()))
			driverInfoSlice := ([]loader.Char)(unsafe.Slice(driverInfoPtr, 256))
			driverInfo := "Whooo Info"
			for i, r := range []byte(driverInfo) {
				driverInfoSlice[i] = loader.Char(r)
			}
			driverInfoSlice[len(driverInfo)] = 0
		})

	var driverOutData PhysicalDeviceDriverProperties
	err := extension.PhysicalDeviceProperties2(
		physicalDevice,
		&khr_get_physical_device_properties2.PhysicalDeviceProperties2{
			NextOutData: common.NextOutData{&driverOutData},
		})
	require.NoError(t, err)
	require.Equal(t, PhysicalDeviceDriverProperties{
		DriverID:           DriverIDGoogleSwiftshader,
		DriverName:         "Some Loader",
		DriverInfo:         "Whooo Info",
		ConformanceVersion: ConformanceVersion{Major: 1, Minor: 3, Subminor: 5, Patch: 7},
	}, driverOutData)
}
