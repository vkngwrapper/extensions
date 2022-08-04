package khr_external_memory_capabilities

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

// ExternalMemoryProperties specifies external memory handle type capabilities
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryProperties.html
type ExternalMemoryProperties struct {
	// ExternalMemoryFeatures specifies the features of the handle type
	ExternalMemoryFeatures ExternalMemoryFeatureFlags
	// ExportFromImportedHandleTypes specifies which types of imported handle the handle type can
	// be exported from
	ExportFromImportedHandleTypes ExternalMemoryHandleTypeFlags
	// CompatibleHandleTypes specifies handle types which can be specified at the same time as the
	// handle type which creating an Image compatible with external memory
	CompatibleHandleTypes ExternalMemoryHandleTypeFlags
}

func (o ExternalMemoryProperties) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkExternalMemoryPropertiesKHR{})))
	}

	info := (*C.VkExternalMemoryPropertiesKHR)(preallocatedPointer)
	info.externalMemoryFeatures = C.VkExternalMemoryFeatureFlags(o.ExternalMemoryFeatures)
	info.exportFromImportedHandleTypes = C.VkExternalMemoryHandleTypeFlags(o.ExportFromImportedHandleTypes)
	info.compatibleHandleTypes = C.VkExternalMemoryHandleTypeFlags(o.CompatibleHandleTypes)

	return preallocatedPointer, nil
}

func (o *ExternalMemoryProperties) PopulateOutData(cDataPointer unsafe.Pointer) error {
	info := (*C.VkExternalMemoryPropertiesKHR)(cDataPointer)
	o.ExternalMemoryFeatures = ExternalMemoryFeatureFlags(info.externalMemoryFeatures)
	o.ExportFromImportedHandleTypes = ExternalMemoryHandleTypeFlags(info.exportFromImportedHandleTypes)
	o.CompatibleHandleTypes = ExternalMemoryHandleTypeFlags(info.compatibleHandleTypes)

	return nil
}
