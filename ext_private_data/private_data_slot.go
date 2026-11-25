package ext_private_data

import (
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/loader"
	ext_private_data_loader "github.com/vkngwrapper/extensions/v3/ext_private_data/loader"
)

// PrivateDataSlot is an opaque representation of a private data slot object
//
// https://docs.vulkan.org/refpages/latest/refpages/source/VkPrivateDataSlot.html
type PrivateDataSlot struct {
	device                loader.VkDevice
	privateDataSlotHandle ext_private_data_loader.VkPrivateDataSlot

	apiVersion common.APIVersion
}

func (s PrivateDataSlot) Handle() ext_private_data_loader.VkPrivateDataSlot {
	return s.privateDataSlotHandle
}

func (s PrivateDataSlot) DeviceHandle() loader.VkDevice {
	return s.device
}

func (s PrivateDataSlot) APIVersion() common.APIVersion {
	return s.apiVersion
}

func (s PrivateDataSlot) Initialized() bool {
	return s.privateDataSlotHandle != 0
}

func InternalPrivateDataSlot(device loader.VkDevice, handle ext_private_data_loader.VkPrivateDataSlot, version common.APIVersion) PrivateDataSlot {
	return PrivateDataSlot{
		device:                device,
		privateDataSlotHandle: handle,
		apiVersion:            version,
	}
}
