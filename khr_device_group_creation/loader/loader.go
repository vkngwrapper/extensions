package khr_device_group_creation_loader

/*
#include <stdlib.h>
#include "../../vulkan/vulkan.h"

VkResult cgoEnumeratePhysicalDeviceGroupsKHR(PFN_vkEnumeratePhysicalDeviceGroupsKHR fn, VkInstance instance, uint32_t *pPhysicalDeviceGroupCount, VkPhysicalDeviceGroupPropertiesKHR *pPhysicalDeviceGroupProperties) {
	return fn(instance, pPhysicalDeviceGroupCount, pPhysicalDeviceGroupProperties);
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

//go:generate mockgen -source loader.go -destination ../mocks/loader.go -package mock_device_group_creation

type Loader interface {
	VkEnumeratePhysicalDeviceGroupsKHR(instance loader.VkInstance, pPhysicalDeviceGroupCount *loader.Uint32, pPhysicalDeviceGroupProperties *VkPhysicalDeviceGroupPropertiesKHR) (common.VkResult, error)
}

type VkPhysicalDeviceGroupPropertiesKHR C.VkPhysicalDeviceGroupPropertiesKHR
type VkDeviceGroupDeviceCreateInfoKHR C.VkDeviceGroupDeviceCreateInfoKHR

type CLoader struct {
	coreLoader loader.Loader

	enumeratePhysicalDeviceGroups C.PFN_vkEnumeratePhysicalDeviceGroupsKHR
}

func CreateLoaderFromCore(coreLoader loader.Loader) *CLoader {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	return &CLoader{
		coreLoader: coreLoader,

		enumeratePhysicalDeviceGroups: (C.PFN_vkEnumeratePhysicalDeviceGroupsKHR)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkEnumeratePhysicalDeviceGroupsKHR")))),
	}
}

func (d *CLoader) VkEnumeratePhysicalDeviceGroupsKHR(instance loader.VkInstance, pPhysicalDeviceGroupCount *loader.Uint32, pPhysicalDeviceGroupProperties *VkPhysicalDeviceGroupPropertiesKHR) (common.VkResult, error) {
	if d.enumeratePhysicalDeviceGroups == nil {
		panic("attempt to call extension method vkEnumeratePhysicalDeviceGroupsKHR when extension not present")
	}

	res := common.VkResult(C.cgoEnumeratePhysicalDeviceGroupsKHR(
		d.enumeratePhysicalDeviceGroups,
		(C.VkInstance)(unsafe.Pointer(instance)),
		(*C.uint32_t)(pPhysicalDeviceGroupCount),
		(*C.VkPhysicalDeviceGroupPropertiesKHR)(pPhysicalDeviceGroupProperties),
	))

	return res, res.ToError()
}
