package khr_timeline_semaphore

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/vkngwrapper/core/v3/core1_2"
)

type SemaphoreType = core1_2.SemaphoreType

////

type SemaphoreWaitFlags = core1_2.SemaphoreWaitFlags

////

const (
	// ExtensionName is "VK_KHR_timeline_semaphore"
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_timeline_semaphore.html
	ExtensionName string = C.VK_KHR_TIMELINE_SEMAPHORE_EXTENSION_NAME

	// SemaphoreTypeBinary specifies a binary Semaphore type that has a boolean payload
	// indicating whether the Semaphore is currently signaled or unsignaled
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreType.html
	SemaphoreTypeBinary SemaphoreType = C.VK_SEMAPHORE_TYPE_BINARY_KHR
	// SemaphoreTypeTimeline specifies a timeline Semaphore type that has a strictly
	// increasing 64-bit unsigned integer payload indicating whether the Semaphore is signaled
	// with respect to a particular reference value
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreType.html
	SemaphoreTypeTimeline SemaphoreType = C.VK_SEMAPHORE_TYPE_TIMELINE_KHR

	// SemaphoreWaitAny specifies that the Semaphore wait condition is that at least one of
	// the Semaphore objects in SemaphoreWaitInfo.Semaphores has reached the value specified
	// by the corresponding element of SemaphoreWaitInfo.Values
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreWaitFlagBits.html
	SemaphoreWaitAny SemaphoreWaitFlags = C.VK_SEMAPHORE_WAIT_ANY_BIT_KHR
)

func init() {
	SemaphoreTypeBinary.Register("Binary")
	SemaphoreTypeTimeline.Register("Timeline")

	SemaphoreWaitAny.Register("Any")
}
