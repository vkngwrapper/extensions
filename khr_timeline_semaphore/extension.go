package khr_timeline_semaphore

import (
	"time"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
	"github.com/vkngwrapper/extensions/v3/khr_timeline_semaphore/loader"
)

// VulkanExtension is an implementation of the Extension interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanExtension struct {
	driver khr_timeline_semaphore_loader.Loader
}

// CreateExtensionFromDevice produces an Extension object from a Device with
// khr_timeline_semaphore loaded
func CreateExtensionFromDevice(device core.Device) *VulkanExtension {
	if !device.IsDeviceExtensionActive(ExtensionName) {
		return nil
	}

	return &VulkanExtension{
		driver: khr_timeline_semaphore_loader.CreateLoaderFromCore(device.Driver()),
	}
}

// CreateExtensionFromDriver generates an Extension from a loader.Loader object- this is usually
// used in tests to build an Extension from mock drivers
func CreateExtensionFromDriver(driver khr_timeline_semaphore_loader.Loader) *VulkanExtension {
	return &VulkanExtension{
		driver: driver,
	}
}

func (e *VulkanExtension) SemaphoreCounterValue(semaphore core.Semaphore) (uint64, common.VkResult, error) {
	if semaphore.Handle() == 0 {
		panic("semaphore cannot be uninitialized")
	}

	var value loader.Uint64
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

func (e *VulkanExtension) SignalSemaphore(device core.Device, o SemaphoreSignalInfo) (common.VkResult, error) {
	if device.Handle() == 0 {
		panic("device cannot be uninitialized")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	signalPtr, err := common.AllocOptions(arena, o)
	if err != nil {
		return core1_0.VKErrorUnknown, err
	}

	return e.driver.VkSignalSemaphoreKHR(
		device.Handle(),
		(*khr_timeline_semaphore_loader.VkSemaphoreSignalInfoKHR)(signalPtr),
	)
}

func (e *VulkanExtension) WaitSemaphores(device core.Device, timeout time.Duration, o SemaphoreWaitInfo) (common.VkResult, error) {
	if device.Handle() == 0 {
		panic("device cannot be uninitialized")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	waitPtr, err := common.AllocOptions(arena, o)
	if err != nil {
		return core1_0.VKErrorUnknown, err
	}

	return e.driver.VkWaitSemaphoresKHR(
		device.Handle(),
		(*khr_timeline_semaphore_loader.VkSemaphoreWaitInfoKHR)(waitPtr),
		loader.Uint64(common.TimeoutNanoseconds(timeout)),
	)
}
