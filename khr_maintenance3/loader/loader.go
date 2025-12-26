package khr_maintenance3_loader

//go:generate mockgen -source loader.go -destination ../mocks/loader.go -package mock_maintenance3

/*
#include <stdlib.h>
#include "../../vulkan/vulkan.h"

void cgoGetDescriptorSetLayoutSupportKHR(PFN_vkGetDescriptorSetLayoutSupportKHR fn, VkDevice device, VkDescriptorSetLayoutCreateInfo *pCreateInfo, VkDescriptorSetLayoutSupportKHR *pSupport) {
	fn(device, pCreateInfo, pSupport);
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
	VkGetDescriptorSetLayoutSupportKHR(device loader.VkDevice, pCreateInfo *loader.VkDescriptorSetLayoutCreateInfo, pSupport *VkDescriptorSetLayoutSupportKHR)
}

type VkDescriptorSetLayoutSupportKHR C.VkDescriptorSetLayoutSupportKHR
type VkPhysicalDeviceMaintenance3PropertiesKHR C.VkPhysicalDeviceMaintenance3PropertiesKHR

type CLoader struct {
	coreLoader                    loader.Loader
	getDescriptorSetLayoutSupport C.PFN_vkGetDescriptorSetLayoutSupportKHR
}

func CreateLoaderFromCore(coreLoader loader.Loader) *CLoader {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	return &CLoader{
		coreLoader:                    coreLoader,
		getDescriptorSetLayoutSupport: (C.PFN_vkGetDescriptorSetLayoutSupportKHR)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkGetDescriptorSetLayoutSupportKHR")))),
	}
}

func (d *CLoader) VkGetDescriptorSetLayoutSupportKHR(device loader.VkDevice, pCreateInfo *loader.VkDescriptorSetLayoutCreateInfo, pSupport *VkDescriptorSetLayoutSupportKHR) {
	if d.getDescriptorSetLayoutSupport == nil {
		panic("attempt to call extension method vkGetDescriptorSetLayoutSupportKHR when extension not present")
	}

	C.cgoGetDescriptorSetLayoutSupportKHR(d.getDescriptorSetLayoutSupport,
		(C.VkDevice)(unsafe.Pointer(device)),
		(*C.VkDescriptorSetLayoutCreateInfo)(unsafe.Pointer(pCreateInfo)),
		(*C.VkDescriptorSetLayoutSupportKHR)(pSupport),
	)
}
