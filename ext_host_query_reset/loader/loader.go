package ext_host_query_reset_loader

//go:generate mockgen -source loader.go -destination ../mocks/loader.go -package mock_host_query_reset

/*
#include <stdlib.h>
#include "../../vulkan/vulkan.h"

void cgoResetQueryPoolEXT(PFN_vkResetQueryPoolEXT fn, VkDevice device, VkQueryPool queryPool, uint32_t firstQuery, uint32_t queryCount) {
	fn(device, queryPool, firstQuery, queryCount);
}
*/
import "C"
import (
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3/loader"
	_ "github.com/vkngwrapper/extensions/v3/vulkan"
)

type Loader interface {
	VkResetQueryPoolEXT(device loader.VkDevice, queryPool loader.VkQueryPool, firstQuery loader.Uint32, queryCount loader.Uint32)
}

type VkPhysicalDeviceHostQueryResetFeaturesEXT C.VkPhysicalDeviceHostQueryResetFeaturesEXT

type CLoader struct {
	coreLoader loader.Loader

	resetQueryPool C.PFN_vkResetQueryPoolEXT
}

func CreateLoaderFromCore(coreLoader loader.Loader) *CLoader {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	return &CLoader{
		coreLoader:     coreLoader,
		resetQueryPool: (C.PFN_vkResetQueryPoolEXT)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkResetQueryPoolEXT")))),
	}
}

func (d *CLoader) VkResetQueryPoolEXT(device loader.VkDevice, queryPool loader.VkQueryPool, firstQuery loader.Uint32, queryCount loader.Uint32) {
	C.cgoResetQueryPoolEXT(
		d.resetQueryPool,
		C.VkDevice(unsafe.Pointer(device)),
		C.VkQueryPool(unsafe.Pointer(queryPool)),
		C.uint32_t(firstQuery),
		C.uint32_t(queryCount),
	)
}
