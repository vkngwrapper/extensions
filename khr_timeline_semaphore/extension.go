package khr_timeline_semaphore

import (
	"time"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/driver"
	khr_timeline_semaphore_driver "github.com/vkngwrapper/extensions/v3/khr_timeline_semaphore/driver"
)

// VulkanExtension is an implementation of the Extension interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanExtension struct {
	driver khr_timeline_semaphore_driver.Driver
}

// CreateExtensionFromDevice produces an Extension object from a Device with
// khr_timeline_semaphore loaded
func CreateExtensionFromDevice(device core1_0.Device) *VulkanExtension {
	if !device.IsDeviceExtensionActive(ExtensionName) {
		return nil
	}

	return &VulkanExtension{
		driver: khr_timeline_semaphore_driver.CreateDriverFromCore(device.Driver()),
	}
}

// CreateExtensionFromDriver generates an Extension from a driver.Driver object- this is usually
// used in tests to build an Extension from mock drivers
func CreateExtensionFromDriver(driver khr_timeline_semaphore_driver.Driver) *VulkanExtension {
	return &VulkanExtension{
		driver: driver,
	}
}

func (e *VulkanExtension) SemaphoreCounterValue(semaphore core1_0.Semaphore) (uint64, common.VkResult, error) {
	if semaphore == nil {
		panic("semaphore cannot be nil")
	}

	var value driver.Uint64
	res, err := e.driver.VkGetSemaphoreCounterValueKHR(
		semaphore.DeviceHandle(),
		semaphore.Handle(),
		&value,
	)
	if err != nil {
		return 0, res, err
	}

	return uint64(value), res, nil
}

func (e *VulkanExtension) SignalSemaphore(device core1_0.Device, o SemaphoreSignalInfo) (common.VkResult, error) {
	if device == nil {
		panic("device cannot be nil")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	signalPtr, err := common.AllocOptions(arena, o)
	if err != nil {
		return core1_0.VKErrorUnknown, err
	}

	return e.driver.VkSignalSemaphoreKHR(
		device.Handle(),
		(*khr_timeline_semaphore_driver.VkSemaphoreSignalInfoKHR)(signalPtr),
	)
}

func (e *VulkanExtension) WaitSemaphores(device core1_0.Device, timeout time.Duration, o SemaphoreWaitInfo) (common.VkResult, error) {
	if device == nil {
		panic("device cannot be nil")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	waitPtr, err := common.AllocOptions(arena, o)
	if err != nil {
		return core1_0.VKErrorUnknown, err
	}

	return e.driver.VkWaitSemaphoresKHR(
		device.Handle(),
		(*khr_timeline_semaphore_driver.VkSemaphoreWaitInfoKHR)(waitPtr),
		driver.Uint64(common.TimeoutNanoseconds(timeout)),
	)
}
