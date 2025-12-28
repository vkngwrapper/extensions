package khr_timeline_semaphore_test

import (
	"reflect"
	"testing"
	"time"
	"unsafe"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
	mock_driver "github.com/vkngwrapper/core/v3/loader/mocks"
	"github.com/vkngwrapper/core/v3/mocks"
	"github.com/vkngwrapper/core/v3/mocks/mocks1_0"
	"github.com/vkngwrapper/extensions/v3/khr_get_physical_device_properties2"
	khr_get_physical_device_properties2_driver "github.com/vkngwrapper/extensions/v3/khr_get_physical_device_properties2/loader"
	mock_get_physical_device_properties2 "github.com/vkngwrapper/extensions/v3/khr_get_physical_device_properties2/mocks"
	"github.com/vkngwrapper/extensions/v3/khr_timeline_semaphore"
	khr_timeline_semaphore_driver "github.com/vkngwrapper/extensions/v3/khr_timeline_semaphore/loader"
	mock_timeline_semaphore "github.com/vkngwrapper/extensions/v3/khr_timeline_semaphore/mocks"
	"go.uber.org/mock/gomock"
)

func TestVulkanExtension_SemaphoreCounterValue(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	semaphore := mocks.NewDummySemaphore(device)

	extDriver := mock_timeline_semaphore.NewMockLoader(ctrl)
	extension := khr_timeline_semaphore.CreateExtensionDriverFromLoader(extDriver, device)

	extDriver.EXPECT().VkGetSemaphoreCounterValueKHR(
		device.Handle(),
		semaphore.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device loader.VkDevice,
		semaphore loader.VkSemaphore,
		pValue *loader.Uint64) (common.VkResult, error) {

		*pValue = loader.Uint64(37)
		return core1_0.VKSuccess, nil
	})

	value, _, err := extension.GetSemaphoreCounterValue(
		semaphore,
	)
	require.NoError(t, err)
	require.Equal(t, uint64(37), value)
}

func TestVulkanExtension_SignalSemaphore(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	semaphore := mocks.NewDummySemaphore(device)

	extDriver := mock_timeline_semaphore.NewMockLoader(ctrl)
	extension := khr_timeline_semaphore.CreateExtensionDriverFromLoader(extDriver, device)

	extDriver.EXPECT().VkSignalSemaphoreKHR(
		device.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device loader.VkDevice,
		pSignalInfo *khr_timeline_semaphore_driver.VkSemaphoreSignalInfoKHR) (common.VkResult, error) {

		val := reflect.ValueOf(pSignalInfo).Elem()
		require.Equal(t, uint64(1000207005), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SEMAPHORE_SIGNAL_INFO_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, semaphore.Handle(), loader.VkSemaphore(val.FieldByName("semaphore").UnsafePointer()))
		require.Equal(t, uint64(13), val.FieldByName("value").Uint())

		return core1_0.VKSuccess, nil
	})

	_, err := extension.SignalSemaphore(
		khr_timeline_semaphore.SemaphoreSignalInfo{
			Semaphore: semaphore,
			Value:     uint64(13),
		})
	require.NoError(t, err)
}

func TestVulkanExtension_WaitSemaphores(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})

	extDriver := mock_timeline_semaphore.NewMockLoader(ctrl)
	extension := khr_timeline_semaphore.CreateExtensionDriverFromLoader(extDriver, device)

	semaphore1 := mocks.NewDummySemaphore(device)
	semaphore2 := mocks.NewDummySemaphore(device)

	extDriver.EXPECT().VkWaitSemaphoresKHR(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		loader.Uint64(60000000000),
	).DoAndReturn(func(device loader.VkDevice,
		pWaitInfo *khr_timeline_semaphore_driver.VkSemaphoreWaitInfoKHR,
		timeout loader.Uint64) (common.VkResult, error) {

		val := reflect.ValueOf(pWaitInfo).Elem()
		require.Equal(t, uint64(1000207004), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SEMAPHORE_WAIT_INFO_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1), val.FieldByName("flags").Uint()) // VK_SEMAPHORE_WAIT_ANY_BIT_KHR
		require.Equal(t, uint64(2), val.FieldByName("semaphoreCount").Uint())

		semaphorePtr := (*loader.VkSemaphore)(val.FieldByName("pSemaphores").UnsafePointer())
		semaphoreSlice := unsafe.Slice(semaphorePtr, 2)
		require.Equal(t, []loader.VkSemaphore{semaphore1.Handle(), semaphore2.Handle()}, semaphoreSlice)

		valuesPtr := (*loader.Uint64)(val.FieldByName("pValues").UnsafePointer())
		valuesSlice := unsafe.Slice(valuesPtr, 2)
		require.Equal(t, []loader.Uint64{13, 19}, valuesSlice)

		return core1_0.VKSuccess, nil
	})

	_, err := extension.WaitSemaphores(
		time.Minute,
		khr_timeline_semaphore.SemaphoreWaitInfo{
			Flags: khr_timeline_semaphore.SemaphoreWaitAny,
			Semaphores: []core.Semaphore{
				semaphore1,
				semaphore2,
			},
			Values: []uint64{
				13,
				19,
			},
		})
	require.NoError(t, err)
}

func TestPhysicalDeviceTimelineSemaphoreFeaturesOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	instance := mocks.NewDummyInstance(common.Vulkan1_0, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_0)
	mockDevice := mocks.NewDummyDevice(common.Vulkan1_0, []string{})

	coreLoader := mock_driver.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := mocks1_0.InternalCoreInstanceDriver(instance, coreLoader)

	coreLoader.EXPECT().VkCreateDevice(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice loader.VkPhysicalDevice,
		pCreateInfo *loader.VkDeviceCreateInfo,
		pAllocator *loader.VkAllocationCallbacks,
		pDevice *loader.VkDevice) (common.VkResult, error) {
		*pDevice = mockDevice.Handle()

		val := reflect.ValueOf(pCreateInfo).Elem()
		require.Equal(t, uint64(3), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_CREATE_INFO

		next := (*khr_timeline_semaphore_driver.VkPhysicalDeviceTimelineSemaphoreFeaturesKHR)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000207000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_TIMELINE_SEMAPHORE_FEATURES_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1), val.FieldByName("timelineSemaphore").Uint())

		return core1_0.VKSuccess, nil
	})

	device, _, err := driver.CreateDevice(
		physicalDevice,
		nil,
		core1_0.DeviceCreateInfo{
			QueueCreateInfos: []core1_0.DeviceQueueCreateInfo{
				{
					QueuePriorities: []float32{0},
				},
			},
			NextOptions: common.NextOptions{
				khr_timeline_semaphore.PhysicalDeviceTimelineSemaphoreFeatures{
					TimelineSemaphore: true,
				},
			},
		})
	require.NoError(t, err)
	require.Equal(t, mockDevice.Handle(), device.Handle())
}

func TestPhysicalDeviceTimelineSemaphoreFeaturesOutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_get_physical_device_properties2.NewMockLoader(ctrl)
	extension := khr_get_physical_device_properties2.CreateExtensionDriverFromLoader(extDriver)

	instance := mocks.NewDummyInstance(common.Vulkan1_0, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_0)

	extDriver.EXPECT().VkGetPhysicalDeviceFeatures2KHR(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice loader.VkPhysicalDevice,
		pFeatures *khr_get_physical_device_properties2_driver.VkPhysicalDeviceFeatures2KHR) {

		val := reflect.ValueOf(pFeatures).Elem()
		require.Equal(t, uint64(1000059000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FEATURES_2_KHR

		next := (*khr_timeline_semaphore_driver.VkPhysicalDeviceTimelineSemaphoreFeaturesKHR)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000207000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_TIMELINE_SEMAPHORE_FEATURES_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("timelineSemaphore").UnsafeAddr())) = loader.VkBool32(1)
	})

	var outData khr_timeline_semaphore.PhysicalDeviceTimelineSemaphoreFeatures
	err := extension.GetPhysicalDeviceFeatures2(
		physicalDevice,
		&khr_get_physical_device_properties2.PhysicalDeviceFeatures2{
			NextOutData: common.NextOutData{&outData},
		},
	)
	require.NoError(t, err)
	require.Equal(t, khr_timeline_semaphore.PhysicalDeviceTimelineSemaphoreFeatures{
		TimelineSemaphore: true,
	}, outData)
}

func TestSemaphoreTypeCreateOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	mockSemaphore := mocks.NewDummySemaphore(device)

	coreLoader := mock_driver.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := mocks1_0.InternalDeviceDriver(device, coreLoader)

	coreLoader.EXPECT().VkCreateSemaphore(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device loader.VkDevice,
		pCreateInfo *loader.VkSemaphoreCreateInfo,
		pAllocator *loader.VkAllocationCallbacks,
		pSemaphore *loader.VkSemaphore) (common.VkResult, error) {

		*pSemaphore = mockSemaphore.Handle()

		val := reflect.ValueOf(pCreateInfo).Elem()
		require.Equal(t, uint64(9), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SEMAPHORE_CREATE_INFO

		next := (*khr_timeline_semaphore_driver.VkSemaphoreTypeCreateInfoKHR)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000207002), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SEMAPHORE_TYPE_CREATE_INFO_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1), val.FieldByName("semaphoreType").Uint()) // VK_SEMAPHORE_TYPE_TIMELINE_KHR
		require.Equal(t, uint64(13), val.FieldByName("initialValue").Uint())

		return core1_0.VKSuccess, nil
	})

	semaphore, _, err := driver.CreateSemaphore(
		nil,
		core1_0.SemaphoreCreateInfo{
			NextOptions: common.NextOptions{khr_timeline_semaphore.SemaphoreTypeCreateInfo{
				SemaphoreType: khr_timeline_semaphore.SemaphoreTypeTimeline,
				InitialValue:  uint64(13),
			}},
		})
	require.NoError(t, err)
	require.Equal(t, mockSemaphore.Handle(), semaphore.Handle())
}

func TestTimelineSemaphoreSubmitOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	queue := mocks.NewDummyQueue(device)
	fence := mocks.NewDummyFence(device)

	coreLoader := mock_driver.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := mocks1_0.InternalDeviceDriver(device, coreLoader)

	coreLoader.EXPECT().VkQueueSubmit(
		queue.Handle(),
		loader.Uint32(1),
		gomock.Not(gomock.Nil()),
		fence.Handle(),
	).DoAndReturn(func(queue loader.VkQueue,
		submitCount loader.Uint32,
		pSubmits *loader.VkSubmitInfo,
		fence loader.VkFence) (common.VkResult, error) {

		val := reflect.ValueOf(pSubmits).Elem()
		require.Equal(t, uint64(4), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SUBMIT_INFO

		next := (*khr_timeline_semaphore_driver.VkTimelineSemaphoreSubmitInfoKHR)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()
		require.Equal(t, uint64(1000207003), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_TIMELINE_SEMAPHORE_SUBMIT_INFO_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(2), val.FieldByName("waitSemaphoreValueCount").Uint())
		require.Equal(t, uint64(3), val.FieldByName("signalSemaphoreValueCount").Uint())

		waitPtr := (*loader.Uint64)(val.FieldByName("pWaitSemaphoreValues").UnsafePointer())
		waitSlice := unsafe.Slice(waitPtr, 2)
		require.Equal(t, []loader.Uint64{3, 5}, waitSlice)

		signalPtr := (*loader.Uint64)(val.FieldByName("pSignalSemaphoreValues").UnsafePointer())
		signalSlice := unsafe.Slice(signalPtr, 3)
		require.Equal(t, []loader.Uint64{7, 11, 13}, signalSlice)

		return core1_0.VKSuccess, nil
	})

	_, err := driver.QueueSubmit(
		queue,
		&fence,
		core1_0.SubmitInfo{
			NextOptions: common.NextOptions{
				khr_timeline_semaphore.TimelineSemaphoreSubmitInfo{
					WaitSemaphoreValues:   []uint64{3, 5},
					SignalSemaphoreValues: []uint64{7, 11, 13},
				},
			},
		},
	)
	require.NoError(t, err)
}

func TestPhysicalDeviceTimelineSemaphoreOutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_get_physical_device_properties2.NewMockLoader(ctrl)
	extension := khr_get_physical_device_properties2.CreateExtensionDriverFromLoader(extDriver)

	instance := mocks.NewDummyInstance(common.Vulkan1_0, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_0)

	extDriver.EXPECT().VkGetPhysicalDeviceProperties2KHR(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice loader.VkPhysicalDevice, pProperties *khr_get_physical_device_properties2_driver.VkPhysicalDeviceProperties2KHR) {
		val := reflect.ValueOf(pProperties).Elem()
		require.Equal(t, uint64(1000059001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROPERTIES_2_KHR

		next := (*khr_timeline_semaphore_driver.VkPhysicalDeviceTimelineSemaphorePropertiesKHR)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000207001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_TIMELINE_SEMAPHORE_PROPERTIES_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		*(*loader.Uint64)(unsafe.Pointer(val.FieldByName("maxTimelineSemaphoreValueDifference").UnsafeAddr())) = loader.Uint64(3)
	})

	var outData khr_timeline_semaphore.PhysicalDeviceTimelineSemaphoreProperties
	err := extension.GetPhysicalDeviceProperties2(
		physicalDevice,
		&khr_get_physical_device_properties2.PhysicalDeviceProperties2{
			NextOutData: common.NextOutData{&outData},
		})
	require.NoError(t, err)
	require.Equal(t, khr_timeline_semaphore.PhysicalDeviceTimelineSemaphoreProperties{
		MaxTimelineSemaphoreValueDifference: 3,
	}, outData)
}
