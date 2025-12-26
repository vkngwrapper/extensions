package khr_external_fence_capabilities_loader

/*
#include <stdlib.h>
#include "../../vulkan/vulkan.h"

void cgoGetPhysicalDeviceExternalFencePropertiesKHR(PFN_vkGetPhysicalDeviceExternalFencePropertiesKHR fn, VkPhysicalDevice physicalDevice, VkPhysicalDeviceExternalFenceInfoKHR *pExternalFenceInfo, VkExternalFencePropertiesKHR *pExternalFenceProperties) {
	fn(physicalDevice, pExternalFenceInfo, pExternalFenceProperties);
}
*/
import "C"
import (
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3/loader"
	_ "github.com/vkngwrapper/extensions/v3/vulkan"
)

//go:generate mockgen -source loader.go -destination ../mocks/loader.go -package mock_external_fence_capabilities

type Loader interface {
	VkGetPhysicalDeviceExternalFencePropertiesKHR(physicalDevice loader.VkPhysicalDevice, pExternalFenceInfo *VkPhysicalDeviceExternalFenceInfoKHR, pExternalFenceProperties *VkExternalFencePropertiesKHR)
}

type VkPhysicalDeviceExternalFenceInfoKHR C.VkPhysicalDeviceExternalFenceInfoKHR
type VkExternalFencePropertiesKHR C.VkExternalFencePropertiesKHR
type VkPhysicalDeviceIDPropertiesKHR C.VkPhysicalDeviceIDPropertiesKHR

type CLoader struct {
	coreLoader loader.Loader

	getPhysicalDeviceExternalFenceProperties C.PFN_vkGetPhysicalDeviceExternalFencePropertiesKHR
}

func CreateLoaderFromCore(coreLoader loader.Loader) *CLoader {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	return &CLoader{
		coreLoader: coreLoader,

		getPhysicalDeviceExternalFenceProperties: (C.PFN_vkGetPhysicalDeviceExternalFencePropertiesKHR)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkGetPhysicalDeviceExternalFencePropertiesKHR")))),
	}
}

func (d *CLoader) VkGetPhysicalDeviceExternalFencePropertiesKHR(physicalDevice loader.VkPhysicalDevice, pExternalFenceInfo *VkPhysicalDeviceExternalFenceInfoKHR, pExternalFenceProperties *VkExternalFencePropertiesKHR) {
	if d.getPhysicalDeviceExternalFenceProperties == nil {
		panic("attempt to call extension method vkGetPhysicalDeviceExternalFencePropertiesKHR when extension not present")
	}

	C.cgoGetPhysicalDeviceExternalFencePropertiesKHR(
		d.getPhysicalDeviceExternalFenceProperties,
		C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)),
		(*C.VkPhysicalDeviceExternalFenceInfoKHR)(pExternalFenceInfo),
		(*C.VkExternalFencePropertiesKHR)(pExternalFenceProperties),
	)
}
