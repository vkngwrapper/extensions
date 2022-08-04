package khr_external_semaphore_capabilities

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

// ExternalSemaphoreProperties describes supported external Semaphore handle features
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalSemaphoreProperties.html
type ExternalSemaphoreProperties struct {
	// ExportFromImportedHandleTypes specifies which types of imported handle HandleType can
	// be exported from
	ExportFromImportedHandleTypes ExternalSemaphoreHandleTypeFlags
	// CompatibleHandleTypes specifies handle types which can be specified at the same time as
	// HandleType when creating a Semaphore
	CompatibleHandleTypes ExternalSemaphoreHandleTypeFlags
	// ExternalSemaphoreFeatures describes the features of HandleType
	ExternalSemaphoreFeatures ExternalSemaphoreFeatureFlags

	common.NextOutData
}

func (o *ExternalSemaphoreProperties) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkExternalSemaphorePropertiesKHR{})))
	}

	info := (*C.VkExternalSemaphorePropertiesKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_EXTERNAL_SEMAPHORE_PROPERTIES_KHR
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *ExternalSemaphoreProperties) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkExternalSemaphorePropertiesKHR)(cDataPointer)

	o.ExportFromImportedHandleTypes = ExternalSemaphoreHandleTypeFlags(info.exportFromImportedHandleTypes)
	o.CompatibleHandleTypes = ExternalSemaphoreHandleTypeFlags(info.compatibleHandleTypes)
	o.ExternalSemaphoreFeatures = ExternalSemaphoreFeatureFlags(info.externalSemaphoreFeatures)

	return info.pNext, nil
}
