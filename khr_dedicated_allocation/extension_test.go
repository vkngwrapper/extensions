package khr_dedicated_allocation_test

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/common/extensions"
	"github.com/vkngwrapper/core/v2/core1_0"
	"github.com/vkngwrapper/core/v2/driver"
	mock_driver "github.com/vkngwrapper/core/v2/driver/mocks"
	"github.com/vkngwrapper/core/v2/mocks"
	"github.com/vkngwrapper/extensions/v2/khr_dedicated_allocation"
	khr_dedicated_allocation_driver "github.com/vkngwrapper/extensions/v2/khr_dedicated_allocation/driver"
	"github.com/vkngwrapper/extensions/v2/khr_get_memory_requirements2"
	khr_get_memory_requirements2_driver "github.com/vkngwrapper/extensions/v2/khr_get_memory_requirements2/driver"
	mock_get_memory_requirements2 "github.com/vkngwrapper/extensions/v2/khr_get_memory_requirements2/mocks"
	"reflect"
	"testing"
	"unsafe"
)

func TestDedicatedMemoryRequirementsOutData_Buffer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_get_memory_requirements2.NewMockDriver(ctrl)
	extension := khr_get_memory_requirements2.CreateExtensionFromDriver(extDriver)

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := mocks.EasyMockDevice(ctrl, coreDriver)
	buffer := mocks.EasyMockBuffer(ctrl)

	extDriver.EXPECT().VkGetBufferMemoryRequirements2KHR(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil())).DoAndReturn(
		func(device driver.VkDevice,
			pInfo *khr_get_memory_requirements2_driver.VkBufferMemoryRequirementsInfo2KHR,
			pMemoryRequirements *khr_get_memory_requirements2_driver.VkMemoryRequirements2KHR,
		) {
			options := reflect.ValueOf(pInfo).Elem()
			require.Equal(t, uint64(1000146000), options.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_BUFFER_MEMORY_REQUIREMENTS_INFO_2_KHR
			require.True(t, options.FieldByName("pNext").IsNil())
			require.Equal(t, buffer.Handle(), driver.VkBuffer(options.FieldByName("buffer").UnsafePointer()))

			outData := reflect.ValueOf(pMemoryRequirements).Elem()
			require.Equal(t, uint64(1000146003), outData.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_MEMORY_REQUIREMENTS_2_KHR

			memoryRequirements := outData.FieldByName("memoryRequirements")
			*(*driver.VkDeviceSize)(unsafe.Pointer(memoryRequirements.FieldByName("size").UnsafeAddr())) = driver.VkDeviceSize(1)
			*(*driver.VkDeviceSize)(unsafe.Pointer(memoryRequirements.FieldByName("alignment").UnsafeAddr())) = driver.VkDeviceSize(3)
			*(*driver.Uint32)(unsafe.Pointer(memoryRequirements.FieldByName("memoryTypeBits").UnsafeAddr())) = driver.Uint32(5)

			dedicatedPtr := (*khr_dedicated_allocation_driver.VkMemoryDedicatedRequirementsKHR)(outData.FieldByName("pNext").UnsafePointer())
			dedicated := reflect.ValueOf(dedicatedPtr).Elem()
			require.Equal(t, uint64(1000127000), dedicated.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_MEMORY_DEDICATED_REQUIREMENTS_KHR
			require.True(t, dedicated.FieldByName("pNext").IsNil())
			*(*driver.VkBool32)(unsafe.Pointer(dedicated.FieldByName("prefersDedicatedAllocation").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(dedicated.FieldByName("requiresDedicatedAllocation").UnsafeAddr())) = driver.VkBool32(0)
		})

	var memReqs khr_dedicated_allocation.MemoryDedicatedRequirements
	var outData = khr_get_memory_requirements2.MemoryRequirements2{
		NextOutData: common.NextOutData{Next: &memReqs},
	}
	err := extension.BufferMemoryRequirements2(device,
		khr_get_memory_requirements2.BufferMemoryRequirementsInfo2{
			Buffer: buffer,
		}, &outData)
	require.NoError(t, err)
	require.False(t, memReqs.RequiresDedicatedAllocation)
	require.True(t, memReqs.PrefersDedicatedAllocation)

	require.Equal(t, 1, outData.MemoryRequirements.Size)
	require.Equal(t, 3, outData.MemoryRequirements.Alignment)
	require.Equal(t, uint32(5), outData.MemoryRequirements.MemoryTypeBits)
}

func TestDedicatedMemoryRequirementsOutData_Image(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_get_memory_requirements2.NewMockDriver(ctrl)
	extension := khr_get_memory_requirements2.CreateExtensionFromDriver(extDriver)

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := mocks.EasyMockDevice(ctrl, coreDriver)
	image := mocks.EasyMockImage(ctrl)

	extDriver.EXPECT().VkGetImageMemoryRequirements2KHR(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil())).DoAndReturn(
		func(device driver.VkDevice,
			pInfo *khr_get_memory_requirements2_driver.VkImageMemoryRequirementsInfo2KHR,
			pMemoryRequirements *khr_get_memory_requirements2_driver.VkMemoryRequirements2KHR,
		) {
			options := reflect.ValueOf(pInfo).Elem()
			require.Equal(t, uint64(1000146001), options.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_IMAGE_MEMORY_REQUIREMENTS_INFO_2_KHR
			require.True(t, options.FieldByName("pNext").IsNil())
			require.Equal(t, image.Handle(), driver.VkImage(options.FieldByName("image").UnsafePointer()))

			outData := reflect.ValueOf(pMemoryRequirements).Elem()
			require.Equal(t, uint64(1000146003), outData.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_MEMORY_REQUIREMENTS_2_KHR

			memoryRequirements := outData.FieldByName("memoryRequirements")
			*(*driver.VkDeviceSize)(unsafe.Pointer(memoryRequirements.FieldByName("size").UnsafeAddr())) = driver.VkDeviceSize(1)
			*(*driver.VkDeviceSize)(unsafe.Pointer(memoryRequirements.FieldByName("alignment").UnsafeAddr())) = driver.VkDeviceSize(3)
			*(*driver.Uint32)(unsafe.Pointer(memoryRequirements.FieldByName("memoryTypeBits").UnsafeAddr())) = driver.Uint32(5)

			dedicatedPtr := (*khr_dedicated_allocation_driver.VkMemoryDedicatedRequirementsKHR)(outData.FieldByName("pNext").UnsafePointer())
			dedicated := reflect.ValueOf(dedicatedPtr).Elem()
			require.Equal(t, uint64(1000127000), dedicated.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_MEMORY_DEDICATED_REQUIREMENTS_KHR
			require.True(t, dedicated.FieldByName("pNext").IsNil())
			*(*driver.VkBool32)(unsafe.Pointer(dedicated.FieldByName("prefersDedicatedAllocation").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(dedicated.FieldByName("requiresDedicatedAllocation").UnsafeAddr())) = driver.VkBool32(0)
		})

	var memReqs khr_dedicated_allocation.MemoryDedicatedRequirements
	var outData = khr_get_memory_requirements2.MemoryRequirements2{
		NextOutData: common.NextOutData{Next: &memReqs},
	}
	err := extension.ImageMemoryRequirements2(device,
		khr_get_memory_requirements2.ImageMemoryRequirementsInfo2{
			Image: image,
		}, &outData)
	require.NoError(t, err)
	require.False(t, memReqs.RequiresDedicatedAllocation)
	require.True(t, memReqs.PrefersDedicatedAllocation)

	require.Equal(t, 1, outData.MemoryRequirements.Size)
	require.Equal(t, 3, outData.MemoryRequirements.Alignment)
	require.Equal(t, uint32(5), outData.MemoryRequirements.MemoryTypeBits)
}

func TestMemoryDedicatedAllocateOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	coreDriver.EXPECT().CreateDeviceDriver(gomock.Any()).Return(coreDriver, nil).AnyTimes()
	device := extensions.CreateDeviceObject(coreDriver, mocks.NewFakeDeviceHandle(), common.Vulkan1_0)

	buffer := mocks.EasyMockBuffer(ctrl)
	expectedMemory := mocks.EasyMockDeviceMemory(ctrl)

	coreDriver.EXPECT().VkAllocateMemory(device.Handle(), gomock.Not(gomock.Nil()), gomock.Nil(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(device driver.VkDevice, pAllocateInfo *driver.VkMemoryAllocateInfo, pAllocator *driver.VkAllocationCallbacks, pMemory *driver.VkDeviceMemory) (common.VkResult, error) {
			*pMemory = expectedMemory.Handle()

			options := reflect.ValueOf(pAllocateInfo).Elem()
			require.Equal(t, uint64(5), options.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_MEMORY_ALLOCATE_INFO
			require.Equal(t, uint64(1), options.FieldByName("allocationSize").Uint())
			require.Equal(t, uint64(3), options.FieldByName("memoryTypeIndex").Uint())

			dedicatedPtr := (*khr_dedicated_allocation_driver.VkMemoryDedicatedAllocateInfoKHR)(options.FieldByName("pNext").UnsafePointer())
			dedicated := reflect.ValueOf(dedicatedPtr).Elem()

			require.Equal(t, uint64(1000127001), dedicated.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_MEMORY_DEDICATED_ALLOCATE_INFO_KHR
			require.True(t, dedicated.FieldByName("pNext").IsNil())
			require.Equal(t, buffer.Handle(), driver.VkBuffer(dedicated.FieldByName("buffer").UnsafePointer()))
			require.True(t, dedicated.FieldByName("image").IsNil())

			return core1_0.VKSuccess, nil
		})

	memory, _, err := device.AllocateMemory(nil, core1_0.MemoryAllocateInfo{
		AllocationSize:  1,
		MemoryTypeIndex: 3,
		NextOptions: common.NextOptions{Next: khr_dedicated_allocation.MemoryDedicatedAllocateInfo{
			Buffer: buffer,
		}},
	})
	require.NoError(t, err)
	require.Equal(t, expectedMemory.Handle(), memory.Handle())

}
