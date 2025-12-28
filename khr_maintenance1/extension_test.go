package khr_maintenance1_test

import (
	"testing"

	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/mocks"
	"github.com/vkngwrapper/extensions/v3/khr_maintenance1"
	khr_maintenance1_driver "github.com/vkngwrapper/extensions/v3/khr_maintenance1/loader"
	mock_maintenance1 "github.com/vkngwrapper/extensions/v3/khr_maintenance1/mocks"
	"go.uber.org/mock/gomock"
)

func TestVulkanExtension_TrimCommandPool(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	device := mocks.NewDummyDevice(common.Vulkan1_0, []string{})
	commandPool := mocks.NewDummyCommandPool(device)

	maintDriver := mock_maintenance1.NewMockLoader(ctrl)
	extension := khr_maintenance1.CreateExtensionDriverFromLoader(maintDriver)

	maintDriver.EXPECT().VkTrimCommandPoolKHR(device.Handle(), commandPool.Handle(), khr_maintenance1_driver.VkCommandPoolTrimFlagsKHR(0))

	extension.TrimCommandPool(commandPool, 0)
}
