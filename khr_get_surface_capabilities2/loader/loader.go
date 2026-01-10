package khr_get_surface_capabilities2_loader

/*
#include <stdlib.h>
#include "../../vulkan/vulkan.h"

VkResult cgoGetPhysicalDeviceSurfaceCapabilities2KHR(PFN_vkGetPhysicalDeviceSurfaceCapabilities2KHR fn, VkPhysicalDevice physicalDevice, VkPhysicalDeviceSurfaceInfo2KHR *pSurfaceInfo, VkSurfaceCapabilities2KHR *pSurfaceCapabilities) {
	return fn(physicalDevice, pSurfaceInfo, pSurfaceCapabilities);
}

VkResult cgoGetPhysicalDeviceSurfaceFormats2KHR(PFN_vkGetPhysicalDeviceSurfaceFormats2KHR fn, VkPhysicalDevice physicalDevice, VkPhysicalDeviceSurfaceInfo2KHR *pSurfaceInfo, uint32_t *pSurfaceFormatCount, VkSurfaceFormat2KHR *pSurfaceFormats) {
	return fn(physicalDevice, pSurfaceInfo, pSurfaceFormatCount, pSurfaceFormats);
}
*/
import "C"
import (
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/loader"
)

//go:generate mockgen -source loader.go -destination ../mocks/loader.go -package mock_get_surface_capabilities2

type Loader interface {
	VkGetPhysicalDeviceSurfaceCapabilities2KHR(physicalDevice loader.VkPhysicalDevice, pSurfaceInfo *VkPhysicalDeviceSurfaceInfo2KHR, pSurfaceCapabilities *VkSurfaceCapabilities2KHR) (common.VkResult, error)
	VkGetPhysicalDeviceSurfaceFormats2KHR(physicalDevice loader.VkPhysicalDevice, pSurfaceInfo *VkPhysicalDeviceSurfaceInfo2KHR, pSurfaceFormatCount *loader.Uint32, pSurfaceFormats *VkSurfaceFormat2KHR) (common.VkResult, error)
}

type VkPhysicalDeviceSurfaceInfo2KHR C.VkPhysicalDeviceSurfaceInfo2KHR
type VkSurfaceCapabilities2KHR C.VkSurfaceCapabilities2KHR
type VkSurfaceFormat2KHR C.VkSurfaceFormat2KHR

type CLoader struct {
	coreLoader loader.Loader

	getPhysicalDeviceSurfaceCapabilities2 C.PFN_vkGetPhysicalDeviceSurfaceCapabilities2KHR
	getPhysicalDeviceSurfaceFormats2      C.PFN_vkGetPhysicalDeviceSurfaceFormats2KHR
}

func CreateLoaderFromCore(coreLoader loader.Loader) *CLoader {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	return &CLoader{
		coreLoader: coreLoader,

		getPhysicalDeviceSurfaceCapabilities2: (C.PFN_vkGetPhysicalDeviceSurfaceCapabilities2KHR)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkGetPhysicalDeviceSurfaceCapabilities2KHR")))),
		getPhysicalDeviceSurfaceFormats2:      (C.PFN_vkGetPhysicalDeviceSurfaceFormats2KHR)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkGetPhysicalDeviceSurfaceFormats2KHR")))),
	}
}

func (d *CLoader) VkGetPhysicalDeviceSurfaceCapabilities2KHR(physicalDevice loader.VkPhysicalDevice, pSurfaceInfo *VkPhysicalDeviceSurfaceInfo2KHR, pSurfaceCapabilities *VkSurfaceCapabilities2KHR) (common.VkResult, error) {
	if d.getPhysicalDeviceSurfaceCapabilities2 == nil {
		panic("attempt to call extension method vkGetPhysicalDeviceSurfaceCapabilities2KHR when extension not present")
	}

	res := common.VkResult(C.cgoGetPhysicalDeviceSurfaceCapabilities2KHR(d.getPhysicalDeviceSurfaceCapabilities2,
		C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)),
		(*C.VkPhysicalDeviceSurfaceInfo2KHR)(pSurfaceInfo),
		(*C.VkSurfaceCapabilities2KHR)(pSurfaceCapabilities),
	))

	return res, res.ToError()
}

func (d *CLoader) VkGetPhysicalDeviceSurfaceFormats2KHR(physicalDevice loader.VkPhysicalDevice, pSurfaceInfo *VkPhysicalDeviceSurfaceInfo2KHR, pSurfaceFormatCount *loader.Uint32, pSurfaceFormats *VkSurfaceFormat2KHR) (common.VkResult, error) {
	if d.getPhysicalDeviceSurfaceFormats2 == nil {
		panic("attempt to call extension method vkGetPhysicalDeviceSurfaceFormats2KHR when extension not present")
	}

	res := common.VkResult(C.cgoGetPhysicalDeviceSurfaceFormats2KHR(d.getPhysicalDeviceSurfaceFormats2,
		C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)),
		(*C.VkPhysicalDeviceSurfaceInfo2KHR)(pSurfaceInfo),
		(*C.uint32_t)(pSurfaceFormatCount),
		(*C.VkSurfaceFormat2KHR)(pSurfaceFormats),
	))
	return res, res.ToError()
}
