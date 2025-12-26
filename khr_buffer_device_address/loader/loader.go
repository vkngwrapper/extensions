package khr_buffer_device_address_loader

/*
#include <stdlib.h>
#include "../../vulkan/vulkan.h"

VkDeviceAddress cgoGetBufferDeviceAddressKHR(PFN_vkGetBufferDeviceAddressKHR fn, VkDevice device, VkBufferDeviceAddressInfoKHR *pInfo) {
	return fn(device, pInfo);
}

uint64_t cgoGetBufferOpaqueCaptureAddressKHR(PFN_vkGetBufferOpaqueCaptureAddressKHR fn, VkDevice device, VkBufferDeviceAddressInfoKHR *pInfo) {
	return fn(device, pInfo);
}

uint64_t cgoGetDeviceMemoryOpaqueCaptureAddressKHR(PFN_vkGetDeviceMemoryOpaqueCaptureAddressKHR fn, VkDevice device, VkDeviceMemoryOpaqueCaptureAddressInfoKHR *pInfo) {
	return fn(device, pInfo);
}
*/
import "C"
import (
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3/loader"
	_ "github.com/vkngwrapper/extensions/v3/vulkan"
)

//go:generate mockgen -source loader.go -destination ../mocks/loader.go -package mock_buffer_device_address

type Loader interface {
	VkGetBufferDeviceAddressKHR(device loader.VkDevice, pInfo *VkBufferDeviceAddressInfoKHR) loader.VkDeviceAddress
	VkGetBufferOpaqueCaptureAddressKHR(device loader.VkDevice, pInfo *VkBufferDeviceAddressInfoKHR) loader.Uint64
	VkGetDeviceMemoryOpaqueCaptureAddressKHR(device loader.VkDevice, pInfo *VkDeviceMemoryOpaqueCaptureAddressInfoKHR) loader.Uint64
}

type VkBufferDeviceAddressInfoKHR C.VkBufferDeviceAddressInfoKHR
type VkDeviceMemoryOpaqueCaptureAddressInfoKHR C.VkDeviceMemoryOpaqueCaptureAddressInfoKHR
type VkBufferOpaqueCaptureAddressCreateInfoKHR C.VkBufferOpaqueCaptureAddressCreateInfoKHR
type VkMemoryOpaqueCaptureAddressAllocateInfoKHR C.VkMemoryOpaqueCaptureAddressAllocateInfoKHR
type VkPhysicalDeviceBufferDeviceAddressFeaturesKHR C.VkPhysicalDeviceBufferDeviceAddressFeaturesKHR

type CLoader struct {
	coreLoader loader.Loader

	getBufferDeviceAddress              C.PFN_vkGetBufferDeviceAddressKHR
	getBufferOpaqueCaptureAddress       C.PFN_vkGetBufferOpaqueCaptureAddressKHR
	getDeviceMemoryOpaqueCaptureAddress C.PFN_vkGetDeviceMemoryOpaqueCaptureAddressKHR
}

func CreateLoaderFromCore(coreLoader loader.Loader) *CLoader {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	return &CLoader{
		coreLoader: coreLoader,

		getBufferDeviceAddress:              (C.PFN_vkGetBufferDeviceAddressKHR)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkGetBufferDeviceAddressKHR")))),
		getBufferOpaqueCaptureAddress:       (C.PFN_vkGetBufferOpaqueCaptureAddressKHR)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkGetBufferOpaqueCaptureAddressKHR")))),
		getDeviceMemoryOpaqueCaptureAddress: (C.PFN_vkGetDeviceMemoryOpaqueCaptureAddressKHR)(coreLoader.LoadProcAddr((*loader.Char)(arena.CString("vkGetDeviceMemoryOpaqueCaptureAddressKHR")))),
	}
}

func (d *CLoader) VkGetBufferDeviceAddressKHR(device loader.VkDevice, pInfo *VkBufferDeviceAddressInfoKHR) loader.VkDeviceAddress {
	if d.getBufferDeviceAddress == nil {
		panic("attempt to call extension method vkGetBufferDeviceAddressKHR when extension not present")
	}

	return loader.VkDeviceAddress(C.cgoGetBufferDeviceAddressKHR(
		d.getBufferDeviceAddress,
		C.VkDevice(unsafe.Pointer(device)),
		(*C.VkBufferDeviceAddressInfoKHR)(pInfo),
	))
}

func (d *CLoader) VkGetBufferOpaqueCaptureAddressKHR(device loader.VkDevice, pInfo *VkBufferDeviceAddressInfoKHR) loader.Uint64 {
	if d.getBufferOpaqueCaptureAddress == nil {
		panic("attempt to call extension method vkGetBufferOpaqueCaptureAddressKHR when extension not present")
	}

	return loader.Uint64(C.cgoGetBufferOpaqueCaptureAddressKHR(
		d.getBufferOpaqueCaptureAddress,
		C.VkDevice(unsafe.Pointer(device)),
		(*C.VkBufferDeviceAddressInfoKHR)(pInfo),
	))
}

func (d *CLoader) VkGetDeviceMemoryOpaqueCaptureAddressKHR(device loader.VkDevice, pInfo *VkDeviceMemoryOpaqueCaptureAddressInfoKHR) loader.Uint64 {
	if d.getDeviceMemoryOpaqueCaptureAddress == nil {
		panic("attempt to call extension method vkGetDeviceMemoryOpaqueCaptureAddressKHR when extension not present")
	}

	return loader.Uint64(C.cgoGetDeviceMemoryOpaqueCaptureAddressKHR(
		d.getDeviceMemoryOpaqueCaptureAddress,
		C.VkDevice(unsafe.Pointer(device)),
		(*C.VkDeviceMemoryOpaqueCaptureAddressInfoKHR)(pInfo),
	))
}
