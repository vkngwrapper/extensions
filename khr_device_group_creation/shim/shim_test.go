package khr_device_group_creation_shim

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/core1_1"
	core_mocks "github.com/vkngwrapper/core/v3/mocks/mocks1_0"
	"github.com/vkngwrapper/extensions/v3/khr_device_group_creation"
	mock_device_group_creation "github.com/vkngwrapper/extensions/v3/khr_device_group_creation/mocks"
	"go.uber.org/mock/gomock"
)

func TestVulkanShim_EnumeratePhysicalDeviceGroups(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extension := mock_device_group_creation.NewMockExtension(ctrl)
	instance := core_mocks.NewMockInstance(ctrl)
	shim := NewShim(extension, instance)

	device1 := core_mocks.NewMockPhysicalDevice(ctrl)
	device2 := core_mocks.NewMockPhysicalDevice(ctrl)
	device3 := core_mocks.NewMockPhysicalDevice(ctrl)

	extension.EXPECT().EnumeratePhysicalDeviceGroups(instance, gomock.Any()).Return(
		[]*khr_device_group_creation.PhysicalDeviceGroupProperties{
			{
				PhysicalDevices:  []core1_0.PhysicalDevice{device1, device2},
				SubsetAllocation: true,
			},
			{
				PhysicalDevices: []core1_0.PhysicalDevice{device3},
			},
		}, core1_0.VKSuccess, nil)

	properties, _, err := shim.EnumeratePhysicalDeviceGroups(func() *core1_1.PhysicalDeviceGroupProperties {
		return nil
	})
	require.NoError(t, err)
	require.Len(t, properties, 2)
}
