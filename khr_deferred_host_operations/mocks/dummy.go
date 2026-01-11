package mock_deferred_host_operations

import (
	"math/rand"
	"unsafe"

	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/extensions/v3/khr_deferred_host_operations"
	khr_deferred_host_operations_loader "github.com/vkngwrapper/extensions/v3/khr_deferred_host_operations/loader"
)

func NewDummyDeferredOperation(device core1_0.Device) khr_deferred_host_operations.DeferredOperation {
	return khr_deferred_host_operations.InternalDeferredOperation(device.Handle(), NewFakeDeferredOperationHandle(), device.APIVersion())
}

func fakePointer() unsafe.Pointer {
	return unsafe.Pointer(uintptr(rand.Int()))
}

func NewFakeDeferredOperationHandle() khr_deferred_host_operations_loader.VkDeferredOperationKHR {
	return khr_deferred_host_operations_loader.VkDeferredOperationKHR(fakePointer())
}
