package khr_acceleration_structure_loader

//go:generate mockgen -source loader.go -destination ../mocks/loader.go -package mock_acceleration_structure

/*
#include <stdlib.h>
#include "../../vulkan/vulkan.h"

VkResult cgoBuildAccelerationStructuresKHR(PFN_vkBuildAccelerationStructuresKHR fn, VkDevice device, VkDeferredOperationKHR deferredOperation, uint32_t infoCount, VkAccelerationStructureBuildGeometryInfoKHR *pInfos, const VkAccelerationStructureBuildRangeInfoKHR **ppBuildRangeInfos) {
	return fn(device, deferredOperation, infoCount, pInfos, ppBuildRangeInfos);
}

void cgoCmdBuildAccelerationStructuresIndirectKHR(PFN_vkCmdBuildAccelerationStructuresIndirectKHR fn, VkCommandBuffer commandBuffer, uint32_t infoCount, VkAccelerationStructureBuildGeometryInfoKHR *pInfos, VkDeviceAddress *pIndirectDeviceAddresses, uint32_t *pIndirectStrides, const uint32_t **ppMaxPrimitiveCounts) {
	fn(commandBuffer, infoCount, pInfos, pIndirectDeviceAddresses, pIndirectStrides, ppMaxPrimitiveCounts);
}

void cgoCmdBuildAccelerationStructuresKHR(PFN_vkCmdBuildAccelerationStructuresKHR fn, VkCommandBuffer commandBuffer, uint32_t infoCount, VkAccelerationStructureBuildGeometryInfoKHR *pInfos, const VkAccelerationStructureBuildRangeInfoKHR **ppBuildRangeInfos) {
	fn(commandBuffer, infoCount, pInfos, ppBuildRangeInfos);
}

void cgoCmdCopyAccelerationStructureKHR(PFN_vkCmdCopyAccelerationStructureKHR fn, VkCommandBuffer commandBuffer, VkCopyAccelerationStructureInfoKHR *pInfo) {
	fn(commandBuffer, pInfo);
}

void cgoCmdCopyAccelerationStructureToMemoryKHR(PFN_vkCmdCopyAccelerationStructureToMemoryKHR fn, VkCommandBuffer commandBuffer, VkCopyAccelerationStructureToMemoryInfoKHR *pInfo) {
	fn(commandBuffer, pInfo);
}

void cgoCmdCopyMemoryToAccelerationStructureKHR(PFN_vkCmdCopyMemoryToAccelerationStructureKHR fn, VkCommandBuffer commandBuffer, VkCopyMemoryToAccelerationStructureInfoKHR *pInfo) {
	fn(commandBuffer, pInfo);
}

void cgoCmdWriteAccelerationStructuresPropertiesKHR(PFN_vkCmdWriteAccelerationStructuresPropertiesKHR fn, VkCommandBuffer commandBuffer, uint32_t accelerationStructureCount, VkAccelerationStructureKHR *pAccelerationStructures, VkQueryType queryType, VkQueryPool queryPool, uint32_t firstQuery) {
	fn(commandBuffer, accelerationStructureCount, pAccelerationStructures, queryType, queryPool, firstQuery);
}

VkResult cgoCopyAccelerationStructureKHR(PFN_vkCopyAccelerationStructureKHR fn, VkDevice device, VkDeferredOperationKHR deferredOperation, VkCopyAccelerationStructureInfoKHR *pInfo) {
	return fn(device, deferredOperation, pInfo);
}

VkResult cgoCopyAccelerationStructureToMemoryKHR(PFN_vkCopyAccelerationStructureToMemoryKHR fn, VkDevice device, VkDeferredOperationKHR deferredOperation, VkCopyAccelerationStructureToMemoryInfoKHR *pInfo) {
	return fn(device, deferredOperation, pInfo);
}

VkResult cgoCopyMemoryToAccelerationStructureKHR(PFN_vkCopyMemoryToAccelerationStructureKHR fn, VkDevice device, VkDeferredOperationKHR deferredOperation, VkCopyMemoryToAccelerationStructureInfoKHR *pInfo) {
	return fn(device, deferredOperation, pInfo);
}

VkResult cgoCreateAccelerationStructureKHR(PFN_vkCreateAccelerationStructureKHR fn, VkDevice device, VkAccelerationStructureCreateInfoKHR *pCreateInfo, VkAllocationCallbacks *pAllocator, VkAccelerationStructureKHR *pAccelerationStructure) {
	return fn(device, pCreateInfo, pAllocator, pAccelerationStructure);
}

void cgoDestroyAccelerationStructureKHR(PFN_vkDestroyAccelerationStructureKHR fn, VkDevice device, VkAccelerationStructureKHR accelerationStructure, VkAllocationCallbacks *pAllocator) {
	fn(device, accelerationStructure, pAllocator);
}

void cgoGetAccelerationStructureBuildSizesKHR(PFN_vkGetAccelerationStructureBuildSizesKHR fn, VkDevice device, VkAccelerationStructureBuildTypeKHR buildType, VkAccelerationStructureBuildGeometryInfoKHR *pBuildInfo, uint32_t *pMaxPrimitiveCounts, VkAccelerationStructureBuildSizesInfoKHR *pSizeInfo) {
	fn(device, buildType, pBuildInfo, pMaxPrimitiveCounts, pSizeInfo);
}

VkDeviceAddress cgoGetAccelerationStructureDeviceAddressKHR(PFN_vkGetAccelerationStructureDeviceAddressKHR fn, VkDevice device, VkAccelerationStructureDeviceAddressInfoKHR *pInfo) {
	fn(device, pInfo);
}

void cgoGetDeviceAccelerationStructureCompatibilityKHR(PFN_vkGetDeviceAccelerationStructureCompatibilityKHR fn, VkDevice device, VkAccelerationStructureVersionInfoKHR *pVersionInfo, VkAccelerationStructureCompatibilityKHR *pCompatibility) {
	fn(device, pVersionInfo, pCompatibility);
}

VkResult cgoWriteAccelerationStructuresPropertiesKHR(PFN_vkWriteAccelerationStructuresPropertiesKHR fn, VkDevice device, uint32_t accelerationStructureCount, VkAccelerationStructureKHR *pAccelerationStructures, VkQueryType queryType, size_t dataSize, void *pData, size_t stride) {
	return fn(device, accelerationStructureCount, pAccelerationStructures, queryType, dataSize, pData, stride);
}

*/
import "C"
import (
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/loader"
	khr_deferred_host_operations_loader "github.com/vkngwrapper/extensions/v3/khr_deferred_host_operations/loader"
)

type Loader interface {
	VkBuildAccelerationStructuresKHR(
		device loader.VkDevice,
		deferredOperation khr_deferred_host_operations_loader.VkDeferredOperationKHR,
		infoCount loader.Uint32,
		pInfos *VkAccelerationStructureBuildGeometryInfoKHR,
		ppBuildRangeInfos **VkAccelerationStructureBuildRangeInfoKHR,
	) (common.VkResult, error)
	VkCmdBuildAccelerationStructuresIndirectKHR(
		commandBuffer loader.VkCommandBuffer,
		infoCount loader.Uint32,
		pInfos *VkAccelerationStructureBuildGeometryInfoKHR,
		pIndirectDeviceAddresses *loader.VkDeviceAddress,
		pIndirectStrides *loader.Uint32,
		ppMaxPrimitiveCountws **loader.Uint32,
	)
	VkCmdBuildAccelerationStructuresKHR(
		commandBuffer loader.VkCommandBuffer,
		infoCount loader.Uint32,
		pInfos *VkAccelerationStructureBuildGeometryInfoKHR,
		ppBuildRangeInfos **VkAccelerationStructureBuildRangeInfoKHR,
	)
	VkCmdCopyAccelerationStructureKHR(
		commandBuffer loader.VkCommandBuffer,
		pInfo *VkCopyAccelerationStructureInfoKHR,
	)
	VkCmdCopyAccelerationStructureToMemoryKHR(
		commandBuffer loader.VkCommandBuffer,
		pInfo *VkCopyAccelerationStructureToMemoryInfoKHR,
	)
	VkCmdCopyMemoryToAccelerationStructureKHR(
		commandBuffer loader.VkCommandBuffer,
		pInfo *VkCopyMemoryToAccelerationStructureInfoKHR,
	)
	VkCmdWriteAccelerationStructuresPropertiesKHR(
		commandBuffer loader.VkCommandBuffer,
		accelerationStructureCount loader.Uint32,
		pAccelerationStructures *VkAccelerationStructureKHR,
		queryType loader.VkQueryType,
		queryPool loader.VkQueryPool,
		firstQuery loader.Uint32,
	)
	VkCopyAccelerationStructureKHR(
		device loader.VkDevice,
		deferredOperation khr_deferred_host_operations_loader.VkDeferredOperationKHR,
		pInfo *VkCopyAccelerationStructureInfoKHR,
	) (common.VkResult, error)
	VkCopyAccelerationStructureToMemoryKHR(
		device loader.VkDevice,
		deferredOperation khr_deferred_host_operations_loader.VkDeferredOperationKHR,
		pInfo *VkCopyAccelerationStructureToMemoryInfoKHR,
	) (common.VkResult, error)
	VkCopyMemoryToAccelerationStructureKHR(
		device loader.VkDevice,
		deferredOperation khr_deferred_host_operations_loader.VkDeferredOperationKHR,
		pInfo *VkCopyMemoryToAccelerationStructureInfoKHR,
	) (common.VkResult, error)
	VkCreateAccelerationStructureKHR(
		device loader.VkDevice,
		pCreateInfo *VkAccelerationStructureCreateInfoKHR,
		pAllocator *loader.VkAllocationCallbacks,
		pAccelerationStructure *VkAccelerationStructureKHR,
	) (common.VkResult, error)
	VkDestroyAccelerationStructureKHR(
		device loader.VkDevice,
		accelerationStructure VkAccelerationStructureKHR,
		pAllocator *loader.VkAllocationCallbacks,
	)
	VkGetAccelerationStructureBuildSizesKHR(
		device loader.VkDevice,
		buildType VkAccelerationStructureBuildTypeKHR,
		pBuildInfo *VkAccelerationStructureBuildGeometryInfoKHR,
		pMaxPrimitiveCounts *loader.Uint32,
		pSizeInfo *VkAccelerationStructureBuildSizesInfoKHR,
	)
	VkGetAccelerationStructureDeviceAddressKHR(
		device loader.VkDevice,
		pInfo *VkAccelerationStructureDeviceAddressInfoKHR,
	) loader.VkDeviceAddress
	VkGetDeviceAccelerationStructureCompatibilityKHR(
		device loader.VkDevice,
		pVersionInfo *VkAccelerationStructureVersionInfoKHR,
		pCompatibility *VkAccelerationStructureCompatibilityKHR,
	)
	VkWriteAccelerationStructuresPropertiesKHR(
		device loader.VkDevice,
		accelerationStructureCount loader.Uint32,
		pAccelerationStructures *VkAccelerationStructureKHR,
		queryType loader.VkQueryType,
		dataSize loader.Size,
		pData unsafe.Pointer,
		stride loader.Size,
	) (common.VkResult, error)
}

type VkAccelerationStructureKHR loader.VulkanHandle

type VkAccelerationStructureGeometryKHR C.VkAccelerationStructureGeometryKHR
type VkAccelerationStructureVersionInfoKHR C.VkAccelerationStructureVersionInfoKHR
type VkAccelerationStructureCompatibilityKHR C.VkAccelerationStructureCompatibilityKHR
type VkAccelerationStructureDeviceAddressInfoKHR C.VkAccelerationStructureDeviceAddressInfoKHR
type VkAccelerationStructureBuildTypeKHR C.VkAccelerationStructureBuildTypeKHR
type VkAccelerationStructureBuildSizesInfoKHR C.VkAccelerationStructureBuildSizesInfoKHR
type VkAccelerationStructureCreateInfoKHR C.VkAccelerationStructureCreateInfoKHR
type VkAccelerationStructureBuildGeometryInfoKHR C.VkAccelerationStructureBuildGeometryInfoKHR
type VkAccelerationStructureBuildRangeInfoKHR C.VkAccelerationStructureBuildRangeInfoKHR
type VkCopyAccelerationStructureInfoKHR C.VkCopyAccelerationStructureInfoKHR
type VkCopyAccelerationStructureToMemoryInfoKHR C.VkCopyAccelerationStructureToMemoryInfoKHR
type VkCopyMemoryToAccelerationStructureInfoKHR C.VkCopyMemoryToAccelerationStructureInfoKHR
type VkAccelerationStructureGeometryInstancesDataKHR C.VkAccelerationStructureGeometryInstancesDataKHR
type VkAccelerationStructureGeometryAabbsDataKHR C.VkAccelerationStructureGeometryAabbsDataKHR
type VkAccelerationStructureGeometryTrianglesDataKHR C.VkAccelerationStructureGeometryTrianglesDataKHR
type VkPhysicalDeviceAccelerationStructureFeaturesKHR C.VkPhysicalDeviceAccelerationStructureFeaturesKHR
type VkPhysicalDeviceAccelerationStructurePropertiesKHR C.VkPhysicalDeviceAccelerationStructurePropertiesKHR
type VkWriteDescriptorSetAccelerationStructureKHR C.VkWriteDescriptorSetAccelerationStructureKHR

type CLoader struct {
	coreLoader loader.Loader

	buildAccelerationStructures            C.PFN_vkBuildAccelerationStructuresKHR
	cmdBuildAccelerationStructuresIndirect C.PFN_vkCmdBuildAccelerationStructuresIndirectKHR
	cmdBuildAccelerationStructures         C.PFN_vkCmdBuildAccelerationStructuresKHR
	createAccelerationStructure            C.PFN_vkCreateAccelerationStructureKHR
	destroyAccelerationStructure           C.PFN_vkDestroyAccelerationStructureKHR

	cmdCopyAccelerationStructure         C.PFN_vkCmdCopyAccelerationStructureKHR
	cmdCopyAccelerationStructureToMemory C.PFN_vkCmdCopyAccelerationStructureToMemoryKHR
	cmdCopyMemoryToAccelerationStructure C.PFN_vkCmdCopyMemoryToAccelerationStructureKHR
	copyAccelerationStructure            C.PFN_vkCopyAccelerationStructureKHR
	copyAccelerationStructureToMemory    C.PFN_vkCopyAccelerationStructureToMemoryKHR
	copyMemoryToAccelerationStructure    C.PFN_vkCopyMemoryToAccelerationStructureKHR

	cmdWriteAccelerationStructureProperties     C.PFN_vkCmdWriteAccelerationStructuresPropertiesKHR
	getAccelerationStructureBuildSizes          C.PFN_vkGetAccelerationStructureBuildSizesKHR
	getAccelerationStructureDeviceAddresses     C.PFN_vkGetAccelerationStructureDeviceAddressKHR
	getDeviceAccelerationStructureCompatibility C.PFN_vkGetDeviceAccelerationStructureCompatibilityKHR
	writeAccelerationStructuresProperties       C.PFN_vkWriteAccelerationStructuresPropertiesKHR
}

func CreateLoaderFromCore(coreLoader loader.Loader) *CLoader {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	return &CLoader{
		coreLoader: coreLoader,

		buildAccelerationStructures:            (C.PFN_vkBuildAccelerationStructuresKHR)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkBuildAccelerationStructuresKHR")))),
		cmdBuildAccelerationStructuresIndirect: (C.PFN_vkCmdBuildAccelerationStructuresIndirectKHR)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkCmdBuildAccelerationStructuresIndirectKHR")))),
		cmdBuildAccelerationStructures:         (C.PFN_vkCmdBuildAccelerationStructuresKHR)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkCmdBuildAccelerationStructuresKHR")))),
		createAccelerationStructure:            (C.PFN_vkCreateAccelerationStructureKHR)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkCreateAccelerationStructureKHR")))),
		destroyAccelerationStructure:           (C.PFN_vkDestroyAccelerationStructureKHR)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkDestroyAccelerationStructureKHR")))),

		cmdCopyAccelerationStructure:         (C.PFN_vkCmdCopyAccelerationStructureKHR)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkCmdCopyAccelerationStructureKHR")))),
		cmdCopyAccelerationStructureToMemory: (C.PFN_vkCmdCopyAccelerationStructureToMemoryKHR)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkCmdCopyAccelerationStructureToMemoryKHR")))),
		cmdCopyMemoryToAccelerationStructure: (C.PFN_vkCmdCopyMemoryToAccelerationStructureKHR)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkCmdCopyMemoryToAccelerationStructureKHR")))),
		copyAccelerationStructure:            (C.PFN_vkCopyAccelerationStructureKHR)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkCopyAccelerationStructureKHR")))),
		copyAccelerationStructureToMemory:    (C.PFN_vkCopyAccelerationStructureToMemoryKHR)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkCopyAccelerationStructureToMemoryKHR")))),
		copyMemoryToAccelerationStructure:    (C.PFN_vkCopyMemoryToAccelerationStructureKHR)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkCopyMemoryToAccelerationStructureKHR")))),

		cmdWriteAccelerationStructureProperties:     (C.PFN_vkCmdWriteAccelerationStructuresPropertiesKHR)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkCmdWriteAccelerationStructuresPropertiesKHR")))),
		getAccelerationStructureBuildSizes:          (C.PFN_vkGetAccelerationStructureBuildSizesKHR)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkGetAccelerationStructureBuildSizesKHR")))),
		getAccelerationStructureDeviceAddresses:     (C.PFN_vkGetAccelerationStructureDeviceAddressKHR)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkGetAccelerationStructureDeviceAddressKHR")))),
		getDeviceAccelerationStructureCompatibility: (C.PFN_vkGetDeviceAccelerationStructureCompatibilityKHR)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkGetDeviceAccelerationStructureCompatibilityKHR")))),
		writeAccelerationStructuresProperties:       (C.PFN_vkWriteAccelerationStructuresPropertiesKHR)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkWriteAccelerationStructuresPropertiesKHR")))),
	}
}

func (c *CLoader) VkBuildAccelerationStructuresKHR(
	device loader.VkDevice,
	deferredOperation khr_deferred_host_operations_loader.VkDeferredOperationKHR,
	infoCount loader.Uint32,
	pInfos *VkAccelerationStructureBuildGeometryInfoKHR,
	ppBuildRangeInfos **VkAccelerationStructureBuildRangeInfoKHR,
) (common.VkResult, error) {
	if c.buildAccelerationStructures == nil {
		panic("attempt to call extension method vkBuildAccelerationStructuresKHR when extension not present")
	}

	res := common.VkResult(C.cgoBuildAccelerationStructuresKHR(
		c.buildAccelerationStructures,
		C.VkDevice(unsafe.Pointer(device)),
		C.VkDeferredOperationKHR(unsafe.Pointer(deferredOperation)),
		C.uint32_t(infoCount),
		(*C.VkAccelerationStructureBuildGeometryInfoKHR)(unsafe.Pointer(pInfos)),
		(**C.VkAccelerationStructureBuildRangeInfoKHR)(unsafe.Pointer(ppBuildRangeInfos)),
	))
	return res, res.ToError()
}

func (c *CLoader) VkCmdBuildAccelerationStructuresIndirectKHR(
	commandBuffer loader.VkCommandBuffer,
	infoCount loader.Uint32,
	pInfos *VkAccelerationStructureBuildGeometryInfoKHR,
	pIndirectDeviceAddresses *loader.VkDeviceAddress,
	pIndirectStrides *loader.Uint32,
	ppMaxPrimitiveCounts **loader.Uint32,
) {
	if c.cmdBuildAccelerationStructuresIndirect == nil {
		panic("attempt to call extension method vkCmdBuildAccelerationStructuresIndirectKHR when extension not present")
	}

	C.cgoCmdBuildAccelerationStructuresIndirectKHR(
		c.cmdBuildAccelerationStructuresIndirect,
		C.VkCommandBuffer(unsafe.Pointer(commandBuffer)),
		C.uint32_t(infoCount),
		(*C.VkAccelerationStructureBuildGeometryInfoKHR)(unsafe.Pointer(pInfos)),
		(*C.VkDeviceAddress)(unsafe.Pointer(pIndirectDeviceAddresses)),
		(*C.uint32_t)(unsafe.Pointer(pIndirectStrides)),
		(**C.uint32_t)(unsafe.Pointer(ppMaxPrimitiveCounts)),
	)
}

func (c *CLoader) VkCmdBuildAccelerationStructuresKHR(
	commandBuffer loader.VkCommandBuffer,
	infoCount loader.Uint32,
	pInfos *VkAccelerationStructureBuildGeometryInfoKHR,
	ppBuildRangeInfos **VkAccelerationStructureBuildRangeInfoKHR,
) {
	if c.cmdBuildAccelerationStructures == nil {
		panic("attempt to call extension method vkCmdBuildAccelerationStructuresKHR when extension not present")
	}

	C.cgoCmdBuildAccelerationStructuresKHR(
		c.cmdBuildAccelerationStructures,
		C.VkCommandBuffer(unsafe.Pointer(commandBuffer)),
		C.uint32_t(infoCount),
		(*C.VkAccelerationStructureBuildGeometryInfoKHR)(unsafe.Pointer(pInfos)),
		(**C.VkAccelerationStructureBuildRangeInfoKHR)(unsafe.Pointer(ppBuildRangeInfos)),
	)
}

func (c *CLoader) VkCmdCopyAccelerationStructureKHR(
	commandBuffer loader.VkCommandBuffer,
	pInfo *VkCopyAccelerationStructureInfoKHR,
) {
	if c.cmdCopyAccelerationStructure == nil {
		panic("attempt to call extension method vkCmdCopyAccelerationStructureKHR when extension not present")
	}

	C.cgoCmdCopyAccelerationStructureKHR(
		c.cmdCopyAccelerationStructure,
		C.VkCommandBuffer(unsafe.Pointer(commandBuffer)),
		(*C.VkCopyAccelerationStructureInfoKHR)(unsafe.Pointer(pInfo)),
	)
}

func (c *CLoader) VkCmdCopyAccelerationStructureToMemoryKHR(
	commandBuffer loader.VkCommandBuffer,
	pInfo *VkCopyAccelerationStructureToMemoryInfoKHR,
) {
	if c.cmdCopyAccelerationStructureToMemory == nil {
		panic("attempt to call extension method vkCmdCopyAccelerationStructureToMemoryKHR when extension not present")
	}

	C.cgoCmdCopyAccelerationStructureToMemoryKHR(
		c.cmdCopyAccelerationStructureToMemory,
		C.VkCommandBuffer(unsafe.Pointer(commandBuffer)),
		(*C.VkCopyAccelerationStructureToMemoryInfoKHR)(pInfo),
	)
}

func (c *CLoader) VkCmdCopyMemoryToAccelerationStructureKHR(
	commandBuffer loader.VkCommandBuffer,
	pInfo *VkCopyMemoryToAccelerationStructureInfoKHR,
) {
	if c.cmdCopyMemoryToAccelerationStructure == nil {
		panic("attempt to call extension method vkCmdCopyMemoryToAccelerationStructureKHR when extension not present")
	}

	C.cgoCmdCopyMemoryToAccelerationStructureKHR(
		c.cmdCopyMemoryToAccelerationStructure,
		C.VkCommandBuffer(unsafe.Pointer(commandBuffer)),
		(*C.VkCopyMemoryToAccelerationStructureInfoKHR)(unsafe.Pointer(pInfo)),
	)
}

func (c *CLoader) VkCmdWriteAccelerationStructuresPropertiesKHR(
	commandBuffer loader.VkCommandBuffer,
	accelerationStructureCount loader.Uint32,
	pAccelerationStructures *VkAccelerationStructureKHR,
	queryType loader.VkQueryType,
	queryPool loader.VkQueryPool,
	firstQuery loader.Uint32,
) {
	if c.cmdWriteAccelerationStructureProperties == nil {
		panic("attempt to call extension method vkCmdWriteAccelerationStructuresPropertiesKHR when extension not present")
	}

	C.cgoCmdWriteAccelerationStructuresPropertiesKHR(
		c.cmdWriteAccelerationStructureProperties,
		C.VkCommandBuffer(unsafe.Pointer(commandBuffer)),
		C.uint32_t(accelerationStructureCount),
		(*C.VkAccelerationStructureKHR)(unsafe.Pointer(pAccelerationStructures)),
		C.VkQueryType(queryType),
		C.VkQueryPool(unsafe.Pointer(queryPool)),
		C.uint32_t(firstQuery),
	)
}

func (c *CLoader) VkCopyAccelerationStructureKHR(
	device loader.VkDevice,
	deferredOperation khr_deferred_host_operations_loader.VkDeferredOperationKHR,
	pInfo *VkCopyAccelerationStructureInfoKHR,
) (common.VkResult, error) {
	if c.copyAccelerationStructure == nil {
		panic("attempt to call extension method vkCopyAccelerationStructureKHR when extension not present")
	}

	res := common.VkResult(C.cgoCopyAccelerationStructureKHR(
		c.copyAccelerationStructure,
		C.VkDevice(unsafe.Pointer(device)),
		C.VkDeferredOperationKHR(unsafe.Pointer(deferredOperation)),
		(*C.VkCopyAccelerationStructureInfoKHR)(unsafe.Pointer(pInfo)),
	))

	return res, res.ToError()
}

func (c *CLoader) VkCopyAccelerationStructureToMemoryKHR(
	device loader.VkDevice,
	deferredOperation khr_deferred_host_operations_loader.VkDeferredOperationKHR,
	pInfo *VkCopyAccelerationStructureToMemoryInfoKHR,
) (common.VkResult, error) {
	if c.copyAccelerationStructureToMemory == nil {
		panic("attempt to call extension method vkCopyAccelerationStructureToMemoryKHR when extension not present")
	}

	res := common.VkResult(C.cgoCopyAccelerationStructureToMemoryKHR(
		c.copyAccelerationStructureToMemory,
		C.VkDevice(unsafe.Pointer(device)),
		C.VkDeferredOperationKHR(unsafe.Pointer(deferredOperation)),
		(*C.VkCopyAccelerationStructureToMemoryInfoKHR)(unsafe.Pointer(pInfo)),
	))

	return res, res.ToError()
}

func (c *CLoader) VkCopyMemoryToAccelerationStructureKHR(
	device loader.VkDevice,
	deferredOperation khr_deferred_host_operations_loader.VkDeferredOperationKHR,
	pInfo *VkCopyMemoryToAccelerationStructureInfoKHR,
) (common.VkResult, error) {
	if c.copyMemoryToAccelerationStructure == nil {
		panic("attempt to call extension method vkCopyMemoryToAccelerationStructureKHR when extension not present")
	}

	res := common.VkResult(C.cgoCopyMemoryToAccelerationStructureKHR(
		c.copyMemoryToAccelerationStructure,
		C.VkDevice(unsafe.Pointer(device)),
		C.VkDeferredOperationKHR(unsafe.Pointer(deferredOperation)),
		(*C.VkCopyMemoryToAccelerationStructureInfoKHR)(pInfo),
	))

	return res, res.ToError()
}

func (c *CLoader) VkCreateAccelerationStructureKHR(
	device loader.VkDevice,
	pCreateInfo *VkAccelerationStructureCreateInfoKHR,
	pAllocator *loader.VkAllocationCallbacks,
	pAccelerationStructure *VkAccelerationStructureKHR,
) (common.VkResult, error) {
	if c.createAccelerationStructure == nil {
		panic("attempt to call extension method vkCreateAccelerationStructureKHR when extension not present")
	}

	res := common.VkResult(C.cgoCreateAccelerationStructureKHR(
		c.createAccelerationStructure,
		C.VkDevice(unsafe.Pointer(device)),
		(*C.VkAccelerationStructureCreateInfoKHR)(unsafe.Pointer(pCreateInfo)),
		(*C.VkAllocationCallbacks)(unsafe.Pointer(pAllocator)),
		(*C.VkAccelerationStructureKHR)(unsafe.Pointer(pAccelerationStructure)),
	))

	return res, res.ToError()
}

func (c *CLoader) VkDestroyAccelerationStructureKHR(
	device loader.VkDevice,
	accelerationStructure VkAccelerationStructureKHR,
	pAllocator *loader.VkAllocationCallbacks,
) {
	if c.destroyAccelerationStructure == nil {
		panic("attempt to call extension method vkDestroyAccelerationStructureKHR when extension not present")
	}

	C.cgoDestroyAccelerationStructureKHR(
		c.destroyAccelerationStructure,
		C.VkDevice(unsafe.Pointer(device)),
		C.VkAccelerationStructureKHR(unsafe.Pointer(accelerationStructure)),
		(*C.VkAllocationCallbacks)(unsafe.Pointer(pAllocator)),
	)
}

func (c *CLoader) VkGetAccelerationStructureBuildSizesKHR(
	device loader.VkDevice,
	buildType VkAccelerationStructureBuildTypeKHR,
	pBuildInfo *VkAccelerationStructureBuildGeometryInfoKHR,
	pMaxPrimitiveCounts *loader.Uint32,
	pSizeInfo *VkAccelerationStructureBuildSizesInfoKHR,
) {
	if c.getAccelerationStructureBuildSizes == nil {
		panic("attempt to call extension method vkGetAccelerationStructureBuildSizesKHR when extension not present")
	}

	C.cgoGetAccelerationStructureBuildSizesKHR(
		c.getAccelerationStructureBuildSizes,
		C.VkDevice(unsafe.Pointer(device)),
		C.VkAccelerationStructureBuildTypeKHR(buildType),
		(*C.VkAccelerationStructureBuildGeometryInfoKHR)(unsafe.Pointer(pBuildInfo)),
		(*C.uint32_t)(unsafe.Pointer(pMaxPrimitiveCounts)),
		(*C.VkAccelerationStructureBuildSizesInfoKHR)(unsafe.Pointer(pSizeInfo)),
	)
}

func (c *CLoader) VkGetAccelerationStructureDeviceAddressKHR(
	device loader.VkDevice,
	pInfo *VkAccelerationStructureDeviceAddressInfoKHR,
) loader.VkDeviceAddress {
	if c.getAccelerationStructureDeviceAddresses == nil {
		panic("attempt to call extension method vkGetAccelerationStructureDeviceAddressKHR when extension not present")
	}

	return loader.VkDeviceAddress(C.cgoGetAccelerationStructureDeviceAddressKHR(
		c.getAccelerationStructureDeviceAddresses,
		C.VkDevice(unsafe.Pointer(device)),
		(*C.VkAccelerationStructureDeviceAddressInfoKHR)(unsafe.Pointer(pInfo)),
	))
}

func (c *CLoader) VkGetDeviceAccelerationStructureCompatibilityKHR(
	device loader.VkDevice,
	pVersionInfo *VkAccelerationStructureVersionInfoKHR,
	pCompatibility *VkAccelerationStructureCompatibilityKHR,
) {
	if c.getDeviceAccelerationStructureCompatibility == nil {
		panic("attempt to call extension method vkGetDeviceAccelerationStructureCompatibilityKHR when extension not present")
	}

	C.cgoGetDeviceAccelerationStructureCompatibilityKHR(
		c.getDeviceAccelerationStructureCompatibility,
		C.VkDevice(unsafe.Pointer(device)),
		(*C.VkAccelerationStructureVersionInfoKHR)(unsafe.Pointer(pVersionInfo)),
		(*C.VkAccelerationStructureCompatibilityKHR)(unsafe.Pointer(pCompatibility)),
	)
}

func (c *CLoader) VkWriteAccelerationStructuresPropertiesKHR(
	device loader.VkDevice,
	accelerationStructureCount loader.Uint32,
	pAccelerationStructures *VkAccelerationStructureKHR,
	queryType loader.VkQueryType,
	dataSize loader.Size,
	pData unsafe.Pointer,
	stride loader.Size,
) (common.VkResult, error) {
	if c.writeAccelerationStructuresProperties == nil {
		panic("attempt to call extension method vkWriteAccelerationStructuresPropertiesKHR when extension not present")
	}

	res := common.VkResult(C.cgoWriteAccelerationStructuresPropertiesKHR(
		c.writeAccelerationStructuresProperties,
		C.VkDevice(unsafe.Pointer(device)),
		C.uint32_t(accelerationStructureCount),
		(*C.VkAccelerationStructureKHR)(unsafe.Pointer(pAccelerationStructures)),
		C.VkQueryType(queryType),
		C.size_t(dataSize),
		pData,
		C.size_t(stride),
	))

	return res, res.ToError()
}
