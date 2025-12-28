package khr_external_semaphore_capabilities

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/vkngwrapper/core/v3/core1_1"
)

// ExternalSemaphoreFeatureFlags describes features of an external Semaphore handle type
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalSemaphoreFeatureFlagBits.html
type ExternalSemaphoreFeatureFlags = core1_1.ExternalSemaphoreFeatureFlags

////

// ExternalSemaphoreHandleTypeFlags is a bitmask of valid external Semaphore handle types
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalSemaphoreHandleTypeFlagBits.html
type ExternalSemaphoreHandleTypeFlags = core1_1.ExternalSemaphoreHandleTypeFlags

////

const (
	// ExtensionName is "VK_KHR_external_semaphore_capabilities"
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_external_semaphore_capabilities.html
	ExtensionName string = C.VK_KHR_EXTERNAL_SEMAPHORE_CAPABILITIES_EXTENSION_NAME

	// LUIDSize is the length of a locally unique Device identifier
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_LUID_SIZE.html
	LUIDSize int = C.VK_LUID_SIZE_KHR

	// ExternalSemaphoreFeatureExportable specifies that handles of this type can be exported
	// from Vulkan Semaphore objects
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalSemaphoreFeatureFlagBits.html
	ExternalSemaphoreFeatureExportable ExternalSemaphoreFeatureFlags = C.VK_EXTERNAL_SEMAPHORE_FEATURE_EXPORTABLE_BIT_KHR
	// ExternalSemaphoreFeatureImportable specifies that handles of this type can be imported
	// as Vulkan Semaphore objects
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalSemaphoreFeatureFlagBits.html
	ExternalSemaphoreFeatureImportable ExternalSemaphoreFeatureFlags = C.VK_EXTERNAL_SEMAPHORE_FEATURE_IMPORTABLE_BIT_KHR

	// ExternalSemaphoreHandleTypeOpaqueFD specifies a POSIX file descriptor handle that has
	// only limited valid usage outside of Vulkan and other compatible APIs
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalSemaphoreHandleTypeFlagBits.html
	ExternalSemaphoreHandleTypeOpaqueFD ExternalSemaphoreHandleTypeFlags = C.VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_FD_BIT_KHR
	// ExternalSemaphoreHandleTypeOpaqueWin32 specifies an NT handle that has only limited
	// valid usage outside of Vulkan and other compatible APIs
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalSemaphoreHandleTypeFlagBits.html
	ExternalSemaphoreHandleTypeOpaqueWin32 ExternalSemaphoreHandleTypeFlags = C.VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_WIN32_BIT_KHR
	// ExternalSemaphoreHandleTypeOpaqueWin32KMT specifies a global share handle that has only
	// limited valid usage outside of Vulkan and other compatible APIs
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalSemaphoreHandleTypeFlagBits.html
	ExternalSemaphoreHandleTypeOpaqueWin32KMT ExternalSemaphoreHandleTypeFlags = C.VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT_KHR
	// ExternalSemaphoreHandleTypeD3D12Fence specifies an NT handle returned by
	// ID3D12Device::CreateSharedHandle referring to a Direct3D 12 fence, or
	// ID3D11Device5::CreateFence referring to a Direct3D 11 fence
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalSemaphoreHandleTypeFlagBits.html
	ExternalSemaphoreHandleTypeD3D12Fence ExternalSemaphoreHandleTypeFlags = C.VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_D3D12_FENCE_BIT_KHR
	// ExternalSemaphoreHandleTypeSyncFD specifies a POSIX file descriptor handle to a Linux Sync
	// File or Android Fence object. It can be used with any native API accepting a valid sync file or
	// Fence as input
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalSemaphoreHandleTypeFlagBits.html
	ExternalSemaphoreHandleTypeSyncFD ExternalSemaphoreHandleTypeFlags = C.VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_SYNC_FD_BIT_KHR
)

func init() {
	ExternalSemaphoreFeatureExportable.Register("Exportable")
	ExternalSemaphoreFeatureImportable.Register("Importable")

	ExternalSemaphoreHandleTypeOpaqueFD.Register("Opaque File Descriptor")
	ExternalSemaphoreHandleTypeOpaqueWin32.Register("Opaque Win32 Handle")
	ExternalSemaphoreHandleTypeOpaqueWin32KMT.Register("Opaque Win32 Handle (Kernel Mode)")
	ExternalSemaphoreHandleTypeD3D12Fence.Register("D3D Fence")
	ExternalSemaphoreHandleTypeSyncFD.Register("Sync File Descriptor")
}
