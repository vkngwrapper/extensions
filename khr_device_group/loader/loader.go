package khr_device_group_loader

/*
#include <stdlib.h>
#include "../../vulkan/vulkan.h"

void cgoCmdDispatchBaseKHR(PFN_vkCmdDispatchBaseKHR fn, VkCommandBuffer commandBuffer, uint32_t baseGroupX, uint32_t baseGroupY, uint32_t baseGroupZ, uint32_t groupCountX, uint32_t groupCountY, uint32_t groupCountZ) {
	fn(commandBuffer, baseGroupX, baseGroupY, baseGroupZ, groupCountX, groupCountY, groupCountZ);
}

void cgoCmdSetDeviceMaskKHR(PFN_vkCmdSetDeviceMaskKHR fn, VkCommandBuffer commandBuffer, uint32_t deviceMask) {
	fn(commandBuffer, deviceMask);
}

void cgoGetDeviceGroupPeerMemoryFeaturesKHR(PFN_vkGetDeviceGroupPeerMemoryFeatures fn, VkDevice device, uint32_t heapIndex, uint32_t localDeviceIndex, uint32_t remoteDeviceIndex, VkPeerMemoryFeatureFlags *pPeerMemoryFeatures) {
	fn(device, heapIndex, localDeviceIndex, remoteDeviceIndex, pPeerMemoryFeatures);
}

VkResult cgoGetDeviceGroupPresentCapabilitiesKHRDeviceGroup(PFN_vkGetDeviceGroupPresentCapabilitiesKHR fn, VkDevice device, VkDeviceGroupPresentCapabilitiesKHR *pDeviceGroupPresentCapabilites) {
	return fn(device, pDeviceGroupPresentCapabilites);
}

VkResult cgoGetDeviceGroupSurfacePresentModesKHRDeviceGroup(PFN_vkGetDeviceGroupSurfacePresentModesKHR fn, VkDevice device, VkSurfaceKHR surface, VkDeviceGroupPresentModeFlagsKHR *pModes) {
	return fn(device, surface, pModes);
}

VkResult cgoGetPhysicalDevicePresentRectanglesKHRDeviceGroup(PFN_vkGetPhysicalDevicePresentRectanglesKHR fn, VkPhysicalDevice physicalDevice, VkSurfaceKHR surface, uint32_t *pRectCount, VkRect2D *pRects) {
	return fn(physicalDevice, surface, pRectCount, pRects);
}

VkResult cgoAcquireNextImage2KHRDeviceGroup(PFN_vkAcquireNextImage2KHR fn, VkDevice device, VkAcquireNextImageInfoKHR *pAcquireInfo, uint32_t *pImageIndex) {
	return fn(device, pAcquireInfo, pImageIndex);
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

//go:generate mockgen -source loader.go -destination ../mocks/loader.go -package mock_device_group

type Loader interface {
	VkCmdDispatchBaseKHR(commandBuffer loader.VkCommandBuffer, baseGroupX, baseGroupY, baseGroupZ, groupCountX, groupCountY, groupCountZ loader.Uint32)
	VkCmdSetDeviceMaskKHR(commandBuffer loader.VkCommandBuffer, deviceMask loader.Uint32)
	VkGetDeviceGroupPeerMemoryFeaturesKHR(device loader.VkDevice, heapIndex, localDeviceIndex, remoteDeviceIndex loader.Uint32, pPeerMemoryFeatures *VkPeerMemoryFeatureFlagsKHR)

	VkGetDeviceGroupPresentCapabilitiesKHR(device loader.VkDevice, pDeviceGroupPresentCapabilities *VkDeviceGroupPresentCapabilitiesKHR) (common.VkResult, error)
	VkGetDeviceGroupSurfacePresentModesKHR(device loader.VkDevice, surface khr_surface_driver.VkSurfaceKHR, pModes *VkDeviceGroupPresentModeFlagsKHR) (common.VkResult, error)
	VkGetPhysicalDevicePresentRectanglesKHR(physicalDevice loader.VkPhysicalDevice, surface khr_surface_driver.VkSurfaceKHR, pRectCount *loader.Uint32, pRects *loader.VkRect2D) (common.VkResult, error)

	VkAcquireNextImage2KHR(device loader.VkDevice, pAcquireInfo *VkAcquireNextImageInfoKHR, pImageIndex *loader.Uint32) (common.VkResult, error)
}

type VkDeviceGroupBindSparseInfoKHR C.VkDeviceGroupBindSparseInfoKHR
type VkDeviceGroupCommandBufferBeginInfoKHR C.VkDeviceGroupCommandBufferBeginInfoKHR
type VkMemoryAllocateFlagsInfoKHR C.VkMemoryAllocateFlagsInfoKHR
type VkDeviceGroupRenderPassBeginInfoKHR C.VkDeviceGroupRenderPassBeginInfoKHR
type VkDeviceGroupSubmitInfoKHR C.VkDeviceGroupSubmitInfoKHR
type VkBindBufferMemoryDeviceGroupInfoKHR C.VkBindBufferMemoryDeviceGroupInfoKHR
type VkBindImageMemoryDeviceGroupInfoKHR C.VkBindImageMemoryDeviceGroupInfoKHR
type VkDeviceGroupPresentCapabilitiesKHR C.VkDeviceGroupPresentCapabilitiesKHR
type VkAcquireNextImageInfoKHR C.VkAcquireNextImageInfoKHR
type VkBindImageMemorySwapchainInfoKHR C.VkBindImageMemorySwapchainInfoKHR
type VkImageSwapchainCreateInfoKHR C.VkImageSwapchainCreateInfoKHR
type VkDeviceGroupPresentInfoKHR C.VkDeviceGroupPresentInfoKHR
type VkDeviceGroupSwapchainCreateInfoKHR C.VkDeviceGroupSwapchainCreateInfoKHR
type VkMemoryAllocateFlagsKHR C.VkMemoryAllocateFlagsKHR
type VkPeerMemoryFeatureFlagsKHR C.VkPeerMemoryFeatureFlagsKHR
type VkDeviceGroupPresentModeFlagsKHR C.VkDeviceGroupPresentModeFlagsKHR

type CLoader struct {
	coreLoader loader.Loader

	cmdDispatchBase                    C.PFN_vkCmdDispatchBaseKHR
	cmdSetDeviceMask                   C.PFN_vkCmdSetDeviceMaskKHR
	getDeviceGroupPeerMemoryFeatures   C.PFN_vkGetDeviceGroupPeerMemoryFeatures
	getDeviceGroupPresentCapabilities  C.PFN_vkGetDeviceGroupPresentCapabilitiesKHR
	getDeviceGroupSurfacePresentModes  C.PFN_vkGetDeviceGroupSurfacePresentModesKHR
	getPhysicalDevicePresentRectangles C.PFN_vkGetPhysicalDevicePresentRectanglesKHR
	acquireNextImage                   C.PFN_vkAcquireNextImage2KHR
}

func CreateLoaderFromCore(coreLoader loader.Loader) *CLoader {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	return &CLoader{
		coreLoader: coreLoader,

		cmdDispatchBase:                    (C.PFN_vkCmdDispatchBaseKHR)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkCmdDispatchBaseKHR")))),
		cmdSetDeviceMask:                   (C.PFN_vkCmdSetDeviceMaskKHR)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkCmdSetDeviceMaskKHR")))),
		getDeviceGroupPeerMemoryFeatures:   (C.PFN_vkGetDeviceGroupPeerMemoryFeatures)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkGetDeviceGroupPeerMemoryFeatures")))),
		getDeviceGroupPresentCapabilities:  (C.PFN_vkGetDeviceGroupPresentCapabilitiesKHR)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkGetDeviceGroupPresentCapabilitiesKHR")))),
		getDeviceGroupSurfacePresentModes:  (C.PFN_vkGetDeviceGroupSurfacePresentModesKHR)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkGetDeviceGroupSurfacePresentModesKHR")))),
		getPhysicalDevicePresentRectangles: (C.PFN_vkGetPhysicalDevicePresentRectanglesKHR)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkGetPhysicalDevicePresentRectanglesKHR")))),
		acquireNextImage:                   (C.PFN_vkAcquireNextImage2KHR)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkAcquireNextImage2KHR")))),
	}
}

func (d *CLoader) VkCmdDispatchBaseKHR(commandBuffer loader.VkCommandBuffer, baseGroupX, baseGroupY, baseGroupZ, groupCountX, groupCountY, groupCountZ loader.Uint32) {
	if d.cmdDispatchBase == nil {
		panic("attempt to call extension method vkCmdDispatchBaseKHR when extension not present")
	}

	C.cgoCmdDispatchBaseKHR(d.cmdDispatchBase,
		C.VkCommandBuffer(unsafe.Pointer(commandBuffer)),
		C.uint32_t(baseGroupX),
		C.uint32_t(baseGroupY),
		C.uint32_t(baseGroupZ),
		C.uint32_t(groupCountX),
		C.uint32_t(groupCountY),
		C.uint32_t(groupCountZ),
	)
}

func (d *CLoader) VkCmdSetDeviceMaskKHR(commandBuffer loader.VkCommandBuffer, deviceMask loader.Uint32) {
	if d.cmdSetDeviceMask == nil {
		panic("attempt to call extension method vkCmdSetDeviceMaskKHR when extension not present")
	}

	C.cgoCmdSetDeviceMaskKHR(d.cmdSetDeviceMask,
		C.VkCommandBuffer(unsafe.Pointer(commandBuffer)),
		C.uint32_t(deviceMask),
	)
}

func (d *CLoader) VkGetDeviceGroupPeerMemoryFeaturesKHR(device loader.VkDevice, heapIndex, localDeviceIndex, remoteDeviceIndex loader.Uint32, pPeerMemoryFeatures *VkPeerMemoryFeatureFlagsKHR) {
	if d.getDeviceGroupPeerMemoryFeatures == nil {
		panic("attempt to call extension method vkGetDeviceGroupPeerMemoryFeaturesKHR when extension not present")
	}

	C.cgoGetDeviceGroupPeerMemoryFeaturesKHR(d.getDeviceGroupPeerMemoryFeatures,
		C.VkDevice(unsafe.Pointer(device)),
		C.uint32_t(heapIndex),
		C.uint32_t(localDeviceIndex),
		C.uint32_t(remoteDeviceIndex),
		(*C.VkPeerMemoryFeatureFlagsKHR)(pPeerMemoryFeatures),
	)
}

func (d *CLoader) VkGetDeviceGroupPresentCapabilitiesKHR(device loader.VkDevice, pDeviceGroupPresentCapabilities *VkDeviceGroupPresentCapabilitiesKHR) (common.VkResult, error) {
	if d.getDeviceGroupPresentCapabilities == nil {
		panic("attempt to call extension method vkGetDeviceGroupPresentCapabilitiesKHR when extension and/or prerequisite not present")
	}

	res := common.VkResult(C.cgoGetDeviceGroupPresentCapabilitiesKHRDeviceGroup(
		d.getDeviceGroupPresentCapabilities,
		C.VkDevice(unsafe.Pointer(device)),
		(*C.VkDeviceGroupPresentCapabilitiesKHR)(pDeviceGroupPresentCapabilities),
	))

	return res, res.ToError()
}

func (d *CLoader) VkGetDeviceGroupSurfacePresentModesKHR(device loader.VkDevice, surface khr_surface_driver.VkSurfaceKHR, pModes *VkDeviceGroupPresentModeFlagsKHR) (common.VkResult, error) {
	if d.getDeviceGroupSurfacePresentModes == nil {
		panic("attempt to call extension method vkGetDeviceGroupSurfacePresentModesKHR when extension and/or prerequisite not present")
	}

	res := common.VkResult(C.cgoGetDeviceGroupSurfacePresentModesKHRDeviceGroup(
		d.getDeviceGroupSurfacePresentModes,
		C.VkDevice(unsafe.Pointer(device)),
		C.VkSurfaceKHR(unsafe.Pointer(surface)),
		(*C.VkDeviceGroupPresentModeFlagsKHR)(pModes),
	))

	return res, res.ToError()
}

func (d *CLoader) VkGetPhysicalDevicePresentRectanglesKHR(physicalDevice loader.VkPhysicalDevice, surface khr_surface_driver.VkSurfaceKHR, pRectCount *loader.Uint32, pRects *loader.VkRect2D) (common.VkResult, error) {
	if d.getPhysicalDevicePresentRectangles == nil {
		panic("attempt to call extension method vkGetPhysicalDevicePresentRectanglesKHR when extension and/or prerequisite not present")
	}

	res := common.VkResult(C.cgoGetPhysicalDevicePresentRectanglesKHRDeviceGroup(
		d.getPhysicalDevicePresentRectangles,
		C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)),
		C.VkSurfaceKHR(unsafe.Pointer(surface)),
		(*C.uint32_t)(pRectCount),
		(*C.VkRect2D)(unsafe.Pointer(pRects)),
	))
	return res, res.ToError()
}

func (d *CLoader) VkAcquireNextImage2KHR(device loader.VkDevice, pAcquireInfo *VkAcquireNextImageInfoKHR, pImageIndex *loader.Uint32) (common.VkResult, error) {
	if d.acquireNextImage == nil {
		panic("attempt to call extension method vkAcquireNextImage2KHR when extension and/or prerequisite not present")
	}

	res := common.VkResult(C.cgoAcquireNextImage2KHRDeviceGroup(
		d.acquireNextImage,
		C.VkDevice(unsafe.Pointer(device)),
		(*C.VkAcquireNextImageInfoKHR)(pAcquireInfo),
		(*C.uint32_t)(pImageIndex),
	))
	return res, res.ToError()
}
