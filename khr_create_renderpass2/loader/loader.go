package khr_create_renderpass2_loader

/*
#include <stdlib.h>
#include "../../vulkan/vulkan.h"

void cgoCmdBeginRenderPass2KHR(PFN_vkCmdBeginRenderPass2KHR fn, VkCommandBuffer commandBuffer, VkRenderPassBeginInfo *pRenderPassBegin, VkSubpassBeginInfoKHR *pSubpassBegininfo) {
	fn(commandBuffer, pRenderPassBegin, pSubpassBegininfo);
}

void cgoCmdEndRenderPass2KHR(PFN_vkCmdEndRenderPass2KHR fn, VkCommandBuffer commandBuffer, VkSubpassEndInfoKHR *pSubpassEndInfo) {
	fn(commandBuffer, pSubpassEndInfo);
}

void cgoCmdNextSubpass2KHR(PFN_vkCmdNextSubpass2KHR fn, VkCommandBuffer commandBuffer, VkSubpassBeginInfoKHR *pSubpassBeginInfo, VkSubpassEndInfoKHR *pSubpassEndInfo) {
	fn(commandBuffer, pSubpassBeginInfo, pSubpassEndInfo);
}

VkResult cgoCreateRenderPass2KHR(PFN_vkCreateRenderPass2KHR fn, VkDevice device, VkRenderPassCreateInfo2KHR *pCreateInfo, VkAllocationCallbacks *pAllocator, VkRenderPass *pRenderPass) {
	return fn(device, pCreateInfo, pAllocator, pRenderPass);
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

//go:generate mockgen -source loader.go -destination ../mocks/loader.go -package mock_create_renderpass2

type Loader interface {
	VkCmdBeginRenderPass2KHR(commandBuffer loader.VkCommandBuffer, pRenderPassBegin *loader.VkRenderPassBeginInfo, pSubpassBeginInfo *VkSubpassBeginInfoKHR)
	VkCmdEndRenderPass2KHR(commandBuffer loader.VkCommandBuffer, pSubpassEndInfo *VkSubpassEndInfoKHR)
	VkCmdNextSubpass2KHR(commandBuffer loader.VkCommandBuffer, pSubpassBeginInfo *VkSubpassBeginInfoKHR, pSubpassEndInfo *VkSubpassEndInfoKHR)
	VkCreateRenderPass2KHR(device loader.VkDevice, pCreateInfo *VkRenderPassCreateInfo2KHR, pAllocator *loader.VkAllocationCallbacks, pRenderPass *loader.VkRenderPass) (common.VkResult, error)
}

type VkAttachmentDescription2KHR C.VkAttachmentDescription2KHR
type VkAttachmentReference2KHR C.VkAttachmentReference2KHR
type VkRenderPassCreateInfo2KHR C.VkRenderPassCreateInfo2KHR
type VkSubpassBeginInfoKHR C.VkSubpassBeginInfoKHR
type VkSubpassDependency2KHR C.VkSubpassDependency2KHR
type VkSubpassDescription2KHR C.VkSubpassDescription2KHR
type VkSubpassEndInfoKHR C.VkSubpassEndInfoKHR

type CLoader struct {
	coreLoader loader.Loader

	beginRenderPass  C.PFN_vkCmdBeginRenderPass2KHR
	endRenderPass    C.PFN_vkCmdEndRenderPass2KHR
	nextSubpass      C.PFN_vkCmdNextSubpass2KHR
	createRenderPass C.PFN_vkCreateRenderPass2KHR
}

func CreateLoaderFromCore(coreLoader loader.Loader) *CLoader {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	return &CLoader{
		coreLoader: coreLoader,

		beginRenderPass:  (C.PFN_vkCmdBeginRenderPass2KHR)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkCmdBeginRenderPass2KHR")))),
		endRenderPass:    (C.PFN_vkCmdEndRenderPass2KHR)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkCmdEndRenderPass2KHR")))),
		nextSubpass:      (C.PFN_vkCmdNextSubpass2KHR)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkCmdNextSubpass2KHR")))),
		createRenderPass: (C.PFN_vkCreateRenderPass2KHR)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkCreateRenderPass2KHR")))),
	}
}

func (d *CLoader) VkCmdBeginRenderPass2KHR(commandBuffer loader.VkCommandBuffer, pRenderPassBegin *loader.VkRenderPassBeginInfo, pSubpassBeginInfo *VkSubpassBeginInfoKHR) {
	if d.beginRenderPass == nil {
		panic("attempt to call extension method vkCmdBeginRenderPass2KHR when extension not present")
	}

	C.cgoCmdBeginRenderPass2KHR(
		d.beginRenderPass,
		C.VkCommandBuffer(unsafe.Pointer(commandBuffer)),
		(*C.VkRenderPassBeginInfo)(unsafe.Pointer(pRenderPassBegin)),
		(*C.VkSubpassBeginInfoKHR)(pSubpassBeginInfo),
	)
}

func (d *CLoader) VkCmdEndRenderPass2KHR(commandBuffer loader.VkCommandBuffer, pSubpassEndInfo *VkSubpassEndInfoKHR) {
	if d.endRenderPass == nil {
		panic("attempt to call extension method vkCmdEndRenderPass2KHR when extension not present")
	}

	C.cgoCmdEndRenderPass2KHR(
		d.endRenderPass,
		C.VkCommandBuffer(unsafe.Pointer(commandBuffer)),
		(*C.VkSubpassEndInfo)(pSubpassEndInfo),
	)
}

func (d *CLoader) VkCmdNextSubpass2KHR(commandBuffer loader.VkCommandBuffer, pSubpassBeginInfo *VkSubpassBeginInfoKHR, pSubpassEndInfo *VkSubpassEndInfoKHR) {
	if d.nextSubpass == nil {
		panic("attempt to call extension method vkCmdNextSubpass2KHR when extension not present")
	}

	C.cgoCmdNextSubpass2KHR(
		d.nextSubpass,
		C.VkCommandBuffer(unsafe.Pointer(commandBuffer)),
		(*C.VkSubpassBeginInfoKHR)(pSubpassBeginInfo),
		(*C.VkSubpassEndInfoKHR)(pSubpassEndInfo),
	)
}

func (d *CLoader) VkCreateRenderPass2KHR(device loader.VkDevice, pCreateInfo *VkRenderPassCreateInfo2KHR, pAllocator *loader.VkAllocationCallbacks, pRenderPass *loader.VkRenderPass) (common.VkResult, error) {
	if d.createRenderPass == nil {
		panic("attempt to call extension method vkCreateRenderPass2KHR when extension not present")
	}

	res := common.VkResult(C.cgoCreateRenderPass2KHR(
		d.createRenderPass,
		C.VkDevice(unsafe.Pointer(device)),
		(*C.VkRenderPassCreateInfo2KHR)(pCreateInfo),
		(*C.VkAllocationCallbacks)(unsafe.Pointer(pAllocator)),
		(*C.VkRenderPass)(unsafe.Pointer(pRenderPass)),
	))

	return res, res.ToError()
}
