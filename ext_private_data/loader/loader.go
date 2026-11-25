package ext_private_data_loader

/*
#include <stdlib.h>
#include "../../vulkan/vulkan.h"

#cgo noescape cgoCreatePrivateDataSlotEXT
#cgo noescape cgoDestroyPrivateDataSlotEXT
#cgo noescape cgoGetPrivateDataEXT
#cgo noescape cgoSetPrivateDataEXT

VkResult cgoCreatePrivateDataSlotEXT(
	PFN_vkCreatePrivateDataSlotEXT fn,
    VkDevice                                    device,
    const VkPrivateDataSlotCreateInfoEXT*       pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkPrivateDataSlot*                          pPrivateDataSlot) {

	return fn(device, pCreateInfo, pAllocator, pPrivateDataSlot);
}

void cgoDestroyPrivateDataSlotEXT(
	PFN_vkDestroyPrivateDataSlotEXT fn,
    VkDevice                                    device,
    VkPrivateDataSlot                           privateDataSlot,
    const VkAllocationCallbacks*                pAllocator) {

	fn(device, privateDataSlot, pAllocator);
}

void cgoGetPrivateDataEXT(
	PFN_vkGetPrivateDataEXT fn,
    VkDevice                                    device,
    VkObjectType                                objectType,
    uint64_t                                    objectHandle,
    VkPrivateDataSlot                           privateDataSlot,
    uint64_t*                                   pData) {

	fn(device, objectType, objectHandle, privateDataSlot, pData);
}

VkResult cgoSetPrivateDataEXT(
	PFN_vkSetPrivateDataEXT fn,
    VkDevice                                    device,
    VkObjectType                                objectType,
    uint64_t                                    objectHandle,
    VkPrivateDataSlot                           privateDataSlot,
    uint64_t                                    data) {

	fn(device, objectType, objectHandle, privateDataSlot, data);
}
*/
import "C"
import (
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/loader"
)

type VkPrivateDataSlot loader.VulkanHandle
type VkPrivateDataSlotCreateInfoEXT C.VkPrivateDataSlotCreateInfoEXT
type VkDevicePrivateDataCreateInfoEXT C.VkDevicePrivateDataCreateInfoEXT
type VkPhysicalDevicePrivateDataFeaturesEXT C.VkPhysicalDevicePrivateDataFeaturesEXT

//go:generate go run go.uber.org/mock/mockgen -source loader.go -destination ../mocks/loader.go -package mock_private_data

type Loader interface {
	VkCreatePrivateDataSlotEXT(device loader.VkDevice, pCreateInfo *VkPrivateDataSlotCreateInfoEXT, pAllocator *loader.VkAllocationCallbacks, pPrivateDataSlot *VkPrivateDataSlot) (common.VkResult, error)
	VkDestroyPrivateDataSlotEXT(device loader.VkDevice, privateDataSlot VkPrivateDataSlot, pAllocator *loader.VkAllocationCallbacks)
	VkGetPrivateDataEXT(device loader.VkDevice, objectType loader.VkObjectType, objectHandle loader.Uint64, privateDataSlot VkPrivateDataSlot, data *loader.Uint64)
	VkSetPrivateDataEXT(device loader.VkDevice, objectType loader.VkObjectType, objectHandle loader.Uint64, privateDataSlot VkPrivateDataSlot, data loader.Uint64) (common.VkResult, error)
}

type CLoader struct {
	coreLoader loader.Loader

	cmdCreatePrivateDataSlot  C.PFN_vkCreatePrivateDataSlotEXT
	cmdDestroyPrivateDataSlot C.PFN_vkDestroyPrivateDataSlotEXT
	cmdGetPrivateData         C.PFN_vkGetPrivateDataEXT
	cmdSetPrivateData         C.PFN_vkSetPrivateDataEXT
}

func CreateLoaderFromCore(coreLoader loader.Loader) *CLoader {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	return &CLoader{
		coreLoader: coreLoader,

		cmdCreatePrivateDataSlot:  (C.PFN_vkCreatePrivateDataSlotEXT)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkCreatePrivateDataSlotEXT")))),
		cmdDestroyPrivateDataSlot: (C.PFN_vkDestroyPrivateDataSlotEXT)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkDestroyPrivateDataSlotEXT")))),
		cmdGetPrivateData:         (C.PFN_vkGetPrivateDataEXT)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkGetPrivateDataEXT")))),
		cmdSetPrivateData:         (C.PFN_vkSetPrivateDataEXT)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkSetPrivateDataEXT")))),
	}
}

var _ Loader = &CLoader{}

func (c *CLoader) VkCreatePrivateDataSlotEXT(device loader.VkDevice, pCreateInfo *VkPrivateDataSlotCreateInfoEXT, pAllocator *loader.VkAllocationCallbacks, pPrivateDataSlot *VkPrivateDataSlot) (common.VkResult, error) {
	if c.cmdCreatePrivateDataSlot == nil {
		panic("attempt to call extension method vkCreatePrivateDataSlotEXT when extension not present")
	}

	res := common.VkResult(C.cgoCreatePrivateDataSlotEXT(c.cmdCreatePrivateDataSlot, (C.VkDevice)(unsafe.Pointer(device)), (*C.VkPrivateDataSlotCreateInfoEXT)(pCreateInfo), (*C.VkAllocationCallbacks)(unsafe.Pointer(pAllocator)), (*C.VkPrivateDataSlot)(unsafe.Pointer(pPrivateDataSlot))))
	return res, res.ToError()
}

func (c *CLoader) VkDestroyPrivateDataSlotEXT(device loader.VkDevice, privateDataSlot VkPrivateDataSlot, pAllocator *loader.VkAllocationCallbacks) {
	if c.cmdDestroyPrivateDataSlot == nil {
		panic("attempt to call extension method vkDestroyPrivateDataSlotEXT when extension not present")
	}

	C.cgoDestroyPrivateDataSlotEXT(c.cmdDestroyPrivateDataSlot, (C.VkDevice)(unsafe.Pointer(device)), (C.VkPrivateDataSlot)(unsafe.Pointer(privateDataSlot)), (*C.VkAllocationCallbacks)(unsafe.Pointer(pAllocator)))
}

func (c *CLoader) VkGetPrivateDataEXT(device loader.VkDevice, objectType loader.VkObjectType, objectHandle loader.Uint64, privateDataSlot VkPrivateDataSlot, data *loader.Uint64) {
	if c.cmdGetPrivateData == nil {
		panic("attempt to call extension method vkGetPrivateDataEXT when extension not present")
	}

	C.cgoGetPrivateDataEXT(c.cmdGetPrivateData, (C.VkDevice)(unsafe.Pointer(device)), C.VkObjectType(objectType), C.uint64_t(objectHandle), (C.VkPrivateDataSlot)(unsafe.Pointer(privateDataSlot)), (*C.uint64_t)(data))
}

func (c *CLoader) VkSetPrivateDataEXT(device loader.VkDevice, objectType loader.VkObjectType, objectHandle loader.Uint64, privateDataSlot VkPrivateDataSlot, data loader.Uint64) (common.VkResult, error) {
	if c.cmdSetPrivateData == nil {
		panic("attempt to call extension method vkSetPrivateDataEXT when extension not present")
	}

	res := common.VkResult(C.cgoSetPrivateDataEXT(c.cmdSetPrivateData, (C.VkDevice)(unsafe.Pointer(device)), C.VkObjectType(objectType), C.uint64_t(objectHandle), (C.VkPrivateDataSlot)(unsafe.Pointer(privateDataSlot)), C.uint64_t(data)))
	return res, res.ToError()
}
