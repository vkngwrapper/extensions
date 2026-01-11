package khr_deferred_host_operations

import (
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/loader"
)

//go:generate mockgen -source extiface.go -destination ./mocks/extension.go -package mock_deferred_host_operations

type ExtensionDriver interface {
	CreateDeferredOperation(callbacks *loader.AllocationCallbacks) (DeferredOperation, common.VkResult, error)
	DestroyDeferredOperation(operation DeferredOperation, callbacks *loader.AllocationCallbacks)
	DeferredOperationJoin(operation DeferredOperation) (common.VkResult, error)
	GetDeferredOperationResult(operation DeferredOperation) (common.VkResult, error)
	GetDeferredOperationMaxConcurrency(operation DeferredOperation) int
}
