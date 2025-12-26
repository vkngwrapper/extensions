package khr_timeline_semaphore

import (
	"time"

	"github.com/vkngwrapper/core/v3"
	"github.com/vkngwrapper/core/v3/common"
)

//go:generate mockgen -source extiface.go -destination ./mocks/extension.go -package mock_timeline_semaphore

// Extension contains all commands for the khr_timeline_semaphore extension
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_timeline_semaphore.html
type Extension interface {
	// SemaphoreCounterValue queries the current state of a timeline Semaphore
	//
	// semaphore - The Semaphore to query
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetSemaphoreCounterValue.html
	SemaphoreCounterValue(semaphore core.Semaphore) (uint64, common.VkResult, error)
	// SignalSemaphore signals a timeline Semaphore on the host
	//
	// device - The Device which owns the Semaphore being signaled
	//
	// o - Contains information about the signal operation
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkSignalSemaphore.html
	SignalSemaphore(device core.Device, o SemaphoreSignalInfo) (common.VkResult, error)
	// WaitSemaphores waits for timeline Semaphore objects on the host
	//
	// device - The Device which owns the Semaphore objects being waited on
	//
	// timeout - How long to wait before returning VKTimeout. May be common.NoTimeout to wait indefinitely.
	// The timeout is adjusted to the closest value allowed by the implementation timeout accuracy,
	// which may be substantially longer than the requested timeout.
	//
	// o - Contains information about the wait condition
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkWaitSemaphores.html
	WaitSemaphores(device core.Device, timeout time.Duration, o SemaphoreWaitInfo) (common.VkResult, error)
}
