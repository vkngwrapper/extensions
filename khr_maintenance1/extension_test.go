package khr_maintenance1_test

import (
	"github.com/golang/mock/gomock"
	"github.com/vkngwrapper/core/v2/common"
	mock_driver "github.com/vkngwrapper/core/v2/driver/mocks"
	"github.com/vkngwrapper/core/v2/mocks"
	"github.com/vkngwrapper/extensions/v2/khr_maintenance1"
	khr_maintenance1_driver "github.com/vkngwrapper/extensions/v2/khr_maintenance1/driver"
	mock_maintenance1 "github.com/vkngwrapper/extensions/v2/khr_maintenance1/mocks"
	"testing"
)

func TestVulkanExtension_TrimCommandPool(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := mocks.EasyMockDevice(ctrl, coreDriver)
	commandPool := mocks.EasyMockCommandPool(ctrl, device)

	maintDriver := mock_maintenance1.NewMockDriver(ctrl)
	extension := khr_maintenance1.CreateExtensionFromDriver(maintDriver)

	maintDriver.EXPECT().VkTrimCommandPoolKHR(device.Handle(), commandPool.Handle(), khr_maintenance1_driver.VkCommandPoolTrimFlagsKHR(0))

	extension.TrimCommandPool(commandPool, 0)
}
