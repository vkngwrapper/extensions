package khr_deferred_host_operations

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import "github.com/vkngwrapper/core/v3/common"

const (
	// ExtensionName is "VK_KHR_deferred_host_operations"
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VK_KHR_deferred_host_operations.html
	ExtensionName string = C.VK_KHR_DEFERRED_HOST_OPERATIONS_EXTENSION_NAME

	// VkOperationDeferred indicates that a deferred operation was requested and at least some of the work
	// was deferred.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkResult.html
	VkOperationDeferred common.VkResult = C.VK_OPERATION_DEFERRED_KHR
	// VkOperationNotDeferred indicates that a operation was requested and no operations were deferred.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkResult.html
	VkOperationNotDeferred common.VkResult = C.VK_OPERATION_NOT_DEFERRED_KHR
	// VkThreadDone indicates that a deferred operation is not complete but there is no work remaining to assign
	// to additional threads.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkResult.html
	VkThreadDone common.VkResult = C.VK_THREAD_DONE_KHR
	// VkThreadIdle indicates that a deferred operation is not complete but there is currently
	// no work for this thread to do at the time of this call.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkResult.html
	VkThreadIdle common.VkResult = C.VK_THREAD_IDLE_KHR
)

func init() {
	VkOperationDeferred.Register("deferred")
	VkOperationNotDeferred.Register("not deferred")
	VkThreadDone.Register("thread done")
	VkThreadIdle.Register("thread idle")
}
