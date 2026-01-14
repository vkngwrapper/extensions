package mock_acceleration_structure

import (
	"math/rand"
	"unsafe"

	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/extensions/v3/khr_acceleration_structure"
	khr_acceleration_structure_loader "github.com/vkngwrapper/extensions/v3/khr_acceleration_structure/loader"
)

func NewDummyAccelerationStructure(device core1_0.Device) khr_acceleration_structure.AccelerationStructure {
	return khr_acceleration_structure.InternalAccelerationStructure(device.Handle(), NewFakeAccelerationStructure(), device.APIVersion())
}

func fakePointer() unsafe.Pointer {
	return unsafe.Pointer(uintptr(rand.Int()))
}

func NewFakeAccelerationStructure() khr_acceleration_structure_loader.VkAccelerationStructureKHR {
	return khr_acceleration_structure_loader.VkAccelerationStructureKHR(fakePointer())
}
