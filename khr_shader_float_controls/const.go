package khr_shader_float_controls

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import _ "github.com/vkngwrapper/extensions/vulkan"

// ShaderFloatControlsIndependence specifies whether, and how, shader float controls
// can be set separately
type ShaderFloatControlsIndependence int32

var shaderFloatControlsIndependenceMapping = make(map[ShaderFloatControlsIndependence]string)

func (e ShaderFloatControlsIndependence) Register(str string) {
	shaderFloatControlsIndependenceMapping[e] = str
}

func (e ShaderFloatControlsIndependence) String() string {
	return shaderFloatControlsIndependenceMapping[e]
}

////

const (
	// ExtensionName is "VK_KHR_shader_float_controls"
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_shader_float_controls.html
	ExtensionName string = C.VK_KHR_SHADER_FLOAT_CONTROLS_EXTENSION_NAME

	// ShaderFloatControlsIndependence32BitOnly specifies that shader float controls for 32-bit
	// floating point can be set independently
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderFloatControlsIndependence.html
	ShaderFloatControlsIndependence32BitOnly ShaderFloatControlsIndependence = C.VK_SHADER_FLOAT_CONTROLS_INDEPENDENCE_32_BIT_ONLY_KHR
	// ShaderFloatControlsIndependenceAll specifies that shader float controls for all
	// bit widths can be set independently
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderFloatControlsIndependence.html
	ShaderFloatControlsIndependenceAll ShaderFloatControlsIndependence = C.VK_SHADER_FLOAT_CONTROLS_INDEPENDENCE_ALL_KHR
	// ShaderFloatControlsIndependenceNone specifies that shader float controls for all bit widths
	// must be set identically
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderFloatControlsIndependence.html
	ShaderFloatControlsIndependenceNone ShaderFloatControlsIndependence = C.VK_SHADER_FLOAT_CONTROLS_INDEPENDENCE_NONE_KHR
)

func init() {
	ShaderFloatControlsIndependenceAll.Register("All")
	ShaderFloatControlsIndependenceNone.Register("None")
	ShaderFloatControlsIndependence32BitOnly.Register("32-Bit Only")
}
