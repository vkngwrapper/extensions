package khr_external_fence_capabilities

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/common"
	"unsafe"
)

// ExternalFenceProperties describes supported external Fence handle features
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFenceProperties.html
type ExternalFenceProperties struct {
	// ExportFromImportedHandleTypes indicates which type of imported handle HandleType can be exported from
	ExportFromImportedHandleTypes ExternalFenceHandleTypeFlags
	// CompatibleHandleTypes specifies handle types which can be specified at the same time as HandleType when creating a Fence
	CompatibleHandleTypes ExternalFenceHandleTypeFlags
	// ExternalFenceFeatures indicates the features of HandleType
	ExternalFenceFeatures ExternalFenceFeatureFlags

	common.NextOutData
}

func (o *ExternalFenceProperties) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkExternalFencePropertiesKHR{})))
	}

	info := (*C.VkExternalFencePropertiesKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_EXTERNAL_FENCE_PROPERTIES_KHR
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *ExternalFenceProperties) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkExternalFencePropertiesKHR)(cDataPointer)

	o.ExportFromImportedHandleTypes = ExternalFenceHandleTypeFlags(info.exportFromImportedHandleTypes)
	o.CompatibleHandleTypes = ExternalFenceHandleTypeFlags(info.compatibleHandleTypes)
	o.ExternalFenceFeatures = ExternalFenceFeatureFlags(info.externalFenceFeatures)

	return info.pNext, nil
}
