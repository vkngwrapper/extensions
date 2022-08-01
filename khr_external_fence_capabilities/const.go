package khr_external_fence_capabilities

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import "github.com/vkngwrapper/core/common"

// ExternalFenceFeatureFlags describes features of an external Fence handle type
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFenceFeatureFlagBits.html
type ExternalFenceFeatureFlags int32

var externalFenceFeaturesMapping = common.NewFlagStringMapping[ExternalFenceFeatureFlags]()

func (f ExternalFenceFeatureFlags) Register(str string) {
	externalFenceFeaturesMapping.Register(f, str)
}

func (f ExternalFenceFeatureFlags) String() string {
	return externalFenceFeaturesMapping.FlagsToString(f)
}

////

// ExternalFenceHandleTypeFlags is a bitmask of valid external Fence handle types
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFenceHandleTypeFlagBits.html
type ExternalFenceHandleTypeFlags int32

var externalFenceHandleTypesMapping = common.NewFlagStringMapping[ExternalFenceHandleTypeFlags]()

func (f ExternalFenceHandleTypeFlags) Register(str string) {
	externalFenceHandleTypesMapping.Register(f, str)
}

func (f ExternalFenceHandleTypeFlags) String() string {
	return externalFenceHandleTypesMapping.FlagsToString(f)
}

////

const (
	// ExtensionName is "VK_KHR_external_fence_capabilities"
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_external_fence_capabilities.html
	ExtensionName string = C.VK_KHR_EXTERNAL_FENCE_CAPABILITIES_EXTENSION_NAME

	// LUIDSize is the length of a locally unique Device identifier
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_LUID_SIZE.html
	LUIDSize int = C.VK_LUID_SIZE_KHR

	// ExternalFenceFeatureExportable specifies handles of this type can be exported from Vulkan
	// Fence objects
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFenceFeatureFlagBits.html
	ExternalFenceFeatureExportable ExternalFenceFeatureFlags = C.VK_EXTERNAL_FENCE_FEATURE_EXPORTABLE_BIT_KHR
	// ExternalFenceFeatureImportable specifies handles of this type can be imported to Vulkan Fence
	// objects
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFenceFeatureFlagBits.html
	ExternalFenceFeatureImportable ExternalFenceFeatureFlags = C.VK_EXTERNAL_FENCE_FEATURE_IMPORTABLE_BIT_KHR

	// ExternalFenceHandleTypeOpaqueFD specifies a POSIX file descriptor handle that has only limited
	// valid usage outside of Vulkan and other compatible APIs
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFenceHandleTypeFlagBits.html
	ExternalFenceHandleTypeOpaqueFD ExternalFenceHandleTypeFlags = C.VK_EXTERNAL_FENCE_HANDLE_TYPE_OPAQUE_FD_BIT_KHR
	// ExternalFenceHandleTypeOpaqueWin32 specifies an NT handle that has only limited valid usage
	// outside of Vulkan and other compatible APIs
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFenceHandleTypeFlagBits.html
	ExternalFenceHandleTypeOpaqueWin32 ExternalFenceHandleTypeFlags = C.VK_EXTERNAL_FENCE_HANDLE_TYPE_OPAQUE_WIN32_BIT_KHR
	// ExternalFenceHandleTypeOpaqueWin32KMT specifies a global share handle that has only limited
	// usage outside of Vulkan and other compatible APIs
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFenceHandleTypeFlagBits.html
	ExternalFenceHandleTypeOpaqueWin32KMT ExternalFenceHandleTypeFlags = C.VK_EXTERNAL_FENCE_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT_KHR
	// ExternalFenceHandleTypeSyncFD specifies a POSIX file descriptor handle to a Linux Sync File
	// or Android Fence
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFenceHandleTypeFlagBits.html
	ExternalFenceHandleTypeSyncFD ExternalFenceHandleTypeFlags = C.VK_EXTERNAL_FENCE_HANDLE_TYPE_SYNC_FD_BIT_KHR
)

func init() {
	ExternalFenceFeatureExportable.Register("Exportable")
	ExternalFenceFeatureImportable.Register("Importable")

	ExternalFenceHandleTypeOpaqueFD.Register("Opaque File Descriptor")
	ExternalFenceHandleTypeOpaqueWin32.Register("Opaque Win32")
	ExternalFenceHandleTypeOpaqueWin32KMT.Register("Opaque Win32 Kernel-Mode")
	ExternalFenceHandleTypeSyncFD.Register("Sync File Descriptor")
}
