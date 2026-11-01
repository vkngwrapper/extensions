package ext_inline_uniform_block_test

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
	"github.com/vkngwrapper/extensions/v3/ext_inline_uniform_block"
	ext_inline_uniform_block_loader "github.com/vkngwrapper/extensions/v3/ext_inline_uniform_block/loader"
	"go.uber.org/mock/gomock"
)

func TestPhysicalDeviceInlineUniformBlockFeatures_OutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	instance := mocks.NewDummyInstance(common.Vulkan1_1, []string{})
	mockLoader := mock_loader.NewMockLoader(ctrl)
	driver := mocks1_1.InternalCoreInstanceDriver(instance, mockLoader)

	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_1)

	mockLoader.EXPECT().VkGetPhysicalDeviceFeatures2(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice loader.VkPhysicalDevice,
		pFeatures *loader.VkPhysicalDeviceFeatures2) {

		info := reflect.ValueOf(pFeatures).Elem()
		require.Equal(t, uint64(1000059000), info.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FEATURES_2
		require.False(t, info.FieldByName("pNext").IsNil())

		pInnerFeatures := (*ext_inline_uniform_block_loader.VkPhysicalDeviceInlineUniformBlockFeaturesEXT)(info.FieldByName("pNext").UnsafePointer())
		features := reflect.ValueOf(pInnerFeatures).Elem()
		require.Equal(t, uint64(0x3b9ce510), features.FieldByName("sType").Uint())
		require.True(t, features.FieldByName("pNext").IsNil())
		*(*uint64)(unsafe.Pointer(features.FieldByName("inlineUniformBlock").UnsafeAddr())) = uint64(1)
		*(*uint64)(unsafe.Pointer(features.FieldByName("descriptorBindingInlineUniformBlockUpdateAfterBind").UnsafeAddr())) = uint64(1)
	})

	var baseFeatures core1_1.PhysicalDeviceFeatures2
	var features ext_inline_uniform_block.PhysicalDeviceInlineUniformBlockFeatures

	baseFeatures.NextOutData.Next = &features

	err := driver.GetPhysicalDeviceFeatures2(physicalDevice, &baseFeatures)
	require.NoError(t, err)

	require.Equal(t, ext_inline_uniform_block.PhysicalDeviceInlineUniformBlockFeatures{
		InlineUniformBlock: true,
		DescriptorBindingInlineUniformBlockUpdateAfterBind: true,
	}, features)
}

func TestPhysicalDeviceInlineUniformBlockFeatures_Options(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	instance := mocks.NewDummyInstance(common.Vulkan1_1, []string{})
	mockLoader := mock_loader.NewMockLoader(ctrl)
	driver := mocks1_1.InternalCoreInstanceDriver(instance, mockLoader)

	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_1)
	expectedDevice := mocks.NewDummyDevice(common.Vulkan1_1, []string{})

	mockLoader.EXPECT().VkCreateDevice(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice loader.VkPhysicalDevice,
		pCreateInfo *loader.VkDeviceCreateInfo,
		pAllocator *loader.VkAllocationCallbacks,
		pDevice *loader.VkDevice) (common.VkResult, error) {

		info := reflect.ValueOf(pCreateInfo).Elem()
		require.Equal(t, uint64(3), info.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_CREATE_INFO
		require.False(t, info.FieldByName("pNext").IsNil())

		pInnerFeatures := (*ext_inline_uniform_block_loader.VkPhysicalDeviceInlineUniformBlockFeaturesEXT)(info.FieldByName("pNext").UnsafePointer())
		features := reflect.ValueOf(pInnerFeatures).Elem()

		require.Equal(t, uint64(0x3b9ce510), features.FieldByName("sType").Uint())
		require.True(t, features.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1), features.FieldByName("inlineUniformBlock").Uint())
		require.Equal(t, uint64(1), features.FieldByName("descriptorBindingInlineUniformBlockUpdateAfterBind").Uint())

		*pDevice = expectedDevice.Handle()

		return core1_0.VKSuccess, nil
	})

	device, res, err := driver.CreateDevice(physicalDevice, nil, core1_0.DeviceCreateInfo{
		QueueCreateInfos: []core1_0.DeviceQueueCreateInfo{
			{
				QueueFamilyIndex: 0,
				QueuePriorities:  []float32{1},
			},
		},
		NextOptions: common.NextOptions{Next: ext_inline_uniform_block.PhysicalDeviceInlineUniformBlockFeatures{
			InlineUniformBlock: true,
			DescriptorBindingInlineUniformBlockUpdateAfterBind: true,
		}},
	})
	require.NoError(t, err)
	require.Equal(t, core1_0.VKSuccess, res)
	require.NotNil(t, device)
	require.Equal(t, expectedDevice.Handle(), device.Device().Handle())
}

func TestDescriptorPoolInlineUniformBlockCreateInfo_Options(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader := mock_loader.NewMockLoader(ctrl)
	device := mocks.NewDummyDevice(common.Vulkan1_1, []string{})
	driver := mocks1_1.InternalDeviceDriver(device, mockLoader)
	expectedPool := mocks.NewDummyDescriptorPool(device)

	mockLoader.EXPECT().VkCreateDescriptorPool(device.Handle(), gomock.Not(gomock.Nil()), gomock.Nil(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(device loader.VkDevice, pCreateInfo *loader.VkDescriptorPoolCreateInfo, pAllocator *loader.VkAllocationCallbacks, pDescriptorPool *loader.VkDescriptorPool) (common.VkResult, error) {
			info := reflect.ValueOf(pCreateInfo).Elem()
			require.Equal(t, uint64(0x21), info.FieldByName("sType").Uint())
			require.False(t, info.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(3), info.FieldByName("maxSets").Uint())
			require.Equal(t, uint64(1), info.FieldByName("poolSizeCount").Uint())

			pPoolSizes := (*loader.VkDescriptorPoolSize)(info.FieldByName("pPoolSizes").UnsafePointer())
			poolSizesSlice := unsafe.Slice(pPoolSizes, 1)
			poolSize := reflect.ValueOf(poolSizesSlice).Index(0)
			require.Equal(t, uint64(0x3b9ce510), poolSize.FieldByName("_type").Uint())
			require.Equal(t, uint64(5), poolSize.FieldByName("descriptorCount").Uint())

			pNext := (*ext_inline_uniform_block_loader.VkDescriptorPoolInlineUniformBlockCreateInfoEXT)(info.FieldByName("pNext").UnsafePointer())
			info = reflect.ValueOf(pNext).Elem()

			require.Equal(t, uint64(0x3b9ce513), info.FieldByName("sType").Uint())
			require.True(t, info.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(7), info.FieldByName("maxInlineUniformBlockBindings").Uint())

			*pDescriptorPool = expectedPool.Handle()

			return core1_0.VKSuccess, nil
		})

	pool, _, err := driver.CreateDescriptorPool(nil, core1_0.DescriptorPoolCreateInfo{
		Flags:   core1_0.DescriptorPoolCreateFreeDescriptorSet,
		MaxSets: 3,
		PoolSizes: []core1_0.DescriptorPoolSize{
			{
				Type:            ext_inline_uniform_block.DescriptorTypeInlineUniformBlock,
				DescriptorCount: 5,
			},
		},
		Next: ext_inline_uniform_block.DescriptorPoolInlineUniformBlockCreateInfo{
			MaxInlineUniformBlockBindings: 7,
		},
	})
	require.Equal(t, expectedPool.Handle(), pool.Handle())
	require.NoError(t, err)
}

func TestWriteDescriptorSetInlineUniformBlock_Options(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLoader := mock_loader.NewMockLoader(ctrl)
	device := mocks.NewDummyDevice(common.Vulkan1_1, []string{})
	driver := mocks1_1.InternalDeviceDriver(device, mockLoader)
	descriptorPool := mocks.NewDummyDescriptorPool(device)
	descriptorSet := mocks.NewDummyDescriptorSet(descriptorPool, device)

	mockLoader.EXPECT().VkUpdateDescriptorSets(device.Handle(), loader.Uint32(1), gomock.Not(gomock.Nil()), loader.Uint32(0), gomock.Nil()).
		DoAndReturn(func(device loader.VkDevice, descriptorWriteCount loader.Uint32, pDescriptorWrites *loader.VkWriteDescriptorSet, descriptorCopyCount loader.Uint32, pDescriptorCopies *loader.VkCopyDescriptorSet) {
			writeSlice := unsafe.Slice(pDescriptorWrites, 1)
			info := reflect.ValueOf(writeSlice).Index(0)
			require.Equal(t, uint64(0x23), info.FieldByName("sType").Uint())
			require.False(t, info.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(8), info.FieldByName("descriptorCount").Uint())

			pNext := (*ext_inline_uniform_block_loader.VkWriteDescriptorSetInlineUniformBlockEXT)(info.FieldByName("pNext").UnsafePointer())
			info = reflect.ValueOf(pNext).Elem()

			require.Equal(t, uint64(0x3b9ce512), info.FieldByName("sType").Uint())
			require.True(t, info.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(8), info.FieldByName("dataSize").Uint())

			dataPtr := (*byte)(info.FieldByName("pData").UnsafePointer())
			dataSlice := unsafe.Slice(dataPtr, 8)
			text := string(dataSlice)
			require.Equal(t, "testData", text)
		})

	err := driver.UpdateDescriptorSets([]core1_0.WriteDescriptorSet{
		{
			DstSet:         descriptorSet,
			DescriptorType: ext_inline_uniform_block.DescriptorTypeInlineUniformBlock,
			Next: ext_inline_uniform_block.WriteDescriptorSetInlineUniformBlock{
				Data: []byte("testData"),
			},
		},
	}, nil)
	require.NoError(t, err)
}

func TestPhysicalDeviceInlineUniformBlockProperties_OutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	instance := mocks.NewDummyInstance(common.Vulkan1_1, []string{})
	mockLoader := mock_loader.NewMockLoader(ctrl)
	driver := mocks1_1.InternalCoreInstanceDriver(instance, mockLoader)

	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_1)

	mockLoader.EXPECT().VkGetPhysicalDeviceProperties2(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice loader.VkPhysicalDevice, pBaseProperties *loader.VkPhysicalDeviceProperties2) {
		baseProperties := reflect.ValueOf(pBaseProperties).Elem()

		require.Equal(t, uint64(1000059001), baseProperties.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROPERTIES_2
		require.False(t, baseProperties.FieldByName("pNext").IsNil())

		pProperties := (*ext_inline_uniform_block_loader.VkPhysicalDeviceInlineUniformBlockPropertiesEXT)(baseProperties.FieldByName("pNext").UnsafePointer())
		properties := reflect.ValueOf(pProperties).Elem()

		require.Equal(t, uint64(0x3b9ce510), properties.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_ACCELERATION_STRUCTURE_PROPERTIES_KHR
		require.True(t, properties.FieldByName("pNext").IsNil())

		*(*uint64)(unsafe.Pointer(properties.FieldByName("maxInlineUniformBlockSize").UnsafeAddr())) = uint64(1)
		*(*uint64)(unsafe.Pointer(properties.FieldByName("maxPerStageDescriptorInlineUniformBlocks").UnsafeAddr())) = uint64(3)
		*(*uint64)(unsafe.Pointer(properties.FieldByName("maxPerStageDescriptorUpdateAfterBindInlineUniformBlocks").UnsafeAddr())) = uint64(5)
		*(*uint64)(unsafe.Pointer(properties.FieldByName("maxDescriptorSetInlineUniformBlocks").UnsafeAddr())) = uint64(7)
		*(*uint64)(unsafe.Pointer(properties.FieldByName("maxDescriptorSetUpdateAfterBindInlineUniformBlocks").UnsafeAddr())) = uint64(11)
	})

	var baseProperties core1_1.PhysicalDeviceProperties2
	var properties ext_inline_uniform_block.PhysicalDeviceInlineUniformBlockProperties

	baseProperties.Next = &properties

	err := driver.GetPhysicalDeviceProperties2(physicalDevice, &baseProperties)
	require.NoError(t, err)

	require.Equal(t, ext_inline_uniform_block.PhysicalDeviceInlineUniformBlockProperties{
		MaxInlineUniformBlockSize:                               1,
		MaxPerStageDescriptorInlineUniformBlocks:                3,
		MaxPerStageDescriptorUpdateAfterBindInlineUniformBlocks: 5,
		MaxDescriptorSetInlineUniformBlocks:                     7,
		MaxDescriptorSetUpdateAfterBindInlineUniformBlocks:      11,
	}, properties)
}
