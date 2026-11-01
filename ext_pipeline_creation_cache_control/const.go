package ext_pipeline_creation_cache_control

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
)

const (
	// ExtensionName is "VK_EXT_pipeline_creation_cache_control"
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VK_EXT_pipeline_creation_cache_control.html
	ExtensionName string = C.VK_EXT_PIPELINE_CREATION_CACHE_CONTROL_EXTENSION_NAME
)

// PipelineCacheCreateFlags specifies the behavior of the pipeline cache
//
// https://docs.vulkan.org/refpages/latest/refpages/source/VkPipelineCacheCreateFlagBits.html
type PipelineCacheCreateFlags uint32

var pipelineCacheCreateFlags = common.NewFlagStringMapping[PipelineCacheCreateFlags]()

func (f PipelineCacheCreateFlags) Register(str string) {
	pipelineCacheCreateFlags.Register(f, str)
}

func (f PipelineCacheCreateFlags) String() string {
	return pipelineCacheCreateFlags.FlagsToString(f)
}

const (
	// PipelineCacheCreateExternallySynchronized specifies that all commands that modify the
	// created PipelineCache will be externally synchronized. When set, the implementation may
	// skip any unnecessary processing needed to support simultaneous modification from multiple
	// threads where allowed.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkPipelineCacheCreateFlagBits.html
	PipelineCacheCreateExternallySynchronized PipelineCacheCreateFlags = C.VK_PIPELINE_CACHE_CREATE_EXTERNALLY_SYNCHRONIZED_BIT_EXT

	// PipelineCreateEarlyReturnOnFailure specifies that control will be returned to the
	// application if any individual pipeline returns a result which is not VKSuccess rather
	// than continuing to create additional pipelines.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkPipelineCreateFlagBits.html
	PipelineCreateEarlyReturnOnFailure core1_0.PipelineCreateFlags = C.VK_PIPELINE_CREATE_EARLY_RETURN_ON_FAILURE_BIT_EXT

	// PipelineCreateFailOnPipelineCompileRequired specifies that pipeline creation will fail
	// if a compile is required for creation of a valid Pipeline object;
	// VKErrorPipelineCompileRequired will be returned by pipeline creation, and the Pipeline will be
	// nil.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkResult.html
	PipelineCreateFailOnPipelineCompileRequired core1_0.PipelineCreateFlags = C.VK_PIPELINE_CREATE_FAIL_ON_PIPELINE_COMPILE_REQUIRED_BIT_EXT

	// VKPipelineCompileRequired indicates that a requested pipeline would have required
	// compilation, but the application requested compilation to not be performed.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkResult.html
	VKPipelineCompileRequired common.VkResult = C.VK_PIPELINE_COMPILE_REQUIRED_EXT
)

func init() {
	PipelineCacheCreateExternallySynchronized.Register("ExternallySynchronized")

	PipelineCreateEarlyReturnOnFailure.Register("EarlyReturnOnFailure")
	PipelineCreateFailOnPipelineCompileRequired.Register("FailOnPipelineCompileRequired")

	VKPipelineCompileRequired.Register("Pipeline Compile Required")
}
