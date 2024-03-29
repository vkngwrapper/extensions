package khr_maintenance3_driver

//go:generate mockgen -source driver.go -destination ../dummies/driver.go -package mock_maintenance3

/*
#include <stdlib.h>
#include "../../vulkan/vulkan.h"

void cgoGetDescriptorSetLayoutSupportKHR(PFN_vkGetDescriptorSetLayoutSupportKHR fn, VkDevice device, VkDescriptorSetLayoutCreateInfo *pCreateInfo, VkDescriptorSetLayoutSupportKHR *pSupport) {
	fn(device, pCreateInfo, pSupport);
}
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v2/driver"
	_ "github.com/vkngwrapper/extensions/v2/vulkan"
	"unsafe"
)

type Driver interface {
	VkGetDescriptorSetLayoutSupportKHR(device driver.VkDevice, pCreateInfo *driver.VkDescriptorSetLayoutCreateInfo, pSupport *VkDescriptorSetLayoutSupportKHR)
}

type VkDescriptorSetLayoutSupportKHR C.VkDescriptorSetLayoutSupportKHR
type VkPhysicalDeviceMaintenance3PropertiesKHR C.VkPhysicalDeviceMaintenance3PropertiesKHR

type CDriver struct {
	coreDriver                    driver.Driver
	getDescriptorSetLayoutSupport C.PFN_vkGetDescriptorSetLayoutSupportKHR
}

func CreateDriverFromCore(coreDriver driver.Driver) *CDriver {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	return &CDriver{
		coreDriver:                    coreDriver,
		getDescriptorSetLayoutSupport: (C.PFN_vkGetDescriptorSetLayoutSupportKHR)(coreDriver.LoadProcAddr((*driver.Char)(arena.CString("vkGetDescriptorSetLayoutSupportKHR")))),
	}
}

func (d *CDriver) VkGetDescriptorSetLayoutSupportKHR(device driver.VkDevice, pCreateInfo *driver.VkDescriptorSetLayoutCreateInfo, pSupport *VkDescriptorSetLayoutSupportKHR) {
	if d.getDescriptorSetLayoutSupport == nil {
		panic("attempt to call extension method vkGetDescriptorSetLayoutSupportKHR when extension not present")
	}

	C.cgoGetDescriptorSetLayoutSupportKHR(d.getDescriptorSetLayoutSupport,
		(C.VkDevice)(unsafe.Pointer(device)),
		(*C.VkDescriptorSetLayoutCreateInfo)(unsafe.Pointer(pCreateInfo)),
		(*C.VkDescriptorSetLayoutSupportKHR)(pSupport),
	)
}
