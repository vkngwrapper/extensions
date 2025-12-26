package khr_sampler_ycbcr_conversion_loader

/*
#include <stdlib.h>
#include "../../vulkan/vulkan.h"

VkResult cgoCreateSamplerYcbcrConversionKHR(PFN_vkCreateSamplerYcbcrConversionKHR fn, VkDevice device, VkSamplerYcbcrConversionCreateInfoKHR *pCreateInfo, VkAllocationCallbacks *pAllocator, VkSamplerYcbcrConversionKHR *pYcbcrConversion) {
	return fn(device, pCreateInfo, pAllocator, pYcbcrConversion);
}

void cgoDestroySamplerYcbcrConversionKHR(PFN_vkDestroySamplerYcbcrConversionKHR fn, VkDevice device, VkSamplerYcbcrConversionKHR ycbcrConversion, VkAllocationCallbacks *pAllocator) {
	fn(device, ycbcrConversion, pAllocator);
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

//go:generate mockgen -source loader.go -destination ../mocks/loader.go -package mock_sampler_ycbcr_conversion

type Loader interface {
	VkCreateSamplerYcbcrConversionKHR(device loader.VkDevice, pCreateInfo *VkSamplerYcbcrConversionCreateInfoKHR, pAllocator *loader.VkAllocationCallbacks, pYcbcrConversion *VkSamplerYcbcrConversionKHR) (common.VkResult, error)
	VkDestroySamplerYcbcrConversionKHR(device loader.VkDevice, ycbcrConversion VkSamplerYcbcrConversionKHR, pAllocator *loader.VkAllocationCallbacks)
}

type VkSamplerYcbcrConversionKHR loader.VulkanHandle
type VkSamplerYcbcrConversionCreateInfoKHR C.VkSamplerYcbcrConversionCreateInfoKHR
type VkBindImagePlaneMemoryInfoKHR C.VkBindImagePlaneMemoryInfoKHR
type VkSamplerYcbcrConversionImageFormatPropertiesKHR C.VkSamplerYcbcrConversionImageFormatPropertiesKHR
type VkImagePlaneMemoryRequirementsInfoKHR C.VkImagePlaneMemoryRequirementsInfoKHR
type VkPhysicalDeviceSamplerYcbcrConversionFeaturesKHR C.VkPhysicalDeviceSamplerYcbcrConversionFeaturesKHR
type VkSamplerYcbcrConversionInfoKHR C.VkSamplerYcbcrConversionInfoKHR

type CLoader struct {
	coreLoader loader.Loader

	createSamplerYcbcrConversion  C.PFN_vkCreateSamplerYcbcrConversionKHR
	destroySamplerYcbcrConversion C.PFN_vkDestroySamplerYcbcrConversionKHR
}

func CreateLoaderFromCore(coreLoader loader.Loader) *CLoader {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	return &CLoader{
		coreLoader: coreLoader,

		createSamplerYcbcrConversion:  (C.PFN_vkCreateSamplerYcbcrConversionKHR)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkCreateSamplerYcbcrConversionKHR")))),
		destroySamplerYcbcrConversion: (C.PFN_vkDestroySamplerYcbcrConversionKHR)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkDestroySamplerYcbcrConversionKHR")))),
	}
}

func (d *CLoader) VkCreateSamplerYcbcrConversionKHR(device loader.VkDevice, pCreateInfo *VkSamplerYcbcrConversionCreateInfoKHR, pAllocator *loader.VkAllocationCallbacks, pYcbcrConversion *VkSamplerYcbcrConversionKHR) (common.VkResult, error) {
	if d.createSamplerYcbcrConversion == nil {
		panic("attempt to call extension method vkCreateSamplerYcbcrConversionKHR when extension not present")
	}

	res := common.VkResult(C.cgoCreateSamplerYcbcrConversionKHR(
		d.createSamplerYcbcrConversion,
		C.VkDevice(unsafe.Pointer(device)),
		(*C.VkSamplerYcbcrConversionCreateInfoKHR)(pCreateInfo),
		(*C.VkAllocationCallbacks)(unsafe.Pointer(pAllocator)),
		(*C.VkSamplerYcbcrConversionKHR)(unsafe.Pointer(pYcbcrConversion)),
	))
	return res, res.ToError()
}

func (d *CLoader) VkDestroySamplerYcbcrConversionKHR(device loader.VkDevice, ycbcrConversion VkSamplerYcbcrConversionKHR, pAllocator *loader.VkAllocationCallbacks) {
	if d.destroySamplerYcbcrConversion == nil {
		panic("attempt to call extension method vkDestroySamplerYcbcrConversionKHR when extension not present")
	}

	C.cgoDestroySamplerYcbcrConversionKHR(
		d.destroySamplerYcbcrConversion,
		C.VkDevice(unsafe.Pointer(device)),
		C.VkSamplerYcbcrConversionKHR(unsafe.Pointer(ycbcrConversion)),
		(*C.VkAllocationCallbacks)(unsafe.Pointer(pAllocator)),
	)
}
