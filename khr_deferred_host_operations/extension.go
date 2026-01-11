package khr_deferred_host_operations

import (
	"errors"

	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
	khr_deferred_host_operations_loader "github.com/vkngwrapper/extensions/v3/khr_deferred_host_operations/loader"
)

type VulkanExtensionDriver struct {
	loader khr_deferred_host_operations_loader.Loader

	device core1_0.Device
}

// CreateExtensionDriverFromCoreDriver produces an ExtensionDriver object from a Device with
// khr_deferred_host_operations loaded
func CreateExtensionDriverFromCoreDriver(driver core1_0.DeviceDriver) ExtensionDriver {
	device := driver.Device()
	if !device.IsDeviceExtensionActive(ExtensionName) {
		return nil
	}

	return &VulkanExtensionDriver{
		loader: khr_deferred_host_operations_loader.CreateLoaderFromCore(driver.Loader()),
		device: device,
	}
}

// CreateExtensionDriverFromLoader generates an ExtensionDriver from a loader.Loader object- this is usually
// used in tests to build an ExtensionDriver from mock drivers
func CreateExtensionDriverFromLoader(loader khr_deferred_host_operations_loader.Loader, device core1_0.Device) *VulkanExtensionDriver {
	return &VulkanExtensionDriver{
		loader: loader,
		device: device,
	}
}

func (v *VulkanExtensionDriver) CreateDeferredOperation(callbacks *loader.AllocationCallbacks) (DeferredOperation, common.VkResult, error) {
	var deferredOperation DeferredOperation
	var deferredOperationHandle khr_deferred_host_operations_loader.VkDeferredOperationKHR
	res, err := v.loader.VkCreateDeferredOperationKHR(
		v.device.Handle(),
		callbacks.Handle(),
		&deferredOperationHandle,
	)
	if err != nil {
		return deferredOperation, res, err
	}

	return DeferredOperation{
		handle:     deferredOperationHandle,
		device:     v.device.Handle(),
		apiVersion: v.device.APIVersion(),
	}, res, nil
}

func (v *VulkanExtensionDriver) DestroyDeferredOperation(operation DeferredOperation, callbacks *loader.AllocationCallbacks) {
	if !operation.Initialized() {
		panic("operation cannot be uninitialized")
	}

	v.loader.VkDestroyDeferredOperationKHR(
		v.device.Handle(),
		operation.Handle(),
		callbacks.Handle(),
	)
}

func (v *VulkanExtensionDriver) DeferredOperationJoin(operation DeferredOperation) (common.VkResult, error) {
	if !operation.Initialized() {
		return core1_0.VKErrorUnknown, errors.New("operation cannot be uninitialized")
	}

	return v.loader.VkDeferredOperationJoinKHR(
		v.device.Handle(),
		operation.Handle(),
	)
}

func (v *VulkanExtensionDriver) GetDeferredOperationResult(operation DeferredOperation) (common.VkResult, error) {
	if !operation.Initialized() {
		return core1_0.VKErrorUnknown, errors.New("operation cannot be uninitialized")
	}

	return v.loader.VkGetDeferredOperationResultKHR(
		v.device.Handle(),
		operation.Handle(),
	)
}

func (v *VulkanExtensionDriver) GetDeferredOperationMaxConcurrency(operation DeferredOperation) int {
	if !operation.Initialized() {
		panic("operation cannot be uninitialized")
	}

	return int(v.loader.VkGetDeferredOperationMaxConcurrencyKHR(
		v.device.Handle(),
		operation.Handle(),
	))
}
