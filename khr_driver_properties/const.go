package khr_driver_properties

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"

// DriverID specifies khronos driver id's
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDriverId.html
type DriverID int32

var driverIDMapping = make(map[DriverID]string)

func (e DriverID) Register(str string) {
	driverIDMapping[e] = str
}

func (e DriverID) String() string {
	return driverIDMapping[e]
}

////

const (
	// ExtensionName is "VK_KHR_driver_properties"
	ExtensionName string = C.VK_KHR_DRIVER_PROPERTIES_EXTENSION_NAME

	// MaxDriverInfoSize is the length of a PhysicalDevice driver information string
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_MAX_DRIVER_INFO_SIZE.html
	MaxDriverInfoSize int = C.VK_MAX_DRIVER_INFO_SIZE_KHR
	// MaxDriverNameSize is the maximum length of a PhysicalDevice driver name string
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_MAX_DRIVER_NAME_SIZE.html
	MaxDriverNameSize int = C.VK_MAX_DRIVER_NAME_SIZE_KHR

	// DriverIDAmdOpenSource indicates open-source AMD drivers
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDriverId.html
	DriverIDAmdOpenSource DriverID = C.VK_DRIVER_ID_AMD_OPEN_SOURCE_KHR
	// DriverIDAmdProprietary indicates proprietary AMD drivers
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDriverId.html
	DriverIDAmdProprietary DriverID = C.VK_DRIVER_ID_AMD_PROPRIETARY_KHR
	// DriverIDArmProprietary indicates proprietary ARM drivers
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDriverId.html
	DriverIDArmProprietary DriverID = C.VK_DRIVER_ID_ARM_PROPRIETARY_KHR
	// DriverIDBroadcomProprietary indicates proprietary Broadcom drivers
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDriverId.html
	DriverIDBroadcomProprietary DriverID = C.VK_DRIVER_ID_BROADCOM_PROPRIETARY_KHR
	// DriverIDGgpProprietary indicates proprietary GGP drivers
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDriverId.html
	DriverIDGgpProprietary DriverID = C.VK_DRIVER_ID_GGP_PROPRIETARY_KHR
	// DriverIDGoogleSwiftshader indicates Google Swiftshader drivers
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDriverId.html
	DriverIDGoogleSwiftshader DriverID = C.VK_DRIVER_ID_GOOGLE_SWIFTSHADER_KHR
	// DriverIDImaginationProprietary indicates proprietary Imagination drivers
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDriverId.html
	DriverIDImaginationProprietary DriverID = C.VK_DRIVER_ID_IMAGINATION_PROPRIETARY_KHR
	// DriverIDIntelOpenSourceMesa indicates open-source Mesa drivers
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDriverId.html
	DriverIDIntelOpenSourceMesa DriverID = C.VK_DRIVER_ID_INTEL_OPEN_SOURCE_MESA_KHR
	// DriverIDIntelProprietaryWindows indicates proprietary Intel drivers for Windows
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDriverId.html
	DriverIDIntelProprietaryWindows DriverID = C.VK_DRIVER_ID_INTEL_PROPRIETARY_WINDOWS_KHR
	// DriverIDMesaRadV indicates Mesa Rad-V drivers
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDriverId.html
	DriverIDMesaRadV DriverID = C.VK_DRIVER_ID_MESA_RADV_KHR
	// DriverIDNvidiaProprietary indicates proprietary NVidia drivers
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDriverId.html
	DriverIDNvidiaProprietary DriverID = C.VK_DRIVER_ID_NVIDIA_PROPRIETARY_KHR
	// DriverIDQualcommProprietary indicates proprietary Qualcomm drivers
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDriverId.html
	DriverIDQualcommProprietary DriverID = C.VK_DRIVER_ID_QUALCOMM_PROPRIETARY_KHR
)

func init() {
	DriverIDAmdOpenSource.Register("AMD Open-Source")
	DriverIDAmdProprietary.Register("AMD Proprietary")
	DriverIDArmProprietary.Register("ARM Proprietary")
	DriverIDBroadcomProprietary.Register("Broadcom Proprietary")
	DriverIDGgpProprietary.Register("GGP Proprietary")
	DriverIDGoogleSwiftshader.Register("Google Swiftshader")
	DriverIDImaginationProprietary.Register("Imagination Proprietary")
	DriverIDIntelOpenSourceMesa.Register("Intel Open-Source (Mesa)")
	DriverIDIntelProprietaryWindows.Register("Intel Proprietary (Windows)")
	DriverIDMesaRadV.Register("Mesa RADV")
	DriverIDNvidiaProprietary.Register("Nvidia Proprietary")
	DriverIDQualcommProprietary.Register("Qualcomm Proprietary")
}
