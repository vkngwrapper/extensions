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

// VulkanExtensionDriver is an implementation of the ExtensionDriver interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanExtensionDriver struct {
	driver khr_timeline_semaphore_loader.Loader
	device core.Device
}

// CreateExtensionDriverFromCoreDriver produces an ExtensionDriver object from a Device with
// khr_timeline_semaphore loaded
func CreateExtensionDriverFromCoreDriver(driver core1_0.DeviceDriver) ExtensionDriver {
	device := driver.Device()
	if !device.IsDeviceExtensionActive(ExtensionName) {
		return nil
	}

	return &VulkanExtensionDriver{
		driver: khr_timeline_semaphore_loader.CreateLoaderFromCore(driver.Loader()),
		device: device,
	}
}

// CreateExtensionDriverFromLoader generates an ExtensionDriver from a loader.Loader object- this is usually
// used in tests to build an ExtensionDriver from mock drivers
func CreateExtensionDriverFromLoader(driver khr_timeline_semaphore_loader.Loader, device core.Device) *VulkanExtensionDriver {
	return &VulkanExtensionDriver{
		driver: driver,
		device: device,
	}
}

func (e *VulkanExtensionDriver) GetSemaphoreCounterValue(semaphore core.Semaphore) (uint64, common.VkResult, error) {
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

func (e *VulkanExtensionDriver) SignalSemaphore(o SemaphoreSignalInfo) (common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	signalPtr, err := common.AllocOptions(arena, o)
	if err != nil {
		return core1_0.VKErrorUnknown, err
	}

	return e.driver.VkSignalSemaphoreKHR(
		e.device.Handle(),
		(*khr_timeline_semaphore_loader.VkSemaphoreSignalInfoKHR)(signalPtr),
	)
}

func (e *VulkanExtensionDriver) WaitSemaphores(timeout time.Duration, o SemaphoreWaitInfo) (common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	waitPtr, err := common.AllocOptions(arena, o)
	if err != nil {
		return core1_0.VKErrorUnknown, err
	}

	return e.driver.VkWaitSemaphoresKHR(
		e.device.Handle(),
		(*khr_timeline_semaphore_loader.VkSemaphoreWaitInfoKHR)(waitPtr),
		loader.Uint64(common.TimeoutNanoseconds(timeout)),
	)
}
