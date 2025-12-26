package khr_draw_indirect_count_loader

/*
#include <stdlib.h>
#include "../../vulkan/vulkan.h"

void cgoCmdDrawIndexedIndirectCountKHR(PFN_vkCmdDrawIndexedIndirectCountKHR fn, VkCommandBuffer commandBuffer, VkBuffer buffer, VkDeviceSize offset, VkBuffer countBuffer, VkDeviceSize countBufferOffset, uint32_t maxDrawCount, uint32_t stride) {
	fn(commandBuffer, buffer, offset, countBuffer, countBufferOffset, maxDrawCount, stride);
}

void cgoCmdDrawIndirectCountKHR(PFN_vkCmdDrawIndirectCountKHR fn, VkCommandBuffer commandBuffer, VkBuffer buffer, VkDeviceSize offset, VkBuffer countBuffer, VkDeviceSize countBufferOffset, uint32_t maxDrawCount, uint32_t stride) {
	fn(commandBuffer, buffer, offset, countBuffer, countBufferOffset, maxDrawCount, stride);
}
*/
import "C"
import (
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3/loader"
	_ "github.com/vkngwrapper/extensions/v3/vulkan"
)

//go:generate mockgen -source loader.go -destination ../mocks/loader.go -package mock_draw_indirect_count

type Loader interface {
	VkCmdDrawIndexedIndirectCountKHR(commandBuffer loader.VkCommandBuffer, buffer loader.VkBuffer, offset loader.VkDeviceSize, countBuffer loader.VkBuffer, countBufferOffset loader.VkDeviceSize, maxDrawCount loader.Uint32, stride loader.Uint32)
	VkCmdDrawIndirectCountKHR(commandBuffer loader.VkCommandBuffer, buffer loader.VkBuffer, offset loader.VkDeviceSize, countBuffer loader.VkBuffer, countBufferOffset loader.VkDeviceSize, maxDrawCount loader.Uint32, stride loader.Uint32)
}

type CLoader struct {
	coreLoader loader.Loader

	drawIndexedIndirectCount C.PFN_vkCmdDrawIndexedIndirectCountKHR
	drawIndirectCount        C.PFN_vkCmdDrawIndirectCountKHR
}

func CreateDriverFromCore(coreLoader loader.Loader) *CLoader {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	return &CLoader{
		coreLoader: coreLoader,

		drawIndexedIndirectCount: (C.PFN_vkCmdDrawIndexedIndirectCountKHR)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkCmdDrawIndexedIndirectCountKHR")))),
		drawIndirectCount:        (C.PFN_vkCmdDrawIndirectCountKHR)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkCmdDrawIndirectCountKHR")))),
	}
}

func (d *CLoader) VkCmdDrawIndexedIndirectCountKHR(commandBuffer loader.VkCommandBuffer, buffer loader.VkBuffer, offset loader.VkDeviceSize, countBuffer loader.VkBuffer, countBufferOffset loader.VkDeviceSize, maxDrawCount loader.Uint32, stride loader.Uint32) {
	if d.drawIndexedIndirectCount == nil {
		panic("attempt to call extension method vkCmdDrawIndexedIndirectCountKHR when extension not present")
	}

	C.cgoCmdDrawIndexedIndirectCountKHR(
		d.drawIndexedIndirectCount,
		C.VkCommandBuffer(unsafe.Pointer(commandBuffer)),
		C.VkBuffer(unsafe.Pointer(buffer)),
		C.VkDeviceSize(offset),
		C.VkBuffer(unsafe.Pointer(countBuffer)),
		C.VkDeviceSize(countBufferOffset),
		C.uint32_t(maxDrawCount),
		C.uint32_t(stride),
	)
}

func (d *CLoader) VkCmdDrawIndirectCountKHR(commandBuffer loader.VkCommandBuffer, buffer loader.VkBuffer, offset loader.VkDeviceSize, countBuffer loader.VkBuffer, countBufferOffset loader.VkDeviceSize, maxDrawCount loader.Uint32, stride loader.Uint32) {
	if d.drawIndirectCount == nil {
		panic("attempt to call extension method vkCmdDrawIndirectCountKHR when extension not present")
	}

	C.cgoCmdDrawIndirectCountKHR(
		d.drawIndirectCount,
		C.VkCommandBuffer(unsafe.Pointer(commandBuffer)),
		C.VkBuffer(unsafe.Pointer(buffer)),
		C.VkDeviceSize(offset),
		C.VkBuffer(unsafe.Pointer(countBuffer)),
		C.VkDeviceSize(countBufferOffset),
		C.uint32_t(maxDrawCount),
		C.uint32_t(stride),
	)
}
