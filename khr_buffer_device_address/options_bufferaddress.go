package khr_buffer_device_address

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/vkngwrapper/core/v3/core1_2"
)

// BufferDeviceAddressInfo specifies the Buffer to query an address for
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferDeviceAddressInfo.html
type BufferDeviceAddressInfo = core1_2.BufferDeviceAddressInfo
