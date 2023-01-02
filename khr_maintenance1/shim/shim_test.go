package khr_maintenance1_shim

import (
	"github.com/golang/mock/gomock"
	core_mocks "github.com/vkngwrapper/core/v2/mocks"
	"github.com/vkngwrapper/extensions/v2/khr_maintenance1"
	mock_maintenance1 "github.com/vkngwrapper/extensions/v2/khr_maintenance1/mocks"
	"testing"
)

func TestVulkanShim_TrimCommandPool(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extension := mock_maintenance1.NewMockExtension(ctrl)
	commandPool := core_mocks.NewMockCommandPool(ctrl)
	shim := NewShim(extension, commandPool)

	extension.EXPECT().TrimCommandPool(
		commandPool,
		khr_maintenance1.CommandPoolTrimFlags(0),
	)

	shim.TrimCommandPool(0)
}
