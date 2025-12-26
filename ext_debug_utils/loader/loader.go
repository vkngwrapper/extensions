package ext_debug_utils_loader

/*
#include <stdlib.h>
#include "../../vulkan/vulkan.h"

VkResult cgoCreateDebugUtilsMessengerEXT(PFN_vkCreateDebugUtilsMessengerEXT fn, VkInstance instance, VkDebugUtilsMessengerCreateInfoEXT *pCreateInfo, VkAllocationCallbacks *pAllocator, VkDebugUtilsMessengerEXT* pDebugMessenger) {
	return fn(instance, pCreateInfo, pAllocator, pDebugMessenger);
}

void cgoDestroyDebugUtilsMessengerEXT(PFN_vkDestroyDebugUtilsMessengerEXT fn, VkInstance instance, VkDebugUtilsMessengerEXT debugMessenger, VkAllocationCallbacks* pAllocator) {
	fn(instance, debugMessenger, pAllocator);
}

void cgoCmdBeginDebugUtilsLabelEXT(PFN_vkCmdBeginDebugUtilsLabelEXT fn, VkCommandBuffer commandBuffer, VkDebugUtilsLabelEXT *pLabelInfo) {
	fn(commandBuffer, pLabelInfo);
}

void cgoCmdEndDebugUtilsLabelEXT(PFN_vkCmdEndDebugUtilsLabelEXT fn, VkCommandBuffer commandBuffer) {
	fn(commandBuffer);
}

void cgoCmdInsertDebugUtilsLabelEXT(PFN_vkCmdInsertDebugUtilsLabelEXT fn, VkCommandBuffer commandBuffer, VkDebugUtilsLabelEXT *pLabelInfo) {
	fn(commandBuffer, pLabelInfo);
}

void cgoQueueBeginDebugUtilsLabelEXT(PFN_vkQueueBeginDebugUtilsLabelEXT fn, VkQueue queue, VkDebugUtilsLabelEXT *pLabelInfo) {
	fn(queue, pLabelInfo);
}

void cgoQueueEndDebugUtilsLabelEXT(PFN_vkQueueEndDebugUtilsLabelEXT fn, VkQueue queue) {
	fn(queue);
}

void cgoQueueInsertDebugUtilsLabelEXT(PFN_vkQueueInsertDebugUtilsLabelEXT fn, VkQueue queue, VkDebugUtilsLabelEXT *pLabelInfo) {
	fn(queue, pLabelInfo);
}

VkResult cgoSetDebugUtilsObjectNameEXT(PFN_vkSetDebugUtilsObjectNameEXT fn, VkDevice device, VkDebugUtilsObjectNameInfoEXT *pNameInfo) {
	return fn(device, pNameInfo);
}

VkResult cgoSetDebugUtilsObjectTagEXT(PFN_vkSetDebugUtilsObjectTagEXT fn, VkDevice device, VkDebugUtilsObjectTagInfoEXT *pTagInfo) {
	return fn(device, pTagInfo);
}

void cgoSubmitDebugUtilsMessageEXT(PFN_vkSubmitDebugUtilsMessageEXT fn, VkInstance instance, VkDebugUtilsMessageSeverityFlagBitsEXT messageSeverity, VkDebugUtilsMessageTypeFlagsEXT messageTypes, VkDebugUtilsMessengerCallbackDataEXT *pCallbackData) {
	fn(instance, messageSeverity, messageTypes, pCallbackData);
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

//go:generate mockgen -source loader.go -destination ../mocks/loader.go -package mock_debugutils

type CLoader struct {
	createDebugUtilsMessenger  C.PFN_vkCreateDebugUtilsMessengerEXT
	destroyDebugUtilsMessenger C.PFN_vkDestroyDebugUtilsMessengerEXT
	cmdBeginDebugUtilsLabel    C.PFN_vkCmdBeginDebugUtilsLabelEXT
	cmdEndDebugUtilsLabel      C.PFN_vkCmdEndDebugUtilsLabelEXT
	cmdInsertDebugUtilsLabel   C.PFN_vkCmdInsertDebugUtilsLabelEXT
	queueBeginDebugUtilsLabel  C.PFN_vkQueueBeginDebugUtilsLabelEXT
	queueEndDebugUtilsLabel    C.PFN_vkQueueEndDebugUtilsLabelEXT
	queueInsertDebugUtilsLabel C.PFN_vkQueueInsertDebugUtilsLabelEXT
	setDebugUtilsObjectName    C.PFN_vkSetDebugUtilsObjectNameEXT
	setDebugUtilsObjectTag     C.PFN_vkSetDebugUtilsObjectTagEXT
	submitDebugUtilsMessage    C.PFN_vkSubmitDebugUtilsMessageEXT
}

type VkDebugUtilsMessengerCreateInfoEXT C.VkDebugUtilsMessengerCreateInfoEXT
type VkDebugUtilsMessengerEXT loader.VulkanHandle
type VkDebugUtilsLabelEXT C.VkDebugUtilsLabelEXT
type VkDebugUtilsObjectNameInfoEXT C.VkDebugUtilsObjectNameInfoEXT
type VkDebugUtilsObjectTagInfoEXT C.VkDebugUtilsObjectTagInfoEXT
type VkDebugUtilsMessageSeverityFlagBitsEXT C.VkDebugUtilsMessageSeverityFlagBitsEXT
type VkDebugUtilsMessageTypeFlagsEXT C.VkDebugUtilsMessageTypeFlagsEXT
type VkDebugUtilsMessengerCallbackDataEXT C.VkDebugUtilsMessengerCallbackDataEXT
type Loader interface {
	VkCreateDebugUtilsMessengerEXT(instance loader.VkInstance, pCreateInfo *VkDebugUtilsMessengerCreateInfoEXT, pAllocator *loader.VkAllocationCallbacks, pDebugMessenger *VkDebugUtilsMessengerEXT) (common.VkResult, error)
	VkDestroyDebugUtilsMessengerEXT(instance loader.VkInstance, debugMessenger VkDebugUtilsMessengerEXT, pAllocator *loader.VkAllocationCallbacks)
	VKCmdBeginDebugUtilsLabelEXT(commandBuffer loader.VkCommandBuffer, pLabelInfo *VkDebugUtilsLabelEXT)
	VkCmdEndDebugUtilsLabelEXT(commandBuffer loader.VkCommandBuffer)
	VkCmdInsertDebugUtilsLabelEXT(commandBuffer loader.VkCommandBuffer, pLabelInfo *VkDebugUtilsLabelEXT)
	VkQueueBeginDebugUtilsLabelEXT(queue loader.VkQueue, pLabelInfo *VkDebugUtilsLabelEXT)
	VkQueueEndDebugUtilsLabelEXT(queue loader.VkQueue)
	VkQueueInsertDebugUtilsLabelEXT(queue loader.VkQueue, pLabelInfo *VkDebugUtilsLabelEXT)
	VkSetDebugUtilsObjectNameEXT(device loader.VkDevice, pNameInfo *VkDebugUtilsObjectNameInfoEXT) (common.VkResult, error)
	VkSetDebugUtilsObjectTagEXT(device loader.VkDevice, pTagInfo *VkDebugUtilsObjectTagInfoEXT) (common.VkResult, error)
	VkSubmitDebugUtilsMessageEXT(instance loader.VkInstance, messageSeverity VkDebugUtilsMessageSeverityFlagBitsEXT, messageTypes VkDebugUtilsMessageTypeFlagsEXT, pCallbackData *VkDebugUtilsMessengerCallbackDataEXT)
}

func CreateLoaderFromCore(coreLoader loader.Loader) *CLoader {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	return &CLoader{
		createDebugUtilsMessenger:  (C.PFN_vkCreateDebugUtilsMessengerEXT)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkCreateDebugUtilsMessengerEXT")))),
		destroyDebugUtilsMessenger: (C.PFN_vkDestroyDebugUtilsMessengerEXT)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkDestroyDebugUtilsMessengerEXT")))),
		cmdBeginDebugUtilsLabel:    (C.PFN_vkCmdBeginDebugUtilsLabelEXT)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkCmdBeginDebugUtilsLabelEXT")))),
		cmdEndDebugUtilsLabel:      (C.PFN_vkCmdEndDebugUtilsLabelEXT)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkCmdEndDebugUtilsLabelEXT")))),
		cmdInsertDebugUtilsLabel:   (C.PFN_vkCmdInsertDebugUtilsLabelEXT)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkCmdInsertDebugUtilsLabelEXT")))),
		queueBeginDebugUtilsLabel:  (C.PFN_vkQueueBeginDebugUtilsLabelEXT)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkQueueBeginDebugUtilsLabelEXT")))),
		queueEndDebugUtilsLabel:    (C.PFN_vkQueueEndDebugUtilsLabelEXT)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkQueueEndDebugUtilsLabelEXT")))),
		queueInsertDebugUtilsLabel: (C.PFN_vkQueueInsertDebugUtilsLabelEXT)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkQueueInsertDebugUtilsLabelEXT")))),
		setDebugUtilsObjectName:    (C.PFN_vkSetDebugUtilsObjectNameEXT)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkSetDebugUtilsObjectNameEXT")))),
		setDebugUtilsObjectTag:     (C.PFN_vkSetDebugUtilsObjectTagEXT)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkSetDebugUtilsObjectTagEXT")))),
		submitDebugUtilsMessage:    (C.PFN_vkSubmitDebugUtilsMessageEXT)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkSubmitDebugUtilsMessageEXT")))),
	}
}

func (d *CLoader) VkCreateDebugUtilsMessengerEXT(instance loader.VkInstance, pCreateInfo *VkDebugUtilsMessengerCreateInfoEXT, pAllocator *loader.VkAllocationCallbacks, pDebugMessenger *VkDebugUtilsMessengerEXT) (common.VkResult, error) {
	if d.createDebugUtilsMessenger == nil {
		panic("attempt to call extension method vkCreateDebugUtilsMessengerEXT when extension not present")
	}

	res := common.VkResult(C.cgoCreateDebugUtilsMessengerEXT(d.createDebugUtilsMessenger,
		C.VkInstance(unsafe.Pointer(instance)),
		(*C.VkDebugUtilsMessengerCreateInfoEXT)(pCreateInfo),
		(*C.VkAllocationCallbacks)(unsafe.Pointer(pAllocator)),
		(*C.VkDebugUtilsMessengerEXT)(unsafe.Pointer(pDebugMessenger))))

	return res, res.ToError()
}

func (d *CLoader) VkDestroyDebugUtilsMessengerEXT(instance loader.VkInstance, debugMessenger VkDebugUtilsMessengerEXT, pAllocator *loader.VkAllocationCallbacks) {
	if d.destroyDebugUtilsMessenger == nil {
		panic("attempt to call extension method vkDestroyDebugUtilsMessengerEXT when extension not present")
	}

	C.cgoDestroyDebugUtilsMessengerEXT(d.destroyDebugUtilsMessenger,
		C.VkInstance(unsafe.Pointer(instance)),
		C.VkDebugUtilsMessengerEXT(unsafe.Pointer(debugMessenger)),
		(*C.VkAllocationCallbacks)(unsafe.Pointer(pAllocator)))
}

func (d *CLoader) VKCmdBeginDebugUtilsLabelEXT(commandBuffer loader.VkCommandBuffer, pLabelInfo *VkDebugUtilsLabelEXT) {
	if d.cmdBeginDebugUtilsLabel == nil {
		panic("attempt to call extension method VKCmdBeginDebugUtilsLabelEXT when extension not present")
	}

	C.cgoCmdBeginDebugUtilsLabelEXT(d.cmdBeginDebugUtilsLabel,
		C.VkCommandBuffer(unsafe.Pointer(commandBuffer)),
		(*C.VkDebugUtilsLabelEXT)(unsafe.Pointer(pLabelInfo)))
}

func (d *CLoader) VkCmdEndDebugUtilsLabelEXT(commandBuffer loader.VkCommandBuffer) {
	if d.cmdEndDebugUtilsLabel == nil {
		panic("attempt to call extension method VkCmdEndDebugUtilsLabelEXT when extension not present")
	}

	C.cgoCmdEndDebugUtilsLabelEXT(d.cmdEndDebugUtilsLabel,
		C.VkCommandBuffer(unsafe.Pointer(commandBuffer)))
}

func (d *CLoader) VkCmdInsertDebugUtilsLabelEXT(commandBuffer loader.VkCommandBuffer, pLabelInfo *VkDebugUtilsLabelEXT) {
	if d.cmdInsertDebugUtilsLabel == nil {
		panic("attempt to call extension method VkCmdInsertDebugUtilsLabelEXT when extension not present")
	}

	C.cgoCmdInsertDebugUtilsLabelEXT(d.cmdInsertDebugUtilsLabel,
		C.VkCommandBuffer(unsafe.Pointer(commandBuffer)),
		(*C.VkDebugUtilsLabelEXT)(unsafe.Pointer(pLabelInfo)))
}

func (d *CLoader) VkQueueBeginDebugUtilsLabelEXT(queue loader.VkQueue, pLabelInfo *VkDebugUtilsLabelEXT) {
	if d.queueBeginDebugUtilsLabel == nil {
		panic("attempt to call extension method VkQueueBeginDebugUtilsLabelEXT when extension not present")
	}

	C.cgoQueueBeginDebugUtilsLabelEXT(d.queueBeginDebugUtilsLabel,
		C.VkQueue(unsafe.Pointer(queue)),
		(*C.VkDebugUtilsLabelEXT)(unsafe.Pointer(pLabelInfo)))
}

func (d *CLoader) VkQueueEndDebugUtilsLabelEXT(queue loader.VkQueue) {
	if d.queueEndDebugUtilsLabel == nil {
		panic("attempt to call extension method VkQueueEndDebugUtilsLabelEXT when extension not present")
	}

	C.cgoQueueEndDebugUtilsLabelEXT(d.queueEndDebugUtilsLabel,
		C.VkQueue(unsafe.Pointer(queue)))
}

func (d *CLoader) VkQueueInsertDebugUtilsLabelEXT(queue loader.VkQueue, pLabelInfo *VkDebugUtilsLabelEXT) {
	if d.queueInsertDebugUtilsLabel == nil {
		panic("attempt to call extension method VkQueueInsertDebugUtilsLabelEXT when extension not present")
	}

	C.cgoQueueInsertDebugUtilsLabelEXT(d.queueInsertDebugUtilsLabel,
		C.VkQueue(unsafe.Pointer(queue)),
		(*C.VkDebugUtilsLabelEXT)(unsafe.Pointer(pLabelInfo)))
}

func (d *CLoader) VkSetDebugUtilsObjectNameEXT(device loader.VkDevice, pNameInfo *VkDebugUtilsObjectNameInfoEXT) (common.VkResult, error) {
	if d.setDebugUtilsObjectName == nil {
		panic("attempt to call extension method VkSetDebugUtilsObjectNameEXT when extension not present")
	}

	res := common.VkResult(C.cgoSetDebugUtilsObjectNameEXT(d.setDebugUtilsObjectName,
		C.VkDevice(unsafe.Pointer(device)),
		(*C.VkDebugUtilsObjectNameInfoEXT)(unsafe.Pointer(pNameInfo))))

	return res, res.ToError()
}

func (d *CLoader) VkSetDebugUtilsObjectTagEXT(device loader.VkDevice, pTagInfo *VkDebugUtilsObjectTagInfoEXT) (common.VkResult, error) {
	if d.setDebugUtilsObjectTag == nil {
		panic("attempt to call extension method VkSetDebugUtilsObjectTagEXT when extension not present")
	}

	res := common.VkResult(C.cgoSetDebugUtilsObjectTagEXT(d.setDebugUtilsObjectTag,
		C.VkDevice(unsafe.Pointer(device)),
		(*C.VkDebugUtilsObjectTagInfoEXT)(unsafe.Pointer(pTagInfo))))

	return res, res.ToError()
}

func (d *CLoader) VkSubmitDebugUtilsMessageEXT(instance loader.VkInstance, messageSeverity VkDebugUtilsMessageSeverityFlagBitsEXT, messageTypes VkDebugUtilsMessageTypeFlagsEXT, pCallbackData *VkDebugUtilsMessengerCallbackDataEXT) {
	if d.submitDebugUtilsMessage == nil {
		panic("attempt to call extension method VkSubmitDebugUtilsMessageEXT when extension not present")
	}

	C.cgoSubmitDebugUtilsMessageEXT(d.submitDebugUtilsMessage,
		C.VkInstance(unsafe.Pointer(instance)),
		C.VkDebugUtilsMessageSeverityFlagBitsEXT(messageSeverity),
		C.VkDebugUtilsMessageTypeFlagsEXT(messageTypes),
		(*C.VkDebugUtilsMessengerCallbackDataEXT)(unsafe.Pointer(pCallbackData)))
}
