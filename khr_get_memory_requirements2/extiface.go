package khr_get_memory_requirements2

import "github.com/vkngwrapper/core/core1_0"

//go:generate mockgen -source extiface.go -destination ./mocks/extension.go -package mock_get_memory_requirements2

// Extension contains all commands for the khr_get_memory_requirements2 extension
type Extension interface {
	// BufferMemoryRequirements2 returns the memory requirements for the specified Vulkan object
	//
	// device - The Device which owns the Buffer
	//
	// o - Contains parameters required for the memory requirements query
	//
	// out - A pre-allocated object in which the memory requirements of the Buffer object will be
	// populated. It should include any desired chained OutData objects
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetBufferMemoryRequirements2.html
	BufferMemoryRequirements2(device core1_0.Device, o BufferMemoryRequirementsInfo2, out *MemoryRequirements2) error
	// ImageMemoryRequirements2 returns the memory requirements for the specified Vulkan object
	//
	// device - The Device which owns the Image
	//
	// o - Contains parameters required for the memory requirements query
	//
	// out - A pre-allocated object in which the memory requirements of the Image object will be
	// populated. It should include any desired chained OutData objects
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetImageMemoryRequirements2.html
	ImageMemoryRequirements2(device core1_0.Device, o ImageMemoryRequirementsInfo2, out *MemoryRequirements2) error
	// ImageSparseMemoryRequirements2 queries the memory requirements for a sparse Image
	//
	// device - The Device which owns the Image
	//
	// o - Contains parameters required for the memory requirements query
	//
	// outDataFactory - This method can be provided to allocate each SparseImageMemoryRequirements2 object
	// that is returned, along with any chained OutData structures. It can also be left nil, in which case
	// SparseImageMemoryRequirements2 will be allocated with no chained structures.
	ImageSparseMemoryRequirements2(device core1_0.Device, o ImageSparseMemoryRequirementsInfo2, outDataFactory func() *SparseImageMemoryRequirements2) ([]*SparseImageMemoryRequirements2, error)
}
