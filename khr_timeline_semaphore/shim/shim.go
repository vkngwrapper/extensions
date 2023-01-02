package khr_timeline_semaphore_shim

import (
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/core1_0"
	"github.com/vkngwrapper/core/v2/core1_2"
	"github.com/vkngwrapper/extensions/v2/khr_timeline_semaphore"
	"time"
)

//go:generate mockgen -source shim.go -destination ../mocks/shim.go -package mock_timeline_semaphore

// SemaphoreShim contains all commands for the khr_timeline_semaphore extension for semaphores
type SemaphoreShim interface {
	// CounterValue queries the current state of this timeline Semaphore
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetSemaphoreCounterValue.html
	CounterValue() (uint64, common.VkResult, error)
}

type VulkanSemaphoreShim struct {
	extension khr_timeline_semaphore.Extension
	semaphore core1_0.Semaphore
}

// Compiler check that VulkanSemaphoreShim satisfies SemaphoreShim
var _ SemaphoreShim = &VulkanSemaphoreShim{}

func NewSemaphoreShim(extension khr_timeline_semaphore.Extension, semaphore core1_0.Semaphore) *VulkanSemaphoreShim {
	if extension == nil {
		panic("extension cannot be nil")
	}
	if semaphore == nil {
		panic("semaphore cannot be nil")
	}
	return &VulkanSemaphoreShim{
		extension: extension,
		semaphore: semaphore,
	}
}

func (s *VulkanSemaphoreShim) CounterValue() (uint64, common.VkResult, error) {
	return s.extension.SemaphoreCounterValue(s.semaphore)
}

// DeviceShim contains all commands for the khr_timeline_semaphore extension for devices
type DeviceShim interface {
	// SignalSemaphore signals a timeline Semaphore on the host
	//
	// o - Contains information about the signal operation
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkSignalSemaphore.html
	SignalSemaphore(o core1_2.SemaphoreSignalInfo) (common.VkResult, error)
	// WaitSemaphores waits for timeline Semaphore objects on the host
	//
	// timeout - How long to wait before returning VKTimeout. May be common.NoTimeout to wait indefinitely.
	// The timeout is adjusted to the closest value allowed by the implementation timeout accuracy,
	// which may be substantially longer than the requested timeout.
	//
	// o - Contains information about the wait condition
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkWaitSemaphores.html
	WaitSemaphores(timeout time.Duration, o core1_2.SemaphoreWaitInfo) (common.VkResult, error)
}

type VulkanDeviceShim struct {
	extension khr_timeline_semaphore.Extension
	device    core1_0.Device
}

// Compiler check that VulkanDeviceShim satisfies DeviceShim
var _ DeviceShim = &VulkanDeviceShim{}

func NewDeviceShim(extension khr_timeline_semaphore.Extension, device core1_0.Device) *VulkanDeviceShim {
	if extension == nil {
		panic("extension cannot be nil")
	}
	if device == nil {
		panic("device cannot be nil")
	}
	return &VulkanDeviceShim{
		extension: extension,
		device:    device,
	}
}

func (s *VulkanDeviceShim) SignalSemaphore(o core1_2.SemaphoreSignalInfo) (common.VkResult, error) {
	return s.extension.SignalSemaphore(s.device, khr_timeline_semaphore.SemaphoreSignalInfo(o))
}

func (s *VulkanDeviceShim) WaitSemaphores(timeout time.Duration, o core1_2.SemaphoreWaitInfo) (common.VkResult, error) {
	inOptions := khr_timeline_semaphore.SemaphoreWaitInfo{
		Flags:       khr_timeline_semaphore.SemaphoreWaitFlags(o.Flags),
		Semaphores:  o.Semaphores,
		Values:      o.Values,
		NextOptions: o.NextOptions,
	}

	return s.extension.WaitSemaphores(s.device, timeout, inOptions)
}
