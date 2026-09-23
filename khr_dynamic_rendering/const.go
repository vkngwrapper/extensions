package khr_dynamic_rendering

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	_ "github.com/vkngwrapper/extensions/v3/vulkan"
)

const (
	// ExtensionName is "VK_KHR_dynamic_rendering"
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VK_KHR_dynamic_rendering.html
	ExtensionName string = C.VK_KHR_DYNAMIC_RENDERING_EXTENSION_NAME

	// AttachmentStoreOpNone specifies that the contents of the attachment are not stored
	// when rendering ends.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkAttachmentStoreOp.html
	AttachmentStoreOpNone core1_0.AttachmentStoreOp = C.VK_ATTACHMENT_STORE_OP_NONE_KHR
)

// RenderingFlags specifies additional parameters for rendering
//
// https://docs.vulkan.org/refpages/latest/refpages/source/VkRenderingFlagBitsKHR.html
type RenderingFlags int32

var renderingFlagsMapping = common.NewFlagStringMapping[RenderingFlags]()

func (f RenderingFlags) Register(str string) {
	renderingFlagsMapping.Register(f, str)
}

func (f RenderingFlags) String() string {
	return renderingFlagsMapping.FlagsToString(f)
}

const (
	// RenderingContentsSecondaryCommandBuffers specifies that rendering commands are recorded
	// in secondary CommandBuffer objects.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkRenderingFlagBitsKHR.html
	RenderingContentsSecondaryCommandBuffers RenderingFlags = C.VK_RENDERING_CONTENTS_SECONDARY_COMMAND_BUFFERS_BIT_KHR
	// RenderingSuspending specifies that rendering is suspended when rendering ends.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkRenderingFlagBitsKHR.html
	RenderingSuspending RenderingFlags = C.VK_RENDERING_SUSPENDING_BIT_KHR
	// RenderingResuming specifies that rendering resumes a previously suspended render pass.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkRenderingFlagBitsKHR.html
	RenderingResuming RenderingFlags = C.VK_RENDERING_RESUMING_BIT_KHR
)

func init() {
	RenderingContentsSecondaryCommandBuffers.Register("ContentsSecondaryCommandBuffers")
	RenderingSuspending.Register("Suspending")
	RenderingResuming.Register("Resuming")
	AttachmentStoreOpNone.Register("None")
}
