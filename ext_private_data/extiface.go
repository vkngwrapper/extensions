package ext_private_data

import (
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
)

//go:generate go run go.uber.org/mock/mockgen -source extiface.go -destination ./mocks/extension.go -package mock_private_data

// ExtensionDriver contains all commands for the ext_private_data
//
// https://docs.vulkan.org/refpages/latest/refpages/source/VK_EXT_private_data.html
type ExtensionDriver interface {
	// CreatePrivateDataSlot creates a slot for private data storage
	//
	// allocationCallbacks - controls host memory allocation
	//
	// options - PrivateDataSlotCreateInfo to configure the new PriavateDataSlot
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/vkCreatePrivateDataSlot.html
	CreatePrivateDataSlot(allocationCallbacks *loader.AllocationCallbacks, options PrivateDataSlotCreateInfo) (PrivateDataSlot, common.VkResult, error)

	// DestroyPrivateDataSlot destroys a private data slot
	//
	// privateDataSlot - the private data slot to destroy
	//
	// allocator - controls host memory allocation
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/vkDestroyPrivateDataSlot.html
	DestroyPrivateDataSlot(privateDataSlot PrivateDataSlot, allocator *loader.AllocationCallbacks)

	// GetPrivateData retrieves data associated with a Vulkan object
	//
	// objectType - specifies the type of object data is associated with
	//
	// handle - a handle to the object data is associated with
	//
	// privateDataSlot - specifies location of private data pointer storage. the configuration of this object
	// also determines what value is returned in the case that no data has been set for the object
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/vkGetPrivateData.html
	GetPrivateData(objectType core1_0.ObjectType, handle loader.VulkanHandle, privateDataSlot PrivateDataSlot) uint64

	// SetPrivateData associates data with a Vulkan object
	//
	// objectType - specifies the type of object to associate data with
	//
	// handle - a handle to the object to associate data with
	//
	// privateDataSlot - specifies location of private data storage
	//
	// data - application-defined data to associate the object with
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/vkSetPrivateData.html
	SetPrivateData(objectType core1_0.ObjectType, handle loader.VulkanHandle, privateDataSlot PrivateDataSlot, data uint64) (common.VkResult, error)
}
