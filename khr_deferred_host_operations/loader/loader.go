package khr_deferred_host_operations_loader

/*
#include <stdlib.h>
#include "../../vulkan/vulkan.h"

VkResult cgoCreateDeferredOperationKHR(PFN_vkCreateDeferredOperationKHR fn, VkDevice device, VkAllocationCallbacks *pAllocator, VkDeferredOperationKHR *pDeferredOperation) {
	return fn(device, pAllocator, pDeferredOperation);
}

void cgoDestroyDeferredOperationKHR(PFN_vkDestroyDeferredOperationKHR fn, VkDevice device, VkDeferredOperationKHR operation, VkAllocationCallbacks *pAllocator) {
	fn(device, operation, pAllocator);
}

VkResult cgoDeferredOperationJoinKHR(PFN_vkDeferredOperationJoinKHR fn, VkDevice device, VkDeferredOperationKHR operation) {
	return fn(device, operation);
}

uint32_t cgoGetDeferredOperationMaxConcurrencyKHR(PFN_vkGetDeferredOperationMaxConcurrencyKHR fn, VkDevice device, VkDeferredOperationKHR operation) {
	return fn(device, operation);
}

VkResult cgoGetDeferredOperationResultKHR(PFN_vkGetDeferredOperationResultKHR fn, VkDevice device, VkDeferredOperationKHR operation) {
	return fn(device, operation);
}

*/
import "C"
import (
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/loader"
)

//go:generate mockgen -source loader.go -destination ../mocks/loader.go -package mock_deferred_host_operations

type Loader interface {
	VkCreateDeferredOperationKHR(device loader.VkDevice, pAllocator *loader.VkAllocationCallbacks, pDeferredOperation *VkDeferredOperationKHR) (common.VkResult, error)
	VkDeferredOperationJoinKHR(device loader.VkDevice, operation VkDeferredOperationKHR) (common.VkResult, error)
	VkDestroyDeferredOperationKHR(device loader.VkDevice, operation VkDeferredOperationKHR, pAllocator *loader.VkAllocationCallbacks)
	VkGetDeferredOperationMaxConcurrencyKHR(device loader.VkDevice, operation VkDeferredOperationKHR) loader.Uint32
	VkGetDeferredOperationResultKHR(device loader.VkDevice, operation VkDeferredOperationKHR) (common.VkResult, error)
}

type VkDeferredOperationKHR loader.VulkanHandle

type CLoader struct {
	coreLoader loader.Loader

	createDeferredOperation  C.PFN_vkCreateDeferredOperationKHR
	deferredOperationJoin    C.PFN_vkDeferredOperationJoinKHR
	destroyDeferredOperation C.PFN_vkDestroyDeferredOperationKHR

	getDeferredOperationMaxConcurrency C.PFN_vkGetDeferredOperationMaxConcurrencyKHR
	getDeferredOperationResult         C.PFN_vkGetDeferredOperationResultKHR
}

func CreateLoaderFromCore(coreLoader loader.Loader) *CLoader {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	return &CLoader{
		coreLoader: coreLoader,

		createDeferredOperation:  (C.PFN_vkCreateDeferredOperationKHR)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkCreateDeferredOperationKHR")))),
		deferredOperationJoin:    (C.PFN_vkDeferredOperationJoinKHR)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkDeferredOperationJoinKHR")))),
		destroyDeferredOperation: (C.PFN_vkDestroyDeferredOperationKHR)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkDestroyDeferredOperationKHR")))),

		getDeferredOperationMaxConcurrency: (C.PFN_vkGetDeferredOperationMaxConcurrencyKHR)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkGetDeferredOperationMaxConcurrencyKHR")))),
		getDeferredOperationResult:         (C.PFN_vkGetDeferredOperationResultKHR)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkGetDeferredOperationResultKHR")))),
	}
}

func (c *CLoader) VkCreateDeferredOperationKHR(device loader.VkDevice, pAllocator *loader.VkAllocationCallbacks, pDeferredOperation *VkDeferredOperationKHR) (common.VkResult, error) {
	if c.createDeferredOperation == nil {
		panic("attempt to call extension method vkCreateDeferredOperationKHR when extension not present")
	}

	res := common.VkResult(C.cgoCreateDeferredOperationKHR(
		c.createDeferredOperation,
		C.VkDevice(unsafe.Pointer(device)),
		(*C.VkAllocationCallbacks)(unsafe.Pointer(pAllocator)),
		(*C.VkDeferredOperationKHR)(unsafe.Pointer(pDeferredOperation)),
	))
	return res, res.ToError()
}

func (c *CLoader) VkDestroyDeferredOperationKHR(device loader.VkDevice, operation VkDeferredOperationKHR, pAllocator *loader.VkAllocationCallbacks) {
	if c.destroyDeferredOperation == nil {
		panic("attempt to call extension method vkDestroyDeferredOperationKHR when extension not present")
	}

	C.cgoDestroyDeferredOperationKHR(
		c.destroyDeferredOperation,
		C.VkDevice(unsafe.Pointer(device)),
		C.VkDeferredOperationKHR(unsafe.Pointer(operation)),
		(*C.VkAllocationCallbacks)(unsafe.Pointer(pAllocator)),
	)
}

func (c *CLoader) VkDeferredOperationJoinKHR(device loader.VkDevice, operation VkDeferredOperationKHR) (common.VkResult, error) {
	if c.deferredOperationJoin == nil {
		panic("attempt to call extension method vkDeferredOperationJoinKHR when extension not present")
	}

	res := common.VkResult(C.cgoDeferredOperationJoinKHR(
		c.deferredOperationJoin,
		C.VkDevice(unsafe.Pointer(device)),
		C.VkDeferredOperationKHR(unsafe.Pointer(operation)),
	))
	return res, res.ToError()
}

func (c *CLoader) VkGetDeferredOperationResultKHR(device loader.VkDevice, operation VkDeferredOperationKHR) (common.VkResult, error) {
	if c.getDeferredOperationResult == nil {
		panic("attempt to call extension method vkGetDeferredOperationResultKHR when extension not present")
	}

	res := common.VkResult(C.cgoGetDeferredOperationResultKHR(
		c.getDeferredOperationResult,
		C.VkDevice(unsafe.Pointer(device)),
		C.VkDeferredOperationKHR(unsafe.Pointer(operation)),
	))
	return res, res.ToError()
}

func (c *CLoader) VkGetDeferredOperationMaxConcurrencyKHR(device loader.VkDevice, operation VkDeferredOperationKHR) loader.Uint32 {
	if c.getDeferredOperationMaxConcurrency == nil {
		panic("attempt to call extension method vkGetDeferredOperationMaxConcurrencyKHR when extension not present")
	}

	result := C.cgoGetDeferredOperationMaxConcurrencyKHR(
		c.getDeferredOperationMaxConcurrency,
		C.VkDevice(unsafe.Pointer(device)),
		C.VkDeferredOperationKHR(unsafe.Pointer(operation)),
	)

	return loader.Uint32(result)
}
