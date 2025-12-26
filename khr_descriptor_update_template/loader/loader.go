package khr_descriptor_update_template_loader

/*
#include <stdlib.h>
#include "../../vulkan/vulkan.h"

VkResult cgoCreateDescriptorUpdateTemplateKHR(PFN_vkCreateDescriptorUpdateTemplateKHR fn, VkDevice device, VkDescriptorUpdateTemplateCreateInfoKHR *pCreateInfo, VkAllocationCallbacks *pAllocator, VkDescriptorUpdateTemplateKHR *pDescriptorUpdateTemplate) {
	return fn(device, pCreateInfo, pAllocator, pDescriptorUpdateTemplate);
}

void cgoDestroyDescriptorUpdateTemplateKHR(PFN_vkDestroyDescriptorUpdateTemplateKHR fn, VkDevice device, VkDescriptorUpdateTemplateKHR descriptorUpdateTemplate, VkAllocationCallbacks *pAllocator) {
	fn(device, descriptorUpdateTemplate, pAllocator);
}

void cgoUpdateDescriptorSetWithTemplateKHR(PFN_vkUpdateDescriptorSetWithTemplateKHR fn, VkDevice device, VkDescriptorSet descriptorSet, VkDescriptorUpdateTemplateKHR descriptorUpdateTemplate, void *pData) {
	fn(device, descriptorSet, descriptorUpdateTemplate, pData);
}

void cgoCmdPushDescriptorSetWithTemplateKHR(PFN_vkCmdPushDescriptorSetWithTemplateKHR fn, VkCommandBuffer commandBuffer, VkDescriptorUpdateTemplateKHR descriptorUpdateTemplate, VkPipelineLayout layout, uint32_t set, void *pData) {
	fn(commandBuffer, descriptorUpdateTemplate, layout, set, pData);
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

//go:generate mockgen -source loader.go -destination ../mocks/loader.go -package mock_descriptor_update_template

type Loader interface {
	VkCreateDescriptorUpdateTemplateKHR(device loader.VkDevice, pCreateInfo *VkDescriptorUpdateTemplateCreateInfoKHR, pAllocator *loader.VkAllocationCallbacks, pDescriptorUpdateTemplate *VkDescriptorUpdateTemplateKHR) (common.VkResult, error)
	VkDestroyDescriptorUpdateTemplateKHR(device loader.VkDevice, descriptorUpdateTemplate VkDescriptorUpdateTemplateKHR, pAllocator *loader.VkAllocationCallbacks)
	VkUpdateDescriptorSetWithTemplateKHR(device loader.VkDevice, descriptorSet loader.VkDescriptorSet, descriptorUpdateTemplate VkDescriptorUpdateTemplateKHR, pData unsafe.Pointer)
	VkCmdPushDescriptorSetWithTemplateKHR(commandBuffer loader.VkCommandBuffer, descriptorUpdateTemplate VkDescriptorUpdateTemplateKHR, layout loader.VkPipelineLayout, set loader.Uint32, pData unsafe.Pointer)
}

type VkDescriptorUpdateTemplateKHR loader.VulkanHandle
type VkDescriptorUpdateTemplateCreateInfoKHR C.VkDescriptorUpdateTemplateCreateInfoKHR
type VkDescriptorUpdateTemplateEntryKHR C.VkDescriptorUpdateTemplateEntryKHR

type CLoader struct {
	coreLoader loader.Loader

	createDescriptorUpdateTemplate   C.PFN_vkCreateDescriptorUpdateTemplateKHR
	destroyDescriptorUpdateTemplate  C.PFN_vkDestroyDescriptorUpdateTemplateKHR
	updateDescriptorSetWithTemplate  C.PFN_vkUpdateDescriptorSetWithTemplateKHR
	cmdPushDescriptorSetWithTemplate C.PFN_vkCmdPushDescriptorSetWithTemplateKHR
}

func CreateLoaderFromCore(coreLoader loader.Loader) *CLoader {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	return &CLoader{
		coreLoader: coreLoader,

		createDescriptorUpdateTemplate:   (C.PFN_vkCreateDescriptorUpdateTemplateKHR)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkCreateDescriptorUpdateTemplateKHR")))),
		destroyDescriptorUpdateTemplate:  (C.PFN_vkDestroyDescriptorUpdateTemplateKHR)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkDestroyDescriptorUpdateTemplateKHR")))),
		updateDescriptorSetWithTemplate:  (C.PFN_vkUpdateDescriptorSetWithTemplateKHR)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkUpdateDescriptorSetWithTemplateKHR")))),
		cmdPushDescriptorSetWithTemplate: (C.PFN_vkCmdPushDescriptorSetWithTemplateKHR)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkCmdPushDescriptorSetWithTemplateKHR")))),
	}
}

func (d *CLoader) VkCreateDescriptorUpdateTemplateKHR(device loader.VkDevice, pCreateInfo *VkDescriptorUpdateTemplateCreateInfoKHR, pAllocator *loader.VkAllocationCallbacks, pDescriptorUpdateTemplate *VkDescriptorUpdateTemplateKHR) (common.VkResult, error) {
	if d.createDescriptorUpdateTemplate == nil {
		panic("attempt to call extension method vkCreateDescriptorUpdateTemplateKHR when extension not present")
	}

	res := common.VkResult(C.cgoCreateDescriptorUpdateTemplateKHR(d.createDescriptorUpdateTemplate,
		C.VkDevice(unsafe.Pointer(device)),
		(*C.VkDescriptorUpdateTemplateCreateInfoKHR)(pCreateInfo),
		(*C.VkAllocationCallbacks)(unsafe.Pointer(pAllocator)),
		(*C.VkDescriptorUpdateTemplateKHR)(unsafe.Pointer(pDescriptorUpdateTemplate))))

	return res, res.ToError()
}

func (d *CLoader) VkDestroyDescriptorUpdateTemplateKHR(device loader.VkDevice, descriptorUpdateTemplate VkDescriptorUpdateTemplateKHR, pAllocator *loader.VkAllocationCallbacks) {
	if d.destroyDescriptorUpdateTemplate == nil {
		panic("attempt to call extension method vkDestroyDescriptorUpdateTemplateKHR when extension not present")
	}

	C.cgoDestroyDescriptorUpdateTemplateKHR(d.destroyDescriptorUpdateTemplate,
		C.VkDevice(unsafe.Pointer(device)),
		C.VkDescriptorUpdateTemplateKHR(unsafe.Pointer(descriptorUpdateTemplate)),
		(*C.VkAllocationCallbacks)(unsafe.Pointer(pAllocator)))
}

func (d *CLoader) VkUpdateDescriptorSetWithTemplateKHR(device loader.VkDevice, descriptorSet loader.VkDescriptorSet, descriptorUpdateTemplate VkDescriptorUpdateTemplateKHR, pData unsafe.Pointer) {
	if d.updateDescriptorSetWithTemplate == nil {
		panic("attempt to call extension method vkUpdateDescriptorSetWithTemplateKHR when extension not present")
	}

	C.cgoUpdateDescriptorSetWithTemplateKHR(d.updateDescriptorSetWithTemplate,
		C.VkDevice(unsafe.Pointer(device)),
		C.VkDescriptorSet(unsafe.Pointer(descriptorSet)),
		C.VkDescriptorUpdateTemplateKHR(unsafe.Pointer(descriptorUpdateTemplate)),
		pData)
}

func (d *CLoader) VkCmdPushDescriptorSetWithTemplateKHR(commandBuffer loader.VkCommandBuffer, descriptorUpdateTemplate VkDescriptorUpdateTemplateKHR, layout loader.VkPipelineLayout, set loader.Uint32, pData unsafe.Pointer) {
	if d.cmdPushDescriptorSetWithTemplate == nil {
		panic("attempt to call extension method vkUpdateDescriptorSetWithTemplateKHR when prerequisite not present")
	}

	C.cgoCmdPushDescriptorSetWithTemplateKHR(d.cmdPushDescriptorSetWithTemplate,
		C.VkCommandBuffer(unsafe.Pointer(commandBuffer)),
		C.VkDescriptorUpdateTemplateKHR(unsafe.Pointer(descriptorUpdateTemplate)),
		C.VkPipelineLayout(unsafe.Pointer(layout)),
		C.uint32_t(set),
		pData)
}
