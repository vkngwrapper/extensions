package khr_bind_memory2_test

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
	"github.com/vkngwrapper/core/v3/mocks"
	"github.com/vkngwrapper/extensions/v3/khr_bind_memory2"
	khr_bind_memory2_driver "github.com/vkngwrapper/extensions/v3/khr_bind_memory2/loader"
	mock_bind_memory2 "github.com/vkngwrapper/extensions/v3/khr_bind_memory2/mocks"
	"go.uber.org/mock/gomock"
)

func TestVulkanExtension_BindBufferMemory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	extDriver := mock_bind_memory2.NewMockLoader(ctrl)
	extension := khr_bind_memory2.CreateExtensionDriverFromLoader(extDriver, device)

	buffer1 := mocks.NewDummyBuffer(device)
	buffer2 := mocks.NewDummyBuffer(device)

	memory1 := mocks.NewDummyDeviceMemory(device, 1)
	memory2 := mocks.NewDummyDeviceMemory(device, 1)

	extDriver.EXPECT().VkBindBufferMemory2KHR(device.Handle(), loader.Uint32(2), gomock.Not(gomock.Nil())).
		DoAndReturn(func(device loader.VkDevice, bindInfoCount loader.Uint32, pBindInfos *khr_bind_memory2_driver.VkBindBufferMemoryInfoKHR) (common.VkResult, error) {
			bindInfoSlice := unsafe.Slice(pBindInfos, 2)
			val := reflect.ValueOf(bindInfoSlice)

			bind := val.Index(0)
			require.Equal(t, uint64(1000157000), bind.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_BIND_BUFFER_MEMORY_INFO_KHR
			require.True(t, bind.FieldByName("pNext").IsNil())
			require.Equal(t, buffer1.Handle(), (loader.VkBuffer)(bind.FieldByName("buffer").UnsafePointer()))
			require.Equal(t, memory1.Handle(), (loader.VkDeviceMemory)(bind.FieldByName("memory").UnsafePointer()))
			require.Equal(t, uint64(1), bind.FieldByName("memoryOffset").Uint())

			bind = val.Index(1)
			require.Equal(t, uint64(1000157000), bind.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_BIND_BUFFER_MEMORY_INFO_KHR
			require.True(t, bind.FieldByName("pNext").IsNil())
			require.Equal(t, buffer2.Handle(), (loader.VkBuffer)(bind.FieldByName("buffer").UnsafePointer()))
			require.Equal(t, memory2.Handle(), (loader.VkDeviceMemory)(bind.FieldByName("memory").UnsafePointer()))
			require.Equal(t, uint64(3), bind.FieldByName("memoryOffset").Uint())

			return core1_0.VKSuccess, nil
		})

	_, err := extension.BindBufferMemory2([]khr_bind_memory2.BindBufferMemoryInfo{
		{
			Buffer:       buffer1,
			Memory:       memory1,
			MemoryOffset: 1,
		},
		{
			Buffer:       buffer2,
			Memory:       memory2,
			MemoryOffset: 3,
		},
	})
	require.NoError(t, err)
}

func TestVulkanExtension_BindImageMemory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_bind_memory2.NewMockLoader(ctrl)
	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	extension := khr_bind_memory2.CreateExtensionDriverFromLoader(extDriver, device)

	image1 := mocks.NewDummyImage(device)
	image2 := mocks.NewDummyImage(device)

	memory1 := mocks.NewDummyDeviceMemory(device, 1)
	memory2 := mocks.NewDummyDeviceMemory(device, 1)

	extDriver.EXPECT().VkBindImageMemory2KHR(device.Handle(), loader.Uint32(2), gomock.Not(gomock.Nil())).
		DoAndReturn(func(device loader.VkDevice, bindInfoCount loader.Uint32, pBindInfos *khr_bind_memory2_driver.VkBindImageMemoryInfoKHR) (common.VkResult, error) {
			bindInfoSlice := unsafe.Slice(pBindInfos, 2)
			val := reflect.ValueOf(bindInfoSlice)

			bind := val.Index(0)
			require.Equal(t, uint64(1000157001), bind.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_BIND_IMAGE_MEMORY_INFO_KHR
			require.True(t, bind.FieldByName("pNext").IsNil())
			require.Equal(t, image1.Handle(), (loader.VkImage)(bind.FieldByName("image").UnsafePointer()))
			require.Equal(t, memory1.Handle(), (loader.VkDeviceMemory)(bind.FieldByName("memory").UnsafePointer()))
			require.Equal(t, uint64(1), bind.FieldByName("memoryOffset").Uint())

			bind = val.Index(1)
			require.Equal(t, uint64(1000157001), bind.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_BIND_IMAGE_MEMORY_INFO_KHR
			require.True(t, bind.FieldByName("pNext").IsNil())
			require.Equal(t, image2.Handle(), (loader.VkImage)(bind.FieldByName("image").UnsafePointer()))
			require.Equal(t, memory2.Handle(), (loader.VkDeviceMemory)(bind.FieldByName("memory").UnsafePointer()))
			require.Equal(t, uint64(3), bind.FieldByName("memoryOffset").Uint())

			return core1_0.VKSuccess, nil
		})

	_, err := extension.BindImageMemory2([]khr_bind_memory2.BindImageMemoryInfo{
		{
			Image:        image1,
			Memory:       memory1,
			MemoryOffset: 1,
		},
		{
			Image:        image2,
			Memory:       memory2,
			MemoryOffset: 3,
		},
	})
	require.NoError(t, err)
}
