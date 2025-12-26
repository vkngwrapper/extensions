package khr_external_memory_capabilities_loader

/*
#include <stdlib.h>
#include "../../vulkan/vulkan.h"

void cgoGetPhysicalDeviceExternalBufferPropertiesKHR(PFN_vkGetPhysicalDeviceExternalBufferPropertiesKHR fn, VkPhysicalDevice physicalDevice, VkPhysicalDeviceExternalBufferInfoKHR *pExternalBufferInfo, VkExternalBufferPropertiesKHR *pExternalBufferProperties) {
	fn(physicalDevice, pExternalBufferInfo, pExternalBufferProperties);
}
*/
import "C"
import (
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3/loader"
	_ "github.com/vkngwrapper/extensions/v3/vulkan"
)

//go:generate mockgen -source loader.go -destination ../mocks/loader.go -package mock_external_memory_capabilities

type Loader interface {
	VkGetPhysicalDeviceExternalBufferPropertiesKHR(physicalDevice loader.VkPhysicalDevice, pExternalBufferInfo *VkPhysicalDeviceExternalBufferInfoKHR, pExternalBufferProperties *VkExternalBufferPropertiesKHR)
}

type VkPhysicalDeviceExternalBufferInfoKHR C.VkPhysicalDeviceExternalBufferInfoKHR
type VkExternalBufferPropertiesKHR C.VkExternalBufferPropertiesKHR
type VkExternalMemoryPropertiesKHR C.VkExternalMemoryPropertiesKHR
type VkExternalImageFormatPropertiesKHR C.VkExternalImageFormatPropertiesKHR
type VkPhysicalDeviceExternalImageFormatInfoKHR C.VkPhysicalDeviceExternalImageFormatInfoKHR
type VkPhysicalDeviceIDPropertiesKHR C.VkPhysicalDeviceIDPropertiesKHR

type CLoader struct {
	coreLoader loader.Loader

	getPhysicalDeviceExternalBufferProperties C.PFN_vkGetPhysicalDeviceExternalBufferPropertiesKHR
}

func CreateLoaderFromCore(coreLoader loader.Loader) *CLoader {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	return &CLoader{
		coreLoader: coreLoader,

		getPhysicalDeviceExternalBufferProperties: (C.PFN_vkGetPhysicalDeviceExternalBufferPropertiesKHR)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkGetPhysicalDeviceExternalBufferPropertiesKHR")))),
	}
}

func (d *CLoader) VkGetPhysicalDeviceExternalBufferPropertiesKHR(physicalDevice loader.VkPhysicalDevice, pExternalBufferInfo *VkPhysicalDeviceExternalBufferInfoKHR, pExternalBufferProperties *VkExternalBufferPropertiesKHR) {
	if d.getPhysicalDeviceExternalBufferProperties == nil {
		panic("attempt to call extension method vkGetPhysicalDeviceExternalBufferPropertiesKHR when extension not present")
	}

	C.cgoGetPhysicalDeviceExternalBufferPropertiesKHR(
		d.getPhysicalDeviceExternalBufferProperties,
		C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)),
		(*C.VkPhysicalDeviceExternalBufferInfoKHR)(pExternalBufferInfo),
		(*C.VkExternalBufferPropertiesKHR)(pExternalBufferProperties),
	)
}
