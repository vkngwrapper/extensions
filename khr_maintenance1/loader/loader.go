package khr_maintenance1_loader

//go:generate mockgen -source loader.go -destination ../mocks/loader.go -package mock_maintenance1

/*
#include <stdlib.h>
#include "../../vulkan/vulkan.h"

void cgoTrimCommandPoolKHR(PFN_vkTrimCommandPoolKHR fn, VkDevice device, VkCommandPool commandPool, VkCommandPoolTrimFlagsKHR flags) {
	fn(device, commandPool, flags);
}
*/
import "C"
import (
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3/loader"
	_ "github.com/vkngwrapper/extensions/v3/vulkan"
)

type VkCommandPoolTrimFlagsKHR C.VkCommandPoolTrimFlagsKHR
type Loader interface {
	VkTrimCommandPoolKHR(device loader.VkDevice, commandPool loader.VkCommandPool, flags VkCommandPoolTrimFlagsKHR)
}

type CLoader struct {
	driver loader.Loader

	trimCommandPool C.PFN_vkTrimCommandPoolKHR
}

func CreateLoaderFromCore(coreLoader loader.Loader) *CLoader {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	return &CLoader{
		driver: coreLoader,

		trimCommandPool: (C.PFN_vkTrimCommandPoolKHR)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkTrimCommandPoolKHR")))),
	}
}

func (d *CLoader) VkTrimCommandPoolKHR(device loader.VkDevice, commandPool loader.VkCommandPool, flags VkCommandPoolTrimFlagsKHR) {
	if d.trimCommandPool == nil {
		panic("attempt to call extension method vkTrimCommandPoolKHR when extension not present")
	}

	C.cgoTrimCommandPoolKHR(d.trimCommandPool,
		C.VkDevice(unsafe.Pointer(device)),
		C.VkCommandPool(unsafe.Pointer(commandPool)),
		C.VkCommandPoolTrimFlags(flags))
}
