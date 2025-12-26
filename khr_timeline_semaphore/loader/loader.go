package khr_timeline_semaphore_loader

/*
#include <stdlib.h>
#include "../../vulkan/vulkan.h"

VkResult cgoGetSemaphoreCounterValueKHR(PFN_vkGetSemaphoreCounterValueKHR fn, VkDevice device, VkSemaphore semaphore, uint64_t *pValue) {
	return fn(device, semaphore, pValue);
}

VkResult cgoSignalSemaphoreKHR(PFN_vkSignalSemaphoreKHR fn, VkDevice device, VkSemaphoreSignalInfo *pSignalInfo) {
	return fn(device, pSignalInfo);
}

VkResult cgoWaitSemaphoresKHR(PFN_vkWaitSemaphoresKHR fn, VkDevice device, VkSemaphoreWaitInfo *pWaitInfo, uint64_t timeout) {
	return fn(device, pWaitInfo, timeout);
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

//go:generate mockgen -source loader.go -destination ../mocks/loader.go -package mock_timeline_semaphore

type Loader interface {
	VkGetSemaphoreCounterValueKHR(device loader.VkDevice, semaphore loader.VkSemaphore, pValue *loader.Uint64) (common.VkResult, error)
	VkSignalSemaphoreKHR(device loader.VkDevice, pSignalInfo *VkSemaphoreSignalInfoKHR) (common.VkResult, error)
	VkWaitSemaphoresKHR(device loader.VkDevice, pWaitInfo *VkSemaphoreWaitInfoKHR, timeout loader.Uint64) (common.VkResult, error)
}

type VkSemaphoreSignalInfoKHR C.VkSemaphoreSignalInfoKHR
type VkSemaphoreWaitInfoKHR C.VkSemaphoreWaitInfoKHR
type VkPhysicalDeviceTimelineSemaphoreFeaturesKHR C.VkPhysicalDeviceTimelineSemaphoreFeaturesKHR
type VkPhysicalDeviceTimelineSemaphorePropertiesKHR C.VkPhysicalDeviceTimelineSemaphorePropertiesKHR
type VkSemaphoreTypeCreateInfoKHR C.VkSemaphoreTypeCreateInfoKHR
type VkTimelineSemaphoreSubmitInfoKHR C.VkTimelineSemaphoreSubmitInfoKHR

type CLoader struct {
	coreLoader loader.Loader

	getSemaphoreCounterValue C.PFN_vkGetSemaphoreCounterValueKHR
	signalSemaphore          C.PFN_vkSignalSemaphoreKHR
	waitSemaphores           C.PFN_vkWaitSemaphoresKHR
}

func CreateLoaderFromCore(coreLoader loader.Loader) *CLoader {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	return &CLoader{
		coreLoader: coreLoader,

		getSemaphoreCounterValue: (C.PFN_vkGetSemaphoreCounterValueKHR)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkGetSemaphoreCounterValueKHR")))),
		signalSemaphore:          (C.PFN_vkSignalSemaphoreKHR)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkSignalSemaphoreKHR")))),
		waitSemaphores:           (C.PFN_vkWaitSemaphoresKHR)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkWaitSemaphoresKHR")))),
	}
}

func (d *CLoader) VkGetSemaphoreCounterValueKHR(device loader.VkDevice, semaphore loader.VkSemaphore, pValue *loader.Uint64) (common.VkResult, error) {
	if d.getSemaphoreCounterValue == nil {
		panic("attempt to call extension method vkGetSemaphoreCounterValueKHR when extension not present")
	}

	res := common.VkResult(C.cgoGetSemaphoreCounterValueKHR(
		d.getSemaphoreCounterValue,
		C.VkDevice(unsafe.Pointer(device)),
		C.VkSemaphore(unsafe.Pointer(semaphore)),
		(*C.uint64_t)(pValue),
	))
	return res, res.ToError()
}

func (d *CLoader) VkSignalSemaphoreKHR(device loader.VkDevice, pSignalInfo *VkSemaphoreSignalInfoKHR) (common.VkResult, error) {
	if d.signalSemaphore == nil {
		panic("attempt to call extension method vkSignalSemaphoreKHR when extension not present")
	}

	res := common.VkResult(C.cgoSignalSemaphoreKHR(
		d.signalSemaphore,
		C.VkDevice(unsafe.Pointer(device)),
		(*C.VkSemaphoreSignalInfoKHR)(pSignalInfo),
	))

	return res, res.ToError()
}

func (d *CLoader) VkWaitSemaphoresKHR(device loader.VkDevice, pWaitInfo *VkSemaphoreWaitInfoKHR, timeout loader.Uint64) (common.VkResult, error) {
	if d.waitSemaphores == nil {
		panic("attempt to call extension method vkWaitSemaphoresKHR when extension not present")
	}

	res := common.VkResult(C.cgoWaitSemaphoresKHR(
		d.waitSemaphores,
		C.VkDevice(unsafe.Pointer(device)),
		(*C.VkSemaphoreWaitInfoKHR)(pWaitInfo),
		C.uint64_t(timeout),
	))

	return res, res.ToError()
}
