package khr_surface

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/core1_0"
)

// SurfaceTransformFlags represents presentation transforms supported on a Device
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceTransformFlagBitsKHR.html
type SurfaceTransformFlags int32

var surfaceTransformsMapping = common.NewFlagStringMapping[SurfaceTransformFlags]()

func (f SurfaceTransformFlags) Register(str string) {
	surfaceTransformsMapping.Register(f, str)
}
func (f SurfaceTransformFlags) String() string {
	return surfaceTransformsMapping.FlagsToString(f)
}

////

// CompositeAlphaFlags represents alpha-compositing modes supported on a Device
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCompositeAlphaFlagBitsKHR.html
type CompositeAlphaFlags int32

var compositeAlphaModeMapping = make(map[CompositeAlphaFlags]string)

func (e CompositeAlphaFlags) Register(str string) {
	compositeAlphaModeMapping[e] = str
}

func (e CompositeAlphaFlags) String() string {
	return compositeAlphaModeMapping[e]
}

////

// PresentMode represents presentation modes supported for a Surface
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentModeKHR.html
type PresentMode int32

var presentModeMapping = make(map[PresentMode]string)

func (e PresentMode) Register(str string) {
	presentModeMapping[e] = str
}

func (e PresentMode) String() string {
	return presentModeMapping[e]
}

////

// ColorSpace represents the supported color space of the presentation engine
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkColorSpaceKHR.html
type ColorSpace int32

var colorSpaceMapping = make(map[ColorSpace]string)

func (e ColorSpace) Register(str string) {
	colorSpaceMapping[e] = str
}

func (e ColorSpace) String() string {
	return colorSpaceMapping[e]
}

////

const (
	// ExtensionName is "VK_KHR_surface"
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_surface.html
	ExtensionName string = C.VK_KHR_SURFACE_EXTENSION_NAME

	// ObjectTypeSurface specifies a Surface handle
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkObjectType.html
	ObjectTypeSurface core1_0.ObjectType = C.VK_OBJECT_TYPE_SURFACE_KHR

	// TransformIdentity specifies that Image content is presented without being transformed
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceTransformFlagBitsKHR.html
	TransformIdentity SurfaceTransformFlags = C.VK_SURFACE_TRANSFORM_IDENTITY_BIT_KHR
	// TransformRotate90 specifies that Image content is rotated 90 degrees clockwise
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceTransformFlagBitsKHR.html
	TransformRotate90 SurfaceTransformFlags = C.VK_SURFACE_TRANSFORM_ROTATE_90_BIT_KHR
	// TransformRotate180 specifies that Image content is rotated 180 degrees clockwise
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceTransformFlagBitsKHR.html
	TransformRotate180 SurfaceTransformFlags = C.VK_SURFACE_TRANSFORM_ROTATE_180_BIT_KHR
	// TransformRotate270 specifies that Image content is rotated 270 degrees clockwise
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceTransformFlagBitsKHR.html
	TransformRotate270 SurfaceTransformFlags = C.VK_SURFACE_TRANSFORM_ROTATE_270_BIT_KHR
	// TransformHorizontalMirror specifies that Image content is mirrored horizontally
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceTransformFlagBitsKHR.html
	TransformHorizontalMirror SurfaceTransformFlags = C.VK_SURFACE_TRANSFORM_HORIZONTAL_MIRROR_BIT_KHR
	// TransformHorizontalMirrorRotate90 specifies that Image content is mirrored horizontally, then rotated 90
	// degrees clockwise
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceTransformFlagBitsKHR.html
	TransformHorizontalMirrorRotate90 SurfaceTransformFlags = C.VK_SURFACE_TRANSFORM_HORIZONTAL_MIRROR_ROTATE_90_BIT_KHR
	// TransformHorizontalMirrorRotate180 specifies that Image content is mirrored horizontally, then rotated 180
	// degrees clockwise
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceTransformFlagBitsKHR.html
	TransformHorizontalMirrorRotate180 SurfaceTransformFlags = C.VK_SURFACE_TRANSFORM_HORIZONTAL_MIRROR_ROTATE_180_BIT_KHR
	// TransformHorizontalMirrorRotate270 specifies that Image content is mirrored horizontally, then rotated 270
	// degrees clockwise
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceTransformFlagBitsKHR.html
	TransformHorizontalMirrorRotate270 SurfaceTransformFlags = C.VK_SURFACE_TRANSFORM_HORIZONTAL_MIRROR_ROTATE_270_BIT_KHR
	// TransformInherit specifies that the presentation transform is not specified, and is instead determined
	// by platform-specific considerations and mechanisms outside Vulkan
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceTransformFlagBitsKHR.html
	TransformInherit SurfaceTransformFlags = C.VK_SURFACE_TRANSFORM_INHERIT_BIT_KHR

	// CompositeAlphaOpaque indicates that the Image is treated as if it has a constant alpha of 1.0
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCompositeAlphaFlagBitsKHR.html
	CompositeAlphaOpaque CompositeAlphaFlags = C.VK_COMPOSITE_ALPHA_OPAQUE_BIT_KHR
	// CompositeAlphaPreMultiplied indicates that the alpha component of the Image is respected in
	// the compositing process. The non-alpha components of the Image are expected to already be
	// multiplied by the alpha component by the application
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCompositeAlphaFlagBitsKHR.html
	CompositeAlphaPreMultiplied CompositeAlphaFlags = C.VK_COMPOSITE_ALPHA_PRE_MULTIPLIED_BIT_KHR
	// CompositeAlphaPostMultiplied indicates that the alpha component of the Image is respected in
	// the compositing process.  The non-alpha components of the Image are not expected to already be
	// multiplied by the alpha component in the application; the compositor will multiply the non-alpha
	// components of the Image by the alpha component during compositing
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCompositeAlphaFlagBitsKHR.html
	CompositeAlphaPostMultiplied CompositeAlphaFlags = C.VK_COMPOSITE_ALPHA_POST_MULTIPLIED_BIT_KHR
	// CompositeAlphaInherit indicates that the application is responsible for setting the composite alpha
	// blending mode using native window system commands
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCompositeAlphaFlagBitsKHR.html
	CompositeAlphaInherit CompositeAlphaFlags = C.VK_COMPOSITE_ALPHA_INHERIT_BIT_KHR

	// PresentModeImmediate specifies that the presentation engine does not wait for a vertical blanking
	// period to update the current Image, meaning this mode may result in visible tearing. The requests
	// are applied immediately.
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentModeKHR.html
	PresentModeImmediate PresentMode = C.VK_PRESENT_MODE_IMMEDIATE_KHR
	// PresentModeMailbox specifies that the presentation engine waits for the next vertical blanking period
	// to update the current Image. Tearing cannot be observed. An internal single-entry queue is used to
	// hold pending presentation requests. If the queue is full when a new presentation request is received,
	// the new request replaces the existing entry
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentModeKHR.html
	PresentModeMailbox PresentMode = C.VK_PRESENT_MODE_MAILBOX_KHR
	// PresentModeFIFO specifies that the presentation engine waits for the next vertical blanking period to
	// update the current Image. Tearing cannot be observed. An internal queue is used to hold pending presentation
	// requests. New requests are appended to the end of the queue, and one request is removed from the beginning of
	// the queue during each vertical blanking period in which the queue is non-empty. This value is required
	// to be supported.
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentModeKHR.html
	PresentModeFIFO PresentMode = C.VK_PRESENT_MODE_FIFO_KHR
	// PresentModeFIFORelaxed specifies that the presentation engine generally waits for the next vertical
	// blanking period to update the current Image. If a vertical blanking period has already passed since
	// the last update of the current Image, then the presentation engine does not wait for another
	// vertical blanking period for the update, meaning this mode may result in visible tearing in this
	// case. This mode is useful for reducing visual stutter with an application that will mostly
	// present a new Image before the next vertical blanking period, but may occasionally be late, and
	// present a new Image just after the next vertical blanking period. An internal queue is used to hold
	// pending presentation requests.
	PresentModeFIFORelaxed PresentMode = C.VK_PRESENT_MODE_FIFO_RELAXED_KHR

	// ColorSpaceSRGBNonlinear specifies support for the sRGB color space
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkColorSpaceKHR.html
	ColorSpaceSRGBNonlinear ColorSpace = C.VK_COLOR_SPACE_SRGB_NONLINEAR_KHR

	// VKErrorSurfaceLost indicates a surface is no longer available
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResult.html
	VKErrorSurfaceLost common.VkResult = C.VK_ERROR_SURFACE_LOST_KHR
	// VKErrorNativeWindowInUse indicates the requested window is already in use by Vulkan or
	// another API in a manner which prevents it from being used again
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResult.html
	VKErrorNativeWindowInUse common.VkResult = C.VK_ERROR_NATIVE_WINDOW_IN_USE_KHR
)

func init() {
	ObjectTypeSurface.Register("Surface")

	TransformIdentity.Register("Identity")
	TransformRotate90.Register("Rotate 90")
	TransformRotate180.Register("Rotate 180")
	TransformRotate270.Register("Rotate 270")
	TransformHorizontalMirror.Register("Horizontal Mirror")
	TransformHorizontalMirrorRotate90.Register("Horizontal Mirror & Rotate 90")
	TransformHorizontalMirrorRotate180.Register("Horizontal Mirror & Rotate 180")
	TransformHorizontalMirrorRotate270.Register("Horizontal Mirror & Rotate 270")
	TransformInherit.Register("Inherit")

	CompositeAlphaOpaque.Register("Opaque")
	CompositeAlphaPreMultiplied.Register("Pre-Multiplied")
	CompositeAlphaPostMultiplied.Register("Post-Multiplied")
	CompositeAlphaInherit.Register("Inherited")

	PresentModeImmediate.Register("Immediate")
	PresentModeMailbox.Register("Mailbox")
	PresentModeFIFO.Register("FIFO")
	PresentModeFIFORelaxed.Register("FIFO Relaxed")

	ColorSpaceSRGBNonlinear.Register("sRGB Non-Linear")

	VKErrorSurfaceLost.Register("surface lost")
	VKErrorNativeWindowInUse.Register("native window in use")
}
