//go:build windows

package ext_full_screen_exclusive_loader

/*
#define VK_USE_PLATFORM_WIN32_KHR
#include <stdlib.h>
#include "../../vulkan/vulkan.h"

VkResult cgoAcquireFullScreenExclusiveModeEXT(PFN_vkAcquireFullScreenExclusiveModeEXT fn, VkDevice device, VkSwapchainKHR swapchain) {
	return fn(device, swapchain);
}

VkResult cgoGetPhysicalDeviceSurfacePresentModes2EXT(PFN_vkGetPhysicalDeviceSurfacePresentModes2EXT fn, VkPhysicalDevice physicalDevice, VkPhysicalDeviceSurfaceInfo2KHR *pSurfaceInfo, uint32_t *pPresentModeCount, VkPresentModeKHR *pPresentModes) {
	return fn(physicalDevice, pSurfaceInfo, pPresentModeCount, pPresentModes);
}

VkResult cgoReleaseFullScreenExclusiveModeEXT(PFN_vkReleaseFullScreenExclusiveModeEXT fn, VkDevice device, VkSwapchainKHR swapchain) {
	return fn(device, swapchain);
}

VkResult cgoGetDeviceGroupSurfacePresentModes2EXT(PFN_vkGetDeviceGroupSurfacePresentModes2EXT fn, VkDevice device, VkPhysicalDeviceSurfaceInfo2KHR *pSurfaceInfo, VkDeviceGroupPresentModeFlagsKHR *pModes) {
	return fn(device, pSurfaceInfo, pModes);
}

*/
import "C"
import (
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/loader"
	khr_device_group_loader "github.com/vkngwrapper/extensions/v3/khr_device_group/loader"
	khr_get_surface_capabilities2_loader "github.com/vkngwrapper/extensions/v3/khr_get_surface_capabilities2/loader"
	khr_surface_loader "github.com/vkngwrapper/extensions/v3/khr_surface/loader"
	khr_swapchain_loader "github.com/vkngwrapper/extensions/v3/khr_swapchain/loader"
)

//go:generate mockgen -source loader.go -destination ../mocks/loader.go -package mock_full_screen_exclusive

type Loader interface {
	VkAcquireFullScreenExclusiveModeEXT(device loader.VkDevice, swapchain khr_swapchain_loader.VkSwapchainKHR) (common.VkResult, error)
	VkGetPhysicalDeviceSurfacePresentModes2EXT(physicalDevice loader.VkPhysicalDevice, pSurfaceInfo *khr_get_surface_capabilities2_loader.VkPhysicalDeviceSurfaceInfo2KHR, pPresentModeCount *loader.Uint32, pPresentModes *khr_surface_loader.VkPresentModeKHR) (common.VkResult, error)
	VkReleaseFullScreenExclusiveModeEXT(device loader.VkDevice, swapchain khr_swapchain_loader.VkSwapchainKHR) (common.VkResult, error)

	VkGetDeviceGroupSurfacePresentModes2EXT(device loader.VkDevice, pSurfaceInfo *khr_get_surface_capabilities2_loader.VkPhysicalDeviceSurfaceInfo2KHR, pModes *khr_device_group_loader.VkDeviceGroupPresentModeFlagsKHR) (common.VkResult, error)
}

type VkSurfaceFullScreenExclusiveInfoEXT C.VkSurfaceFullScreenExclusiveInfoEXT
type VkSurfaceCapabilitiesFullScreenExclusiveEXT C.VkSurfaceCapabilitiesFullScreenExclusiveEXT
type VkSurfaceFullScreenExclusiveWin32InfoEXT C.VkSurfaceFullScreenExclusiveWin32InfoEXT
type VkFullScreenExclusiveEXT C.VkFullScreenExclusiveEXT

type CLoader struct {
	coreLoader loader.Loader

	acquireFullScreenExclusiveMode        C.PFN_vkAcquireFullScreenExclusiveModeEXT
	getPhysicalDeviceSurfacePresentModes2 C.PFN_vkGetPhysicalDeviceSurfacePresentModes2EXT
	releaseFullScreenExclusiveMode        C.PFN_vkReleaseFullScreenExclusiveModeEXT

	getDeviceGroupSurfacePresentModes2 C.PFN_vkGetDeviceGroupSurfacePresentModes2EXT
}

func CreateLoaderFromCore(coreLoader loader.Loader) *CLoader {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	return &CLoader{
		coreLoader: coreLoader,

		acquireFullScreenExclusiveMode:        (C.PFN_vkAcquireFullScreenExclusiveModeEXT)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkAcquireFullScreenExclusiveModeEXT")))),
		getPhysicalDeviceSurfacePresentModes2: (C.PFN_vkGetPhysicalDeviceSurfacePresentModes2EXT)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkGetPhysicalDeviceSurfacePresentModes2EXT")))),
		releaseFullScreenExclusiveMode:        (C.PFN_vkReleaseFullScreenExclusiveModeEXT)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkReleaseFullScreenExclusiveModeEXT")))),

		getDeviceGroupSurfacePresentModes2: (C.PFN_vkGetDeviceGroupSurfacePresentModes2EXT)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkGetDeviceGroupSurfacePresentModes2EXT")))),
	}
}

func (c *CLoader) VkAcquireFullScreenExclusiveModeEXT(device loader.VkDevice, swapchain khr_swapchain_loader.VkSwapchainKHR) (common.VkResult, error) {
	if c.acquireFullScreenExclusiveMode == nil {
		panic("attempt to call extension method vkAcquireFullScreenExclusiveModeEXT when extension not present")
	}

	res := common.VkResult(C.cgoAcquireFullScreenExclusiveModeEXT(
		c.acquireFullScreenExclusiveMode,
		C.VkDevice(unsafe.Pointer(device)),
		C.VkSwapchainKHR(unsafe.Pointer(swapchain)),
	))

	return res, res.ToError()
}

func (c *CLoader) VkGetPhysicalDeviceSurfacePresentModes2EXT(physicalDevice loader.VkPhysicalDevice, pSurfaceInfo *khr_get_surface_capabilities2_loader.VkPhysicalDeviceSurfaceInfo2KHR, pPresentModeCount *loader.Uint32, pPresentModes *khr_surface_loader.VkPresentModeKHR) (common.VkResult, error) {
	if c.getPhysicalDeviceSurfacePresentModes2 == nil {
		panic("attempt to call extension method vkGetPhysicalDeviceSurfacePresentModes2EXT when extension not present")
	}

	res := common.VkResult(C.cgoGetPhysicalDeviceSurfacePresentModes2EXT(
		c.getPhysicalDeviceSurfacePresentModes2,
		C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)),
		(*C.VkPhysicalDeviceSurfaceInfo2KHR)(unsafe.Pointer(pSurfaceInfo)),
		(*C.uint32_t)(unsafe.Pointer(pPresentModeCount)),
		(*C.VkPresentModeKHR)(unsafe.Pointer(pPresentModes)),
	))
	return res, res.ToError()
}

func (c *CLoader) VkReleaseFullScreenExclusiveModeEXT(device loader.VkDevice, swapchain khr_swapchain_loader.VkSwapchainKHR) (common.VkResult, error) {
	if c.releaseFullScreenExclusiveMode == nil {
		panic("attempt to call extension method vkReleaseFullScreenExclusiveModeEXT when extension not present")
	}

	res := common.VkResult(C.cgoReleaseFullScreenExclusiveModeEXT(
		c.releaseFullScreenExclusiveMode,
		C.VkDevice(unsafe.Pointer(device)),
		C.VkSwapchainKHR(unsafe.Pointer(swapchain)),
	))
	return res, res.ToError()
}

func (c *CLoader) VkGetDeviceGroupSurfacePresentModes2EXT(device loader.VkDevice, pSurfaceInfo *khr_get_surface_capabilities2_loader.VkPhysicalDeviceSurfaceInfo2KHR, pModes *khr_device_group_loader.VkDeviceGroupPresentModeFlagsKHR) (common.VkResult, error) {
	if c.getDeviceGroupSurfacePresentModes2 == nil {
		panic("attempt to call extension method vkGetDeviceGroupSurfacePresentModes2EXT when extension not present")
	}

	res := common.VkResult(C.cgoGetDeviceGroupSurfacePresentModes2EXT(
		c.getDeviceGroupSurfacePresentModes2,
		C.VkDevice(unsafe.Pointer(device)),
		(*C.VkPhysicalDeviceSurfaceInfo2KHR)(unsafe.Pointer(pSurfaceInfo)),
		(*C.VkDeviceGroupPresentModeFlagsKHR)(unsafe.Pointer(pModes)),
	))
	return res, res.ToError()
}
