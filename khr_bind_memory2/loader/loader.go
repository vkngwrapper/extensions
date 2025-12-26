package khr_bind_memory2_loader

//go:generate mockgen -source loader.go -destination ../mocks/loader.go -package mock_bind_memory2

/*
#include <stdlib.h>
#include "../../vulkan/vulkan.h"

VkResult cgoBindBufferMemory2KHR(PFN_vkBindBufferMemory2KHR fn, VkDevice device, uint32_t bindInfoCount, VkBindBufferMemoryInfoKHR *pBindInfos) {
	return fn(device, bindInfoCount, pBindInfos);
}

VkResult cgoBindImageMemory2KHR(PFN_vkBindImageMemory2KHR fn, VkDevice device, uint32_t bindInfoCount, VkBindImageMemoryInfoKHR *pBindInfos) {
	return fn(device, bindInfoCount, pBindInfos);
}
*/
import "C"
import (
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/loader"
	_ "github.com/vkngwrapper/extensions/v3/vulkan"
)

type CLoader struct {
	driver           loader.Loader
	bindBufferMemory C.PFN_vkBindBufferMemory2KHR
	bindImageMemory  C.PFN_vkBindImageMemory2KHR
}

type VkBindBufferMemoryInfoKHR C.VkBindBufferMemoryInfoKHR
type VkBindImageMemoryInfoKHR C.VkBindImageMemoryInfoKHR
type Loader interface {
	VkBindBufferMemory2KHR(device loader.VkDevice, bindInfoCount loader.Uint32, pBindInfos *VkBindBufferMemoryInfoKHR) (common.VkResult, error)
	VkBindImageMemory2KHR(device loader.VkDevice, bindInfoCount loader.Uint32, pBindInfos *VkBindImageMemoryInfoKHR) (common.VkResult, error)
}

func CreateLoaderFromCore(coreLoader loader.Loader) *CLoader {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	return &CLoader{
		driver:           coreLoader,
		bindBufferMemory: (C.PFN_vkBindBufferMemory2KHR)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkBindBufferMemory2KHR")))),
		bindImageMemory:  (C.PFN_vkBindImageMemory2KHR)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkBindImageMemory2KHR")))),
	}
}

func (d *CLoader) VkBindBufferMemory2KHR(device loader.VkDevice, bindInfoCount loader.Uint32, pBindInfos *VkBindBufferMemoryInfoKHR) (common.VkResult, error) {
	if d.bindBufferMemory == nil {
		panic("attempt to call extension method vkBindBufferMemory2KHR when extension not present")
	}

	res := common.VkResult(C.cgoBindBufferMemory2KHR(d.bindBufferMemory,
		C.VkDevice(unsafe.Pointer(device)),
		C.uint32_t(bindInfoCount),
		(*C.VkBindBufferMemoryInfo)(pBindInfos)))

	return res, res.ToError()
}

func (d *CLoader) VkBindImageMemory2KHR(device loader.VkDevice, bindInfoCount loader.Uint32, pBindInfos *VkBindImageMemoryInfoKHR) (common.VkResult, error) {
	if d.bindImageMemory == nil {
		panic("attempt to call extension method vkBindImageMemory2KHR when extension not present")
	}

	res := common.VkResult(C.cgoBindImageMemory2KHR(d.bindImageMemory,
		C.VkDevice(unsafe.Pointer(device)),
		C.uint32_t(bindInfoCount),
		(*C.VkBindImageMemoryInfo)(pBindInfos)))

	return res, res.ToError()
}
