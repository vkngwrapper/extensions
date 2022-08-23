package khr_external_fence_capabilities

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/cockroachdb/errors"
	"github.com/google/uuid"
	"github.com/vkngwrapper/core/v2/common"
	"unsafe"
)

// PhysicalDeviceIDProperties speicifes IDs related to the PhysicalDevice
type PhysicalDeviceIDProperties struct {
	// DeviceUUID represents a universally-unique identifier for the device
	DeviceUUID uuid.UUID
	// DriverUUID represents a universally-unique identifier for the driver build
	// in use by the device
	DriverUUID uuid.UUID
	// DeviceLUID represents a locally-unique identifier for the device
	DeviceLUID uint64
	// DeviceNodeMask identifies the node within a linked device adapter corresponding to the
	// Device
	DeviceNodeMask uint32
	// DeviceLUIDValid is true if DeviceLUID contains a valid LUID and DeviceNodeMask contains
	// a valid node mask
	DeviceLUIDValid bool

	common.NextOutData
}

func (o *PhysicalDeviceIDProperties) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceIDPropertiesKHR{})))
	}
	info := (*C.VkPhysicalDeviceIDPropertiesKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_ID_PROPERTIES_KHR
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceIDProperties) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceIDPropertiesKHR)(cDataPointer)

	deviceUUIDBytes := C.GoBytes(unsafe.Pointer(&info.deviceUUID[0]), C.VK_UUID_SIZE)
	o.DeviceUUID, err = uuid.FromBytes(deviceUUIDBytes)
	if err != nil {
		return nil, errors.Wrap(err, "vulkan provided invalid device uuid")
	}

	driverUUIDBytes := C.GoBytes(unsafe.Pointer(&info.driverUUID[0]), C.VK_UUID_SIZE)
	o.DriverUUID, err = uuid.FromBytes(driverUUIDBytes)
	if err != nil {
		return nil, errors.Wrap(err, "vulkan provided invalid driver uuid")
	}

	o.DeviceLUID = *(*uint64)(unsafe.Pointer(&info.deviceLUID[0]))
	o.DeviceNodeMask = uint32(info.deviceNodeMask)
	o.DeviceLUIDValid = info.deviceLUIDValid != C.VkBool32(0)

	return info.pNext, nil
}
