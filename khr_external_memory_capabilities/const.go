package khr_external_memory_capabilities

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import "github.com/vkngwrapper/core/v2/common"

// ExternalMemoryFeatureFlags specifies features of an external memory handle type
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryFeatureFlagBits.html
type ExternalMemoryFeatureFlags int32

var externalMemoryFeaturesMapping = common.NewFlagStringMapping[ExternalMemoryFeatureFlags]()

func (f ExternalMemoryFeatureFlags) Register(str string) {
	externalMemoryFeaturesMapping.Register(f, str)
}

func (f ExternalMemoryFeatureFlags) String() string {
	return externalMemoryFeaturesMapping.FlagsToString(f)
}

////

// ExternalMemoryHandleTypeFlags specifies external memory handle types
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlagBits.html
type ExternalMemoryHandleTypeFlags int32

var externalMemoryHandleTypesMapping = common.NewFlagStringMapping[ExternalMemoryHandleTypeFlags]()

func (f ExternalMemoryHandleTypeFlags) Register(str string) {
	externalMemoryHandleTypesMapping.Register(f, str)
}

func (f ExternalMemoryHandleTypeFlags) String() string {
	return externalMemoryHandleTypesMapping.FlagsToString(f)
}

////

const (
	// ExtensionName is "VK_KHR_external_memory_capabilities"
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_external_memory_capabilities.html
	ExtensionName string = C.VK_KHR_EXTERNAL_MEMORY_CAPABILITIES_EXTENSION_NAME

	// LUIDSize is the length of a locally unique Device identifier
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_LUID_SIZE.html
	LUIDSize int = C.VK_LUID_SIZE_KHR

	// ExternalMemoryFeatureDedicatedOnly specifies that Image or Buffer objects created with the
	// specified parameters and handle type must create or import a dedicated allocation for
	// the Image or Buffer object
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryFeatureFlagBits.html
	ExternalMemoryFeatureDedicatedOnly ExternalMemoryFeatureFlags = C.VK_EXTERNAL_MEMORY_FEATURE_DEDICATED_ONLY_BIT_KHR
	// ExternalMemoryFeatureExportable specifies that handles of this type can be exported from
	// Vulkan memory objects
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryFeatureFlagBits.html
	ExternalMemoryFeatureExportable ExternalMemoryFeatureFlags = C.VK_EXTERNAL_MEMORY_FEATURE_EXPORTABLE_BIT_KHR
	// ExternalMemoryFeatureImportable specifies that handles of this type can be imported as Vulkan
	// memory objects
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryFeatureFlagBits.html
	ExternalMemoryFeatureImportable ExternalMemoryFeatureFlags = C.VK_EXTERNAL_MEMORY_FEATURE_IMPORTABLE_BIT_KHR

	// ExternalMemoryHandleTypeD3D11Texture specifies an NT handle returned by
	// IDXGIResource1::CreateSharedHandle referring to a Direct3D 10 or 11 texture resource
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlagBits.html
	ExternalMemoryHandleTypeD3D11Texture ExternalMemoryHandleTypeFlags = C.VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D11_TEXTURE_BIT_KHR
	// ExternalMemoryHandleTypeD3D11TextureKMT specifies a global share handle returned by
	// IDXGIResource::GetSharedHandle referring to a Direct3D 10 or 11 texture resource
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlagBits.html
	ExternalMemoryHandleTypeD3D11TextureKMT ExternalMemoryHandleTypeFlags = C.VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D11_TEXTURE_KMT_BIT_KHR
	// ExternalMemoryHandleTypeD3D12Heap specifies an NT handle returned by
	// ID3D12Device::CreateSharedHandle referring to a Direct3D 12 heap resource
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlagBits.html
	ExternalMemoryHandleTypeD3D12Heap ExternalMemoryHandleTypeFlags = C.VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D12_HEAP_BIT_KHR
	// ExternalMemoryHandleTypeD3D12Resource specifies an NT handle returned by
	// ID3D12Device::CreateSharedHandle referring to a Direct3D 12 committed resource
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlagBits.html
	ExternalMemoryHandleTypeD3D12Resource ExternalMemoryHandleTypeFlags = C.VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D12_RESOURCE_BIT_KHR
	// ExternalMemoryHandleTypeOpaqueFD specifies a POSIX file descriptor handle that has only limited
	// valid usage outside of Vulkan and other compatible APIs
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlagBits.html
	ExternalMemoryHandleTypeOpaqueFD ExternalMemoryHandleTypeFlags = C.VK_EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_FD_BIT_KHR
	// ExternalMemoryHandleTypeOpaqueWin32 specifies an NT handle that has only limited valid usage
	// outside of Vulkan and other compatible APIs
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlagBits.html
	ExternalMemoryHandleTypeOpaqueWin32 ExternalMemoryHandleTypeFlags = C.VK_EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_WIN32_BIT_KHR
	// ExternalMemoryHandleTypeOpaqueWin32KMT specifies a global share handle that has only
	// limited valid usage outside of Vulkan and other compatible APIs
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlagBits.html
	ExternalMemoryHandleTypeOpaqueWin32KMT ExternalMemoryHandleTypeFlags = C.VK_EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT_KHR
)

func init() {
	ExternalMemoryFeatureDedicatedOnly.Register("Dedicated Only")
	ExternalMemoryFeatureExportable.Register("Exportable")
	ExternalMemoryFeatureImportable.Register("Importable")

	ExternalMemoryHandleTypeD3D11Texture.Register("D3D11 Texture")
	ExternalMemoryHandleTypeD3D11TextureKMT.Register("D3D11 Texture (Kernel-Mode)")
	ExternalMemoryHandleTypeD3D12Heap.Register("D3D12 Heap")
	ExternalMemoryHandleTypeD3D12Resource.Register("D3D12 Resource")
	ExternalMemoryHandleTypeOpaqueFD.Register("Opaque File Descriptor")
	ExternalMemoryHandleTypeOpaqueWin32.Register("Opaque Win32")
	ExternalMemoryHandleTypeOpaqueWin32KMT.Register("Opaque Win32 (Kernel-Mode)")
}
