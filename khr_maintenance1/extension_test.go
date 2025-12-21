package khr_maintenance1_test

import (
	"testing"

	"github.com/vkngwrapper/core/v3/common"
	mock_driver "github.com/vkngwrapper/core/v3/driver/mocks"
	"github.com/vkngwrapper/core/v3/mocks/mocks1_0"
	"github.com/vkngwrapper/extensions/v3/khr_maintenance1"
	khr_maintenance1_driver "github.com/vkngwrapper/extensions/v3/khr_maintenance1/driver"
	mock_maintenance1 "github.com/vkngwrapper/extensions/v3/khr_maintenance1/mocks"
	"go.uber.org/mock/gomock"
)

func TestVulkanExtension_TrimCommandPool(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := mocks1_0.EasyMockDevice(ctrl, coreDriver)
	commandPool := mocks1_0.EasyMockCommandPool(ctrl, device)

	maintDriver := mock_maintenance1.NewMockDriver(ctrl)
	extension := khr_maintenance1.CreateExtensionFromDriver(maintDriver)

	maintDriver.EXPECT().VkTrimCommandPoolKHR(device.Handle(), commandPool.Handle(), khr_maintenance1_driver.VkCommandPoolTrimFlagsKHR(0))

	extension.TrimCommandPool(commandPool, 0)
}
