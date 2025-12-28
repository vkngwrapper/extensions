package khr_maintenance4

import (
	"github.com/vkngwrapper/core/v3/core1_1"
)

//go:generate mockgen -source extiface.go -destination ./mocks/extension.go -package mock_maintenance4

type ExtensionDriver interface {
	// GetDeviceBufferMemoryRequirements returns the memory requirements for a specified Vulkan
	// object.
	//
	// options - Contains the parameters required for the memory requirements query
	//
	// outData - A pre-allocated object in which memory requirements will be populated.
	// It should include any desired chained OutData objects
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDeviceBufferMemoryRequirementsKHR.html
	GetDeviceBufferMemoryRequirements(options DeviceBufferMemoryRequirements, outData *core1_1.MemoryRequirements2) error
	// GetDeviceImageMemoryRequirements returns the memory requirements for a specified Vulkan
	// object.
	//
	// options - Contains the parameters required for the memory requirements query
	//
	// outData - A pre-allocated object in which memory requirements will be populated.
	// It should include any desired chained OutData objects
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDeviceImageMemoryRequirementsKHR.html
	GetDeviceImageMemoryRequirements(options DeviceImageMemoryRequirements, outData *core1_1.MemoryRequirements2) error
	// GetDeviceImageSparseMemoryRequirements queries the memory requirements for a sparse image
	//
	// options - Contains the parameters required for the memory requirements query
	//
	// outDataFactory - This method can be provided to allocate each SparseImageMemoryRequirements2 object
	// that is returned, along with any chained OutData structures. It can also be left nil, in which case
	// SparseImageMemoryRequirements2 will be allocated with no chained structures.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/vkGetDeviceImageSparseMemoryRequirements.html
	GetDeviceImageSparseMemoryRequirements(options DeviceImageMemoryRequirements, outDataFactory func() *core1_1.SparseImageMemoryRequirements2) ([]*core1_1.SparseImageMemoryRequirements2, error)
}
