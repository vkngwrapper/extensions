package khr_external_semaphore_capabilities_loader

/*
#include <stdlib.h>
#include "../../vulkan/vulkan.h"

void cgoGetPhysicalDeviceExternalSemaphorePropertiesKHR(PFN_vkGetPhysicalDeviceExternalSemaphorePropertiesKHR fn, VkPhysicalDevice physicalDevice, VkPhysicalDeviceExternalSemaphoreInfoKHR *pExternalSemaphoreInfo, VkExternalSemaphorePropertiesKHR *pExternalSemaphoreProperties) {
	fn(physicalDevice, pExternalSemaphoreInfo, pExternalSemaphoreProperties);
}
*/
import "C"
import (
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3/loader"
	_ "github.com/vkngwrapper/extensions/v3/vulkan"
)

//go:generate mockgen -source loader.go -destination ../mocks/loader.go -package mock_external_semaphore_capabilities

type Loader interface {
	VkGetPhysicalDeviceExternalSemaphorePropertiesKHR(physicalDevice loader.VkPhysicalDevice, pExternalSemaphoreInfo *VkPhysicalDeviceExternalSemaphoreInfoKHR, pExternalSemaphoreProperties *VkExternalSemaphorePropertiesKHR)
}

type VkPhysicalDeviceExternalSemaphoreInfoKHR C.VkPhysicalDeviceExternalSemaphoreInfoKHR
type VkExternalSemaphorePropertiesKHR C.VkExternalSemaphorePropertiesKHR
type VkPhysicalDeviceIDPropertiesKHR C.VkPhysicalDeviceIDPropertiesKHR

type CLoader struct {
	coreLoader loader.Loader

	getPhysicalDeviceExternalSemaphoreProperties C.PFN_vkGetPhysicalDeviceExternalSemaphorePropertiesKHR
}

func CreateLoaderFromCore(coreLoader loader.Loader) *CLoader {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	return &CLoader{
		coreLoader: coreLoader,

		getPhysicalDeviceExternalSemaphoreProperties: (C.PFN_vkGetPhysicalDeviceExternalSemaphorePropertiesKHR)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkGetPhysicalDeviceExternalSemaphorePropertiesKHR")))),
	}
}

func (d *CLoader) VkGetPhysicalDeviceExternalSemaphorePropertiesKHR(physicalDevice loader.VkPhysicalDevice, pExternalSemaphoreInfo *VkPhysicalDeviceExternalSemaphoreInfoKHR, pExternalSemaphoreProperties *VkExternalSemaphorePropertiesKHR) {
	if d.getPhysicalDeviceExternalSemaphoreProperties == nil {
		panic("attempt to call extension method vkGetPhysicalDeviceExternalSemaphorePropertiesKHR when extension not present")
	}

	C.cgoGetPhysicalDeviceExternalSemaphorePropertiesKHR(
		d.getPhysicalDeviceExternalSemaphoreProperties,
		C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)),
		(*C.VkPhysicalDeviceExternalSemaphoreInfoKHR)(pExternalSemaphoreInfo),
		(*C.VkExternalSemaphorePropertiesKHR)(pExternalSemaphoreProperties),
	)
}
