package khr_8bit_storage

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/common"
	"unsafe"
)

// PhysicalDevice8BitStorageFeatures describes features supported by khr_8bit_storage
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice8BitStorageFeatures.html
type PhysicalDevice8BitStorageFeatures struct {
	// StorageBuffer8BitAccess indicates whether objects in the StorageBuffer, ShaderRecordBufferKHR,
	// or PhysicalStorageBuffer storage class with the Block decoration can have 8-bit integer members
	StorageBuffer8BitAccess bool
	// UniformAndStorageBuffer8BitAccess indicates whether objects in the Uniform storage class
	// with the Block decoration can have 8-bit integer members
	UniformAndStorageBuffer8BitAccess bool
	// StoragePushConstant8 indicates whether objects in the PushConstant storage class can have 8-bit
	// integer members
	StoragePushConstant8 bool

	common.NextOptions
	common.NextOutData
}

func (o *PhysicalDevice8BitStorageFeatures) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDevice8BitStorageFeaturesKHR{})))
	}

	outData := (*C.VkPhysicalDevice8BitStorageFeaturesKHR)(preallocatedPointer)
	outData.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_8BIT_STORAGE_FEATURES_KHR
	outData.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDevice8BitStorageFeatures) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	outData := (*C.VkPhysicalDevice8BitStorageFeaturesKHR)(cDataPointer)
	o.StoragePushConstant8 = outData.storagePushConstant8 != C.VkBool32(0)
	o.UniformAndStorageBuffer8BitAccess = outData.uniformAndStorageBuffer8BitAccess != C.VkBool32(0)
	o.StorageBuffer8BitAccess = outData.storageBuffer8BitAccess != C.VkBool32(0)

	return outData.pNext, nil
}

func (o PhysicalDevice8BitStorageFeatures) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDevice8BitStorageFeatures{})))
	}

	info := (*C.VkPhysicalDevice8BitStorageFeatures)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_8BIT_STORAGE_FEATURES_KHR
	info.pNext = next
	info.storageBuffer8BitAccess = C.VkBool32(0)
	info.uniformAndStorageBuffer8BitAccess = C.VkBool32(0)
	info.storagePushConstant8 = C.VkBool32(0)

	if o.StorageBuffer8BitAccess {
		info.storageBuffer8BitAccess = C.VkBool32(1)
	}

	if o.UniformAndStorageBuffer8BitAccess {
		info.uniformAndStorageBuffer8BitAccess = C.VkBool32(1)
	}

	if o.StoragePushConstant8 {
		info.storagePushConstant8 = C.VkBool32(1)
	}

	return preallocatedPointer, nil
}
