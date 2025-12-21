package khr_swapchain

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/pkg/errors"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
)

// PresentOutData represents optionally-returned data from the Extension.QueuePresent command
type PresentOutData struct {
	// Results is a slice of status codes, one for each Swapchain provided in PresentInfo.
	// This slice allows the caller to inspect the results of each individual Swapchain presentation
	// executed by Extension.QueuePresent
	Results []common.VkResult
}

// PresentInfo describes parameters of a Queue presentation
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentInfoKHR.html
type PresentInfo struct {
	// WaitSemaphores is a slice of Semaphore objects to wait for before issuing the present request
	WaitSemaphores []core1_0.Semaphore
	// Swapchains is a slice of Swapchain objects being presented to by this command
	Swapchains []Swapchain
	// ImageIndices is a slice of indices into the array of each Swapchain object's presentable Image objects.
	// Each entry in this slice identifies the Image to present on the corresponding entry in the Swapchains
	// slice.
	ImageIndices []int

	common.NextOptions

	// OutData is a struct whose contents will be populated by the present command. It may be left nil
	// if the caller is uninterested in this data.
	OutData *PresentOutData
}

func (o PresentInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkPresentInfoKHR)
	}
	if len(o.Swapchains) != len(o.ImageIndices) {
		return nil, errors.Errorf("present: specified %d swapchains and %d image indices, but they should match", len(o.Swapchains), len(o.ImageIndices))
	}

	createInfo := (*C.VkPresentInfoKHR)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_PRESENT_INFO_KHR
	createInfo.pNext = next

	waitSemaphoreCount := len(o.WaitSemaphores)
	createInfo.waitSemaphoreCount = C.uint32_t(waitSemaphoreCount)
	createInfo.pWaitSemaphores = nil
	if waitSemaphoreCount > 0 {
		semaphorePtr := (*C.VkSemaphore)(allocator.Malloc(waitSemaphoreCount * int(unsafe.Sizeof([1]C.VkSemaphore{}))))
		semaphoreSlice := ([]C.VkSemaphore)(unsafe.Slice(semaphorePtr, waitSemaphoreCount))

		for i := 0; i < waitSemaphoreCount; i++ {
			if o.WaitSemaphores[i] == nil {
				return nil, errors.Errorf("khr_swapchain.PresentInfo.WaitSemaphores cannot contain nil "+
					"elements, but element %d is nil", i)
			}
			semaphoreHandle := (C.VkSemaphore)(unsafe.Pointer(o.WaitSemaphores[i].Handle()))
			semaphoreSlice[i] = semaphoreHandle
		}

		createInfo.pWaitSemaphores = semaphorePtr
	}

	swapchainCount := len(o.Swapchains)
	createInfo.swapchainCount = C.uint32_t(swapchainCount)
	createInfo.pSwapchains = nil
	createInfo.pImageIndices = nil
	createInfo.pResults = nil
	if swapchainCount > 0 {
		swapchainPtr := (*C.VkSwapchainKHR)(allocator.Malloc(swapchainCount * int(unsafe.Sizeof([1]C.VkSwapchainKHR{}))))
		swapchainSlice := ([]C.VkSwapchainKHR)(unsafe.Slice(swapchainPtr, swapchainCount))

		imageIndexPtr := (*C.uint32_t)(allocator.Malloc(swapchainCount * int(unsafe.Sizeof(C.uint32_t(0)))))
		imageIndexSlice := ([]C.uint32_t)(unsafe.Slice(imageIndexPtr, swapchainCount))

		resultPtr := (*C.VkResult)(allocator.Malloc(swapchainCount * int(unsafe.Sizeof(C.VkResult(0)))))

		for i := 0; i < swapchainCount; i++ {
			if o.Swapchains[i] == nil {
				return nil, errors.Errorf("khr_swapchain.PresentInfo.Swapchains cannot contain nil "+
					"elements, but element %d is nil", i)
			}
			swapchainSlice[i] = (C.VkSwapchainKHR)(unsafe.Pointer(o.Swapchains[i].Handle()))
			imageIndexSlice[i] = (C.uint32_t)(o.ImageIndices[i])
		}

		createInfo.pSwapchains = swapchainPtr
		createInfo.pImageIndices = imageIndexPtr
		createInfo.pResults = resultPtr
	}

	return preallocatedPointer, nil
}

func (o PresentInfo) PopulateOutData(cDataPointer unsafe.Pointer) error {
	createInfo := (*C.VkPresentInfoKHR)(cDataPointer)

	if o.OutData == nil {
		return nil
	}

	resultCount := len(o.Swapchains)
	o.OutData.Results = make([]common.VkResult, resultCount)

	resultSlice := ([]C.VkResult)(unsafe.Slice(createInfo.pResults, resultCount))
	for i := 0; i < resultCount; i++ {
		o.OutData.Results[i] = common.VkResult(resultSlice[i])
	}

	return nil
}
