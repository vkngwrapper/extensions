package khr_external_memory_test

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
	mock_driver "github.com/vkngwrapper/core/v3/loader/mocks"
	"github.com/vkngwrapper/core/v3/mocks"
	"github.com/vkngwrapper/core/v3/mocks/mocks1_0"
	"github.com/vkngwrapper/extensions/v3/khr_external_memory"
	khr_external_memory_driver "github.com/vkngwrapper/extensions/v3/khr_external_memory/loader"
	"github.com/vkngwrapper/extensions/v3/khr_external_memory_capabilities"
	"go.uber.org/mock/gomock"
)

func TestExternalMemoryAllocateOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreLoader := mock_driver.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := mocks1_0.InternalDeviceDriver(coreLoader)
	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	mockMemory := mocks.NewDummyDeviceMemory(device, 1)

	coreLoader.EXPECT().VkAllocateMemory(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(device loader.VkDevice,
			pAllocateInfo *loader.VkMemoryAllocateInfo,
			pAllocator *loader.VkAllocationCallbacks,
			pMemory *loader.VkDeviceMemory,
		) (common.VkResult, error) {
			*pMemory = mockMemory.Handle()

			val := reflect.ValueOf(pAllocateInfo).Elem()
			require.Equal(t, uint64(5), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_MEMORY_ALLOCATE_INFO
			require.Equal(t, uint64(1), val.FieldByName("allocationSize").Uint())
			require.Equal(t, uint64(3), val.FieldByName("memoryTypeIndex").Uint())

			next := (*khr_external_memory_driver.VkExportMemoryAllocateInfoKHR)(val.FieldByName("pNext").UnsafePointer())
			val = reflect.ValueOf(next).Elem()

			require.Equal(t, uint64(1000072002), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_EXPORT_MEMORY_ALLOCATE_INFO_KHR
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0x10), val.FieldByName("handleTypes").Uint()) // VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D11_TEXTURE_KMT_BIT_KHR

			return core1_0.VKSuccess, nil
		})

	memory, _, err := driver.AllocateMemory(device, nil, core1_0.MemoryAllocateInfo{
		AllocationSize:  1,
		MemoryTypeIndex: 3,
		NextOptions: common.NextOptions{
			khr_external_memory.ExportMemoryAllocateInfo{
				HandleTypes: khr_external_memory_capabilities.ExternalMemoryHandleTypeD3D11TextureKMT,
			},
		},
	})
	require.NoError(t, err)
	require.Equal(t, mockMemory.Handle(), memory.Handle())
}

func TestExternalMemoryImageOptions(t *testing.T) {
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
	).DoAndReturn(
		func(device loader.VkDevice,
			pCreateInfo *loader.VkImageCreateInfo,
			pAllocator *loader.VkAllocationCallbacks,
			pImage *loader.VkImage,
		) (common.VkResult, error) {
			*pImage = mockImage.Handle()

			val := reflect.ValueOf(pCreateInfo).Elem()
			require.Equal(t, uint64(14), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_IMAGE_CREATE_INFO
			require.Equal(t, uint64(1), val.FieldByName("mipLevels").Uint())
			require.Equal(t, uint64(3), val.FieldByName("arrayLayers").Uint())

			next := (*khr_external_memory_driver.VkExternalMemoryImageCreateInfoKHR)(val.FieldByName("pNext").UnsafePointer())
			val = reflect.ValueOf(next).Elem()

			require.Equal(t, uint64(1000072001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_EXTERNAL_MEMORY_IMAGE_CREATE_INFO_KHR
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0x20), val.FieldByName("handleTypes").Uint()) // VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D12_HEAP_BIT_KHR

			return core1_0.VKSuccess, nil
		})

	image, _, err := driver.CreateImage(
		device,
		nil,
		core1_0.ImageCreateInfo{
			MipLevels:   1,
			ArrayLayers: 3,

			NextOptions: common.NextOptions{
				khr_external_memory.ExternalMemoryImageCreateInfo{
					HandleTypes: khr_external_memory_capabilities.ExternalMemoryHandleTypeD3D12Heap,
				},
			},
		})
	require.NoError(t, err)
	require.Equal(t, mockImage.Handle(), image.Handle())
}

func TestExternalMemoryBufferOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreLoader := mock_driver.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := mocks1_0.InternalDeviceDriver(coreLoader)
	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	mockBuffer := mocks.NewDummyBuffer(device)

	coreLoader.EXPECT().VkCreateBuffer(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(device loader.VkDevice,
			pCreateInfo *loader.VkBufferCreateInfo,
			pAllocator *loader.VkAllocationCallbacks,
			pImage *loader.VkBuffer,
		) (common.VkResult, error) {
			*pImage = mockBuffer.Handle()

			val := reflect.ValueOf(pCreateInfo).Elem()
			require.Equal(t, uint64(12), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_BUFFER_CREATE_INFO
			require.Equal(t, uint64(1), val.FieldByName("size").Uint())
			require.Equal(t, uint64(8), val.FieldByName("usage").Uint())

			next := (*khr_external_memory_driver.VkExternalMemoryImageCreateInfoKHR)(val.FieldByName("pNext").UnsafePointer())
			val = reflect.ValueOf(next).Elem()

			require.Equal(t, uint64(1000072000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_EXTERNAL_MEMORY_BUFFER_CREATE_INFO_KHR
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(8), val.FieldByName("handleTypes").Uint()) // VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D11_TEXTURE_BIT_KHR

			return core1_0.VKSuccess, nil
		})

	buffer, _, err := driver.CreateBuffer(
		device,
		nil,
		core1_0.BufferCreateInfo{
			Size:  1,
			Usage: core1_0.BufferUsageStorageTexelBuffer,

			NextOptions: common.NextOptions{
				khr_external_memory.ExternalMemoryBufferCreateInfo{
					HandleTypes: khr_external_memory_capabilities.ExternalMemoryHandleTypeD3D11Texture,
				},
			},
		})
	require.NoError(t, err)
	require.Equal(t, mockBuffer.Handle(), buffer.Handle())
}
