package ext_private_data

import (
	"errors"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
	ext_private_data_loader "github.com/vkngwrapper/extensions/v3/ext_private_data/loader"
)

type VulkanExtensionDriver struct {
	device core1_0.Device
	driver ext_private_data_loader.Loader
}

// CreateExtensionDriverFromCoreDriver produces an ExtensionDriver object from a Device with
// ext_extended_dynamic_state2 loaded
func CreateExtensionDriverFromCoreDriver(coreDriver core1_0.DeviceDriver) ExtensionDriver {
	device := coreDriver.Device()

	if !device.IsDeviceExtensionActive(ExtensionName) {
		return nil
	}

	return CreateExtensionDriverFromLoader(ext_private_data_loader.CreateLoaderFromCore(coreDriver.Loader()), device)
}

// CreateExtensionDriverFromLoader generates an ExtensionDriver from a loader.Loader object- this is usually
// used in tests to build an ExtensionDriver from mock drivers
func CreateExtensionDriverFromLoader(driver ext_private_data_loader.Loader, device core1_0.Device) *VulkanExtensionDriver {
	return &VulkanExtensionDriver{
		device: device,
		driver: driver,
	}
}

func (v *VulkanExtensionDriver) CreatePrivateDataSlot(allocationCallbacks *loader.AllocationCallbacks, options PrivateDataSlotCreateInfo) (PrivateDataSlot, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	optionPtr, err := common.AllocOptions(arena, options)
	if err != nil {
		return PrivateDataSlot{}, core1_0.VKErrorUnknown, err
	}

	var privateDataSlotHandle ext_private_data_loader.VkPrivateDataSlot
	res, err := v.driver.VkCreatePrivateDataSlotEXT(v.device.Handle(), (*ext_private_data_loader.VkPrivateDataSlotCreateInfoEXT)(optionPtr), allocationCallbacks.Handle(), &privateDataSlotHandle)
	if err != nil {
		return PrivateDataSlot{}, res, err
	}

	return InternalPrivateDataSlot(v.device.Handle(), privateDataSlotHandle, v.device.APIVersion()), res, err
}

func (v *VulkanExtensionDriver) DestroyPrivateDataSlot(privateDataSlot PrivateDataSlot, allocator *loader.AllocationCallbacks) {
	v.driver.VkDestroyPrivateDataSlotEXT(v.device.Handle(), privateDataSlot.Handle(), allocator.Handle())
}

func (v *VulkanExtensionDriver) GetPrivateData(objectType core1_0.ObjectType, handle loader.VulkanHandle, privateDataSlot PrivateDataSlot) uint64 {
	if !privateDataSlot.Initialized() {
		panic("privateDataSlot cannot be left uninitialized")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	dataPtr := arena.Malloc(8)

	v.driver.VkGetPrivateDataEXT(v.device.Handle(), loader.VkObjectType(objectType), loader.Uint64(handle), privateDataSlot.Handle(), (*loader.Uint64)(dataPtr))

	return *(*uint64)(dataPtr)
}

func (v *VulkanExtensionDriver) SetPrivateData(objectType core1_0.ObjectType, handle loader.VulkanHandle, privateDataSlot PrivateDataSlot, data uint64) (common.VkResult, error) {
	if !privateDataSlot.Initialized() {
		return core1_0.VKErrorUnknown, errors.New("privateDataSlot cannot be left uninitialized")
	}

	return v.driver.VkSetPrivateDataEXT(v.device.Handle(), loader.VkObjectType(objectType), loader.Uint64(handle), privateDataSlot.Handle(), loader.Uint64(data))
}
