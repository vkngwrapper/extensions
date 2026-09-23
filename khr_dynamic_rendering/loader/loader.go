package khr_dynamic_rendering_loader

//go:generate mockgen -source loader.go -destination ../mocks/loader.go -package mock_dynamic_rendering

/*
#include <stdlib.h>
#include "../../vulkan/vulkan.h"

void cgoCmdBeginRenderingKHR(PFN_vkCmdBeginRenderingKHR fn, VkCommandBuffer commandBuffer, VkRenderingInfoKHR *pRenderingInfo) {
	fn(commandBuffer, pRenderingInfo);
}

void cgoCmdEndRenderingKHR(PFN_vkCmdEndRenderingKHR fn, VkCommandBuffer commandBuffer) {
	fn(commandBuffer);
}
*/
import "C"
import (
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3/loader"
)

type Loader interface {
	VkCmdBeginRenderingKHR(commandBuffer loader.VkCommandBuffer, pRenderingInfo *VkRenderingInfoKHR)
	VkCmdEndRenderingKHR(commandBuffer loader.VkCommandBuffer)
}

type VkRenderingInfoKHR C.VkRenderingInfoKHR
type VkRenderingAttachmentInfoKHR C.VkRenderingAttachmentInfoKHR
type VkPipelineRenderingCreateInfoKHR C.VkPipelineRenderingCreateInfoKHR
type VkCommandBufferInheritanceRenderingInfoKHR C.VkCommandBufferInheritanceRenderingInfoKHR
type VkPhysicalDeviceDynamicRenderingFeaturesKHR C.VkPhysicalDeviceDynamicRenderingFeaturesKHR

type CLoader struct {
	coreLoader loader.Loader

	beginRendering C.PFN_vkCmdBeginRenderingKHR
	endRendering   C.PFN_vkCmdEndRenderingKHR
}

func CreateLoaderFromCore(coreLoader loader.Loader) *CLoader {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	return &CLoader{
		coreLoader: coreLoader,

		beginRendering: (C.PFN_vkCmdBeginRenderingKHR)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkCmdBeginRenderingKHR")))),
		endRendering:   (C.PFN_vkCmdEndRenderingKHR)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkCmdEndRenderingKHR")))),
	}
}

func (d *CLoader) VkCmdBeginRenderingKHR(commandBuffer loader.VkCommandBuffer, pRenderingInfo *VkRenderingInfoKHR) {
	if d.beginRendering == nil {
		panic("attempt to call extension method vkCmdBeginRenderingKHR when extension not present")
	}

	C.cgoCmdBeginRenderingKHR(
		d.beginRendering,
		C.VkCommandBuffer(unsafe.Pointer(commandBuffer)),
		(*C.VkRenderingInfoKHR)(pRenderingInfo),
	)
}

func (d *CLoader) VkCmdEndRenderingKHR(commandBuffer loader.VkCommandBuffer) {
	if d.endRendering == nil {
		panic("attempt to call extension method vkCmdEndRenderingKHR when extension not present")
	}

	C.cgoCmdEndRenderingKHR(d.endRendering, C.VkCommandBuffer(unsafe.Pointer(commandBuffer)))
}

var _ Loader = &CLoader{}
