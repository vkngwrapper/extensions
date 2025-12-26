package khr_get_memory_requirements2_loader

/*
#include <stdlib.h>
#include "../../vulkan/vulkan.h"

void cgoGetBufferMemoryRequirements2KHR(PFN_vkGetBufferMemoryRequirements2 fn, VkDevice device, VkBufferMemoryRequirementsInfo2KHR *pInfo, VkMemoryRequirements2KHR *pMemoryRequirements) {
	fn(device, pInfo, pMemoryRequirements);
}

void cgoGetImageMemoryRequirements2KHR(PFN_vkGetImageMemoryRequirements2KHR fn, VkDevice device, VkImageMemoryRequirementsInfo2KHR *pInfo, VkMemoryRequirements2KHR *pMemoryRequirements) {
	fn(device, pInfo, pMemoryRequirements);
}

void cgoGetImageSparseMemoryRequirements2KHR(PFN_vkGetImageSparseMemoryRequirements2KHR fn, VkDevice device, VkImageSparseMemoryRequirementsInfo2KHR *pInfo, uint32_t *pSparseMemoryRequirementsCount, VkSparseImageMemoryRequirements2KHR *pSparseMemoryRequirements) {
	fn(device, pInfo, pSparseMemoryRequirementsCount, pSparseMemoryRequirements);
}
*/
import "C"
import (
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3/loader"
	_ "github.com/vkngwrapper/extensions/v3/vulkan"
)

//go:generate mockgen -source loader.go -destination ../mocks/loader.go -package mock_get_memory_requirements2

type Loader interface {
	VkGetBufferMemoryRequirements2KHR(device loader.VkDevice, pInfo *VkBufferMemoryRequirementsInfo2KHR, pMemoryRequirements *VkMemoryRequirements2KHR)
	VkGetImageMemoryRequirements2KHR(device loader.VkDevice, pInfo *VkImageMemoryRequirementsInfo2KHR, pMemoryRequirements *VkMemoryRequirements2KHR)
	VkGetImageSparseMemoryRequirements2KHR(device loader.VkDevice, pInfo *VkImageSparseMemoryRequirementsInfo2KHR, pSparseMemoryRequirementCount *loader.Uint32, pSparseMemoryRequirements *VkSparseImageMemoryRequirements2KHR)
}

type VkBufferMemoryRequirementsInfo2KHR C.VkBufferMemoryRequirementsInfo2KHR
type VkImageMemoryRequirementsInfo2KHR C.VkImageMemoryRequirementsInfo2KHR
type VkImageSparseMemoryRequirementsInfo2KHR C.VkImageSparseMemoryRequirementsInfo2KHR
type VkMemoryRequirements2KHR C.VkMemoryRequirements2KHR
type VkSparseImageMemoryRequirements2KHR C.VkSparseImageMemoryRequirements2KHR

type CLoader struct {
	coreLoader loader.Loader

	getBufferMemoryRequirements2      C.PFN_vkGetBufferMemoryRequirements2KHR
	getImageMemoryRequirements2       C.PFN_vkGetImageMemoryRequirements2KHR
	getImageSparseMemoryRequirements2 C.PFN_vkGetImageSparseMemoryRequirements2KHR
}

func CreateLoaderFromCore(coreLoader loader.Loader) *CLoader {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	return &CLoader{
		coreLoader: coreLoader,

		getBufferMemoryRequirements2:      (C.PFN_vkGetBufferMemoryRequirements2KHR)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkGetBufferMemoryRequirements2KHR")))),
		getImageMemoryRequirements2:       (C.PFN_vkGetImageMemoryRequirements2KHR)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkGetImageMemoryRequirements2KHR")))),
		getImageSparseMemoryRequirements2: (C.PFN_vkGetImageSparseMemoryRequirements2KHR)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkGetImageSparseMemoryRequirements2KHR")))),
	}
}

func (d *CLoader) VkGetBufferMemoryRequirements2KHR(device loader.VkDevice, pInfo *VkBufferMemoryRequirementsInfo2KHR, pMemoryRequirements *VkMemoryRequirements2KHR) {
	if d.getBufferMemoryRequirements2 == nil {
		panic("attempt to call extension method vkGetBufferMemoryRequirements2KHR when extension not present")
	}

	C.cgoGetBufferMemoryRequirements2KHR(d.getBufferMemoryRequirements2,
		C.VkDevice(unsafe.Pointer(device)),
		(*C.VkBufferMemoryRequirementsInfo2KHR)(pInfo),
		(*C.VkMemoryRequirements2KHR)(pMemoryRequirements))
}

func (d *CLoader) VkGetImageMemoryRequirements2KHR(device loader.VkDevice, pInfo *VkImageMemoryRequirementsInfo2KHR, pMemoryRequirements *VkMemoryRequirements2KHR) {
	if d.getImageMemoryRequirements2 == nil {
		panic("attempt to call extension method vkGetImageMemoryRequirements2KHR when extension not present")
	}

	C.cgoGetImageMemoryRequirements2KHR(d.getImageMemoryRequirements2,
		C.VkDevice(unsafe.Pointer(device)),
		(*C.VkImageMemoryRequirementsInfo2KHR)(pInfo),
		(*C.VkMemoryRequirements2KHR)(pMemoryRequirements))
}

func (d *CLoader) VkGetImageSparseMemoryRequirements2KHR(device loader.VkDevice, pInfo *VkImageSparseMemoryRequirementsInfo2KHR, pSparseMemoryRequirementCount *loader.Uint32, pSparseMemoryRequirements *VkSparseImageMemoryRequirements2KHR) {
	if d.getImageSparseMemoryRequirements2 == nil {
		panic("attempt to call extension method vkGetImageSparseMemoryRequirements2KHR when extension not present")
	}

	C.cgoGetImageSparseMemoryRequirements2KHR(d.getImageSparseMemoryRequirements2,
		C.VkDevice(unsafe.Pointer(device)),
		(*C.VkImageSparseMemoryRequirementsInfo2KHR)(pInfo),
		(*C.uint32_t)(pSparseMemoryRequirementCount),
		(*C.VkSparseImageMemoryRequirements2KHR)(pSparseMemoryRequirements),
	)
}
