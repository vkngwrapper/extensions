package ext_descriptor_indexing

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
	mock_loader "github.com/vkngwrapper/core/v3/loader/mocks"
	"github.com/vkngwrapper/core/v3/mocks"
	"github.com/vkngwrapper/core/v3/mocks/mocks1_0"
	ext_descriptor_indexing_driver "github.com/vkngwrapper/extensions/v3/ext_descriptor_indexing/loader"
	"github.com/vkngwrapper/extensions/v3/khr_get_physical_device_properties2"
	khr_get_physical_device_properties2_driver "github.com/vkngwrapper/extensions/v3/khr_get_physical_device_properties2/loader"
	mock_get_physical_device_properties2 "github.com/vkngwrapper/extensions/v3/khr_get_physical_device_properties2/mocks"
	"github.com/vkngwrapper/extensions/v3/khr_maintenance3"
	khr_maintenance3_driver "github.com/vkngwrapper/extensions/v3/khr_maintenance3/loader"
	mock_maintenance3 "github.com/vkngwrapper/extensions/v3/khr_maintenance3/mocks"
	"go.uber.org/mock/gomock"
)

func TestDescriptorSetLayoutBindingFlagsCreateOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	mockDescriptorSetLayout := mocks.NewDummyDescriptorSetLayout(device)

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := mocks1_0.InternalDeviceDriver(device, coreLoader)

	coreLoader.EXPECT().VkCreateDescriptorSetLayout(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device loader.VkDevice, pCreateInfo *loader.VkDescriptorSetLayoutCreateInfo, pAllocator *loader.VkAllocationCallbacks, pSetLayout *loader.VkDescriptorSetLayout) (common.VkResult, error) {
		*pSetLayout = mockDescriptorSetLayout.Handle()
		val := reflect.ValueOf(pCreateInfo).Elem()

		require.Equal(t, uint64(32), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_CREATE_INFO

		next := (*ext_descriptor_indexing_driver.VkDescriptorSetLayoutBindingFlagsCreateInfoEXT)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000161000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_BINDING_FLAGS_CREATE_INFO_EXT
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(2), val.FieldByName("bindingCount").Uint())
		flagsPtr := (*ext_descriptor_indexing_driver.VkDescriptorBindingFlagsEXT)(val.FieldByName("pBindingFlags").UnsafePointer())
		flagSlice := unsafe.Slice(flagsPtr, 2)

		require.Equal(t, []ext_descriptor_indexing_driver.VkDescriptorBindingFlagsEXT{8, 1}, flagSlice)

		return core1_0.VKSuccess, nil
	})

	descriptorSetLayout, _, err := driver.CreateDescriptorSetLayout(
		nil,
		core1_0.DescriptorSetLayoutCreateInfo{
			NextOptions: common.NextOptions{
				DescriptorSetLayoutBindingFlagsCreateInfo{
					BindingFlags: []DescriptorBindingFlags{
						DescriptorBindingVariableDescriptorCount,
						DescriptorBindingUpdateAfterBind,
					},
				},
			},
		})
	require.NoError(t, err)
	require.Equal(t, mockDescriptorSetLayout.Handle(), descriptorSetLayout.Handle())
}

func TestDescriptorSetVariableDescriptorCountAllocateOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	descriptorPool := mocks.NewDummyDescriptorPool(device)
	descriptorLayout1 := mocks.NewDummyDescriptorSetLayout(device)
	descriptorLayout2 := mocks.NewDummyDescriptorSetLayout(device)
	descriptorLayout3 := mocks.NewDummyDescriptorSetLayout(device)
	descriptorLayout4 := mocks.NewDummyDescriptorSetLayout(device)

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := mocks1_0.InternalDeviceDriver(device, coreLoader)

	mockDescriptorSet1 := mocks.NewDummyDescriptorSet(descriptorPool, device)
	mockDescriptorSet2 := mocks.NewDummyDescriptorSet(descriptorPool, device)
	mockDescriptorSet3 := mocks.NewDummyDescriptorSet(descriptorPool, device)
	mockDescriptorSet4 := mocks.NewDummyDescriptorSet(descriptorPool, device)

	coreLoader.EXPECT().VkAllocateDescriptorSets(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device loader.VkDevice,
		pAllocateInfo *loader.VkDescriptorSetAllocateInfo,
		pDescriptorSets *loader.VkDescriptorSet) (common.VkResult, error) {

		sets := unsafe.Slice(pDescriptorSets, 4)
		sets[0] = mockDescriptorSet1.Handle()
		sets[1] = mockDescriptorSet2.Handle()
		sets[2] = mockDescriptorSet3.Handle()
		sets[3] = mockDescriptorSet4.Handle()

		val := reflect.ValueOf(pAllocateInfo).Elem()
		require.Equal(t, uint64(34), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DESCRIPTOR_SET_ALLOCATE_INFO

		next := (*ext_descriptor_indexing_driver.VkDescriptorSetVariableDescriptorCountAllocateInfoEXT)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000161003), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DESCRIPTOR_SET_VARIABLE_DESCRIPTOR_COUNT_ALLOCATE_INFO_EXT
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(4), val.FieldByName("descriptorSetCount").Uint())

		countsPtr := (*loader.Uint32)(val.FieldByName("pDescriptorCounts").UnsafePointer())
		countSlice := unsafe.Slice(countsPtr, 4)

		require.Equal(t, []loader.Uint32{1, 3, 5, 7}, countSlice)

		return core1_0.VKSuccess, nil
	})

	sets, _, err := driver.AllocateDescriptorSets(core1_0.DescriptorSetAllocateInfo{
		DescriptorPool: descriptorPool,
		SetLayouts: []core1_0.DescriptorSetLayout{
			descriptorLayout1,
			descriptorLayout2,
			descriptorLayout3,
			descriptorLayout4,
		},
		NextOptions: common.NextOptions{
			DescriptorSetVariableDescriptorCountAllocateInfo{
				DescriptorCounts: []int{1, 3, 5, 7},
			},
		},
	})
	require.NoError(t, err)
	require.Len(t, sets, 4)
	require.Equal(t, []loader.VkDescriptorSet{
		mockDescriptorSet1.Handle(),
		mockDescriptorSet2.Handle(),
		mockDescriptorSet3.Handle(),
		mockDescriptorSet4.Handle(),
	}, []loader.VkDescriptorSet{
		sets[0].Handle(),
		sets[1].Handle(),
		sets[2].Handle(),
		sets[3].Handle(),
	})
}

func TestDescriptorSetVariableDescriptorCountLayoutSupportOutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})

	extDriver := mock_maintenance3.NewMockLoader(ctrl)
	extension := khr_maintenance3.CreateExtensionDriverFromLoader(extDriver, device)

	extDriver.EXPECT().VkGetDescriptorSetLayoutSupportKHR(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device loader.VkDevice,
		pCreateInfo *loader.VkDescriptorSetLayoutCreateInfo,
		pSupport *khr_maintenance3_driver.VkDescriptorSetLayoutSupportKHR) {
		val := reflect.ValueOf(pSupport).Elem()

		require.Equal(t, uint64(1000168001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_SUPPORT
		next := (*ext_descriptor_indexing_driver.VkDescriptorSetVariableDescriptorCountLayoutSupportEXT)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000161004), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DESCRIPTOR_SET_VARIABLE_DESCRIPTOR_COUNT_LAYOUT_SUPPORT_EXT
		require.True(t, val.FieldByName("pNext").IsNil())
		*(*loader.Uint32)(unsafe.Pointer(val.FieldByName("maxVariableDescriptorCount").UnsafeAddr())) = loader.Uint32(7)
	})

	var outData DescriptorSetVariableDescriptorCountLayoutSupport
	err := extension.DescriptorSetLayoutSupport(
		core1_0.DescriptorSetLayoutCreateInfo{},
		&khr_maintenance3.DescriptorSetLayoutSupport{
			NextOutData: common.NextOutData{&outData},
		})
	require.NoError(t, err)
	require.Equal(t, DescriptorSetVariableDescriptorCountLayoutSupport{
		MaxVariableDescriptorCount: 7,
	}, outData)
}

func TestPhysicalDeviceDescriptorIndexingFeaturesOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	instance := mocks.NewDummyInstance(common.Vulkan1_0, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_0)
	mockDevice := mocks.NewDummyDevice(common.Vulkan1_0, []string{})

	coreLoader := mock_loader.LoaderForVersion(ctrl, common.Vulkan1_0)
	driver := mocks1_0.InternalCoreInstanceDriver(instance, coreLoader)

	coreLoader.EXPECT().VkCreateDevice(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(physicalDevice loader.VkPhysicalDevice,
			pCreateInfo *loader.VkDeviceCreateInfo,
			pAllocator *loader.VkAllocationCallbacks,
			pDevice *loader.VkDevice) (common.VkResult, error) {

			*pDevice = mockDevice.Handle()

			val := reflect.ValueOf(pCreateInfo).Elem()
			require.Equal(t, uint64(3), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_CREATE_INFO

			next := (*ext_descriptor_indexing_driver.VkPhysicalDeviceDescriptorIndexingFeaturesEXT)(val.FieldByName("pNext").UnsafePointer())
			val = reflect.ValueOf(next).Elem()

			require.Equal(t, uint64(1000161001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DESCRIPTOR_INDEXING_FEATURES_EXT
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(1), val.FieldByName("shaderInputAttachmentArrayDynamicIndexing").Uint())
			require.Equal(t, uint64(0), val.FieldByName("shaderUniformTexelBufferArrayDynamicIndexing").Uint())
			require.Equal(t, uint64(1), val.FieldByName("shaderStorageTexelBufferArrayDynamicIndexing").Uint())
			require.Equal(t, uint64(0), val.FieldByName("shaderUniformBufferArrayNonUniformIndexing").Uint())
			require.Equal(t, uint64(1), val.FieldByName("shaderSampledImageArrayNonUniformIndexing").Uint())
			require.Equal(t, uint64(0), val.FieldByName("shaderStorageBufferArrayNonUniformIndexing").Uint())
			require.Equal(t, uint64(1), val.FieldByName("shaderStorageImageArrayNonUniformIndexing").Uint())
			require.Equal(t, uint64(0), val.FieldByName("shaderInputAttachmentArrayNonUniformIndexing").Uint())
			require.Equal(t, uint64(1), val.FieldByName("shaderUniformTexelBufferArrayNonUniformIndexing").Uint())
			require.Equal(t, uint64(0), val.FieldByName("shaderStorageTexelBufferArrayNonUniformIndexing").Uint())
			require.Equal(t, uint64(1), val.FieldByName("descriptorBindingUniformBufferUpdateAfterBind").Uint())
			require.Equal(t, uint64(0), val.FieldByName("descriptorBindingSampledImageUpdateAfterBind").Uint())
			require.Equal(t, uint64(1), val.FieldByName("descriptorBindingStorageImageUpdateAfterBind").Uint())
			require.Equal(t, uint64(0), val.FieldByName("descriptorBindingStorageBufferUpdateAfterBind").Uint())
			require.Equal(t, uint64(1), val.FieldByName("descriptorBindingUniformTexelBufferUpdateAfterBind").Uint())
			require.Equal(t, uint64(0), val.FieldByName("descriptorBindingStorageTexelBufferUpdateAfterBind").Uint())
			require.Equal(t, uint64(1), val.FieldByName("descriptorBindingUpdateUnusedWhilePending").Uint())
			require.Equal(t, uint64(0), val.FieldByName("descriptorBindingPartiallyBound").Uint())
			require.Equal(t, uint64(1), val.FieldByName("descriptorBindingVariableDescriptorCount").Uint())
			require.Equal(t, uint64(0), val.FieldByName("runtimeDescriptorArray").Uint())

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
			NextOptions: common.NextOptions{PhysicalDeviceDescriptorIndexingFeatures{
				ShaderInputAttachmentArrayDynamicIndexing:          true,
				ShaderUniformTexelBufferArrayDynamicIndexing:       false,
				ShaderStorageTexelBufferArrayDynamicIndexing:       true,
				ShaderUniformBufferArrayNonUniformIndexing:         false,
				ShaderSampledImageArrayNonUniformIndexing:          true,
				ShaderStorageBufferArrayNonUniformIndexing:         false,
				ShaderStorageImageArrayNonUniformIndexing:          true,
				ShaderInputAttachmentArrayNonUniformIndexing:       false,
				ShaderUniformTexelBufferArrayNonUniformIndexing:    true,
				ShaderStorageTexelBufferArrayNonUniformIndexing:    false,
				DescriptorBindingUniformBufferUpdateAfterBind:      true,
				DescriptorBindingSampledImageUpdateAfterBind:       false,
				DescriptorBindingStorageImageUpdateAfterBind:       true,
				DescriptorBindingStorageBufferUpdateAfterBind:      false,
				DescriptorBindingUniformTexelBufferUpdateAfterBind: true,
				DescriptorBindingStorageTexelBufferUpdateAfterBind: false,
				DescriptorBindingUpdateUnusedWhilePending:          true,
				DescriptorBindingPartiallyBound:                    false,
				DescriptorBindingVariableDescriptorCount:           true,
				RuntimeDescriptorArray:                             false,
			}},
		})
	require.NoError(t, err)
	require.Equal(t, mockDevice.Handle(), device.Handle())
}

func TestPhysicalDeviceDescriptorIndexingFeaturesOutData(t *testing.T) {
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

		next := (*ext_descriptor_indexing_driver.VkPhysicalDeviceDescriptorIndexingFeaturesEXT)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()
		require.Equal(t, uint64(1000161001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DESCRIPTOR_INDEXING_FEATURES_EXT
		require.True(t, val.FieldByName("pNext").IsNil())
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderInputAttachmentArrayDynamicIndexing").UnsafeAddr())) = loader.VkBool32(1)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderUniformTexelBufferArrayDynamicIndexing").UnsafeAddr())) = loader.VkBool32(0)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderStorageTexelBufferArrayDynamicIndexing").UnsafeAddr())) = loader.VkBool32(1)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderUniformBufferArrayNonUniformIndexing").UnsafeAddr())) = loader.VkBool32(0)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderSampledImageArrayNonUniformIndexing").UnsafeAddr())) = loader.VkBool32(1)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderStorageBufferArrayNonUniformIndexing").UnsafeAddr())) = loader.VkBool32(0)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderStorageImageArrayNonUniformIndexing").UnsafeAddr())) = loader.VkBool32(1)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderInputAttachmentArrayNonUniformIndexing").UnsafeAddr())) = loader.VkBool32(0)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderUniformTexelBufferArrayNonUniformIndexing").UnsafeAddr())) = loader.VkBool32(1)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderStorageTexelBufferArrayNonUniformIndexing").UnsafeAddr())) = loader.VkBool32(0)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("descriptorBindingUniformBufferUpdateAfterBind").UnsafeAddr())) = loader.VkBool32(1)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("descriptorBindingSampledImageUpdateAfterBind").UnsafeAddr())) = loader.VkBool32(0)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("descriptorBindingStorageImageUpdateAfterBind").UnsafeAddr())) = loader.VkBool32(1)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("descriptorBindingStorageBufferUpdateAfterBind").UnsafeAddr())) = loader.VkBool32(0)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("descriptorBindingUniformTexelBufferUpdateAfterBind").UnsafeAddr())) = loader.VkBool32(1)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("descriptorBindingStorageTexelBufferUpdateAfterBind").UnsafeAddr())) = loader.VkBool32(0)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("descriptorBindingUpdateUnusedWhilePending").UnsafeAddr())) = loader.VkBool32(1)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("descriptorBindingPartiallyBound").UnsafeAddr())) = loader.VkBool32(0)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("descriptorBindingVariableDescriptorCount").UnsafeAddr())) = loader.VkBool32(1)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("runtimeDescriptorArray").UnsafeAddr())) = loader.VkBool32(0)
	})

	var outData PhysicalDeviceDescriptorIndexingFeatures
	err := extension.GetPhysicalDeviceFeatures2(
		physicalDevice,
		&khr_get_physical_device_properties2.PhysicalDeviceFeatures2{
			NextOutData: common.NextOutData{&outData},
		})
	require.NoError(t, err)
	require.Equal(t, PhysicalDeviceDescriptorIndexingFeatures{
		ShaderInputAttachmentArrayDynamicIndexing:          true,
		ShaderUniformTexelBufferArrayDynamicIndexing:       false,
		ShaderStorageTexelBufferArrayDynamicIndexing:       true,
		ShaderUniformBufferArrayNonUniformIndexing:         false,
		ShaderSampledImageArrayNonUniformIndexing:          true,
		ShaderStorageBufferArrayNonUniformIndexing:         false,
		ShaderStorageImageArrayNonUniformIndexing:          true,
		ShaderInputAttachmentArrayNonUniformIndexing:       false,
		ShaderUniformTexelBufferArrayNonUniformIndexing:    true,
		ShaderStorageTexelBufferArrayNonUniformIndexing:    false,
		DescriptorBindingUniformBufferUpdateAfterBind:      true,
		DescriptorBindingSampledImageUpdateAfterBind:       false,
		DescriptorBindingStorageImageUpdateAfterBind:       true,
		DescriptorBindingStorageBufferUpdateAfterBind:      false,
		DescriptorBindingUniformTexelBufferUpdateAfterBind: true,
		DescriptorBindingStorageTexelBufferUpdateAfterBind: false,
		DescriptorBindingUpdateUnusedWhilePending:          true,
		DescriptorBindingPartiallyBound:                    false,
		DescriptorBindingVariableDescriptorCount:           true,
		RuntimeDescriptorArray:                             false,
	}, outData)
}

func TestPhysicalDeviceDescriptorIndexingOutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_get_physical_device_properties2.NewMockLoader(ctrl)
	extension := khr_get_physical_device_properties2.CreateExtensionDriverFromLoader(extDriver)

	instance := mocks.NewDummyInstance(common.Vulkan1_0, []string{})
	physicalDevice := mocks.NewDummyPhysicalDevice(instance, common.Vulkan1_0)

	extDriver.EXPECT().VkGetPhysicalDeviceProperties2KHR(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice loader.VkPhysicalDevice,
		pProperties *khr_get_physical_device_properties2_driver.VkPhysicalDeviceProperties2KHR) {

		val := reflect.ValueOf(pProperties).Elem()
		require.Equal(t, uint64(1000059001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROPERTIES_2_KHR
		next := (*ext_descriptor_indexing_driver.VkPhysicalDeviceDescriptorIndexingPropertiesEXT)(val.FieldByName("pNext").UnsafePointer())

		val = reflect.ValueOf(next).Elem()
		require.Equal(t, uint64(1000161002), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DESCRIPTOR_INDEXING_PROPERTIES_EXT
		require.True(t, val.FieldByName("pNext").IsNil())
		*(*loader.Uint32)(unsafe.Pointer(val.FieldByName("maxUpdateAfterBindDescriptorsInAllPools").UnsafeAddr())) = loader.Uint32(1)

		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderUniformBufferArrayNonUniformIndexingNative").UnsafeAddr())) = loader.VkBool32(1)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderSampledImageArrayNonUniformIndexingNative").UnsafeAddr())) = loader.VkBool32(0)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderStorageBufferArrayNonUniformIndexingNative").UnsafeAddr())) = loader.VkBool32(1)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderStorageImageArrayNonUniformIndexingNative").UnsafeAddr())) = loader.VkBool32(0)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("shaderInputAttachmentArrayNonUniformIndexingNative").UnsafeAddr())) = loader.VkBool32(1)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("robustBufferAccessUpdateAfterBind").UnsafeAddr())) = loader.VkBool32(0)
		*(*loader.VkBool32)(unsafe.Pointer(val.FieldByName("quadDivergentImplicitLod").UnsafeAddr())) = loader.VkBool32(1)

		*(*loader.Uint32)(unsafe.Pointer(val.FieldByName("maxPerStageDescriptorUpdateAfterBindSamplers").UnsafeAddr())) = loader.Uint32(3)
		*(*loader.Uint32)(unsafe.Pointer(val.FieldByName("maxPerStageDescriptorUpdateAfterBindUniformBuffers").UnsafeAddr())) = loader.Uint32(5)
		*(*loader.Uint32)(unsafe.Pointer(val.FieldByName("maxPerStageDescriptorUpdateAfterBindStorageBuffers").UnsafeAddr())) = loader.Uint32(7)
		*(*loader.Uint32)(unsafe.Pointer(val.FieldByName("maxPerStageDescriptorUpdateAfterBindSampledImages").UnsafeAddr())) = loader.Uint32(11)
		*(*loader.Uint32)(unsafe.Pointer(val.FieldByName("maxPerStageDescriptorUpdateAfterBindStorageImages").UnsafeAddr())) = loader.Uint32(13)
		*(*loader.Uint32)(unsafe.Pointer(val.FieldByName("maxPerStageDescriptorUpdateAfterBindInputAttachments").UnsafeAddr())) = loader.Uint32(17)
		*(*loader.Uint32)(unsafe.Pointer(val.FieldByName("maxPerStageUpdateAfterBindResources").UnsafeAddr())) = loader.Uint32(19)
		*(*loader.Uint32)(unsafe.Pointer(val.FieldByName("maxDescriptorSetUpdateAfterBindSamplers").UnsafeAddr())) = loader.Uint32(23)
		*(*loader.Uint32)(unsafe.Pointer(val.FieldByName("maxDescriptorSetUpdateAfterBindUniformBuffers").UnsafeAddr())) = loader.Uint32(29)
		*(*loader.Uint32)(unsafe.Pointer(val.FieldByName("maxDescriptorSetUpdateAfterBindUniformBuffersDynamic").UnsafeAddr())) = loader.Uint32(31)
		*(*loader.Uint32)(unsafe.Pointer(val.FieldByName("maxDescriptorSetUpdateAfterBindStorageBuffers").UnsafeAddr())) = loader.Uint32(37)
		*(*loader.Uint32)(unsafe.Pointer(val.FieldByName("maxDescriptorSetUpdateAfterBindStorageBuffersDynamic").UnsafeAddr())) = loader.Uint32(41)
		*(*loader.Uint32)(unsafe.Pointer(val.FieldByName("maxDescriptorSetUpdateAfterBindSampledImages").UnsafeAddr())) = loader.Uint32(43)
		*(*loader.Uint32)(unsafe.Pointer(val.FieldByName("maxDescriptorSetUpdateAfterBindStorageImages").UnsafeAddr())) = loader.Uint32(47)
		*(*loader.Uint32)(unsafe.Pointer(val.FieldByName("maxDescriptorSetUpdateAfterBindInputAttachments").UnsafeAddr())) = loader.Uint32(51)
	})

	var outData PhysicalDeviceDescriptorIndexingProperties
	err := extension.GetPhysicalDeviceProperties2(
		physicalDevice,
		&khr_get_physical_device_properties2.PhysicalDeviceProperties2{
			NextOutData: common.NextOutData{&outData},
		})
	require.NoError(t, err)
	require.Equal(t,
		PhysicalDeviceDescriptorIndexingProperties{
			MaxUpdateAfterBindDescriptorsInAllPools: 1,

			ShaderUniformBufferArrayNonUniformIndexingNative:   true,
			ShaderSampledImageArrayNonUniformIndexingNative:    false,
			ShaderStorageBufferArrayNonUniformIndexingNative:   true,
			ShaderStorageImageArrayNonUniformIndexingNative:    false,
			ShaderInputAttachmentArrayNonUniformIndexingNative: true,
			RobustBufferAccessUpdateAfterBind:                  false,
			QuadDivergentImplicitLod:                           true,

			MaxPerStageDescriptorUpdateAfterBindSamplers:         3,
			MaxPerStageDescriptorUpdateAfterBindUniformBuffers:   5,
			MaxPerStageDescriptorUpdateAfterBindStorageBuffers:   7,
			MaxPerStageDescriptorUpdateAfterBindSampledImages:    11,
			MaxPerStageDescriptorUpdateAfterBindStorageImages:    13,
			MaxPerStageDescriptorUpdateAfterBindInputAttachments: 17,
			MaxPerStageUpdateAfterBindResources:                  19,
			MaxDescriptorSetUpdateAfterBindSamplers:              23,
			MaxDescriptorSetUpdateAfterBindUniformBuffers:        29,
			MaxDescriptorSetUpdateAfterBindUniformBuffersDynamic: 31,
			MaxDescriptorSetUpdateAfterBindStorageBuffers:        37,
			MaxDescriptorSetUpdateAfterBindStorageBuffersDynamic: 41,
			MaxDescriptorSetUpdateAfterBindSampledImages:         43,
			MaxDescriptorSetUpdateAfterBindStorageImages:         47,
			MaxDescriptorSetUpdateAfterBindInputAttachments:      51,
		},
		outData)
}
