package khr_16bit_storage

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v2/common"
	"unsafe"
)

// PhysicalDevice16BitStorageFeatures describes features supported by khr_16bit_storage
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice16BitStorageFeatures.html
type PhysicalDevice16BitStorageFeatures struct {
	// StorageBuffer16BitAccess specifies whether objects in the StorageBuffer, ShaderRecordBufferKHR,
	// or PhysicalStorageBuffer storage class with the Block decoration can have 16-bit integer and
	// 16-bit floating-point members
	StorageBuffer16BitAccess bool
	// UniformAndStorageBuffer16BitAccess specifies whether objects in the Uniform storage class with
	// the Block decoration can have 16-bit integer and 16-bit floating-point members
	UniformAndStorageBuffer16BitAccess bool
	// StoragePushConstant16 specifies whether objects in the PushConstant storage class can have
	// 16-bit integer and 16-bit floating-point members
	StoragePushConstant16 bool
	// StorageInputOutput16 specifies whether objects in the Input and Output storage classes can
	// have 16-bit integer and 16-bit floating-point members
	StorageInputOutput16 bool

	common.NextOptions
	common.NextOutData
}

func (o *PhysicalDevice16BitStorageFeatures) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDevice16BitStorageFeaturesKHR{})))
	}

	data := (*C.VkPhysicalDevice16BitStorageFeaturesKHR)(preallocatedPointer)
	data.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_16BIT_STORAGE_FEATURES_KHR
	data.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDevice16BitStorageFeatures) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	data := (*C.VkPhysicalDevice16BitStorageFeaturesKHR)(cDataPointer)

	o.StorageBuffer16BitAccess = data.storageBuffer16BitAccess != C.VkBool32(0)
	o.UniformAndStorageBuffer16BitAccess = data.uniformAndStorageBuffer16BitAccess != C.VkBool32(0)
	o.StoragePushConstant16 = data.storagePushConstant16 != C.VkBool32(0)
	o.StorageInputOutput16 = data.storageInputOutput16 != C.VkBool32(0)

	return data.pNext, nil
}

func (o PhysicalDevice16BitStorageFeatures) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDevice16BitStorageFeaturesKHR{})))
	}

	data := (*C.VkPhysicalDevice16BitStorageFeaturesKHR)(preallocatedPointer)
	data.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_16BIT_STORAGE_FEATURES_KHR
	data.pNext = next

	data.storageBuffer16BitAccess = C.VkBool32(0)
	data.uniformAndStorageBuffer16BitAccess = C.VkBool32(0)
	data.storagePushConstant16 = C.VkBool32(0)
	data.storageInputOutput16 = C.VkBool32(0)

	if o.StorageBuffer16BitAccess {
		data.storageBuffer16BitAccess = C.VkBool32(1)
	}

	if o.UniformAndStorageBuffer16BitAccess {
		data.uniformAndStorageBuffer16BitAccess = C.VkBool32(1)
	}

	if o.StoragePushConstant16 {
		data.storagePushConstant16 = C.VkBool32(1)
	}

	if o.StorageInputOutput16 {
		data.storageInputOutput16 = C.VkBool32(1)
	}

	return preallocatedPointer, nil
}
