package ext_host_query_reset_shim

import (
	"github.com/golang/mock/gomock"
	core_mocks "github.com/vkngwrapper/core/v2/mocks"
	mock_host_query_reset "github.com/vkngwrapper/extensions/v2/ext_host_query_reset/mocks"
	"testing"
)

func TestVulkanShim_Reset(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	queryPool := core_mocks.NewMockQueryPool(ctrl)
	extension := mock_host_query_reset.NewMockExtension(ctrl)
	shim := NewShim(extension, queryPool)

	extension.EXPECT().ResetQueryPool(queryPool, 1, 3)
	shim.Reset(1, 3)
}
