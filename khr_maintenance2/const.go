package khr_maintenance2

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/vkngwrapper/core/v2/core1_0"
)

// PointClippingBehavior specifies the point clipping behavior
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPointClippingBehavior.html
type PointClippingBehavior int32

var pointClippingBehaviorMapping = make(map[PointClippingBehavior]string)

func (e PointClippingBehavior) Register(str string) {
	pointClippingBehaviorMapping[e] = str
}

func (e PointClippingBehavior) String() string {
	return pointClippingBehaviorMapping[e]
}

////

// TessellationDomainOrigin describes tessellation domain origin
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkTessellationDomainOrigin.html
type TessellationDomainOrigin int32

var tessellationDomainOriginMapping = make(map[TessellationDomainOrigin]string)

func (e TessellationDomainOrigin) Register(str string) {
	tessellationDomainOriginMapping[e] = str
}

func (e TessellationDomainOrigin) String() string {
	return tessellationDomainOriginMapping[e]
}

////

const (
	// ExtensionName is "VK_KHR_maintenance2"
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_maintenance2.html
	ExtensionName string = C.VK_KHR_MAINTENANCE2_EXTENSION_NAME

	// ImageCreateBlockTexelViewCompatible specifies that the Image having a compressed format can be
	// used to create an ImageView with an uncompressed format where each texel in the ImageView
	// corresponds to a compressed texel block of the Image
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateFlagBits.html
	ImageCreateBlockTexelViewCompatible core1_0.ImageCreateFlags = C.VK_IMAGE_CREATE_BLOCK_TEXEL_VIEW_COMPATIBLE_BIT_KHR
	// ImageCreateExtendedUsage specifies that the Image can be created with usage flags that are not
	// supported for the format the Image is created with but are supported for at least one format
	// an ImageView created from this Image can have
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateFlagBits.html
	ImageCreateExtendedUsage core1_0.ImageCreateFlags = C.VK_IMAGE_CREATE_EXTENDED_USAGE_BIT_KHR

	// ImageLayoutDepthAttachmentStencilReadOnlyOptimal specifies a layout for depth/stencil format
	// Image objects allowing read and write access to the depth aspect as a depth attachment, and read-only
	// access to the stencil aspect as a stencil attachment or in shaders as a sampled Image, combined
	// Image/Sampler, or input attachment
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html
	ImageLayoutDepthAttachmentStencilReadOnlyOptimal core1_0.ImageLayout = C.VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_STENCIL_READ_ONLY_OPTIMAL_KHR
	// ImageLayoutDepthReadOnlyStencilAttachmentOptimal specifies a layout for depth/stencil format Image objects
	// allowing read and write access to the stencil aspect as a stencil attachment, and read-only access
	// to the depth aspect as a depth attachment or in shaders as a sampled Image, combined Image/Sampler,
	// or input attachment
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html
	ImageLayoutDepthReadOnlyStencilAttachmentOptimal core1_0.ImageLayout = C.VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_STENCIL_ATTACHMENT_OPTIMAL_KHR

	// PointClippingAllClipPlanes specifies that the primitive is discarded if the vertex lies
	// outside any clip plane, including the planes bounding the view volume
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPointClippingBehavior.html
	PointClippingAllClipPlanes PointClippingBehavior = C.VK_POINT_CLIPPING_BEHAVIOR_ALL_CLIP_PLANES_KHR
	// PointClippingUserClipPlanesOnly specifies that the primitive is discarded only if the vertex
	// lies outside any user clip plane
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPointClippingBehavior.html
	PointClippingUserClipPlanesOnly PointClippingBehavior = C.VK_POINT_CLIPPING_BEHAVIOR_USER_CLIP_PLANES_ONLY_KHR

	// TessellationDomainOriginUpperLeft specifies that the origin of the domain space
	// is in the upper left corner, as shown in figure
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkTessellationDomainOrigin.html
	TessellationDomainOriginUpperLeft TessellationDomainOrigin = C.VK_TESSELLATION_DOMAIN_ORIGIN_UPPER_LEFT_KHR
	// TessellationDomainOriginLowerLeft specifies that the origin of the domain space
	// is in the lower left corner, as shown in figure
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkTessellationDomainOrigin.html
	TessellationDomainOriginLowerLeft TessellationDomainOrigin = C.VK_TESSELLATION_DOMAIN_ORIGIN_LOWER_LEFT_KHR
)

func init() {
	ImageCreateBlockTexelViewCompatible.Register("Block Texel View Compatible")
	ImageCreateExtendedUsage.Register("Extended Usage")

	ImageLayoutDepthReadOnlyStencilAttachmentOptimal.Register("Depth Read-Only Stencil Attachment Optimal")
	ImageLayoutDepthAttachmentStencilReadOnlyOptimal.Register("Depth Attachment Stencil Read-Only Optimal")

	PointClippingAllClipPlanes.Register("All Clip Planes")
	PointClippingUserClipPlanesOnly.Register("User Clip Planes Only")

	TessellationDomainOriginUpperLeft.Register("Upper Left")
	TessellationDomainOriginLowerLeft.Register("Lower Left")
}
