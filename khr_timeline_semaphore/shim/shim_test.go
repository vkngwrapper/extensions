package khr_timeline_semaphore_shim

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/core1_2"
	core_mocks "github.com/vkngwrapper/core/v3/mocks/mocks1_0"
	"github.com/vkngwrapper/extensions/v3/khr_timeline_semaphore"
	mock_timeline_semaphore "github.com/vkngwrapper/extensions/v3/khr_timeline_semaphore/mocks"
	"go.uber.org/mock/gomock"
)

func TestVulkanSemaphoreShim_CounterValue(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extension := mock_timeline_semaphore.NewMockExtension(ctrl)
	semaphore := core_mocks.NewMockSemaphore(ctrl)
	shim := NewSemaphoreShim(extension, semaphore)

	extension.EXPECT().SemaphoreCounterValue(semaphore).Return(uint64(3), core1_0.VKSuccess, nil)

	val, _, err := shim.CounterValue()
	require.NoError(t, err)
	require.Equal(t, uint64(3), val)
}

func TestVulkanDeviceShim_SignalSemaphore(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extension := mock_timeline_semaphore.NewMockExtension(ctrl)
	device := core_mocks.NewMockDevice(ctrl)
	shim := NewDeviceShim(extension, device)

	semaphore := core_mocks.NewMockSemaphore(ctrl)

	extension.EXPECT().SignalSemaphore(
		device,
		khr_timeline_semaphore.SemaphoreSignalInfo{
			Semaphore: semaphore,
			Value:     3,
		}).Return(core1_0.VKSuccess, nil)

	_, err := shim.SignalSemaphore(core1_2.SemaphoreSignalInfo{
		Semaphore: semaphore,
		Value:     3,
	})
	require.NoError(t, err)
}

func TestVulkanDeviceShim_WaitSemaphores(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extension := mock_timeline_semaphore.NewMockExtension(ctrl)
	device := core_mocks.NewMockDevice(ctrl)
	shim := NewDeviceShim(extension, device)

	semaphore1 := core_mocks.NewMockSemaphore(ctrl)
	semaphore2 := core_mocks.NewMockSemaphore(ctrl)

	extension.EXPECT().WaitSemaphores(
		device,
		time.Minute,
		khr_timeline_semaphore.SemaphoreWaitInfo{
			Flags:      khr_timeline_semaphore.SemaphoreWaitAny,
			Semaphores: []core1_0.Semaphore{semaphore1, semaphore2},
			Values:     []uint64{3, 5},
		}).Return(core1_0.VKSuccess, nil)

	_, err := shim.WaitSemaphores(
		time.Minute,
		core1_2.SemaphoreWaitInfo{
			Flags:      core1_2.SemaphoreWaitAny,
			Semaphores: []core1_0.Semaphore{semaphore1, semaphore2},
			Values:     []uint64{3, 5},
		})
	require.NoError(t, err)
}
