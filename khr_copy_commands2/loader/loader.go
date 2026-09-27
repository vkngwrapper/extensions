package khr_copy_commands2_loader

/*
#include <stdlib.h>
#include "../../vulkan/vulkan.h"

void cgoCmdBlitImage2KHR(PFN_vkCmdBlitImage2KHR fn, VkCommandBuffer commandBuffer, VkBlitImageInfo2KHR *pBlitImageInfo) {
	fn(commandBuffer, pBlitImageInfo);
}

void cgoCmdCopyBuffer2KHR(PFN_vkCmdCopyBuffer2KHR fn, VkCommandBuffer commandBuffer, VkCopyBufferInfo2KHR *pCopyBufferInfo) {
	fn(commandBuffer, pCopyBufferInfo);
}

void cgoCmdCopyBufferToImage2KHR(PFN_vkCmdCopyBufferToImage2KHR fn, VkCommandBuffer commandBuffer, VkCopyBufferToImageInfo2KHR *pCopyBufferToImageInfo) {
	fn(commandBuffer, pCopyBufferToImageInfo);
}

void cgoCmdCopyImage2KHR(PFN_vkCmdCopyImage2KHR fn, VkCommandBuffer commandBuffer, VkCopyImageInfo2KHR *pCopyImageInfo) {
	fn(commandBuffer, pCopyImageInfo);
}

void cgoCmdCopyImageToBuffer2KHR(PFN_vkCmdCopyImageToBuffer2KHR fn, VkCommandBuffer commandBuffer, VkCopyImageToBufferInfo2KHR *pCopyImageToBufferInfo) {
	fn(commandBuffer, pCopyImageToBufferInfo);
}

void cgoCmdResolveImage2KHR(PFN_vkCmdResolveImage2KHR fn, VkCommandBuffer commandBuffer, VkResolveImageInfo2KHR *pResolveImageInfo) {
	fn(commandBuffer, pResolveImageInfo);
}

*/
import "C"
import (
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3/loader"
)

//go:generate go run go.uber.org/mock/mockgen -source loader.go -destination ../mocks/loader.go -package mock_copy_commands2

type Loader interface {
	VkCmdBlitImage2KHR(commandBuffer loader.VkCommandBuffer, pBlitImageInfo *VkBlitImageInfo2KHR)
	VkCmdCopyBuffer2KHR(commandBuffer loader.VkCommandBuffer, pCopyBufferInfo *VkCopyBufferInfo2KHR)
	VkCmdCopyBufferToImage2KHR(commandBuffer loader.VkCommandBuffer, pCopyBufferToImageInfo *VkCopyBufferToImageInfo2KHR)
	VkCmdCopyImage2KHR(commandBuffer loader.VkCommandBuffer, pCopyImageInfo *VkCopyImageInfo2KHR)
	VkCmdCopyImageToBuffer2KHR(commandBuffer loader.VkCommandBuffer, pCopyImageToBufferInfo *VkCopyImageToBufferInfo2KHR)
	VkCmdResolveImage2KHR(commandBuffer loader.VkCommandBuffer, pResolveImageInfo *VkResolveImageInfo2KHR)
}

type VkBlitImageInfo2KHR C.VkBlitImageInfo2KHR
type VkBufferCopy2KHR C.VkBufferCopy2KHR
type VkBufferImageCopy2KHR C.VkBufferImageCopy2KHR
type VkCopyBufferInfo2KHR C.VkCopyBufferInfo2KHR
type VkCopyBufferToImageInfo2KHR C.VkCopyBufferToImageInfo2KHR
type VkCopyImageInfo2KHR C.VkCopyImageInfo2KHR
type VkCopyImageToBufferInfo2KHR C.VkCopyImageToBufferInfo2KHR
type VkImageBlit2KHR C.VkImageBlit2KHR
type VkImageCopy2KHR C.VkImageCopy2KHR
type VkImageResolve2KHR C.VkImageResolve2KHR
type VkResolveImageInfo2KHR C.VkResolveImageInfo2KHR

type CLoader struct {
	coreLoader loader.Loader

	cmdBlitImage2         C.PFN_vkCmdBlitImage2KHR
	cmdCopyBuffer2        C.PFN_vkCmdCopyBuffer2KHR
	cmdCopyBufferToImage2 C.PFN_vkCmdCopyBufferToImage2KHR
	cmdCopyImage2         C.PFN_vkCmdCopyImage2KHR
	cmdCopyImageToBuffer2 C.PFN_vkCmdCopyImageToBuffer2KHR
	cmdResolveImage2      C.PFN_vkCmdResolveImage2KHR
}

var _ Loader = &CLoader{}

func CreateLoaderFromCore(coreLoader loader.Loader) *CLoader {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	return &CLoader{
		coreLoader: coreLoader,

		cmdBlitImage2:         (C.PFN_vkCmdBlitImage2KHR)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkCmdBlitImage2KHR")))),
		cmdCopyBuffer2:        (C.PFN_vkCmdCopyBuffer2KHR)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkCmdCopyBuffer2KHR")))),
		cmdCopyBufferToImage2: (C.PFN_vkCmdCopyBufferToImage2KHR)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkCmdCopyBufferToImage2KHR")))),
		cmdCopyImage2:         (C.PFN_vkCmdCopyImage2KHR)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkCmdCopyImage2KHR")))),
		cmdCopyImageToBuffer2: (C.PFN_vkCmdCopyImageToBuffer2KHR)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkCmdCopyImageToBuffer2KHR")))),
		cmdResolveImage2:      (C.PFN_vkCmdResolveImage2KHR)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkCmdResolveImage2KHR")))),
	}
}

func (c *CLoader) VkCmdBlitImage2KHR(commandBuffer loader.VkCommandBuffer, pBlitImageInfo *VkBlitImageInfo2KHR) {
	if c.cmdBlitImage2 == nil {
		panic("attempt to call extension method vkCmdBlitImage2KHR when extension not present")
	}

	C.cgoCmdBlitImage2KHR(c.cmdBlitImage2,
		C.VkCommandBuffer(unsafe.Pointer(commandBuffer)),
		(*C.VkBlitImageInfo2KHR)(unsafe.Pointer(pBlitImageInfo)))
}

func (c *CLoader) VkCmdCopyBuffer2KHR(commandBuffer loader.VkCommandBuffer, pCopyBufferInfo *VkCopyBufferInfo2KHR) {
	if c.cmdCopyBuffer2 == nil {
		panic("attempt to call extension method vkCmdCopyBuffer2KHR when extension not present")
	}

	C.cgoCmdCopyBuffer2KHR(c.cmdCopyBuffer2,
		C.VkCommandBuffer(unsafe.Pointer(commandBuffer)),
		(*C.VkCopyBufferInfo2KHR)(unsafe.Pointer(pCopyBufferInfo)))
}

func (c *CLoader) VkCmdCopyBufferToImage2KHR(commandBuffer loader.VkCommandBuffer, pCopyBufferToImageInfo *VkCopyBufferToImageInfo2KHR) {
	if c.cmdCopyBufferToImage2 == nil {
		panic("attempt to call extension method vkCmdCopyBufferToImage2KHR when extension not present")
	}

	C.cgoCmdCopyBufferToImage2KHR(c.cmdCopyBufferToImage2,
		C.VkCommandBuffer(unsafe.Pointer(commandBuffer)),
		(*C.VkCopyBufferToImageInfo2KHR)(unsafe.Pointer(pCopyBufferToImageInfo)))
}

func (c *CLoader) VkCmdCopyImage2KHR(commandBuffer loader.VkCommandBuffer, pCopyImageInfo *VkCopyImageInfo2KHR) {
	if c.cmdCopyImage2 == nil {
		panic("attempt to call extension method vkCmdCopyImage2KHR when extension not present")
	}

	C.cgoCmdCopyImage2KHR(c.cmdCopyImage2,
		C.VkCommandBuffer(unsafe.Pointer(commandBuffer)),
		(*C.VkCopyImageInfo2KHR)(unsafe.Pointer(pCopyImageInfo)))
}

func (c *CLoader) VkCmdCopyImageToBuffer2KHR(commandBuffer loader.VkCommandBuffer, pCopyImageToBufferInfo *VkCopyImageToBufferInfo2KHR) {
	if c.cmdCopyImageToBuffer2 == nil {
		panic("attempt to call extension method vkCmdCopyImageToBuffer2 when extension not present")
	}

	C.cgoCmdCopyImageToBuffer2KHR(c.cmdCopyImageToBuffer2,
		C.VkCommandBuffer(unsafe.Pointer(commandBuffer)),
		(*C.VkCopyImageToBufferInfo2KHR)(unsafe.Pointer(pCopyImageToBufferInfo)),
	)
}

func (c *CLoader) VkCmdResolveImage2KHR(commandBuffer loader.VkCommandBuffer, pResolveImageInfo *VkResolveImageInfo2KHR) {
	if c.cmdResolveImage2 == nil {
		panic("attempt to call extension method vkCmdResolveImage2KHR when extension not present")
	}

	C.cgoCmdResolveImage2KHR(c.cmdResolveImage2,
		C.VkCommandBuffer(unsafe.Pointer(commandBuffer)),
		(*C.VkResolveImageInfo2KHR)(unsafe.Pointer(pResolveImageInfo)),
	)
}
