package khr_acceleration_structure

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import "unsafe"

type DeviceOrHostAddressConst interface {
	PopulateAddressUnion(v unsafe.Pointer)
}

type DeviceAddressConst struct {
	DeviceAddress uint64
}

func (a DeviceAddressConst) PopulateAddressUnion(v unsafe.Pointer) {
	deviceAddress := (*C.VkDeviceAddress)(v)
	*deviceAddress = C.VkDeviceAddress(a.DeviceAddress)
}

type HostAddressConst struct {
	HostAddress unsafe.Pointer
}

func (a HostAddressConst) PopulateAddressUnion(v unsafe.Pointer) {
	hostAddress := (*unsafe.Pointer)(v)
	*hostAddress = a.HostAddress
}
