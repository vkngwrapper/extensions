package khr_synchronization2_loader

/*
#include <stdlib.h>
#include "../../vulkan/vulkan.h"

void cgoCmdPipelineBarrier2KHR(PFN_vkCmdPipelineBarrier2KHR fn, VkCommandBuffer commandBuffer, VkDependencyInfo *pDependencyInfo) {
	fn(commandBuffer, pDependencyInfo);
}

void cgoCmdResetEvent2KHR(PFN_vkCmdResetEvent2KHR fn, VkCommandBuffer commandBuffer, VkEvent event, VkPipelineStageFlags2 stageMask) {
	fn(commandBuffer, event, stageMask);
}

void cgoCmdSetEvent2KHR(PFN_vkCmdSetEvent2KHR fn, VkCommandBuffer commandBuffer, VkEvent event, VkDependencyInfo *pDependencyInfo) {
	fn(commandBuffer, event, pDependencyInfo);
}

void cgoCmdWaitEvents2KHR(PFN_vkCmdWaitEvents2KHR fn, VkCommandBuffer commandBuffer, uint32_t eventCount, VkEvent *pEvents, VkDependencyInfo *pDependencyInfos) {
	fn(commandBuffer, eventCount, pEvents, pDependencyInfos);
}

void cgoCmdWriteTimestamp2KHR(PFN_vkCmdWriteTimestamp2KHR fn, VkCommandBuffer commandBuffer, VkPipelineStageFlags2 stage, VkQueryPool queryPool, uint32_t query) {
	fn(commandBuffer, stage, queryPool, query);
}

VkResult cgoQueueSubmit2KHR(PFN_vkQueueSubmit2KHR fn, VkQueue queue, uint32_t submitCount, VkSubmitInfo2 *pSubmits, VkFence fence) {
	return fn(queue, submitCount, pSubmits, fence);
}
*/
import "C"
import (
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/loader"
)

//go:generate go run go.uber.org/mock/mockgen -source loader.go -destination ../mocks/loader.go -package mock_synchronization2

type Loader interface {
	VkCmdPipelineBarrier2KHR(commandBuffer loader.VkCommandBuffer, pDependencyInfo *VkDependencyInfo)
	VkCmdResetEvent2KHR(commandBuffer loader.VkCommandBuffer, event loader.VkEvent, stageMask VkPipelineStageFlags2)
	VkCmdSetEvent2KHR(commandBuffer loader.VkCommandBuffer, event loader.VkEvent, pDependencyInfo *VkDependencyInfo)
	VkCmdWaitEvents2KHR(commandBuffer loader.VkCommandBuffer, eventCount loader.Uint32, pEvents *loader.VkEvent, pDependencyInfos *VkDependencyInfo)
	VkCmdWriteTimestamp2KHR(commandBuffer loader.VkCommandBuffer, stage VkPipelineStageFlags2, queryPool loader.VkQueryPool, query loader.Uint32)
	VkQueueSubmit2(queue loader.VkQueue, submitCount loader.Uint32, pSubmits *VkSubmitInfo2, fence loader.VkFence) (common.VkResult, error)
}

type VkDependencyInfo C.VkDependencyInfo
type VkPipelineStageFlags2 C.VkPipelineStageFlags2
type VkSubmitInfo2 C.VkSubmitInfo2
type VkMemoryBarrier2 C.VkMemoryBarrier2
type VkBufferMemoryBarrier2 C.VkBufferMemoryBarrier2
type VkImageMemoryBarrier2 C.VkImageMemoryBarrier2
type VkSemaphoreSubmitInfo C.VkSemaphoreSubmitInfo
type VkCommandBufferSubmitInfo C.VkCommandBufferSubmitInfo
type VkPhysicalDeviceSynchronization2Features C.VkPhysicalDeviceSynchronization2Features

type CLoader struct {
	coreLoader loader.Loader

	cmdPipelineBarrier2 C.PFN_vkCmdPipelineBarrier2KHR
	cmdResetEvent2      C.PFN_vkCmdResetEvent2KHR
	cmdSetEvent2        C.PFN_vkCmdSetEvent2KHR
	cmdWaitEvents2      C.PFN_vkCmdWaitEvents2KHR
	cmdWriteTimestamp2  C.PFN_vkCmdWriteTimestamp2KHR
	queueSubmit2        C.PFN_vkQueueSubmit2KHR
}

func CreateLoaderFromCore(coreLoader loader.Loader) *CLoader {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	return &CLoader{
		coreLoader: coreLoader,

		cmdPipelineBarrier2: (C.PFN_vkCmdPipelineBarrier2KHR)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkCmdPipelineBarrier2KHR")))),
		cmdResetEvent2:      (C.PFN_vkCmdResetEvent2KHR)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkCmdResetEvent2KHR")))),
		cmdSetEvent2:        (C.PFN_vkCmdSetEvent2KHR)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkCmdSetEvent2KHR")))),
		cmdWaitEvents2:      (C.PFN_vkCmdWaitEvents2KHR)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkCmdWaitEvents2KHR")))),
		cmdWriteTimestamp2:  (C.PFN_vkCmdWriteTimestamp2KHR)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkCmdWriteTimestamp2KHR")))),
		queueSubmit2:        (C.PFN_vkQueueSubmit2KHR)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkQueueSubmit2KHR")))),
	}
}

var _ Loader = &CLoader{}

func (c *CLoader) VkCmdPipelineBarrier2KHR(commandBuffer loader.VkCommandBuffer, pDependencyInfo *VkDependencyInfo) {
	if c.cmdPipelineBarrier2 == nil {
		panic("attempt to call extension method vkCmdPipelineBarrier2KHR when extension not present")
	}

	C.cgoCmdPipelineBarrier2KHR(
		c.cmdPipelineBarrier2,
		(C.VkCommandBuffer)(unsafe.Pointer(commandBuffer)),
		(*C.VkDependencyInfo)(unsafe.Pointer(pDependencyInfo)),
	)
}

func (c *CLoader) VkCmdResetEvent2KHR(commandBuffer loader.VkCommandBuffer, event loader.VkEvent, stageMask VkPipelineStageFlags2) {
	if c.cmdResetEvent2 == nil {
		panic("attempt to call extension method vkCmdResetEvent2KHR when extension not present")
	}

	C.cgoCmdResetEvent2KHR(
		c.cmdResetEvent2,
		(C.VkCommandBuffer)(unsafe.Pointer(commandBuffer)),
		(C.VkEvent)(unsafe.Pointer(event)),
		(C.VkPipelineStageFlags2)(stageMask),
	)
}

func (c *CLoader) VkCmdSetEvent2KHR(commandBuffer loader.VkCommandBuffer, event loader.VkEvent, pDependencyInfo *VkDependencyInfo) {
	if c.cmdSetEvent2 == nil {
		panic("attempt to call extension method vkCmdSetEvent2KHR when extension not present")
	}

	C.cgoCmdSetEvent2KHR(
		c.cmdSetEvent2,
		(C.VkCommandBuffer)(unsafe.Pointer(commandBuffer)),
		(C.VkEvent)(unsafe.Pointer(event)),
		(*C.VkDependencyInfo)(unsafe.Pointer(pDependencyInfo)),
	)
}

func (c *CLoader) VkCmdWaitEvents2KHR(commandBuffer loader.VkCommandBuffer, eventCount loader.Uint32, pEvents *loader.VkEvent, pDependencyInfos *VkDependencyInfo) {
	if c.cmdWaitEvents2 == nil {
		panic("attempt to call extension method vkCmdWaitEvents2KHR when extension not present")
	}

	C.cgoCmdWaitEvents2KHR(
		c.cmdWaitEvents2,
		(C.VkCommandBuffer)(unsafe.Pointer(commandBuffer)),
		(C.uint32_t)(eventCount),
		(*C.VkEvent)(unsafe.Pointer(pEvents)),
		(*C.VkDependencyInfo)(unsafe.Pointer(pDependencyInfos)),
	)
}

func (c *CLoader) VkCmdWriteTimestamp2KHR(commandBuffer loader.VkCommandBuffer, stage VkPipelineStageFlags2, queryPool loader.VkQueryPool, query loader.Uint32) {
	if c.cmdWriteTimestamp2 == nil {
		panic("attempt to call extension method vkCmdWriteTimestamp2KHR when extension not present")
	}

	C.cgoCmdWriteTimestamp2KHR(
		c.cmdWriteTimestamp2,
		(C.VkCommandBuffer)(unsafe.Pointer(commandBuffer)),
		(C.VkPipelineStageFlags2)(stage),
		(C.VkQueryPool)(unsafe.Pointer(queryPool)),
		(C.uint32_t)(query),
	)
}

func (c *CLoader) VkQueueSubmit2(queue loader.VkQueue, submitCount loader.Uint32, pSubmits *VkSubmitInfo2, fence loader.VkFence) (common.VkResult, error) {
	if c.queueSubmit2 == nil {
		panic("attempt to call extension method vkQueueSubmit2KHR when extension not present")
	}

	res := common.VkResult(C.cgoQueueSubmit2KHR(
		c.queueSubmit2,
		(C.VkQueue)(unsafe.Pointer(queue)),
		(C.uint32_t)(submitCount),
		(*C.VkSubmitInfo2)(unsafe.Pointer(pSubmits)),
		(C.VkFence)(unsafe.Pointer(fence)),
	))

	return res, res.ToError()
}
