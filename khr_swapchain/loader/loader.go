package khr_swapchain_loader

//go:generate mockgen -source loader.go -destination ../mocks/loader.go -package mock_swapchain

/*
#include <stdlib.h>
#include "../../vulkan/vulkan.h"

VkResult cgoAcquireNextImageKHR(PFN_vkAcquireNextImageKHR fn, VkDevice device, VkSwapchainKHR swapchain, uint64_t timeout, VkSemaphore semaphore, VkFence fence, uint32_t* pImageIndex) {
	return fn(device, swapchain, timeout, semaphore, fence, pImageIndex);
}

VkResult cgoCreateSwapchainKHR(PFN_vkCreateSwapchainKHR fn, VkDevice device, VkSwapchainCreateInfoKHR* pCreateInfo, VkAllocationCallbacks* pAllocator, VkSwapchainKHR* pSwapchain) {
	return fn(device, pCreateInfo, pAllocator, pSwapchain);
}

void cgoDestroySwapchainKHR(PFN_vkDestroySwapchainKHR fn, VkDevice device, VkSwapchainKHR swapchain, VkAllocationCallbacks* pAllocator) {
	fn(device, swapchain, pAllocator);
}

VkResult cgoGetSwapchainImagesKHR(PFN_vkGetSwapchainImagesKHR fn, VkDevice device, VkSwapchainKHR swapchain, uint32_t* pSwapchainImageCount, VkImage* pSwapchainImages) {
	return fn(device, swapchain, pSwapchainImageCount, pSwapchainImages);
}

VkResult cgoQueuePresentKHR(PFN_vkQueuePresentKHR fn, VkQueue queue, VkPresentInfoKHR* pPresentInfo) {
	return fn(queue, pPresentInfo);
}

VkResult cgoSwapchainAcquireNextImage2KHR(PFN_vkAcquireNextImage2KHR fn, VkDevice device, VkAcquireNextImageInfoKHR *pAcquireInfo, uint32_t *pImageIndex) {
	return fn(device, pAcquireInfo, pImageIndex);
}

VkResult cgoSwapchainGetDeviceGroupPresentCapabilitiesKHR(PFN_vkGetDeviceGroupPresentCapabilitiesKHR fn, VkDevice device, VkDeviceGroupPresentCapabilitiesKHR *pDeviceGroupPresentCapabilities) {
	return fn(device, pDeviceGroupPresentCapabilities);
}

VkResult cgoSwapchainGetDeviceGroupSurfacePresentModesKHR(PFN_vkGetDeviceGroupSurfacePresentModesKHR fn, VkDevice device, VkSurfaceKHR surface, VkDeviceGroupPresentModeFlagsKHR *pModes) {
	return fn(device, surface, pModes);
}

VkResult cgoSwapchainGetPhysicalDevicePresentRectanglesKHR(PFN_vkGetPhysicalDevicePresentRectanglesKHR fn, VkPhysicalDevice physicalDevice, VkSurfaceKHR surface, uint32_t *pRectCount, VkRect2D *pRects) {
	return fn(physicalDevice, surface, pRectCount, pRects);
}
*/
import "C"
import (
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/loader"
	khr_surface_driver "github.com/vkngwrapper/extensions/v3/khr_surface/loader"
	_ "github.com/vkngwrapper/extensions/v3/vulkan"
)

type CLoader struct {
	driver           loader.Loader
	createFunc       C.PFN_vkCreateSwapchainKHR
	destroyFunc      C.PFN_vkDestroySwapchainKHR
	getImagesFunc    C.PFN_vkGetSwapchainImagesKHR
	acquireNextFunc  C.PFN_vkAcquireNextImageKHR
	queuePresentFunc C.PFN_vkQueuePresentKHR

	acquireNextImageFunc                  C.PFN_vkAcquireNextImage2KHR
	getDeviceGroupPresentCapsFunc         C.PFN_vkGetDeviceGroupPresentCapabilitiesKHR
	getDeviceGroupSurfacePresentModesFunc C.PFN_vkGetDeviceGroupSurfacePresentModesKHR
	getPhysicalDevicePresentRectsFunc     C.PFN_vkGetPhysicalDevicePresentRectanglesKHR
}

type VkSwapchainKHR loader.VulkanHandle
type VkSwapchainCreateInfoKHR C.VkSwapchainCreateInfoKHR
type VkPresentInfoKHR C.VkPresentInfoKHR
type VkAcquireNextImageInfoKHR C.VkAcquireNextImageInfoKHR
type VkDeviceGroupPresentCapabilitiesKHR C.VkDeviceGroupPresentCapabilitiesKHR
type VkBindImageMemorySwapchainInfoKHR C.VkBindImageMemorySwapchainInfoKHR
type VkImageSwapchainCreateInfoKHR C.VkImageSwapchainCreateInfoKHR
type VkDeviceGroupPresentInfoKHR C.VkDeviceGroupPresentInfoKHR
type VkDeviceGroupSwapchainCreateInfoKHR C.VkDeviceGroupSwapchainCreateInfoKHR
type VkDeviceGroupPresentModeFlagsKHR C.VkDeviceGroupPresentModeFlagsKHR

type Loader interface {
	VkCreateSwapchainKHR(device loader.VkDevice, pCreateInfo *VkSwapchainCreateInfoKHR, pAllocator *loader.VkAllocationCallbacks, pSwapchain *VkSwapchainKHR) (common.VkResult, error)
	VkDestroySwapchainKHR(device loader.VkDevice, swapchain VkSwapchainKHR, pAllocator *loader.VkAllocationCallbacks)
	VkGetSwapchainImagesKHR(device loader.VkDevice, swapchain VkSwapchainKHR, pSwapchainImageCount *loader.Uint32, pSwapchainImages *loader.VkImage) (common.VkResult, error)
	VkAcquireNextImageKHR(device loader.VkDevice, swapchain VkSwapchainKHR, timeout loader.Uint64, semaphore loader.VkSemaphore, fence loader.VkFence, pImageIndex *loader.Uint32) (common.VkResult, error)
	VkQueuePresentKHR(queue loader.VkQueue, pPresentInfo *VkPresentInfoKHR) (common.VkResult, error)
	VkAcquireNextImage2KHR(device loader.VkDevice, pAcquireInfo *VkAcquireNextImageInfoKHR, pImageIndex *loader.Uint32) (common.VkResult, error)
	VkGetDeviceGroupPresentCapabilitiesKHR(device loader.VkDevice, pDeviceGroupPresentCapabilities *VkDeviceGroupPresentCapabilitiesKHR) (common.VkResult, error)
	VkGetDeviceGroupSurfacePresentModesKHR(device loader.VkDevice, surface khr_surface_driver.VkSurfaceKHR, pModes *VkDeviceGroupPresentModeFlagsKHR) (common.VkResult, error)
	VkGetPhysicalDevicePresentRectanglesKHR(physicalDevice loader.VkPhysicalDevice, surface khr_surface_driver.VkSurfaceKHR, pRectCount *loader.Uint32, pRects *loader.VkRect2D) (common.VkResult, error)
}

func CreateLoaderFromCore(coreLoader loader.Loader) *CLoader {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	return &CLoader{
		driver:           coreLoader,
		createFunc:       (C.PFN_vkCreateSwapchainKHR)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkCreateSwapchainKHR")))),
		destroyFunc:      (C.PFN_vkDestroySwapchainKHR)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkDestroySwapchainKHR")))),
		getImagesFunc:    (C.PFN_vkGetSwapchainImagesKHR)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkGetSwapchainImagesKHR")))),
		acquireNextFunc:  (C.PFN_vkAcquireNextImageKHR)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkAcquireNextImageKHR")))),
		queuePresentFunc: (C.PFN_vkQueuePresentKHR)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkQueuePresentKHR")))),

		acquireNextImageFunc:                  (C.PFN_vkAcquireNextImage2KHR)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkAcquireNextImage2KHR")))),
		getDeviceGroupPresentCapsFunc:         (C.PFN_vkGetDeviceGroupPresentCapabilitiesKHR)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkGetDeviceGroupPresentCapabilitiesKHR")))),
		getDeviceGroupSurfacePresentModesFunc: (C.PFN_vkGetDeviceGroupSurfacePresentModesKHR)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkGetDeviceGroupSurfacePresentModesKHR")))),
		getPhysicalDevicePresentRectsFunc:     (C.PFN_vkGetPhysicalDevicePresentRectanglesKHR)(coreLoader.LoadInstanceProcAddr((*loader.Char)(arena.CString("vkGetPhysicalDevicePresentRectanglesKHR")))),
	}
}

func (d *CLoader) VkCreateSwapchainKHR(device loader.VkDevice, pCreateInfo *VkSwapchainCreateInfoKHR, pAllocator *loader.VkAllocationCallbacks, pSwapchain *VkSwapchainKHR) (common.VkResult, error) {
	if d.createFunc == nil {
		panic("attempt to call extension method vkCreateSwapchainKHR when extension not present")
	}

	res := common.VkResult(C.cgoCreateSwapchainKHR(d.createFunc,
		C.VkDevice(unsafe.Pointer(device)),
		(*C.VkSwapchainCreateInfoKHR)(pCreateInfo),
		(*C.VkAllocationCallbacks)(unsafe.Pointer(pAllocator)),
		(*C.VkSwapchainKHR)(unsafe.Pointer(pSwapchain))))

	return res, res.ToError()
}

func (d *CLoader) VkDestroySwapchainKHR(device loader.VkDevice, swapchain VkSwapchainKHR, pAllocator *loader.VkAllocationCallbacks) {
	if d.destroyFunc == nil {
		panic("attempt to call extension method vkDestroySwapchainKHR when extension not present")
	}

	C.cgoDestroySwapchainKHR(d.destroyFunc,
		C.VkDevice(unsafe.Pointer(device)),
		C.VkSwapchainKHR(unsafe.Pointer(swapchain)),
		(*C.VkAllocationCallbacks)(unsafe.Pointer(pAllocator)))
}

func (d *CLoader) VkGetSwapchainImagesKHR(device loader.VkDevice, swapchain VkSwapchainKHR, pSwapchainImageCount *loader.Uint32, pSwapchainImages *loader.VkImage) (common.VkResult, error) {
	if d.getImagesFunc == nil {
		panic("attempt to call extension method vkGetSwapchainImagesKHR when extension not present")
	}

	res := common.VkResult(C.cgoGetSwapchainImagesKHR(d.getImagesFunc,
		C.VkDevice(unsafe.Pointer(device)),
		C.VkSwapchainKHR(unsafe.Pointer(swapchain)),
		(*C.uint32_t)(unsafe.Pointer(pSwapchainImageCount)),
		(*C.VkImage)(unsafe.Pointer(pSwapchainImages))))

	return res, res.ToError()
}

func (d *CLoader) VkAcquireNextImageKHR(device loader.VkDevice, swapchain VkSwapchainKHR, timeout loader.Uint64, semaphore loader.VkSemaphore, fence loader.VkFence, pImageIndex *loader.Uint32) (common.VkResult, error) {
	if d.acquireNextFunc == nil {
		panic("attempt to call extension method vkAcquireNextImageKHR when extension not present")
	}

	res := common.VkResult(C.cgoAcquireNextImageKHR(d.acquireNextFunc,
		C.VkDevice(unsafe.Pointer(device)),
		C.VkSwapchainKHR(unsafe.Pointer(swapchain)),
		C.uint64_t(timeout),
		C.VkSemaphore(unsafe.Pointer(semaphore)),
		C.VkFence(unsafe.Pointer(fence)),
		(*C.uint32_t)(unsafe.Pointer(pImageIndex)),
	))

	return res, res.ToError()
}

func (d *CLoader) VkQueuePresentKHR(queue loader.VkQueue, pPresentInfo *VkPresentInfoKHR) (common.VkResult, error) {
	if d.queuePresentFunc == nil {
		panic("attempt to call extension method vkQueuePresentKHR when extension not present")
	}

	res := common.VkResult(C.cgoQueuePresentKHR(d.queuePresentFunc,
		C.VkQueue(unsafe.Pointer(queue)),
		(*C.VkPresentInfoKHR)(pPresentInfo)))

	return res, res.ToError()
}

func (d *CLoader) VkAcquireNextImage2KHR(device loader.VkDevice, pAcquireInfo *VkAcquireNextImageInfoKHR, pImageIndex *loader.Uint32) (common.VkResult, error) {
	if d.acquireNextImageFunc == nil {
		panic("attempt to call extension method vkAcquireNextImage2KHR when extension (or core1.1) not present")
	}

	res := common.VkResult(C.cgoSwapchainAcquireNextImage2KHR(d.acquireNextImageFunc,
		C.VkDevice(unsafe.Pointer(device)),
		(*C.VkAcquireNextImageInfoKHR)(pAcquireInfo),
		(*C.uint32_t)(pImageIndex),
	))

	return res, res.ToError()
}

func (d *CLoader) VkGetDeviceGroupPresentCapabilitiesKHR(device loader.VkDevice, pDeviceGroupPresentCapabilities *VkDeviceGroupPresentCapabilitiesKHR) (common.VkResult, error) {
	if d.getDeviceGroupPresentCapsFunc == nil {
		panic("attempt to call extension method vkGetDeviceGroupPresentCapabilitiesKHR when extension (or core1.1) not present")
	}

	res := common.VkResult(C.cgoSwapchainGetDeviceGroupPresentCapabilitiesKHR(d.getDeviceGroupPresentCapsFunc,
		C.VkDevice(unsafe.Pointer(device)),
		(*C.VkDeviceGroupPresentCapabilitiesKHR)(pDeviceGroupPresentCapabilities),
	))

	return res, res.ToError()
}

func (d *CLoader) VkGetDeviceGroupSurfacePresentModesKHR(device loader.VkDevice, surface khr_surface_driver.VkSurfaceKHR, pModes *VkDeviceGroupPresentModeFlagsKHR) (common.VkResult, error) {
	if d.getDeviceGroupSurfacePresentModesFunc == nil {
		panic("attempt to call extension method vkGetDeviceGroupSurfacePresentModesKHR when extension (or core1.1) not present")
	}

	res := common.VkResult(C.cgoSwapchainGetDeviceGroupSurfacePresentModesKHR(d.getDeviceGroupSurfacePresentModesFunc,
		C.VkDevice(unsafe.Pointer(device)),
		C.VkSurfaceKHR(unsafe.Pointer(surface)),
		(*C.VkDeviceGroupPresentModeFlagsKHR)(pModes),
	))

	return res, res.ToError()
}

func (d *CLoader) VkGetPhysicalDevicePresentRectanglesKHR(physicalDevice loader.VkPhysicalDevice, surface khr_surface_driver.VkSurfaceKHR, pRectCount *loader.Uint32, pRects *loader.VkRect2D) (common.VkResult, error) {
	if d.getPhysicalDevicePresentRectsFunc == nil {
		panic("attempt to call extension method vkGetDeviceGroupSurfacePresentModesKHR when extension (or core1.1) not present")
	}

	res := common.VkResult(C.cgoSwapchainGetPhysicalDevicePresentRectanglesKHR(d.getPhysicalDevicePresentRectsFunc,
		C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)),
		C.VkSurfaceKHR(unsafe.Pointer(surface)),
		(*C.uint32_t)(pRectCount),
		(*C.VkRect2D)(unsafe.Pointer(pRects)),
	))

	return res, res.ToError()
}
