package ext_memory_budget_test

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/driver"
	mock_driver "github.com/vkngwrapper/core/v2/driver/mocks"
	"github.com/vkngwrapper/core/v2/mocks"
	"github.com/vkngwrapper/extensions/v2/ext_memory_budget"
	ext_memory_budget_driver "github.com/vkngwrapper/extensions/v2/ext_memory_budget/driver"
	"github.com/vkngwrapper/extensions/v2/khr_get_physical_device_properties2"
	khr_get_physical_device_properties2_driver "github.com/vkngwrapper/extensions/v2/khr_get_physical_device_properties2/driver"
	mock_get_physical_device_properties2 "github.com/vkngwrapper/extensions/v2/khr_get_physical_device_properties2/mocks"
	"reflect"
	"testing"
	"unsafe"
)

func TestMemoryBudgetPropertiesOutData(t *testing.T) {
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
		pProperties *khr_get_physical_device_properties2_driver.VkPhysicalDeviceProperties2KHR,
	) {
		val := reflect.ValueOf(pProperties).Elem()
		require.Equal(t, uint64(1000059001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROPERTIES_2_KHR

		next := (*ext_memory_budget_driver.VkPhysicalDeviceMemoryBudgetPropertiesEXT)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000237000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MEMORY_BUDGET_PROPERTIES_EXT
		require.True(t, val.FieldByName("pNext").IsNil())

		heapBudgetSlice := unsafe.Slice((*driver.VkDeviceSize)(unsafe.Pointer(val.FieldByName("heapBudget").UnsafeAddr())), common.MaxMemoryHeaps)
		heapUsageSlice := unsafe.Slice((*driver.VkDeviceSize)(unsafe.Pointer(val.FieldByName("heapUsage").UnsafeAddr())), common.MaxMemoryHeaps)
		for i := 0; i < common.MaxMemoryHeaps; i++ {
			heapBudgetSlice[i] = driver.VkDeviceSize(i)
			heapUsageSlice[i] = driver.VkDeviceSize(i + 10)
		}

	})

	var outData ext_memory_budget.PhysicalDeviceMemoryBudgetProperties
	properties := khr_get_physical_device_properties2.PhysicalDeviceProperties2{
		NextOutData: common.NextOutData{&outData},
	}

	err := extension.PhysicalDeviceProperties2(physicalDevice, &properties)
	require.NoError(t, err)
	require.Equal(t, ext_memory_budget.PhysicalDeviceMemoryBudgetProperties{
		HeapBudget: [common.MaxMemoryHeaps]int{
			0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15,
		},
		HeapUsage: [common.MaxMemoryHeaps]int{
			10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25,
		},
	}, outData)
}
