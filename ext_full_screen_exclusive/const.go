//go:build windows

package ext_full_screen_exclusive

/*
#define VK_USE_PLATFORM_WIN32_KHR
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import "github.com/vkngwrapper/core/v3/common"

const (
	// ExtensionName is "VK_EXT_full_screen_exclusive"
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VK_EXT_full_screen_exclusive.html
	ExtensionName string = C.VK_EXT_FULL_SCREEN_EXCLUSIVE_EXTENSION_NAME

	// VkErrorFullScreenExclusiveModeLost indicates that an operation on a swapchain created
	// with VK_FULL_SCREEN_EXCLUSIVE_APPLICATION_CONTROLLED_EXT failed as it did not have
	// exclusive full-screen access. This may occur due to implementation-dependent reasons,
	// outside of the application’s control.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkResult.html#
	VkErrorFullScreenExclusiveModeLost common.VkResult = C.VK_ERROR_FULL_SCREEN_EXCLUSIVE_MODE_LOST_EXT
)

type FullScreenExclusive int

var fullScreenExclusiveMapping = make(map[FullScreenExclusive]string)

func (e FullScreenExclusive) Register(str string) {
	fullScreenExclusiveMapping[e] = str
}

func (e FullScreenExclusive) String() string {
	return fullScreenExclusiveMapping[e]
}

const (
	// FullScreenExclusiveDefault specifies that the implementation should determine the
	// appropriate full-screen method by whatever means it deems appropriate.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkFullScreenExclusiveEXT.html#
	FullScreenExclusiveDefault FullScreenExclusive = C.VK_FULL_SCREEN_EXCLUSIVE_DEFAULT_EXT
	// FullScreenExclusiveAllowed specifies that the implementation may use full-screen
	// exclusive mechanisms when available. Such mechanisms may result in better performance
	// and/or the availability of different presentation capabilities, but may require a more
	// disruptive transition during swapchain initialization, first presentation and/or
	// destruction.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkFullScreenExclusiveEXT.html#
	FullScreenExclusiveAllowed FullScreenExclusive = C.VK_FULL_SCREEN_EXCLUSIVE_ALLOWED_EXT
	// FullScreenExclusiveDisallowed specifies that the implementation should avoid using
	// full-screen mechanisms which rely on disruptive transitions.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkFullScreenExclusiveEXT.html#
	FullScreenExclusiveDisallowed FullScreenExclusive = C.VK_FULL_SCREEN_EXCLUSIVE_DISALLOWED_EXT
	// FullScreenExclusiveApplicationControlled specifies that the application will manage
	// full-screen exclusive mode by using the vkAcquireFullScreenExclusiveModeEXT and
	// vkReleaseFullScreenExclusiveModeEXT commands.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkFullScreenExclusiveEXT.html#
	FullScreenExclusiveApplicationControlled FullScreenExclusive = C.VK_FULL_SCREEN_EXCLUSIVE_APPLICATION_CONTROLLED_EXT
)

func init() {
	FullScreenExclusiveDefault.Register("Default")
	FullScreenExclusiveAllowed.Register("Allowed")
	FullScreenExclusiveDisallowed.Register("Disallowed")
	FullScreenExclusiveApplicationControlled.Register("ApplicationControlled")

	VkErrorFullScreenExclusiveModeLost.Register("exclusive mode lost")
}
