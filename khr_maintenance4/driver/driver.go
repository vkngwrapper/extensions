package khr_maintenance4_driver

//go:generate mockgen -source driver.go -destination ../dummies/driver.go -package mock_maintenance4

/*
#include <stdlib.h>
#include "../../vulkan/vulkan.h"

void cgoGetDeviceBufferMemoryRequirmentsKHR(PFN_vkGetDeviceBufferMemoryRequirementsKHR fn, VkDevice device, VkDeviceBufferMemoryRequirementsKHR *pInfo, VkMemoryRequirements2 *pMemoryRequirements) {
	fn(device, pInfo, pMemoryRequirements);
}

void cgoGetDeviceImageMemoryRequirementsKHR(PFN_vkGetDeviceImageMemoryRequirementsKHR fn, VkDevice device, VkDeviceImageMemoryRequirementsKHR *pInfo, VkMemoryRequirements2 *pMemoryRequirements) {
	fn(device, pInfo, pMemoryRequirements);
}

void cgoGetDeviceImageSparseMemoryRequirementsKHR(PFN_vkGetDeviceImageSparseMemoryRequirementsKHR fn, VkDevice device, VkDeviceImageMemoryRequirementsKHR *pInfo, uint32_t *pSparseMemoryRequirementCount, VkSparseImageMemoryRequirements2 *pSparseMemoryRequirements) {
	fn(device, pInfo, pSparseMemoryRequirementCount, pSparseMemoryRequirements);
}

*/
import "C"
import (
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3/driver"
)

type Driver interface {
	VkGetDeviceBufferMemoryRequirementsKHR(device driver.VkDevice, pInfo *VkDeviceBufferMemoryRequirementsKHR, pMemoryRequirements *driver.VkMemoryRequirements2)
	VkGetDeviceImageMemoryRequirementsKHR(device driver.VkDevice, pInfo *VkDeviceImageMemoryRequirementsKHR, pMemoryRequirements *driver.VkMemoryRequirements2)
	VkGetDeviceImageSparseMemoryRequirementsKHR(device driver.VkDevice, pInfo *VkDeviceImageMemoryRequirementsKHR, pSparseMemoryRequirementCount *driver.Uint32, pSparseMemoryRequirements *driver.VkSparseImageMemoryRequirements2)
}

type VkDeviceImageMemoryRequirementsKHR C.VkDeviceImageMemoryRequirementsKHR
type VkDeviceBufferMemoryRequirementsKHR C.VkDeviceBufferMemoryRequirementsKHR
type VkPhysicalDeviceMaintenance4FeaturesKHR C.VkPhysicalDeviceMaintenance4FeaturesKHR
type VkPhysicalDeviceMaintenance4PropertiesKHR C.VkPhysicalDeviceMaintenance4PropertiesKHR

type CDriver struct {
	coreDriver                                driver.Driver
	getDeviceBufferMemoryRequirementsKHR      C.PFN_vkGetDeviceBufferMemoryRequirementsKHR
	getDeviceImageMemoryRequirementsKHR       C.PFN_vkGetDeviceImageMemoryRequirementsKHR
	getDeviceImageSparseMemoryRequirementsKHR C.PFN_vkGetDeviceImageSparseMemoryRequirementsKHR
}

func CreateDriverFromCore(coreDriver driver.Driver) *CDriver {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	return &CDriver{
		coreDriver:                                coreDriver,
		getDeviceBufferMemoryRequirementsKHR:      (C.PFN_vkGetDeviceBufferMemoryRequirementsKHR)(coreDriver.LoadProcAddr((*driver.Char)(arena.CString("vkGetDeviceBufferMemoryRequirementsKHR")))),
		getDeviceImageMemoryRequirementsKHR:       (C.PFN_vkGetDeviceImageMemoryRequirementsKHR)(coreDriver.LoadProcAddr((*driver.Char)(arena.CString("vkGetDeviceImageMemoryRequirementsKHR")))),
		getDeviceImageSparseMemoryRequirementsKHR: (C.PFN_vkGetDeviceImageSparseMemoryRequirementsKHR)(coreDriver.LoadProcAddr((*driver.Char)(arena.CString("vkGetDeviceImageSparseMemoryRequirementsKHR")))),
	}
}

func (d *CDriver) VkGetDeviceBufferMemoryRequirementsKHR(device driver.VkDevice, pInfo *VkDeviceBufferMemoryRequirementsKHR, pMemoryRequirements *driver.VkMemoryRequirements2) {
	if d.getDeviceBufferMemoryRequirementsKHR == nil {
		panic("attempt to call extension method vkGetDeviceBufferMemoryRequirementsKHR when extension not present")
	}

	C.cgoGetDeviceBufferMemoryRequirmentsKHR(
		d.getDeviceBufferMemoryRequirementsKHR,
		(C.VkDevice)(unsafe.Pointer(device)),
		(*C.VkDeviceBufferMemoryRequirementsKHR)(pInfo),
		(*C.VkMemoryRequirements2)(unsafe.Pointer(pMemoryRequirements)),
	)
}

func (d *CDriver) VkGetDeviceImageMemoryRequirementsKHR(device driver.VkDevice, pInfo *VkDeviceImageMemoryRequirementsKHR, pMemoryRequirements *driver.VkMemoryRequirements2) {
	if d.getDeviceImageMemoryRequirementsKHR == nil {
		panic("attempt to call extension method vkGetDeviceImageMemoryRequirementsKHR when extension not present")
	}

	C.cgoGetDeviceImageMemoryRequirementsKHR(
		d.getDeviceImageMemoryRequirementsKHR,
		(C.VkDevice)(unsafe.Pointer(device)),
		(*C.VkDeviceImageMemoryRequirementsKHR)(pInfo),
		(*C.VkMemoryRequirements2)(unsafe.Pointer(pMemoryRequirements)),
	)
}

func (d *CDriver) VkGetDeviceImageSparseMemoryRequirementsKHR(device driver.VkDevice, pInfo *VkDeviceImageMemoryRequirementsKHR, pSparseMemoryRequirementCount *driver.Uint32, pSparseMemoryRequirements *driver.VkSparseImageMemoryRequirements2) {
	if d.getDeviceImageSparseMemoryRequirementsKHR == nil {
		panic("attempt to call extension method vkGetDeviceImageSparseMemoryRequirementsKHR when extension not present")
	}

	C.cgoGetDeviceImageSparseMemoryRequirementsKHR(
		d.getDeviceImageSparseMemoryRequirementsKHR,
		(C.VkDevice)(unsafe.Pointer(device)),
		(*C.VkDeviceImageMemoryRequirementsKHR)(pInfo),
		(*C.uint32_t)(pSparseMemoryRequirementCount),
		(*C.VkSparseImageMemoryRequirements2)(unsafe.Pointer(pSparseMemoryRequirements)),
	)
}

var _ Driver = &CDriver{}
