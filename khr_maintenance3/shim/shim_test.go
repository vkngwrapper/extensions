package khr_maintenance3_shim

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v2/core1_0"
	"github.com/vkngwrapper/core/v2/core1_1"
	core_mocks "github.com/vkngwrapper/core/v2/mocks"
	"github.com/vkngwrapper/extensions/v3/khr_maintenance3"
	mock_maintenance3 "github.com/vkngwrapper/extensions/v3/khr_maintenance3/mocks"
	"go.uber.org/mock/gomock"
)

func TestVulkanShim_DescriptorSetLayoutSupport(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extension := mock_maintenance3.NewMockExtension(ctrl)
	device := core_mocks.NewMockDevice(ctrl)
	shim := NewShim(extension, device)

	extension.EXPECT().DescriptorSetLayoutSupport(
		device,
		core1_0.DescriptorSetLayoutCreateInfo{
			Bindings: []core1_0.DescriptorSetLayoutBinding{
				{
					Binding: 3,
				},
				{
					Binding: 5,
				},
			},
		},
		gomock.Any()).DoAndReturn(func(device core1_0.Device, setLayoutOptions core1_0.DescriptorSetLayoutCreateInfo, outData *khr_maintenance3.DescriptorSetLayoutSupport) error {
		outData.Supported = true
		return nil
	})

	var out core1_1.DescriptorSetLayoutSupport
	err := shim.DescriptorSetLayoutSupport(core1_0.DescriptorSetLayoutCreateInfo{
		Bindings: []core1_0.DescriptorSetLayoutBinding{
			{
				Binding: 3,
			},
			{
				Binding: 5,
			},
		},
	}, &out)
	require.NoError(t, err)
	require.Equal(t, core1_1.DescriptorSetLayoutSupport{
		Supported: true,
	}, out)
}
