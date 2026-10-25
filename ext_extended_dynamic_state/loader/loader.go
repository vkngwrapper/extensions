package ext_extended_dynamic_state_loader

/*
#include <stdlib.h>
#include "../../vulkan/vulkan.h"

void cgoCmdBindVertexBuffers2EXT(PFN_vkCmdBindVertexBuffers2EXT fn,
	VkCommandBuffer                             commandBuffer,
    uint32_t                                    firstBinding,
    uint32_t                                    bindingCount,
    const VkBuffer*                             pBuffers,
    const VkDeviceSize*                         pOffsets,
    const VkDeviceSize*                         pSizes,
    const VkDeviceSize*                         pStrides) {
	fn(commandBuffer, firstBinding, bindingCount, pBuffers, pOffsets, pSizes, pStrides);
}

void cgoCmdSetCullModeEXT(PFN_vkCmdSetCullModeEXT fn,
    VkCommandBuffer                             commandBuffer,
    VkCullModeFlags                             cullMode) {

	fn(commandBuffer, cullMode);
}

void cgoCmdSetDepthBoundsTestEnableEXT(PFN_vkCmdSetDepthBoundsTestEnableEXT fn,
    VkCommandBuffer                             commandBuffer,
    VkBool32                                    depthBoundsTestEnable) {

	fn(commandBuffer, depthBoundsTestEnable);
}

void cgoCmdSetDepthCompareOpEXT(PFN_vkCmdSetDepthCompareOpEXT fn,
    VkCommandBuffer                             commandBuffer,
    VkCompareOp                                 depthCompareOp) {

	fn(commandBuffer, depthCompareOp);
}

void cgoCmdSetDepthTestEnableEXT(PFN_vkCmdSetDepthTestEnableEXT fn,
    VkCommandBuffer                             commandBuffer,
    VkBool32                                    depthTestEnable) {

	fn(commandBuffer, depthTestEnable);
}

void cgoCmdSetDepthWriteEnableEXT(PFN_vkCmdSetDepthWriteEnableEXT fn,
    VkCommandBuffer                             commandBuffer,
    VkBool32                                    depthWriteEnable) {

	fn(commandBuffer, depthWriteEnable);
}

void cgoCmdSetFrontFaceEXT(PFN_vkCmdSetFrontFaceEXT fn,
    VkCommandBuffer                             commandBuffer,
    VkFrontFace                                 frontFace) {

	fn(commandBuffer, frontFace);
}

void cgoCmdSetPrimitiveTopologyEXT(PFN_vkCmdSetPrimitiveTopologyEXT fn,
    VkCommandBuffer                             commandBuffer,
    VkPrimitiveTopology                         primitiveTopology) {

	fn(commandBuffer, primitiveTopology);
}

void cgoCmdSetScissorWithCountEXT(PFN_vkCmdSetScissorWithCountEXT fn,
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    scissorCount,
    const VkRect2D*                             pScissors) {

	fn(commandBuffer, scissorCount, pScissors);
}

void cgoCmdSetStencilOpEXT(PFN_vkCmdSetStencilOpEXT fn,
    VkCommandBuffer                             commandBuffer,
    VkStencilFaceFlags                          faceMask,
    VkStencilOp                                 failOp,
    VkStencilOp                                 passOp,
    VkStencilOp                                 depthFailOp,
    VkCompareOp                                 compareOp) {

	fn(commandBuffer, faceMask, failOp, passOp, depthFailOp, compareOp);
}

void cgoCmdSetStencilTestEnableEXT(PFN_vkCmdSetStencilTestEnableEXT fn,
    VkCommandBuffer                             commandBuffer,
    VkBool32                                    stencilTestEnable) {

	fn(commandBuffer, stencilTestEnable);
}

void cgoCmdSetViewportWithCountEXT(PFN_vkCmdSetViewportWithCountEXT fn,
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    viewportCount,
    const VkViewport*                           pViewports) {

	fn(commandBuffer, viewportCount, pViewports);
}
*/
import "C"
import (
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3/loader"
)

//go:generate go run go.uber.org/mock/mockgen -source loader.go -destination ../mocks/loader.go -package mock_extended_dynamic_state

type Loader interface {
	VkCmdBindVertexBuffers2(commandBuffer loader.VkCommandBuffer, firstBinding loader.Uint32, bindingCount loader.Uint32, pBuffers *loader.VkBuffer, pOffsets *loader.VkDeviceSize, pSizes *loader.VkDeviceSize, pStrides *loader.VkDeviceSize)
	VkCmdSetCullMode(commandBuffer loader.VkCommandBuffer, cullMode loader.VkCullModeFlags)
	VkCmdSetDepthBoundsTestEnable(commandBuffer loader.VkCommandBuffer, depthBoundsTestEnable loader.VkBool32)
	VkCmdSetDepthCompareOp(commandBuffer loader.VkCommandBuffer, depthCompareOp loader.VkCompareOp)
	VkCmdSetDepthTestEnable(commandBuffer loader.VkCommandBuffer, depthTestEnable loader.VkBool32)
	VkCmdSetDepthWriteEnable(commandBuffer loader.VkCommandBuffer, depthWriteEnable loader.VkBool32)
	VkCmdSetFrontFace(commandBuffer loader.VkCommandBuffer, frontFace loader.VkFrontFace)
	VkCmdSetPrimitiveTopology(commandBuffer loader.VkCommandBuffer, primitiveTopology loader.VkPrimitiveTopology)
	VkCmdSetScissorWithCount(commandBuffer loader.VkCommandBuffer, scissorCount loader.Uint32, pScissors *loader.VkRect2D)
	VkCmdSetStencilOp(commandBuffer loader.VkCommandBuffer, faceMask loader.VkStencilFaceFlags, failOp loader.VkStencilOp, passOp loader.VkStencilOp, depthFailOp loader.VkStencilOp, compareOp loader.VkCompareOp)
	VkCmdSetStencilTestEnable(commandBuffer loader.VkCommandBuffer, stencilTestEnable loader.VkBool32)
	VkCmdSetViewportWithCount(commandBuffer loader.VkCommandBuffer, viewportCount loader.Uint32, pViewports *loader.VkViewport)
}

type VkPhysicalDeviceExtendedDynamicStateFeaturesEXT C.VkPhysicalDeviceExtendedDynamicStateFeaturesEXT

type CLoader struct {
	coreLoader loader.Loader

	cmdBindVertexBuffers2       C.PFN_vkCmdBindVertexBuffers2EXT
	cmdSetCullMode              C.PFN_vkCmdSetCullModeEXT
	cmdSetDepthBoundsTestEnable C.PFN_vkCmdSetDepthBoundsTestEnableEXT
	cmdSetDepthCompareOp        C.PFN_vkCmdSetDepthCompareOpEXT
	cmdSetDepthTestEnable       C.PFN_vkCmdSetDepthTestEnableEXT
	cmdSetDepthWriteEnable      C.PFN_vkCmdSetDepthWriteEnableEXT
	cmdSetFrontFace             C.PFN_vkCmdSetFrontFaceEXT
	cmdSetPrimitiveTopology     C.PFN_vkCmdSetPrimitiveTopologyEXT
	cmdSetScissorWithCount      C.PFN_vkCmdSetScissorWithCountEXT
	cmdSetStencilOp             C.PFN_vkCmdSetStencilOpEXT
	cmdSetStencilTestEnable     C.PFN_vkCmdSetStencilTestEnableEXT
	cmdSetViewportWithCount     C.PFN_vkCmdSetViewportWithCountEXT
}

func CreateLoaderFromCore(coreLoader loader.Loader) *CLoader {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	return &CLoader{
		coreLoader: coreLoader,

		cmdBindVertexBuffers2:       (C.PFN_vkCmdBindVertexBuffers2EXT)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkCmdBindVertexBuffers2EXT")))),
		cmdSetCullMode:              (C.PFN_vkCmdSetCullModeEXT)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkCmdSetCullModeEXT")))),
		cmdSetDepthBoundsTestEnable: (C.PFN_vkCmdSetDepthBoundsTestEnableEXT)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkCmdSetDepthBoundsTestEnableEXT")))),
		cmdSetDepthCompareOp:        (C.PFN_vkCmdSetDepthCompareOpEXT)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkCmdSetDepthCompareOpEXT")))),
		cmdSetDepthTestEnable:       (C.PFN_vkCmdSetDepthTestEnableEXT)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkCmdSetDepthTestEnableEXT")))),
		cmdSetDepthWriteEnable:      (C.PFN_vkCmdSetDepthWriteEnableEXT)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkCmdSetDepthWriteEnableEXT")))),
		cmdSetFrontFace:             (C.PFN_vkCmdSetFrontFaceEXT)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkCmdSetFrontFaceEXT")))),
		cmdSetPrimitiveTopology:     (C.PFN_vkCmdSetPrimitiveTopologyEXT)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkCmdSetPrimitiveTopologyEXT")))),
		cmdSetScissorWithCount:      (C.PFN_vkCmdSetScissorWithCountEXT)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkCmdSetScissorWithCountEXT")))),
		cmdSetStencilOp:             (C.PFN_vkCmdSetStencilOpEXT)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkCmdSetStencilOpEXT")))),
		cmdSetStencilTestEnable:     (C.PFN_vkCmdSetStencilTestEnableEXT)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkCmdSetStencilTestEnableEXT")))),
		cmdSetViewportWithCount:     (C.PFN_vkCmdSetViewportWithCountEXT)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkCmdSetViewportWithCountEXT")))),
	}
}

var _ Loader = &CLoader{}

func (c *CLoader) VkCmdBindVertexBuffers2(commandBuffer loader.VkCommandBuffer, firstBinding loader.Uint32, bindingCount loader.Uint32, pBuffers *loader.VkBuffer, pOffsets *loader.VkDeviceSize, pSizes *loader.VkDeviceSize, pStrides *loader.VkDeviceSize) {
	if c.cmdBindVertexBuffers2 == nil {
		panic("attempt to call extension method vkCmdBindVertexBuffers2EXT when extension not present")
	}

	C.cgoCmdBindVertexBuffers2EXT(c.cmdBindVertexBuffers2, (C.VkCommandBuffer)(unsafe.Pointer(commandBuffer)), C.uint32_t(firstBinding), C.uint32_t(bindingCount), (*C.VkBuffer)(unsafe.Pointer(pBuffers)), (*C.VkDeviceSize)(unsafe.Pointer(pOffsets)), (*C.VkDeviceSize)(unsafe.Pointer(pSizes)), (*C.VkDeviceSize)(unsafe.Pointer(pStrides)))
}

func (c *CLoader) VkCmdSetCullMode(commandBuffer loader.VkCommandBuffer, cullMode loader.VkCullModeFlags) {
	if c.cmdSetCullMode == nil {
		panic("attempt to call extension method vkCmdSetCullModeEXT when extension not present")
	}

	C.cgoCmdSetCullModeEXT(c.cmdSetCullMode, (C.VkCommandBuffer)(unsafe.Pointer(commandBuffer)), C.VkCullModeFlags(cullMode))
}

func (c *CLoader) VkCmdSetDepthBoundsTestEnable(commandBuffer loader.VkCommandBuffer, depthBoundsTestEnable loader.VkBool32) {
	if c.cmdSetDepthBoundsTestEnable == nil {
		panic("attempt to call extension method vkCmdSetDepthBoundsTestEnableEXT when extension not present")
	}

	C.cgoCmdSetDepthBoundsTestEnableEXT(c.cmdSetDepthBoundsTestEnable, (C.VkCommandBuffer)(unsafe.Pointer(commandBuffer)), C.VkBool32(depthBoundsTestEnable))
}

func (c *CLoader) VkCmdSetDepthCompareOp(commandBuffer loader.VkCommandBuffer, depthCompareOp loader.VkCompareOp) {
	if c.cmdSetDepthCompareOp == nil {
		panic("attempt to call extension method vkCmdSetDepthCompareOpEXT when extension not present")
	}

	C.cgoCmdSetDepthCompareOpEXT(c.cmdSetDepthCompareOp, (C.VkCommandBuffer)(unsafe.Pointer(commandBuffer)), C.VkCompareOp(depthCompareOp))
}

func (c *CLoader) VkCmdSetDepthTestEnable(commandBuffer loader.VkCommandBuffer, depthTestEnable loader.VkBool32) {
	if c.cmdSetDepthTestEnable == nil {
		panic("attempt to call extension method vkCmdSetDepthTestEnableEXT when extension not present")
	}

	C.cgoCmdSetDepthBoundsTestEnableEXT(c.cmdSetDepthTestEnable, (C.VkCommandBuffer)(unsafe.Pointer(commandBuffer)), C.VkBool32(depthTestEnable))
}

func (c *CLoader) VkCmdSetDepthWriteEnable(commandBuffer loader.VkCommandBuffer, depthWriteEnable loader.VkBool32) {
	if c.cmdSetDepthWriteEnable == nil {
		panic("attempt to call extension method vkCmdSetDepthWriteEnableEXT when extension not present")
	}

	C.cgoCmdSetDepthWriteEnableEXT(c.cmdSetDepthWriteEnable, (C.VkCommandBuffer)(unsafe.Pointer(commandBuffer)), C.VkBool32(depthWriteEnable))
}

func (c *CLoader) VkCmdSetFrontFace(commandBuffer loader.VkCommandBuffer, frontFace loader.VkFrontFace) {
	if c.cmdSetFrontFace == nil {
		panic("attempt to call extension method vkCmdSetFrontFaceEXT when extension not present")
	}

	C.cgoCmdSetFrontFaceEXT(c.cmdSetFrontFace, (C.VkCommandBuffer)(unsafe.Pointer(commandBuffer)), C.VkFrontFace(frontFace))
}

func (c *CLoader) VkCmdSetPrimitiveTopology(commandBuffer loader.VkCommandBuffer, primitiveTopology loader.VkPrimitiveTopology) {
	if c.cmdSetPrimitiveTopology == nil {
		panic("attempt to call extension method vkCmdSetPrimitiveTopologyEXT when extension not present")
	}

	C.cgoCmdSetPrimitiveTopologyEXT(c.cmdSetPrimitiveTopology, (C.VkCommandBuffer)(unsafe.Pointer(commandBuffer)), C.VkPrimitiveTopology(primitiveTopology))
}

func (c *CLoader) VkCmdSetScissorWithCount(commandBuffer loader.VkCommandBuffer, scissorCount loader.Uint32, pScissors *loader.VkRect2D) {
	if c.cmdSetScissorWithCount == nil {
		panic("attempt to call extension method vkCmdSetScissorWithCountEXT when extension not present")
	}

	C.cgoCmdSetScissorWithCountEXT(c.cmdSetScissorWithCount, (C.VkCommandBuffer)(unsafe.Pointer(commandBuffer)), C.uint32_t(scissorCount), (*C.VkRect2D)(unsafe.Pointer(pScissors)))
}

func (c *CLoader) VkCmdSetStencilOp(commandBuffer loader.VkCommandBuffer, faceMask loader.VkStencilFaceFlags, failOp loader.VkStencilOp, passOp loader.VkStencilOp, depthFailOp loader.VkStencilOp, compareOp loader.VkCompareOp) {
	if c.cmdSetStencilOp == nil {
		panic("attempt to call extension method vkCmdSetStencilOpEXT when extension not present")
	}

	C.cgoCmdSetStencilOpEXT(c.cmdSetStencilOp, (C.VkCommandBuffer)(unsafe.Pointer(commandBuffer)), C.VkStencilFaceFlags(faceMask), C.VkStencilOp(failOp), C.VkStencilOp(passOp), C.VkStencilOp(depthFailOp), C.VkCompareOp(compareOp))
}

func (c *CLoader) VkCmdSetStencilTestEnable(commandBuffer loader.VkCommandBuffer, stencilTestEnable loader.VkBool32) {
	if c.cmdSetStencilTestEnable == nil {
		panic("attempt to call extension method vkCmdSetStencilTestEnableEXT when extension not present")
	}

	C.cgoCmdSetStencilTestEnableEXT(c.cmdSetStencilTestEnable, (C.VkCommandBuffer)(unsafe.Pointer(commandBuffer)), C.VkBool32(stencilTestEnable))
}

func (c *CLoader) VkCmdSetViewportWithCount(commandBuffer loader.VkCommandBuffer, viewportCount loader.Uint32, pViewports *loader.VkViewport) {
	if c.cmdSetViewportWithCount == nil {
		panic("attempt to call extension method vkCmdSetViewportWithCountEXT when extension not present")
	}

	C.cgoCmdSetViewportWithCountEXT(c.cmdSetViewportWithCount, (C.VkCommandBuffer)(unsafe.Pointer(commandBuffer)), C.uint32_t(viewportCount), (*C.VkViewport)(unsafe.Pointer(pViewports)))
}
