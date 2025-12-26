package khr_surface_loader

//go:generate mockgen -source loader.go -destination ../mocks/loader.go -package mock_surface

/*
#include <stdlib.h>
#include "../../vulkan/vulkan.h"

VkResult cgoGetPhysicalDeviceSurfaceCapabilitiesKHR(PFN_vkGetPhysicalDeviceSurfaceCapabilitiesKHR fn, VkPhysicalDevice physicalDevice, VkSurfaceKHR surface, VkSurfaceCapabilitiesKHR* pSurfaceCapabilities) {
	return fn(physicalDevice, surface, pSurfaceCapabilities);
}

VkResult cgoGetPhysicalDeviceSurfaceSupportKHR(PFN_vkGetPhysicalDeviceSurfaceSupportKHR fn, VkPhysicalDevice physicalDevice, uint32_t queueFamilyIndex, VkSurfaceKHR surface, VkBool32* pSupported) {
	return fn(physicalDevice, queueFamilyIndex, surface, pSupported);
}

void cgoDestroySurfaceKHR(PFN_vkDestroySurfaceKHR fn, VkInstance instance, VkSurfaceKHR surface, VkAllocationCallbacks* pAllocator) {
	fn(instance, surface, pAllocator);
}

VkResult cgoGetPhysicalDeviceSurfaceFormatsKHR(PFN_vkGetPhysicalDeviceSurfaceFormatsKHR fn,VkPhysicalDevice physicalDevice, VkSurfaceKHR surface, uint32_t* pSurfaceFormatCount, VkSurfaceFormatKHR* pSurfaceFormats) {
	return fn(physicalDevice, surface, pSurfaceFormatCount, pSurfaceFormats);
}

VkResult cgoGetPhysicalDeviceSurfacePresentModesKHR(PFN_vkGetPhysicalDeviceSurfacePresentModesKHR fn, VkPhysicalDevice physicalDevice, VkSurfaceKHR surface, uint32_t* pPresentModeCount, VkPresentModeKHR* pPresentModes) {
	return fn(physicalDevice, surface, pPresentModeCount, pPresentModes);
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

type CLoader struct {
	physicalSurfaceCapabilitiesFunc C.PFN_vkGetPhysicalDeviceSurfaceCapabilitiesKHR
	physicalSurfaceSupportFunc      C.PFN_vkGetPhysicalDeviceSurfaceSupportKHR
	surfaceFormatsFunc              C.PFN_vkGetPhysicalDeviceSurfaceFormatsKHR
	presentModesFunc                C.PFN_vkGetPhysicalDeviceSurfacePresentModesKHR
	destroyFunc                     C.PFN_vkDestroySurfaceKHR
}

type VkSurfaceKHR loader.VulkanHandle
type VkSurfaceCapabilitiesKHR C.VkSurfaceCapabilitiesKHR
type VkSurfaceFormatKHR C.VkSurfaceFormatKHR
type VkPresentModeKHR C.VkPresentModeKHR

type Loader interface {
	VkDestroySurfaceKHR(instance loader.VkInstance, surface VkSurfaceKHR, pAllocator *loader.VkAllocationCallbacks)
	VkGetPhysicalDeviceSurfaceCapabilitiesKHR(physicalDevice loader.VkPhysicalDevice, surface VkSurfaceKHR, pSurfaceCapabilities *VkSurfaceCapabilitiesKHR) (common.VkResult, error)
	VkGetPhysicalDeviceSurfaceSupportKHR(physicalDevice loader.VkPhysicalDevice, queueFamilyIndex loader.Uint32, surface VkSurfaceKHR, pSupported *loader.VkBool32) (common.VkResult, error)
	VkGetPhysicalDeviceSurfaceFormatsKHR(physicalDevice loader.VkPhysicalDevice, surface VkSurfaceKHR, pSurfaceFormatCount *loader.Uint32, pSurfaceFormats *VkSurfaceFormatKHR) (common.VkResult, error)
	VkGetPhysicalDeviceSurfacePresentModesKHR(physicalDevice loader.VkPhysicalDevice, surface VkSurfaceKHR, pPresentModeCount *loader.Uint32, pPresentModes *VkPresentModeKHR) (common.VkResult, error)
}

func CreateDriverFromCore(coreLoader loader.Loader) *CLoader {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	physicalSurfaceCapabilitiesFunc := (C.PFN_vkGetPhysicalDeviceSurfaceCapabilitiesKHR)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkGetPhysicalDeviceSurfaceCapabilitiesKHR"))))
	physicalSurfaceSupportFunc := (C.PFN_vkGetPhysicalDeviceSurfaceSupportKHR)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkGetPhysicalDeviceSurfaceSupportKHR"))))
	surfaceFormatsFunc := (C.PFN_vkGetPhysicalDeviceSurfaceFormatsKHR)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkGetPhysicalDeviceSurfaceFormatsKHR"))))
	presentModesFunc := (C.PFN_vkGetPhysicalDeviceSurfacePresentModesKHR)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkGetPhysicalDeviceSurfacePresentModesKHR"))))
	destroyFunc := (C.PFN_vkDestroySurfaceKHR)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkDestroySurfaceKHR"))))

	return &CLoader{
		physicalSurfaceSupportFunc:      physicalSurfaceSupportFunc,
		physicalSurfaceCapabilitiesFunc: physicalSurfaceCapabilitiesFunc,
		surfaceFormatsFunc:              surfaceFormatsFunc,
		presentModesFunc:                presentModesFunc,
		destroyFunc:                     destroyFunc,
	}
}

func (d *CLoader) VkDestroySurfaceKHR(instance loader.VkInstance, surface VkSurfaceKHR, pAllocator *loader.VkAllocationCallbacks) {
	if d.destroyFunc == nil {
		panic("attempt to call extension method vkDestroySurfaceKHR when extension not present")
	}

	C.cgoDestroySurfaceKHR(d.destroyFunc,
		C.VkInstance(unsafe.Pointer(instance)),
		C.VkSurfaceKHR(unsafe.Pointer(surface)),
		(*C.VkAllocationCallbacks)(unsafe.Pointer(pAllocator)))
}

func (d *CLoader) VkGetPhysicalDeviceSurfaceCapabilitiesKHR(physicalDevice loader.VkPhysicalDevice, surface VkSurfaceKHR, pSurfaceCapabilities *VkSurfaceCapabilitiesKHR) (common.VkResult, error) {
	if d.physicalSurfaceCapabilitiesFunc == nil {
		panic("attempt to call extension method vkGetPhysicalDeviceSurfaceCapabilitiesKHR")
	}

	res := common.VkResult(C.cgoGetPhysicalDeviceSurfaceCapabilitiesKHR(d.physicalSurfaceCapabilitiesFunc,
		C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)),
		C.VkSurfaceKHR(unsafe.Pointer(surface)),
		(*C.VkSurfaceCapabilitiesKHR)(pSurfaceCapabilities)))

	return res, res.ToError()
}

func (d *CLoader) VkGetPhysicalDeviceSurfaceSupportKHR(physicalDevice loader.VkPhysicalDevice, queueFamilyIndex loader.Uint32, surface VkSurfaceKHR, pSupported *loader.VkBool32) (common.VkResult, error) {
	if d.physicalSurfaceSupportFunc == nil {
		panic("attempt to call extension method vkGetPhysicalDeviceSurfaceSupportKHR")
	}

	res := common.VkResult(C.cgoGetPhysicalDeviceSurfaceSupportKHR(d.physicalSurfaceSupportFunc,
		C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)),
		C.uint32_t(queueFamilyIndex),
		C.VkSurfaceKHR(unsafe.Pointer(surface)),
		(*C.VkBool32)(unsafe.Pointer(pSupported))))

	return res, res.ToError()
}

func (d *CLoader) VkGetPhysicalDeviceSurfaceFormatsKHR(physicalDevice loader.VkPhysicalDevice, surface VkSurfaceKHR, pSurfaceFormatCount *loader.Uint32, pSurfaceFormats *VkSurfaceFormatKHR) (common.VkResult, error) {
	if d.surfaceFormatsFunc == nil {
		panic("attempt to call extension method vkGetPhysicalDeviceSurfaceFormatsKHR")
	}

	res := common.VkResult(C.cgoGetPhysicalDeviceSurfaceFormatsKHR(d.surfaceFormatsFunc,
		C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)),
		C.VkSurfaceKHR(unsafe.Pointer(surface)),
		(*C.uint32_t)(unsafe.Pointer(pSurfaceFormatCount)),
		(*C.VkSurfaceFormatKHR)(pSurfaceFormats)))
	return res, res.ToError()
}

func (d *CLoader) VkGetPhysicalDeviceSurfacePresentModesKHR(physicalDevice loader.VkPhysicalDevice, surface VkSurfaceKHR, pPresentModeCount *loader.Uint32, pPresentModes *VkPresentModeKHR) (common.VkResult, error) {
	if d.presentModesFunc == nil {
		panic("attempt to call extension method vkGetPhysicalDeviceSurfacePresentModesKHR")
	}

	res := common.VkResult(C.cgoGetPhysicalDeviceSurfacePresentModesKHR(d.presentModesFunc,
		C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)),
		C.VkSurfaceKHR(unsafe.Pointer(surface)),
		(*C.uint32_t)(unsafe.Pointer(pPresentModeCount)),
		(*C.VkPresentModeKHR)(pPresentModes)))

	return res, res.ToError()
}
