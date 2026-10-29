package ext_extended_dynamic_state2_loader

/*
#include <stdlib.h>
#include "../../vulkan/vulkan.h"

void cgoCmdSetDepthBiasEnableEXT(PFN_vkCmdSetDepthBiasEnableEXT fn, VkCommandBuffer commandBuffer, VkBool32 depthBiasEnable) {
	fn(commandBuffer, depthBiasEnable);
}

void cgoCmdSetLogicOpEXT(PFN_vkCmdSetLogicOpEXT fn, VkCommandBuffer commandBuffer, VkLogicOp logicOp) {
	fn(commandBuffer, logicOp);
}

void cgoCmdSetPatchControlPointsEXT(PFN_vkCmdSetPatchControlPointsEXT fn, VkCommandBuffer commandBuffer, uint32_t patchControlPoints) {
	fn(commandBuffer, patchControlPoints);
}

void cgoCmdSetPrimitiveRestartEnableEXT(PFN_vkCmdSetPrimitiveRestartEnableEXT fn, VkCommandBuffer commandBuffer, VkBool32 primitiveRestartEnable) {
	fn(commandBuffer, primitiveRestartEnable);
}

void cgoCmdSetRasterizerDiscardEnableEXT(PFN_vkCmdSetRasterizerDiscardEnableEXT fn, VkCommandBuffer commandBuffer, VkBool32 rasterizerDiscardEnable) {
	fn(commandBuffer, rasterizerDiscardEnable);
}
*/
import "C"
import (
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3/loader"
)

//go:generate go run go.uber.org/mock/mockgen -source loader.go -destination ../mocks/loader.go -package mock_extended_dynamic_state2

type Loader interface {
	VkCmdSetDepthBiasEnableEXT(commandBuffer loader.VkCommandBuffer, depthBiasEnable loader.VkBool32)
	VkCmdSetLogicOpEXT(commandBuffer loader.VkCommandBuffer, logicOp loader.VkLogicOp)
	VkCmdSetPatchControlPointsEXT(commandBuffer loader.VkCommandBuffer, patchControlPoints loader.Uint32)
	VkCmdSetPrimitiveRestartEnableEXT(commandBuffer loader.VkCommandBuffer, primitiveRestartEnable loader.VkBool32)
	VkCmdSetRasterizerDiscardEnableEXT(commandBuffer loader.VkCommandBuffer, rasterizerDiscardEnable loader.VkBool32)
}

type VkPhysicalDeviceExtendedDynamicState2FeaturesEXT C.VkPhysicalDeviceExtendedDynamicState2FeaturesEXT

type CLoader struct {
	coreLoader loader.Loader

	cmdSetDepthBiasEnable         C.PFN_vkCmdSetDepthBiasEnableEXT
	cmdSetLogicOp                 C.PFN_vkCmdSetLogicOpEXT
	cmdSetPatchControlPoints      C.PFN_vkCmdSetPatchControlPointsEXT
	cmdSetPrimitiveRestartEnable  C.PFN_vkCmdSetPrimitiveRestartEnableEXT
	cmdSetRasterizerDiscardEnable C.PFN_vkCmdSetRasterizerDiscardEnableEXT
}

func CreateLoaderFromCore(coreLoader loader.Loader) *CLoader {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	return &CLoader{
		coreLoader: coreLoader,
	}
}

var _ Loader = &CLoader{}

func (c *CLoader) VkCmdSetDepthBiasEnableEXT(commandBuffer loader.VkCommandBuffer, depthBiasEnable loader.VkBool32) {
	if c.cmdSetDepthBiasEnable == nil {
		panic("attempt to call extension method vkCmdSetDepthBiasEnableEXT when extension not present")
	}

	C.cgoCmdSetDepthBiasEnableEXT(c.cmdSetDepthBiasEnable, (C.VkCommandBuffer)(unsafe.Pointer(commandBuffer)), C.VkBool32(depthBiasEnable))
}

// VkCmdSetLogicOpEXT implements [Loader].
func (c *CLoader) VkCmdSetLogicOpEXT(commandBuffer loader.VkCommandBuffer, logicOp loader.VkLogicOp) {
	if c.cmdSetLogicOp == nil {
		panic("attempt to call extension method vkCmdSetLogicOpEXT when extension not present")
	}

	C.cgoCmdSetLogicOpEXT(c.cmdSetLogicOp, (C.VkCommandBuffer)(unsafe.Pointer(commandBuffer)), C.VkLogicOp(logicOp))
}

// VkCmdSetPatchControlPointsEXT implements [Loader].
func (c *CLoader) VkCmdSetPatchControlPointsEXT(commandBuffer loader.VkCommandBuffer, patchControlPoints loader.Uint32) {
	if c.cmdSetPatchControlPoints == nil {
		panic("attempt to call extension method vkCmdSetPatchControlPointsEXT when extension not present")
	}

	C.cgoCmdSetPatchControlPointsEXT(c.cmdSetPatchControlPoints, (C.VkCommandBuffer)(unsafe.Pointer(commandBuffer)), C.uint32_t(patchControlPoints))
}

// VkCmdSetPrimitiveRestartEnableEXT implements [Loader].
func (c *CLoader) VkCmdSetPrimitiveRestartEnableEXT(commandBuffer loader.VkCommandBuffer, primitiveRestartEnable loader.VkBool32) {
	if c.cmdSetPrimitiveRestartEnable == nil {
		panic("attempt to call extension method vkCmdSetPrimitiveRestartEnableEXT when extension not present")
	}

	C.cgoCmdSetPrimitiveRestartEnableEXT(c.cmdSetPrimitiveRestartEnable, (C.VkCommandBuffer)(unsafe.Pointer(commandBuffer)), C.VkBool32(primitiveRestartEnable))
}

// VkCmdSetRasterizerDiscardEnableEXT implements [Loader].
func (c *CLoader) VkCmdSetRasterizerDiscardEnableEXT(commandBuffer loader.VkCommandBuffer, rasterizerDiscardEnable loader.VkBool32) {
	if c.cmdSetRasterizerDiscardEnable == nil {
		panic("attempt to call extension method vkCmdSetRasterizerDiscardEnableEXT when extension not present")
	}

	C.cgoCmdSetRasterizerDiscardEnableEXT(c.cmdSetRasterizerDiscardEnable, (C.VkCommandBuffer)(unsafe.Pointer(commandBuffer)), C.VkBool32(rasterizerDiscardEnable))
}
