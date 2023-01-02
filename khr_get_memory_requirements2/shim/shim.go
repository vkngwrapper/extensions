package khr_get_memory_requirements2_shim

import (
	"github.com/vkngwrapper/core/v2/core1_0"
	"github.com/vkngwrapper/core/v2/core1_1"
	"github.com/vkngwrapper/extensions/v2/khr_get_memory_requirements2"
)

//go:generate mockgen -source shim.go -destination ../mocks/shim.go -package mock_get_memory_requirements2

// Shim contains all commands for the khr_get_memory_requirements2 extension
type Shim interface {
	// BufferMemoryRequirements2 returns the memory requirements for the specified Vulkan object
	//
	// o - Contains parameters required for the memory requirements query
	//
	// out - A pre-allocated object in which the memory requirements of the Buffer object will be
	// populated. It should include any desired chained OutData objects
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetBufferMemoryRequirements2.html
	BufferMemoryRequirements2(o core1_1.BufferMemoryRequirementsInfo2, out *core1_1.MemoryRequirements2) error
	// ImageMemoryRequirements2 returns the memory requirements for the specified Vulkan object
	//
	// o - Contains parameters required for the memory requirements query
	//
	// out - A pre-allocated object in which the memory requirements of the Image object will be
	// populated. It should include any desired chained OutData objects
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetImageMemoryRequirements2.html
	ImageMemoryRequirements2(o core1_1.ImageMemoryRequirementsInfo2, out *core1_1.MemoryRequirements2) error
	// ImageSparseMemoryRequirements2 queries the memory requirements for a sparse Image
	//
	// o - Contains parameters required for the memory requirements query
	//
	// outDataFactory - This method can be provided to allocate each SparseImageMemoryRequirements2 object
	// that is returned, along with any chained OutData structures. It can also be left nil, in which case
	// SparseImageMemoryRequirements2 will be allocated with no chained structures.
	ImageSparseMemoryRequirements2(o core1_1.ImageSparseMemoryRequirementsInfo2, outDataFactory func() *core1_1.SparseImageMemoryRequirements2) ([]*core1_1.SparseImageMemoryRequirements2, error)
}

// Compiler check that VulkanShim satisfies Shim
var _ Shim = &VulkanShim{}

type VulkanShim struct {
	extension khr_get_memory_requirements2.Extension
	device    core1_0.Device
}

func NewShim(extension khr_get_memory_requirements2.Extension, device core1_0.Device) *VulkanShim {
	if extension == nil {
		panic("extension cannot be nil")
	}
	if device == nil {
		panic("device cannot be nil")
	}

	return &VulkanShim{
		extension: extension,
		device:    device,
	}
}

func (s *VulkanShim) BufferMemoryRequirements2(o core1_1.BufferMemoryRequirementsInfo2, out *core1_1.MemoryRequirements2) error {
	return s.extension.BufferMemoryRequirements2(
		s.device,
		khr_get_memory_requirements2.BufferMemoryRequirementsInfo2(o),
		(*khr_get_memory_requirements2.MemoryRequirements2)(out))
}

func (s *VulkanShim) ImageMemoryRequirements2(o core1_1.ImageMemoryRequirementsInfo2, out *core1_1.MemoryRequirements2) error {
	return s.extension.ImageMemoryRequirements2(
		s.device,
		khr_get_memory_requirements2.ImageMemoryRequirementsInfo2(o),
		(*khr_get_memory_requirements2.MemoryRequirements2)(out),
	)
}

func (s *VulkanShim) ImageSparseMemoryRequirements2(o core1_1.ImageSparseMemoryRequirementsInfo2, outDataFactory func() *core1_1.SparseImageMemoryRequirements2) ([]*core1_1.SparseImageMemoryRequirements2, error) {
	factory := func() *khr_get_memory_requirements2.SparseImageMemoryRequirements2 {
		return (*khr_get_memory_requirements2.SparseImageMemoryRequirements2)(outDataFactory())
	}

	retVal, err := s.extension.ImageSparseMemoryRequirements2(
		s.device,
		khr_get_memory_requirements2.ImageSparseMemoryRequirementsInfo2(o),
		factory,
	)

	castRetVal := make([]*core1_1.SparseImageMemoryRequirements2, 0, len(retVal))
	for _, req := range retVal {
		castRetVal = append(castRetVal, (*core1_1.SparseImageMemoryRequirements2)(req))
	}

	return castRetVal, err
}
