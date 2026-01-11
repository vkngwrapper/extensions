package khr_deferred_host_operations

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/loader"
	khr_deferred_host_operations_loader "github.com/vkngwrapper/extensions/v3/khr_deferred_host_operations/loader"
)

type DeferredOperation struct {
	handle khr_deferred_host_operations_loader.VkDeferredOperationKHR
	device loader.VkDevice

	apiVersion common.APIVersion
}

func (s DeferredOperation) Handle() khr_deferred_host_operations_loader.VkDeferredOperationKHR {
	return s.handle
}

func (s DeferredOperation) DeviceHandle() loader.VkDevice {
	return s.device
}

func (s DeferredOperation) APIVersion() common.APIVersion {
	return s.apiVersion
}

func (s DeferredOperation) Initialized() bool {
	return s.handle != 0
}

func InternalDeferredOperation(device loader.VkDevice, handle khr_deferred_host_operations_loader.VkDeferredOperationKHR, version common.APIVersion) DeferredOperation {
	return DeferredOperation{
		device:     device,
		handle:     handle,
		apiVersion: version,
	}
}
